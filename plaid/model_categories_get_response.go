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

// CategoriesGetResponse CategoriesGetResponse defines the response schema for `/categories/get`
type CategoriesGetResponse struct {
	// An array of all of the transaction categories used by Plaid.
	Categories []Category `json:"categories"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _CategoriesGetResponse CategoriesGetResponse

// NewCategoriesGetResponse instantiates a new CategoriesGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoriesGetResponse(categories []Category, requestId string) *CategoriesGetResponse {
	this := CategoriesGetResponse{}
	this.Categories = categories
	this.RequestId = requestId
	return &this
}

// NewCategoriesGetResponseWithDefaults instantiates a new CategoriesGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoriesGetResponseWithDefaults() *CategoriesGetResponse {
	this := CategoriesGetResponse{}
	return &this
}

// GetCategories returns the Categories field value
func (o *CategoriesGetResponse) GetCategories() []Category {
	if o == nil {
		var ret []Category
		return ret
	}

	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value
// and a boolean to check if the value has been set.
func (o *CategoriesGetResponse) GetCategoriesOk() (*[]Category, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Categories, true
}

// SetCategories sets field value
func (o *CategoriesGetResponse) SetCategories(v []Category) {
	o.Categories = v
}

// GetRequestId returns the RequestId field value
func (o *CategoriesGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *CategoriesGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *CategoriesGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o CategoriesGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["categories"] = o.Categories
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CategoriesGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCategoriesGetResponse := _CategoriesGetResponse{}

	if err = json.Unmarshal(bytes, &varCategoriesGetResponse); err == nil {
		*o = CategoriesGetResponse(varCategoriesGetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "categories")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCategoriesGetResponse struct {
	value *CategoriesGetResponse
	isSet bool
}

func (v NullableCategoriesGetResponse) Get() *CategoriesGetResponse {
	return v.value
}

func (v *NullableCategoriesGetResponse) Set(val *CategoriesGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoriesGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoriesGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoriesGetResponse(val *CategoriesGetResponse) *NullableCategoriesGetResponse {
	return &NullableCategoriesGetResponse{value: val, isSet: true}
}

func (v NullableCategoriesGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoriesGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


