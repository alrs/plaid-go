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

// PaymentProfileStatus The status of the given Payment Profile.  `READY`: This Payment Profile is ready to be used to create transfers using `/transfer/authorization/create` and `/transfer/create`.  `PENDING`: This Payment Profile is not ready to be used. You’ll need to call `/link/token/create` and provide the `payment_profile_token` in the `transfer.payment_profile_token` field to initiate the account linking experience.  `REMOVED`: This Payment Profile has been removed.
type PaymentProfileStatus string

var _ = fmt.Printf

// List of PaymentProfileStatus
const (
	PAYMENTPROFILESTATUS_PENDING PaymentProfileStatus = "PENDING"
	PAYMENTPROFILESTATUS_READY PaymentProfileStatus = "READY"
	PAYMENTPROFILESTATUS_REMOVED PaymentProfileStatus = "REMOVED"
)

var allowedPaymentProfileStatusEnumValues = []PaymentProfileStatus{
	"PENDING",
	"READY",
	"REMOVED",
}

func (v *PaymentProfileStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := PaymentProfileStatus(value)


	*v = enumTypeValue
	return nil
}

// NewPaymentProfileStatusFromValue returns a pointer to a valid PaymentProfileStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentProfileStatusFromValue(v string) (*PaymentProfileStatus, error) {
	ev := PaymentProfileStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentProfileStatus) IsValid() bool {
	for _, existing := range allowedPaymentProfileStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentProfileStatus value
func (v PaymentProfileStatus) Ptr() *PaymentProfileStatus {
	return &v
}

type NullablePaymentProfileStatus struct {
	value *PaymentProfileStatus
	isSet bool
}

func (v NullablePaymentProfileStatus) Get() *PaymentProfileStatus {
	return v.value
}

func (v *NullablePaymentProfileStatus) Set(val *PaymentProfileStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentProfileStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentProfileStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentProfileStatus(val *PaymentProfileStatus) *NullablePaymentProfileStatus {
	return &NullablePaymentProfileStatus{value: val, isSet: true}
}

func (v NullablePaymentProfileStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentProfileStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

