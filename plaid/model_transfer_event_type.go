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

// TransferEventType The type of event that this transfer represents.  `pending`: A new transfer was created; it is in the pending state.  `cancelled`: The transfer was cancelled by the client.  `failed`: The transfer failed, no funds were moved.  `posted`: The transfer has been successfully submitted to the payment network.  `settled`: Credits are available to be withdrawn or debits have been deducted from the Plaid linked account.  `returned`: A posted transfer was returned.  `swept`: The transfer was swept to / from the sweep account.  `swept_settled`: Credits are available to be withdrawn or debits have been deducted from the customer’s business checking account.  `return_swept`: Due to the transfer being returned, funds were pulled from or pushed back to the sweep account.
type TransferEventType string

var _ = fmt.Printf

// List of TransferEventType
const (
	TRANSFEREVENTTYPE_PENDING TransferEventType = "pending"
	TRANSFEREVENTTYPE_CANCELLED TransferEventType = "cancelled"
	TRANSFEREVENTTYPE_FAILED TransferEventType = "failed"
	TRANSFEREVENTTYPE_POSTED TransferEventType = "posted"
	TRANSFEREVENTTYPE_SETTLED TransferEventType = "settled"
	TRANSFEREVENTTYPE_RETURNED TransferEventType = "returned"
	TRANSFEREVENTTYPE_SWEPT TransferEventType = "swept"
	TRANSFEREVENTTYPE_SWEPT_SETTLED TransferEventType = "swept_settled"
	TRANSFEREVENTTYPE_RETURN_SWEPT TransferEventType = "return_swept"
)

var allowedTransferEventTypeEnumValues = []TransferEventType{
	"pending",
	"cancelled",
	"failed",
	"posted",
	"settled",
	"returned",
	"swept",
	"swept_settled",
	"return_swept",
}

func (v *TransferEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := TransferEventType(value)


	*v = enumTypeValue
	return nil
}

// NewTransferEventTypeFromValue returns a pointer to a valid TransferEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferEventTypeFromValue(v string) (*TransferEventType, error) {
	ev := TransferEventType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferEventType) IsValid() bool {
	for _, existing := range allowedTransferEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferEventType value
func (v TransferEventType) Ptr() *TransferEventType {
	return &v
}

type NullableTransferEventType struct {
	value *TransferEventType
	isSet bool
}

func (v NullableTransferEventType) Get() *TransferEventType {
	return v.value
}

func (v *NullableTransferEventType) Set(val *TransferEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferEventType(val *TransferEventType) *NullableTransferEventType {
	return &NullableTransferEventType{value: val, isSet: true}
}

func (v NullableTransferEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

