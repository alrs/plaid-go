/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.205.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// Source A type indicating whether a dashboard user, an API-based user, or Plaid last touched this object.
type Source string

var _ = fmt.Printf

// List of Source
const (
	SOURCE_DASHBOARD Source = "dashboard"
	SOURCE_LINK Source = "link"
	SOURCE_API Source = "api"
	SOURCE_SYSTEM Source = "system"
)

var allowedSourceEnumValues = []Source{
	"dashboard",
	"link",
	"api",
	"system",
}

func (v *Source) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := Source(value)


	*v = enumTypeValue
	return nil
}

// NewSourceFromValue returns a pointer to a valid Source
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSourceFromValue(v string) (*Source, error) {
	ev := Source(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Source) IsValid() bool {
	for _, existing := range allowedSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Source value
func (v Source) Ptr() *Source {
	return &v
}

type NullableSource struct {
	value *Source
	isSet bool
}

func (v NullableSource) Get() *Source {
	return v.value
}

func (v *NullableSource) Set(val *Source) {
	v.value = val
	v.isSet = true
}

func (v NullableSource) IsSet() bool {
	return v.isSet
}

func (v *NullableSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSource(val *Source) *NullableSource {
	return &NullableSource{value: val, isSet: true}
}

func (v NullableSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

