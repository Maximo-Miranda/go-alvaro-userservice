package user

import (
	utl "github.com/Maximo-Miranda/go-alvaro-userservice/utils"
)

// UserModel
type UserModel struct {
	ID                   int64
	Name                 string
	Lastname             string
	Email                string
	Password             string
	Phonenumber          string
}

type Users []UserModel

// TableName
func (UserModel) TableName() string {
	return "users"
}

// Create
func (m *UserModel) Create() error {

	conn, err := utl.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	errQuery := conn.Create(m)
	if errQuery.Error != nil {
		return errQuery.Error
	}

	return nil
}

// All
func (m *UserModel) All() (Users, error) {

	users := Users{}

	conn, err := utl.Connect()
	if err != nil {
		return users, err
	}
	defer conn.Close()

	result := conn.Find(&users)
	if result.Error != nil {
		return users, err
	}

	return users, nil
}

// Delete
func (m *UserModel) Delete() error {

	conn, err := utl.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	result := conn.Delete(m)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
