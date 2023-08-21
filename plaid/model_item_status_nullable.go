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

// ItemStatusNullable Information about the last successful and failed transactions update for the Item.
type ItemStatusNullable struct {
	Investments NullableItemStatusInvestments `json:"investments,omitempty"`
	Transactions NullableItemStatusTransactions `json:"transactions,omitempty"`
	LastWebhook NullableItemStatusLastWebhook `json:"last_webhook,omitempty"`
}

// NewItemStatusNullable instantiates a new ItemStatusNullable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemStatusNullable() *ItemStatusNullable {
	this := ItemStatusNullable{}
	return &this
}

// NewItemStatusNullableWithDefaults instantiates a new ItemStatusNullable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemStatusNullableWithDefaults() *ItemStatusNullable {
	this := ItemStatusNullable{}
	return &this
}

// GetInvestments returns the Investments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemStatusNullable) GetInvestments() ItemStatusInvestments {
	if o == nil || o.Investments.Get() == nil {
		var ret ItemStatusInvestments
		return ret
	}
	return *o.Investments.Get()
}

// GetInvestmentsOk returns a tuple with the Investments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemStatusNullable) GetInvestmentsOk() (*ItemStatusInvestments, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Investments.Get(), o.Investments.IsSet()
}

// HasInvestments returns a boolean if a field has been set.
func (o *ItemStatusNullable) HasInvestments() bool {
	if o != nil && o.Investments.IsSet() {
		return true
	}

	return false
}

// SetInvestments gets a reference to the given NullableItemStatusInvestments and assigns it to the Investments field.
func (o *ItemStatusNullable) SetInvestments(v ItemStatusInvestments) {
	o.Investments.Set(&v)
}
// SetInvestmentsNil sets the value for Investments to be an explicit nil
func (o *ItemStatusNullable) SetInvestmentsNil() {
	o.Investments.Set(nil)
}

// UnsetInvestments ensures that no value is present for Investments, not even an explicit nil
func (o *ItemStatusNullable) UnsetInvestments() {
	o.Investments.Unset()
}

// GetTransactions returns the Transactions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemStatusNullable) GetTransactions() ItemStatusTransactions {
	if o == nil || o.Transactions.Get() == nil {
		var ret ItemStatusTransactions
		return ret
	}
	return *o.Transactions.Get()
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemStatusNullable) GetTransactionsOk() (*ItemStatusTransactions, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Transactions.Get(), o.Transactions.IsSet()
}

// HasTransactions returns a boolean if a field has been set.
func (o *ItemStatusNullable) HasTransactions() bool {
	if o != nil && o.Transactions.IsSet() {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given NullableItemStatusTransactions and assigns it to the Transactions field.
func (o *ItemStatusNullable) SetTransactions(v ItemStatusTransactions) {
	o.Transactions.Set(&v)
}
// SetTransactionsNil sets the value for Transactions to be an explicit nil
func (o *ItemStatusNullable) SetTransactionsNil() {
	o.Transactions.Set(nil)
}

// UnsetTransactions ensures that no value is present for Transactions, not even an explicit nil
func (o *ItemStatusNullable) UnsetTransactions() {
	o.Transactions.Unset()
}

// GetLastWebhook returns the LastWebhook field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemStatusNullable) GetLastWebhook() ItemStatusLastWebhook {
	if o == nil || o.LastWebhook.Get() == nil {
		var ret ItemStatusLastWebhook
		return ret
	}
	return *o.LastWebhook.Get()
}

// GetLastWebhookOk returns a tuple with the LastWebhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemStatusNullable) GetLastWebhookOk() (*ItemStatusLastWebhook, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastWebhook.Get(), o.LastWebhook.IsSet()
}

// HasLastWebhook returns a boolean if a field has been set.
func (o *ItemStatusNullable) HasLastWebhook() bool {
	if o != nil && o.LastWebhook.IsSet() {
		return true
	}

	return false
}

// SetLastWebhook gets a reference to the given NullableItemStatusLastWebhook and assigns it to the LastWebhook field.
func (o *ItemStatusNullable) SetLastWebhook(v ItemStatusLastWebhook) {
	o.LastWebhook.Set(&v)
}
// SetLastWebhookNil sets the value for LastWebhook to be an explicit nil
func (o *ItemStatusNullable) SetLastWebhookNil() {
	o.LastWebhook.Set(nil)
}

// UnsetLastWebhook ensures that no value is present for LastWebhook, not even an explicit nil
func (o *ItemStatusNullable) UnsetLastWebhook() {
	o.LastWebhook.Unset()
}

func (o ItemStatusNullable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Investments.IsSet() {
		toSerialize["investments"] = o.Investments.Get()
	}
	if o.Transactions.IsSet() {
		toSerialize["transactions"] = o.Transactions.Get()
	}
	if o.LastWebhook.IsSet() {
		toSerialize["last_webhook"] = o.LastWebhook.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableItemStatusNullable struct {
	value *ItemStatusNullable
	isSet bool
}

func (v NullableItemStatusNullable) Get() *ItemStatusNullable {
	return v.value
}

func (v *NullableItemStatusNullable) Set(val *ItemStatusNullable) {
	v.value = val
	v.isSet = true
}

func (v NullableItemStatusNullable) IsSet() bool {
	return v.isSet
}

func (v *NullableItemStatusNullable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemStatusNullable(val *ItemStatusNullable) *NullableItemStatusNullable {
	return &NullableItemStatusNullable{value: val, isSet: true}
}

func (v NullableItemStatusNullable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemStatusNullable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


