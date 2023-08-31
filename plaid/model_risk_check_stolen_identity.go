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

// RiskCheckStolenIdentity Field containing the data used in determining the outcome of the stolen identity risk check.  Contains the following fields:  `score` - A score from 0 to 100 indicating the likelihood that the user is a stolen identity.
type RiskCheckStolenIdentity struct {
	// A score from 0 to 100 indicating the likelihood that the user is a stolen identity.
	Score int32 `json:"score"`
	AdditionalProperties map[string]interface{}
}

type _RiskCheckStolenIdentity RiskCheckStolenIdentity

// NewRiskCheckStolenIdentity instantiates a new RiskCheckStolenIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskCheckStolenIdentity(score int32) *RiskCheckStolenIdentity {
	this := RiskCheckStolenIdentity{}
	this.Score = score
	return &this
}

// NewRiskCheckStolenIdentityWithDefaults instantiates a new RiskCheckStolenIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskCheckStolenIdentityWithDefaults() *RiskCheckStolenIdentity {
	this := RiskCheckStolenIdentity{}
	return &this
}

// GetScore returns the Score field value
func (o *RiskCheckStolenIdentity) GetScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *RiskCheckStolenIdentity) GetScoreOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *RiskCheckStolenIdentity) SetScore(v int32) {
	o.Score = v
}

func (o RiskCheckStolenIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["score"] = o.Score
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiskCheckStolenIdentity) UnmarshalJSON(bytes []byte) (err error) {
	varRiskCheckStolenIdentity := _RiskCheckStolenIdentity{}

	if err = json.Unmarshal(bytes, &varRiskCheckStolenIdentity); err == nil {
		*o = RiskCheckStolenIdentity(varRiskCheckStolenIdentity)
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

type NullableRiskCheckStolenIdentity struct {
	value *RiskCheckStolenIdentity
	isSet bool
}

func (v NullableRiskCheckStolenIdentity) Get() *RiskCheckStolenIdentity {
	return v.value
}

func (v *NullableRiskCheckStolenIdentity) Set(val *RiskCheckStolenIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskCheckStolenIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskCheckStolenIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskCheckStolenIdentity(val *RiskCheckStolenIdentity) *NullableRiskCheckStolenIdentity {
	return &NullableRiskCheckStolenIdentity{value: val, isSet: true}
}

func (v NullableRiskCheckStolenIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskCheckStolenIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

