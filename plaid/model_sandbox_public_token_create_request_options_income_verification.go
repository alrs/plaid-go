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

// SandboxPublicTokenCreateRequestOptionsIncomeVerification A set of parameters for income verification options. This field is required if `income_verification` is included in the `initial_products` array.
type SandboxPublicTokenCreateRequestOptionsIncomeVerification struct {
	// The types of source income data that users will be permitted to share. Options include `bank` and `payroll`. Currently you can only specify one of these options.
	IncomeSourceTypes *[]IncomeVerificationSourceType `json:"income_source_types,omitempty"`
	BankIncome *SandboxPublicTokenCreateRequestIncomeVerificationBankIncome `json:"bank_income,omitempty"`
}

// NewSandboxPublicTokenCreateRequestOptionsIncomeVerification instantiates a new SandboxPublicTokenCreateRequestOptionsIncomeVerification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxPublicTokenCreateRequestOptionsIncomeVerification() *SandboxPublicTokenCreateRequestOptionsIncomeVerification {
	this := SandboxPublicTokenCreateRequestOptionsIncomeVerification{}
	return &this
}

// NewSandboxPublicTokenCreateRequestOptionsIncomeVerificationWithDefaults instantiates a new SandboxPublicTokenCreateRequestOptionsIncomeVerification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxPublicTokenCreateRequestOptionsIncomeVerificationWithDefaults() *SandboxPublicTokenCreateRequestOptionsIncomeVerification {
	this := SandboxPublicTokenCreateRequestOptionsIncomeVerification{}
	return &this
}

// GetIncomeSourceTypes returns the IncomeSourceTypes field value if set, zero value otherwise.
func (o *SandboxPublicTokenCreateRequestOptionsIncomeVerification) GetIncomeSourceTypes() []IncomeVerificationSourceType {
	if o == nil || o.IncomeSourceTypes == nil {
		var ret []IncomeVerificationSourceType
		return ret
	}
	return *o.IncomeSourceTypes
}

// GetIncomeSourceTypesOk returns a tuple with the IncomeSourceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxPublicTokenCreateRequestOptionsIncomeVerification) GetIncomeSourceTypesOk() (*[]IncomeVerificationSourceType, bool) {
	if o == nil || o.IncomeSourceTypes == nil {
		return nil, false
	}
	return o.IncomeSourceTypes, true
}

// HasIncomeSourceTypes returns a boolean if a field has been set.
func (o *SandboxPublicTokenCreateRequestOptionsIncomeVerification) HasIncomeSourceTypes() bool {
	if o != nil && o.IncomeSourceTypes != nil {
		return true
	}

	return false
}

// SetIncomeSourceTypes gets a reference to the given []IncomeVerificationSourceType and assigns it to the IncomeSourceTypes field.
func (o *SandboxPublicTokenCreateRequestOptionsIncomeVerification) SetIncomeSourceTypes(v []IncomeVerificationSourceType) {
	o.IncomeSourceTypes = &v
}

// GetBankIncome returns the BankIncome field value if set, zero value otherwise.
func (o *SandboxPublicTokenCreateRequestOptionsIncomeVerification) GetBankIncome() SandboxPublicTokenCreateRequestIncomeVerificationBankIncome {
	if o == nil || o.BankIncome == nil {
		var ret SandboxPublicTokenCreateRequestIncomeVerificationBankIncome
		return ret
	}
	return *o.BankIncome
}

// GetBankIncomeOk returns a tuple with the BankIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxPublicTokenCreateRequestOptionsIncomeVerification) GetBankIncomeOk() (*SandboxPublicTokenCreateRequestIncomeVerificationBankIncome, bool) {
	if o == nil || o.BankIncome == nil {
		return nil, false
	}
	return o.BankIncome, true
}

// HasBankIncome returns a boolean if a field has been set.
func (o *SandboxPublicTokenCreateRequestOptionsIncomeVerification) HasBankIncome() bool {
	if o != nil && o.BankIncome != nil {
		return true
	}

	return false
}

// SetBankIncome gets a reference to the given SandboxPublicTokenCreateRequestIncomeVerificationBankIncome and assigns it to the BankIncome field.
func (o *SandboxPublicTokenCreateRequestOptionsIncomeVerification) SetBankIncome(v SandboxPublicTokenCreateRequestIncomeVerificationBankIncome) {
	o.BankIncome = &v
}

func (o SandboxPublicTokenCreateRequestOptionsIncomeVerification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IncomeSourceTypes != nil {
		toSerialize["income_source_types"] = o.IncomeSourceTypes
	}
	if o.BankIncome != nil {
		toSerialize["bank_income"] = o.BankIncome
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxPublicTokenCreateRequestOptionsIncomeVerification struct {
	value *SandboxPublicTokenCreateRequestOptionsIncomeVerification
	isSet bool
}

func (v NullableSandboxPublicTokenCreateRequestOptionsIncomeVerification) Get() *SandboxPublicTokenCreateRequestOptionsIncomeVerification {
	return v.value
}

func (v *NullableSandboxPublicTokenCreateRequestOptionsIncomeVerification) Set(val *SandboxPublicTokenCreateRequestOptionsIncomeVerification) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxPublicTokenCreateRequestOptionsIncomeVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxPublicTokenCreateRequestOptionsIncomeVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxPublicTokenCreateRequestOptionsIncomeVerification(val *SandboxPublicTokenCreateRequestOptionsIncomeVerification) *NullableSandboxPublicTokenCreateRequestOptionsIncomeVerification {
	return &NullableSandboxPublicTokenCreateRequestOptionsIncomeVerification{value: val, isSet: true}
}

func (v NullableSandboxPublicTokenCreateRequestOptionsIncomeVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxPublicTokenCreateRequestOptionsIncomeVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


