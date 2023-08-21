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

// CraBankIncomeSource Detailed information for the income source.
type CraBankIncomeSource struct {
	// A unique identifier for an income source.
	IncomeSourceId *string `json:"income_source_id,omitempty"`
	// The most common name or original description for the underlying income transactions.
	IncomeDescription *string `json:"income_description,omitempty"`
	IncomeCategory *CreditBankIncomeCategory `json:"income_category,omitempty"`
	// Minimum of all dates within the specific income sources in the user's bank account for days requested by the client. The date will be returned in an ISO 8601 format (YYYY-MM-DD).
	StartDate *string `json:"start_date,omitempty"`
	// Maximum of all dates within the specific income sources in the user’s bank account for days requested by the client. The date will be returned in an ISO 8601 format (YYYY-MM-DD).
	EndDate *string `json:"end_date,omitempty"`
	PayFrequency *CreditBankIncomePayFrequency `json:"pay_frequency,omitempty"`
	// Total amount of earnings in the user’s bank account for the specific income source for days requested by the client.
	TotalAmount *float32 `json:"total_amount,omitempty"`
	// The ISO 4217 currency code of the amount or balance.
	IsoCurrencyCode NullableString `json:"iso_currency_code,omitempty"`
	// The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.
	UnofficialCurrencyCode NullableString `json:"unofficial_currency_code,omitempty"`
	// Number of transactions for the income source within the start and end date.
	TransactionCount *int32 `json:"transaction_count,omitempty"`
	// The expected date of the end user’s next paycheck for the income source. The date will be returned in an ISO 8601 format (YYYY-MM-DD).
	NextPaymentDate NullableString `json:"next_payment_date,omitempty"`
	// An estimate of the average gross monthly income based on the historical net amount and income category for the income source(s).
	HistoricalAverageMonthlyGrossIncome NullableFloat32 `json:"historical_average_monthly_gross_income,omitempty"`
	// The average monthly net income amount estimated based on the historical data for the income source(s).
	HistoricalAverageMonthlyIncome NullableFloat32 `json:"historical_average_monthly_income,omitempty"`
	// The predicted average monthly net income amount for the income source(s).
	ForecastedAverageMonthlyIncome NullableFloat32 `json:"forecasted_average_monthly_income,omitempty"`
	Employer *CraBankIncomeEmployer `json:"employer,omitempty"`
	HistoricalSummary *[]CraBankIncomeHistoricalSummary `json:"historical_summary,omitempty"`
}

// NewCraBankIncomeSource instantiates a new CraBankIncomeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCraBankIncomeSource() *CraBankIncomeSource {
	this := CraBankIncomeSource{}
	return &this
}

// NewCraBankIncomeSourceWithDefaults instantiates a new CraBankIncomeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCraBankIncomeSourceWithDefaults() *CraBankIncomeSource {
	this := CraBankIncomeSource{}
	return &this
}

// GetIncomeSourceId returns the IncomeSourceId field value if set, zero value otherwise.
func (o *CraBankIncomeSource) GetIncomeSourceId() string {
	if o == nil || o.IncomeSourceId == nil {
		var ret string
		return ret
	}
	return *o.IncomeSourceId
}

// GetIncomeSourceIdOk returns a tuple with the IncomeSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraBankIncomeSource) GetIncomeSourceIdOk() (*string, bool) {
	if o == nil || o.IncomeSourceId == nil {
		return nil, false
	}
	return o.IncomeSourceId, true
}

// HasIncomeSourceId returns a boolean if a field has been set.
func (o *CraBankIncomeSource) HasIncomeSourceId() bool {
	if o != nil && o.IncomeSourceId != nil {
		return true
	}

	return false
}

// SetIncomeSourceId gets a reference to the given string and assigns it to the IncomeSourceId field.
func (o *CraBankIncomeSource) SetIncomeSourceId(v string) {
	o.IncomeSourceId = &v
}

