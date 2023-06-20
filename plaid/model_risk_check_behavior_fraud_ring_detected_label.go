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
	"fmt"
)

// RiskCheckBehaviorFraudRingDetectedLabel Field describing the outcome of a fraud ring behavior risk check.  `yes` indicates that fraud ring activity was detected.  `no` indicates that fraud ring activity was not detected.  `no_data` indicates there was not enough information available to give an accurate signal.
type RiskCheckBehaviorFraudRingDetectedLabel string

var _ = fmt.Printf

// List of RiskCheckBehaviorFraudRingDetectedLabel
const (
	RISKCHECKBEHAVIORFRAUDRINGDETECTEDLABEL_YES RiskCheckBehaviorFraudRingDetectedLabel = "yes"
	RISKCHECKBEHAVIORFRAUDRINGDETECTEDLABEL_NO RiskCheckBehaviorFraudRingDetectedLabel = "no"
	RISKCHECKBEHAVIORFRAUDRINGDETECTEDLABEL_NO_DATA RiskCheckBehaviorFraudRingDetectedLabel = "no_data"
)

var allowedRiskCheckBehaviorFraudRingDetectedLabelEnumValues = []RiskCheckBehaviorFraudRingDetectedLabel{
	"yes",
	"no",
	"no_data",
}

func (v *RiskCheckBehaviorFraudRingDetectedLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := RiskCheckBehaviorFraudRingDetectedLabel(value)


	*v = enumTypeValue
	return nil
}

// NewRiskCheckBehaviorFraudRingDetectedLabelFromValue returns a pointer to a valid RiskCheckBehaviorFraudRingDetectedLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRiskCheckBehaviorFraudRingDetectedLabelFromValue(v string) (*RiskCheckBehaviorFraudRingDetectedLabel, error) {
	ev := RiskCheckBehaviorFraudRingDetectedLabel(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RiskCheckBehaviorFraudRingDetectedLabel) IsValid() bool {
	for _, existing := range allowedRiskCheckBehaviorFraudRingDetectedLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RiskCheckBehaviorFraudRingDetectedLabel value
func (v RiskCheckBehaviorFraudRingDetectedLabel) Ptr() *RiskCheckBehaviorFraudRingDetectedLabel {
	return &v
}

type NullableRiskCheckBehaviorFraudRingDetectedLabel struct {
	value *RiskCheckBehaviorFraudRingDetectedLabel
	isSet bool
}

func (v NullableRiskCheckBehaviorFraudRingDetectedLabel) Get() *RiskCheckBehaviorFraudRingDetectedLabel {
	return v.value
}

func (v *NullableRiskCheckBehaviorFraudRingDetectedLabel) Set(val *RiskCheckBehaviorFraudRingDetectedLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskCheckBehaviorFraudRingDetectedLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskCheckBehaviorFraudRingDetectedLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskCheckBehaviorFraudRingDetectedLabel(val *RiskCheckBehaviorFraudRingDetectedLabel) *NullableRiskCheckBehaviorFraudRingDetectedLabel {
	return &NullableRiskCheckBehaviorFraudRingDetectedLabel{value: val, isSet: true}
}

func (v NullableRiskCheckBehaviorFraudRingDetectedLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskCheckBehaviorFraudRingDetectedLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

