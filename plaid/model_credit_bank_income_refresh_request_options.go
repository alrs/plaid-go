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

// CreditBankIncomeRefreshRequestOptions An optional object for `/credit/bank_income/refresh` request options.
type CreditBankIncomeRefreshRequestOptions struct {
	// How many days of data to include in the refresh. If not specified, this will default to the days requested in the most recently generated bank income report for the user.
	DaysRequested *int32 `json:"days_requested,omitempty"`
	// The URL where Plaid will send the bank income webhook.
	Webhook *string `json:"webhook,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreditBankIncomeRefreshRequestOptions CreditBankIncomeRefreshRequestOptions

// NewCreditBankIncomeRefreshRequestOptions instantiates a new CreditBankIncomeRefreshRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankIncomeRefreshRequestOptions() *CreditBankIncomeRefreshRequestOptions {
	this := CreditBankIncomeRefreshRequestOptions{}
	return &this
}

// NewCreditBankIncomeRefreshRequestOptionsWithDefaults instantiates a new CreditBankIncomeRefreshRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankIncomeRefreshRequestOptionsWithDefaults() *CreditBankIncomeRefreshRequestOptions {
	this := CreditBankIncomeRefreshRequestOptions{}
	return &this
}

// GetDaysRequested returns the DaysRequested field value if set, zero value otherwise.
func (o *CreditBankIncomeRefreshRequestOptions) GetDaysRequested() int32 {
	if o == nil || o.DaysRequested == nil {
		var ret int32
		return ret
	}
	return *o.DaysRequested
}

// GetDaysRequestedOk returns a tuple with the DaysRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeRefreshRequestOptions) GetDaysRequestedOk() (*int32, bool) {
	if o == nil || o.DaysRequested == nil {
		return nil, false
	}
	return o.DaysRequested, true
}

// HasDaysRequested returns a boolean if a field has been set.
func (o *CreditBankIncomeRefreshRequestOptions) HasDaysRequested() bool {
	if o != nil && o.DaysRequested != nil {
		return true
	}

	return false
}

// SetDaysRequested gets a reference to the given int32 and assigns it to the DaysRequested field.
func (o *CreditBankIncomeRefreshRequestOptions) SetDaysRequested(v int32) {
	o.DaysRequested = &v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *CreditBankIncomeRefreshRequestOptions) GetWebhook() string {
	if o == nil || o.Webhook == nil {
		var ret string
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeRefreshRequestOptions) GetWebhookOk() (*string, bool) {
	if o == nil || o.Webhook == nil {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *CreditBankIncomeRefreshRequestOptions) HasWebhook() bool {
	if o != nil && o.Webhook != nil {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given string and assigns it to the Webhook field.
func (o *CreditBankIncomeRefreshRequestOptions) SetWebhook(v string) {
	o.Webhook = &v
}

func (o CreditBankIncomeRefreshRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DaysRequested != nil {
		toSerialize["days_requested"] = o.DaysRequested
	}
	if o.Webhook != nil {
		toSerialize["webhook"] = o.Webhook
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditBankIncomeRefreshRequestOptions) UnmarshalJSON(bytes []byte) (err error) {
	varCreditBankIncomeRefreshRequestOptions := _CreditBankIncomeRefreshRequestOptions{}

	if err = json.Unmarshal(bytes, &varCreditBankIncomeRefreshRequestOptions); err == nil {
		*o = CreditBankIncomeRefreshRequestOptions(varCreditBankIncomeRefreshRequestOptions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "days_requested")
		delete(additionalProperties, "webhook")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditBankIncomeRefreshRequestOptions struct {
	value *CreditBankIncomeRefreshRequestOptions
	isSet bool
}

func (v NullableCreditBankIncomeRefreshRequestOptions) Get() *CreditBankIncomeRefreshRequestOptions {
	return v.value
}

func (v *NullableCreditBankIncomeRefreshRequestOptions) Set(val *CreditBankIncomeRefreshRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankIncomeRefreshRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankIncomeRefreshRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankIncomeRefreshRequestOptions(val *CreditBankIncomeRefreshRequestOptions) *NullableCreditBankIncomeRefreshRequestOptions {
	return &NullableCreditBankIncomeRefreshRequestOptions{value: val, isSet: true}
}

func (v NullableCreditBankIncomeRefreshRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankIncomeRefreshRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


