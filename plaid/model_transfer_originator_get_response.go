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

// TransferOriginatorGetResponse Defines the response schema for `/transfer/originator/get`
type TransferOriginatorGetResponse struct {
	Originator DetailedOriginator `json:"originator"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferOriginatorGetResponse TransferOriginatorGetResponse

// NewTransferOriginatorGetResponse instantiates a new TransferOriginatorGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferOriginatorGetResponse(originator DetailedOriginator, requestId string) *TransferOriginatorGetResponse {
	this := TransferOriginatorGetResponse{}
	this.Originator = originator
	this.RequestId = requestId
	return &this
}

// NewTransferOriginatorGetResponseWithDefaults instantiates a new TransferOriginatorGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferOriginatorGetResponseWithDefaults() *TransferOriginatorGetResponse {
	this := TransferOriginatorGetResponse{}
	return &this
}

// GetOriginator returns the Originator field value
func (o *TransferOriginatorGetResponse) GetOriginator() DetailedOriginator {
	if o == nil {
		var ret DetailedOriginator
		return ret
	}

	return o.Originator
}

// GetOriginatorOk returns a tuple with the Originator field value
// and a boolean to check if the value has been set.
func (o *TransferOriginatorGetResponse) GetOriginatorOk() (*DetailedOriginator, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Originator, true
}

// SetOriginator sets field value
func (o *TransferOriginatorGetResponse) SetOriginator(v DetailedOriginator) {
	o.Originator = v
}

// GetRequestId returns the RequestId field value
func (o *TransferOriginatorGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferOriginatorGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferOriginatorGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferOriginatorGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["originator"] = o.Originator
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferOriginatorGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferOriginatorGetResponse := _TransferOriginatorGetResponse{}

	if err = json.Unmarshal(bytes, &varTransferOriginatorGetResponse); err == nil {
		*o = TransferOriginatorGetResponse(varTransferOriginatorGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "originator")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferOriginatorGetResponse struct {
	value *TransferOriginatorGetResponse
	isSet bool
}

func (v NullableTransferOriginatorGetResponse) Get() *TransferOriginatorGetResponse {
	return v.value
}

func (v *NullableTransferOriginatorGetResponse) Set(val *TransferOriginatorGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferOriginatorGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferOriginatorGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferOriginatorGetResponse(val *TransferOriginatorGetResponse) *NullableTransferOriginatorGetResponse {
	return &NullableTransferOriginatorGetResponse{value: val, isSet: true}
}

func (v NullableTransferOriginatorGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferOriginatorGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


