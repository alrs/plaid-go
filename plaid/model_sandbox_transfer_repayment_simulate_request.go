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

// SandboxTransferRepaymentSimulateRequest Defines the request schema for `/sandbox/transfer/repayment/simulate`
type SandboxTransferRepaymentSimulateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
}

// NewSandboxTransferRepaymentSimulateRequest instantiates a new SandboxTransferRepaymentSimulateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxTransferRepaymentSimulateRequest() *SandboxTransferRepaymentSimulateRequest {
	this := SandboxTransferRepaymentSimulateRequest{}
	return &this
}

// NewSandboxTransferRepaymentSimulateRequestWithDefaults instantiates a new SandboxTransferRepaymentSimulateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxTransferRepaymentSimulateRequestWithDefaults() *SandboxTransferRepaymentSimulateRequest {
	this := SandboxTransferRepaymentSimulateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SandboxTransferRepaymentSimulateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxTransferRepaymentSimulateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SandboxTransferRepaymentSimulateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SandboxTransferRepaymentSimulateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *SandboxTransferRepaymentSimulateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxTransferRepaymentSimulateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *SandboxTransferRepaymentSimulateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *SandboxTransferRepaymentSimulateRequest) SetSecret(v string) {
	o.Secret = &v
}

func (o SandboxTransferRepaymentSimulateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxTransferRepaymentSimulateRequest struct {
	value *SandboxTransferRepaymentSimulateRequest
	isSet bool
}

func (v NullableSandboxTransferRepaymentSimulateRequest) Get() *SandboxTransferRepaymentSimulateRequest {
	return v.value
}

func (v *NullableSandboxTransferRepaymentSimulateRequest) Set(val *SandboxTransferRepaymentSimulateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxTransferRepaymentSimulateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxTransferRepaymentSimulateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxTransferRepaymentSimulateRequest(val *SandboxTransferRepaymentSimulateRequest) *NullableSandboxTransferRepaymentSimulateRequest {
	return &NullableSandboxTransferRepaymentSimulateRequest{value: val, isSet: true}
}

func (v NullableSandboxTransferRepaymentSimulateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxTransferRepaymentSimulateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


