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
)

// PartnerEndCustomerWithSecrets The details for the newly created end customer, including secrets for non-Production environments.
type PartnerEndCustomerWithSecrets struct {
	ClientId *string `json:"client_id,omitempty"`
	CompanyName *string `json:"company_name,omitempty"`
	Status *PartnerEndCustomerStatus `json:"status,omitempty"`
	Secrets *PartnerEndCustomerSecrets `json:"secrets,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PartnerEndCustomerWithSecrets PartnerEndCustomerWithSecrets

// NewPartnerEndCustomerWithSecrets instantiates a new PartnerEndCustomerWithSecrets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerEndCustomerWithSecrets() *PartnerEndCustomerWithSecrets {
	this := PartnerEndCustomerWithSecrets{}
	return &this
}

// NewPartnerEndCustomerWithSecretsWithDefaults instantiates a new PartnerEndCustomerWithSecrets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerEndCustomerWithSecretsWithDefaults() *PartnerEndCustomerWithSecrets {
	this := PartnerEndCustomerWithSecrets{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PartnerEndCustomerWithSecrets) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerWithSecrets) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PartnerEndCustomerWithSecrets) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PartnerEndCustomerWithSecrets) SetClientId(v string) {
	o.ClientId = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *PartnerEndCustomerWithSecrets) GetCompanyName() string {
	if o == nil || o.CompanyName == nil {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerWithSecrets) GetCompanyNameOk() (*string, bool) {
	if o == nil || o.CompanyName == nil {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *PartnerEndCustomerWithSecrets) HasCompanyName() bool {
	if o != nil && o.CompanyName != nil {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *PartnerEndCustomerWithSecrets) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PartnerEndCustomerWithSecrets) GetStatus() PartnerEndCustomerStatus {
	if o == nil || o.Status == nil {
		var ret PartnerEndCustomerStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerWithSecrets) GetStatusOk() (*PartnerEndCustomerStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PartnerEndCustomerWithSecrets) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PartnerEndCustomerStatus and assigns it to the Status field.
func (o *PartnerEndCustomerWithSecrets) SetStatus(v PartnerEndCustomerStatus) {
	o.Status = &v
}

// GetSecrets returns the Secrets field value if set, zero value otherwise.
func (o *PartnerEndCustomerWithSecrets) GetSecrets() PartnerEndCustomerSecrets {
	if o == nil || o.Secrets == nil {
		var ret PartnerEndCustomerSecrets
		return ret
	}
	return *o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerWithSecrets) GetSecretsOk() (*PartnerEndCustomerSecrets, bool) {
	if o == nil || o.Secrets == nil {
		return nil, false
	}
	return o.Secrets, true
}

// HasSecrets returns a boolean if a field has been set.
func (o *PartnerEndCustomerWithSecrets) HasSecrets() bool {
	if o != nil && o.Secrets != nil {
		return true
	}

	return false
}

// SetSecrets gets a reference to the given PartnerEndCustomerSecrets and assigns it to the Secrets field.
func (o *PartnerEndCustomerWithSecrets) SetSecrets(v PartnerEndCustomerSecrets) {
	o.Secrets = &v
}

func (o PartnerEndCustomerWithSecrets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.CompanyName != nil {
		toSerialize["company_name"] = o.CompanyName
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Secrets != nil {
		toSerialize["secrets"] = o.Secrets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PartnerEndCustomerWithSecrets) UnmarshalJSON(bytes []byte) (err error) {
	varPartnerEndCustomerWithSecrets := _PartnerEndCustomerWithSecrets{}

	if err = json.Unmarshal(bytes, &varPartnerEndCustomerWithSecrets); err == nil {
		*o = PartnerEndCustomerWithSecrets(varPartnerEndCustomerWithSecrets)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "company_name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "secrets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePartnerEndCustomerWithSecrets struct {
	value *PartnerEndCustomerWithSecrets
	isSet bool
}

func (v NullablePartnerEndCustomerWithSecrets) Get() *PartnerEndCustomerWithSecrets {
	return v.value
}

func (v *NullablePartnerEndCustomerWithSecrets) Set(val *PartnerEndCustomerWithSecrets) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerEndCustomerWithSecrets) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerEndCustomerWithSecrets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerEndCustomerWithSecrets(val *PartnerEndCustomerWithSecrets) *NullablePartnerEndCustomerWithSecrets {
	return &NullablePartnerEndCustomerWithSecrets{value: val, isSet: true}
}

func (v NullablePartnerEndCustomerWithSecrets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerEndCustomerWithSecrets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


