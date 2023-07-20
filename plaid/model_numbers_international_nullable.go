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

// NumbersInternationalNullable Identifying information for transferring money to or from an international bank account via wire transfer.
type NumbersInternationalNullable struct {
	// The Plaid account ID associated with the account numbers
	AccountId string `json:"account_id"`
	// The International Bank Account Number (IBAN) for the account
	Iban string `json:"iban"`
	// The Bank Identifier Code (BIC) for the account
	Bic string `json:"bic"`
}

// NewNumbersInternationalNullable instantiates a new NumbersInternationalNullable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumbersInternationalNullable(accountId string, iban string, bic string) *NumbersInternationalNullable {
	this := NumbersInternationalNullable{}
	this.AccountId = accountId
	this.Iban = iban
	this.Bic = bic
	return &this
}

// NewNumbersInternationalNullableWithDefaults instantiates a new NumbersInternationalNullable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumbersInternationalNullableWithDefaults() *NumbersInternationalNullable {
	this := NumbersInternationalNullable{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *NumbersInternationalNullable) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *NumbersInternationalNullable) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *NumbersInternationalNullable) SetAccountId(v string) {
	o.AccountId = v
}

// GetIban returns the Iban field value
func (o *NumbersInternationalNullable) GetIban() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Iban
}

// GetIbanOk returns a tuple with the Iban field value
// and a boolean to check if the value has been set.
func (o *NumbersInternationalNullable) GetIbanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Iban, true
}

// SetIban sets field value
func (o *NumbersInternationalNullable) SetIban(v string) {
	o.Iban = v
}

// GetBic returns the Bic field value
func (o *NumbersInternationalNullable) GetBic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bic
}

// GetBicOk returns a tuple with the Bic field value
// and a boolean to check if the value has been set.
func (o *NumbersInternationalNullable) GetBicOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Bic, true
}

// SetBic sets field value
func (o *NumbersInternationalNullable) SetBic(v string) {
	o.Bic = v
}

func (o NumbersInternationalNullable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["iban"] = o.Iban
	}
	if true {
		toSerialize["bic"] = o.Bic
	}
	return json.Marshal(toSerialize)
}

type NullableNumbersInternationalNullable struct {
	value *NumbersInternationalNullable
	isSet bool
}

func (v NullableNumbersInternationalNullable) Get() *NumbersInternationalNullable {
	return v.value
}

func (v *NullableNumbersInternationalNullable) Set(val *NumbersInternationalNullable) {
	v.value = val
	v.isSet = true
}

func (v NullableNumbersInternationalNullable) IsSet() bool {
	return v.isSet
}

func (v *NullableNumbersInternationalNullable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumbersInternationalNullable(val *NumbersInternationalNullable) *NullableNumbersInternationalNullable {
	return &NullableNumbersInternationalNullable{value: val, isSet: true}
}

func (v NullableNumbersInternationalNullable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumbersInternationalNullable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


