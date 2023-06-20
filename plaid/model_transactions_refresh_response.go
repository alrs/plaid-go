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

// TransactionsRefreshResponse TransactionsRefreshResponse defines the response schema for `/transactions/refresh`
type TransactionsRefreshResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransactionsRefreshResponse TransactionsRefreshResponse

// NewTransactionsRefreshResponse instantiates a new TransactionsRefreshResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsRefreshResponse(requestId string) *TransactionsRefreshResponse {
	this := TransactionsRefreshResponse{}
	this.RequestId = requestId
	return &this
}

// NewTransactionsRefreshResponseWithDefaults instantiates a new TransactionsRefreshResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsRefreshResponseWithDefaults() *TransactionsRefreshResponse {
	this := TransactionsRefreshResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *TransactionsRefreshResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransactionsRefreshResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransactionsRefreshResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransactionsRefreshResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransactionsRefreshResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransactionsRefreshResponse := _TransactionsRefreshResponse{}

	if err = json.Unmarshal(bytes, &varTransactionsRefreshResponse); err == nil {
		*o = TransactionsRefreshResponse(varTransactionsRefreshResponse)
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

type NullableTransactionsRefreshResponse struct {
	value *TransactionsRefreshResponse
	isSet bool
}

func (v NullableTransactionsRefreshResponse) Get() *TransactionsRefreshResponse {
	return v.value
}

func (v *NullableTransactionsRefreshResponse) Set(val *TransactionsRefreshResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsRefreshResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsRefreshResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsRefreshResponse(val *TransactionsRefreshResponse) *NullableTransactionsRefreshResponse {
	return &NullableTransactionsRefreshResponse{value: val, isSet: true}
}

func (v NullableTransactionsRefreshResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsRefreshResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


