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
	Customer                    string                   `json:"customer" binding:"required"`
	BillingType                 BillingType              `json:"billingType" binding:"required"`
	Value                       float64                  `json:"value" binding:"required"`
	DueDate                     string                   `json:"dueDate" binding:"required"`
	Description                 *string                  `json:"description,omitempty"`
	BankSlipExpirationAfterDays *int32                   `json:"daysAfterDueDateToRegistrationCancellation,omitempty"`
	ExternalReference           *string                  `json:"externalReference,omitempty"`
	InstallmentCount            *int32                   `json:"installmentCount,omitempty"`
	TotalValue                  *float64                 `json:"totalValue,omitempty"`
	InstallmentValue            *float64                 `json:"installmentValue,omitempty"`
	Discount                    *PaymentDiscount         `json:"discount,omitempty"`
	Interest                    *PaymentInterest         `json:"interest,omitempty"`
	Fine                        *PaymentFine             `json:"fine,omitempty"`
	PostalService               *bool                    `json:"postalService,omitempty"`
	Split                       []PaymentSplit           `json:"split,omitempty"`
	CreditCard                  *PaymentCreditCard       `json:"creditCard,omitempty"`
	CreditCardHolder            *PaymentCreditCardHolder `json:"creditCardHolderInfo,omitempty"`
	CreditCardToken             *string                  `json:"creditCardToken,omitempty"`
	CreditCardAuthorizeOnly     *bool                    `json:"authorizeOnly,omitempty"`
	UserRemoteIp                *string                  `json:"remoteIp,omitempty"`
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
	OriginalDueDate       string      `json:"originalDueDate" binding:"required"`
	PaymentDate           *string     `json:"paymentDate,omitempty"`
	ClientPaymentDate     *string     `json:"clientPaymentDate,omitempty"`
	InstallmentNumber     *int        `json:"installmentNumber,omitempty"`
	TransactionReceiptURL *string     `json:"transactionReceiptUrl"`
	BankSlipNumber        *string     `json:"nossoNumero,omitempty"`
	InvoiceURL            *string     `json:"invoiceUrl,omitempty"`
	BankSlipURL           *string     `json:"bankSlipUrl,omitempty"`
	InvoiceNumber         *string     `json:"invoiceNumber,omitempty"`
	Discount              *struct {
		Value            int64 `json:"value" binding:"required"`
		DueDateLimitDays int32 `json:"dueDateLimitDays" binding:"required"`
	} `json:"discount,omitempty"`
	Fine *struct {
		Value float64 `json:"value" binding:"required"`
	} `json:"fine,omitempty"`
	Interest *struct {
		Value float64 `json:"value" binding:"required"`
	} `json:"interest,omitempty"`
	Deleted       bool         `json:"deleted" binding:"required"`
	PostalService bool         `json:"postalService" binding:"required"`
	Anticipated   bool         `json:"anticipated" binding:"required"`
	Anticipable   bool         `json:"anticipable" binding:"required"`
	Refunds       *interface{} `json:"refunds"`
}

type PaymentDiscountFineType string

const (
	PaymentFineTypeFixed      PaymentDiscountFineType = "FIXED"
	PaymentFineTypePercentage PaymentDiscountFineType = "PERCENTAGE"
)

type PaymentCreditCard struct {
	HolderName  string `json:"holderName" binding:"required"`
	Number      string `json:"number" binding:"required"`
	ExpiryMonth string `json:"expiryMonth" binding:"required"`
	ExpiryYear  string `json:"expiryYear" binding:"required"`
	CCV         string `json:"ccv" binding:"required"`
}

