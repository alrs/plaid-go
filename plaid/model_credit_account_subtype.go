/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.391.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// CreditAccountSubtype Valid account subtypes for credit accounts. For a list containing descriptions of each subtype, see [Account schemas](https://plaid.com/docs/api/accounts/#StandaloneAccountType-credit).
type CreditAccountSubtype string

var _ = fmt.Printf

// List of CreditAccountSubtype
const (
	CREDITACCOUNTSUBTYPE_CREDIT_CARD CreditAccountSubtype = "credit card"
	CREDITACCOUNTSUBTYPE_PAYPAL CreditAccountSubtype = "paypal"
	CREDITACCOUNTSUBTYPE_ALL CreditAccountSubtype = "all"
)

var allowedCreditAccountSubtypeEnumValues = []CreditAccountSubtype{
	"credit card",
	"paypal",
	"all",
}

func (v *CreditAccountSubtype) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := CreditAccountSubtype(value)


	*v = enumTypeValue
	return nil
}

// NewCreditAccountSubtypeFromValue returns a pointer to a valid CreditAccountSubtype
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreditAccountSubtypeFromValue(v string) (*CreditAccountSubtype, error) {
	ev := CreditAccountSubtype(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreditAccountSubtype) IsValid() bool {
	for _, existing := range allowedCreditAccountSubtypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreditAccountSubtype value
func (v CreditAccountSubtype) Ptr() *CreditAccountSubtype {
	return &v
}

type NullableCreditAccountSubtype struct {
	value *CreditAccountSubtype
	isSet bool
}

func (v NullableCreditAccountSubtype) Get() *CreditAccountSubtype {
	return v.value
}

func (v *NullableCreditAccountSubtype) Set(val *CreditAccountSubtype) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditAccountSubtype) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditAccountSubtype) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditAccountSubtype(val *CreditAccountSubtype) *NullableCreditAccountSubtype {
	return &NullableCreditAccountSubtype{value: val, isSet: true}
}

func (v NullableCreditAccountSubtype) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditAccountSubtype) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

