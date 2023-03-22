/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.343.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// DocumentNameMatchCode A match summary describing the cross comparison between the subject's name, extracted from the document image, and the name they separately provided to identity verification attempt.
type DocumentNameMatchCode string

var _ = fmt.Printf

// List of DocumentNameMatchCode
const (
	DOCUMENTNAMEMATCHCODE_MATCH DocumentNameMatchCode = "match"
	DOCUMENTNAMEMATCHCODE_PARTIAL_MATCH DocumentNameMatchCode = "partial_match"
	DOCUMENTNAMEMATCHCODE_NO_MATCH DocumentNameMatchCode = "no_match"
)

var allowedDocumentNameMatchCodeEnumValues = []DocumentNameMatchCode{
	"match",
	"partial_match",
	"no_match",
}

func (v *DocumentNameMatchCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := DocumentNameMatchCode(value)


	*v = enumTypeValue
	return nil
}

// NewDocumentNameMatchCodeFromValue returns a pointer to a valid DocumentNameMatchCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDocumentNameMatchCodeFromValue(v string) (*DocumentNameMatchCode, error) {
	ev := DocumentNameMatchCode(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DocumentNameMatchCode) IsValid() bool {
	for _, existing := range allowedDocumentNameMatchCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DocumentNameMatchCode value
func (v DocumentNameMatchCode) Ptr() *DocumentNameMatchCode {
	return &v
}

type NullableDocumentNameMatchCode struct {
	value *DocumentNameMatchCode
	isSet bool
}

func (v NullableDocumentNameMatchCode) Get() *DocumentNameMatchCode {
	return v.value
}

func (v *NullableDocumentNameMatchCode) Set(val *DocumentNameMatchCode) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentNameMatchCode) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentNameMatchCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentNameMatchCode(val *DocumentNameMatchCode) *NullableDocumentNameMatchCode {
	return &NullableDocumentNameMatchCode{value: val, isSet: true}
}

func (v NullableDocumentNameMatchCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentNameMatchCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

