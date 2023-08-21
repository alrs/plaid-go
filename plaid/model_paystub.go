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
)

// Paystub An object representing data extracted from the end user's paystub.
type Paystub struct {
	Deductions Deductions `json:"deductions"`
	// An identifier of the document referenced by the document metadata.
	DocId string `json:"doc_id"`
	Earnings Earnings `json:"earnings"`
	Employee Employee `json:"employee"`
	Employer PaystubEmployer `json:"employer"`
	EmploymentDetails *EmploymentDetails `json:"employment_details,omitempty"`
	NetPay NetPay `json:"net_pay"`
	PayPeriodDetails PayPeriodDetails `json:"pay_period_details"`
	PaystubDetails *PaystubDetails `json:"paystub_details,omitempty"`
	IncomeBreakdown *[]IncomeBreakdown `json:"income_breakdown,omitempty"`
	YtdEarnings *PaystubYTDDetails `json:"ytd_earnings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Paystub Paystub

// NewPaystub instantiates a new Paystub object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaystub(deductions Deductions, docId string, earnings Earnings, employee Employee, employer PaystubEmployer, netPay NetPay, payPeriodDetails PayPeriodDetails) *Paystub {
	this := Paystub{}
	this.Deductions = deductions
	this.DocId = docId
	this.Earnings = earnings
	this.Employee = employee
	this.Employer = employer
	this.NetPay = netPay
	this.PayPeriodDetails = payPeriodDetails
	return &this
}

// NewPaystubWithDefaults instantiates a new Paystub object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaystubWithDefaults() *Paystub {
	this := Paystub{}
	return &this
}

// GetDeductions returns the Deductions field value
func (o *Paystub) GetDeductions() Deductions {
	if o == nil {
		var ret Deductions
		return ret
	}

	return o.Deductions
}

// GetDeductionsOk returns a tuple with the Deductions field value
// and a boolean to check if the value has been set.
func (o *Paystub) GetDeductionsOk() (*Deductions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Deductions, true
}

// SetDeductions sets field value
func (o *Paystub) SetDeductions(v Deductions) {
	o.Deductions = v
}

// GetDocId returns the DocId field value
func (o *Paystub) GetDocId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocId
}

// GetDocIdOk returns a tuple with the DocId field value
// and a boolean to check if the value has been set.
func (o *Paystub) GetDocIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocId, true
}

// SetDocId sets field value
func (o *Paystub) SetDocId(v string) {
	o.DocId = v
}

// GetEarnings returns the Earnings field value
func (o *Paystub) GetEarnings() Earnings {
	if o == nil {
		var ret Earnings
		return ret
	}

	return o.Earnings
}

// GetEarningsOk returns a tuple with the Earnings field value
// and a boolean to check if the value has been set.
func (o *Paystub) GetEarningsOk() (*Earnings, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Earnings, true
}

// SetEarnings sets field value
func (o *Paystub) SetEarnings(v Earnings) {
	o.Earnings = v
}

// GetEmployee returns the Employee field value
func (o *Paystub) GetEmployee() Employee {
	if o == nil {
		var ret Employee
		return ret
	}

	return o.Employee
}

// GetEmployeeOk returns a tuple with the Employee field value
// and a boolean to check if the value has been set.
func (o *Paystub) GetEmployeeOk() (*Employee, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Employee, true
}

// SetEmployee sets field value
func (o *Paystub) SetEmployee(v Employee) {
	o.Employee = v
}

// GetEmployer returns the Employer field value
func (o *Paystub) GetEmployer() PaystubEmployer {
	if o == nil {
		var ret PaystubEmployer
		return ret
	}

	return o.Employer
}

// GetEmployerOk returns a tuple with the Employer field value
// and a boolean to check if the value has been set.
func (o *Paystub) GetEmployerOk() (*PaystubEmployer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Employer, true
}

// SetEmployer sets field value
func (o *Paystub) SetEmployer(v PaystubEmployer) {
	o.Employer = v
}

// GetEmploymentDetails returns the EmploymentDetails field value if set, zero value otherwise.
func (o *Paystub) GetEmploymentDetails() EmploymentDetails {
	if o == nil || o.EmploymentDetails == nil {
		var ret EmploymentDetails
		return ret
	}
	return *o.EmploymentDetails
}

// GetEmploymentDetailsOk returns a tuple with the EmploymentDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Paystub) GetEmploymentDetailsOk() (*EmploymentDetails, bool) {
	if o == nil || o.EmploymentDetails == nil {
		return nil, false
	}
	return o.EmploymentDetails, true
}

// HasEmploymentDetails returns a boolean if a field has been set.
func (o *Paystub) HasEmploymentDetails() bool {
	if o != nil && o.EmploymentDetails != nil {
		return true
	}

	return false
}

// SetEmploymentDetails gets a reference to the given EmploymentDetails and assigns it to the EmploymentDetails field.
func (o *Paystub) SetEmploymentDetails(v EmploymentDetails) {
	o.EmploymentDetails = &v
}

// GetNetPay returns the NetPay field value
func (o *Paystub) GetNetPay() NetPay {
	if o == nil {
		var ret NetPay
		return ret
	}

	return o.NetPay
}

// GetNetPayOk returns a tuple with the NetPay field value
// and a boolean to check if the value has been set.
func (o *Paystub) GetNetPayOk() (*NetPay, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NetPay, true
}

// SetNetPay sets field value
func (o *Paystub) SetNetPay(v NetPay) {
	o.NetPay = v
}

// GetPayPeriodDetails returns the PayPeriodDetails field value
func (o *Paystub) GetPayPeriodDetails() PayPeriodDetails {
	if o == nil {
		var ret PayPeriodDetails
		return ret
	}

	return o.PayPeriodDetails
}

// GetPayPeriodDetailsOk returns a tuple with the PayPeriodDetails field value
// and a boolean to check if the value has been set.
func (o *Paystub) GetPayPeriodDetailsOk() (*PayPeriodDetails, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PayPeriodDetails, true
}

// SetPayPeriodDetails sets field value
func (o *Paystub) SetPayPeriodDetails(v PayPeriodDetails) {
	o.PayPeriodDetails = v
}

// GetPaystubDetails returns the PaystubDetails field value if set, zero value otherwise.
func (o *Paystub) GetPaystubDetails() PaystubDetails {
	if o == nil || o.PaystubDetails == nil {
		var ret PaystubDetails
		return ret
	}
	return *o.PaystubDetails
}

// GetPaystubDetailsOk returns a tuple with the PaystubDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Paystub) GetPaystubDetailsOk() (*PaystubDetails, bool) {
	if o == nil || o.PaystubDetails == nil {
		return nil, false
	}
	return o.PaystubDetails, true
}

// HasPaystubDetails returns a boolean if a field has been set.
func (o *Paystub) HasPaystubDetails() bool {
	if o != nil && o.PaystubDetails != nil {
		return true
	}

	return false
}

// SetPaystubDetails gets a reference to the given PaystubDetails and assigns it to the PaystubDetails field.
func (o *Paystub) SetPaystubDetails(v PaystubDetails) {
	o.PaystubDetails = &v
}

// GetIncomeBreakdown returns the IncomeBreakdown field value if set, zero value otherwise.
func (o *Paystub) GetIncomeBreakdown() []IncomeBreakdown {
	if o == nil || o.IncomeBreakdown == nil {
		var ret []IncomeBreakdown
		return ret
	}
	return *o.IncomeBreakdown
}

// GetIncomeBreakdownOk returns a tuple with the IncomeBreakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Paystub) GetIncomeBreakdownOk() (*[]IncomeBreakdown, bool) {
	if o == nil || o.IncomeBreakdown == nil {
		return nil, false
	}
	return o.IncomeBreakdown, true
}

// HasIncomeBreakdown returns a boolean if a field has been set.
func (o *Paystub) HasIncomeBreakdown() bool {
	if o != nil && o.IncomeBreakdown != nil {
		return true
	}

	return false
}

// SetIncomeBreakdown gets a reference to the given []IncomeBreakdown and assigns it to the IncomeBreakdown field.
func (o *Paystub) SetIncomeBreakdown(v []IncomeBreakdown) {
	o.IncomeBreakdown = &v
}

// GetYtdEarnings returns the YtdEarnings field value if set, zero value otherwise.
func (o *Paystub) GetYtdEarnings() PaystubYTDDetails {
	if o == nil || o.YtdEarnings == nil {
		var ret PaystubYTDDetails
		return ret
	}
	return *o.YtdEarnings
}

// GetYtdEarningsOk returns a tuple with the YtdEarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Paystub) GetYtdEarningsOk() (*PaystubYTDDetails, bool) {
	if o == nil || o.YtdEarnings == nil {
		return nil, false
	}
	return o.YtdEarnings, true
}

// HasYtdEarnings returns a boolean if a field has been set.
func (o *Paystub) HasYtdEarnings() bool {
	if o != nil && o.YtdEarnings != nil {
		return true
	}

	return false
}

// SetYtdEarnings gets a reference to the given PaystubYTDDetails and assigns it to the YtdEarnings field.
func (o *Paystub) SetYtdEarnings(v PaystubYTDDetails) {
	o.YtdEarnings = &v
}

func (o Paystub) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["deductions"] = o.Deductions
	}
	if true {
		toSerialize["doc_id"] = o.DocId
	}
	if true {
		toSerialize["earnings"] = o.Earnings
	}
	if true {
		toSerialize["employee"] = o.Employee
	}
	if true {
		toSerialize["employer"] = o.Employer
	}
	if o.EmploymentDetails != nil {
		toSerialize["employment_details"] = o.EmploymentDetails
	}
	if true {
		toSerialize["net_pay"] = o.NetPay
	}
	if true {
		toSerialize["pay_period_details"] = o.PayPeriodDetails
	}
	if o.PaystubDetails != nil {
		toSerialize["paystub_details"] = o.PaystubDetails
	}
	if o.IncomeBreakdown != nil {
		toSerialize["income_breakdown"] = o.IncomeBreakdown
	}
	if o.YtdEarnings != nil {
		toSerialize["ytd_earnings"] = o.YtdEarnings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Paystub) UnmarshalJSON(bytes []byte) (err error) {
	varPaystub := _Paystub{}

	if err = json.Unmarshal(bytes, &varPaystub); err == nil {
		*o = Paystub(varPaystub)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "deductions")
		delete(additionalProperties, "doc_id")
		delete(additionalProperties, "earnings")
		delete(additionalProperties, "employee")
		delete(additionalProperties, "employer")
		delete(additionalProperties, "employment_details")
		delete(additionalProperties, "net_pay")
		delete(additionalProperties, "pay_period_details")
		delete(additionalProperties, "paystub_details")
		delete(additionalProperties, "income_breakdown")
		delete(additionalProperties, "ytd_earnings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaystub struct {
	value *Paystub
	isSet bool
}

func (v NullablePaystub) Get() *Paystub {
	return v.value
}

func (v *NullablePaystub) Set(val *Paystub) {
	v.value = val
	v.isSet = true
}

func (v NullablePaystub) IsSet() bool {
	return v.isSet
}

func (v *NullablePaystub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaystub(val *Paystub) *NullablePaystub {
	return &NullablePaystub{value: val, isSet: true}
}

func (v NullablePaystub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaystub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


