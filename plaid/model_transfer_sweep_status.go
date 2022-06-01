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

// TransferSweepStatus The status of the sweep for the transfer. `unswept`: The transfer hasn't been swept yet. `swept`: The transfer was swept to the sweep account. `return_swept`: The transfer was returned, funds were pulled back or pushed back to the sweep account. `null`: The transfer will never be swept (e.g. if the transfer is cancelled or returned before being swept)
type TransferSweepStatus string

// List of TransferSweepStatus
const (
	TRANSFERSWEEPSTATUS_NULL TransferSweepStatus = "null"
	TRANSFERSWEEPSTATUS_UNSWEPT TransferSweepStatus = "unswept"
	TRANSFERSWEEPSTATUS_SWEPT TransferSweepStatus = "swept"
	TRANSFERSWEEPSTATUS_REVERSE_SWEPT TransferSweepStatus = "reverse_swept"
	TRANSFERSWEEPSTATUS_RETURN_SWEPT TransferSweepStatus = "return_swept"
)

var allowedTransferSweepStatusEnumValues = []TransferSweepStatus{
	"null",
	"unswept",
	"swept",
	"reverse_swept",
	"return_swept",
}

func (v *TransferSweepStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransferSweepStatus(value)
	for _, existing := range allowedTransferSweepStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransferSweepStatus", value)
}

// NewTransferSweepStatusFromValue returns a pointer to a valid TransferSweepStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferSweepStatusFromValue(v string) (*TransferSweepStatus, error) {
	ev := TransferSweepStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransferSweepStatus: valid values are %v", v, allowedTransferSweepStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferSweepStatus) IsValid() bool {
	for _, existing := range allowedTransferSweepStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferSweepStatus value
func (v TransferSweepStatus) Ptr() *TransferSweepStatus {
	return &v
}

type NullableTransferSweepStatus struct {
	value *TransferSweepStatus
	isSet bool
}

func (v NullableTransferSweepStatus) Get() *TransferSweepStatus {
	return v.value
}

func (v *NullableTransferSweepStatus) Set(val *TransferSweepStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferSweepStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferSweepStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferSweepStatus(val *TransferSweepStatus) *NullableTransferSweepStatus {
	return &NullableTransferSweepStatus{value: val, isSet: true}
}

func (v NullableTransferSweepStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferSweepStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

