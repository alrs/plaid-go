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

// WalletTransactionCounterpartyBACS The account number and sort code of the counterparty's account
type WalletTransactionCounterpartyBACS struct {
	// The account number of the account. Maximum of 10 characters.
	Account *string `json:"account,omitempty"`
	// The 6-character sort code of the account.
	SortCode *string `json:"sort_code,omitempty"`
}

// NewWalletTransactionCounterpartyBACS instantiates a new WalletTransactionCounterpartyBACS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletTransactionCounterpartyBACS() *WalletTransactionCounterpartyBACS {
	this := WalletTransactionCounterpartyBACS{}
	return &this
}

// NewWalletTransactionCounterpartyBACSWithDefaults instantiates a new WalletTransactionCounterpartyBACS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletTransactionCounterpartyBACSWithDefaults() *WalletTransactionCounterpartyBACS {
	this := WalletTransactionCounterpartyBACS{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *WalletTransactionCounterpartyBACS) GetAccount() string {
	if o == nil || o.Account == nil {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletTransactionCounterpartyBACS) GetAccountOk() (*string, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *WalletTransactionCounterpartyBACS) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *WalletTransactionCounterpartyBACS) SetAccount(v string) {
	o.Account = &v
}

// GetSortCode returns the SortCode field value if set, zero value otherwise.
func (o *WalletTransactionCounterpartyBACS) GetSortCode() string {
	if o == nil || o.SortCode == nil {
		var ret string
		return ret
	}
	return *o.SortCode
}

// GetSortCodeOk returns a tuple with the SortCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletTransactionCounterpartyBACS) GetSortCodeOk() (*string, bool) {
	if o == nil || o.SortCode == nil {
		return nil, false
	}
	return o.SortCode, true
}

// HasSortCode returns a boolean if a field has been set.
func (o *WalletTransactionCounterpartyBACS) HasSortCode() bool {
	if o != nil && o.SortCode != nil {
		return true
	}

	return false
}

// SetSortCode gets a reference to the given string and assigns it to the SortCode field.
func (o *WalletTransactionCounterpartyBACS) SetSortCode(v string) {
	o.SortCode = &v
}

func (o WalletTransactionCounterpartyBACS) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.SortCode != nil {
		toSerialize["sort_code"] = o.SortCode
	}
	return json.Marshal(toSerialize)
}

type NullableWalletTransactionCounterpartyBACS struct {
	value *WalletTransactionCounterpartyBACS
	isSet bool
}

func (v NullableWalletTransactionCounterpartyBACS) Get() *WalletTransactionCounterpartyBACS {
	return v.value
}

func (v *NullableWalletTransactionCounterpartyBACS) Set(val *WalletTransactionCounterpartyBACS) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletTransactionCounterpartyBACS) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletTransactionCounterpartyBACS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletTransactionCounterpartyBACS(val *WalletTransactionCounterpartyBACS) *NullableWalletTransactionCounterpartyBACS {
	return &NullableWalletTransactionCounterpartyBACS{value: val, isSet: true}
}

func (v NullableWalletTransactionCounterpartyBACS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletTransactionCounterpartyBACS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


