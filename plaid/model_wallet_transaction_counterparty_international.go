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

// WalletTransactionCounterpartyInternational International Bank Account Number for a Wallet Transaction
type WalletTransactionCounterpartyInternational struct {
	// International Bank Account Number (IBAN).
	Iban *string `json:"iban,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WalletTransactionCounterpartyInternational WalletTransactionCounterpartyInternational

// NewWalletTransactionCounterpartyInternational instantiates a new WalletTransactionCounterpartyInternational object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletTransactionCounterpartyInternational() *WalletTransactionCounterpartyInternational {
	this := WalletTransactionCounterpartyInternational{}
	return &this
}

// NewWalletTransactionCounterpartyInternationalWithDefaults instantiates a new WalletTransactionCounterpartyInternational object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletTransactionCounterpartyInternationalWithDefaults() *WalletTransactionCounterpartyInternational {
	this := WalletTransactionCounterpartyInternational{}
	return &this
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *WalletTransactionCounterpartyInternational) GetIban() string {
	if o == nil || o.Iban == nil {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletTransactionCounterpartyInternational) GetIbanOk() (*string, bool) {
	if o == nil || o.Iban == nil {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *WalletTransactionCounterpartyInternational) HasIban() bool {
	if o != nil && o.Iban != nil {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *WalletTransactionCounterpartyInternational) SetIban(v string) {
	o.Iban = &v
}

func (o WalletTransactionCounterpartyInternational) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Iban != nil {
		toSerialize["iban"] = o.Iban
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WalletTransactionCounterpartyInternational) UnmarshalJSON(bytes []byte) (err error) {
	varWalletTransactionCounterpartyInternational := _WalletTransactionCounterpartyInternational{}

	if err = json.Unmarshal(bytes, &varWalletTransactionCounterpartyInternational); err == nil {
		*o = WalletTransactionCounterpartyInternational(varWalletTransactionCounterpartyInternational)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "iban")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWalletTransactionCounterpartyInternational struct {
	value *WalletTransactionCounterpartyInternational
	isSet bool
}

func (v NullableWalletTransactionCounterpartyInternational) Get() *WalletTransactionCounterpartyInternational {
	return v.value
}

func (v *NullableWalletTransactionCounterpartyInternational) Set(val *WalletTransactionCounterpartyInternational) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletTransactionCounterpartyInternational) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletTransactionCounterpartyInternational) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletTransactionCounterpartyInternational(val *WalletTransactionCounterpartyInternational) *NullableWalletTransactionCounterpartyInternational {
	return &NullableWalletTransactionCounterpartyInternational{value: val, isSet: true}
}

func (v NullableWalletTransactionCounterpartyInternational) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletTransactionCounterpartyInternational) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


