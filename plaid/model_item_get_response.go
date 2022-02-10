/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.64.14
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// ItemGetResponse ItemGetResponse defines the response schema for `/item/get` and `/item/webhook/update`
type ItemGetResponse struct {
	Item Item `json:"item"`
	Status NullableItemStatusNullable `json:"status,omitempty"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _ItemGetResponse ItemGetResponse

// NewItemGetResponse instantiates a new ItemGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemGetResponse(item Item, requestId string) *ItemGetResponse {
	this := ItemGetResponse{}
	this.Item = item
	this.RequestId = requestId
	return &this
}

// NewItemGetResponseWithDefaults instantiates a new ItemGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemGetResponseWithDefaults() *ItemGetResponse {
	this := ItemGetResponse{}
	return &this
}

// GetItem returns the Item field value
func (o *ItemGetResponse) GetItem() Item {
	if o == nil {
		var ret Item
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *ItemGetResponse) GetItemOk() (*Item, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *ItemGetResponse) SetItem(v Item) {
	o.Item = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemGetResponse) GetStatus() ItemStatusNullable {
	if o == nil || o.Status.Get() == nil {
		var ret ItemStatusNullable
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemGetResponse) GetStatusOk() (*ItemStatusNullable, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *ItemGetResponse) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableItemStatusNullable and assigns it to the Status field.
func (o *ItemGetResponse) SetStatus(v ItemStatusNullable) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *ItemGetResponse) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *ItemGetResponse) UnsetStatus() {
	o.Status.Unset()
}

// GetRequestId returns the RequestId field value
func (o *ItemGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ItemGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ItemGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o ItemGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ItemGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varItemGetResponse := _ItemGetResponse{}

	if err = json.Unmarshal(bytes, &varItemGetResponse); err == nil {
		*o = ItemGetResponse(varItemGetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "item")
		delete(additionalProperties, "status")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableItemGetResponse struct {
	value *ItemGetResponse
	isSet bool
}

func (v NullableItemGetResponse) Get() *ItemGetResponse {
	return v.value
}

func (v *NullableItemGetResponse) Set(val *ItemGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableItemGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableItemGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemGetResponse(val *ItemGetResponse) *NullableItemGetResponse {
	return &NullableItemGetResponse{value: val, isSet: true}
}

func (v NullableItemGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


