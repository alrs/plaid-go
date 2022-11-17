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

// PaystubDetails An object representing details that can be found on the paystub.
type PaystubDetails struct {
	// Beginning date of the pay period on the paystub in the 'YYYY-MM-DD' format.
	PayPeriodStartDate NullableString `json:"pay_period_start_date,omitempty"`
	// Ending date of the pay period on the paystub in the 'YYYY-MM-DD' format.
	PayPeriodEndDate NullableString `json:"pay_period_end_date,omitempty"`
	// Pay date on the paystub in the 'YYYY-MM-DD' format.
	PayDate NullableString `json:"pay_date,omitempty"`
	// The name of the payroll provider that generated the paystub, e.g. ADP
	PaystubProvider NullableString `json:"paystub_provider,omitempty"`
	PayFrequency NullablePaystubPayFrequency `json:"pay_frequency,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaystubDetails PaystubDetails

// NewPaystubDetails instantiates a new PaystubDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaystubDetails() *PaystubDetails {
	this := PaystubDetails{}
	return &this
}

// NewPaystubDetailsWithDefaults instantiates a new PaystubDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaystubDetailsWithDefaults() *PaystubDetails {
	this := PaystubDetails{}
	return &this
}

// GetPayPeriodStartDate returns the PayPeriodStartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaystubDetails) GetPayPeriodStartDate() string {
	if o == nil || o.PayPeriodStartDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.PayPeriodStartDate.Get()
}

// GetPayPeriodStartDateOk returns a tuple with the PayPeriodStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaystubDetails) GetPayPeriodStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayPeriodStartDate.Get(), o.PayPeriodStartDate.IsSet()
}

// HasPayPeriodStartDate returns a boolean if a field has been set.
func (o *PaystubDetails) HasPayPeriodStartDate() bool {
	if o != nil && o.PayPeriodStartDate.IsSet() {
		return true
	}

	return false
}

// SetPayPeriodStartDate gets a reference to the given NullableString and assigns it to the PayPeriodStartDate field.
func (o *PaystubDetails) SetPayPeriodStartDate(v string) {
	o.PayPeriodStartDate.Set(&v)
}
// SetPayPeriodStartDateNil sets the value for PayPeriodStartDate to be an explicit nil
func (o *PaystubDetails) SetPayPeriodStartDateNil() {
	o.PayPeriodStartDate.Set(nil)
}

// UnsetPayPeriodStartDate ensures that no value is present for PayPeriodStartDate, not even an explicit nil
func (o *PaystubDetails) UnsetPayPeriodStartDate() {
	o.PayPeriodStartDate.Unset()
}

// GetPayPeriodEndDate returns the PayPeriodEndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaystubDetails) GetPayPeriodEndDate() string {
	if o == nil || o.PayPeriodEndDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.PayPeriodEndDate.Get()
}

// GetPayPeriodEndDateOk returns a tuple with the PayPeriodEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaystubDetails) GetPayPeriodEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayPeriodEndDate.Get(), o.PayPeriodEndDate.IsSet()
}

// HasPayPeriodEndDate returns a boolean if a field has been set.
func (o *PaystubDetails) HasPayPeriodEndDate() bool {
	if o != nil && o.PayPeriodEndDate.IsSet() {
		return true
	}

	return false
}

// SetPayPeriodEndDate gets a reference to the given NullableString and assigns it to the PayPeriodEndDate field.
func (o *PaystubDetails) SetPayPeriodEndDate(v string) {
	o.PayPeriodEndDate.Set(&v)
}
// SetPayPeriodEndDateNil sets the value for PayPeriodEndDate to be an explicit nil
func (o *PaystubDetails) SetPayPeriodEndDateNil() {
	o.PayPeriodEndDate.Set(nil)
}

// UnsetPayPeriodEndDate ensures that no value is present for PayPeriodEndDate, not even an explicit nil
func (o *PaystubDetails) UnsetPayPeriodEndDate() {
	o.PayPeriodEndDate.Unset()
}

// GetPayDate returns the PayDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaystubDetails) GetPayDate() string {
	if o == nil || o.PayDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.PayDate.Get()
}

// GetPayDateOk returns a tuple with the PayDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaystubDetails) GetPayDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayDate.Get(), o.PayDate.IsSet()
}

// HasPayDate returns a boolean if a field has been set.
func (o *PaystubDetails) HasPayDate() bool {
	if o != nil && o.PayDate.IsSet() {
		return true
	}

	return false
}

// SetPayDate gets a reference to the given NullableString and assigns it to the PayDate field.
func (o *PaystubDetails) SetPayDate(v string) {
	o.PayDate.Set(&v)
}
// SetPayDateNil sets the value for PayDate to be an explicit nil
func (o *PaystubDetails) SetPayDateNil() {
	o.PayDate.Set(nil)
}

// UnsetPayDate ensures that no value is present for PayDate, not even an explicit nil
func (o *PaystubDetails) UnsetPayDate() {
	o.PayDate.Unset()
}

// GetPaystubProvider returns the PaystubProvider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaystubDetails) GetPaystubProvider() string {
	if o == nil || o.PaystubProvider.Get() == nil {
		var ret string
		return ret
	}
	return *o.PaystubProvider.Get()
}

// GetPaystubProviderOk returns a tuple with the PaystubProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaystubDetails) GetPaystubProviderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaystubProvider.Get(), o.PaystubProvider.IsSet()
}

// HasPaystubProvider returns a boolean if a field has been set.
func (o *PaystubDetails) HasPaystubProvider() bool {
	if o != nil && o.PaystubProvider.IsSet() {
		return true
	}

	return false
}

// SetPaystubProvider gets a reference to the given NullableString and assigns it to the PaystubProvider field.
func (o *PaystubDetails) SetPaystubProvider(v string) {
	o.PaystubProvider.Set(&v)
}
// SetPaystubProviderNil sets the value for PaystubProvider to be an explicit nil
func (o *PaystubDetails) SetPaystubProviderNil() {
	o.PaystubProvider.Set(nil)
}

// UnsetPaystubProvider ensures that no value is present for PaystubProvider, not even an explicit nil
func (o *PaystubDetails) UnsetPaystubProvider() {
	o.PaystubProvider.Unset()
}

// GetPayFrequency returns the PayFrequency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaystubDetails) GetPayFrequency() PaystubPayFrequency {
	if o == nil || o.PayFrequency.Get() == nil {
		var ret PaystubPayFrequency
		return ret
	}
	return *o.PayFrequency.Get()
}

// GetPayFrequencyOk returns a tuple with the PayFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaystubDetails) GetPayFrequencyOk() (*PaystubPayFrequency, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayFrequency.Get(), o.PayFrequency.IsSet()
}

// HasPayFrequency returns a boolean if a field has been set.
func (o *PaystubDetails) HasPayFrequency() bool {
	if o != nil && o.PayFrequency.IsSet() {
		return true
	}

	return false
}

// SetPayFrequency gets a reference to the given NullablePaystubPayFrequency and assigns it to the PayFrequency field.
func (o *PaystubDetails) SetPayFrequency(v PaystubPayFrequency) {
	o.PayFrequency.Set(&v)
}
// SetPayFrequencyNil sets the value for PayFrequency to be an explicit nil
func (o *PaystubDetails) SetPayFrequencyNil() {
	o.PayFrequency.Set(nil)
}

// UnsetPayFrequency ensures that no value is present for PayFrequency, not even an explicit nil
func (o *PaystubDetails) UnsetPayFrequency() {
	o.PayFrequency.Unset()
}

func (o PaystubDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PayPeriodStartDate.IsSet() {
		toSerialize["pay_period_start_date"] = o.PayPeriodStartDate.Get()
	}
	if o.PayPeriodEndDate.IsSet() {
		toSerialize["pay_period_end_date"] = o.PayPeriodEndDate.Get()
	}
	if o.PayDate.IsSet() {
		toSerialize["pay_date"] = o.PayDate.Get()
	}
	if o.PaystubProvider.IsSet() {
		toSerialize["paystub_provider"] = o.PaystubProvider.Get()
	}
	if o.PayFrequency.IsSet() {
		toSerialize["pay_frequency"] = o.PayFrequency.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaystubDetails) UnmarshalJSON(bytes []byte) (err error) {
	varPaystubDetails := _PaystubDetails{}

	if err = json.Unmarshal(bytes, &varPaystubDetails); err == nil {
		*o = PaystubDetails(varPaystubDetails)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "pay_period_start_date")
		delete(additionalProperties, "pay_period_end_date")
		delete(additionalProperties, "pay_date")
		delete(additionalProperties, "paystub_provider")
		delete(additionalProperties, "pay_frequency")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaystubDetails struct {
	value *PaystubDetails
	isSet bool
}

func (v NullablePaystubDetails) Get() *PaystubDetails {
	return v.value
}

func (v *NullablePaystubDetails) Set(val *PaystubDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePaystubDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePaystubDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaystubDetails(val *PaystubDetails) *NullablePaystubDetails {
	return &NullablePaystubDetails{value: val, isSet: true}
}

func (v NullablePaystubDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaystubDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


