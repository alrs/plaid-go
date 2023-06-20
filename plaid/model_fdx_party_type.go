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

// FDXPartyType Identifies the type of a party
type FDXPartyType string

var _ = fmt.Printf

// List of FDXPartyType
const (
	FDXPARTYTYPE_DATA_ACCESS_PLATFORM FDXPartyType = "DATA_ACCESS_PLATFORM"
	FDXPARTYTYPE_DATA_PROVIDER FDXPartyType = "DATA_PROVIDER"
	FDXPARTYTYPE_DATA_RECIPIENT FDXPartyType = "DATA_RECIPIENT"
	FDXPARTYTYPE_INDIVIDUAL FDXPartyType = "INDIVIDUAL"
	FDXPARTYTYPE_MERCHANT FDXPartyType = "MERCHANT"
	FDXPARTYTYPE_VENDOR FDXPartyType = "VENDOR"
)

var allowedFDXPartyTypeEnumValues = []FDXPartyType{
	"DATA_ACCESS_PLATFORM",
	"DATA_PROVIDER",
	"DATA_RECIPIENT",
	"INDIVIDUAL",
	"MERCHANT",
	"VENDOR",
}

func (v *FDXPartyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := FDXPartyType(value)


	*v = enumTypeValue
	return nil
}

// NewFDXPartyTypeFromValue returns a pointer to a valid FDXPartyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFDXPartyTypeFromValue(v string) (*FDXPartyType, error) {
	ev := FDXPartyType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FDXPartyType) IsValid() bool {
	for _, existing := range allowedFDXPartyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FDXPartyType value
func (v FDXPartyType) Ptr() *FDXPartyType {
	return &v
}

type NullableFDXPartyType struct {
	value *FDXPartyType
	isSet bool
}

func (v NullableFDXPartyType) Get() *FDXPartyType {
	return v.value
}

func (v *NullableFDXPartyType) Set(val *FDXPartyType) {
	v.value = val
	v.isSet = true
}

func (v NullableFDXPartyType) IsSet() bool {
	return v.isSet
}

func (v *NullableFDXPartyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFDXPartyType(val *FDXPartyType) *NullableFDXPartyType {
	return &NullableFDXPartyType{value: val, isSet: true}
}

func (v NullableFDXPartyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFDXPartyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

