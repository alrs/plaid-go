/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.79.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// EmploymentVerificationStatus Current employment status.
type EmploymentVerificationStatus string

// List of EmploymentVerificationStatus
const (
	EMPLOYMENTVERIFICATIONSTATUS_EMPLOYMENT_STATUS_ACTIVE EmploymentVerificationStatus = "EMPLOYMENT_STATUS_ACTIVE"
	EMPLOYMENTVERIFICATIONSTATUS_EMPLOYMENT_STATUS_INACTIVE EmploymentVerificationStatus = "EMPLOYMENT_STATUS_INACTIVE"
	EMPLOYMENTVERIFICATIONSTATUS_NULL EmploymentVerificationStatus = "null"
)

var allowedEmploymentVerificationStatusEnumValues = []EmploymentVerificationStatus{
	"EMPLOYMENT_STATUS_ACTIVE",
	"EMPLOYMENT_STATUS_INACTIVE",
	"null",
}

func (v *EmploymentVerificationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EmploymentVerificationStatus(value)
	for _, existing := range allowedEmploymentVerificationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EmploymentVerificationStatus", value)
}

// NewEmploymentVerificationStatusFromValue returns a pointer to a valid EmploymentVerificationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEmploymentVerificationStatusFromValue(v string) (*EmploymentVerificationStatus, error) {
	ev := EmploymentVerificationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EmploymentVerificationStatus: valid values are %v", v, allowedEmploymentVerificationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EmploymentVerificationStatus) IsValid() bool {
	for _, existing := range allowedEmploymentVerificationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EmploymentVerificationStatus value
func (v EmploymentVerificationStatus) Ptr() *EmploymentVerificationStatus {
	return &v
}

type NullableEmploymentVerificationStatus struct {
	value *EmploymentVerificationStatus
	isSet bool
}

func (v NullableEmploymentVerificationStatus) Get() *EmploymentVerificationStatus {
	return v.value
}

func (v *NullableEmploymentVerificationStatus) Set(val *EmploymentVerificationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEmploymentVerificationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEmploymentVerificationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmploymentVerificationStatus(val *EmploymentVerificationStatus) *NullableEmploymentVerificationStatus {
	return &NullableEmploymentVerificationStatus{value: val, isSet: true}
}

func (v NullableEmploymentVerificationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmploymentVerificationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

