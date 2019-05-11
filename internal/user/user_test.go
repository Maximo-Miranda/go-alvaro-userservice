package user

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"
	goalvaro "github.com/Maximo-Miranda/go-alvaro-userservice/protos/compiled"
)

var ctx context.Context
var emp *empty.Empty
var user *goalvaro.User

// TestServer_CreateUser
func TestServer_CreateUser(t *testing.T) {

	s := Server{}

	user = &goalvaro.User{
		Name:"TestName",
		Lastname:"TestLastname",
		Email:"test@corp.com",
		Phonenumber: "583872",
		Password: "secret",
	}

	_, err := s.CreateUser(ctx, user)
	if err != nil {
		t.Errorf("Failed to test TestCreateUser - %v", err)
	}

	t.Log(user)
}

// TestServer_ListUser
func TestServer_ListUser(t *testing.T) {

	s := Server{}

	response, err := s.ListUser(ctx, emp)
	if err != nil {
		t.Errorf("Failed to test TestListUsers - %v", err.Error())
	}

	if response.Data[0].Id != user.Id {
		t.Error("Failed to test TestListUsers")
	}
}

// TestServer_DeleteUser
func TestServer_DeleteUser(t *testing.T) {

	s := Server{}

	_, err := s.DeleteUser(ctx, &goalvaro.UserId{Id: user.Id})
	if err != nil {
		t.Errorf("Failed to test TestServer_DeleteUser - %v", err)
	}

}