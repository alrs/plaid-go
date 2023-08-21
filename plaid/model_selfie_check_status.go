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

// SelfieCheckStatus The outcome status for the associated Identity Verification attempt's `selfie_check` step. This field will always have the same value as `steps.selfie_check`.
type SelfieCheckStatus string

var _ = fmt.Printf

// List of SelfieCheckStatus
const (
	SELFIECHECKSTATUS_SUCCESS SelfieCheckStatus = "success"
	SELFIECHECKSTATUS_FAILED SelfieCheckStatus = "failed"
)

var allowedSelfieCheckStatusEnumValues = []SelfieCheckStatus{
	"success",
	"failed",
}

func (v *SelfieCheckStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := SelfieCheckStatus(value)


	*v = enumTypeValue
	return nil
}

// NewSelfieCheckStatusFromValue returns a pointer to a valid SelfieCheckStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSelfieCheckStatusFromValue(v string) (*SelfieCheckStatus, error) {
	ev := SelfieCheckStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SelfieCheckStatus) IsValid() bool {
	for _, existing := range allowedSelfieCheckStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SelfieCheckStatus value
func (v SelfieCheckStatus) Ptr() *SelfieCheckStatus {
	return &v
}

type NullableSelfieCheckStatus struct {
	value *SelfieCheckStatus
	isSet bool
}

func (v NullableSelfieCheckStatus) Get() *SelfieCheckStatus {
	return v.value
}

func (v *NullableSelfieCheckStatus) Set(val *SelfieCheckStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfieCheckStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfieCheckStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfieCheckStatus(val *SelfieCheckStatus) *NullableSelfieCheckStatus {
	return &NullableSelfieCheckStatus{value: val, isSet: true}
}

func (v NullableSelfieCheckStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfieCheckStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

