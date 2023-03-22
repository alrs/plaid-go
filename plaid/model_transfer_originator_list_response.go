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

// TransferOriginatorListResponse Defines the response schema for `/transfer/originator/list`
type TransferOriginatorListResponse struct {
	Originators []Originator `json:"originators"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferOriginatorListResponse TransferOriginatorListResponse

// NewTransferOriginatorListResponse instantiates a new TransferOriginatorListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferOriginatorListResponse(originators []Originator, requestId string) *TransferOriginatorListResponse {
	this := TransferOriginatorListResponse{}
	this.Originators = originators
	this.RequestId = requestId
	return &this
}

// NewTransferOriginatorListResponseWithDefaults instantiates a new TransferOriginatorListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferOriginatorListResponseWithDefaults() *TransferOriginatorListResponse {
	this := TransferOriginatorListResponse{}
	return &this
}

// GetOriginators returns the Originators field value
func (o *TransferOriginatorListResponse) GetOriginators() []Originator {
	if o == nil {
		var ret []Originator
		return ret
	}

	return o.Originators
}

// GetOriginatorsOk returns a tuple with the Originators field value
// and a boolean to check if the value has been set.
func (o *TransferOriginatorListResponse) GetOriginatorsOk() (*[]Originator, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Originators, true
}

// SetOriginators sets field value
func (o *TransferOriginatorListResponse) SetOriginators(v []Originator) {
	o.Originators = v
}

// GetRequestId returns the RequestId field value
func (o *TransferOriginatorListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferOriginatorListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferOriginatorListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferOriginatorListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["originators"] = o.Originators
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferOriginatorListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferOriginatorListResponse := _TransferOriginatorListResponse{}

	if err = json.Unmarshal(bytes, &varTransferOriginatorListResponse); err == nil {
		*o = TransferOriginatorListResponse(varTransferOriginatorListResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "originators")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferOriginatorListResponse struct {
	value *TransferOriginatorListResponse
	isSet bool
}

func (v NullableTransferOriginatorListResponse) Get() *TransferOriginatorListResponse {
	return v.value
}

func (v *NullableTransferOriginatorListResponse) Set(val *TransferOriginatorListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferOriginatorListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferOriginatorListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferOriginatorListResponse(val *TransferOriginatorListResponse) *NullableTransferOriginatorListResponse {
	return &NullableTransferOriginatorListResponse{value: val, isSet: true}
}

func (v NullableTransferOriginatorListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferOriginatorListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


