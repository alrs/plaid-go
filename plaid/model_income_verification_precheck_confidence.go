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
	"fmt"
)

// IncomeVerificationPrecheckConfidence The confidence that Plaid can support the user in the digital income verification flow instead of requiring a manual paystub upload. One of the following:  `\"HIGH\"`: It is very likely that this user can use the digital income verification flow.  \"`LOW`\": It is unlikely that this user can use the digital income verification flow.  `\"UNKNOWN\"`: It was not possible to determine if the user is supportable with the information passed.
type IncomeVerificationPrecheckConfidence string

var _ = fmt.Printf

// List of IncomeVerificationPrecheckConfidence
const (
	INCOMEVERIFICATIONPRECHECKCONFIDENCE_HIGH IncomeVerificationPrecheckConfidence = "HIGH"
	INCOMEVERIFICATIONPRECHECKCONFIDENCE_LOW IncomeVerificationPrecheckConfidence = "LOW"
	INCOMEVERIFICATIONPRECHECKCONFIDENCE_UNKNOWN IncomeVerificationPrecheckConfidence = "UNKNOWN"
)

var allowedIncomeVerificationPrecheckConfidenceEnumValues = []IncomeVerificationPrecheckConfidence{
	"HIGH",
	"LOW",
	"UNKNOWN",
}

func (v *IncomeVerificationPrecheckConfidence) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := IncomeVerificationPrecheckConfidence(value)


	*v = enumTypeValue
	return nil
}

// NewIncomeVerificationPrecheckConfidenceFromValue returns a pointer to a valid IncomeVerificationPrecheckConfidence
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIncomeVerificationPrecheckConfidenceFromValue(v string) (*IncomeVerificationPrecheckConfidence, error) {
	ev := IncomeVerificationPrecheckConfidence(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IncomeVerificationPrecheckConfidence) IsValid() bool {
	for _, existing := range allowedIncomeVerificationPrecheckConfidenceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncomeVerificationPrecheckConfidence value
func (v IncomeVerificationPrecheckConfidence) Ptr() *IncomeVerificationPrecheckConfidence {
	return &v
}

type NullableIncomeVerificationPrecheckConfidence struct {
	value *IncomeVerificationPrecheckConfidence
	isSet bool
}

func (v NullableIncomeVerificationPrecheckConfidence) Get() *IncomeVerificationPrecheckConfidence {
	return v.value
}

func (v *NullableIncomeVerificationPrecheckConfidence) Set(val *IncomeVerificationPrecheckConfidence) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationPrecheckConfidence) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationPrecheckConfidence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationPrecheckConfidence(val *IncomeVerificationPrecheckConfidence) *NullableIncomeVerificationPrecheckConfidence {
	return &NullableIncomeVerificationPrecheckConfidence{value: val, isSet: true}
}

func (v NullableIncomeVerificationPrecheckConfidence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationPrecheckConfidence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

