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

// TransferAuthorizationGuaranteeDecisionRationaleCode A code representing the reason Plaid declined to guarantee this transfer:  `RETURN_BANK`: The risk of a bank-initiated return (for example, an R01/NSF) is too high to guarantee this transfer.  `RETURN_CUSTOMER`: The risk of a customer-initiated return (for example, a R10/Unauthorized) is too high to guarantee this transfer.  `GUARANTEE_LIMIT_REACHED`: This transfer is low-risk, but Guarantee has exhausted an internal limit on the number or rate of guarantees that applies to this transfer.  `RISK_ESTIMATE_UNAVAILABLE`: A risk estimate is unavailable for this Item.  `REQUIRED_PARAM_MISSING`: Required fields are missing.
type TransferAuthorizationGuaranteeDecisionRationaleCode string

var _ = fmt.Printf

// List of TransferAuthorizationGuaranteeDecisionRationaleCode
const (
	TRANSFERAUTHORIZATIONGUARANTEEDECISIONRATIONALECODE_RETURN_BANK TransferAuthorizationGuaranteeDecisionRationaleCode = "RETURN_BANK"
	TRANSFERAUTHORIZATIONGUARANTEEDECISIONRATIONALECODE_RETURN_CUSTOMER TransferAuthorizationGuaranteeDecisionRationaleCode = "RETURN_CUSTOMER"
	TRANSFERAUTHORIZATIONGUARANTEEDECISIONRATIONALECODE_GUARANTEE_LIMIT_REACHED TransferAuthorizationGuaranteeDecisionRationaleCode = "GUARANTEE_LIMIT_REACHED"
	TRANSFERAUTHORIZATIONGUARANTEEDECISIONRATIONALECODE_RISK_ESTIMATE_UNAVAILABLE TransferAuthorizationGuaranteeDecisionRationaleCode = "RISK_ESTIMATE_UNAVAILABLE"
	TRANSFERAUTHORIZATIONGUARANTEEDECISIONRATIONALECODE_REQUIRED_PARAM_MISSING TransferAuthorizationGuaranteeDecisionRationaleCode = "REQUIRED_PARAM_MISSING"
)

var allowedTransferAuthorizationGuaranteeDecisionRationaleCodeEnumValues = []TransferAuthorizationGuaranteeDecisionRationaleCode{
	"RETURN_BANK",
	"RETURN_CUSTOMER",
	"GUARANTEE_LIMIT_REACHED",
	"RISK_ESTIMATE_UNAVAILABLE",
	"REQUIRED_PARAM_MISSING",
}

func (v *TransferAuthorizationGuaranteeDecisionRationaleCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := TransferAuthorizationGuaranteeDecisionRationaleCode(value)


	*v = enumTypeValue
	return nil
}

// NewTransferAuthorizationGuaranteeDecisionRationaleCodeFromValue returns a pointer to a valid TransferAuthorizationGuaranteeDecisionRationaleCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferAuthorizationGuaranteeDecisionRationaleCodeFromValue(v string) (*TransferAuthorizationGuaranteeDecisionRationaleCode, error) {
	ev := TransferAuthorizationGuaranteeDecisionRationaleCode(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferAuthorizationGuaranteeDecisionRationaleCode) IsValid() bool {
	for _, existing := range allowedTransferAuthorizationGuaranteeDecisionRationaleCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferAuthorizationGuaranteeDecisionRationaleCode value
func (v TransferAuthorizationGuaranteeDecisionRationaleCode) Ptr() *TransferAuthorizationGuaranteeDecisionRationaleCode {
	return &v
}

type NullableTransferAuthorizationGuaranteeDecisionRationaleCode struct {
	value *TransferAuthorizationGuaranteeDecisionRationaleCode
	isSet bool
}

func (v NullableTransferAuthorizationGuaranteeDecisionRationaleCode) Get() *TransferAuthorizationGuaranteeDecisionRationaleCode {
	return v.value
}

func (v *NullableTransferAuthorizationGuaranteeDecisionRationaleCode) Set(val *TransferAuthorizationGuaranteeDecisionRationaleCode) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferAuthorizationGuaranteeDecisionRationaleCode) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferAuthorizationGuaranteeDecisionRationaleCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferAuthorizationGuaranteeDecisionRationaleCode(val *TransferAuthorizationGuaranteeDecisionRationaleCode) *NullableTransferAuthorizationGuaranteeDecisionRationaleCode {
	return &NullableTransferAuthorizationGuaranteeDecisionRationaleCode{value: val, isSet: true}
}

func (v NullableTransferAuthorizationGuaranteeDecisionRationaleCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferAuthorizationGuaranteeDecisionRationaleCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

