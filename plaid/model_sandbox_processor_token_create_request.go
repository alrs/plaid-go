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

// SandboxProcessorTokenCreateRequest SandboxProcessorTokenCreateRequest defines the request schema for `/sandbox/processor_token/create`
type SandboxProcessorTokenCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The ID of the institution the Item will be associated with
	InstitutionId string `json:"institution_id"`
	Options *SandboxProcessorTokenCreateRequestOptions `json:"options,omitempty"`
}

// NewSandboxProcessorTokenCreateRequest instantiates a new SandboxProcessorTokenCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxProcessorTokenCreateRequest(institutionId string) *SandboxProcessorTokenCreateRequest {
	this := SandboxProcessorTokenCreateRequest{}
	this.InstitutionId = institutionId
	return &this
}

// NewSandboxProcessorTokenCreateRequestWithDefaults instantiates a new SandboxProcessorTokenCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxProcessorTokenCreateRequestWithDefaults() *SandboxProcessorTokenCreateRequest {
	this := SandboxProcessorTokenCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SandboxProcessorTokenCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxProcessorTokenCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SandboxProcessorTokenCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SandboxProcessorTokenCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *SandboxProcessorTokenCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxProcessorTokenCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *SandboxProcessorTokenCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *SandboxProcessorTokenCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetInstitutionId returns the InstitutionId field value
func (o *SandboxProcessorTokenCreateRequest) GetInstitutionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value
// and a boolean to check if the value has been set.
func (o *SandboxProcessorTokenCreateRequest) GetInstitutionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstitutionId, true
}

// SetInstitutionId sets field value
func (o *SandboxProcessorTokenCreateRequest) SetInstitutionId(v string) {
	o.InstitutionId = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SandboxProcessorTokenCreateRequest) GetOptions() SandboxProcessorTokenCreateRequestOptions {
	if o == nil || o.Options == nil {
		var ret SandboxProcessorTokenCreateRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxProcessorTokenCreateRequest) GetOptionsOk() (*SandboxProcessorTokenCreateRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SandboxProcessorTokenCreateRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given SandboxProcessorTokenCreateRequestOptions and assigns it to the Options field.
func (o *SandboxProcessorTokenCreateRequest) SetOptions(v SandboxProcessorTokenCreateRequestOptions) {
	o.Options = &v
}

func (o SandboxProcessorTokenCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["institution_id"] = o.InstitutionId
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxProcessorTokenCreateRequest struct {
	value *SandboxProcessorTokenCreateRequest
	isSet bool
}

func (v NullableSandboxProcessorTokenCreateRequest) Get() *SandboxProcessorTokenCreateRequest {
	return v.value
}

func (v *NullableSandboxProcessorTokenCreateRequest) Set(val *SandboxProcessorTokenCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxProcessorTokenCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxProcessorTokenCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxProcessorTokenCreateRequest(val *SandboxProcessorTokenCreateRequest) *NullableSandboxProcessorTokenCreateRequest {
	return &NullableSandboxProcessorTokenCreateRequest{value: val, isSet: true}
}

func (v NullableSandboxProcessorTokenCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxProcessorTokenCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


