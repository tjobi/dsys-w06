package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"
	gRPC "w06/proto"

	"google.golang.org/grpc"
)

type Server struct {
	gRPC.UnimplementedPhysicalTimeServiceServer
	name string
	port string

	//mutex sync.Mutex
}

var serverName = flag.String("name", "default", "Senders name")
var port = flag.String("port", "123", "Server port")

func main() {
	flag.Parse()
	fmt.Printf(".:server is starting:.")

	go launchServer()

	for {
		time.Sleep(5 * time.Second)
	}
}

func launchServer() {
	log.Printf("Server %s: Attempts to create listener on port %s\n", *serverName, *port)

	// Create listener tcp on given port or default port 5400
	list, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", *port))
	if err != nil {
		log.Printf("Server %s: Failed to listen on port %s: %v", *serverName, *port, err) //If it fails to listen on the port, run launchServer method again with the next value/port in ports array
		return
	}

	// makes gRPC server using the options
	// you can add options here if you want or remove the options part entirely
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	// makes a new server instance using the name and port from the flags.
	server := &Server{
		name: *serverName,
		port: *port,
	}

	gRPC.RegisterPhysicalTimeServiceServer(grpcServer, server) //Registers the server to the gRPC server.

	log.Printf("Server %s: Listening on port %s\n", *serverName, *port)

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
	// code here is unreachable because grpcServer.Serve occupies the current thread.
}

func (s *Server) CurrentTime(context.Context, *gRPC.ClientRequest) (*gRPC.Time, error) {
	return &gRPC.Time{Timestamp: time.Now().Local().String()}, nil
}
