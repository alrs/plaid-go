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

// PlaidErrorType A broad categorization of the error. Safe for programmatic use.
type PlaidErrorType string

var _ = fmt.Printf

// List of PlaidErrorType
const (
	PLAIDERRORTYPE_INVALID_REQUEST PlaidErrorType = "INVALID_REQUEST"
	PLAIDERRORTYPE_INVALID_RESULT PlaidErrorType = "INVALID_RESULT"
	PLAIDERRORTYPE_INVALID_INPUT PlaidErrorType = "INVALID_INPUT"
	PLAIDERRORTYPE_INSTITUTION_ERROR PlaidErrorType = "INSTITUTION_ERROR"
	PLAIDERRORTYPE_RATE_LIMIT_EXCEEDED PlaidErrorType = "RATE_LIMIT_EXCEEDED"
	PLAIDERRORTYPE_API_ERROR PlaidErrorType = "API_ERROR"
	PLAIDERRORTYPE_ITEM_ERROR PlaidErrorType = "ITEM_ERROR"
	PLAIDERRORTYPE_ASSET_REPORT_ERROR PlaidErrorType = "ASSET_REPORT_ERROR"
	PLAIDERRORTYPE_RECAPTCHA_ERROR PlaidErrorType = "RECAPTCHA_ERROR"
	PLAIDERRORTYPE_OAUTH_ERROR PlaidErrorType = "OAUTH_ERROR"
	PLAIDERRORTYPE_PAYMENT_ERROR PlaidErrorType = "PAYMENT_ERROR"
	PLAIDERRORTYPE_BANK_TRANSFER_ERROR PlaidErrorType = "BANK_TRANSFER_ERROR"
	PLAIDERRORTYPE_INCOME_VERIFICATION_ERROR PlaidErrorType = "INCOME_VERIFICATION_ERROR"
	PLAIDERRORTYPE_MICRODEPOSITS_ERROR PlaidErrorType = "MICRODEPOSITS_ERROR"
)

var allowedPlaidErrorTypeEnumValues = []PlaidErrorType{
	"INVALID_REQUEST",
	"INVALID_RESULT",
	"INVALID_INPUT",
	"INSTITUTION_ERROR",
	"RATE_LIMIT_EXCEEDED",
	"API_ERROR",
	"ITEM_ERROR",
	"ASSET_REPORT_ERROR",
	"RECAPTCHA_ERROR",
	"OAUTH_ERROR",
	"PAYMENT_ERROR",
	"BANK_TRANSFER_ERROR",
	"INCOME_VERIFICATION_ERROR",
	"MICRODEPOSITS_ERROR",
}

func (v *PlaidErrorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := PlaidErrorType(value)


	*v = enumTypeValue
	return nil
}

// NewPlaidErrorTypeFromValue returns a pointer to a valid PlaidErrorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaidErrorTypeFromValue(v string) (*PlaidErrorType, error) {
	ev := PlaidErrorType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaidErrorType) IsValid() bool {
	for _, existing := range allowedPlaidErrorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PlaidErrorType value
func (v PlaidErrorType) Ptr() *PlaidErrorType {
	return &v
}

type NullablePlaidErrorType struct {
	value *PlaidErrorType
	isSet bool
}

func (v NullablePlaidErrorType) Get() *PlaidErrorType {
	return v.value
}

func (v *NullablePlaidErrorType) Set(val *PlaidErrorType) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaidErrorType) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaidErrorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaidErrorType(val *PlaidErrorType) *NullablePlaidErrorType {
	return &NullablePlaidErrorType{value: val, isSet: true}
}

func (v NullablePlaidErrorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaidErrorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

