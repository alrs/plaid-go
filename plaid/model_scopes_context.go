/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.40.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// ScopesContext An indicator for when scopes are being updated. When scopes are updated via enrollment (i.e. OAuth), the partner must send `ENROLLMENT`. When scopes are updated in a post-enrollment view, the partner must send `PORTAL`.
type ScopesContext string

// List of ScopesContext
const (
	SCOPESCONTEXT_ENROLLMENT ScopesContext = "ENROLLMENT"
	SCOPESCONTEXT_PORTAL ScopesContext = "PORTAL"
)

var allowedScopesContextEnumValues = []ScopesContext{
	"ENROLLMENT",
	"PORTAL",
}

func (v *ScopesContext) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ScopesContext(value)
	for _, existing := range allowedScopesContextEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ScopesContext", value)
}

// NewScopesContextFromValue returns a pointer to a valid ScopesContext
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewScopesContextFromValue(v string) (*ScopesContext, error) {
	ev := ScopesContext(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ScopesContext: valid values are %v", v, allowedScopesContextEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ScopesContext) IsValid() bool {
	for _, existing := range allowedScopesContextEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScopesContext value
func (v ScopesContext) Ptr() *ScopesContext {
	return &v
}

type NullableScopesContext struct {
	value *ScopesContext
	isSet bool
}

func (v NullableScopesContext) Get() *ScopesContext {
	return v.value
}

func (v *NullableScopesContext) Set(val *ScopesContext) {
	v.value = val
	v.isSet = true
}

func (v NullableScopesContext) IsSet() bool {
	return v.isSet
}

func (v *NullableScopesContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopesContext(val *ScopesContext) *NullableScopesContext {
	return &NullableScopesContext{value: val, isSet: true}
}

func (v NullableScopesContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopesContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

