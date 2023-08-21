/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.419.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// ExpirationDate A description of whether the associated document was expired when the verification was performed.  Note: In the case where an expiration date is not present on the document or failed to be extracted, this value will be `no_data`.
type ExpirationDate string

var _ = fmt.Printf

// List of ExpirationDate
const (
	EXPIRATIONDATE_NOT_EXPIRED ExpirationDate = "not_expired"
	EXPIRATIONDATE_EXPIRED ExpirationDate = "expired"
	EXPIRATIONDATE_NO_DATA ExpirationDate = "no_data"
)

var allowedExpirationDateEnumValues = []ExpirationDate{
	"not_expired",
	"expired",
	"no_data",
}

func (v *ExpirationDate) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := ExpirationDate(value)


	*v = enumTypeValue
	return nil
}

// NewExpirationDateFromValue returns a pointer to a valid ExpirationDate
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExpirationDateFromValue(v string) (*ExpirationDate, error) {
	ev := ExpirationDate(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExpirationDate) IsValid() bool {
	for _, existing := range allowedExpirationDateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExpirationDate value
func (v ExpirationDate) Ptr() *ExpirationDate {
	return &v
}

type NullableExpirationDate struct {
	value *ExpirationDate
	isSet bool
}

func (v NullableExpirationDate) Get() *ExpirationDate {
	return v.value
}

func (v *NullableExpirationDate) Set(val *ExpirationDate) {
	v.value = val
	v.isSet = true
}

func (v NullableExpirationDate) IsSet() bool {
	return v.isSet
}

func (v *NullableExpirationDate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpirationDate(val *ExpirationDate) *NullableExpirationDate {
	return &NullableExpirationDate{value: val, isSet: true}
}

func (v NullableExpirationDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpirationDate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

