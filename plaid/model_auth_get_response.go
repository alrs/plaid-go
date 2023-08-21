/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.419.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// AuthGetResponse AuthGetResponse defines the response schema for `/auth/get`
type AuthGetResponse struct {
	// The `accounts` for which numbers are being retrieved.
	Accounts []AccountBase `json:"accounts"`
	Numbers AuthGetNumbers `json:"numbers"`
	Item Item `json:"item"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _AuthGetResponse AuthGetResponse

// NewAuthGetResponse instantiates a new AuthGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthGetResponse(accounts []AccountBase, numbers AuthGetNumbers, item Item, requestId string) *AuthGetResponse {
	this := AuthGetResponse{}
	this.Accounts = accounts
	this.Numbers = numbers
	this.Item = item
	this.RequestId = requestId
	return &this
}

// NewAuthGetResponseWithDefaults instantiates a new AuthGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthGetResponseWithDefaults() *AuthGetResponse {
	this := AuthGetResponse{}
	return &this
}

// GetAccounts returns the Accounts field value
func (o *AuthGetResponse) GetAccounts() []AccountBase {
	if o == nil {
		var ret []AccountBase
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *AuthGetResponse) GetAccountsOk() (*[]AccountBase, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Accounts, true
}

// SetAccounts sets field value
func (o *AuthGetResponse) SetAccounts(v []AccountBase) {
	o.Accounts = v
}

// GetNumbers returns the Numbers field value
func (o *AuthGetResponse) GetNumbers() AuthGetNumbers {
	if o == nil {
		var ret AuthGetNumbers
		return ret
	}

	return o.Numbers
}

// GetNumbersOk returns a tuple with the Numbers field value
// and a boolean to check if the value has been set.
func (o *AuthGetResponse) GetNumbersOk() (*AuthGetNumbers, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Numbers, true
}

// SetNumbers sets field value
func (o *AuthGetResponse) SetNumbers(v AuthGetNumbers) {
	o.Numbers = v
}

// GetItem returns the Item field value
func (o *AuthGetResponse) GetItem() Item {
	if o == nil {
		var ret Item
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *AuthGetResponse) GetItemOk() (*Item, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *AuthGetResponse) SetItem(v Item) {
	o.Item = v
}

// GetRequestId returns the RequestId field value
func (o *AuthGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *AuthGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *AuthGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o AuthGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accounts"] = o.Accounts
	}
	if true {
		toSerialize["numbers"] = o.Numbers
	}
	if true {
		toSerialize["item"] = o.Item
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAuthGetResponse := _AuthGetResponse{}

	if err = json.Unmarshal(bytes, &varAuthGetResponse); err == nil {
		*o = AuthGetResponse(varAuthGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accounts")
		delete(additionalProperties, "numbers")
		delete(additionalProperties, "item")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthGetResponse struct {
	value *AuthGetResponse
	isSet bool
}

func (v NullableAuthGetResponse) Get() *AuthGetResponse {
	return v.value
}

func (v *NullableAuthGetResponse) Set(val *AuthGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthGetResponse(val *AuthGetResponse) *NullableAuthGetResponse {
	return &NullableAuthGetResponse{value: val, isSet: true}
}

func (v NullableAuthGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


