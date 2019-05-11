package user

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"
	goalvaro "github.com/Maximo-Miranda/go-alvaro-userservice/protos/compiled"
)

var ctx context.Context
var emp *empty.Empty

func TestCreateUser(t *testing.T) {

	s := Server{}

	user := &goalvaro.User{
		Name:"Alavar",
		Lastname:"Vega",
		Email:"alvaro@elmistico.com",
		Phonenumber: "3147114206",
		Password: "secret",
	}

	_, err := s.CreateUser(ctx, user)
	if err != nil {
		t.Errorf("Failed to test TestCreateUser - %v", err)
	}
}


func TestListUsers(t *testing.T) {

	s := Server{}

	_, err := s.ListUsers(ctx, emp)
	if err != nil {
		t.Errorf("Failed to test TestListUsers - %v", err)
	}
}