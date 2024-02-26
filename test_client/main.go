/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Modified by Stony Brook University students
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
package main

import (
	"cmp"
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"slices"
	"strconv"

	pb "orcanet/market"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

type clientHandle struct {
	stream pb.MarketService_RequestQueryClient
	user   *pb.User
}

func main() {

	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer conn.Close()
	c := pb.NewMarketServiceClient(conn)

	// Prompt for username in terminal
	var username string
	fmt.Print("Enter username: ")
	fmt.Scanln(&username)

	// Generate a random ID for new user
	userID := fmt.Sprintf("user%d", rand.Intn(10000))

	fmt.Print("Enter a price for supplying files: ")
	var price int32
	_, err = fmt.Scanln(&price)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Make random port and assign it to registered user
	random_port := rand.Intn(65535-49152+1) + 49152

	user := &pb.User{
		Id:    userID,
		Name:  username,
		Ip:    "localhost",
		Port:  int32(random_port),
		Price: price,
	}

	// Create a stream and send it to the client handler
	stream, err := c.RequestQuery(context.Background())
	if err != nil {
		log.Fatalf("Failed to call %v", err)
	}
	ch := clientHandle{stream: stream, user: user}

	// We really need to think about how to peoerply synchronize this 
	// Create channels
	holdersch := make(chan *pb.Holders)
	registrch := make(chan string, 1)

	// Immediately start a goroutine for this (will probably need to handle this better)
	// Now that I think about it, is this okay? 
	go ch.getMessages(holdersch, registrch)

	// Print out list of options 
	for {
		fmt.Println("---------------------------------")
		fmt.Println("1. Request a file")
		fmt.Println("2. Register a file")
		fmt.Println("3. Check holders for a file")
		fmt.Println("4. Exit")
		fmt.Print("Option: ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		if choice == 4 {
			return
		}

		fmt.Print("Enter a file hash: ")
		var fileHash string
		_, err = fmt.Scanln(&fileHash)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		switch choice {
			case 1:
				ch.createRequest(fileHash)
			case 2:
				ch.registerRequest(fileHash, registrch)
			case 3:
				ch.checkHolders(fileHash, holdersch)
			case 4:
				return
			default:
				fmt.Println("Unknown option: ", choice)
		}
		fmt.Println()
	}
}

// Will need to finish this later 
func (ch *clientHandle) createRequest(fileHash string) {

	/* err := ch.stream.Send(&pb.Request{RequestType: &pb.Request_Filerequest{Filerequest: &pb.FileRequest{User: ch.user, FileHash: fileHash}}})
	if err != nil {
		log.Printf("Error while sending message to server :: %v", err)
	} */

	// Might need to do more stuff here
}

// register a file 
func (ch *clientHandle) registerRequest(fileHash string, registrch <-chan string) {

	err := ch.stream.Send(&pb.Request{RequestType: &pb.Request_Supplyfile{Supplyfile: &pb.SupplyFile{User: ch.user, FileHash: fileHash}}})
	if err != nil {
		log.Printf("Error while sending message to server :: %v", err)
	}

	// Let's try the method we saw 
	response := <- registrch // This will hang until we hae read something 
	fmt.Println(response)
}

// check for holders then make a subsequent query to request for a file
func (ch *clientHandle) checkHolders(fileHash string, holdersch chan *pb.Holders) {

	err := ch.stream.Send(&pb.Request{RequestType: &pb.Request_Checkholder{Checkholder: &pb.CheckHolder{FileHash: fileHash}}})
	if err != nil {
		log.Printf("Error while sending message  to server :: %v", err)
	}

	// Will eventually need a case for err channel so we can break out of the loop if an error occurs
	for {
		select {
			// Somehow the response is never received but it was fine before?? how come?? 
			case response := <-holdersch:

				supply_files := response.GetHolders()

				slices.SortFunc(supply_files, func(a, b *pb.SupplyFile) int {
					return cmp.Compare(a.GetUser().GetPrice(), b.GetUser().GetPrice())
				})

				for idx, holder := range supply_files {
					user := holder.GetUser()
					fmt.Printf("(%d) Username: %s, Price: %d\n", idx, user.GetName(), user.GetPrice())
				}
				fmt.Println("Choose which supplier to get file from, or 'n' to cancel:")
				var choice string
				_, err = fmt.Scanln(&choice)
				if err != nil {
					fmt.Println("Error: ", err)
					return
				}
				idx, err := strconv.ParseInt(choice, 10, 32)
				if err != nil {
					return
				}
				if idx < 0 || int(idx) > len(supply_files) {
					fmt.Println("Invalid index chosen")
					return
				}
				fmt.Printf("%v chosen, requesting file\n", idx)

				// Get the file and the user with it
				// May need to modify this so we will have to see
				/*user := supply_files[idx].GetUser()
				err = ch.stream.Send(&pb.Request{RequestType: &pb.Request_Filerequest{Filerequest: &pb.FileRequest{User: user, FileHash: fileHash}}})
				if err != nil {
					log.Fatalf("Error %v", err)
				}  */
				return 
		}
	}
}

// Goroutine that constantly reads messages and writes into respective channels 
func (ch *clientHandle) getMessages(holdersch chan <- *pb.Holders, registrch chan <- string) {

	for {
		// This blocks until a message is received
		mssg, err := ch.stream.Recv() 
		if err != nil {
			log.Printf("Error in receiving message from server :: %v", err)
		}

		// Determine the message type 
		switch mssg.ResponseType.(type) {
			case *pb.Response_Fileresponse:
				// We will do something here later
			case *pb.Response_Holders:
				// Write into the holders channel 
				supply_files := mssg.GetHolders()
				holdersch <- supply_files
			case *pb.Response_Message:
				// Write into the register channel
				message := mssg.GetMessage()
				registrch <- message
		}
	}
}
