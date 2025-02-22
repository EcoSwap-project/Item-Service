package main

import (
	"log"
	"net"
	c "item_service/config"
	user "item_service/genproto/authentication_service"
	pb "item_service/genproto/item_service"
	"item_service/pkg"
	"item_service/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection" // Import the reflection package
)

func main() {
	// Load configuration
	config := c.Load()

	// Connect to the database
	db, err := pkg.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a TCP listener on specified port
	listener, err := net.Listen("tcp", ":"+config.URL_PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Println("Server started on port " + config.URL_PORT)

	// Connect to the user service
	userConn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Cannot connect to user connection on 50051" + err.Error())
	}
	defer userConn.Close()
	userClient := user.NewEcoServiceClient(userConn)

	

	s := grpc.NewServer()

	// Initialize repository and services
	rs := services.NewMainService(db, userClient) // Pass the itemClient here

	// Create a new gRPC server

	pb.RegisterEcoExchangeServiceServer(s, rs)

	// Register reflection service on gRPC server
	reflection.Register(s)

	// Start the gRPC server
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
