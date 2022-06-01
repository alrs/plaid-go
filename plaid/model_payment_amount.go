/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.121.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// PaymentAmount The amount and currency of a payment
type PaymentAmount struct {
	Currency PaymentAmountCurrency `json:"currency"`
	// The amount of the payment. Must contain at most two digits of precision e.g. `1.23`. Minimum accepted value is `1`.
	Value float64 `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _PaymentAmount PaymentAmount

// NewPaymentAmount instantiates a new PaymentAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentAmount(currency PaymentAmountCurrency, value float64) *PaymentAmount {
	this := PaymentAmount{}
	this.Currency = currency
	this.Value = value
	return &this
}

// NewPaymentAmountWithDefaults instantiates a new PaymentAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentAmountWithDefaults() *PaymentAmount {
	this := PaymentAmount{}
	return &this
}

// GetCurrency returns the Currency field value
func (o *PaymentAmount) GetCurrency() PaymentAmountCurrency {
	if o == nil {
		var ret PaymentAmountCurrency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PaymentAmount) GetCurrencyOk() (*PaymentAmountCurrency, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PaymentAmount) SetCurrency(v PaymentAmountCurrency) {
	o.Currency = v
}

// GetValue returns the Value field value
func (o *PaymentAmount) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PaymentAmount) GetValueOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *PaymentAmount) SetValue(v float64) {
	o.Value = v
}

func (o PaymentAmount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentAmount) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentAmount := _PaymentAmount{}

	if err = json.Unmarshal(bytes, &varPaymentAmount); err == nil {
		*o = PaymentAmount(varPaymentAmount)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "currency")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentAmount struct {
	value *PaymentAmount
	isSet bool
}

func (v NullablePaymentAmount) Get() *PaymentAmount {
	return v.value
}

func (v *NullablePaymentAmount) Set(val *PaymentAmount) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentAmount) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentAmount(val *PaymentAmount) *NullablePaymentAmount {
	return &NullablePaymentAmount{value: val, isSet: true}
}

func (v NullablePaymentAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


