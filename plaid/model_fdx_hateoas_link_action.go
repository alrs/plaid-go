/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.181.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// FDXHateoasLinkAction HTTP Method to use for the request
type FDXHateoasLinkAction string

// List of FDXHateoasLinkAction
const (
	FDXHATEOASLINKACTION_GET FDXHateoasLinkAction = "GET"
	FDXHATEOASLINKACTION_POST FDXHateoasLinkAction = "POST"
	FDXHATEOASLINKACTION_PATCH FDXHateoasLinkAction = "PATCH"
	FDXHATEOASLINKACTION_DELETE FDXHateoasLinkAction = "DELETE"
	FDXHATEOASLINKACTION_PUT FDXHateoasLinkAction = "PUT"
)

var allowedFDXHateoasLinkActionEnumValues = []FDXHateoasLinkAction{
	"GET",
	"POST",
	"PATCH",
	"DELETE",
	"PUT",
}

func (v *FDXHateoasLinkAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FDXHateoasLinkAction(value)
	for _, existing := range allowedFDXHateoasLinkActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FDXHateoasLinkAction", value)
}

// NewFDXHateoasLinkActionFromValue returns a pointer to a valid FDXHateoasLinkAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFDXHateoasLinkActionFromValue(v string) (*FDXHateoasLinkAction, error) {
	ev := FDXHateoasLinkAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FDXHateoasLinkAction: valid values are %v", v, allowedFDXHateoasLinkActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FDXHateoasLinkAction) IsValid() bool {
	for _, existing := range allowedFDXHateoasLinkActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FDXHateoasLinkAction value
func (v FDXHateoasLinkAction) Ptr() *FDXHateoasLinkAction {
	return &v
}

type NullableFDXHateoasLinkAction struct {
	value *FDXHateoasLinkAction
	isSet bool
}

func (v NullableFDXHateoasLinkAction) Get() *FDXHateoasLinkAction {
	return v.value
}

func (v *NullableFDXHateoasLinkAction) Set(val *FDXHateoasLinkAction) {
	v.value = val
	v.isSet = true
}

func (v NullableFDXHateoasLinkAction) IsSet() bool {
	return v.isSet
}

func (v *NullableFDXHateoasLinkAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFDXHateoasLinkAction(val *FDXHateoasLinkAction) *NullableFDXHateoasLinkAction {
	return &NullableFDXHateoasLinkAction{value: val, isSet: true}
}

func (v NullableFDXHateoasLinkAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFDXHateoasLinkAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

