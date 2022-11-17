/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.205.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// FDXNotificationPriority Priority of notification
type FDXNotificationPriority string

var _ = fmt.Printf

// List of FDXNotificationPriority
const (
	FDXNOTIFICATIONPRIORITY_HIGH FDXNotificationPriority = "HIGH"
	FDXNOTIFICATIONPRIORITY_MEDIUM FDXNotificationPriority = "MEDIUM"
	FDXNOTIFICATIONPRIORITY_LOW FDXNotificationPriority = "LOW"
)

var allowedFDXNotificationPriorityEnumValues = []FDXNotificationPriority{
	"HIGH",
	"MEDIUM",
	"LOW",
}

func (v *FDXNotificationPriority) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := FDXNotificationPriority(value)


	*v = enumTypeValue
	return nil
}

// NewFDXNotificationPriorityFromValue returns a pointer to a valid FDXNotificationPriority
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFDXNotificationPriorityFromValue(v string) (*FDXNotificationPriority, error) {
	ev := FDXNotificationPriority(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FDXNotificationPriority) IsValid() bool {
	for _, existing := range allowedFDXNotificationPriorityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FDXNotificationPriority value
func (v FDXNotificationPriority) Ptr() *FDXNotificationPriority {
	return &v
}

type NullableFDXNotificationPriority struct {
	value *FDXNotificationPriority
	isSet bool
}

func (v NullableFDXNotificationPriority) Get() *FDXNotificationPriority {
	return v.value
}

func (v *NullableFDXNotificationPriority) Set(val *FDXNotificationPriority) {
	v.value = val
	v.isSet = true
}

func (v NullableFDXNotificationPriority) IsSet() bool {
	return v.isSet
}

func (v *NullableFDXNotificationPriority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFDXNotificationPriority(val *FDXNotificationPriority) *NullableFDXNotificationPriority {
	return &NullableFDXNotificationPriority{value: val, isSet: true}
}

func (v NullableFDXNotificationPriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFDXNotificationPriority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

