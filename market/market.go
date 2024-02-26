package market

import (
	"fmt"
	"log"
	"sync"
)

type messageUnit struct {	
	Request  int32
	Payload  interface{}
}

type messageHandle struct {
	messageQueue []messageUnit
	mu           sync.Mutex
}

// Handler for our queue and mutex
var messageHandleObject = messageHandle{}

// Must use the Unimplemented Service Server in order to work
type MarketServer struct {
	UnimplementedMarketServiceServer
}

type CustomData struct {
	Type int32 // We will use an id for now (I am lazy)
	Data interface{}
}


// Hold the list of files globally (should this be globablly or should i put in the RequestQuery function)
var files = make(map[string][]*SupplyFile)



// Printing the map of the holders on the server side
func printHoldersMap() {
	for hash, holders := range files {
		fmt.Printf("\nFile Hash: %s\n", hash)

		for _, holder := range holders {
			user := holder.GetUser()
			fmt.Printf("Username: %s, Price: %d\n", user.GetName(), user.GetPrice())
		}

	}
}


func (is *MarketServer) RequestQuery(csi MarketService_RequestQueryServer) error {

	errch := make(chan error)

	// Goroutines for receiving and sending
	go receiveRequests(csi, errch)
	go sendResponses(csi, errch)

	return <-errch

}

// Receive Messages regarding File Requests
func receiveRequests(csi_ MarketService_RequestQueryServer, errch chan error) error {

	// Keep receiving from the stream
	for {

		req, err := csi_.Recv()

		if err != nil {
			log.Printf("Error in receiving message from client :: %v", err)
			errch <- err
		} else {

			// lock thread 
			messageHandleObject.mu.Lock()

			switch req.RequestType.(type) {
			// Request a file
			case *Request_Filerequest:
				// File look up
				fileHash := req.GetFilerequest().GetFileHash()

				// I hate Go
				if _, ok := files[fileHash]; !ok {

					// Come back to this later 

					/*payload := make(map[string][]*User)
					payload[fileHash] = append(payload[fileHash], source_user)
					payload[fileHash] = append(payload[fileHash], target_user)

					// Add tp the queue
					messageHandleObject.messageQueue = append(messageHandleObject.messageQueue, messageUnit{
						Request: 0,
						Payload: payload,
					}) */
				}
			// Checking for holders
			case *Request_Checkholder:

				fileHash := req.GetCheckholder().GetFileHash()

				// Get holders
				holders := files[fileHash]
				printHoldersMap()

				// Add to the queue
				messageHandleObject.messageQueue = append(messageHandleObject.messageQueue, messageUnit{
					Request: 1,
					Payload: holders,
				})

			// Registering a User
			case *Request_Supplyfile:

				user := req.GetSupplyfile().GetUser()
				fileHash := req.GetSupplyfile().GetFileHash()

				// Register user 
				files[fileHash] = append(files[fileHash], req.GetSupplyfile())

				// Add to queue 
				messageHandleObject.messageQueue = append(messageHandleObject.messageQueue, messageUnit{
					Request:  2,
					Payload:  user,
				})
			}

			// Unlock thread after done 
			messageHandleObject.mu.Unlock()
		}
	}
}

// I think this function is compeltely awful and doesnt utilize synchronization properly. I definitly want to fix this later. 
// Send Messages regarding File Requests
func sendResponses(csi_ MarketService_RequestQueryServer, errch chan error) error {

	for {

		// Iterate through our message queues
		for {

			messageHandleObject.mu.Lock()
			
			if len(messageHandleObject.messageQueue) == 0 {
				messageHandleObject.mu.Unlock()
				break
			}

			// Get the information from the mesasge queue
			request := messageHandleObject.messageQueue[0].Request
			payload := messageHandleObject.messageQueue[0].Payload

		
			switch request {
				case 0:
					// Code for requesting file
				case 1:
					// Getting file holders and send as a strema 
					holders := payload.([]*SupplyFile)
					err := csi_.Send(&Response{ResponseType: &Response_Holders{Holders: &Holders{Holders: holders}}})
					if err != nil {
						errch <- err 
					}
				case 2:
					// registering a file and send as a stream 
					message := fmt.Sprintf("File has been registered!")
					err := csi_.Send(&Response{ResponseType: &Response_Message{Message: message}})
					if err != nil {
						errch <- err
					}
			}

			// Decrease the message queue (make sure this is locked )
			if len(messageHandleObject.messageQueue) > 1 {
				messageHandleObject.messageQueue = messageHandleObject.messageQueue[1:]
			} else {
				messageHandleObject.messageQueue = []messageUnit{}
			}

			messageHandleObject.mu.Unlock()
		}
	}
}
