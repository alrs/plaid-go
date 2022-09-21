/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.181.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// CreditRelayRemoveResponse CreditRelayRemoveResponse defines the response schema for `/credit/relay/remove`
type CreditRelayRemoveResponse struct {
	// `true` if the Relay token was successfully removed.
	Removed bool `json:"removed"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _CreditRelayRemoveResponse CreditRelayRemoveResponse

// NewCreditRelayRemoveResponse instantiates a new CreditRelayRemoveResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditRelayRemoveResponse(removed bool, requestId string) *CreditRelayRemoveResponse {
	this := CreditRelayRemoveResponse{}
	this.Removed = removed
	this.RequestId = requestId
	return &this
}

// NewCreditRelayRemoveResponseWithDefaults instantiates a new CreditRelayRemoveResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditRelayRemoveResponseWithDefaults() *CreditRelayRemoveResponse {
	this := CreditRelayRemoveResponse{}
	return &this
}

// GetRemoved returns the Removed field value
func (o *CreditRelayRemoveResponse) GetRemoved() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Removed
}

// GetRemovedOk returns a tuple with the Removed field value
// and a boolean to check if the value has been set.
func (o *CreditRelayRemoveResponse) GetRemovedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Removed, true
}

// SetRemoved sets field value
func (o *CreditRelayRemoveResponse) SetRemoved(v bool) {
	o.Removed = v
}

// GetRequestId returns the RequestId field value
func (o *CreditRelayRemoveResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *CreditRelayRemoveResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *CreditRelayRemoveResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o CreditRelayRemoveResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["removed"] = o.Removed
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditRelayRemoveResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCreditRelayRemoveResponse := _CreditRelayRemoveResponse{}

	if err = json.Unmarshal(bytes, &varCreditRelayRemoveResponse); err == nil {
		*o = CreditRelayRemoveResponse(varCreditRelayRemoveResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "removed")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditRelayRemoveResponse struct {
	value *CreditRelayRemoveResponse
	isSet bool
}

func (v NullableCreditRelayRemoveResponse) Get() *CreditRelayRemoveResponse {
	return v.value
}

func (v *NullableCreditRelayRemoveResponse) Set(val *CreditRelayRemoveResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditRelayRemoveResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditRelayRemoveResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditRelayRemoveResponse(val *CreditRelayRemoveResponse) *NullableCreditRelayRemoveResponse {
	return &NullableCreditRelayRemoveResponse{value: val, isSet: true}
}

func (v NullableCreditRelayRemoveResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditRelayRemoveResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


