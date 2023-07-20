/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.391.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// RiskCheckBehavior Result summary object specifying values for `behavior` attributes of risk check, when available.
type RiskCheckBehavior struct {
	UserInteractions RiskCheckBehaviorUserInteractionsLabel `json:"user_interactions"`
	FraudRingDetected RiskCheckBehaviorFraudRingDetectedLabel `json:"fraud_ring_detected"`
	BotDetected RiskCheckBehaviorBotDetectedLabel `json:"bot_detected"`
	AdditionalProperties map[string]interface{}
}

type _RiskCheckBehavior RiskCheckBehavior

// NewRiskCheckBehavior instantiates a new RiskCheckBehavior object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskCheckBehavior(userInteractions RiskCheckBehaviorUserInteractionsLabel, fraudRingDetected RiskCheckBehaviorFraudRingDetectedLabel, botDetected RiskCheckBehaviorBotDetectedLabel) *RiskCheckBehavior {
	this := RiskCheckBehavior{}
	this.UserInteractions = userInteractions
	this.FraudRingDetected = fraudRingDetected
	this.BotDetected = botDetected
	return &this
}

// NewRiskCheckBehaviorWithDefaults instantiates a new RiskCheckBehavior object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskCheckBehaviorWithDefaults() *RiskCheckBehavior {
	this := RiskCheckBehavior{}
	return &this
}

// GetUserInteractions returns the UserInteractions field value
func (o *RiskCheckBehavior) GetUserInteractions() RiskCheckBehaviorUserInteractionsLabel {
	if o == nil {
		var ret RiskCheckBehaviorUserInteractionsLabel
		return ret
	}

	return o.UserInteractions
}

// GetUserInteractionsOk returns a tuple with the UserInteractions field value
// and a boolean to check if the value has been set.
func (o *RiskCheckBehavior) GetUserInteractionsOk() (*RiskCheckBehaviorUserInteractionsLabel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserInteractions, true
}

// SetUserInteractions sets field value
func (o *RiskCheckBehavior) SetUserInteractions(v RiskCheckBehaviorUserInteractionsLabel) {
	o.UserInteractions = v
}

// GetFraudRingDetected returns the FraudRingDetected field value
func (o *RiskCheckBehavior) GetFraudRingDetected() RiskCheckBehaviorFraudRingDetectedLabel {
	if o == nil {
		var ret RiskCheckBehaviorFraudRingDetectedLabel
		return ret
	}

	return o.FraudRingDetected
}

// GetFraudRingDetectedOk returns a tuple with the FraudRingDetected field value
// and a boolean to check if the value has been set.
func (o *RiskCheckBehavior) GetFraudRingDetectedOk() (*RiskCheckBehaviorFraudRingDetectedLabel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FraudRingDetected, true
}

// SetFraudRingDetected sets field value
func (o *RiskCheckBehavior) SetFraudRingDetected(v RiskCheckBehaviorFraudRingDetectedLabel) {
	o.FraudRingDetected = v
}

// GetBotDetected returns the BotDetected field value
func (o *RiskCheckBehavior) GetBotDetected() RiskCheckBehaviorBotDetectedLabel {
	if o == nil {
		var ret RiskCheckBehaviorBotDetectedLabel
		return ret
	}

	return o.BotDetected
}

// GetBotDetectedOk returns a tuple with the BotDetected field value
// and a boolean to check if the value has been set.
func (o *RiskCheckBehavior) GetBotDetectedOk() (*RiskCheckBehaviorBotDetectedLabel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BotDetected, true
}

// SetBotDetected sets field value
func (o *RiskCheckBehavior) SetBotDetected(v RiskCheckBehaviorBotDetectedLabel) {
	o.BotDetected = v
}

func (o RiskCheckBehavior) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user_interactions"] = o.UserInteractions
	}
	if true {
		toSerialize["fraud_ring_detected"] = o.FraudRingDetected
	}
	if true {
		toSerialize["bot_detected"] = o.BotDetected
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiskCheckBehavior) UnmarshalJSON(bytes []byte) (err error) {
	varRiskCheckBehavior := _RiskCheckBehavior{}

	if err = json.Unmarshal(bytes, &varRiskCheckBehavior); err == nil {
		*o = RiskCheckBehavior(varRiskCheckBehavior)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "user_interactions")
		delete(additionalProperties, "fraud_ring_detected")
		delete(additionalProperties, "bot_detected")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRiskCheckBehavior struct {
	value *RiskCheckBehavior
	isSet bool
}

func (v NullableRiskCheckBehavior) Get() *RiskCheckBehavior {
	return v.value
}

func (v *NullableRiskCheckBehavior) Set(val *RiskCheckBehavior) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskCheckBehavior) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskCheckBehavior) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskCheckBehavior(val *RiskCheckBehavior) *NullableRiskCheckBehavior {
	return &NullableRiskCheckBehavior{value: val, isSet: true}
}

func (v NullableRiskCheckBehavior) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskCheckBehavior) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


