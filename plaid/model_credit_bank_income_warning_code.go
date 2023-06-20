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

// CreditBankIncomeWarningCode The warning code identifies a specific kind of warning. `IDENTITY_UNAVAILABLE`: Unable to extract identity for the Item `TRANSACTIONS_UNAVAILABLE`: Unable to extract transactions for the Item `ITEM_UNAPPROVED`: User exited flow before giving permission to share data for the Item `REPORT_DELETED`: Report deleted due to customer or consumer request `DATA_UNAVAILABLE`: No relevant data was found for the Item
type CreditBankIncomeWarningCode string

var _ = fmt.Printf

// List of CreditBankIncomeWarningCode
const (
	CREDITBANKINCOMEWARNINGCODE_IDENTITY_UNAVAILABLE CreditBankIncomeWarningCode = "IDENTITY_UNAVAILABLE"
	CREDITBANKINCOMEWARNINGCODE_TRANSACTIONS_UNAVAILABLE CreditBankIncomeWarningCode = "TRANSACTIONS_UNAVAILABLE"
	CREDITBANKINCOMEWARNINGCODE_ITEM_UNAPPROVED CreditBankIncomeWarningCode = "ITEM_UNAPPROVED"
	CREDITBANKINCOMEWARNINGCODE_REPORT_DELETED CreditBankIncomeWarningCode = "REPORT_DELETED"
	CREDITBANKINCOMEWARNINGCODE_DATA_UNAVAILABLE CreditBankIncomeWarningCode = "DATA_UNAVAILABLE"
)

var allowedCreditBankIncomeWarningCodeEnumValues = []CreditBankIncomeWarningCode{
	"IDENTITY_UNAVAILABLE",
	"TRANSACTIONS_UNAVAILABLE",
	"ITEM_UNAPPROVED",
	"REPORT_DELETED",
	"DATA_UNAVAILABLE",
}

func (v *CreditBankIncomeWarningCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := CreditBankIncomeWarningCode(value)


	*v = enumTypeValue
	return nil
}

// NewCreditBankIncomeWarningCodeFromValue returns a pointer to a valid CreditBankIncomeWarningCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreditBankIncomeWarningCodeFromValue(v string) (*CreditBankIncomeWarningCode, error) {
	ev := CreditBankIncomeWarningCode(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreditBankIncomeWarningCode) IsValid() bool {
	for _, existing := range allowedCreditBankIncomeWarningCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreditBankIncomeWarningCode value
func (v CreditBankIncomeWarningCode) Ptr() *CreditBankIncomeWarningCode {
	return &v
}

type NullableCreditBankIncomeWarningCode struct {
	value *CreditBankIncomeWarningCode
	isSet bool
}

func (v NullableCreditBankIncomeWarningCode) Get() *CreditBankIncomeWarningCode {
	return v.value
}

func (v *NullableCreditBankIncomeWarningCode) Set(val *CreditBankIncomeWarningCode) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankIncomeWarningCode) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankIncomeWarningCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankIncomeWarningCode(val *CreditBankIncomeWarningCode) *NullableCreditBankIncomeWarningCode {
	return &NullableCreditBankIncomeWarningCode{value: val, isSet: true}
}

func (v NullableCreditBankIncomeWarningCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankIncomeWarningCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