// GetIncomeDescription returns the IncomeDescription field value if set, zero value otherwise.
func (o *CraBankIncomeSource) GetIncomeDescription() string {
	if o == nil || o.IncomeDescription == nil {
		var ret string
		return ret
	}
	return *o.IncomeDescription
}

// GetIncomeDescriptionOk returns a tuple with the IncomeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraBankIncomeSource) GetIncomeDescriptionOk() (*string, bool) {
	if o == nil || o.IncomeDescription == nil {
		return nil, false
	}
	return o.IncomeDescription, true
}

// HasIncomeDescription returns a boolean if a field has been set.
func (o *CraBankIncomeSource) HasIncomeDescription() bool {
	if o != nil && o.IncomeDescription != nil {
		return true
	}

	return false
}

// SetIncomeDescription gets a reference to the given string and assigns it to the IncomeDescription field.
func (o *CraBankIncomeSource) SetIncomeDescription(v string) {
	o.IncomeDescription = &v
}

// GetIncomeCategory returns the IncomeCategory field value if set, zero value otherwise.
func (o *CraBankIncomeSource) GetIncomeCategory() CreditBankIncomeCategory {
	if o == nil || o.IncomeCategory == nil {
		var ret CreditBankIncomeCategory
		return ret
	}
	return *o.IncomeCategory
}

// GetIncomeCategoryOk returns a tuple with the IncomeCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraBankIncomeSource) GetIncomeCategoryOk() (*CreditBankIncomeCategory, bool) {
	if o == nil || o.IncomeCategory == nil {
		return nil, false
	}
	return o.IncomeCategory, true
}

// HasIncomeCategory returns a boolean if a field has been set.
func (o *CraBankIncomeSource) HasIncomeCategory() bool {
	if o != nil && o.IncomeCategory != nil {
		return true
	}

	return false
}

// SetIncomeCategory gets a reference to the given CreditBankIncomeCategory and assigns it to the IncomeCategory field.
func (o *CraBankIncomeSource) SetIncomeCategory(v CreditBankIncomeCategory) {
	o.IncomeCategory = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CraBankIncomeSource) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraBankIncomeSource) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CraBankIncomeSource) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *CraBankIncomeSource) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *CraBankIncomeSource) GetEndDate() string {
	if o == nil || o.EndDate == nil {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraBankIncomeSource) GetEndDateOk() (*string, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *CraBankIncomeSource) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *CraBankIncomeSource) SetEndDate(v string) {
	o.EndDate = &v
}

// GetPayFrequency returns the PayFrequency field value if set, zero value otherwise.
func (o *CraBankIncomeSource) GetPayFrequency() CreditBankIncomePayFrequency {
	if o == nil || o.PayFrequency == nil {
		var ret CreditBankIncomePayFrequency
		return ret
	}
	return *o.PayFrequency
}

// GetPayFrequencyOk returns a tuple with the PayFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraBankIncomeSource) GetPayFrequencyOk() (*CreditBankIncomePayFrequency, bool) {
	if o == nil || o.PayFrequency == nil {
		return nil, false
	}
	return o.PayFrequency, true
}

// HasPayFrequency returns a boolean if a field has been set.
func (o *CraBankIncomeSource) HasPayFrequency() bool {
	if o != nil && o.PayFrequency != nil {
		return true
	}

	return false
}

