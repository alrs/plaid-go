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

// CraBankIncomeHistoricalSummary The end user's monthly summary for the income source(s).
type CraBankIncomeHistoricalSummary struct {
	// Total amount of earnings for the income source(s) of the user for the month in the summary. This can contain multiple amounts, with each amount denominated in one unique currency.
	TotalAmounts *[]CreditAmountWithCurrency `json:"total_amounts,omitempty"`
	// The start date of the period covered in this monthly summary. This date will be the first day of the month, unless the month being covered is a partial month because it is the first month included in the summary and the date range being requested does not begin with the first day of the month. The date will be returned in an ISO 8601 format (YYYY-MM-DD).
	StartDate *string `json:"start_date,omitempty"`
	// The end date of the period included in this monthly summary. This date will be the last day of the month, unless the month being covered is a partial month because it is the last month included in the summary and the date range being requested does not end with the last day of the month. The date will be returned in an ISO 8601 format (YYYY-MM-DD).
	EndDate *string `json:"end_date,omitempty"`
	Transactions *[]CraBankIncomeTransaction `json:"transactions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CraBankIncomeHistoricalSummary CraBankIncomeHistoricalSummary

// NewCraBankIncomeHistoricalSummary instantiates a new CraBankIncomeHistoricalSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCraBankIncomeHistoricalSummary() *CraBankIncomeHistoricalSummary {
	this := CraBankIncomeHistoricalSummary{}
	return &this
}

// NewCraBankIncomeHistoricalSummaryWithDefaults instantiates a new CraBankIncomeHistoricalSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCraBankIncomeHistoricalSummaryWithDefaults() *CraBankIncomeHistoricalSummary {
	this := CraBankIncomeHistoricalSummary{}
	return &this
}

// GetTotalAmounts returns the TotalAmounts field value if set, zero value otherwise.
func (o *CraBankIncomeHistoricalSummary) GetTotalAmounts() []CreditAmountWithCurrency {
	if o == nil || o.TotalAmounts == nil {
		var ret []CreditAmountWithCurrency
		return ret
	}
	return *o.TotalAmounts
}

// GetTotalAmountsOk returns a tuple with the TotalAmounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraBankIncomeHistoricalSummary) GetTotalAmountsOk() (*[]CreditAmountWithCurrency, bool) {
	if o == nil || o.TotalAmounts == nil {
		return nil, false
	}
	return o.TotalAmounts, true
}

// HasTotalAmounts returns a boolean if a field has been set.
func (o *CraBankIncomeHistoricalSummary) HasTotalAmounts() bool {
	if o != nil && o.TotalAmounts != nil {
		return true
	}

	return false
}

// SetTotalAmounts gets a reference to the given []CreditAmountWithCurrency and assigns it to the TotalAmounts field.
func (o *CraBankIncomeHistoricalSummary) SetTotalAmounts(v []CreditAmountWithCurrency) {
	o.TotalAmounts = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CraBankIncomeHistoricalSummary) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraBankIncomeHistoricalSummary) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CraBankIncomeHistoricalSummary) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *CraBankIncomeHistoricalSummary) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *CraBankIncomeHistoricalSummary) GetEndDate() string {
	if o == nil || o.EndDate == nil {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraBankIncomeHistoricalSummary) GetEndDateOk() (*string, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *CraBankIncomeHistoricalSummary) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *CraBankIncomeHistoricalSummary) SetEndDate(v string) {
	o.EndDate = &v
}

// GetTransactions returns the Transactions field value if set, zero value otherwise.
func (o *CraBankIncomeHistoricalSummary) GetTransactions() []CraBankIncomeTransaction {
	if o == nil || o.Transactions == nil {
		var ret []CraBankIncomeTransaction
		return ret
	}
	return *o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraBankIncomeHistoricalSummary) GetTransactionsOk() (*[]CraBankIncomeTransaction, bool) {
	if o == nil || o.Transactions == nil {
		return nil, false
	}
	return o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *CraBankIncomeHistoricalSummary) HasTransactions() bool {
	if o != nil && o.Transactions != nil {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given []CraBankIncomeTransaction and assigns it to the Transactions field.
func (o *CraBankIncomeHistoricalSummary) SetTransactions(v []CraBankIncomeTransaction) {
	o.Transactions = &v
}

func (o CraBankIncomeHistoricalSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalAmounts != nil {
		toSerialize["total_amounts"] = o.TotalAmounts
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["end_date"] = o.EndDate
	}
	if o.Transactions != nil {
		toSerialize["transactions"] = o.Transactions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CraBankIncomeHistoricalSummary) UnmarshalJSON(bytes []byte) (err error) {
	varCraBankIncomeHistoricalSummary := _CraBankIncomeHistoricalSummary{}

	if err = json.Unmarshal(bytes, &varCraBankIncomeHistoricalSummary); err == nil {
		*o = CraBankIncomeHistoricalSummary(varCraBankIncomeHistoricalSummary)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "total_amounts")
		delete(additionalProperties, "start_date")
		delete(additionalProperties, "end_date")
		delete(additionalProperties, "transactions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCraBankIncomeHistoricalSummary struct {
	value *CraBankIncomeHistoricalSummary
	isSet bool
}

func (v NullableCraBankIncomeHistoricalSummary) Get() *CraBankIncomeHistoricalSummary {
	return v.value
}

func (v *NullableCraBankIncomeHistoricalSummary) Set(val *CraBankIncomeHistoricalSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableCraBankIncomeHistoricalSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableCraBankIncomeHistoricalSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCraBankIncomeHistoricalSummary(val *CraBankIncomeHistoricalSummary) *NullableCraBankIncomeHistoricalSummary {
	return &NullableCraBankIncomeHistoricalSummary{value: val, isSet: true}
}

func (v NullableCraBankIncomeHistoricalSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCraBankIncomeHistoricalSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