type PaymentCreditCardHolder struct {
	Name        string  `json:"name" binding:"required"`
	Email       string  `json:"email" validate:"required,email"`
	Document    string  `json:"cpfCnpj" binding:"required"`
	ZipCode     string  `json:"postalCode" binding:"required"`
	Number      string  `json:"addressNumber" binding:"required"`
	Complement  *string `json:"addressComplement,omitempty"`
	Phone       string  `json:"phone" binding:"required"`
	MobilePhone *string `json:"mobilePhone,omitempty"`
}

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
	WalletId        string   `json:"walletId" binding:"required"`
	FixedValue      *float64 `json:"fixedValue,omitempty"`
	PercentualValue *float64 `json:"percentualValue,omitempty"`
	TotalFixedValue *float64 `json:"totalFixedValue,omitempty"`
}

type NewCreditCardToken struct {
	Customer         string                  `json:"customer" binding:"required"`
	CreditCard       PaymentCreditCard       `json:"creditCard,omitempty"`
	CreditCardHolder PaymentCreditCardHolder `json:"creditCardHolderInfo,omitempty"`
	UserRemoteIp     string                  `json:"remoteIp,omitempty"`
}

type CreditCardToken struct {
	Number string `json:"creditCardNumber" binding:"required"`
	Brand  string `json:"creditCardBrand" binding:"required"`
	Token  string `json:"creditCardToken" binding:"required"`
}

type Pix struct {
	QrCode       string `json:"encodedImage" binding:"required"`
	CopyPAste    string `json:"payload" binding:"required"`
	ExpirationAt string `json:"expirationDate" binding:"required"`
}

type BankSlip struct {
	ID             string `json:"identificationField" binding:"required"`
	BankSlipNumber string `json:"nossoNumero" binding:"required"`
	BarCode        string `json:"barCode" binding:"required"`
}

type PaymentSimulation struct {
	Value        float64 `json:"value" binding:"required"`
	Installments int     `json:"installmentCount" binding:"required"`
	CreditCard   *struct {
		NetValue      float64 `json:"netValue,omitempty"`
		FeePercentage float64 `json:"feePercentage,omitempty"`
		OperationFee  float64 `json:"operationFee,omitempty"`
		Installment   struct {
			PaymentNetValue float64 `json:"paymentNetValue,omitempty"`
			PaymentValue    float64 `json:"paymentValue,omitempty"`
		} `json:"installment,omitempty"`
	} `json:"creditCard,omitempty"`
	BankSlip *struct {
		NetValue    float64 `json:"netValue,omitempty"`
		FeeValue    float64 `json:"feeValue,omitempty"`
		Installment struct {
			PaymentNetValue float64 `json:"paymentNetValue,omitempty"`
			PaymentValue    float64 `json:"paymentValue,omitempty"`
		} `json:"installment,omitempty"`
	} `json:"bankSlip,omitempty"`
	Pix *struct {
		NetValue      float64 `json:"netValue,omitempty"`
		FeePercentage float64 `json:"feePercentage,omitempty"`
		FeeValue      float64 `json:"feeValue,omitempty"`
		Installment   struct {
			PaymentNetValue float64 `json:"paymentNetValue,omitempty"`
			PaymentValue    float64 `json:"paymentValue,omitempty"`
		} `json:"installment,omitempty"`
	} `json:"pix,omitempty"`
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
	responseBody, err := asaas.apiRequest("GET", "/payments/", nil, nil)
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
	responseBody, err := asaas.apiRequest("GET", fmt.Sprintf("/payments/%s/", id), nil, nil)
	if err != nil {
		return payment, err
	}
	err = json.Unmarshal(responseBody, &payment)
	if err != nil {
		return payment, err
	}
	return payment, err
}

func (asaas *asaas) ReadPaymentBankSlip(id string) (bankSlip *BankSlip, err error) {
	responseBody, err := asaas.apiRequest("GET", fmt.Sprintf("/payments/%s/identificationField", id), nil, nil)
	if err != nil {
		return bankSlip, err
	}
	err = json.Unmarshal(responseBody, &bankSlip)
	if err != nil {
		return bankSlip, err
	}
	return bankSlip, err
}

