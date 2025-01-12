package entities

import (
	"fmt"

	"github.com/Sinet2000/go-eshop-console/exceptions"
)

type ContactInfo struct {
	PhoneNumber string
	Email       string
	Address     string
}

func NewContactInfo(phoneNumber, email, address string) (*ContactInfo, error) {
	if phoneNumber == "" || email == "" {
		return nil, &exceptions.DomainException{Message: "Phone number and email cannot be empty"}
	}

	// Assuming simple validation (you can improve it based on your use case)
	if len(phoneNumber) < 4 {
		return nil, &exceptions.DomainException{Message: "Phone number must be at least 4 characters long"}
	}

	contactInfo := &ContactInfo{
		PhoneNumber: phoneNumber,
		Email:       email,
		Address:     address,
	}

	return contactInfo, nil
}

func (c *ContactInfo) UpdateAddress(newAddress string) {
	c.Address = newAddress
}

func (c *ContactInfo) GetContactInfo() string {
	return fmt.Sprintf("Phone: %s\nEmail: %s\nAddress: %s", c.PhoneNumber, c.Email, c.Address)
}
