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

// TransferRefundCancelResponse Defines the response schema for `/transfer/refund/cancel`
type TransferRefundCancelResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferRefundCancelResponse TransferRefundCancelResponse

// NewTransferRefundCancelResponse instantiates a new TransferRefundCancelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRefundCancelResponse(requestId string) *TransferRefundCancelResponse {
	this := TransferRefundCancelResponse{}
	this.RequestId = requestId
	return &this
}

// NewTransferRefundCancelResponseWithDefaults instantiates a new TransferRefundCancelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRefundCancelResponseWithDefaults() *TransferRefundCancelResponse {
	this := TransferRefundCancelResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *TransferRefundCancelResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferRefundCancelResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferRefundCancelResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferRefundCancelResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferRefundCancelResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferRefundCancelResponse := _TransferRefundCancelResponse{}

	if err = json.Unmarshal(bytes, &varTransferRefundCancelResponse); err == nil {
		*o = TransferRefundCancelResponse(varTransferRefundCancelResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferRefundCancelResponse struct {
	value *TransferRefundCancelResponse
	isSet bool
}

func (v NullableTransferRefundCancelResponse) Get() *TransferRefundCancelResponse {
	return v.value
}

func (v *NullableTransferRefundCancelResponse) Set(val *TransferRefundCancelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRefundCancelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRefundCancelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRefundCancelResponse(val *TransferRefundCancelResponse) *NullableTransferRefundCancelResponse {
	return &NullableTransferRefundCancelResponse{value: val, isSet: true}
}

func (v NullableTransferRefundCancelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRefundCancelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