// SetPayFrequency gets a reference to the given CreditBankIncomePayFrequency and assigns it to the PayFrequency field.
func (o *CraBankIncomeSource) SetPayFrequency(v CreditBankIncomePayFrequency) {
	o.PayFrequency = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *CraBankIncomeSource) GetTotalAmount() float32 {
	if o == nil || o.TotalAmount == nil {
		var ret float32
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraBankIncomeSource) GetTotalAmountOk() (*float32, bool) {
	if o == nil || o.TotalAmount == nil {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *CraBankIncomeSource) HasTotalAmount() bool {
	if o != nil && o.TotalAmount != nil {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given float32 and assigns it to the TotalAmount field.
func (o *CraBankIncomeSource) SetTotalAmount(v float32) {
	o.TotalAmount = &v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CraBankIncomeSource) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.IsoCurrencyCode.Get()
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CraBankIncomeSource) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsoCurrencyCode.Get(), o.IsoCurrencyCode.IsSet()
}

// HasIsoCurrencyCode returns a boolean if a field has been set.
func (o *CraBankIncomeSource) HasIsoCurrencyCode() bool {
	if o != nil && o.IsoCurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetIsoCurrencyCode gets a reference to the given NullableString and assigns it to the IsoCurrencyCode field.
func (o *CraBankIncomeSource) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode.Set(&v)
}
// SetIsoCurrencyCodeNil sets the value for IsoCurrencyCode to be an explicit nil
func (o *CraBankIncomeSource) SetIsoCurrencyCodeNil() {
	o.IsoCurrencyCode.Set(nil)
}

// UnsetIsoCurrencyCode ensures that no value is present for IsoCurrencyCode, not even an explicit nil
func (o *CraBankIncomeSource) UnsetIsoCurrencyCode() {
	o.IsoCurrencyCode.Unset()
}

// GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CraBankIncomeSource) GetUnofficialCurrencyCode() string {
	if o == nil || o.UnofficialCurrencyCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.UnofficialCurrencyCode.Get()
}

// GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CraBankIncomeSource) GetUnofficialCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnofficialCurrencyCode.Get(), o.UnofficialCurrencyCode.IsSet()
}

// HasUnofficialCurrencyCode returns a boolean if a field has been set.
func (o *CraBankIncomeSource) HasUnofficialCurrencyCode() bool {
	if o != nil && o.UnofficialCurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetUnofficialCurrencyCode gets a reference to the given NullableString and assigns it to the UnofficialCurrencyCode field.
func (o *CraBankIncomeSource) SetUnofficialCurrencyCode(v string) {
	o.UnofficialCurrencyCode.Set(&v)
}
// SetUnofficialCurrencyCodeNil sets the value for UnofficialCurrencyCode to be an explicit nil
func (o *CraBankIncomeSource) SetUnofficialCurrencyCodeNil() {
	o.UnofficialCurrencyCode.Set(nil)
}

// UnsetUnofficialCurrencyCode ensures that no value is present for UnofficialCurrencyCode, not even an explicit nil
func (o *CraBankIncomeSource) UnsetUnofficialCurrencyCode() {
	o.UnofficialCurrencyCode.Unset()
}

// GetTransactionCount returns the TransactionCount field value if set, zero value otherwise.
func (o *CraBankIncomeSource) GetTransactionCount() int32 {
	if o == nil || o.TransactionCount == nil {
		var ret int32
		return ret
	}
	return *o.TransactionCount
}

// GetTransactionCountOk returns a tuple with the TransactionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraBankIncomeSource) GetTransactionCountOk() (*int32, bool) {
	if o == nil || o.TransactionCount == nil {
		return nil, false
	}
	return o.TransactionCount, true
}

// HasTransactionCount returns a boolean if a field has been set.
func (o *CraBankIncomeSource) HasTransactionCount() bool {
	if o != nil && o.TransactionCount != nil {
		return true
	}

	return false
}

// SetTransactionCount gets a reference to the given int32 and assigns it to the TransactionCount field.
func (o *CraBankIncomeSource) SetTransactionCount(v int32) {
	o.TransactionCount = &v
}

// GetNextPaymentDate returns the NextPaymentDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CraBankIncomeSource) GetNextPaymentDate() string {
	if o == nil || o.NextPaymentDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.NextPaymentDate.Get()
}

