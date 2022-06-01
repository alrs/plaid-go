/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.121.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// ProcessorIdentityGetResponse ProcessorIdentityGetResponse defines the response schema for `/processor/identity/get`
type ProcessorIdentityGetResponse struct {
	Account AccountIdentity `json:"account"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _ProcessorIdentityGetResponse ProcessorIdentityGetResponse

// NewProcessorIdentityGetResponse instantiates a new ProcessorIdentityGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorIdentityGetResponse(account AccountIdentity, requestId string) *ProcessorIdentityGetResponse {
	this := ProcessorIdentityGetResponse{}
	this.Account = account
	this.RequestId = requestId
	return &this
}

// NewProcessorIdentityGetResponseWithDefaults instantiates a new ProcessorIdentityGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorIdentityGetResponseWithDefaults() *ProcessorIdentityGetResponse {
	this := ProcessorIdentityGetResponse{}
	return &this
}

// GetAccount returns the Account field value
func (o *ProcessorIdentityGetResponse) GetAccount() AccountIdentity {
	if o == nil {
		var ret AccountIdentity
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *ProcessorIdentityGetResponse) GetAccountOk() (*AccountIdentity, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *ProcessorIdentityGetResponse) SetAccount(v AccountIdentity) {
	o.Account = v
}

// GetRequestId returns the RequestId field value
func (o *ProcessorIdentityGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ProcessorIdentityGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ProcessorIdentityGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o ProcessorIdentityGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account"] = o.Account
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProcessorIdentityGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varProcessorIdentityGetResponse := _ProcessorIdentityGetResponse{}

	if err = json.Unmarshal(bytes, &varProcessorIdentityGetResponse); err == nil {
		*o = ProcessorIdentityGetResponse(varProcessorIdentityGetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProcessorIdentityGetResponse struct {
	value *ProcessorIdentityGetResponse
	isSet bool
}

func (v NullableProcessorIdentityGetResponse) Get() *ProcessorIdentityGetResponse {
	return v.value
}

func (v *NullableProcessorIdentityGetResponse) Set(val *ProcessorIdentityGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorIdentityGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorIdentityGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorIdentityGetResponse(val *ProcessorIdentityGetResponse) *NullableProcessorIdentityGetResponse {
	return &NullableProcessorIdentityGetResponse{value: val, isSet: true}
}

func (v NullableProcessorIdentityGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorIdentityGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


