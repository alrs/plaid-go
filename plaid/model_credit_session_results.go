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
)

// CreditSessionResults The set of results for a Link session.
type CreditSessionResults struct {
	// The set of item adds for the Link session.
	ItemAddResults *[]CreditSessionItemAddResult `json:"item_add_results,omitempty"`
	// The set of bank income verifications for the Link session.
	BankIncomeResults *[]CreditSessionBankIncomeResult `json:"bank_income_results,omitempty"`
	// The set of payroll income verifications for the Link session.
	PayrollIncomeResults *[]CreditSessionPayrollIncomeResult `json:"payroll_income_results,omitempty"`
	DocumentIncomeResults NullableCreditSessionDocumentIncomeResult `json:"document_income_results,omitempty"`
}

// NewCreditSessionResults instantiates a new CreditSessionResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditSessionResults() *CreditSessionResults {
	this := CreditSessionResults{}
	return &this
}

// NewCreditSessionResultsWithDefaults instantiates a new CreditSessionResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditSessionResultsWithDefaults() *CreditSessionResults {
	this := CreditSessionResults{}
	return &this
}

// GetItemAddResults returns the ItemAddResults field value if set, zero value otherwise.
func (o *CreditSessionResults) GetItemAddResults() []CreditSessionItemAddResult {
	if o == nil || o.ItemAddResults == nil {
		var ret []CreditSessionItemAddResult
		return ret
	}
	return *o.ItemAddResults
}

// GetItemAddResultsOk returns a tuple with the ItemAddResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditSessionResults) GetItemAddResultsOk() (*[]CreditSessionItemAddResult, bool) {
	if o == nil || o.ItemAddResults == nil {
		return nil, false
	}
	return o.ItemAddResults, true
}

// HasItemAddResults returns a boolean if a field has been set.
func (o *CreditSessionResults) HasItemAddResults() bool {
	if o != nil && o.ItemAddResults != nil {
		return true
	}

	return false
}

// SetItemAddResults gets a reference to the given []CreditSessionItemAddResult and assigns it to the ItemAddResults field.
func (o *CreditSessionResults) SetItemAddResults(v []CreditSessionItemAddResult) {
	o.ItemAddResults = &v
}

// GetBankIncomeResults returns the BankIncomeResults field value if set, zero value otherwise.
func (o *CreditSessionResults) GetBankIncomeResults() []CreditSessionBankIncomeResult {
	if o == nil || o.BankIncomeResults == nil {
		var ret []CreditSessionBankIncomeResult
		return ret
	}
	return *o.BankIncomeResults
}

// GetBankIncomeResultsOk returns a tuple with the BankIncomeResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditSessionResults) GetBankIncomeResultsOk() (*[]CreditSessionBankIncomeResult, bool) {
	if o == nil || o.BankIncomeResults == nil {
		return nil, false
	}
	return o.BankIncomeResults, true
}

// HasBankIncomeResults returns a boolean if a field has been set.
func (o *CreditSessionResults) HasBankIncomeResults() bool {
	if o != nil && o.BankIncomeResults != nil {
		return true
	}

	return false
}

// SetBankIncomeResults gets a reference to the given []CreditSessionBankIncomeResult and assigns it to the BankIncomeResults field.
func (o *CreditSessionResults) SetBankIncomeResults(v []CreditSessionBankIncomeResult) {
	o.BankIncomeResults = &v
}

// GetPayrollIncomeResults returns the PayrollIncomeResults field value if set, zero value otherwise.
func (o *CreditSessionResults) GetPayrollIncomeResults() []CreditSessionPayrollIncomeResult {
	if o == nil || o.PayrollIncomeResults == nil {
		var ret []CreditSessionPayrollIncomeResult
		return ret
	}
	return *o.PayrollIncomeResults
}

// GetPayrollIncomeResultsOk returns a tuple with the PayrollIncomeResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditSessionResults) GetPayrollIncomeResultsOk() (*[]CreditSessionPayrollIncomeResult, bool) {
	if o == nil || o.PayrollIncomeResults == nil {
		return nil, false
	}
	return o.PayrollIncomeResults, true
}

// HasPayrollIncomeResults returns a boolean if a field has been set.
func (o *CreditSessionResults) HasPayrollIncomeResults() bool {
	if o != nil && o.PayrollIncomeResults != nil {
		return true
	}

	return false
}

// SetPayrollIncomeResults gets a reference to the given []CreditSessionPayrollIncomeResult and assigns it to the PayrollIncomeResults field.
func (o *CreditSessionResults) SetPayrollIncomeResults(v []CreditSessionPayrollIncomeResult) {
	o.PayrollIncomeResults = &v
}

// GetDocumentIncomeResults returns the DocumentIncomeResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditSessionResults) GetDocumentIncomeResults() CreditSessionDocumentIncomeResult {
	if o == nil || o.DocumentIncomeResults.Get() == nil {
		var ret CreditSessionDocumentIncomeResult
		return ret
	}
	return *o.DocumentIncomeResults.Get()
}

// GetDocumentIncomeResultsOk returns a tuple with the DocumentIncomeResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditSessionResults) GetDocumentIncomeResultsOk() (*CreditSessionDocumentIncomeResult, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DocumentIncomeResults.Get(), o.DocumentIncomeResults.IsSet()
}

// HasDocumentIncomeResults returns a boolean if a field has been set.
func (o *CreditSessionResults) HasDocumentIncomeResults() bool {
	if o != nil && o.DocumentIncomeResults.IsSet() {
		return true
	}

	return false
}

// SetDocumentIncomeResults gets a reference to the given NullableCreditSessionDocumentIncomeResult and assigns it to the DocumentIncomeResults field.
func (o *CreditSessionResults) SetDocumentIncomeResults(v CreditSessionDocumentIncomeResult) {
	o.DocumentIncomeResults.Set(&v)
}
// SetDocumentIncomeResultsNil sets the value for DocumentIncomeResults to be an explicit nil
func (o *CreditSessionResults) SetDocumentIncomeResultsNil() {
	o.DocumentIncomeResults.Set(nil)
}

// UnsetDocumentIncomeResults ensures that no value is present for DocumentIncomeResults, not even an explicit nil
func (o *CreditSessionResults) UnsetDocumentIncomeResults() {
	o.DocumentIncomeResults.Unset()
}

func (o CreditSessionResults) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ItemAddResults != nil {
		toSerialize["item_add_results"] = o.ItemAddResults
	}
	if o.BankIncomeResults != nil {
		toSerialize["bank_income_results"] = o.BankIncomeResults
	}
	if o.PayrollIncomeResults != nil {
		toSerialize["payroll_income_results"] = o.PayrollIncomeResults
	}
	if o.DocumentIncomeResults.IsSet() {
		toSerialize["document_income_results"] = o.DocumentIncomeResults.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCreditSessionResults struct {
	value *CreditSessionResults
	isSet bool
}

func (v NullableCreditSessionResults) Get() *CreditSessionResults {
	return v.value
}

func (v *NullableCreditSessionResults) Set(val *CreditSessionResults) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditSessionResults) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditSessionResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditSessionResults(val *CreditSessionResults) *NullableCreditSessionResults {
	return &NullableCreditSessionResults{value: val, isSet: true}
}

func (v NullableCreditSessionResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditSessionResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