// GetNextPaymentDateOk returns a tuple with the NextPaymentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CraBankIncomeSource) GetNextPaymentDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextPaymentDate.Get(), o.NextPaymentDate.IsSet()
}

// HasNextPaymentDate returns a boolean if a field has been set.
func (o *CraBankIncomeSource) HasNextPaymentDate() bool {
	if o != nil && o.NextPaymentDate.IsSet() {
		return true
	}

	return false
}

// SetNextPaymentDate gets a reference to the given NullableString and assigns it to the NextPaymentDate field.
func (o *CraBankIncomeSource) SetNextPaymentDate(v string) {
	o.NextPaymentDate.Set(&v)
}
// SetNextPaymentDateNil sets the value for NextPaymentDate to be an explicit nil
func (o *CraBankIncomeSource) SetNextPaymentDateNil() {
	o.NextPaymentDate.Set(nil)
}

// UnsetNextPaymentDate ensures that no value is present for NextPaymentDate, not even an explicit nil
func (o *CraBankIncomeSource) UnsetNextPaymentDate() {
	o.NextPaymentDate.Unset()
}

// GetHistoricalAverageMonthlyGrossIncome returns the HistoricalAverageMonthlyGrossIncome field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CraBankIncomeSource) GetHistoricalAverageMonthlyGrossIncome() float32 {
	if o == nil || o.HistoricalAverageMonthlyGrossIncome.Get() == nil {
		var ret float32
		return ret
	}
	return *o.HistoricalAverageMonthlyGrossIncome.Get()
}

// GetHistoricalAverageMonthlyGrossIncomeOk returns a tuple with the HistoricalAverageMonthlyGrossIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CraBankIncomeSource) GetHistoricalAverageMonthlyGrossIncomeOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HistoricalAverageMonthlyGrossIncome.Get(), o.HistoricalAverageMonthlyGrossIncome.IsSet()
}

// HasHistoricalAverageMonthlyGrossIncome returns a boolean if a field has been set.
func (o *CraBankIncomeSource) HasHistoricalAverageMonthlyGrossIncome() bool {
	if o != nil && o.HistoricalAverageMonthlyGrossIncome.IsSet() {
		return true
	}

	return false
}

// SetHistoricalAverageMonthlyGrossIncome gets a reference to the given NullableFloat32 and assigns it to the HistoricalAverageMonthlyGrossIncome field.
func (o *CraBankIncomeSource) SetHistoricalAverageMonthlyGrossIncome(v float32) {
	o.HistoricalAverageMonthlyGrossIncome.Set(&v)
}
// SetHistoricalAverageMonthlyGrossIncomeNil sets the value for HistoricalAverageMonthlyGrossIncome to be an explicit nil
func (o *CraBankIncomeSource) SetHistoricalAverageMonthlyGrossIncomeNil() {
	o.HistoricalAverageMonthlyGrossIncome.Set(nil)
}

// UnsetHistoricalAverageMonthlyGrossIncome ensures that no value is present for HistoricalAverageMonthlyGrossIncome, not even an explicit nil
func (o *CraBankIncomeSource) UnsetHistoricalAverageMonthlyGrossIncome() {
	o.HistoricalAverageMonthlyGrossIncome.Unset()
}

// GetHistoricalAverageMonthlyIncome returns the HistoricalAverageMonthlyIncome field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CraBankIncomeSource) GetHistoricalAverageMonthlyIncome() float32 {
	if o == nil || o.HistoricalAverageMonthlyIncome.Get() == nil {
		var ret float32
		return ret
	}
	return *o.HistoricalAverageMonthlyIncome.Get()
}

// GetHistoricalAverageMonthlyIncomeOk returns a tuple with the HistoricalAverageMonthlyIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CraBankIncomeSource) GetHistoricalAverageMonthlyIncomeOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HistoricalAverageMonthlyIncome.Get(), o.HistoricalAverageMonthlyIncome.IsSet()
}

