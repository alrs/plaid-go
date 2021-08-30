/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.26.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// ProcessorNumber An object containing identifying numbers used for making electronic transfers to and from the `account`. The identifying number type (ACH, EFT, IBAN, or BACS) used will depend on the country of the account. An account may have more than one number type. If a particular identifying number type is not used by the `account` for which auth data has been requested, a null value will be returned.
type ProcessorNumber struct {
	Ach NullableNumbersACHNullable `json:"ach,omitempty"`
	Eft NullableNumbersEFTNullable `json:"eft,omitempty"`
	International NullableNumbersInternationalNullable `json:"international,omitempty"`
	Bacs NullableNumbersBACSNullable `json:"bacs,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProcessorNumber ProcessorNumber

// NewProcessorNumber instantiates a new ProcessorNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorNumber() *ProcessorNumber {
	this := ProcessorNumber{}
	return &this
}

// NewProcessorNumberWithDefaults instantiates a new ProcessorNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorNumberWithDefaults() *ProcessorNumber {
	this := ProcessorNumber{}
	return &this
}

// GetAch returns the Ach field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessorNumber) GetAch() NumbersACHNullable {
	if o == nil || o.Ach.Get() == nil {
		var ret NumbersACHNullable
		return ret
	}
	return *o.Ach.Get()
}

// GetAchOk returns a tuple with the Ach field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessorNumber) GetAchOk() (*NumbersACHNullable, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Ach.Get(), o.Ach.IsSet()
}

// HasAch returns a boolean if a field has been set.
func (o *ProcessorNumber) HasAch() bool {
	if o != nil && o.Ach.IsSet() {
		return true
	}

	return false
}

// SetAch gets a reference to the given NullableNumbersACHNullable and assigns it to the Ach field.
func (o *ProcessorNumber) SetAch(v NumbersACHNullable) {
	o.Ach.Set(&v)
}
// SetAchNil sets the value for Ach to be an explicit nil
func (o *ProcessorNumber) SetAchNil() {
	o.Ach.Set(nil)
}

// UnsetAch ensures that no value is present for Ach, not even an explicit nil
func (o *ProcessorNumber) UnsetAch() {
	o.Ach.Unset()
}

// GetEft returns the Eft field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessorNumber) GetEft() NumbersEFTNullable {
	if o == nil || o.Eft.Get() == nil {
		var ret NumbersEFTNullable
		return ret
	}
	return *o.Eft.Get()
}

// GetEftOk returns a tuple with the Eft field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessorNumber) GetEftOk() (*NumbersEFTNullable, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Eft.Get(), o.Eft.IsSet()
}

// HasEft returns a boolean if a field has been set.
func (o *ProcessorNumber) HasEft() bool {
	if o != nil && o.Eft.IsSet() {
		return true
	}

	return false
}

// SetEft gets a reference to the given NullableNumbersEFTNullable and assigns it to the Eft field.
func (o *ProcessorNumber) SetEft(v NumbersEFTNullable) {
	o.Eft.Set(&v)
}
// SetEftNil sets the value for Eft to be an explicit nil
func (o *ProcessorNumber) SetEftNil() {
	o.Eft.Set(nil)
}

// UnsetEft ensures that no value is present for Eft, not even an explicit nil
func (o *ProcessorNumber) UnsetEft() {
	o.Eft.Unset()
}

// GetInternational returns the International field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessorNumber) GetInternational() NumbersInternationalNullable {
	if o == nil || o.International.Get() == nil {
		var ret NumbersInternationalNullable
		return ret
	}
	return *o.International.Get()
}

// GetInternationalOk returns a tuple with the International field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessorNumber) GetInternationalOk() (*NumbersInternationalNullable, bool) {
	if o == nil  {
		return nil, false
	}
	return o.International.Get(), o.International.IsSet()
}

// HasInternational returns a boolean if a field has been set.
func (o *ProcessorNumber) HasInternational() bool {
	if o != nil && o.International.IsSet() {
		return true
	}

	return false
}

// SetInternational gets a reference to the given NullableNumbersInternationalNullable and assigns it to the International field.
func (o *ProcessorNumber) SetInternational(v NumbersInternationalNullable) {
	o.International.Set(&v)
}
// SetInternationalNil sets the value for International to be an explicit nil
func (o *ProcessorNumber) SetInternationalNil() {
	o.International.Set(nil)
}

// UnsetInternational ensures that no value is present for International, not even an explicit nil
func (o *ProcessorNumber) UnsetInternational() {
	o.International.Unset()
}

// GetBacs returns the Bacs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessorNumber) GetBacs() NumbersBACSNullable {
	if o == nil || o.Bacs.Get() == nil {
		var ret NumbersBACSNullable
		return ret
	}
	return *o.Bacs.Get()
}

// GetBacsOk returns a tuple with the Bacs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessorNumber) GetBacsOk() (*NumbersBACSNullable, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Bacs.Get(), o.Bacs.IsSet()
}

// HasBacs returns a boolean if a field has been set.
func (o *ProcessorNumber) HasBacs() bool {
	if o != nil && o.Bacs.IsSet() {
		return true
	}

	return false
}

// SetBacs gets a reference to the given NullableNumbersBACSNullable and assigns it to the Bacs field.
func (o *ProcessorNumber) SetBacs(v NumbersBACSNullable) {
	o.Bacs.Set(&v)
}
// SetBacsNil sets the value for Bacs to be an explicit nil
func (o *ProcessorNumber) SetBacsNil() {
	o.Bacs.Set(nil)
}

// UnsetBacs ensures that no value is present for Bacs, not even an explicit nil
func (o *ProcessorNumber) UnsetBacs() {
	o.Bacs.Unset()
}

func (o ProcessorNumber) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ach.IsSet() {
		toSerialize["ach"] = o.Ach.Get()
	}
	if o.Eft.IsSet() {
		toSerialize["eft"] = o.Eft.Get()
	}
	if o.International.IsSet() {
		toSerialize["international"] = o.International.Get()
	}
	if o.Bacs.IsSet() {
		toSerialize["bacs"] = o.Bacs.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProcessorNumber) UnmarshalJSON(bytes []byte) (err error) {
	varProcessorNumber := _ProcessorNumber{}

	if err = json.Unmarshal(bytes, &varProcessorNumber); err == nil {
		*o = ProcessorNumber(varProcessorNumber)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ach")
		delete(additionalProperties, "eft")
		delete(additionalProperties, "international")
		delete(additionalProperties, "bacs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProcessorNumber struct {
	value *ProcessorNumber
	isSet bool
}

func (v NullableProcessorNumber) Get() *ProcessorNumber {
	return v.value
}

func (v *NullableProcessorNumber) Set(val *ProcessorNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorNumber(val *ProcessorNumber) *NullableProcessorNumber {
	return &NullableProcessorNumber{value: val, isSet: true}
}

func (v NullableProcessorNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


