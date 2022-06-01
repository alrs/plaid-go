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

// TransferAuthorizationGuaranteeDecision Indicates whether the transfer is guaranteed by Plaid (Guaranteed ACH customers only). This field will contain either `GUARANTEED` or `NOT_GUARANTEED` indicating whether Plaid will guarantee the transfer. If the transfer is not guaranteed, additional information will be provided in the `guarantee_decision_rationale` field. Refer to the `code` field in `guarantee_decision_rationale` for details.
type TransferAuthorizationGuaranteeDecision string

// List of TransferAuthorizationGuaranteeDecision
const (
	TRANSFERAUTHORIZATIONGUARANTEEDECISION_GUARANTEED TransferAuthorizationGuaranteeDecision = "GUARANTEED"
	TRANSFERAUTHORIZATIONGUARANTEEDECISION_NOT_GUARANTEED TransferAuthorizationGuaranteeDecision = "NOT_GUARANTEED"
	TRANSFERAUTHORIZATIONGUARANTEEDECISION_NULL TransferAuthorizationGuaranteeDecision = "null"
)

var allowedTransferAuthorizationGuaranteeDecisionEnumValues = []TransferAuthorizationGuaranteeDecision{
	"GUARANTEED",
	"NOT_GUARANTEED",
	"null",
}

func (v *TransferAuthorizationGuaranteeDecision) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransferAuthorizationGuaranteeDecision(value)
	for _, existing := range allowedTransferAuthorizationGuaranteeDecisionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransferAuthorizationGuaranteeDecision", value)
}

// NewTransferAuthorizationGuaranteeDecisionFromValue returns a pointer to a valid TransferAuthorizationGuaranteeDecision
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferAuthorizationGuaranteeDecisionFromValue(v string) (*TransferAuthorizationGuaranteeDecision, error) {
	ev := TransferAuthorizationGuaranteeDecision(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransferAuthorizationGuaranteeDecision: valid values are %v", v, allowedTransferAuthorizationGuaranteeDecisionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferAuthorizationGuaranteeDecision) IsValid() bool {
	for _, existing := range allowedTransferAuthorizationGuaranteeDecisionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferAuthorizationGuaranteeDecision value
func (v TransferAuthorizationGuaranteeDecision) Ptr() *TransferAuthorizationGuaranteeDecision {
	return &v
}

type NullableTransferAuthorizationGuaranteeDecision struct {
	value *TransferAuthorizationGuaranteeDecision
	isSet bool
}

func (v NullableTransferAuthorizationGuaranteeDecision) Get() *TransferAuthorizationGuaranteeDecision {
	return v.value
}

func (v *NullableTransferAuthorizationGuaranteeDecision) Set(val *TransferAuthorizationGuaranteeDecision) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferAuthorizationGuaranteeDecision) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferAuthorizationGuaranteeDecision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferAuthorizationGuaranteeDecision(val *TransferAuthorizationGuaranteeDecision) *NullableTransferAuthorizationGuaranteeDecision {
	return &NullableTransferAuthorizationGuaranteeDecision{value: val, isSet: true}
}

func (v NullableTransferAuthorizationGuaranteeDecision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferAuthorizationGuaranteeDecision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

