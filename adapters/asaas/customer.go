package asaas

import (
	"encoding/json"
	"fmt"
)

type NewCustomer struct {
	ReferenceID          string  `json:"externalReference" binding:"required"`
	Group                *string `json:"groupName,omitempty"`
	Name                 string  `json:"name" binding:"required"`
	Document             string  `json:"cpfCnpj" binding:"required"`
	Email                *string `json:"email,omitempty"`
	PhoneNumber          *string `json:"phone,omitempty"`
	MobilePhoneNumber    *string `json:"mobilePhone,omitempty"`
	Address              *string `json:"address,omitempty"`
	Number               *string `json:"addressNumber,omitempty"`
	Complement           *string `json:"complement,omitempty"`
	State                *string `json:"province,omitempty"`
	CityName             *string `json:"cityName,omitempty"`
	ZipCode              *string `json:"postalCode,omitempty"`
	Company              *string `json:"company,omitempty"`
	CityInscription      *string `json:"municipalInscription,omitempty"`
	StateInscription     *string `json:"stateInscription,omitempty"`
	Comments             *string `json:"observations,omitempty"`
	AdditionalEmails     *string `json:"additionalEmails,omitempty"`
	NotificationDisabled bool    `json:"notificationDisabled" binding:"required"`
}

type Customer struct {
	ID                   string  `json:"id" binding:"required"`
	ReferenceID          *string `json:"externalReference,omitempty"`
	PersonType           string  `json:"personType" binding:"required"`
	Document             *string `json:"cpfCnpj,omitempty"`
	Name                 string  `json:"name" binding:"required"`
	Email                *string `json:"email,omitempty"`
	PhoneNumber          *string `json:"phone,omitempty"`
	MobilePhoneNumber    *string `json:"mobilePhone,omitempty"`
	Address              *string `json:"address,omitempty"`
	Number               *string `json:"addressNumber,omitempty"`
	Complement           *string `json:"complement,omitempty"`
	Province             *string `json:"province,omitempty"`
	ZipCode              *string `json:"postalCode,omitempty"`
	AdditionalEmails     *string `json:"additionalEmails,omitempty"`
	City                 *int    `json:"city,omitempty"`
	CityName             *string `json:"cityName,omitempty"`
	State                *string `json:"state,omitempty"`
	Country              *string `json:"country,omitempty"`
	Observations         *string `json:"observations,omitempty"`
	NotificationDisabled bool    `json:"notificationDisabled" binding:"required"`
	IsDeleted            bool    `json:"deleted"`
	CreatedAt            string  `json:"dateCreated" binding:"required"`
}

func (asaas *asaas) CreateCustomer(newCustomer NewCustomer) (createdCustomer *Customer, err error) {
	requestBody, err := json.Marshal(newCustomer)
	if err != nil {
		return createdCustomer, err
	}
	responseBody, err := asaas.apiRequest("POST", "/customers/", &requestBody, nil)
	if err != nil {
		return createdCustomer, err
	}
	err = json.Unmarshal(responseBody, &createdCustomer)
	if err != nil {
		return createdCustomer, err
	}
	return createdCustomer, err
}

func (asaas *asaas) ReadAllCustomers() (customers *[]Customer, err error) {
	responseBody, err := asaas.apiRequest("GET", "/customers/", nil, nil)
	if err != nil {
		return customers, err
	}
	err = json.Unmarshal(responseBody, &customers)
	if err != nil {
		return customers, err
	}
	return customers, err
}

func (asaas *asaas) ReadCustomer(id string) (customer *Customer, err error) {
	responseBody, err := asaas.apiRequest("GET", fmt.Sprintf("/customers/%s/", id), nil, nil)
	if err != nil {
		return customer, err
	}
	err = json.Unmarshal(responseBody, &customer)
	if err != nil {
		return customer, err
	}
	return customer, err
}

func (asaas *asaas) UpdateCustomer(customer Customer) (updatedCustomer *Customer, err error) {
	id := customer.ID
	requestBody, err := json.Marshal(customer)
	if err != nil {
		return updatedCustomer, err
	}
	responseBody, err := asaas.apiRequest("PUT", fmt.Sprintf("/customers/%s/", id), &requestBody, nil)
	if err != nil {
		return updatedCustomer, err
	}
	err = json.Unmarshal(responseBody, &updatedCustomer)
	if err != nil {
		return updatedCustomer, err
	}
	return updatedCustomer, err
}

func (asaas *asaas) DeleteCustomer(id string) (err error) {
	_, err = asaas.apiRequest("DELETE", fmt.Sprintf("/customers/%s/", id), nil, nil)
	if err != nil {
		return err
	}
	return err
}

func (asaas *asaas) RecoverDeletedCustomer(id string) (err error) {
	_, err = asaas.apiRequest("POST", fmt.Sprintf("/customers/%s/restore/", id), nil, nil)
	if err != nil {
		return err
	}
	return err
}
