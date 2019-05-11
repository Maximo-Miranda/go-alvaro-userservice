package user

import (
	"fmt"
	"net/http"
	"golang.org/x/net/context"

	"github.com/golang/protobuf/ptypes/empty"
	utl "github.com/Maximo-Miranda/go-alvaro-userservice/utils"
	goalvaro "github.com/Maximo-Miranda/go-alvaro-userservice/protos/compiled"
)

// Server Instace
type Server struct {
}

// ListUser
func (s *Server) ListUser(ctx context.Context, in *empty.Empty) (*goalvaro.UsersResponse, error) {

	gRPCResponse := goalvaro.UsersResponse{}

	m := UserModel{}

	users, err := m.All()
	if err != nil {
		gRPCResponse = goalvaro.UsersResponse{
			Meta: &goalvaro.Response{
				Message: err.Error(),
				Code: http.StatusInternalServerError,
			},
			Data: nil,
		}

		return &gRPCResponse, err
	}

	for _, u := range users {

		user := goalvaro.User{}

		user.Id = u.ID
		user.Name = u.Name
		user.Lastname = u.Lastname
		user.Email = u.Email
		user.Phonenumber = u.Phonenumber

		gRPCResponse.Data = append(gRPCResponse.Data, &user)

	}

	gRPCResponse.Meta = &goalvaro.Response{
		Message: "Get users OK",
		Code: http.StatusOK,
	}

	return &gRPCResponse, nil
}

// CreateUser
func (s *Server) CreateUser(ctx context.Context, in *goalvaro.User) (*goalvaro.UserResponse, error) {

	gRPCResponse := goalvaro.UserResponse{}

	fmt.Println(in)

	m := UserModel{}

	passowordHashed, err := utl.HashPassword(in.Password)
	if err != nil {
		gRPCResponse = goalvaro.UserResponse{
			Meta: &goalvaro.Response{
				Message: err.Error(),
				Code: http.StatusInternalServerError,
			},
			Data: &goalvaro.User{},
		}

		return &gRPCResponse, err
	}

	m.Password = passowordHashed
	m.Name = in.Name
	m.Lastname = in.Lastname
	m.Email = in.Email
	m.Phonenumber = in.Phonenumber

	err = m.Create()
	if err != nil {
		gRPCResponse = goalvaro.UserResponse{
			Meta: &goalvaro.Response{
				Message: err.Error(),
				Code: http.StatusInternalServerError,
			},
			Data: &goalvaro.User{},
		}

		return &gRPCResponse, err
	}

	in.Id = m.ID

	gRPCResponse = goalvaro.UserResponse{
		Meta: &goalvaro.Response{
			Message: "User created OK",
			Code: http.StatusOK,
		},
		Data: in,
	}

	return &gRPCResponse, nil
}

// GetUserByID
func (s *Server) GetUserByID(ctx context.Context, in *goalvaro.UserId) (*goalvaro.UserResponse, error) {

	return nil, nil
}

// UpdateUser
func (s *Server) UpdateUser(ctx context.Context, in *goalvaro.User) (*goalvaro.UserResponse, error) {

	return nil, nil
}

// DeleteUser
func (s *Server) DeleteUser(ctx context.Context, in *goalvaro.UserId) (*goalvaro.UserBoolResponse, error) {

	gRPCResponse := goalvaro.UserBoolResponse{}

	m := UserModel{
		ID: in.Id,
	}

	err := m.Delete()
	if err != nil {
		gRPCResponse = goalvaro.UserBoolResponse{
			Meta: &goalvaro.Response{
				Message: err.Error(),
				Code: http.StatusInternalServerError,
			},
			Data: false,
		}

		return &gRPCResponse, err
	}

	gRPCResponse = goalvaro.UserBoolResponse{
		Meta: &goalvaro.Response{
			Message: "User deleted OK",
			Code: http.StatusOK,
		},
		Data: true,
	}

	return &gRPCResponse, nil
}
