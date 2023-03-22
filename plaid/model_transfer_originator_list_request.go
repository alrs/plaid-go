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

// TransferOriginatorListRequest Defines the request schema for `/transfer/originator/list`
type TransferOriginatorListRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The maximum number of originators to return.
	Count NullableInt32 `json:"count,omitempty"`
	// The number of originators to skip before returning results.
	Offset NullableInt32 `json:"offset,omitempty"`
}

// NewTransferOriginatorListRequest instantiates a new TransferOriginatorListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferOriginatorListRequest() *TransferOriginatorListRequest {
	this := TransferOriginatorListRequest{}
	var count int32 = 25
	this.Count = *NewNullableInt32(&count)
	var offset int32 = 0
	this.Offset = *NewNullableInt32(&offset)
	return &this
}

// NewTransferOriginatorListRequestWithDefaults instantiates a new TransferOriginatorListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferOriginatorListRequestWithDefaults() *TransferOriginatorListRequest {
	this := TransferOriginatorListRequest{}
	var count int32 = 25
	this.Count = *NewNullableInt32(&count)
	var offset int32 = 0
	this.Offset = *NewNullableInt32(&offset)
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransferOriginatorListRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOriginatorListRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransferOriginatorListRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransferOriginatorListRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransferOriginatorListRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOriginatorListRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransferOriginatorListRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransferOriginatorListRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetCount returns the Count field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferOriginatorListRequest) GetCount() int32 {
	if o == nil || o.Count.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Count.Get()
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferOriginatorListRequest) GetCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Count.Get(), o.Count.IsSet()
}

// HasCount returns a boolean if a field has been set.
func (o *TransferOriginatorListRequest) HasCount() bool {
	if o != nil && o.Count.IsSet() {
		return true
	}

	return false
}

// SetCount gets a reference to the given NullableInt32 and assigns it to the Count field.
func (o *TransferOriginatorListRequest) SetCount(v int32) {
	o.Count.Set(&v)
}
// SetCountNil sets the value for Count to be an explicit nil
func (o *TransferOriginatorListRequest) SetCountNil() {
	o.Count.Set(nil)
}

// UnsetCount ensures that no value is present for Count, not even an explicit nil
func (o *TransferOriginatorListRequest) UnsetCount() {
	o.Count.Unset()
}

// GetOffset returns the Offset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferOriginatorListRequest) GetOffset() int32 {
	if o == nil || o.Offset.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Offset.Get()
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferOriginatorListRequest) GetOffsetOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Offset.Get(), o.Offset.IsSet()
}

// HasOffset returns a boolean if a field has been set.
func (o *TransferOriginatorListRequest) HasOffset() bool {
	if o != nil && o.Offset.IsSet() {
		return true
	}

	return false
}

// SetOffset gets a reference to the given NullableInt32 and assigns it to the Offset field.
func (o *TransferOriginatorListRequest) SetOffset(v int32) {
	o.Offset.Set(&v)
}
// SetOffsetNil sets the value for Offset to be an explicit nil
func (o *TransferOriginatorListRequest) SetOffsetNil() {
	o.Offset.Set(nil)
}

// UnsetOffset ensures that no value is present for Offset, not even an explicit nil
func (o *TransferOriginatorListRequest) UnsetOffset() {
	o.Offset.Unset()
}

func (o TransferOriginatorListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.Count.IsSet() {
		toSerialize["count"] = o.Count.Get()
	}
	if o.Offset.IsSet() {
		toSerialize["offset"] = o.Offset.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTransferOriginatorListRequest struct {
	value *TransferOriginatorListRequest
	isSet bool
}

func (v NullableTransferOriginatorListRequest) Get() *TransferOriginatorListRequest {
	return v.value
}

func (v *NullableTransferOriginatorListRequest) Set(val *TransferOriginatorListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferOriginatorListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferOriginatorListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferOriginatorListRequest(val *TransferOriginatorListRequest) *NullableTransferOriginatorListRequest {
	return &NullableTransferOriginatorListRequest{value: val, isSet: true}
}

func (v NullableTransferOriginatorListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferOriginatorListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


