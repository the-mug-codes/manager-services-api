package pix

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/r10r/crc16"
)

type pixCopyPasteOptions struct {
	AmountValue   float64 `json:"amount_value"`
	Description   string  `json:"description"`
	TransactionID string  `json:"transaction_id"`
}

type intMap map[int]interface{}

func validateData(key string, name string, city string) (err error) {
	if len(key) == 0 {
		return errors.New("key must not be empty")
	}
	if len(name) == 0 {
		return errors.New("name must not be empty")
	}
	if len(city) == 0 {
		return errors.New("city must not be empty")
	}
	if utf8.RuneCountInString(name) > 25 {
		return errors.New("name must be at least 25 characters long")
	}
	if utf8.RuneCountInString(city) > 15 {
		return errors.New("city must be at least 15 characters long")
	}
	return nil
}

func buildDataMap(pix *pix, options pixCopyPasteOptions) intMap {
	data := make(intMap)
	data[0] = "01"                                                              // Payload Format Indicator
	data[26] = intMap{0: "BR.GOV.BCB.PIX", 1: pix.Key, 2: options.Description}  // Merchant Account Information
	data[52] = "0000"                                                           // Merchant Category Code
	data[53] = "986"                                                            // Transaction Currency - Brazilian Real - ISO4217
	data[54] = options.AmountValue                                              // Transaction Amount
	data[58] = "BR"                                                             // Country Code - ISO3166-1 alpha 2
	data[59] = pix.Name                                                         // Merchant Name. 25 characters maximum
	data[60] = pix.City                                                         // Merchant City. 15 characters maximum
	data[62] = intMap{5: "***", 50: intMap{0: "BR.GOV.BCB.BRCODE", 1: "1.0.0"}} // Transaction ID
	if len(options.TransactionID) != 0 {
		data[62].(intMap)[5] = options.TransactionID
	}
	return data
}

func sortKeys(data intMap) (keys []int) {
	keys = make([]int, len(data))
	index := 0
	for key := range data {
		keys[index] = key
		index++
	}
	sort.Ints(keys)
	return keys
}

func parseData(dataToParse intMap) (parsedData string, err error) {
	dataString := strings.Builder{}
	keys := sortKeys(dataToParse)
	for _, key := range keys {
		keyValue := reflect.ValueOf(dataToParse[key])
		switch keyValue.Kind() {
		case reflect.String:
			value := dataToParse[key].(string)
			_, err = dataString.WriteString(fmt.Sprintf("%02d%02d%s", key, len(value), value))
			if err != nil {
				return parsedData, err
			}
		case reflect.Float64:
			value := strconv.FormatFloat(keyValue.Float(), 'f', 2, 64)
			_, err = dataString.WriteString(fmt.Sprintf("%02d%02d%s", key, len(value), value))
			if err != nil {
				return parsedData, err
			}
		case reflect.Map:
			content, err := parseData(dataToParse[key].(intMap))
			if err != nil {
				return parsedData, err
			}
			_, err = dataString.WriteString(fmt.Sprintf("%02d%02d%s", key, len(content), content))
			if err != nil {
				return parsedData, err
			}
		}
	}
	parsedData = dataString.String()
	return parsedData, err
}

func calculateCRC16(str string) (crc string, err error) {
	table := crc16.MakeTable(crc16.CRC16_CCITT_FALSE)
	h := crc16.New(table)
	_, err = h.Write([]byte(str))
	if err != nil {
		return crc, err
	}
	crc = fmt.Sprintf("%04X", h.Sum16())
	return crc, nil
}

func (pix *pix) CreatePixCopyPasteCode(amountValue *float64, description *string, transactionID *string) (copyPaste string, err error) {
	options := pixCopyPasteOptions{}
	if amountValue != nil {
		options.AmountValue = *amountValue
	}
	if description != nil {
		options.Description = *description
	}
	if transactionID != nil {
		options.TransactionID = strings.ReplaceAll(*transactionID, "-", "")
	}
	codeString := strings.Builder{}
	if err := validateData(pix.Key, pix.Name, pix.City); err != nil {
		return copyPaste, err
	}
	data := buildDataMap(pix, options)
	parsedData, err := parseData(data)
	if err != nil {
		return copyPaste, err
	}
	codeString.WriteString(parsedData)
	if err != nil {
		return copyPaste, err
	}
	codeString.WriteString("6304")
	if err != nil {
		return copyPaste, err
	}
	crc, err := calculateCRC16(codeString.String())
	if err != nil {
		return copyPaste, err
	}
	codeString.WriteString(crc)
	copyPaste = codeString.String()
	return copyPaste, err
}