// HasHistoricalAverageMonthlyIncome returns a boolean if a field has been set.
func (o *CraBankIncomeSource) HasHistoricalAverageMonthlyIncome() bool {
	if o != nil && o.HistoricalAverageMonthlyIncome.IsSet() {
		return true
	}

	return false
}

// SetHistoricalAverageMonthlyIncome gets a reference to the given NullableFloat32 and assigns it to the HistoricalAverageMonthlyIncome field.
func (o *CraBankIncomeSource) SetHistoricalAverageMonthlyIncome(v float32) {
	o.HistoricalAverageMonthlyIncome.Set(&v)
}
// SetHistoricalAverageMonthlyIncomeNil sets the value for HistoricalAverageMonthlyIncome to be an explicit nil
func (o *CraBankIncomeSource) SetHistoricalAverageMonthlyIncomeNil() {
	o.HistoricalAverageMonthlyIncome.Set(nil)
}

// UnsetHistoricalAverageMonthlyIncome ensures that no value is present for HistoricalAverageMonthlyIncome, not even an explicit nil
func (o *CraBankIncomeSource) UnsetHistoricalAverageMonthlyIncome() {
	o.HistoricalAverageMonthlyIncome.Unset()
}

// GetForecastedAverageMonthlyIncome returns the ForecastedAverageMonthlyIncome field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CraBankIncomeSource) GetForecastedAverageMonthlyIncome() float32 {
	if o == nil || o.ForecastedAverageMonthlyIncome.Get() == nil {
		var ret float32
		return ret
	}
	return *o.ForecastedAverageMonthlyIncome.Get()
}

// GetForecastedAverageMonthlyIncomeOk returns a tuple with the ForecastedAverageMonthlyIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CraBankIncomeSource) GetForecastedAverageMonthlyIncomeOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ForecastedAverageMonthlyIncome.Get(), o.ForecastedAverageMonthlyIncome.IsSet()
}

// HasForecastedAverageMonthlyIncome returns a boolean if a field has been set.
func (o *CraBankIncomeSource) HasForecastedAverageMonthlyIncome() bool {
	if o != nil && o.ForecastedAverageMonthlyIncome.IsSet() {
		return true
	}

	return false
}

// SetForecastedAverageMonthlyIncome gets a reference to the given NullableFloat32 and assigns it to the ForecastedAverageMonthlyIncome field.
func (o *CraBankIncomeSource) SetForecastedAverageMonthlyIncome(v float32) {
	o.ForecastedAverageMonthlyIncome.Set(&v)
}
// SetForecastedAverageMonthlyIncomeNil sets the value for ForecastedAverageMonthlyIncome to be an explicit nil
func (o *CraBankIncomeSource) SetForecastedAverageMonthlyIncomeNil() {
	o.ForecastedAverageMonthlyIncome.Set(nil)
}

// UnsetForecastedAverageMonthlyIncome ensures that no value is present for ForecastedAverageMonthlyIncome, not even an explicit nil
func (o *CraBankIncomeSource) UnsetForecastedAverageMonthlyIncome() {
	o.ForecastedAverageMonthlyIncome.Unset()
}

// GetEmployer returns the Employer field value if set, zero value otherwise.
func (o *CraBankIncomeSource) GetEmployer() CraBankIncomeEmployer {
	if o == nil || o.Employer == nil {
		var ret CraBankIncomeEmployer
		return ret
	}
	return *o.Employer
}

// GetEmployerOk returns a tuple with the Employer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraBankIncomeSource) GetEmployerOk() (*CraBankIncomeEmployer, bool) {
	if o == nil || o.Employer == nil {
		return nil, false
	}
	return o.Employer, true
}

// HasEmployer returns a boolean if a field has been set.
func (o *CraBankIncomeSource) HasEmployer() bool {
	if o != nil && o.Employer != nil {
		return true
	}

	return false
}

