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

// KYCCheckNameSummary Result summary object specifying how the `name` field matched.
type KYCCheckNameSummary struct {
	Summary MatchSummaryCode `json:"summary"`
	AdditionalProperties map[string]interface{}
}

type _KYCCheckNameSummary KYCCheckNameSummary

// NewKYCCheckNameSummary instantiates a new KYCCheckNameSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKYCCheckNameSummary(summary MatchSummaryCode) *KYCCheckNameSummary {
	this := KYCCheckNameSummary{}
	this.Summary = summary
	return &this
}

// NewKYCCheckNameSummaryWithDefaults instantiates a new KYCCheckNameSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKYCCheckNameSummaryWithDefaults() *KYCCheckNameSummary {
	this := KYCCheckNameSummary{}
	return &this
}

// GetSummary returns the Summary field value
func (o *KYCCheckNameSummary) GetSummary() MatchSummaryCode {
	if o == nil {
		var ret MatchSummaryCode
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *KYCCheckNameSummary) GetSummaryOk() (*MatchSummaryCode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *KYCCheckNameSummary) SetSummary(v MatchSummaryCode) {
	o.Summary = v
}

func (o KYCCheckNameSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["summary"] = o.Summary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KYCCheckNameSummary) UnmarshalJSON(bytes []byte) (err error) {
	varKYCCheckNameSummary := _KYCCheckNameSummary{}

	if err = json.Unmarshal(bytes, &varKYCCheckNameSummary); err == nil {
		*o = KYCCheckNameSummary(varKYCCheckNameSummary)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "summary")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKYCCheckNameSummary struct {
	value *KYCCheckNameSummary
	isSet bool
}

func (v NullableKYCCheckNameSummary) Get() *KYCCheckNameSummary {
	return v.value
}

func (v *NullableKYCCheckNameSummary) Set(val *KYCCheckNameSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableKYCCheckNameSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableKYCCheckNameSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKYCCheckNameSummary(val *KYCCheckNameSummary) *NullableKYCCheckNameSummary {
	return &NullableKYCCheckNameSummary{value: val, isSet: true}
}

func (v NullableKYCCheckNameSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKYCCheckNameSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


