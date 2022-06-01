/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.121.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// POBoxStatus Field describing whether the associated address is a post office box. Will be `yes` when a P.O. box is detected, `no` when Plaid confirmed the address is not a P.O. box, and `no_data` when Plaid was not able to determine if the address is a P.O. box.
type POBoxStatus string

// List of POBoxStatus
const (
	POBOXSTATUS_YES POBoxStatus = "yes"
	POBOXSTATUS_NO POBoxStatus = "no"
	POBOXSTATUS_NO_DATA POBoxStatus = "no_data"
)

var allowedPOBoxStatusEnumValues = []POBoxStatus{
	"yes",
	"no",
	"no_data",
}

func (v *POBoxStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := POBoxStatus(value)
	for _, existing := range allowedPOBoxStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid POBoxStatus", value)
}

// NewPOBoxStatusFromValue returns a pointer to a valid POBoxStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPOBoxStatusFromValue(v string) (*POBoxStatus, error) {
	ev := POBoxStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for POBoxStatus: valid values are %v", v, allowedPOBoxStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v POBoxStatus) IsValid() bool {
	for _, existing := range allowedPOBoxStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to POBoxStatus value
func (v POBoxStatus) Ptr() *POBoxStatus {
	return &v
}

type NullablePOBoxStatus struct {
	value *POBoxStatus
	isSet bool
}

func (v NullablePOBoxStatus) Get() *POBoxStatus {
	return v.value
}

func (v *NullablePOBoxStatus) Set(val *POBoxStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePOBoxStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePOBoxStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOBoxStatus(val *POBoxStatus) *NullablePOBoxStatus {
	return &NullablePOBoxStatus{value: val, isSet: true}
}

func (v NullablePOBoxStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOBoxStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

