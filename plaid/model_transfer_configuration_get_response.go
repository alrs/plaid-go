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

// TransferConfigurationGetResponse Defines the response schema for `/transfer/configuration/get`
type TransferConfigurationGetResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	// The max limit of dollar amount of a single transfer (decimal string with two digits of precision e.g. \"10.00\").
	MaxSingleTransferAmount string `json:"max_single_transfer_amount"`
	// The max limit of dollar amount of a single credit transfer (decimal string with two digits of precision e.g. \"10.00\").
	MaxSingleTransferCreditAmount string `json:"max_single_transfer_credit_amount"`
	// The max limit of dollar amount of a single debit transfer (decimal string with two digits of precision e.g. \"10.00\").
	MaxSingleTransferDebitAmount string `json:"max_single_transfer_debit_amount"`
	// The max limit of sum of dollar amount of credit transfers in last 24 hours (decimal string with two digits of precision e.g. \"10.00\").
	MaxDailyCreditAmount string `json:"max_daily_credit_amount"`
	// The max limit of sum of dollar amount of debit transfers in last 24 hours (decimal string with two digits of precision e.g. \"10.00\").
	MaxDailyDebitAmount string `json:"max_daily_debit_amount"`
	// The max limit of sum of dollar amount of credit and debit transfers in one calendar month (decimal string with two digits of precision e.g. \"10.00\").
	MaxMonthlyAmount string `json:"max_monthly_amount"`
	// The max limit of sum of dollar amount of credit transfers in one calendar month (decimal string with two digits of precision e.g. \"10.00\").
	MaxMonthlyCreditAmount string `json:"max_monthly_credit_amount"`
	// The max limit of sum of dollar amount of debit transfers in one calendar month (decimal string with two digits of precision e.g. \"10.00\").
	MaxMonthlyDebitAmount string `json:"max_monthly_debit_amount"`
	// The currency of the dollar amount, e.g. \"USD\".
	IsoCurrencyCode string `json:"iso_currency_code"`
	AdditionalProperties map[string]interface{}
}

type _TransferConfigurationGetResponse TransferConfigurationGetResponse

// NewTransferConfigurationGetResponse instantiates a new TransferConfigurationGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferConfigurationGetResponse(requestId string, maxSingleTransferAmount string, maxSingleTransferCreditAmount string, maxSingleTransferDebitAmount string, maxDailyCreditAmount string, maxDailyDebitAmount string, maxMonthlyAmount string, maxMonthlyCreditAmount string, maxMonthlyDebitAmount string, isoCurrencyCode string) *TransferConfigurationGetResponse {
	this := TransferConfigurationGetResponse{}
	this.RequestId = requestId
	this.MaxSingleTransferAmount = maxSingleTransferAmount
	this.MaxSingleTransferCreditAmount = maxSingleTransferCreditAmount
	this.MaxSingleTransferDebitAmount = maxSingleTransferDebitAmount
	this.MaxDailyCreditAmount = maxDailyCreditAmount
	this.MaxDailyDebitAmount = maxDailyDebitAmount
	this.MaxMonthlyAmount = maxMonthlyAmount
	this.MaxMonthlyCreditAmount = maxMonthlyCreditAmount
	this.MaxMonthlyDebitAmount = maxMonthlyDebitAmount
	this.IsoCurrencyCode = isoCurrencyCode
	return &this
}

// NewTransferConfigurationGetResponseWithDefaults instantiates a new TransferConfigurationGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferConfigurationGetResponseWithDefaults() *TransferConfigurationGetResponse {
	this := TransferConfigurationGetResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *TransferConfigurationGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferConfigurationGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferConfigurationGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

// GetMaxSingleTransferAmount returns the MaxSingleTransferAmount field value
func (o *TransferConfigurationGetResponse) GetMaxSingleTransferAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxSingleTransferAmount
}

