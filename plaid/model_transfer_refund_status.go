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

// TransferRefundStatus The status of the refund.  `pending`: A new refund was created; it is in the pending state. `posted`: The refund has been successfully submitted to the payment network. `cancelled`: The refund was cancelled by the client. `failed`: The refund failed or was returned.
type TransferRefundStatus string

var _ = fmt.Printf

// List of TransferRefundStatus
const (
	TRANSFERREFUNDSTATUS_PENDING TransferRefundStatus = "pending"
	TRANSFERREFUNDSTATUS_POSTED TransferRefundStatus = "posted"
	TRANSFERREFUNDSTATUS_CANCELLED TransferRefundStatus = "cancelled"
	TRANSFERREFUNDSTATUS_FAILED TransferRefundStatus = "failed"
)

var allowedTransferRefundStatusEnumValues = []TransferRefundStatus{
	"pending",
	"posted",
	"cancelled",
	"failed",
}

func (v *TransferRefundStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := TransferRefundStatus(value)


	*v = enumTypeValue
	return nil
}

// NewTransferRefundStatusFromValue returns a pointer to a valid TransferRefundStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferRefundStatusFromValue(v string) (*TransferRefundStatus, error) {
	ev := TransferRefundStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferRefundStatus) IsValid() bool {
	for _, existing := range allowedTransferRefundStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferRefundStatus value
func (v TransferRefundStatus) Ptr() *TransferRefundStatus {
	return &v
}

type NullableTransferRefundStatus struct {
	value *TransferRefundStatus
	isSet bool
}

func (v NullableTransferRefundStatus) Get() *TransferRefundStatus {
	return v.value
}

func (v *NullableTransferRefundStatus) Set(val *TransferRefundStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRefundStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRefundStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRefundStatus(val *TransferRefundStatus) *NullableTransferRefundStatus {
	return &NullableTransferRefundStatus{value: val, isSet: true}
}

func (v NullableTransferRefundStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRefundStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

