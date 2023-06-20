/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.379.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// CreditBankIncomeWarningType The warning type which will always be `BANK_INCOME_WARNING`.
type CreditBankIncomeWarningType string

var _ = fmt.Printf

// List of CreditBankIncomeWarningType
const (
	CREDITBANKINCOMEWARNINGTYPE_BANK_INCOME_WARNING CreditBankIncomeWarningType = "BANK_INCOME_WARNING"
)

var allowedCreditBankIncomeWarningTypeEnumValues = []CreditBankIncomeWarningType{
	"BANK_INCOME_WARNING",
}

func (v *CreditBankIncomeWarningType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := CreditBankIncomeWarningType(value)


	*v = enumTypeValue
	return nil
}

// NewCreditBankIncomeWarningTypeFromValue returns a pointer to a valid CreditBankIncomeWarningType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreditBankIncomeWarningTypeFromValue(v string) (*CreditBankIncomeWarningType, error) {
	ev := CreditBankIncomeWarningType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreditBankIncomeWarningType) IsValid() bool {
	for _, existing := range allowedCreditBankIncomeWarningTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreditBankIncomeWarningType value
func (v CreditBankIncomeWarningType) Ptr() *CreditBankIncomeWarningType {
	return &v
}

type NullableCreditBankIncomeWarningType struct {
	value *CreditBankIncomeWarningType
	isSet bool
}

func (v NullableCreditBankIncomeWarningType) Get() *CreditBankIncomeWarningType {
	return v.value
}

func (v *NullableCreditBankIncomeWarningType) Set(val *CreditBankIncomeWarningType) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankIncomeWarningType) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankIncomeWarningType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankIncomeWarningType(val *CreditBankIncomeWarningType) *NullableCreditBankIncomeWarningType {
	return &NullableCreditBankIncomeWarningType{value: val, isSet: true}
}

func (v NullableCreditBankIncomeWarningType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankIncomeWarningType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

