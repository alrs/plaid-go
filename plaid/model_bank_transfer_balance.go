/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.343.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// BankTransferBalance Information about the balance of a bank transfer
type BankTransferBalance struct {
	// The total available balance - the sum of all successful debit transfer amounts minus all credit transfer amounts.
	Available string `json:"available"`
	// The transactable balance shows the amount in your account that you are able to use for transfers, and is essentially your available balance minus your minimum balance.
	Transactable string `json:"transactable"`
	AdditionalProperties map[string]interface{}
}

type _BankTransferBalance BankTransferBalance

// NewBankTransferBalance instantiates a new BankTransferBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferBalance(available string, transactable string) *BankTransferBalance {
	this := BankTransferBalance{}
	this.Available = available
	this.Transactable = transactable
	return &this
}

// NewBankTransferBalanceWithDefaults instantiates a new BankTransferBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferBalanceWithDefaults() *BankTransferBalance {
	this := BankTransferBalance{}
	return &this
}

// GetAvailable returns the Available field value
func (o *BankTransferBalance) GetAvailable() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value
// and a boolean to check if the value has been set.
func (o *BankTransferBalance) GetAvailableOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Available, true
}

// SetAvailable sets field value
func (o *BankTransferBalance) SetAvailable(v string) {
	o.Available = v
}

// GetTransactable returns the Transactable field value
func (o *BankTransferBalance) GetTransactable() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Transactable
}

// GetTransactableOk returns a tuple with the Transactable field value
// and a boolean to check if the value has been set.
func (o *BankTransferBalance) GetTransactableOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Transactable, true
}

// SetTransactable sets field value
func (o *BankTransferBalance) SetTransactable(v string) {
	o.Transactable = v
}

func (o BankTransferBalance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["available"] = o.Available
	}
	if true {
		toSerialize["transactable"] = o.Transactable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BankTransferBalance) UnmarshalJSON(bytes []byte) (err error) {
	varBankTransferBalance := _BankTransferBalance{}

	if err = json.Unmarshal(bytes, &varBankTransferBalance); err == nil {
		*o = BankTransferBalance(varBankTransferBalance)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "available")
		delete(additionalProperties, "transactable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBankTransferBalance struct {
	value *BankTransferBalance
	isSet bool
}

func (v NullableBankTransferBalance) Get() *BankTransferBalance {
	return v.value
}

func (v *NullableBankTransferBalance) Set(val *BankTransferBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferBalance(val *BankTransferBalance) *NullableBankTransferBalance {
	return &NullableBankTransferBalance{value: val, isSet: true}
}

func (v NullableBankTransferBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


