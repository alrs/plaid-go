/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.121.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// IssuingCountry A binary match indicator specifying whether the country that issued the provided document matches the country that the user separately provided to Plaid.  Note: You can configure whether a `no_match` on `issuing_country` fails the `documentary_verification` by editing your Plaid Template.
type IssuingCountry string

// List of IssuingCountry
const (
	ISSUINGCOUNTRY_MATCH IssuingCountry = "match"
	ISSUINGCOUNTRY_NO_MATCH IssuingCountry = "no_match"
)

var allowedIssuingCountryEnumValues = []IssuingCountry{
	"match",
	"no_match",
}

func (v *IssuingCountry) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IssuingCountry(value)
	for _, existing := range allowedIssuingCountryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IssuingCountry", value)
}

// NewIssuingCountryFromValue returns a pointer to a valid IssuingCountry
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIssuingCountryFromValue(v string) (*IssuingCountry, error) {
	ev := IssuingCountry(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IssuingCountry: valid values are %v", v, allowedIssuingCountryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IssuingCountry) IsValid() bool {
	for _, existing := range allowedIssuingCountryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssuingCountry value
func (v IssuingCountry) Ptr() *IssuingCountry {
	return &v
}

type NullableIssuingCountry struct {
	value *IssuingCountry
	isSet bool
}

func (v NullableIssuingCountry) Get() *IssuingCountry {
	return v.value
}

func (v *NullableIssuingCountry) Set(val *IssuingCountry) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuingCountry) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuingCountry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuingCountry(val *IssuingCountry) *NullableIssuingCountry {
	return &NullableIssuingCountry{value: val, isSet: true}
}

func (v NullableIssuingCountry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuingCountry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

