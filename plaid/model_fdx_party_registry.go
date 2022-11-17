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

// FDXPartyRegistry The registry containing the party’s registration with name and id
type FDXPartyRegistry string

var _ = fmt.Printf

// List of FDXPartyRegistry
const (
	FDXPARTYREGISTRY_FDX FDXPartyRegistry = "FDX"
	FDXPARTYREGISTRY_GLEIF FDXPartyRegistry = "GLEIF"
	FDXPARTYREGISTRY_ICANN FDXPartyRegistry = "ICANN"
	FDXPARTYREGISTRY_PRIVATE FDXPartyRegistry = "PRIVATE"
)

var allowedFDXPartyRegistryEnumValues = []FDXPartyRegistry{
	"FDX",
	"GLEIF",
	"ICANN",
	"PRIVATE",
}

func (v *FDXPartyRegistry) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := FDXPartyRegistry(value)


	*v = enumTypeValue
	return nil
}

// NewFDXPartyRegistryFromValue returns a pointer to a valid FDXPartyRegistry
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFDXPartyRegistryFromValue(v string) (*FDXPartyRegistry, error) {
	ev := FDXPartyRegistry(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FDXPartyRegistry) IsValid() bool {
	for _, existing := range allowedFDXPartyRegistryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FDXPartyRegistry value
func (v FDXPartyRegistry) Ptr() *FDXPartyRegistry {
	return &v
}

type NullableFDXPartyRegistry struct {
	value *FDXPartyRegistry
	isSet bool
}

func (v NullableFDXPartyRegistry) Get() *FDXPartyRegistry {
	return v.value
}

func (v *NullableFDXPartyRegistry) Set(val *FDXPartyRegistry) {
	v.value = val
	v.isSet = true
}

func (v NullableFDXPartyRegistry) IsSet() bool {
	return v.isSet
}

func (v *NullableFDXPartyRegistry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFDXPartyRegistry(val *FDXPartyRegistry) *NullableFDXPartyRegistry {
	return &NullableFDXPartyRegistry{value: val, isSet: true}
}

func (v NullableFDXPartyRegistry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFDXPartyRegistry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

