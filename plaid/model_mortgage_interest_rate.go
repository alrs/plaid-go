/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.79.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// MortgageInterestRate Object containing metadata about the interest rate for the mortgage.
type MortgageInterestRate struct {
	// Percentage value (interest rate of current mortgage, not APR) of interest payable on a loan.
	Percentage NullableFloat64 `json:"percentage"`
	// The type of interest charged (fixed or variable).
	Type NullableString `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _MortgageInterestRate MortgageInterestRate

// NewMortgageInterestRate instantiates a new MortgageInterestRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMortgageInterestRate(percentage NullableFloat64, type_ NullableString) *MortgageInterestRate {
	this := MortgageInterestRate{}
	this.Percentage = percentage
	this.Type = type_
	return &this
}

// NewMortgageInterestRateWithDefaults instantiates a new MortgageInterestRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMortgageInterestRateWithDefaults() *MortgageInterestRate {
	this := MortgageInterestRate{}
	return &this
}

// GetPercentage returns the Percentage field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *MortgageInterestRate) GetPercentage() float64 {
	if o == nil || o.Percentage.Get() == nil {
		var ret float64
		return ret
	}

	return *o.Percentage.Get()
}

// GetPercentageOk returns a tuple with the Percentage field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MortgageInterestRate) GetPercentageOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Percentage.Get(), o.Percentage.IsSet()
}

// SetPercentage sets field value
func (o *MortgageInterestRate) SetPercentage(v float64) {
	o.Percentage.Set(&v)
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MortgageInterestRate) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MortgageInterestRate) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *MortgageInterestRate) SetType(v string) {
	o.Type.Set(&v)
}

func (o MortgageInterestRate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["percentage"] = o.Percentage.Get()
	}
	if true {
		toSerialize["type"] = o.Type.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MortgageInterestRate) UnmarshalJSON(bytes []byte) (err error) {
	varMortgageInterestRate := _MortgageInterestRate{}

	if err = json.Unmarshal(bytes, &varMortgageInterestRate); err == nil {
		*o = MortgageInterestRate(varMortgageInterestRate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "percentage")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMortgageInterestRate struct {
	value *MortgageInterestRate
	isSet bool
}

func (v NullableMortgageInterestRate) Get() *MortgageInterestRate {
	return v.value
}

func (v *NullableMortgageInterestRate) Set(val *MortgageInterestRate) {
	v.value = val
	v.isSet = true
}

func (v NullableMortgageInterestRate) IsSet() bool {
	return v.isSet
}

func (v *NullableMortgageInterestRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMortgageInterestRate(val *MortgageInterestRate) *NullableMortgageInterestRate {
	return &NullableMortgageInterestRate{value: val, isSet: true}
}

func (v NullableMortgageInterestRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMortgageInterestRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


