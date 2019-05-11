package main

import (
	"log"

	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"github.com/golang/protobuf/ptypes/empty"
	goalvaro "github.com/Maximo-Miranda/go-alvaro-userservice/protos/compiled"
)

func main(){

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8007", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

/*	str := goalvaro.R
	str.Name = "Elisa"
	str.Lastname = "Arrieta"
	str.Auth_0Id = "dfge5t45t45"
	str.Email = "valery4@hotmail.com"
	str.ImageUrl = "http://localhost"
	str.IsSocial = "Google"
	str.PhoneNumber = "+573008316105"
*/
	c := goalvaro.NewUserServiceClient(conn)

	response, err := c.ListUser(context.Background(), &empty.Empty{})

	if err != nil {
		log.Fatalf("Error when calling: %s", err)
	}

	log.Printf("Response from server: %v", response)

}