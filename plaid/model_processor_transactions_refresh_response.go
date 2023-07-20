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

// ProcessorTransactionsRefreshResponse ProcessorTransactionsRefreshResponse defines the response schema for `/processor/transactions/refresh`
type ProcessorTransactionsRefreshResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _ProcessorTransactionsRefreshResponse ProcessorTransactionsRefreshResponse

// NewProcessorTransactionsRefreshResponse instantiates a new ProcessorTransactionsRefreshResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorTransactionsRefreshResponse(requestId string) *ProcessorTransactionsRefreshResponse {
	this := ProcessorTransactionsRefreshResponse{}
	this.RequestId = requestId
	return &this
}

// NewProcessorTransactionsRefreshResponseWithDefaults instantiates a new ProcessorTransactionsRefreshResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorTransactionsRefreshResponseWithDefaults() *ProcessorTransactionsRefreshResponse {
	this := ProcessorTransactionsRefreshResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *ProcessorTransactionsRefreshResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ProcessorTransactionsRefreshResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ProcessorTransactionsRefreshResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o ProcessorTransactionsRefreshResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProcessorTransactionsRefreshResponse) UnmarshalJSON(bytes []byte) (err error) {
	varProcessorTransactionsRefreshResponse := _ProcessorTransactionsRefreshResponse{}

	if err = json.Unmarshal(bytes, &varProcessorTransactionsRefreshResponse); err == nil {
		*o = ProcessorTransactionsRefreshResponse(varProcessorTransactionsRefreshResponse)
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

type NullableProcessorTransactionsRefreshResponse struct {
	value *ProcessorTransactionsRefreshResponse
	isSet bool
}

func (v NullableProcessorTransactionsRefreshResponse) Get() *ProcessorTransactionsRefreshResponse {
	return v.value
}

func (v *NullableProcessorTransactionsRefreshResponse) Set(val *ProcessorTransactionsRefreshResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorTransactionsRefreshResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorTransactionsRefreshResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorTransactionsRefreshResponse(val *ProcessorTransactionsRefreshResponse) *NullableProcessorTransactionsRefreshResponse {
	return &NullableProcessorTransactionsRefreshResponse{value: val, isSet: true}
}

func (v NullableProcessorTransactionsRefreshResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorTransactionsRefreshResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


