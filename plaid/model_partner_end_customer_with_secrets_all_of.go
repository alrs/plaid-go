/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.391.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// PartnerEndCustomerWithSecretsAllOf struct for PartnerEndCustomerWithSecretsAllOf
type PartnerEndCustomerWithSecretsAllOf struct {
	Secrets *PartnerEndCustomerSecrets `json:"secrets,omitempty"`
}

// NewPartnerEndCustomerWithSecretsAllOf instantiates a new PartnerEndCustomerWithSecretsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerEndCustomerWithSecretsAllOf() *PartnerEndCustomerWithSecretsAllOf {
	this := PartnerEndCustomerWithSecretsAllOf{}
	return &this
}

// NewPartnerEndCustomerWithSecretsAllOfWithDefaults instantiates a new PartnerEndCustomerWithSecretsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerEndCustomerWithSecretsAllOfWithDefaults() *PartnerEndCustomerWithSecretsAllOf {
	this := PartnerEndCustomerWithSecretsAllOf{}
	return &this
}

// GetSecrets returns the Secrets field value if set, zero value otherwise.
func (o *PartnerEndCustomerWithSecretsAllOf) GetSecrets() PartnerEndCustomerSecrets {
	if o == nil || o.Secrets == nil {
		var ret PartnerEndCustomerSecrets
		return ret
	}
	return *o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerWithSecretsAllOf) GetSecretsOk() (*PartnerEndCustomerSecrets, bool) {
	if o == nil || o.Secrets == nil {
		return nil, false
	}
	return o.Secrets, true
}

// HasSecrets returns a boolean if a field has been set.
func (o *PartnerEndCustomerWithSecretsAllOf) HasSecrets() bool {
	if o != nil && o.Secrets != nil {
		return true
	}

	return false
}

// SetSecrets gets a reference to the given PartnerEndCustomerSecrets and assigns it to the Secrets field.
func (o *PartnerEndCustomerWithSecretsAllOf) SetSecrets(v PartnerEndCustomerSecrets) {
	o.Secrets = &v
}

func (o PartnerEndCustomerWithSecretsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Secrets != nil {
		toSerialize["secrets"] = o.Secrets
	}
	return json.Marshal(toSerialize)
}

type NullablePartnerEndCustomerWithSecretsAllOf struct {
	value *PartnerEndCustomerWithSecretsAllOf
	isSet bool
}

func (v NullablePartnerEndCustomerWithSecretsAllOf) Get() *PartnerEndCustomerWithSecretsAllOf {
	return v.value
}

func (v *NullablePartnerEndCustomerWithSecretsAllOf) Set(val *PartnerEndCustomerWithSecretsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerEndCustomerWithSecretsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerEndCustomerWithSecretsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerEndCustomerWithSecretsAllOf(val *PartnerEndCustomerWithSecretsAllOf) *NullablePartnerEndCustomerWithSecretsAllOf {
	return &NullablePartnerEndCustomerWithSecretsAllOf{value: val, isSet: true}
}

func (v NullablePartnerEndCustomerWithSecretsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerEndCustomerWithSecretsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


