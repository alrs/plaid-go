/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.419.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// PartnerEndCustomerSecrets The secrets for the newly created end customer in non-Production environments.
type PartnerEndCustomerSecrets struct {
	// The end customer's secret key for the Sandbox environment.
	Sandbox *string `json:"sandbox,omitempty"`
	// The end customer's secret key for the Development environment.
	Development *string `json:"development,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PartnerEndCustomerSecrets PartnerEndCustomerSecrets

// NewPartnerEndCustomerSecrets instantiates a new PartnerEndCustomerSecrets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerEndCustomerSecrets() *PartnerEndCustomerSecrets {
	this := PartnerEndCustomerSecrets{}
	return &this
}

// NewPartnerEndCustomerSecretsWithDefaults instantiates a new PartnerEndCustomerSecrets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerEndCustomerSecretsWithDefaults() *PartnerEndCustomerSecrets {
	this := PartnerEndCustomerSecrets{}
	return &this
}

// GetSandbox returns the Sandbox field value if set, zero value otherwise.
func (o *PartnerEndCustomerSecrets) GetSandbox() string {
	if o == nil || o.Sandbox == nil {
		var ret string
		return ret
	}
	return *o.Sandbox
}

// GetSandboxOk returns a tuple with the Sandbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerSecrets) GetSandboxOk() (*string, bool) {
	if o == nil || o.Sandbox == nil {
		return nil, false
	}
	return o.Sandbox, true
}

// HasSandbox returns a boolean if a field has been set.
func (o *PartnerEndCustomerSecrets) HasSandbox() bool {
	if o != nil && o.Sandbox != nil {
		return true
	}

	return false
}

// SetSandbox gets a reference to the given string and assigns it to the Sandbox field.
func (o *PartnerEndCustomerSecrets) SetSandbox(v string) {
	o.Sandbox = &v
}

// GetDevelopment returns the Development field value if set, zero value otherwise.
func (o *PartnerEndCustomerSecrets) GetDevelopment() string {
	if o == nil || o.Development == nil {
		var ret string
		return ret
	}
	return *o.Development
}

// GetDevelopmentOk returns a tuple with the Development field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerSecrets) GetDevelopmentOk() (*string, bool) {
	if o == nil || o.Development == nil {
		return nil, false
	}
	return o.Development, true
}

// HasDevelopment returns a boolean if a field has been set.
func (o *PartnerEndCustomerSecrets) HasDevelopment() bool {
	if o != nil && o.Development != nil {
		return true
	}

	return false
}

// SetDevelopment gets a reference to the given string and assigns it to the Development field.
func (o *PartnerEndCustomerSecrets) SetDevelopment(v string) {
	o.Development = &v
}

func (o PartnerEndCustomerSecrets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Sandbox != nil {
		toSerialize["sandbox"] = o.Sandbox
	}
	if o.Development != nil {
		toSerialize["development"] = o.Development
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PartnerEndCustomerSecrets) UnmarshalJSON(bytes []byte) (err error) {
	varPartnerEndCustomerSecrets := _PartnerEndCustomerSecrets{}

	if err = json.Unmarshal(bytes, &varPartnerEndCustomerSecrets); err == nil {
		*o = PartnerEndCustomerSecrets(varPartnerEndCustomerSecrets)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "sandbox")
		delete(additionalProperties, "development")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePartnerEndCustomerSecrets struct {
	value *PartnerEndCustomerSecrets
	isSet bool
}

func (v NullablePartnerEndCustomerSecrets) Get() *PartnerEndCustomerSecrets {
	return v.value
}

func (v *NullablePartnerEndCustomerSecrets) Set(val *PartnerEndCustomerSecrets) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerEndCustomerSecrets) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerEndCustomerSecrets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerEndCustomerSecrets(val *PartnerEndCustomerSecrets) *NullablePartnerEndCustomerSecrets {
	return &NullablePartnerEndCustomerSecrets{value: val, isSet: true}
}

func (v NullablePartnerEndCustomerSecrets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerEndCustomerSecrets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


