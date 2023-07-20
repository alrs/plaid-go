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
	"time"
)

// CreditBankEmploymentReport The report of the Bank Employment data for an end user.
type CreditBankEmploymentReport struct {
	// The unique identifier associated with the Bank Employment Report.
	BankEmploymentReportId string `json:"bank_employment_report_id"`
	// The time when the Bank Employment Report was generated, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (e.g. \"2018-04-12T03:32:11Z\").
	GeneratedTime time.Time `json:"generated_time"`
	// The number of days requested by the customer for the Bank Employment Report.
	DaysRequested int32 `json:"days_requested"`
	// The list of Items in the report along with the associated metadata about the Item.
	Items []CreditBankEmploymentItem `json:"items"`
	// If data from the Bank Employment report was unable to be retrieved, the warnings will contain information about the error that caused the data to be incomplete.
	Warnings []CreditBankEmploymentWarning `json:"warnings"`
}

// NewCreditBankEmploymentReport instantiates a new CreditBankEmploymentReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankEmploymentReport(bankEmploymentReportId string, generatedTime time.Time, daysRequested int32, items []CreditBankEmploymentItem, warnings []CreditBankEmploymentWarning) *CreditBankEmploymentReport {
	this := CreditBankEmploymentReport{}
	this.BankEmploymentReportId = bankEmploymentReportId
	this.GeneratedTime = generatedTime
	this.DaysRequested = daysRequested
	this.Items = items
	this.Warnings = warnings
	return &this
}

// NewCreditBankEmploymentReportWithDefaults instantiates a new CreditBankEmploymentReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankEmploymentReportWithDefaults() *CreditBankEmploymentReport {
	this := CreditBankEmploymentReport{}
	return &this
}

// GetBankEmploymentReportId returns the BankEmploymentReportId field value
func (o *CreditBankEmploymentReport) GetBankEmploymentReportId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankEmploymentReportId
}

// GetBankEmploymentReportIdOk returns a tuple with the BankEmploymentReportId field value
// and a boolean to check if the value has been set.
func (o *CreditBankEmploymentReport) GetBankEmploymentReportIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankEmploymentReportId, true
}

// SetBankEmploymentReportId sets field value
func (o *CreditBankEmploymentReport) SetBankEmploymentReportId(v string) {
	o.BankEmploymentReportId = v
}

// GetGeneratedTime returns the GeneratedTime field value
func (o *CreditBankEmploymentReport) GetGeneratedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.GeneratedTime
}

// GetGeneratedTimeOk returns a tuple with the GeneratedTime field value
// and a boolean to check if the value has been set.
func (o *CreditBankEmploymentReport) GetGeneratedTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GeneratedTime, true
}

// SetGeneratedTime sets field value
func (o *CreditBankEmploymentReport) SetGeneratedTime(v time.Time) {
	o.GeneratedTime = v
}

// GetDaysRequested returns the DaysRequested field value
func (o *CreditBankEmploymentReport) GetDaysRequested() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DaysRequested
}

// GetDaysRequestedOk returns a tuple with the DaysRequested field value
// and a boolean to check if the value has been set.
func (o *CreditBankEmploymentReport) GetDaysRequestedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DaysRequested, true
}

// SetDaysRequested sets field value
func (o *CreditBankEmploymentReport) SetDaysRequested(v int32) {
	o.DaysRequested = v
}

// GetItems returns the Items field value
func (o *CreditBankEmploymentReport) GetItems() []CreditBankEmploymentItem {
	if o == nil {
		var ret []CreditBankEmploymentItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CreditBankEmploymentReport) GetItemsOk() (*[]CreditBankEmploymentItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *CreditBankEmploymentReport) SetItems(v []CreditBankEmploymentItem) {
	o.Items = v
}

// GetWarnings returns the Warnings field value
func (o *CreditBankEmploymentReport) GetWarnings() []CreditBankEmploymentWarning {
	if o == nil {
		var ret []CreditBankEmploymentWarning
		return ret
	}

	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *CreditBankEmploymentReport) GetWarningsOk() (*[]CreditBankEmploymentWarning, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Warnings, true
}

// SetWarnings sets field value
func (o *CreditBankEmploymentReport) SetWarnings(v []CreditBankEmploymentWarning) {
	o.Warnings = v
}

func (o CreditBankEmploymentReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bank_employment_report_id"] = o.BankEmploymentReportId
	}
	if true {
		toSerialize["generated_time"] = o.GeneratedTime
	}
	if true {
		toSerialize["days_requested"] = o.DaysRequested
	}
	if true {
		toSerialize["items"] = o.Items
	}
	if true {
		toSerialize["warnings"] = o.Warnings
	}
	return json.Marshal(toSerialize)
}

type NullableCreditBankEmploymentReport struct {
	value *CreditBankEmploymentReport
	isSet bool
}

func (v NullableCreditBankEmploymentReport) Get() *CreditBankEmploymentReport {
	return v.value
}

func (v *NullableCreditBankEmploymentReport) Set(val *CreditBankEmploymentReport) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankEmploymentReport) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankEmploymentReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankEmploymentReport(val *CreditBankEmploymentReport) *NullableCreditBankEmploymentReport {
	return &NullableCreditBankEmploymentReport{value: val, isSet: true}
}

func (v NullableCreditBankEmploymentReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankEmploymentReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


