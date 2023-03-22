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
)

// IdentityVerificationRequestUserName You can use this field to pre-populate the user's legal name; if it is provided here, they will not be prompted to enter their name in the identity verification attempt.
type IdentityVerificationRequestUserName struct {
	// A string with at least one non-whitespace character, with a max length of 100 characters.
	GivenName string `json:"given_name"`
	// A string with at least one non-whitespace character, with a max length of 100 characters.
	FamilyName string `json:"family_name"`
}

// NewIdentityVerificationRequestUserName instantiates a new IdentityVerificationRequestUserName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityVerificationRequestUserName(givenName string, familyName string) *IdentityVerificationRequestUserName {
	this := IdentityVerificationRequestUserName{}
	this.GivenName = givenName
	this.FamilyName = familyName
	return &this
}

// NewIdentityVerificationRequestUserNameWithDefaults instantiates a new IdentityVerificationRequestUserName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityVerificationRequestUserNameWithDefaults() *IdentityVerificationRequestUserName {
	this := IdentityVerificationRequestUserName{}
	return &this
}

// GetGivenName returns the GivenName field value
func (o *IdentityVerificationRequestUserName) GetGivenName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationRequestUserName) GetGivenNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GivenName, true
}

// SetGivenName sets field value
func (o *IdentityVerificationRequestUserName) SetGivenName(v string) {
	o.GivenName = v
}

// GetFamilyName returns the FamilyName field value
func (o *IdentityVerificationRequestUserName) GetFamilyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationRequestUserName) GetFamilyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FamilyName, true
}

// SetFamilyName sets field value
func (o *IdentityVerificationRequestUserName) SetFamilyName(v string) {
	o.FamilyName = v
}

func (o IdentityVerificationRequestUserName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["given_name"] = o.GivenName
	}
	if true {
		toSerialize["family_name"] = o.FamilyName
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityVerificationRequestUserName struct {
	value *IdentityVerificationRequestUserName
	isSet bool
}

func (v NullableIdentityVerificationRequestUserName) Get() *IdentityVerificationRequestUserName {
	return v.value
}

func (v *NullableIdentityVerificationRequestUserName) Set(val *IdentityVerificationRequestUserName) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityVerificationRequestUserName) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityVerificationRequestUserName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityVerificationRequestUserName(val *IdentityVerificationRequestUserName) *NullableIdentityVerificationRequestUserName {
	return &NullableIdentityVerificationRequestUserName{value: val, isSet: true}
}

func (v NullableIdentityVerificationRequestUserName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityVerificationRequestUserName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


