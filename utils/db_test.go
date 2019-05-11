package utils

import (

	"testing"
)


// TestConnect DB
func TestConnect(t *testing.T) {

	_, err := Connect()
	if err != nil {
		t.Errorf("Failed to connect database - %v", err)
	}

}