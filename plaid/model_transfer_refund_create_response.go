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

// TransferRefundCreateResponse Defines the response schema for `/transfer/refund/create`
type TransferRefundCreateResponse struct {
	Refund TransferRefund `json:"refund"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferRefundCreateResponse TransferRefundCreateResponse

// NewTransferRefundCreateResponse instantiates a new TransferRefundCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRefundCreateResponse(refund TransferRefund, requestId string) *TransferRefundCreateResponse {
	this := TransferRefundCreateResponse{}
	this.Refund = refund
	this.RequestId = requestId
	return &this
}

// NewTransferRefundCreateResponseWithDefaults instantiates a new TransferRefundCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRefundCreateResponseWithDefaults() *TransferRefundCreateResponse {
	this := TransferRefundCreateResponse{}
	return &this
}

// GetRefund returns the Refund field value
func (o *TransferRefundCreateResponse) GetRefund() TransferRefund {
	if o == nil {
		var ret TransferRefund
		return ret
	}

	return o.Refund
}

// GetRefundOk returns a tuple with the Refund field value
// and a boolean to check if the value has been set.
func (o *TransferRefundCreateResponse) GetRefundOk() (*TransferRefund, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Refund, true
}

// SetRefund sets field value
func (o *TransferRefundCreateResponse) SetRefund(v TransferRefund) {
	o.Refund = v
}

// GetRequestId returns the RequestId field value
func (o *TransferRefundCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferRefundCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferRefundCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferRefundCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["refund"] = o.Refund
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferRefundCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferRefundCreateResponse := _TransferRefundCreateResponse{}

	if err = json.Unmarshal(bytes, &varTransferRefundCreateResponse); err == nil {
		*o = TransferRefundCreateResponse(varTransferRefundCreateResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "refund")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferRefundCreateResponse struct {
	value *TransferRefundCreateResponse
	isSet bool
}

func (v NullableTransferRefundCreateResponse) Get() *TransferRefundCreateResponse {
	return v.value
}

func (v *NullableTransferRefundCreateResponse) Set(val *TransferRefundCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRefundCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRefundCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRefundCreateResponse(val *TransferRefundCreateResponse) *NullableTransferRefundCreateResponse {
	return &NullableTransferRefundCreateResponse{value: val, isSet: true}
}

func (v NullableTransferRefundCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRefundCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


