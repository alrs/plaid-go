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

// IncomeVerificationDocParsingConfig Analysis options to enable for document parsing
type IncomeVerificationDocParsingConfig string

var _ = fmt.Printf

// List of IncomeVerificationDocParsingConfig
const (
	INCOMEVERIFICATIONDOCPARSINGCONFIG_OCR IncomeVerificationDocParsingConfig = "ocr"
	INCOMEVERIFICATIONDOCPARSINGCONFIG_FRAUD_RISK IncomeVerificationDocParsingConfig = "fraud_risk"
)

var allowedIncomeVerificationDocParsingConfigEnumValues = []IncomeVerificationDocParsingConfig{
	"ocr",
	"fraud_risk",
}

func (v *IncomeVerificationDocParsingConfig) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := IncomeVerificationDocParsingConfig(value)


	*v = enumTypeValue
	return nil
}

// NewIncomeVerificationDocParsingConfigFromValue returns a pointer to a valid IncomeVerificationDocParsingConfig
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIncomeVerificationDocParsingConfigFromValue(v string) (*IncomeVerificationDocParsingConfig, error) {
	ev := IncomeVerificationDocParsingConfig(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IncomeVerificationDocParsingConfig) IsValid() bool {
	for _, existing := range allowedIncomeVerificationDocParsingConfigEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncomeVerificationDocParsingConfig value
func (v IncomeVerificationDocParsingConfig) Ptr() *IncomeVerificationDocParsingConfig {
	return &v
}

type NullableIncomeVerificationDocParsingConfig struct {
	value *IncomeVerificationDocParsingConfig
	isSet bool
}

func (v NullableIncomeVerificationDocParsingConfig) Get() *IncomeVerificationDocParsingConfig {
	return v.value
}

func (v *NullableIncomeVerificationDocParsingConfig) Set(val *IncomeVerificationDocParsingConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationDocParsingConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationDocParsingConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationDocParsingConfig(val *IncomeVerificationDocParsingConfig) *NullableIncomeVerificationDocParsingConfig {
	return &NullableIncomeVerificationDocParsingConfig{value: val, isSet: true}
}

func (v NullableIncomeVerificationDocParsingConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationDocParsingConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

