/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.205.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// NumbersBACS Identifying information for transferring money to or from a UK bank account via BACS.
type NumbersBACS struct {
	// The Plaid account ID associated with the account numbers
	AccountId string `json:"account_id"`
	// The BACS account number for the account
	Account string `json:"account"`
	// The BACS sort code for the account
	SortCode string `json:"sort_code"`
	AdditionalProperties map[string]interface{}
}

type _NumbersBACS NumbersBACS

// NewNumbersBACS instantiates a new NumbersBACS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumbersBACS(accountId string, account string, sortCode string) *NumbersBACS {
	this := NumbersBACS{}
	this.AccountId = accountId
	this.Account = account
	this.SortCode = sortCode
	return &this
}

// NewNumbersBACSWithDefaults instantiates a new NumbersBACS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumbersBACSWithDefaults() *NumbersBACS {
	this := NumbersBACS{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *NumbersBACS) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *NumbersBACS) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *NumbersBACS) SetAccountId(v string) {
	o.AccountId = v
}

// GetAccount returns the Account field value
func (o *NumbersBACS) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *NumbersBACS) GetAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *NumbersBACS) SetAccount(v string) {
	o.Account = v
}

// GetSortCode returns the SortCode field value
func (o *NumbersBACS) GetSortCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SortCode
}

// GetSortCodeOk returns a tuple with the SortCode field value
// and a boolean to check if the value has been set.
func (o *NumbersBACS) GetSortCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SortCode, true
}

// SetSortCode sets field value
func (o *NumbersBACS) SetSortCode(v string) {
	o.SortCode = v
}

func (o NumbersBACS) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["account"] = o.Account
	}
	if true {
		toSerialize["sort_code"] = o.SortCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NumbersBACS) UnmarshalJSON(bytes []byte) (err error) {
	varNumbersBACS := _NumbersBACS{}

	if err = json.Unmarshal(bytes, &varNumbersBACS); err == nil {
		*o = NumbersBACS(varNumbersBACS)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "account")
		delete(additionalProperties, "sort_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNumbersBACS struct {
	value *NumbersBACS
	isSet bool
}

func (v NullableNumbersBACS) Get() *NumbersBACS {
	return v.value
}

func (v *NullableNumbersBACS) Set(val *NumbersBACS) {
	v.value = val
	v.isSet = true
}

func (v NullableNumbersBACS) IsSet() bool {
	return v.isSet
}

func (v *NullableNumbersBACS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumbersBACS(val *NumbersBACS) *NullableNumbersBACS {
	return &NullableNumbersBACS{value: val, isSet: true}
}

func (v NullableNumbersBACS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumbersBACS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


