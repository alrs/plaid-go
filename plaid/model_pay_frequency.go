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

// PayFrequency The frequency of the pay period.
type PayFrequency struct {
	Value PayFrequencyValue `json:"value"`
	VerificationStatus VerificationStatus `json:"verification_status"`
	AdditionalProperties map[string]interface{}
}

type _PayFrequency PayFrequency

// NewPayFrequency instantiates a new PayFrequency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayFrequency(value PayFrequencyValue, verificationStatus VerificationStatus) *PayFrequency {
	this := PayFrequency{}
	this.Value = value
	this.VerificationStatus = verificationStatus
	return &this
}

// NewPayFrequencyWithDefaults instantiates a new PayFrequency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayFrequencyWithDefaults() *PayFrequency {
	this := PayFrequency{}
	return &this
}

// GetValue returns the Value field value
func (o *PayFrequency) GetValue() PayFrequencyValue {
	if o == nil {
		var ret PayFrequencyValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PayFrequency) GetValueOk() (*PayFrequencyValue, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *PayFrequency) SetValue(v PayFrequencyValue) {
	o.Value = v
}

// GetVerificationStatus returns the VerificationStatus field value
func (o *PayFrequency) GetVerificationStatus() VerificationStatus {
	if o == nil {
		var ret VerificationStatus
		return ret
	}

	return o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value
// and a boolean to check if the value has been set.
func (o *PayFrequency) GetVerificationStatusOk() (*VerificationStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerificationStatus, true
}

// SetVerificationStatus sets field value
func (o *PayFrequency) SetVerificationStatus(v VerificationStatus) {
	o.VerificationStatus = v
}

func (o PayFrequency) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["verification_status"] = o.VerificationStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PayFrequency) UnmarshalJSON(bytes []byte) (err error) {
	varPayFrequency := _PayFrequency{}

	if err = json.Unmarshal(bytes, &varPayFrequency); err == nil {
		*o = PayFrequency(varPayFrequency)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "verification_status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePayFrequency struct {
	value *PayFrequency
	isSet bool
}

func (v NullablePayFrequency) Get() *PayFrequency {
	return v.value
}

func (v *NullablePayFrequency) Set(val *PayFrequency) {
	v.value = val
	v.isSet = true
}

func (v NullablePayFrequency) IsSet() bool {
	return v.isSet
}

func (v *NullablePayFrequency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayFrequency(val *PayFrequency) *NullablePayFrequency {
	return &NullablePayFrequency{value: val, isSet: true}
}

func (v NullablePayFrequency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayFrequency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


