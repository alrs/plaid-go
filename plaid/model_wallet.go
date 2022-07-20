/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.146.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// Wallet An object representing the e-wallet
type Wallet struct {
	// A unique ID identifying the e-wallet
	WalletId string `json:"wallet_id"`
	Balance WalletBalance `json:"balance"`
	Numbers WalletNumbers `json:"numbers"`
	AdditionalProperties map[string]interface{}
}

type _Wallet Wallet

// NewWallet instantiates a new Wallet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWallet(walletId string, balance WalletBalance, numbers WalletNumbers) *Wallet {
	this := Wallet{}
	this.WalletId = walletId
	this.Balance = balance
	this.Numbers = numbers
	return &this
}

// NewWalletWithDefaults instantiates a new Wallet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletWithDefaults() *Wallet {
	this := Wallet{}
	return &this
}

// GetWalletId returns the WalletId field value
func (o *Wallet) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *Wallet) GetWalletIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *Wallet) SetWalletId(v string) {
	o.WalletId = v
}

// GetBalance returns the Balance field value
func (o *Wallet) GetBalance() WalletBalance {
	if o == nil {
		var ret WalletBalance
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *Wallet) GetBalanceOk() (*WalletBalance, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *Wallet) SetBalance(v WalletBalance) {
	o.Balance = v
}

// GetNumbers returns the Numbers field value
func (o *Wallet) GetNumbers() WalletNumbers {
	if o == nil {
		var ret WalletNumbers
		return ret
	}

	return o.Numbers
}

// GetNumbersOk returns a tuple with the Numbers field value
// and a boolean to check if the value has been set.
func (o *Wallet) GetNumbersOk() (*WalletNumbers, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Numbers, true
}

// SetNumbers sets field value
func (o *Wallet) SetNumbers(v WalletNumbers) {
	o.Numbers = v
}

func (o Wallet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["wallet_id"] = o.WalletId
	}
	if true {
		toSerialize["balance"] = o.Balance
	}
	if true {
		toSerialize["numbers"] = o.Numbers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Wallet) UnmarshalJSON(bytes []byte) (err error) {
	varWallet := _Wallet{}

	if err = json.Unmarshal(bytes, &varWallet); err == nil {
		*o = Wallet(varWallet)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "wallet_id")
		delete(additionalProperties, "balance")
		delete(additionalProperties, "numbers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWallet struct {
	value *Wallet
	isSet bool
}

func (v NullableWallet) Get() *Wallet {
	return v.value
}

func (v *NullableWallet) Set(val *Wallet) {
	v.value = val
	v.isSet = true
}

func (v NullableWallet) IsSet() bool {
	return v.isSet
}

func (v *NullableWallet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWallet(val *Wallet) *NullableWallet {
	return &NullableWallet{value: val, isSet: true}
}

func (v NullableWallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWallet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

