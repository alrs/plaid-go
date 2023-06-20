/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.379.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// ProcessorTokenPermissionsSetResponse ProcessorTokenPermissionsSetResponse defines the response schema for `/processor/token/permissions/set`
type ProcessorTokenPermissionsSetResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _ProcessorTokenPermissionsSetResponse ProcessorTokenPermissionsSetResponse

// NewProcessorTokenPermissionsSetResponse instantiates a new ProcessorTokenPermissionsSetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorTokenPermissionsSetResponse(requestId string) *ProcessorTokenPermissionsSetResponse {
	this := ProcessorTokenPermissionsSetResponse{}
	this.RequestId = requestId
	return &this
}

// NewProcessorTokenPermissionsSetResponseWithDefaults instantiates a new ProcessorTokenPermissionsSetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorTokenPermissionsSetResponseWithDefaults() *ProcessorTokenPermissionsSetResponse {
	this := ProcessorTokenPermissionsSetResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *ProcessorTokenPermissionsSetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ProcessorTokenPermissionsSetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ProcessorTokenPermissionsSetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o ProcessorTokenPermissionsSetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProcessorTokenPermissionsSetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varProcessorTokenPermissionsSetResponse := _ProcessorTokenPermissionsSetResponse{}

	if err = json.Unmarshal(bytes, &varProcessorTokenPermissionsSetResponse); err == nil {
		*o = ProcessorTokenPermissionsSetResponse(varProcessorTokenPermissionsSetResponse)
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

type NullableProcessorTokenPermissionsSetResponse struct {
	value *ProcessorTokenPermissionsSetResponse
	isSet bool
}

func (v NullableProcessorTokenPermissionsSetResponse) Get() *ProcessorTokenPermissionsSetResponse {
	return v.value
}

func (v *NullableProcessorTokenPermissionsSetResponse) Set(val *ProcessorTokenPermissionsSetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorTokenPermissionsSetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorTokenPermissionsSetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorTokenPermissionsSetResponse(val *ProcessorTokenPermissionsSetResponse) *NullableProcessorTokenPermissionsSetResponse {
	return &NullableProcessorTokenPermissionsSetResponse{value: val, isSet: true}
}

func (v NullableProcessorTokenPermissionsSetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorTokenPermissionsSetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


