/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.214.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// CreditFreddieMacLoanVOA24 Information specific to a mortgage loan agreement between one or more borrowers and a mortgage lender.
type CreditFreddieMacLoanVOA24 struct {
	LOAN_IDENTIFIERS CreditFreddieMacLoanIdentifiersVOA24 `json:"LOAN_IDENTIFIERS"`
	// Type of loan. The value can only be \"SubjectLoan\"
	LoanRoleType string `json:"LoanRoleType"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacLoanVOA24 CreditFreddieMacLoanVOA24

// NewCreditFreddieMacLoanVOA24 instantiates a new CreditFreddieMacLoanVOA24 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacLoanVOA24(lOANIDENTIFIERS CreditFreddieMacLoanIdentifiersVOA24, loanRoleType string) *CreditFreddieMacLoanVOA24 {
	this := CreditFreddieMacLoanVOA24{}
	this.LOAN_IDENTIFIERS = lOANIDENTIFIERS
	this.LoanRoleType = loanRoleType
	return &this
}

// NewCreditFreddieMacLoanVOA24WithDefaults instantiates a new CreditFreddieMacLoanVOA24 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacLoanVOA24WithDefaults() *CreditFreddieMacLoanVOA24 {
	this := CreditFreddieMacLoanVOA24{}
	return &this
}

// GetLOAN_IDENTIFIERS returns the LOAN_IDENTIFIERS field value
func (o *CreditFreddieMacLoanVOA24) GetLOAN_IDENTIFIERS() CreditFreddieMacLoanIdentifiersVOA24 {
	if o == nil {
		var ret CreditFreddieMacLoanIdentifiersVOA24
		return ret
	}

	return o.LOAN_IDENTIFIERS
}

// GetLOAN_IDENTIFIERSOk returns a tuple with the LOAN_IDENTIFIERS field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacLoanVOA24) GetLOAN_IDENTIFIERSOk() (*CreditFreddieMacLoanIdentifiersVOA24, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LOAN_IDENTIFIERS, true
}

// SetLOAN_IDENTIFIERS sets field value
func (o *CreditFreddieMacLoanVOA24) SetLOAN_IDENTIFIERS(v CreditFreddieMacLoanIdentifiersVOA24) {
	o.LOAN_IDENTIFIERS = v
}

// GetLoanRoleType returns the LoanRoleType field value
func (o *CreditFreddieMacLoanVOA24) GetLoanRoleType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoanRoleType
}

// GetLoanRoleTypeOk returns a tuple with the LoanRoleType field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacLoanVOA24) GetLoanRoleTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LoanRoleType, true
}

// SetLoanRoleType sets field value
func (o *CreditFreddieMacLoanVOA24) SetLoanRoleType(v string) {
	o.LoanRoleType = v
}

func (o CreditFreddieMacLoanVOA24) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["LOAN_IDENTIFIERS"] = o.LOAN_IDENTIFIERS
	}
	if true {
		toSerialize["LoanRoleType"] = o.LoanRoleType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacLoanVOA24) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacLoanVOA24 := _CreditFreddieMacLoanVOA24{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacLoanVOA24); err == nil {
		*o = CreditFreddieMacLoanVOA24(varCreditFreddieMacLoanVOA24)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "LOAN_IDENTIFIERS")
		delete(additionalProperties, "LoanRoleType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieMacLoanVOA24 struct {
	value *CreditFreddieMacLoanVOA24
	isSet bool
}

func (v NullableCreditFreddieMacLoanVOA24) Get() *CreditFreddieMacLoanVOA24 {
	return v.value
}

func (v *NullableCreditFreddieMacLoanVOA24) Set(val *CreditFreddieMacLoanVOA24) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacLoanVOA24) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacLoanVOA24) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacLoanVOA24(val *CreditFreddieMacLoanVOA24) *NullableCreditFreddieMacLoanVOA24 {
	return &NullableCreditFreddieMacLoanVOA24{value: val, isSet: true}
}

func (v NullableCreditFreddieMacLoanVOA24) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacLoanVOA24) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

