package asaas

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type errorResponse struct {
	Errors []struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"errors"`
}

type AsaasInterface interface {
	CreateCustomer(newCustomer NewCustomer) (createdCustomer *Customer, err error)
	ReadAllCustomers() (customer *[]Customer, err error)
	ReadCustomer(id string) (customer *Customer, err error)
	UpdateCustomer(customer Customer) (updatedCustomer *Customer, err error)
	DeleteCustomer(id string) (err error)
	RecoverDeletedCustomer(id string) (err error)
	CreatePayment(newPayment NewPayment) (createdPayment *Payment, err error)
	ReadAllPayments() (payments *[]Payment, err error)
	ReadPayment(id string) (payments *Payment, err error)
	UpdatePayment(payments Payment) (updatedPayment *Payment, err error)
	DeletePayment(id string) (err error)
	RecoverDeletedPayment(id string) (err error)
}

type asaas struct {
	AppName string
}

func Connect(AppName string) AsaasInterface {
	return &asaas{
		AppName: AppName,
	}
}

func (asaas *asaas) apiRequest(method string, path string, requestBody *[]byte, queryParams *[]map[string]string) (responseBody []byte, err error) {
	body := &bytes.Reader{}
	if requestBody != nil {
		body = bytes.NewReader(*requestBody)
	}
	client := http.Client{}
	url := fmt.Sprintf("https://sandbox.asaas.com/api/v3%s", path)
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return responseBody, err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", asaas.AppName)
	request.Header.Add("access_token", os.Getenv("ASAAS_KEY"))
	if queryParams != nil {
		query := request.URL.Query()
		for _, paramMap := range *queryParams {
			for key, value := range paramMap {
				query.Add(key, value)
			}
		}
		request.URL.RawQuery = query.Encode()
	}
	response, err := client.Do(request)
	if err != nil {
		return responseBody, err
	}
	defer response.Body.Close()
	responseBody, err = io.ReadAll(response.Body)
	if err != nil || (response.StatusCode != 200 && response.StatusCode != 201) {
		var errorsList *errorResponse
		var errorDescription []string
		err := json.Unmarshal(responseBody, &errorsList)
		if err != nil {
			return responseBody, err
		}
		for _, error := range errorsList.Errors {
			errorDescription = append(errorDescription, error.Description)
		}
		return responseBody, errors.New(strings.Join(errorDescription, ", "))
	}
	return responseBody, err
}