// GetMaxSingleTransferAmountOk returns a tuple with the MaxSingleTransferAmount field value
// and a boolean to check if the value has been set.
func (o *TransferConfigurationGetResponse) GetMaxSingleTransferAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MaxSingleTransferAmount, true
}

// SetMaxSingleTransferAmount sets field value
func (o *TransferConfigurationGetResponse) SetMaxSingleTransferAmount(v string) {
	o.MaxSingleTransferAmount = v
}

// GetMaxSingleTransferCreditAmount returns the MaxSingleTransferCreditAmount field value
func (o *TransferConfigurationGetResponse) GetMaxSingleTransferCreditAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxSingleTransferCreditAmount
}

// GetMaxSingleTransferCreditAmountOk returns a tuple with the MaxSingleTransferCreditAmount field value
// and a boolean to check if the value has been set.
func (o *TransferConfigurationGetResponse) GetMaxSingleTransferCreditAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MaxSingleTransferCreditAmount, true
}

// SetMaxSingleTransferCreditAmount sets field value
func (o *TransferConfigurationGetResponse) SetMaxSingleTransferCreditAmount(v string) {
	o.MaxSingleTransferCreditAmount = v
}

// GetMaxSingleTransferDebitAmount returns the MaxSingleTransferDebitAmount field value
func (o *TransferConfigurationGetResponse) GetMaxSingleTransferDebitAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxSingleTransferDebitAmount
}

// GetMaxSingleTransferDebitAmountOk returns a tuple with the MaxSingleTransferDebitAmount field value
// and a boolean to check if the value has been set.
func (o *TransferConfigurationGetResponse) GetMaxSingleTransferDebitAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MaxSingleTransferDebitAmount, true
}

// SetMaxSingleTransferDebitAmount sets field value
func (o *TransferConfigurationGetResponse) SetMaxSingleTransferDebitAmount(v string) {
	o.MaxSingleTransferDebitAmount = v
}

// GetMaxDailyCreditAmount returns the MaxDailyCreditAmount field value
func (o *TransferConfigurationGetResponse) GetMaxDailyCreditAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxDailyCreditAmount
}

// GetMaxDailyCreditAmountOk returns a tuple with the MaxDailyCreditAmount field value
// and a boolean to check if the value has been set.
func (o *TransferConfigurationGetResponse) GetMaxDailyCreditAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MaxDailyCreditAmount, true
}

// SetMaxDailyCreditAmount sets field value
func (o *TransferConfigurationGetResponse) SetMaxDailyCreditAmount(v string) {
	o.MaxDailyCreditAmount = v
}

// GetMaxDailyDebitAmount returns the MaxDailyDebitAmount field value
func (o *TransferConfigurationGetResponse) GetMaxDailyDebitAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxDailyDebitAmount
}

// GetMaxDailyDebitAmountOk returns a tuple with the MaxDailyDebitAmount field value
// and a boolean to check if the value has been set.
func (o *TransferConfigurationGetResponse) GetMaxDailyDebitAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MaxDailyDebitAmount, true
}

// SetMaxDailyDebitAmount sets field value
func (o *TransferConfigurationGetResponse) SetMaxDailyDebitAmount(v string) {
	o.MaxDailyDebitAmount = v
}

// GetMaxMonthlyAmount returns the MaxMonthlyAmount field value
func (o *TransferConfigurationGetResponse) GetMaxMonthlyAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxMonthlyAmount
}

// GetMaxMonthlyAmountOk returns a tuple with the MaxMonthlyAmount field value
// and a boolean to check if the value has been set.
func (o *TransferConfigurationGetResponse) GetMaxMonthlyAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MaxMonthlyAmount, true
}

// SetMaxMonthlyAmount sets field value
func (o *TransferConfigurationGetResponse) SetMaxMonthlyAmount(v string) {
	o.MaxMonthlyAmount = v
}

// GetMaxMonthlyCreditAmount returns the MaxMonthlyCreditAmount field value
func (o *TransferConfigurationGetResponse) GetMaxMonthlyCreditAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxMonthlyCreditAmount
}

