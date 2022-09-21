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

// CreditEmploymentGetResponse CreditEmploymentGetResponse defines the response schema for `/credit/employment/get`.
type CreditEmploymentGetResponse struct {
	// Array of employment items.
	Items []CreditEmploymentItem `json:"items"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _CreditEmploymentGetResponse CreditEmploymentGetResponse

// NewCreditEmploymentGetResponse instantiates a new CreditEmploymentGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditEmploymentGetResponse(items []CreditEmploymentItem, requestId string) *CreditEmploymentGetResponse {
	this := CreditEmploymentGetResponse{}
	this.Items = items
	this.RequestId = requestId
	return &this
}

// NewCreditEmploymentGetResponseWithDefaults instantiates a new CreditEmploymentGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditEmploymentGetResponseWithDefaults() *CreditEmploymentGetResponse {
	this := CreditEmploymentGetResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *CreditEmploymentGetResponse) GetItems() []CreditEmploymentItem {
	if o == nil {
		var ret []CreditEmploymentItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CreditEmploymentGetResponse) GetItemsOk() (*[]CreditEmploymentItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *CreditEmploymentGetResponse) SetItems(v []CreditEmploymentItem) {
	o.Items = v
}

// GetRequestId returns the RequestId field value
func (o *CreditEmploymentGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *CreditEmploymentGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *CreditEmploymentGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o CreditEmploymentGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditEmploymentGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCreditEmploymentGetResponse := _CreditEmploymentGetResponse{}

	if err = json.Unmarshal(bytes, &varCreditEmploymentGetResponse); err == nil {
		*o = CreditEmploymentGetResponse(varCreditEmploymentGetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditEmploymentGetResponse struct {
	value *CreditEmploymentGetResponse
	isSet bool
}

func (v NullableCreditEmploymentGetResponse) Get() *CreditEmploymentGetResponse {
	return v.value
}

func (v *NullableCreditEmploymentGetResponse) Set(val *CreditEmploymentGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditEmploymentGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditEmploymentGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditEmploymentGetResponse(val *CreditEmploymentGetResponse) *NullableCreditEmploymentGetResponse {
	return &NullableCreditEmploymentGetResponse{value: val, isSet: true}
}

func (v NullableCreditEmploymentGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditEmploymentGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


