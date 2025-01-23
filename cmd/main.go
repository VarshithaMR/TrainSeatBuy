package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"TrainSeatBuy/cmd/config"
	"TrainSeatBuy/service"
	"TrainSeatBuy/service/proto/generated"
)

func main() {
	configuration, err := config.LoadConfig("env/properties.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	grpcServer := grpc.NewServer()
	ticketServiceServer := service.NewTicketServiceServer(&configuration.SeatConfig)
	generated.RegisterTicketServiceServer(grpcServer, ticketServiceServer)

	listener, err := net.Listen("tcp", configuration.ServerConfig.Port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", configuration.ServerConfig.Port, err)
	}

	// Start the gRPC server
	log.Printf("Starting gRPC server on %s", configuration.ServerConfig.Port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}

}
