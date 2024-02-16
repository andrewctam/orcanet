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
 *
 */
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "orcanet/market"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// replace string with User and File types later

// maps a file to a list of users who want this file
var fileUsersMap = make(map[string][]string)

// map of files to users holding the file
var fileHolders = make(map[string][]string)

// prints the current hashmap
func printMap() {
	fmt.Println("Current map:")
	for key, value := range fileUsersMap {
		fmt.Printf("File ID: %s, UserIDs: %v\n", key, value)
	}
}

type server struct {
	pb.UnimplementedMarketServer
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMarketServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error %v", err)
	}
}

// Add a request that a user with userId wants file with fileId
func (s *server) RequestFile(ctx context.Context, in *pb.FileRequest) (*pb.FileResponse, error) {
	userId := in.GetUserId()
	fileId := in.GetFileId()

	// Check if file is held by anyone; I hate Go
	if _, ok := fileHolders[fileId]; !ok {
		return &pb.FileResponse{Exists: false, Message: "File not found"}, nil
	}

	fileUsersMap[fileId] = append(fileUsersMap[fileId], userId)

	return &pb.FileResponse{Exists: true, Message: "OK"}, nil
}

// Get a list of userIds who are requesting a file with fileId
func (s *server) CheckRequests(ctx context.Context, in *pb.CheckRequest) (*pb.ListReply, error) {
	fileId := in.GetFileId()

	userIds := fileUsersMap[fileId]
	printMap()

	return &pb.ListReply{Strings: userIds}, nil
}

// register that the userId holds fileId
func (s *server) RegisterFile(ctx context.Context, in *pb.RegisterRequest) (*emptypb.Empty, error) {
	userId := in.GetUserId()
	fileId := in.GetFileId()

	fileUsersMap[fileId] = append(fileUsersMap[fileId], userId)
	fileHolders[fileId] = append(fileHolders[fileId], userId)

	return &emptypb.Empty{}, nil
}
