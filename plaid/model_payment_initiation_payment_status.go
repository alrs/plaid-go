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

// PaymentInitiationPaymentStatus The status of the payment.  `PAYMENT_STATUS_INPUT_NEEDED`: This is the initial state of all payments. It indicates that the payment is waiting on user input to continue processing. A payment may re-enter this state later on if further input is needed.  `PAYMENT_STATUS_INITIATED`: The payment has been successfully authorised and accepted by the financial institution. For successful payments, this is a potential terminal status. Further status transitions can be to REJECTED and, when supported by the institution, to EXECUTED.  `PAYMENT_STATUS_INSUFFICIENT_FUNDS`: The payment has failed due to insufficient funds.  `PAYMENT_STATUS_FAILED`: The payment has failed to be initiated. This error may be caused by transient system outages and is retryable once the root cause is resolved.  `PAYMENT_STATUS_BLOCKED`: The payment has been blocked by Plaid. This can occur, for example, due to Plaid flagging the payment as potentially risky. This is a retryable error.  `PAYMENT_STATUS_AUTHORISING`: The payment is currently being processed. The payment will automatically exit this state when the financial institution has authorised the transaction.  `PAYMENT_STATUS_CANCELLED`: The payment was cancelled (typically by the end user) during authorisation.  `PAYMENT_STATUS_EXECUTED`: The funds have successfully left the payer account and payment is considered complete. Not all institutions support this status: support is more common in the UK, and less common in the EU. For institutions where this status is not supported, the terminal status for a successful payment will be `PAYMENT_STATUS_INITIATED`.  `PAYMENT_STATUS_SETTLED`: The payment has settled and funds are available for use. A payment will typically settle within seconds to several days, depending on which payment rail is used. This status is only available to customers using [Plaid Virtual Accounts](https://plaid.com/docs/virtual-accounts/).  `PAYMENT_STATUS_ESTABLISHED`: Indicates that the standing order has been successfully established. This state is only used for standing orders.  `PAYMENT_STATUS_REJECTED`: The payment was rejected by the financial institution.  Deprecated: These statuses will be removed in a future release.  `PAYMENT_STATUS_UNKNOWN`: The payment status is unknown.  `PAYMENT_STATUS_PROCESSING`: The payment is currently being processed. The payment will automatically exit this state when processing is complete.  `PAYMENT_STATUS_COMPLETED`: Indicates that the standing order has been successfully established. This state is only used for standing orders.
type PaymentInitiationPaymentStatus string

var _ = fmt.Printf

// List of PaymentInitiationPaymentStatus
const (
	PAYMENTINITIATIONPAYMENTSTATUS_INPUT_NEEDED PaymentInitiationPaymentStatus = "PAYMENT_STATUS_INPUT_NEEDED"
	PAYMENTINITIATIONPAYMENTSTATUS_PROCESSING PaymentInitiationPaymentStatus = "PAYMENT_STATUS_PROCESSING"
	PAYMENTINITIATIONPAYMENTSTATUS_INITIATED PaymentInitiationPaymentStatus = "PAYMENT_STATUS_INITIATED"
	PAYMENTINITIATIONPAYMENTSTATUS_COMPLETED PaymentInitiationPaymentStatus = "PAYMENT_STATUS_COMPLETED"
	PAYMENTINITIATIONPAYMENTSTATUS_INSUFFICIENT_FUNDS PaymentInitiationPaymentStatus = "PAYMENT_STATUS_INSUFFICIENT_FUNDS"
	PAYMENTINITIATIONPAYMENTSTATUS_FAILED PaymentInitiationPaymentStatus = "PAYMENT_STATUS_FAILED"
	PAYMENTINITIATIONPAYMENTSTATUS_BLOCKED PaymentInitiationPaymentStatus = "PAYMENT_STATUS_BLOCKED"
	PAYMENTINITIATIONPAYMENTSTATUS_UNKNOWN PaymentInitiationPaymentStatus = "PAYMENT_STATUS_UNKNOWN"
	PAYMENTINITIATIONPAYMENTSTATUS_EXECUTED PaymentInitiationPaymentStatus = "PAYMENT_STATUS_EXECUTED"
	PAYMENTINITIATIONPAYMENTSTATUS_SETTLED PaymentInitiationPaymentStatus = "PAYMENT_STATUS_SETTLED"
	PAYMENTINITIATIONPAYMENTSTATUS_AUTHORISING PaymentInitiationPaymentStatus = "PAYMENT_STATUS_AUTHORISING"
	PAYMENTINITIATIONPAYMENTSTATUS_CANCELLED PaymentInitiationPaymentStatus = "PAYMENT_STATUS_CANCELLED"
	PAYMENTINITIATIONPAYMENTSTATUS_ESTABLISHED PaymentInitiationPaymentStatus = "PAYMENT_STATUS_ESTABLISHED"
	PAYMENTINITIATIONPAYMENTSTATUS_REJECTED PaymentInitiationPaymentStatus = "PAYMENT_STATUS_REJECTED"
)

var allowedPaymentInitiationPaymentStatusEnumValues = []PaymentInitiationPaymentStatus{
	"PAYMENT_STATUS_INPUT_NEEDED",
	"PAYMENT_STATUS_PROCESSING",
	"PAYMENT_STATUS_INITIATED",
	"PAYMENT_STATUS_COMPLETED",
	"PAYMENT_STATUS_INSUFFICIENT_FUNDS",
	"PAYMENT_STATUS_FAILED",
	"PAYMENT_STATUS_BLOCKED",
	"PAYMENT_STATUS_UNKNOWN",
	"PAYMENT_STATUS_EXECUTED",
	"PAYMENT_STATUS_SETTLED",
	"PAYMENT_STATUS_AUTHORISING",
	"PAYMENT_STATUS_CANCELLED",
	"PAYMENT_STATUS_ESTABLISHED",
	"PAYMENT_STATUS_REJECTED",
}

func (v *PaymentInitiationPaymentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := PaymentInitiationPaymentStatus(value)


	*v = enumTypeValue
	return nil
}

// NewPaymentInitiationPaymentStatusFromValue returns a pointer to a valid PaymentInitiationPaymentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentInitiationPaymentStatusFromValue(v string) (*PaymentInitiationPaymentStatus, error) {
	ev := PaymentInitiationPaymentStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentInitiationPaymentStatus) IsValid() bool {
	for _, existing := range allowedPaymentInitiationPaymentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentInitiationPaymentStatus value
func (v PaymentInitiationPaymentStatus) Ptr() *PaymentInitiationPaymentStatus {
	return &v
}

type NullablePaymentInitiationPaymentStatus struct {
	value *PaymentInitiationPaymentStatus
	isSet bool
}

func (v NullablePaymentInitiationPaymentStatus) Get() *PaymentInitiationPaymentStatus {
	return v.value
}

func (v *NullablePaymentInitiationPaymentStatus) Set(val *PaymentInitiationPaymentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationPaymentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationPaymentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationPaymentStatus(val *PaymentInitiationPaymentStatus) *NullablePaymentInitiationPaymentStatus {
	return &NullablePaymentInitiationPaymentStatus{value: val, isSet: true}
}

func (v NullablePaymentInitiationPaymentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationPaymentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

