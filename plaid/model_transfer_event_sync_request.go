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

// TransferEventSyncRequest Defines the request schema for `/transfer/event/sync`
type TransferEventSyncRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The latest (largest) `event_id` fetched via the sync endpoint, or 0 initially.
	AfterId int32 `json:"after_id"`
	// The maximum number of transfer events to return.
	Count NullableInt32 `json:"count,omitempty"`
}

// NewTransferEventSyncRequest instantiates a new TransferEventSyncRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferEventSyncRequest(afterId int32) *TransferEventSyncRequest {
	this := TransferEventSyncRequest{}
	this.AfterId = afterId
	var count int32 = 25
	this.Count = *NewNullableInt32(&count)
	return &this
}

// NewTransferEventSyncRequestWithDefaults instantiates a new TransferEventSyncRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferEventSyncRequestWithDefaults() *TransferEventSyncRequest {
	this := TransferEventSyncRequest{}
	var count int32 = 25
	this.Count = *NewNullableInt32(&count)
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransferEventSyncRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEventSyncRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransferEventSyncRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransferEventSyncRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransferEventSyncRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEventSyncRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransferEventSyncRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransferEventSyncRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAfterId returns the AfterId field value
func (o *TransferEventSyncRequest) GetAfterId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AfterId
}

// GetAfterIdOk returns a tuple with the AfterId field value
// and a boolean to check if the value has been set.
func (o *TransferEventSyncRequest) GetAfterIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AfterId, true
}

// SetAfterId sets field value
func (o *TransferEventSyncRequest) SetAfterId(v int32) {
	o.AfterId = v
}

// GetCount returns the Count field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferEventSyncRequest) GetCount() int32 {
	if o == nil || o.Count.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Count.Get()
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferEventSyncRequest) GetCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Count.Get(), o.Count.IsSet()
}

// HasCount returns a boolean if a field has been set.
func (o *TransferEventSyncRequest) HasCount() bool {
	if o != nil && o.Count.IsSet() {
		return true
	}

	return false
}

// SetCount gets a reference to the given NullableInt32 and assigns it to the Count field.
func (o *TransferEventSyncRequest) SetCount(v int32) {
	o.Count.Set(&v)
}
// SetCountNil sets the value for Count to be an explicit nil
func (o *TransferEventSyncRequest) SetCountNil() {
	o.Count.Set(nil)
}

// UnsetCount ensures that no value is present for Count, not even an explicit nil
func (o *TransferEventSyncRequest) UnsetCount() {
	o.Count.Unset()
}

func (o TransferEventSyncRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["after_id"] = o.AfterId
	}
	if o.Count.IsSet() {
		toSerialize["count"] = o.Count.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTransferEventSyncRequest struct {
	value *TransferEventSyncRequest
	isSet bool
}

func (v NullableTransferEventSyncRequest) Get() *TransferEventSyncRequest {
	return v.value
}

func (v *NullableTransferEventSyncRequest) Set(val *TransferEventSyncRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferEventSyncRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferEventSyncRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferEventSyncRequest(val *TransferEventSyncRequest) *NullableTransferEventSyncRequest {
	return &NullableTransferEventSyncRequest{value: val, isSet: true}
}

func (v NullableTransferEventSyncRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferEventSyncRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


