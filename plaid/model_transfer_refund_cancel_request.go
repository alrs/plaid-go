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

// TransferRefundCancelRequest Defines the request schema for `/transfer/refund/cancel`
type TransferRefundCancelRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Plaid’s unique identifier for a refund.
	RefundId string `json:"refund_id"`
}

// NewTransferRefundCancelRequest instantiates a new TransferRefundCancelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRefundCancelRequest(refundId string) *TransferRefundCancelRequest {
	this := TransferRefundCancelRequest{}
	this.RefundId = refundId
	return &this
}

// NewTransferRefundCancelRequestWithDefaults instantiates a new TransferRefundCancelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRefundCancelRequestWithDefaults() *TransferRefundCancelRequest {
	this := TransferRefundCancelRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransferRefundCancelRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRefundCancelRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransferRefundCancelRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransferRefundCancelRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransferRefundCancelRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRefundCancelRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransferRefundCancelRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransferRefundCancelRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetRefundId returns the RefundId field value
func (o *TransferRefundCancelRequest) GetRefundId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefundId
}

// GetRefundIdOk returns a tuple with the RefundId field value
// and a boolean to check if the value has been set.
func (o *TransferRefundCancelRequest) GetRefundIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RefundId, true
}

// SetRefundId sets field value
func (o *TransferRefundCancelRequest) SetRefundId(v string) {
	o.RefundId = v
}

func (o TransferRefundCancelRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["refund_id"] = o.RefundId
	}
	return json.Marshal(toSerialize)
}

type NullableTransferRefundCancelRequest struct {
	value *TransferRefundCancelRequest
	isSet bool
}

func (v NullableTransferRefundCancelRequest) Get() *TransferRefundCancelRequest {
	return v.value
}

func (v *NullableTransferRefundCancelRequest) Set(val *TransferRefundCancelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRefundCancelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRefundCancelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRefundCancelRequest(val *TransferRefundCancelRequest) *NullableTransferRefundCancelRequest {
	return &NullableTransferRefundCancelRequest{value: val, isSet: true}
}

func (v NullableTransferRefundCancelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRefundCancelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


