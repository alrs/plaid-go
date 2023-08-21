/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.419.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// RiskCheckSyntheticIdentity Field containing the data used in determining the outcome of the synthetic identity risk check.  Contains the following fields:  `score` - A score from 0 to 100 indicating the likelihood that the user is a synthetic identity.
type RiskCheckSyntheticIdentity struct {
	// A score from 0 to 100 indicating the likelihood that the user is a synthetic identity.
	Score int32 `json:"score"`
	AdditionalProperties map[string]interface{}
}

type _RiskCheckSyntheticIdentity RiskCheckSyntheticIdentity

// NewRiskCheckSyntheticIdentity instantiates a new RiskCheckSyntheticIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskCheckSyntheticIdentity(score int32) *RiskCheckSyntheticIdentity {
	this := RiskCheckSyntheticIdentity{}
	this.Score = score
	return &this
}

// NewRiskCheckSyntheticIdentityWithDefaults instantiates a new RiskCheckSyntheticIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskCheckSyntheticIdentityWithDefaults() *RiskCheckSyntheticIdentity {
	this := RiskCheckSyntheticIdentity{}
	return &this
}

// GetScore returns the Score field value
func (o *RiskCheckSyntheticIdentity) GetScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *RiskCheckSyntheticIdentity) GetScoreOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *RiskCheckSyntheticIdentity) SetScore(v int32) {
	o.Score = v
}

func (o RiskCheckSyntheticIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["score"] = o.Score
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiskCheckSyntheticIdentity) UnmarshalJSON(bytes []byte) (err error) {
	varRiskCheckSyntheticIdentity := _RiskCheckSyntheticIdentity{}

	if err = json.Unmarshal(bytes, &varRiskCheckSyntheticIdentity); err == nil {
		*o = RiskCheckSyntheticIdentity(varRiskCheckSyntheticIdentity)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "score")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRiskCheckSyntheticIdentity struct {
	value *RiskCheckSyntheticIdentity
	isSet bool
}

func (v NullableRiskCheckSyntheticIdentity) Get() *RiskCheckSyntheticIdentity {
	return v.value
}

func (v *NullableRiskCheckSyntheticIdentity) Set(val *RiskCheckSyntheticIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskCheckSyntheticIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskCheckSyntheticIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskCheckSyntheticIdentity(val *RiskCheckSyntheticIdentity) *NullableRiskCheckSyntheticIdentity {
	return &NullableRiskCheckSyntheticIdentity{value: val, isSet: true}
}

func (v NullableRiskCheckSyntheticIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskCheckSyntheticIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


