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

// FDXNotificationSeverity Severity level of notification
type FDXNotificationSeverity string

var _ = fmt.Printf

// List of FDXNotificationSeverity
const (
	FDXNOTIFICATIONSEVERITY_EMERGENCY FDXNotificationSeverity = "EMERGENCY"
	FDXNOTIFICATIONSEVERITY_ALERT FDXNotificationSeverity = "ALERT"
	FDXNOTIFICATIONSEVERITY_WARNING FDXNotificationSeverity = "WARNING"
	FDXNOTIFICATIONSEVERITY_NOTICE FDXNotificationSeverity = "NOTICE"
	FDXNOTIFICATIONSEVERITY_INFO FDXNotificationSeverity = "INFO"
)

var allowedFDXNotificationSeverityEnumValues = []FDXNotificationSeverity{
	"EMERGENCY",
	"ALERT",
	"WARNING",
	"NOTICE",
	"INFO",
}

func (v *FDXNotificationSeverity) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := FDXNotificationSeverity(value)


	*v = enumTypeValue
	return nil
}

// NewFDXNotificationSeverityFromValue returns a pointer to a valid FDXNotificationSeverity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFDXNotificationSeverityFromValue(v string) (*FDXNotificationSeverity, error) {
	ev := FDXNotificationSeverity(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FDXNotificationSeverity) IsValid() bool {
	for _, existing := range allowedFDXNotificationSeverityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FDXNotificationSeverity value
func (v FDXNotificationSeverity) Ptr() *FDXNotificationSeverity {
	return &v
}

type NullableFDXNotificationSeverity struct {
	value *FDXNotificationSeverity
	isSet bool
}

func (v NullableFDXNotificationSeverity) Get() *FDXNotificationSeverity {
	return v.value
}

func (v *NullableFDXNotificationSeverity) Set(val *FDXNotificationSeverity) {
	v.value = val
	v.isSet = true
}

func (v NullableFDXNotificationSeverity) IsSet() bool {
	return v.isSet
}

func (v *NullableFDXNotificationSeverity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFDXNotificationSeverity(val *FDXNotificationSeverity) *NullableFDXNotificationSeverity {
	return &NullableFDXNotificationSeverity{value: val, isSet: true}
}

func (v NullableFDXNotificationSeverity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFDXNotificationSeverity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

