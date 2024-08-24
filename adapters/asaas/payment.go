package asaas

import (
	"encoding/json"
	"fmt"
)

type BillingType string

const (
	BillingTypeUndefined  BillingType = "UNDEFINED"
	BillingTypeBankSlip   BillingType = "BOLETO"
	BillingTypeCreditCard BillingType = "CREDIT_CARD"
	BillingTypePIX        BillingType = "PIX"
)

type NewPayment struct {
	Customer                    string           `json:"customer" binding:"required"`
	BillingType                 BillingType      `json:"billingType" binding:"required"`
	Value                       float64          `json:"value" binding:"required"`
	DueDate                     string           `json:"dueDate" binding:"required"`
	Description                 *string          `json:"description,omitempty"`
	BankSlipExpirationAfterDays *int32           `json:"daysAfterDueDateToRegistrationCancellation,omitempty"`
	ExternalReference           *string          `json:"externalReference,omitempty"`
	InstallmentCount            *int32           `json:"installmentCount,omitempty"`
	TotalValue                  *float64         `json:"totalValue,omitempty"`
	InstallmentValue            *float64         `json:"installmentValue,omitempty"`
	Discount                    *PaymentDiscount `json:"discount,omitempty"`
	Interest                    *PaymentInterest `json:"interest,omitempty"`
	Fine                        *PaymentFine     `json:"fine,omitempty"`
	PostalService               *bool            `json:"postalService,omitempty"`
	Split                       []PaymentSplit   `json:"split,omitempty"`
}

type Payment struct {
	ID                    string      `json:"id" binding:"required"`
	DateCreated           string      `json:"dateCreated" binding:"required"`
	Customer              string      `json:"customer" binding:"required"`
	PaymentLink           *string     `json:"paymentLink,omitempty"`
	DueDate               string      `json:"dueDate" binding:"required"`
	Value                 float64     `json:"value" binding:"required"`
	NetValue              float64     `json:"netValue" binding:"required"`
	BillingType           BillingType `json:"billingType" binding:"required"`
	CanBePaidAfterDueDate bool        `json:"canBePaidAfterDueDate" binding:"required"`
	PixTransaction        *string     `json:"pixTransaction,omitempty"`
	Status                string      `json:"status" binding:"required"`
	Description           *string     `json:"description,omitempty"`
	ExternalReference     *string     `json:"externalReference,omitempty"`
	OriginalValue         *float64    `json:"originalValue,omitempty"`
	InterestValue         *float64    `json:"interestValue,omitempty"`
	OriginalDueDate       string      `json:"originalDueDate"  binding:"required"`
	PaymentDate           *string     `json:"paymentDate,omitempty"`
	ClientPaymentDate     *string     `json:"clientPaymentDate,omitempty"`
	InstallmentNumber     *int        `json:"installmentNumber,omitempty"`
	TransactionReceiptURL *string     `json:"transactionReceiptUrl"`
	BankSlipNumber        *string     `json:"nossoNumero,omitempty"`
	InvoiceURL            *string     `json:"invoiceUrl,omitempty"`
	BankSlipURL           *string     `json:"bankSlipUrl,omitempty"`
	InvoiceNumber         *string     `json:"invoiceNumber,omitempty"`
	Discount              *struct {
		Value            int64 `json:"value"  binding:"required"`
		DueDateLimitDays int32 `json:"dueDateLimitDays"  binding:"required"`
	} `json:"discount,omitempty"`
	Fine *struct {
		Value float64 `json:"value"  binding:"required"`
	} `json:"fine,omitempty"`
	Interest *struct {
		Value float64 `json:"value"  binding:"required"`
	} `json:"interest,omitempty"`
	Deleted       bool         `json:"deleted"  binding:"required"`
	PostalService bool         `json:"postalService"  binding:"required"`
	Anticipated   bool         `json:"anticipated"  binding:"required"`
	Anticipable   bool         `json:"anticipable"  binding:"required"`
	Refunds       *interface{} `json:"refunds"` // interface{} para suportar o valor nulo
}

type PaymentDiscountFineType string

const (
	PaymentFineTypeFixed      PaymentDiscountFineType = "FIXED"
	PaymentFineTypePercentage PaymentDiscountFineType = "PERCENTAGE"
)

type PaymentDiscount struct {
	Value            float64                 `json:"value" binding:"required"`
	DueDateLimitDays int                     `json:"dueDateLimitDays" binding:"required"`
	Type             PaymentDiscountFineType `json:"type" binding:"required"`
}

type PaymentInterest struct {
	Value float64 `json:"value" binding:"required"`
}

type PaymentFine struct {
	Value float64                 `json:"value" binding:"required"`
	Type  PaymentDiscountFineType `json:"type" binding:"required"`
}

type PaymentSplit struct {
	WalletId        string   `json:"walletId"  binding:"required"`
	FixedValue      *float64 `json:"fixedValue,omitempty"`
	PercentualValue *float64 `json:"percentualValue,omitempty"`
	TotalFixedValue *float64 `json:"totalFixedValue,omitempty"`
}

func (asaas *asaas) CreatePayment(payment NewPayment) (createdPayment *Payment, err error) {
	requestBody, err := json.Marshal(payment)
	if err != nil {
		return createdPayment, err
	}
	responseBody, err := asaas.apiRequest("POST", "/payments/", &requestBody, nil)
	if err != nil {
		return createdPayment, err
	}
	err = json.Unmarshal(responseBody, &createdPayment)
	if err != nil {
		return createdPayment, err
	}
	return createdPayment, err
}

func (asaas *asaas) ReadAllPayments() (payments *[]Payment, err error) {
	responseBody, err := asaas.apiRequest("put", "/payments/", nil, nil)
	if err != nil {
		return payments, err
	}
	err = json.Unmarshal(responseBody, &payments)
	if err != nil {
		return payments, err
	}
	return payments, err
}

func (asaas *asaas) ReadPayment(id string) (payment *Payment, err error) {
	responseBody, err := asaas.apiRequest("put", fmt.Sprintf("/payments/%s", id), nil, nil)
	if err != nil {
		return payment, err
	}
	err = json.Unmarshal(responseBody, &payment)
	if err != nil {
		return payment, err
	}
	return payment, err
}

func (asaas *asaas) UpdatePayment(payment Payment) (updatedPayment *Payment, err error) {
	id := payment.ID
	requestBody, err := json.Marshal(payment)
	if err != nil {
		return updatedPayment, err
	}
	responseBody, err := asaas.apiRequest("put", fmt.Sprintf("/payments/%s", id), &requestBody, nil)
	if err != nil {
		return updatedPayment, err
	}
	err = json.Unmarshal(responseBody, &updatedPayment)
	if err != nil {
		return updatedPayment, err
	}
	return updatedPayment, err
}

func (asaas *asaas) DeletePayment(id string) (err error) {
	_, err = asaas.apiRequest("delete", fmt.Sprintf("/payments/%s", id), nil, nil)
	if err != nil {
		return err
	}
	return err
}

func (asaas *asaas) RecoverDeletedPayment(id string) (err error) {
	_, err = asaas.apiRequest("post", fmt.Sprintf("/payments/%s/restore/", id), nil, nil)
	if err != nil {
		return err
	}
	return err
}
