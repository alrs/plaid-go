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

// ProcessorTokenPermissionsGetRequest ProcessorTokenPermissionsGetRequest defines the request schema for `/processor/token/permissions/get`
type ProcessorTokenPermissionsGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The processor token obtained from the Plaid integration partner. Processor tokens are in the format: `processor-<environment>-<identifier>`
	ProcessorToken string `json:"processor_token"`
}

// NewProcessorTokenPermissionsGetRequest instantiates a new ProcessorTokenPermissionsGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorTokenPermissionsGetRequest(processorToken string) *ProcessorTokenPermissionsGetRequest {
	this := ProcessorTokenPermissionsGetRequest{}
	this.ProcessorToken = processorToken
	return &this
}

// NewProcessorTokenPermissionsGetRequestWithDefaults instantiates a new ProcessorTokenPermissionsGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorTokenPermissionsGetRequestWithDefaults() *ProcessorTokenPermissionsGetRequest {
	this := ProcessorTokenPermissionsGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ProcessorTokenPermissionsGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorTokenPermissionsGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ProcessorTokenPermissionsGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ProcessorTokenPermissionsGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ProcessorTokenPermissionsGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorTokenPermissionsGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ProcessorTokenPermissionsGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ProcessorTokenPermissionsGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetProcessorToken returns the ProcessorToken field value
func (o *ProcessorTokenPermissionsGetRequest) GetProcessorToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessorToken
}

// GetProcessorTokenOk returns a tuple with the ProcessorToken field value
// and a boolean to check if the value has been set.
func (o *ProcessorTokenPermissionsGetRequest) GetProcessorTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProcessorToken, true
}

// SetProcessorToken sets field value
func (o *ProcessorTokenPermissionsGetRequest) SetProcessorToken(v string) {
	o.ProcessorToken = v
}

func (o ProcessorTokenPermissionsGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["processor_token"] = o.ProcessorToken
	}
	return json.Marshal(toSerialize)
}

type NullableProcessorTokenPermissionsGetRequest struct {
	value *ProcessorTokenPermissionsGetRequest
	isSet bool
}

func (v NullableProcessorTokenPermissionsGetRequest) Get() *ProcessorTokenPermissionsGetRequest {
	return v.value
}

func (v *NullableProcessorTokenPermissionsGetRequest) Set(val *ProcessorTokenPermissionsGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorTokenPermissionsGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorTokenPermissionsGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorTokenPermissionsGetRequest(val *ProcessorTokenPermissionsGetRequest) *NullableProcessorTokenPermissionsGetRequest {
	return &NullableProcessorTokenPermissionsGetRequest{value: val, isSet: true}
}

func (v NullableProcessorTokenPermissionsGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorTokenPermissionsGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


