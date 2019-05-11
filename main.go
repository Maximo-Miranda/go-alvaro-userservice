package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	goalvaro "github.com/Maximo-Miranda/go-alvaro-userservice/protos/compiled"
	user "github.com/Maximo-Miranda/go-alvaro-userservice/internal/user"
)

func main(){


	// create a listener on TCP port
	lis, err := net.Listen("tcp", ":" + os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the Process service to the server
	goalvaro.RegisterUserServiceServer(grpcServer, &user.Server{})

	log.Println("Server running on port: ", os.Getenv("APP_PORT"))

	// start the server
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}


}