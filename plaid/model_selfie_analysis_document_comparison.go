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

// SelfieAnalysisDocumentComparison Information about the comparison between the selfie and the document (if documentary verification also ran).
type SelfieAnalysisDocumentComparison string

var _ = fmt.Printf

// List of SelfieAnalysisDocumentComparison
const (
	SELFIEANALYSISDOCUMENTCOMPARISON_MATCH SelfieAnalysisDocumentComparison = "match"
	SELFIEANALYSISDOCUMENTCOMPARISON_NO_MATCH SelfieAnalysisDocumentComparison = "no_match"
	SELFIEANALYSISDOCUMENTCOMPARISON_NO_INPUT SelfieAnalysisDocumentComparison = "no_input"
)

var allowedSelfieAnalysisDocumentComparisonEnumValues = []SelfieAnalysisDocumentComparison{
	"match",
	"no_match",
	"no_input",
}

func (v *SelfieAnalysisDocumentComparison) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := SelfieAnalysisDocumentComparison(value)


	*v = enumTypeValue
	return nil
}

// NewSelfieAnalysisDocumentComparisonFromValue returns a pointer to a valid SelfieAnalysisDocumentComparison
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSelfieAnalysisDocumentComparisonFromValue(v string) (*SelfieAnalysisDocumentComparison, error) {
	ev := SelfieAnalysisDocumentComparison(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SelfieAnalysisDocumentComparison) IsValid() bool {
	for _, existing := range allowedSelfieAnalysisDocumentComparisonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SelfieAnalysisDocumentComparison value
func (v SelfieAnalysisDocumentComparison) Ptr() *SelfieAnalysisDocumentComparison {
	return &v
}

type NullableSelfieAnalysisDocumentComparison struct {
	value *SelfieAnalysisDocumentComparison
	isSet bool
}

func (v NullableSelfieAnalysisDocumentComparison) Get() *SelfieAnalysisDocumentComparison {
	return v.value
}

func (v *NullableSelfieAnalysisDocumentComparison) Set(val *SelfieAnalysisDocumentComparison) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfieAnalysisDocumentComparison) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfieAnalysisDocumentComparison) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfieAnalysisDocumentComparison(val *SelfieAnalysisDocumentComparison) *NullableSelfieAnalysisDocumentComparison {
	return &NullableSelfieAnalysisDocumentComparison{value: val, isSet: true}
}

func (v NullableSelfieAnalysisDocumentComparison) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfieAnalysisDocumentComparison) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

