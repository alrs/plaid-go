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
	"fmt"
)

// RiskCheckBehaviorBotDetectedLabel Field describing the outcome of a bot detection behavior risk check.  `yes` indicates that automated activity was detected.  `no` indicates that automated activity was not detected.  `no_data` indicates there was not enough information available to give an accurate signal.
type RiskCheckBehaviorBotDetectedLabel string

var _ = fmt.Printf

// List of RiskCheckBehaviorBotDetectedLabel
const (
	RISKCHECKBEHAVIORBOTDETECTEDLABEL_YES RiskCheckBehaviorBotDetectedLabel = "yes"
	RISKCHECKBEHAVIORBOTDETECTEDLABEL_NO RiskCheckBehaviorBotDetectedLabel = "no"
	RISKCHECKBEHAVIORBOTDETECTEDLABEL_NO_DATA RiskCheckBehaviorBotDetectedLabel = "no_data"
)

var allowedRiskCheckBehaviorBotDetectedLabelEnumValues = []RiskCheckBehaviorBotDetectedLabel{
	"yes",
	"no",
	"no_data",
}

func (v *RiskCheckBehaviorBotDetectedLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := RiskCheckBehaviorBotDetectedLabel(value)


	*v = enumTypeValue
	return nil
}

// NewRiskCheckBehaviorBotDetectedLabelFromValue returns a pointer to a valid RiskCheckBehaviorBotDetectedLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRiskCheckBehaviorBotDetectedLabelFromValue(v string) (*RiskCheckBehaviorBotDetectedLabel, error) {
	ev := RiskCheckBehaviorBotDetectedLabel(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RiskCheckBehaviorBotDetectedLabel) IsValid() bool {
	for _, existing := range allowedRiskCheckBehaviorBotDetectedLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RiskCheckBehaviorBotDetectedLabel value
func (v RiskCheckBehaviorBotDetectedLabel) Ptr() *RiskCheckBehaviorBotDetectedLabel {
	return &v
}

type NullableRiskCheckBehaviorBotDetectedLabel struct {
	value *RiskCheckBehaviorBotDetectedLabel
	isSet bool
}

func (v NullableRiskCheckBehaviorBotDetectedLabel) Get() *RiskCheckBehaviorBotDetectedLabel {
	return v.value
}

func (v *NullableRiskCheckBehaviorBotDetectedLabel) Set(val *RiskCheckBehaviorBotDetectedLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskCheckBehaviorBotDetectedLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskCheckBehaviorBotDetectedLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskCheckBehaviorBotDetectedLabel(val *RiskCheckBehaviorBotDetectedLabel) *NullableRiskCheckBehaviorBotDetectedLabel {
	return &NullableRiskCheckBehaviorBotDetectedLabel{value: val, isSet: true}
}

func (v NullableRiskCheckBehaviorBotDetectedLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskCheckBehaviorBotDetectedLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

