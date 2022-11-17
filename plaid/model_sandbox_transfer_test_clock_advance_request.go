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
	"time"
)

// SandboxTransferTestClockAdvanceRequest Defines the request schema for `/sandbox/transfer/test_clock/advance`
type SandboxTransferTestClockAdvanceRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId string `json:"client_id"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret string `json:"secret"`
	// Plaid’s unique identifier for a test clock.
	TestClockId NullableString `json:"test_clock_id"`
	// The frozen timestamp on the test clock. This will be of the form `2006-01-02T15:04:05Z`.
	NewFrozenTimestamp time.Time `json:"new_frozen_timestamp"`
}

// NewSandboxTransferTestClockAdvanceRequest instantiates a new SandboxTransferTestClockAdvanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxTransferTestClockAdvanceRequest(clientId string, secret string, testClockId NullableString, newFrozenTimestamp time.Time) *SandboxTransferTestClockAdvanceRequest {
	this := SandboxTransferTestClockAdvanceRequest{}
	this.ClientId = clientId
	this.Secret = secret
	this.TestClockId = testClockId
	this.NewFrozenTimestamp = newFrozenTimestamp
	return &this
}

// NewSandboxTransferTestClockAdvanceRequestWithDefaults instantiates a new SandboxTransferTestClockAdvanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxTransferTestClockAdvanceRequestWithDefaults() *SandboxTransferTestClockAdvanceRequest {
	this := SandboxTransferTestClockAdvanceRequest{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *SandboxTransferTestClockAdvanceRequest) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *SandboxTransferTestClockAdvanceRequest) GetClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *SandboxTransferTestClockAdvanceRequest) SetClientId(v string) {
	o.ClientId = v
}

// GetSecret returns the Secret field value
func (o *SandboxTransferTestClockAdvanceRequest) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *SandboxTransferTestClockAdvanceRequest) GetSecretOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *SandboxTransferTestClockAdvanceRequest) SetSecret(v string) {
	o.Secret = v
}

// GetTestClockId returns the TestClockId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SandboxTransferTestClockAdvanceRequest) GetTestClockId() string {
	if o == nil || o.TestClockId.Get() == nil {
		var ret string
		return ret
	}

	return *o.TestClockId.Get()
}

// GetTestClockIdOk returns a tuple with the TestClockId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SandboxTransferTestClockAdvanceRequest) GetTestClockIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TestClockId.Get(), o.TestClockId.IsSet()
}

// SetTestClockId sets field value
func (o *SandboxTransferTestClockAdvanceRequest) SetTestClockId(v string) {
	o.TestClockId.Set(&v)
}

// GetNewFrozenTimestamp returns the NewFrozenTimestamp field value
func (o *SandboxTransferTestClockAdvanceRequest) GetNewFrozenTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.NewFrozenTimestamp
}

// GetNewFrozenTimestampOk returns a tuple with the NewFrozenTimestamp field value
// and a boolean to check if the value has been set.
func (o *SandboxTransferTestClockAdvanceRequest) GetNewFrozenTimestampOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NewFrozenTimestamp, true
}

// SetNewFrozenTimestamp sets field value
func (o *SandboxTransferTestClockAdvanceRequest) SetNewFrozenTimestamp(v time.Time) {
	o.NewFrozenTimestamp = v
}

func (o SandboxTransferTestClockAdvanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["test_clock_id"] = o.TestClockId.Get()
	}
	if true {
		toSerialize["new_frozen_timestamp"] = o.NewFrozenTimestamp
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxTransferTestClockAdvanceRequest struct {
	value *SandboxTransferTestClockAdvanceRequest
	isSet bool
}

func (v NullableSandboxTransferTestClockAdvanceRequest) Get() *SandboxTransferTestClockAdvanceRequest {
	return v.value
}

func (v *NullableSandboxTransferTestClockAdvanceRequest) Set(val *SandboxTransferTestClockAdvanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxTransferTestClockAdvanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxTransferTestClockAdvanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxTransferTestClockAdvanceRequest(val *SandboxTransferTestClockAdvanceRequest) *NullableSandboxTransferTestClockAdvanceRequest {
	return &NullableSandboxTransferTestClockAdvanceRequest{value: val, isSet: true}
}

func (v NullableSandboxTransferTestClockAdvanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxTransferTestClockAdvanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


