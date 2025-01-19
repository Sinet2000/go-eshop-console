package entities

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Sinet2000/go-eshop-console/exceptions"
)

type CustomerType int

const (
	Individual CustomerType = iota
	Company
)

type ContactInfo struct {
	Email string `json:"email" bson:"email"`
	Phone string `json:"phone" bson:"phone"`
}

type Address struct {
	Street     string `json:"street" bson:"street"`
	City       string `json:"city" bson:"city"`
	State      string `json:"state" bson:"state"`
	PostalCode string `json:"postal_code" bson:"postal_code"`
	Country    string `json:"country" bson:"country"`
}

type Customer struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName    string             `json:"first_name" bson:"first_name"`
	LastName     string             `json:"last_name" bson:"last_name"`
	CompanyName  string             `json:"company_name" bson:"company_name"`
	CustomerType CustomerType       `json:"customer_type" bson:"customer_type"`
	ContactInfo  ContactInfo        `json:"contact_info" bson:"contact_info"`
	Address      Address            `json:"address" bson:"address"`
}

func NewCustomer(
	customerType CustomerType,
	firstName, lastName, companyName string,
	contactInfo ContactInfo,
	address Address,
) (*Customer, error) {

	if customerType == Individual {
		if firstName == "" || lastName == "" {
			return nil, &exceptions.DomainException{Message: "First name and last name cannot be empty"}
		}
	}

	if customerType == Company {
		if companyName == "" {
			return nil, &exceptions.DomainException{Message: "Company name cannot be empty"}
		}
	}

	if contactInfo.Email == "" || contactInfo.Phone == "" {
		return nil, &exceptions.DomainException{Message: "Email and phone number cannot be empty"}
	}

	if address.Street == "" || address.City == "" || address.Country == "" {
		return nil, &exceptions.DomainException{Message: "Street, city, and country in the address cannot be empty"}
	}

	customer := &Customer{
		ID:           primitive.NewObjectID(),
		CustomerType: customerType,
		FirstName:    firstName,
		LastName:     lastName,
		CompanyName:  companyName,
		ContactInfo:  contactInfo,
		Address:      address,
	}

	return customer, nil
}

func (c *Customer) GetFullNameOrCompany() string {
	if c.CustomerType == Individual {
		return fmt.Sprintf("%s %s", c.FirstName, c.LastName)
	}
	return c.CompanyName
}

// func (c *Customer) UpdateContactInfo(newContactInfo ContactInfo) {
// 	c.ContactInfo = newContactInfo
// }

// func (c *Customer) String() string {
// 	if c.CustomerType == Individual {
// 		return fmt.Sprintf("Customer ID: %d\nName: %s\nEmail: %s\nPhone: %s\nAddress: %s",
// 			c.ID, c.GetFullNameOrCompany(), c.ContactInfo.Email, c.ContactInfo.PhoneNumber, c.ContactInfo.Address)
// 	}
// 	return fmt.Sprintf("Customer ID: %d\nCompany Name: %s\nEmail: %s\nPhone: %s\nAddress: %s",
// 		c.ID, c.GetFullNameOrCompany(), c.ContactInfo.Email, c.ContactInfo.PhoneNumber, c.ContactInfo.Address)
// }
