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

// IdentityMatchResponse IdentityMatchResponse defines the response schema for `/identity/match`
type IdentityMatchResponse struct {
	// The accounts for which Identity match has been requested
	Accounts []AccountIdentityMatchScore `json:"accounts"`
	Item Item `json:"item"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _IdentityMatchResponse IdentityMatchResponse

// NewIdentityMatchResponse instantiates a new IdentityMatchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityMatchResponse(accounts []AccountIdentityMatchScore, item Item, requestId string) *IdentityMatchResponse {
	this := IdentityMatchResponse{}
	this.Accounts = accounts
	this.Item = item
	this.RequestId = requestId
	return &this
}

// NewIdentityMatchResponseWithDefaults instantiates a new IdentityMatchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityMatchResponseWithDefaults() *IdentityMatchResponse {
	this := IdentityMatchResponse{}
	return &this
}

// GetAccounts returns the Accounts field value
func (o *IdentityMatchResponse) GetAccounts() []AccountIdentityMatchScore {
	if o == nil {
		var ret []AccountIdentityMatchScore
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *IdentityMatchResponse) GetAccountsOk() (*[]AccountIdentityMatchScore, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Accounts, true
}

// SetAccounts sets field value
func (o *IdentityMatchResponse) SetAccounts(v []AccountIdentityMatchScore) {
	o.Accounts = v
}

// GetItem returns the Item field value
func (o *IdentityMatchResponse) GetItem() Item {
	if o == nil {
		var ret Item
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *IdentityMatchResponse) GetItemOk() (*Item, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *IdentityMatchResponse) SetItem(v Item) {
	o.Item = v
}

// GetRequestId returns the RequestId field value
func (o *IdentityMatchResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *IdentityMatchResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *IdentityMatchResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o IdentityMatchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accounts"] = o.Accounts
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

func (o *IdentityMatchResponse) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityMatchResponse := _IdentityMatchResponse{}

	if err = json.Unmarshal(bytes, &varIdentityMatchResponse); err == nil {
		*o = IdentityMatchResponse(varIdentityMatchResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accounts")
		delete(additionalProperties, "item")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityMatchResponse struct {
	value *IdentityMatchResponse
	isSet bool
}

func (v NullableIdentityMatchResponse) Get() *IdentityMatchResponse {
	return v.value
}

func (v *NullableIdentityMatchResponse) Set(val *IdentityMatchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityMatchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityMatchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityMatchResponse(val *IdentityMatchResponse) *NullableIdentityMatchResponse {
	return &NullableIdentityMatchResponse{value: val, isSet: true}
}

func (v NullableIdentityMatchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityMatchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


