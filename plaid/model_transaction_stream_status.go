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

// TransactionStreamStatus The current status of the transaction stream.  `MATURE`: A `MATURE` recurring stream should have at least 3 transactions and happen on a regular cadence (For Annual recurring stream, we will mark it `MATURE` after 2 instances).  `EARLY_DETECTION`: When a recurring transaction first appears in the transaction history and before it fulfills the requirement of a mature stream, the status will be `EARLY_DETECTION`.  `TOMBSTONED`: A stream that was previously in the `EARLY_DETECTION` status will move to the `TOMBSTONED` status when no further transactions were found at the next expected date.  `UNKNOWN`: A stream is assigned an `UNKNOWN` status when none of the other statuses are applicable.
type TransactionStreamStatus string

var _ = fmt.Printf

// List of TransactionStreamStatus
const (
	TRANSACTIONSTREAMSTATUS_UNKNOWN TransactionStreamStatus = "UNKNOWN"
	TRANSACTIONSTREAMSTATUS_MATURE TransactionStreamStatus = "MATURE"
	TRANSACTIONSTREAMSTATUS_EARLY_DETECTION TransactionStreamStatus = "EARLY_DETECTION"
	TRANSACTIONSTREAMSTATUS_TOMBSTONED TransactionStreamStatus = "TOMBSTONED"
)

var allowedTransactionStreamStatusEnumValues = []TransactionStreamStatus{
	"UNKNOWN",
	"MATURE",
	"EARLY_DETECTION",
	"TOMBSTONED",
}

func (v *TransactionStreamStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := TransactionStreamStatus(value)


	*v = enumTypeValue
	return nil
}

// NewTransactionStreamStatusFromValue returns a pointer to a valid TransactionStreamStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransactionStreamStatusFromValue(v string) (*TransactionStreamStatus, error) {
	ev := TransactionStreamStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransactionStreamStatus) IsValid() bool {
	for _, existing := range allowedTransactionStreamStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransactionStreamStatus value
func (v TransactionStreamStatus) Ptr() *TransactionStreamStatus {
	return &v
}

type NullableTransactionStreamStatus struct {
	value *TransactionStreamStatus
	isSet bool
}

func (v NullableTransactionStreamStatus) Get() *TransactionStreamStatus {
	return v.value
}

func (v *NullableTransactionStreamStatus) Set(val *TransactionStreamStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionStreamStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionStreamStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionStreamStatus(val *TransactionStreamStatus) *NullableTransactionStreamStatus {
	return &NullableTransactionStreamStatus{value: val, isSet: true}
}

func (v NullableTransactionStreamStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionStreamStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

