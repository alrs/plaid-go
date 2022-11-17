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

// WeakAliasDetermination Names that are explicitly marked as low quality either by their `source` list, or by `plaid` by a series of additional checks done by Plaid. Plaid does not ever surface a hit as a result of a weak name alone. If a name has no quality issues, this value will be `none`.
type WeakAliasDetermination string

var _ = fmt.Printf

// List of WeakAliasDetermination
const (
	WEAKALIASDETERMINATION_NONE WeakAliasDetermination = "none"
	WEAKALIASDETERMINATION_SOURCE WeakAliasDetermination = "source"
	WEAKALIASDETERMINATION_PLAID WeakAliasDetermination = "plaid"
)

var allowedWeakAliasDeterminationEnumValues = []WeakAliasDetermination{
	"none",
	"source",
	"plaid",
}

func (v *WeakAliasDetermination) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := WeakAliasDetermination(value)


	*v = enumTypeValue
	return nil
}

// NewWeakAliasDeterminationFromValue returns a pointer to a valid WeakAliasDetermination
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWeakAliasDeterminationFromValue(v string) (*WeakAliasDetermination, error) {
	ev := WeakAliasDetermination(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WeakAliasDetermination) IsValid() bool {
	for _, existing := range allowedWeakAliasDeterminationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WeakAliasDetermination value
func (v WeakAliasDetermination) Ptr() *WeakAliasDetermination {
	return &v
}

type NullableWeakAliasDetermination struct {
	value *WeakAliasDetermination
	isSet bool
}

func (v NullableWeakAliasDetermination) Get() *WeakAliasDetermination {
	return v.value
}

func (v *NullableWeakAliasDetermination) Set(val *WeakAliasDetermination) {
	v.value = val
	v.isSet = true
}

func (v NullableWeakAliasDetermination) IsSet() bool {
	return v.isSet
}

func (v *NullableWeakAliasDetermination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeakAliasDetermination(val *WeakAliasDetermination) *NullableWeakAliasDetermination {
	return &NullableWeakAliasDetermination{value: val, isSet: true}
}

func (v NullableWeakAliasDetermination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeakAliasDetermination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

