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

// TransferCancelResponse Defines the response schema for `/transfer/cancel`
type TransferCancelResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferCancelResponse TransferCancelResponse

// NewTransferCancelResponse instantiates a new TransferCancelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferCancelResponse(requestId string) *TransferCancelResponse {
	this := TransferCancelResponse{}
	this.RequestId = requestId
	return &this
}

// NewTransferCancelResponseWithDefaults instantiates a new TransferCancelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferCancelResponseWithDefaults() *TransferCancelResponse {
	this := TransferCancelResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *TransferCancelResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferCancelResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferCancelResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferCancelResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferCancelResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferCancelResponse := _TransferCancelResponse{}

	if err = json.Unmarshal(bytes, &varTransferCancelResponse); err == nil {
		*o = TransferCancelResponse(varTransferCancelResponse)
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

type NullableTransferCancelResponse struct {
	value *TransferCancelResponse
	isSet bool
}

func (v NullableTransferCancelResponse) Get() *TransferCancelResponse {
	return v.value
}

func (v *NullableTransferCancelResponse) Set(val *TransferCancelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferCancelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferCancelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferCancelResponse(val *TransferCancelResponse) *NullableTransferCancelResponse {
	return &NullableTransferCancelResponse{value: val, isSet: true}
}

func (v NullableTransferCancelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferCancelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


