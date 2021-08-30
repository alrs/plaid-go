/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.26.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// StudentLoanRepaymentModel Student loan repayment information used to configure Sandbox test data for the Liabilities product
type StudentLoanRepaymentModel struct {
	// The only currently supported value for this field is `standard`.
	Type string `json:"type"`
	// Configures the number of months before repayment starts.
	NonRepaymentMonths float32 `json:"non_repayment_months"`
	// Configures the number of months of repayments before the loan is paid off.
	RepaymentMonths float32 `json:"repayment_months"`
	AdditionalProperties map[string]interface{}
}

type _StudentLoanRepaymentModel StudentLoanRepaymentModel

// NewStudentLoanRepaymentModel instantiates a new StudentLoanRepaymentModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStudentLoanRepaymentModel(type_ string, nonRepaymentMonths float32, repaymentMonths float32) *StudentLoanRepaymentModel {
	this := StudentLoanRepaymentModel{}
	this.Type = type_
	this.NonRepaymentMonths = nonRepaymentMonths
	this.RepaymentMonths = repaymentMonths
	return &this
}

// NewStudentLoanRepaymentModelWithDefaults instantiates a new StudentLoanRepaymentModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStudentLoanRepaymentModelWithDefaults() *StudentLoanRepaymentModel {
	this := StudentLoanRepaymentModel{}
	return &this
}

// GetType returns the Type field value
func (o *StudentLoanRepaymentModel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StudentLoanRepaymentModel) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StudentLoanRepaymentModel) SetType(v string) {
	o.Type = v
}

// GetNonRepaymentMonths returns the NonRepaymentMonths field value
func (o *StudentLoanRepaymentModel) GetNonRepaymentMonths() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NonRepaymentMonths
}

// GetNonRepaymentMonthsOk returns a tuple with the NonRepaymentMonths field value
// and a boolean to check if the value has been set.
func (o *StudentLoanRepaymentModel) GetNonRepaymentMonthsOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NonRepaymentMonths, true
}

// SetNonRepaymentMonths sets field value
func (o *StudentLoanRepaymentModel) SetNonRepaymentMonths(v float32) {
	o.NonRepaymentMonths = v
}

// GetRepaymentMonths returns the RepaymentMonths field value
func (o *StudentLoanRepaymentModel) GetRepaymentMonths() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RepaymentMonths
}

// GetRepaymentMonthsOk returns a tuple with the RepaymentMonths field value
// and a boolean to check if the value has been set.
func (o *StudentLoanRepaymentModel) GetRepaymentMonthsOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RepaymentMonths, true
}

// SetRepaymentMonths sets field value
func (o *StudentLoanRepaymentModel) SetRepaymentMonths(v float32) {
	o.RepaymentMonths = v
}

func (o StudentLoanRepaymentModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["non_repayment_months"] = o.NonRepaymentMonths
	}
	if true {
		toSerialize["repayment_months"] = o.RepaymentMonths
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StudentLoanRepaymentModel) UnmarshalJSON(bytes []byte) (err error) {
	varStudentLoanRepaymentModel := _StudentLoanRepaymentModel{}

	if err = json.Unmarshal(bytes, &varStudentLoanRepaymentModel); err == nil {
		*o = StudentLoanRepaymentModel(varStudentLoanRepaymentModel)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "non_repayment_months")
		delete(additionalProperties, "repayment_months")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStudentLoanRepaymentModel struct {
	value *StudentLoanRepaymentModel
	isSet bool
}

func (v NullableStudentLoanRepaymentModel) Get() *StudentLoanRepaymentModel {
	return v.value
}

func (v *NullableStudentLoanRepaymentModel) Set(val *StudentLoanRepaymentModel) {
	v.value = val
	v.isSet = true
}

func (v NullableStudentLoanRepaymentModel) IsSet() bool {
	return v.isSet
}

func (v *NullableStudentLoanRepaymentModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStudentLoanRepaymentModel(val *StudentLoanRepaymentModel) *NullableStudentLoanRepaymentModel {
	return &NullableStudentLoanRepaymentModel{value: val, isSet: true}
}

func (v NullableStudentLoanRepaymentModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStudentLoanRepaymentModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


