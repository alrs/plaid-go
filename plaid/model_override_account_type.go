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

// OverrideAccountType `investment:` Investment account.  `credit:` Credit card  `depository:` Depository account  `loan:` Loan account  `payroll:` Payroll account  `other:` Non-specified account type  See the [Account type schema](https://plaid.com/docs/api/accounts#account-type-schema) for a full listing of account types and corresponding subtypes.
type OverrideAccountType string

var _ = fmt.Printf

// List of OverrideAccountType
const (
	OVERRIDEACCOUNTTYPE_INVESTMENT OverrideAccountType = "investment"
	OVERRIDEACCOUNTTYPE_CREDIT OverrideAccountType = "credit"
	OVERRIDEACCOUNTTYPE_DEPOSITORY OverrideAccountType = "depository"
	OVERRIDEACCOUNTTYPE_LOAN OverrideAccountType = "loan"
	OVERRIDEACCOUNTTYPE_PAYROLL OverrideAccountType = "payroll"
	OVERRIDEACCOUNTTYPE_OTHER OverrideAccountType = "other"
)

var allowedOverrideAccountTypeEnumValues = []OverrideAccountType{
	"investment",
	"credit",
	"depository",
	"loan",
	"payroll",
	"other",
}

func (v *OverrideAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := OverrideAccountType(value)


	*v = enumTypeValue
	return nil
}

// NewOverrideAccountTypeFromValue returns a pointer to a valid OverrideAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOverrideAccountTypeFromValue(v string) (*OverrideAccountType, error) {
	ev := OverrideAccountType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OverrideAccountType) IsValid() bool {
	for _, existing := range allowedOverrideAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OverrideAccountType value
func (v OverrideAccountType) Ptr() *OverrideAccountType {
	return &v
}

type NullableOverrideAccountType struct {
	value *OverrideAccountType
	isSet bool
}

func (v NullableOverrideAccountType) Get() *OverrideAccountType {
	return v.value
}

func (v *NullableOverrideAccountType) Set(val *OverrideAccountType) {
	v.value = val
	v.isSet = true
}

func (v NullableOverrideAccountType) IsSet() bool {
	return v.isSet
}

func (v *NullableOverrideAccountType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverrideAccountType(val *OverrideAccountType) *NullableOverrideAccountType {
	return &NullableOverrideAccountType{value: val, isSet: true}
}

func (v NullableOverrideAccountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverrideAccountType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

