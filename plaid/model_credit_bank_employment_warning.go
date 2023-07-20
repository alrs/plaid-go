/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.391.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// CreditBankEmploymentWarning The warning associated with the data that was unavailable for the Bank Employment Report.
type CreditBankEmploymentWarning struct {
	WarningType CreditBankEmploymentWarningType `json:"warning_type"`
	WarningCode CreditBankIncomeWarningCode `json:"warning_code"`
	Cause CreditBankIncomeCause `json:"cause"`
}

// NewCreditBankEmploymentWarning instantiates a new CreditBankEmploymentWarning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankEmploymentWarning(warningType CreditBankEmploymentWarningType, warningCode CreditBankIncomeWarningCode, cause CreditBankIncomeCause) *CreditBankEmploymentWarning {
	this := CreditBankEmploymentWarning{}
	this.WarningType = warningType
	this.WarningCode = warningCode
	this.Cause = cause
	return &this
}

// NewCreditBankEmploymentWarningWithDefaults instantiates a new CreditBankEmploymentWarning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankEmploymentWarningWithDefaults() *CreditBankEmploymentWarning {
	this := CreditBankEmploymentWarning{}
	return &this
}

// GetWarningType returns the WarningType field value
func (o *CreditBankEmploymentWarning) GetWarningType() CreditBankEmploymentWarningType {
	if o == nil {
		var ret CreditBankEmploymentWarningType
		return ret
	}

	return o.WarningType
}

// GetWarningTypeOk returns a tuple with the WarningType field value
// and a boolean to check if the value has been set.
func (o *CreditBankEmploymentWarning) GetWarningTypeOk() (*CreditBankEmploymentWarningType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WarningType, true
}

// SetWarningType sets field value
func (o *CreditBankEmploymentWarning) SetWarningType(v CreditBankEmploymentWarningType) {
	o.WarningType = v
}

// GetWarningCode returns the WarningCode field value
func (o *CreditBankEmploymentWarning) GetWarningCode() CreditBankIncomeWarningCode {
	if o == nil {
		var ret CreditBankIncomeWarningCode
		return ret
	}

	return o.WarningCode
}

// GetWarningCodeOk returns a tuple with the WarningCode field value
// and a boolean to check if the value has been set.
func (o *CreditBankEmploymentWarning) GetWarningCodeOk() (*CreditBankIncomeWarningCode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WarningCode, true
}

// SetWarningCode sets field value
func (o *CreditBankEmploymentWarning) SetWarningCode(v CreditBankIncomeWarningCode) {
	o.WarningCode = v
}

// GetCause returns the Cause field value
func (o *CreditBankEmploymentWarning) GetCause() CreditBankIncomeCause {
	if o == nil {
		var ret CreditBankIncomeCause
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *CreditBankEmploymentWarning) GetCauseOk() (*CreditBankIncomeCause, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *CreditBankEmploymentWarning) SetCause(v CreditBankIncomeCause) {
	o.Cause = v
}

func (o CreditBankEmploymentWarning) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["warning_type"] = o.WarningType
	}
	if true {
		toSerialize["warning_code"] = o.WarningCode
	}
	if true {
		toSerialize["cause"] = o.Cause
	}
	return json.Marshal(toSerialize)
}

type NullableCreditBankEmploymentWarning struct {
	value *CreditBankEmploymentWarning
	isSet bool
}

func (v NullableCreditBankEmploymentWarning) Get() *CreditBankEmploymentWarning {
	return v.value
}

func (v *NullableCreditBankEmploymentWarning) Set(val *CreditBankEmploymentWarning) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankEmploymentWarning) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankEmploymentWarning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankEmploymentWarning(val *CreditBankEmploymentWarning) *NullableCreditBankEmploymentWarning {
	return &NullableCreditBankEmploymentWarning{value: val, isSet: true}
}

func (v NullableCreditBankEmploymentWarning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankEmploymentWarning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