// SetEmployer gets a reference to the given CraBankIncomeEmployer and assigns it to the Employer field.
func (o *CraBankIncomeSource) SetEmployer(v CraBankIncomeEmployer) {
	o.Employer = &v
}

// GetHistoricalSummary returns the HistoricalSummary field value if set, zero value otherwise.
func (o *CraBankIncomeSource) GetHistoricalSummary() []CraBankIncomeHistoricalSummary {
	if o == nil || o.HistoricalSummary == nil {
		var ret []CraBankIncomeHistoricalSummary
		return ret
	}
	return *o.HistoricalSummary
}

// GetHistoricalSummaryOk returns a tuple with the HistoricalSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraBankIncomeSource) GetHistoricalSummaryOk() (*[]CraBankIncomeHistoricalSummary, bool) {
	if o == nil || o.HistoricalSummary == nil {
		return nil, false
	}
	return o.HistoricalSummary, true
}

// HasHistoricalSummary returns a boolean if a field has been set.
func (o *CraBankIncomeSource) HasHistoricalSummary() bool {
	if o != nil && o.HistoricalSummary != nil {
		return true
	}

	return false
}

// SetHistoricalSummary gets a reference to the given []CraBankIncomeHistoricalSummary and assigns it to the HistoricalSummary field.
func (o *CraBankIncomeSource) SetHistoricalSummary(v []CraBankIncomeHistoricalSummary) {
	o.HistoricalSummary = &v
}

func (o CraBankIncomeSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IncomeSourceId != nil {
		toSerialize["income_source_id"] = o.IncomeSourceId
	}
	if o.IncomeDescription != nil {
		toSerialize["income_description"] = o.IncomeDescription
	}
	if o.IncomeCategory != nil {
		toSerialize["income_category"] = o.IncomeCategory
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["end_date"] = o.EndDate
	}
	if o.PayFrequency != nil {
		toSerialize["pay_frequency"] = o.PayFrequency
	}
	if o.TotalAmount != nil {
		toSerialize["total_amount"] = o.TotalAmount
	}
	if o.IsoCurrencyCode.IsSet() {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode.Get()
	}
	if o.UnofficialCurrencyCode.IsSet() {
		toSerialize["unofficial_currency_code"] = o.UnofficialCurrencyCode.Get()
	}
	if o.TransactionCount != nil {
		toSerialize["transaction_count"] = o.TransactionCount
	}
	if o.NextPaymentDate.IsSet() {
		toSerialize["next_payment_date"] = o.NextPaymentDate.Get()
	}
	if o.HistoricalAverageMonthlyGrossIncome.IsSet() {
		toSerialize["historical_average_monthly_gross_income"] = o.HistoricalAverageMonthlyGrossIncome.Get()
	}
	if o.HistoricalAverageMonthlyIncome.IsSet() {
		toSerialize["historical_average_monthly_income"] = o.HistoricalAverageMonthlyIncome.Get()
	}
	if o.ForecastedAverageMonthlyIncome.IsSet() {
		toSerialize["forecasted_average_monthly_income"] = o.ForecastedAverageMonthlyIncome.Get()
	}
	if o.Employer != nil {
		toSerialize["employer"] = o.Employer
	}
	if o.HistoricalSummary != nil {
		toSerialize["historical_summary"] = o.HistoricalSummary
	}
	return json.Marshal(toSerialize)
}

type NullableCraBankIncomeSource struct {
	value *CraBankIncomeSource
	isSet bool
}

func (v NullableCraBankIncomeSource) Get() *CraBankIncomeSource {
	return v.value
}

func (v *NullableCraBankIncomeSource) Set(val *CraBankIncomeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableCraBankIncomeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableCraBankIncomeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCraBankIncomeSource(val *CraBankIncomeSource) *NullableCraBankIncomeSource {
	return &NullableCraBankIncomeSource{value: val, isSet: true}
}

func (v NullableCraBankIncomeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCraBankIncomeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


