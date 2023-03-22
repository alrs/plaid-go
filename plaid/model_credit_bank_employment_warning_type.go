/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.343.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// CreditBankEmploymentWarningType The warning type which will always be `BANK_EMPLOYMENT_WARNING`.
type CreditBankEmploymentWarningType string

var _ = fmt.Printf

// List of CreditBankEmploymentWarningType
const (
	CREDITBANKEMPLOYMENTWARNINGTYPE_BANK_EMPLOYMENT_WARNING CreditBankEmploymentWarningType = "BANK_EMPLOYMENT_WARNING"
)

var allowedCreditBankEmploymentWarningTypeEnumValues = []CreditBankEmploymentWarningType{
	"BANK_EMPLOYMENT_WARNING",
}

func (v *CreditBankEmploymentWarningType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := CreditBankEmploymentWarningType(value)


	*v = enumTypeValue
	return nil
}

// NewCreditBankEmploymentWarningTypeFromValue returns a pointer to a valid CreditBankEmploymentWarningType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreditBankEmploymentWarningTypeFromValue(v string) (*CreditBankEmploymentWarningType, error) {
	ev := CreditBankEmploymentWarningType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreditBankEmploymentWarningType) IsValid() bool {
	for _, existing := range allowedCreditBankEmploymentWarningTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreditBankEmploymentWarningType value
func (v CreditBankEmploymentWarningType) Ptr() *CreditBankEmploymentWarningType {
	return &v
}

type NullableCreditBankEmploymentWarningType struct {
	value *CreditBankEmploymentWarningType
	isSet bool
}

func (v NullableCreditBankEmploymentWarningType) Get() *CreditBankEmploymentWarningType {
	return v.value
}

func (v *NullableCreditBankEmploymentWarningType) Set(val *CreditBankEmploymentWarningType) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankEmploymentWarningType) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankEmploymentWarningType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankEmploymentWarningType(val *CreditBankEmploymentWarningType) *NullableCreditBankEmploymentWarningType {
	return &NullableCreditBankEmploymentWarningType{value: val, isSet: true}
}

func (v NullableCreditBankEmploymentWarningType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankEmploymentWarningType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

