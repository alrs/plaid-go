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

// TransferRecurringCancelResponse Defines the response schema for `/transfer/recurring/cancel`
type TransferRecurringCancelResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferRecurringCancelResponse TransferRecurringCancelResponse

// NewTransferRecurringCancelResponse instantiates a new TransferRecurringCancelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRecurringCancelResponse(requestId string) *TransferRecurringCancelResponse {
	this := TransferRecurringCancelResponse{}
	this.RequestId = requestId
	return &this
}

// NewTransferRecurringCancelResponseWithDefaults instantiates a new TransferRecurringCancelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRecurringCancelResponseWithDefaults() *TransferRecurringCancelResponse {
	this := TransferRecurringCancelResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *TransferRecurringCancelResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferRecurringCancelResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferRecurringCancelResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferRecurringCancelResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferRecurringCancelResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferRecurringCancelResponse := _TransferRecurringCancelResponse{}

	if err = json.Unmarshal(bytes, &varTransferRecurringCancelResponse); err == nil {
		*o = TransferRecurringCancelResponse(varTransferRecurringCancelResponse)
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

type NullableTransferRecurringCancelResponse struct {
	value *TransferRecurringCancelResponse
	isSet bool
}

func (v NullableTransferRecurringCancelResponse) Get() *TransferRecurringCancelResponse {
	return v.value
}

func (v *NullableTransferRecurringCancelResponse) Set(val *TransferRecurringCancelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRecurringCancelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRecurringCancelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRecurringCancelResponse(val *TransferRecurringCancelResponse) *NullableTransferRecurringCancelResponse {
	return &NullableTransferRecurringCancelResponse{value: val, isSet: true}
}

func (v NullableTransferRecurringCancelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRecurringCancelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


