/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.343.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// PartyRoleType A value from a MISMO defined list that identifies the role that the party plays in the transaction. Parties may be either a person or legal entity. A party may play multiple roles in a transaction.A value from a MISMO defined list that identifies the role that the party plays in the transaction. Parties may be either a person or legal entity. A party may play multiple roles in a transaction.
type PartyRoleType string

var _ = fmt.Printf

// List of PartyRoleType
const (
	PARTYROLETYPE_BORROWER PartyRoleType = "Borrower"
)

var allowedPartyRoleTypeEnumValues = []PartyRoleType{
	"Borrower",
}

func (v *PartyRoleType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := PartyRoleType(value)


	*v = enumTypeValue
	return nil
}

// NewPartyRoleTypeFromValue returns a pointer to a valid PartyRoleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPartyRoleTypeFromValue(v string) (*PartyRoleType, error) {
	ev := PartyRoleType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PartyRoleType) IsValid() bool {
	for _, existing := range allowedPartyRoleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PartyRoleType value
func (v PartyRoleType) Ptr() *PartyRoleType {
	return &v
}

type NullablePartyRoleType struct {
	value *PartyRoleType
	isSet bool
}

func (v NullablePartyRoleType) Get() *PartyRoleType {
	return v.value
}

func (v *NullablePartyRoleType) Set(val *PartyRoleType) {
	v.value = val
	v.isSet = true
}

func (v NullablePartyRoleType) IsSet() bool {
	return v.isSet
}

func (v *NullablePartyRoleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartyRoleType(val *PartyRoleType) *NullablePartyRoleType {
	return &NullablePartyRoleType{value: val, isSet: true}
}

func (v NullablePartyRoleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartyRoleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

