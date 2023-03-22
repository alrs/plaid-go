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

// SignalPrepareResponse SignalPrepareResponse defines the response schema for `/signal/prepare`
type SignalPrepareResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _SignalPrepareResponse SignalPrepareResponse

// NewSignalPrepareResponse instantiates a new SignalPrepareResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignalPrepareResponse(requestId string) *SignalPrepareResponse {
	this := SignalPrepareResponse{}
	this.RequestId = requestId
	return &this
}

// NewSignalPrepareResponseWithDefaults instantiates a new SignalPrepareResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignalPrepareResponseWithDefaults() *SignalPrepareResponse {
	this := SignalPrepareResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *SignalPrepareResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *SignalPrepareResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *SignalPrepareResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o SignalPrepareResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SignalPrepareResponse) UnmarshalJSON(bytes []byte) (err error) {
	varSignalPrepareResponse := _SignalPrepareResponse{}

	if err = json.Unmarshal(bytes, &varSignalPrepareResponse); err == nil {
		*o = SignalPrepareResponse(varSignalPrepareResponse)
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

type NullableSignalPrepareResponse struct {
	value *SignalPrepareResponse
	isSet bool
}

func (v NullableSignalPrepareResponse) Get() *SignalPrepareResponse {
	return v.value
}

func (v *NullableSignalPrepareResponse) Set(val *SignalPrepareResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSignalPrepareResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSignalPrepareResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignalPrepareResponse(val *SignalPrepareResponse) *NullableSignalPrepareResponse {
	return &NullableSignalPrepareResponse{value: val, isSet: true}
}

func (v NullableSignalPrepareResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignalPrepareResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


