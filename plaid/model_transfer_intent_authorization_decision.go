/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.79.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// TransferIntentAuthorizationDecision  A decision regarding the proposed transfer.  `APPROVED` – The proposed transfer has received the end user's consent and has been approved for processing. Plaid has also reviewed the proposed transfer and has approved it for processing.   `PERMITTED` – Plaid was unable to fetch the information required to approve or decline the proposed transfer. You may proceed with the transfer, but further review is recommended. Plaid is awaiting further instructions from the client.  `DECLINED` – Plaid reviewed the proposed transfer and declined processing. Refer to the `code` field in the `decision_rationale` object for details. Null value otherwise.
type TransferIntentAuthorizationDecision string

// List of TransferIntentAuthorizationDecision
const (
	TRANSFERINTENTAUTHORIZATIONDECISION_APPROVED TransferIntentAuthorizationDecision = "APPROVED"
	TRANSFERINTENTAUTHORIZATIONDECISION_PERMITTED TransferIntentAuthorizationDecision = "PERMITTED"
	TRANSFERINTENTAUTHORIZATIONDECISION_DECLINED TransferIntentAuthorizationDecision = "DECLINED"
)

var allowedTransferIntentAuthorizationDecisionEnumValues = []TransferIntentAuthorizationDecision{
	"APPROVED",
	"PERMITTED",
	"DECLINED",
}

func (v *TransferIntentAuthorizationDecision) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransferIntentAuthorizationDecision(value)
	for _, existing := range allowedTransferIntentAuthorizationDecisionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransferIntentAuthorizationDecision", value)
}

// NewTransferIntentAuthorizationDecisionFromValue returns a pointer to a valid TransferIntentAuthorizationDecision
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferIntentAuthorizationDecisionFromValue(v string) (*TransferIntentAuthorizationDecision, error) {
	ev := TransferIntentAuthorizationDecision(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransferIntentAuthorizationDecision: valid values are %v", v, allowedTransferIntentAuthorizationDecisionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferIntentAuthorizationDecision) IsValid() bool {
	for _, existing := range allowedTransferIntentAuthorizationDecisionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferIntentAuthorizationDecision value
func (v TransferIntentAuthorizationDecision) Ptr() *TransferIntentAuthorizationDecision {
	return &v
}

type NullableTransferIntentAuthorizationDecision struct {
	value *TransferIntentAuthorizationDecision
	isSet bool
}

func (v NullableTransferIntentAuthorizationDecision) Get() *TransferIntentAuthorizationDecision {
	return v.value
}

func (v *NullableTransferIntentAuthorizationDecision) Set(val *TransferIntentAuthorizationDecision) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferIntentAuthorizationDecision) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferIntentAuthorizationDecision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferIntentAuthorizationDecision(val *TransferIntentAuthorizationDecision) *NullableTransferIntentAuthorizationDecision {
	return &NullableTransferIntentAuthorizationDecision{value: val, isSet: true}
}

func (v NullableTransferIntentAuthorizationDecision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferIntentAuthorizationDecision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
