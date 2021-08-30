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

// ProjectedIncomeSummaryFieldNumber struct for ProjectedIncomeSummaryFieldNumber
type ProjectedIncomeSummaryFieldNumber struct {
	// The value of the field.
	Value float32 `json:"value"`
	VerificationStatus VerificationStatus `json:"verification_status"`
}

// NewProjectedIncomeSummaryFieldNumber instantiates a new ProjectedIncomeSummaryFieldNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectedIncomeSummaryFieldNumber(value float32, verificationStatus VerificationStatus) *ProjectedIncomeSummaryFieldNumber {
	this := ProjectedIncomeSummaryFieldNumber{}
	this.Value = value
	this.VerificationStatus = verificationStatus
	return &this
}

// NewProjectedIncomeSummaryFieldNumberWithDefaults instantiates a new ProjectedIncomeSummaryFieldNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectedIncomeSummaryFieldNumberWithDefaults() *ProjectedIncomeSummaryFieldNumber {
	this := ProjectedIncomeSummaryFieldNumber{}
	return &this
}

// GetValue returns the Value field value
func (o *ProjectedIncomeSummaryFieldNumber) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ProjectedIncomeSummaryFieldNumber) GetValueOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ProjectedIncomeSummaryFieldNumber) SetValue(v float32) {
	o.Value = v
}

// GetVerificationStatus returns the VerificationStatus field value
func (o *ProjectedIncomeSummaryFieldNumber) GetVerificationStatus() VerificationStatus {
	if o == nil {
		var ret VerificationStatus
		return ret
	}

	return o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value
// and a boolean to check if the value has been set.
func (o *ProjectedIncomeSummaryFieldNumber) GetVerificationStatusOk() (*VerificationStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerificationStatus, true
}

// SetVerificationStatus sets field value
func (o *ProjectedIncomeSummaryFieldNumber) SetVerificationStatus(v VerificationStatus) {
	o.VerificationStatus = v
}

func (o ProjectedIncomeSummaryFieldNumber) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["verification_status"] = o.VerificationStatus
	}
	return json.Marshal(toSerialize)
}

type NullableProjectedIncomeSummaryFieldNumber struct {
	value *ProjectedIncomeSummaryFieldNumber
	isSet bool
}

func (v NullableProjectedIncomeSummaryFieldNumber) Get() *ProjectedIncomeSummaryFieldNumber {
	return v.value
}

func (v *NullableProjectedIncomeSummaryFieldNumber) Set(val *ProjectedIncomeSummaryFieldNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectedIncomeSummaryFieldNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectedIncomeSummaryFieldNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectedIncomeSummaryFieldNumber(val *ProjectedIncomeSummaryFieldNumber) *NullableProjectedIncomeSummaryFieldNumber {
	return &NullableProjectedIncomeSummaryFieldNumber{value: val, isSet: true}
}

func (v NullableProjectedIncomeSummaryFieldNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectedIncomeSummaryFieldNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


