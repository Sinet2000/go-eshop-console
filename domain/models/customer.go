package models

import (
	"fmt"

	"github.com/Sinet2000/go-eshop-console/domain/exceptions"
)

type CustomerType int

const (
	Individual CustomerType = iota
	Company
)

type Customer struct {
	ID           int
	FirstName    string
	LastName     string
	CompanyName  string
	CustomerType CustomerType
	ContactInfo  ContactInfo
}

func NewCustomer(id int, customerType CustomerType, firstName, lastName, companyName string, contactInfo ContactInfo) (*Customer, error) {
	// Validate for Person (Individual)
	if customerType == Individual {
		if firstName == "" || lastName == "" {
			return nil, &exceptions.DomainException{Message: "First name and last name cannot be empty"}
		}
	}

	// Validate for Company
	if customerType == Company {
		if companyName == "" {
			return nil, &exceptions.DomainException{Message: "Company name cannot be empty"}
		}
	}

	customer := &Customer{
		ID:           id,
		CustomerType: customerType,
		FirstName:    firstName,
		LastName:     lastName,
		CompanyName:  companyName,
		ContactInfo:  contactInfo,
	}

	return customer, nil
}

func (c *Customer) GetFullNameOrCompany() string {
	if c.CustomerType == Individual {
		return fmt.Sprintf("%s %s", c.FirstName, c.LastName)
	}
	return c.CompanyName
}

func (c *Customer) GetCustomerInfo() string {
	if c.CustomerType == Individual {
		return fmt.Sprintf("Customer ID: %d\nName: %s\nEmail: %s\nPhone: %s\nAddress: %s",
			c.ID, c.GetFullNameOrCompany(), c.ContactInfo.Email, c.ContactInfo.PhoneNumber, c.ContactInfo.Address)
	}
	return fmt.Sprintf("Customer ID: %d\nCompany Name: %s\nEmail: %s\nPhone: %s\nAddress: %s",
		c.ID, c.GetFullNameOrCompany(), c.ContactInfo.Email, c.ContactInfo.PhoneNumber, c.ContactInfo.Address)
}

func (c *Customer) UpdateContactInfo(newContactInfo ContactInfo) {
	c.ContactInfo = newContactInfo
}
