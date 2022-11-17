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
)

// TransferIntentGetResponse Defines the response schema for `/transfer/intent/get`
type TransferIntentGetResponse struct {
	TransferIntent TransferIntentGet `json:"transfer_intent"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferIntentGetResponse TransferIntentGetResponse

// NewTransferIntentGetResponse instantiates a new TransferIntentGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferIntentGetResponse(transferIntent TransferIntentGet, requestId string) *TransferIntentGetResponse {
	this := TransferIntentGetResponse{}
	this.TransferIntent = transferIntent
	this.RequestId = requestId
	return &this
}

// NewTransferIntentGetResponseWithDefaults instantiates a new TransferIntentGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferIntentGetResponseWithDefaults() *TransferIntentGetResponse {
	this := TransferIntentGetResponse{}
	return &this
}

// GetTransferIntent returns the TransferIntent field value
func (o *TransferIntentGetResponse) GetTransferIntent() TransferIntentGet {
	if o == nil {
		var ret TransferIntentGet
		return ret
	}

	return o.TransferIntent
}

// GetTransferIntentOk returns a tuple with the TransferIntent field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGetResponse) GetTransferIntentOk() (*TransferIntentGet, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransferIntent, true
}

// SetTransferIntent sets field value
func (o *TransferIntentGetResponse) SetTransferIntent(v TransferIntentGet) {
	o.TransferIntent = v
}

// GetRequestId returns the RequestId field value
func (o *TransferIntentGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferIntentGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferIntentGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transfer_intent"] = o.TransferIntent
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferIntentGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferIntentGetResponse := _TransferIntentGetResponse{}

	if err = json.Unmarshal(bytes, &varTransferIntentGetResponse); err == nil {
		*o = TransferIntentGetResponse(varTransferIntentGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "transfer_intent")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferIntentGetResponse struct {
	value *TransferIntentGetResponse
	isSet bool
}

func (v NullableTransferIntentGetResponse) Get() *TransferIntentGetResponse {
	return v.value
}

func (v *NullableTransferIntentGetResponse) Set(val *TransferIntentGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferIntentGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferIntentGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferIntentGetResponse(val *TransferIntentGetResponse) *NullableTransferIntentGetResponse {
	return &NullableTransferIntentGetResponse{value: val, isSet: true}
}

func (v NullableTransferIntentGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferIntentGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


