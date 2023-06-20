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

// NumbersInternational Identifying information for transferring money to or from an international bank account via wire transfer.
type NumbersInternational struct {
	// The Plaid account ID associated with the account numbers
	AccountId string `json:"account_id"`
	// The International Bank Account Number (IBAN) for the account
	Iban string `json:"iban"`
	// The Bank Identifier Code (BIC) for the account
	Bic string `json:"bic"`
	AdditionalProperties map[string]interface{}
}

type _NumbersInternational NumbersInternational

// NewNumbersInternational instantiates a new NumbersInternational object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumbersInternational(accountId string, iban string, bic string) *NumbersInternational {
	this := NumbersInternational{}
	this.AccountId = accountId
	this.Iban = iban
	this.Bic = bic
	return &this
}

// NewNumbersInternationalWithDefaults instantiates a new NumbersInternational object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumbersInternationalWithDefaults() *NumbersInternational {
	this := NumbersInternational{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *NumbersInternational) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *NumbersInternational) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *NumbersInternational) SetAccountId(v string) {
	o.AccountId = v
}

// GetIban returns the Iban field value
func (o *NumbersInternational) GetIban() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Iban
}

// GetIbanOk returns a tuple with the Iban field value
// and a boolean to check if the value has been set.
func (o *NumbersInternational) GetIbanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Iban, true
}

// SetIban sets field value
func (o *NumbersInternational) SetIban(v string) {
	o.Iban = v
}

// GetBic returns the Bic field value
func (o *NumbersInternational) GetBic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bic
}

// GetBicOk returns a tuple with the Bic field value
// and a boolean to check if the value has been set.
func (o *NumbersInternational) GetBicOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Bic, true
}

// SetBic sets field value
func (o *NumbersInternational) SetBic(v string) {
	o.Bic = v
}

func (o NumbersInternational) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NumbersInternational) UnmarshalJSON(bytes []byte) (err error) {
	varNumbersInternational := _NumbersInternational{}

	if err = json.Unmarshal(bytes, &varNumbersInternational); err == nil {
		*o = NumbersInternational(varNumbersInternational)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "iban")
		delete(additionalProperties, "bic")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNumbersInternational struct {
	value *NumbersInternational
	isSet bool
}

func (v NullableNumbersInternational) Get() *NumbersInternational {
	return v.value
}

func (v *NullableNumbersInternational) Set(val *NumbersInternational) {
	v.value = val
	v.isSet = true
}

func (v NullableNumbersInternational) IsSet() bool {
	return v.isSet
}

func (v *NullableNumbersInternational) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumbersInternational(val *NumbersInternational) *NullableNumbersInternational {
	return &NullableNumbersInternational{value: val, isSet: true}
}

func (v NullableNumbersInternational) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumbersInternational) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


