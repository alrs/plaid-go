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

// KYCCheckDateOfBirthSummary Result summary object specifying how the `date_of_birth` field matched.
type KYCCheckDateOfBirthSummary struct {
	Summary MatchSummaryCode `json:"summary"`
}

// NewKYCCheckDateOfBirthSummary instantiates a new KYCCheckDateOfBirthSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKYCCheckDateOfBirthSummary(summary MatchSummaryCode) *KYCCheckDateOfBirthSummary {
	this := KYCCheckDateOfBirthSummary{}
	this.Summary = summary
	return &this
}

// NewKYCCheckDateOfBirthSummaryWithDefaults instantiates a new KYCCheckDateOfBirthSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKYCCheckDateOfBirthSummaryWithDefaults() *KYCCheckDateOfBirthSummary {
	this := KYCCheckDateOfBirthSummary{}
	return &this
}

// GetSummary returns the Summary field value
func (o *KYCCheckDateOfBirthSummary) GetSummary() MatchSummaryCode {
	if o == nil {
		var ret MatchSummaryCode
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *KYCCheckDateOfBirthSummary) GetSummaryOk() (*MatchSummaryCode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *KYCCheckDateOfBirthSummary) SetSummary(v MatchSummaryCode) {
	o.Summary = v
}

func (o KYCCheckDateOfBirthSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["summary"] = o.Summary
	}
	return json.Marshal(toSerialize)
}

type NullableKYCCheckDateOfBirthSummary struct {
	value *KYCCheckDateOfBirthSummary
	isSet bool
}

func (v NullableKYCCheckDateOfBirthSummary) Get() *KYCCheckDateOfBirthSummary {
	return v.value
}

func (v *NullableKYCCheckDateOfBirthSummary) Set(val *KYCCheckDateOfBirthSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableKYCCheckDateOfBirthSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableKYCCheckDateOfBirthSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKYCCheckDateOfBirthSummary(val *KYCCheckDateOfBirthSummary) *NullableKYCCheckDateOfBirthSummary {
	return &NullableKYCCheckDateOfBirthSummary{value: val, isSet: true}
}

func (v NullableKYCCheckDateOfBirthSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKYCCheckDateOfBirthSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


