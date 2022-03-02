/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.79.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// AuthSupportedMethods Metadata specifically related to which auth methods an institution supports.
type AuthSupportedMethods struct {
	// Indicates if instant auth is supported.
	InstantAuth bool `json:"instant_auth"`
	// Indicates if instant match is supported.
	InstantMatch bool `json:"instant_match"`
	// Indicates if automated microdeposits are supported.
	AutomatedMicroDeposits bool `json:"automated_micro_deposits"`
	AdditionalProperties map[string]interface{}
}

type _AuthSupportedMethods AuthSupportedMethods

// NewAuthSupportedMethods instantiates a new AuthSupportedMethods object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthSupportedMethods(instantAuth bool, instantMatch bool, automatedMicroDeposits bool) *AuthSupportedMethods {
	this := AuthSupportedMethods{}
	this.InstantAuth = instantAuth
	this.InstantMatch = instantMatch
	this.AutomatedMicroDeposits = automatedMicroDeposits
	return &this
}

// NewAuthSupportedMethodsWithDefaults instantiates a new AuthSupportedMethods object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthSupportedMethodsWithDefaults() *AuthSupportedMethods {
	this := AuthSupportedMethods{}
	return &this
}

// GetInstantAuth returns the InstantAuth field value
func (o *AuthSupportedMethods) GetInstantAuth() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InstantAuth
}

// GetInstantAuthOk returns a tuple with the InstantAuth field value
// and a boolean to check if the value has been set.
func (o *AuthSupportedMethods) GetInstantAuthOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstantAuth, true
}

// SetInstantAuth sets field value
func (o *AuthSupportedMethods) SetInstantAuth(v bool) {
	o.InstantAuth = v
}

// GetInstantMatch returns the InstantMatch field value
func (o *AuthSupportedMethods) GetInstantMatch() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InstantMatch
}

// GetInstantMatchOk returns a tuple with the InstantMatch field value
// and a boolean to check if the value has been set.
func (o *AuthSupportedMethods) GetInstantMatchOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstantMatch, true
}

// SetInstantMatch sets field value
func (o *AuthSupportedMethods) SetInstantMatch(v bool) {
	o.InstantMatch = v
}

// GetAutomatedMicroDeposits returns the AutomatedMicroDeposits field value
func (o *AuthSupportedMethods) GetAutomatedMicroDeposits() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutomatedMicroDeposits
}

// GetAutomatedMicroDepositsOk returns a tuple with the AutomatedMicroDeposits field value
// and a boolean to check if the value has been set.
func (o *AuthSupportedMethods) GetAutomatedMicroDepositsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AutomatedMicroDeposits, true
}

// SetAutomatedMicroDeposits sets field value
func (o *AuthSupportedMethods) SetAutomatedMicroDeposits(v bool) {
	o.AutomatedMicroDeposits = v
}

func (o AuthSupportedMethods) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["instant_auth"] = o.InstantAuth
	}
	if true {
		toSerialize["instant_match"] = o.InstantMatch
	}
	if true {
		toSerialize["automated_micro_deposits"] = o.AutomatedMicroDeposits
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthSupportedMethods) UnmarshalJSON(bytes []byte) (err error) {
	varAuthSupportedMethods := _AuthSupportedMethods{}

	if err = json.Unmarshal(bytes, &varAuthSupportedMethods); err == nil {
		*o = AuthSupportedMethods(varAuthSupportedMethods)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "instant_auth")
		delete(additionalProperties, "instant_match")
		delete(additionalProperties, "automated_micro_deposits")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthSupportedMethods struct {
	value *AuthSupportedMethods
	isSet bool
}

func (v NullableAuthSupportedMethods) Get() *AuthSupportedMethods {
	return v.value
}

func (v *NullableAuthSupportedMethods) Set(val *AuthSupportedMethods) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthSupportedMethods) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthSupportedMethods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthSupportedMethods(val *AuthSupportedMethods) *NullableAuthSupportedMethods {
	return &NullableAuthSupportedMethods{value: val, isSet: true}
}

func (v NullableAuthSupportedMethods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthSupportedMethods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