// GetMaxMonthlyCreditAmountOk returns a tuple with the MaxMonthlyCreditAmount field value
// and a boolean to check if the value has been set.
func (o *TransferConfigurationGetResponse) GetMaxMonthlyCreditAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MaxMonthlyCreditAmount, true
}

// SetMaxMonthlyCreditAmount sets field value
func (o *TransferConfigurationGetResponse) SetMaxMonthlyCreditAmount(v string) {
	o.MaxMonthlyCreditAmount = v
}

// GetMaxMonthlyDebitAmount returns the MaxMonthlyDebitAmount field value
func (o *TransferConfigurationGetResponse) GetMaxMonthlyDebitAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxMonthlyDebitAmount
}

// GetMaxMonthlyDebitAmountOk returns a tuple with the MaxMonthlyDebitAmount field value
// and a boolean to check if the value has been set.
func (o *TransferConfigurationGetResponse) GetMaxMonthlyDebitAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MaxMonthlyDebitAmount, true
}

// SetMaxMonthlyDebitAmount sets field value
func (o *TransferConfigurationGetResponse) SetMaxMonthlyDebitAmount(v string) {
	o.MaxMonthlyDebitAmount = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
func (o *TransferConfigurationGetResponse) GetIsoCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *TransferConfigurationGetResponse) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsoCurrencyCode, true
}

// SetIsoCurrencyCode sets field value
func (o *TransferConfigurationGetResponse) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = v
}

func (o TransferConfigurationGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}
	if true {
		toSerialize["max_single_transfer_amount"] = o.MaxSingleTransferAmount
	}
	if true {
		toSerialize["max_single_transfer_credit_amount"] = o.MaxSingleTransferCreditAmount
	}
	if true {
		toSerialize["max_single_transfer_debit_amount"] = o.MaxSingleTransferDebitAmount
	}
	if true {
		toSerialize["max_daily_credit_amount"] = o.MaxDailyCreditAmount
	}
	if true {
		toSerialize["max_daily_debit_amount"] = o.MaxDailyDebitAmount
	}
	if true {
		toSerialize["max_monthly_amount"] = o.MaxMonthlyAmount
	}
	if true {
		toSerialize["max_monthly_credit_amount"] = o.MaxMonthlyCreditAmount
	}
	if true {
		toSerialize["max_monthly_debit_amount"] = o.MaxMonthlyDebitAmount
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferConfigurationGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferConfigurationGetResponse := _TransferConfigurationGetResponse{}

	if err = json.Unmarshal(bytes, &varTransferConfigurationGetResponse); err == nil {
		*o = TransferConfigurationGetResponse(varTransferConfigurationGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		delete(additionalProperties, "max_single_transfer_amount")
		delete(additionalProperties, "max_single_transfer_credit_amount")
		delete(additionalProperties, "max_single_transfer_debit_amount")
		delete(additionalProperties, "max_daily_credit_amount")
		delete(additionalProperties, "max_daily_debit_amount")
		delete(additionalProperties, "max_monthly_amount")
		delete(additionalProperties, "max_monthly_credit_amount")
		delete(additionalProperties, "max_monthly_debit_amount")
		delete(additionalProperties, "iso_currency_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferConfigurationGetResponse struct {
	value *TransferConfigurationGetResponse
	isSet bool
}

func (v NullableTransferConfigurationGetResponse) Get() *TransferConfigurationGetResponse {
	return v.value
}

func (v *NullableTransferConfigurationGetResponse) Set(val *TransferConfigurationGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferConfigurationGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferConfigurationGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferConfigurationGetResponse(val *TransferConfigurationGetResponse) *NullableTransferConfigurationGetResponse {
	return &NullableTransferConfigurationGetResponse{value: val, isSet: true}
}

func (v NullableTransferConfigurationGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferConfigurationGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


