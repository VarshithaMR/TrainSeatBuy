package main

import (
	"fmt"
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

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", configuration.ServerConfig.Host, configuration.ServerConfig.Port))
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
