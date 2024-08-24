package asaas

import (
	"encoding/json"
	"fmt"
)

type Split struct {
	ID              string  `json:"id" binding:"required"`
	ReferenceID     *string `json:"externalReference,omitempty"`
	WalletID        string  `json:"walletId" binding:"required"`
	FixedValue      float64 `json:"fixedValue" binding:"required"`
	PercentualValue float64 `json:"percentualValue" binding:"required"`
	TotalValue      float64 `json:"totalValue" binding:"required"`
	Status          string  `json:"status" binding:"required"`
}

func (asaas *asaas) ReadAllPaidSplits() (splits *[]Split, err error) {
	responseBody, err := asaas.apiRequest("GET", "/payments/splits/paid/", nil, nil)
	if err != nil {
		return splits, err
	}
	err = json.Unmarshal(responseBody, &splits)
	if err != nil {
		return splits, err
	}
	return splits, err
}

func (asaas *asaas) ReadAllPaidSplit(id string) (split *Split, err error) {
	responseBody, err := asaas.apiRequest("GET", fmt.Sprintf("/payments/splits/paid/%s/", id), nil, nil)
	if err != nil {
		return split, err
	}
	err = json.Unmarshal(responseBody, &split)
	if err != nil {
		return split, err
	}
	return split, err
}

func (asaas *asaas) ReadAllReceivedSplits() (splits *[]Split, err error) {
	responseBody, err := asaas.apiRequest("GET", "/payments/splits/received/", nil, nil)
	if err != nil {
		return splits, err
	}
	err = json.Unmarshal(responseBody, &splits)
	if err != nil {
		return splits, err
	}
	return splits, err
}

func (asaas *asaas) ReadAllReceivedSplit(id string) (split *Split, err error) {
	responseBody, err := asaas.apiRequest("GET", fmt.Sprintf("/payments/splits/received/%s/", id), nil, nil)
	if err != nil {
		return split, err
	}
	err = json.Unmarshal(responseBody, &split)
	if err != nil {
		return split, err
	}
	return split, err
}
