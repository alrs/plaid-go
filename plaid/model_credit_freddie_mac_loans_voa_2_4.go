/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.343.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// CreditFreddieMacLoansVOA24 A collection of loans that are part of a single deal.
type CreditFreddieMacLoansVOA24 struct {
	LOAN CreditFreddieMacLoanVOA24 `json:"LOAN"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacLoansVOA24 CreditFreddieMacLoansVOA24

// NewCreditFreddieMacLoansVOA24 instantiates a new CreditFreddieMacLoansVOA24 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacLoansVOA24(lOAN CreditFreddieMacLoanVOA24) *CreditFreddieMacLoansVOA24 {
	this := CreditFreddieMacLoansVOA24{}
	this.LOAN = lOAN
	return &this
}

// NewCreditFreddieMacLoansVOA24WithDefaults instantiates a new CreditFreddieMacLoansVOA24 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacLoansVOA24WithDefaults() *CreditFreddieMacLoansVOA24 {
	this := CreditFreddieMacLoansVOA24{}
	return &this
}

// GetLOAN returns the LOAN field value
func (o *CreditFreddieMacLoansVOA24) GetLOAN() CreditFreddieMacLoanVOA24 {
	if o == nil {
		var ret CreditFreddieMacLoanVOA24
		return ret
	}

	return o.LOAN
}

// GetLOANOk returns a tuple with the LOAN field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacLoansVOA24) GetLOANOk() (*CreditFreddieMacLoanVOA24, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LOAN, true
}

// SetLOAN sets field value
func (o *CreditFreddieMacLoansVOA24) SetLOAN(v CreditFreddieMacLoanVOA24) {
	o.LOAN = v
}

func (o CreditFreddieMacLoansVOA24) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["LOAN"] = o.LOAN
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacLoansVOA24) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacLoansVOA24 := _CreditFreddieMacLoansVOA24{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacLoansVOA24); err == nil {
		*o = CreditFreddieMacLoansVOA24(varCreditFreddieMacLoansVOA24)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "LOAN")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieMacLoansVOA24 struct {
	value *CreditFreddieMacLoansVOA24
	isSet bool
}

func (v NullableCreditFreddieMacLoansVOA24) Get() *CreditFreddieMacLoansVOA24 {
	return v.value
}

func (v *NullableCreditFreddieMacLoansVOA24) Set(val *CreditFreddieMacLoansVOA24) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacLoansVOA24) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacLoansVOA24) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacLoansVOA24(val *CreditFreddieMacLoansVOA24) *NullableCreditFreddieMacLoansVOA24 {
	return &NullableCreditFreddieMacLoansVOA24{value: val, isSet: true}
}

func (v NullableCreditFreddieMacLoansVOA24) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacLoansVOA24) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