func (asaas *asaas) RefundPaymentPix(id string) (pix *Pix, err error) {
	responseBody, err := asaas.apiRequest("GET", fmt.Sprintf("/payments/%s/pixQrCode", id), nil, nil)
	if err != nil {
		return pix, err
	}
	err = json.Unmarshal(responseBody, &pix)
	if err != nil {
		return pix, err
	}
	return pix, err
}

func (asaas *asaas) ReadPaymentStatus(id string) (status *string, err error) {
	var statusData = map[string]string{}
	responseBody, err := asaas.apiRequest("GET", fmt.Sprintf("/payments/%s/status/", id), nil, nil)
	if err != nil {
		return status, err
	}
	err = json.Unmarshal(responseBody, &statusData)
	if err != nil {
		return status, err
	}
	*status = statusData["status"]
	return status, err
}

func (asaas *asaas) UpdatePayment(payment Payment) (updatedPayment *Payment, err error) {
	id := payment.ID
	requestBody, err := json.Marshal(payment)
	if err != nil {
		return updatedPayment, err
	}
	responseBody, err := asaas.apiRequest("PUT", fmt.Sprintf("/payments/%s/", id), &requestBody, nil)
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
	_, err = asaas.apiRequest("DELETE", fmt.Sprintf("/payments/%s/", id), nil, nil)
	if err != nil {
		return err
	}
	return err
}

func (asaas *asaas) RecoverDeletedPayment(id string) (err error) {
	_, err = asaas.apiRequest("POST", fmt.Sprintf("/payments/%s/restore/", id), nil, nil)
	if err != nil {
		return err
	}
	return err
}

func (asaas *asaas) CapturePreAuthorizedPayment(id string) (err error) {
	_, err = asaas.apiRequest("POST", fmt.Sprintf("/payments/%s/captureAuthorized/", id), nil, nil)
	if err != nil {
		return err
	}
	return err
}

func (asaas *asaas) CreateCreditCardToken(newCreditCardToken NewCreditCardToken) (creditCardToken *CreditCardToken, err error) {
	requestBody, err := json.Marshal(newCreditCardToken)
	if err != nil {
		return creditCardToken, err
	}
	responseBody, err := asaas.apiRequest("POST", "/creditCard/tokenize/", &requestBody, nil)
	if err != nil {
		return creditCardToken, err
	}
	err = json.Unmarshal(responseBody, &creditCardToken)
	if err != nil {
		return creditCardToken, err
	}
	return creditCardToken, err
}

func (asaas *asaas) PayWithCreditCard(id string, creditCardToken string) (err error) {
	requestData := map[string]string{
		"creditCardToken": creditCardToken,
	}
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return err
	}
	_, err = asaas.apiRequest("POST", fmt.Sprintf("/payments/%s/payWithCreditCard/", id), &requestBody, nil)
	if err != nil {
		return err
	}
	return err
}

func (asaas *asaas) RefundPayment(id string, value float64, description string) (err error) {
	requestData := map[string]interface{}{
		"value":       value,
		"description": description,
	}
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return err
	}
	_, err = asaas.apiRequest("POST", fmt.Sprintf("/payments/%s/refund/", id), &requestBody, nil)
	if err != nil {
		return err
	}
	return err
}

func (asaas *asaas) SimulatePayment(value float64, installments int, billingType []BillingType) (paymentSimulation *PaymentSimulation, err error) {
	requestData := map[string]interface{}{
		"value":            value,
		"installmentCount": installments,
		"billingTypes":     billingType,
	}
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return paymentSimulation, err
	}
	responseBody, err := asaas.apiRequest("POST", "/payments/simulate", &requestBody, nil)
	if err != nil {
		return paymentSimulation, err
	}
	err = json.Unmarshal(responseBody, &paymentSimulation)
	if err != nil {
		return paymentSimulation, err
	}
	return paymentSimulation, err
}
