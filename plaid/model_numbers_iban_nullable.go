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

// NumbersIBANNullable International Bank Account Number (IBAN).
type NumbersIBANNullable struct {
}

// NewNumbersIBANNullable instantiates a new NumbersIBANNullable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumbersIBANNullable() *NumbersIBANNullable {
	this := NumbersIBANNullable{}
	return &this
}

// NewNumbersIBANNullableWithDefaults instantiates a new NumbersIBANNullable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumbersIBANNullableWithDefaults() *NumbersIBANNullable {
	this := NumbersIBANNullable{}
	return &this
}

func (o NumbersIBANNullable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableNumbersIBANNullable struct {
	value *NumbersIBANNullable
	isSet bool
}

func (v NullableNumbersIBANNullable) Get() *NumbersIBANNullable {
	return v.value
}

func (v *NullableNumbersIBANNullable) Set(val *NumbersIBANNullable) {
	v.value = val
	v.isSet = true
}

func (v NullableNumbersIBANNullable) IsSet() bool {
	return v.isSet
}

func (v *NullableNumbersIBANNullable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumbersIBANNullable(val *NumbersIBANNullable) *NullableNumbersIBANNullable {
	return &NullableNumbersIBANNullable{value: val, isSet: true}
}

func (v NullableNumbersIBANNullable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumbersIBANNullable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


