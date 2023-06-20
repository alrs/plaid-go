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

// IncomeBreakdownType The type of income. Possible values include:   `\"regular\"`: regular income   `\"overtime\"`: overtime income   `\"bonus\"`: bonus income
type IncomeBreakdownType string

var _ = fmt.Printf

// List of IncomeBreakdownType
const (
	INCOMEBREAKDOWNTYPE_BONUS IncomeBreakdownType = "bonus"
	INCOMEBREAKDOWNTYPE_OVERTIME IncomeBreakdownType = "overtime"
	INCOMEBREAKDOWNTYPE_REGULAR IncomeBreakdownType = "regular"
	INCOMEBREAKDOWNTYPE_NULL IncomeBreakdownType = "null"
)

var allowedIncomeBreakdownTypeEnumValues = []IncomeBreakdownType{
	"bonus",
	"overtime",
	"regular",
	"null",
}

func (v *IncomeBreakdownType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := IncomeBreakdownType(value)


	*v = enumTypeValue
	return nil
}

// NewIncomeBreakdownTypeFromValue returns a pointer to a valid IncomeBreakdownType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIncomeBreakdownTypeFromValue(v string) (*IncomeBreakdownType, error) {
	ev := IncomeBreakdownType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IncomeBreakdownType) IsValid() bool {
	for _, existing := range allowedIncomeBreakdownTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncomeBreakdownType value
func (v IncomeBreakdownType) Ptr() *IncomeBreakdownType {
	return &v
}

type NullableIncomeBreakdownType struct {
	value *IncomeBreakdownType
	isSet bool
}

func (v NullableIncomeBreakdownType) Get() *IncomeBreakdownType {
	return v.value
}

func (v *NullableIncomeBreakdownType) Set(val *IncomeBreakdownType) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeBreakdownType) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeBreakdownType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeBreakdownType(val *IncomeBreakdownType) *NullableIncomeBreakdownType {
	return &NullableIncomeBreakdownType{value: val, isSet: true}
}

func (v NullableIncomeBreakdownType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeBreakdownType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

