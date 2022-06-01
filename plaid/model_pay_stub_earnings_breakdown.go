/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.121.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// PayStubEarningsBreakdown An object representing the earnings line items for the pay period.
type PayStubEarningsBreakdown struct {
	// Commonly used term to describe the earning line item.
	CanonicalDescription NullableString `json:"canonical_description"`
	// Raw amount of the earning line item.
	CurrentAmount NullableFloat64 `json:"current_amount"`
	// Description of the earning line item.
	Description NullableString `json:"description"`
	// Number of hours applicable for this earning.
	Hours NullableFloat32 `json:"hours"`
	// The ISO-4217 currency code of the line item. Always `null` if `unofficial_currency_code` is non-null.
	IsoCurrencyCode NullableString `json:"iso_currency_code"`
	// Hourly rate applicable for this earning.
	Rate NullableFloat64 `json:"rate"`
	// The unofficial currency code associated with the line item. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.
	UnofficialCurrencyCode NullableString `json:"unofficial_currency_code"`
	// The year-to-date amount of the deduction.
	YtdAmount NullableFloat64 `json:"ytd_amount"`
	AdditionalProperties map[string]interface{}
}

type _PayStubEarningsBreakdown PayStubEarningsBreakdown

// NewPayStubEarningsBreakdown instantiates a new PayStubEarningsBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayStubEarningsBreakdown(canonicalDescription NullableString, currentAmount NullableFloat64, description NullableString, hours NullableFloat32, isoCurrencyCode NullableString, rate NullableFloat64, unofficialCurrencyCode NullableString, ytdAmount NullableFloat64) *PayStubEarningsBreakdown {
	this := PayStubEarningsBreakdown{}
	this.CanonicalDescription = canonicalDescription
	this.CurrentAmount = currentAmount
	this.Description = description
	this.Hours = hours
	this.IsoCurrencyCode = isoCurrencyCode
	this.Rate = rate
	this.UnofficialCurrencyCode = unofficialCurrencyCode
	this.YtdAmount = ytdAmount
	return &this
}

// NewPayStubEarningsBreakdownWithDefaults instantiates a new PayStubEarningsBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayStubEarningsBreakdownWithDefaults() *PayStubEarningsBreakdown {
	this := PayStubEarningsBreakdown{}
	return &this
}

// GetCanonicalDescription returns the CanonicalDescription field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayStubEarningsBreakdown) GetCanonicalDescription() string {
	if o == nil || o.CanonicalDescription.Get() == nil {
		var ret string
		return ret
	}

	return *o.CanonicalDescription.Get()
}

// GetCanonicalDescriptionOk returns a tuple with the CanonicalDescription field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubEarningsBreakdown) GetCanonicalDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CanonicalDescription.Get(), o.CanonicalDescription.IsSet()
}

// SetCanonicalDescription sets field value
func (o *PayStubEarningsBreakdown) SetCanonicalDescription(v string) {
	o.CanonicalDescription.Set(&v)
}

// GetCurrentAmount returns the CurrentAmount field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *PayStubEarningsBreakdown) GetCurrentAmount() float64 {
	if o == nil || o.CurrentAmount.Get() == nil {
		var ret float64
		return ret
	}

	return *o.CurrentAmount.Get()
}

// GetCurrentAmountOk returns a tuple with the CurrentAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubEarningsBreakdown) GetCurrentAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CurrentAmount.Get(), o.CurrentAmount.IsSet()
}

// SetCurrentAmount sets field value
func (o *PayStubEarningsBreakdown) SetCurrentAmount(v float64) {
	o.CurrentAmount.Set(&v)
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayStubEarningsBreakdown) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubEarningsBreakdown) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *PayStubEarningsBreakdown) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetHours returns the Hours field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *PayStubEarningsBreakdown) GetHours() float32 {
	if o == nil || o.Hours.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Hours.Get()
}

// GetHoursOk returns a tuple with the Hours field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubEarningsBreakdown) GetHoursOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Hours.Get(), o.Hours.IsSet()
}

// SetHours sets field value
func (o *PayStubEarningsBreakdown) SetHours(v float32) {
	o.Hours.Set(&v)
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayStubEarningsBreakdown) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.IsoCurrencyCode.Get()
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubEarningsBreakdown) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsoCurrencyCode.Get(), o.IsoCurrencyCode.IsSet()
}

// SetIsoCurrencyCode sets field value
func (o *PayStubEarningsBreakdown) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode.Set(&v)
}

// GetRate returns the Rate field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *PayStubEarningsBreakdown) GetRate() float64 {
	if o == nil || o.Rate.Get() == nil {
		var ret float64
		return ret
	}

	return *o.Rate.Get()
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubEarningsBreakdown) GetRateOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Rate.Get(), o.Rate.IsSet()
}

// SetRate sets field value
func (o *PayStubEarningsBreakdown) SetRate(v float64) {
	o.Rate.Set(&v)
}

// GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayStubEarningsBreakdown) GetUnofficialCurrencyCode() string {
	if o == nil || o.UnofficialCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.UnofficialCurrencyCode.Get()
}

// GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubEarningsBreakdown) GetUnofficialCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnofficialCurrencyCode.Get(), o.UnofficialCurrencyCode.IsSet()
}

// SetUnofficialCurrencyCode sets field value
func (o *PayStubEarningsBreakdown) SetUnofficialCurrencyCode(v string) {
	o.UnofficialCurrencyCode.Set(&v)
}

// GetYtdAmount returns the YtdAmount field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *PayStubEarningsBreakdown) GetYtdAmount() float64 {
	if o == nil || o.YtdAmount.Get() == nil {
		var ret float64
		return ret
	}

	return *o.YtdAmount.Get()
}

// GetYtdAmountOk returns a tuple with the YtdAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubEarningsBreakdown) GetYtdAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.YtdAmount.Get(), o.YtdAmount.IsSet()
}

// SetYtdAmount sets field value
func (o *PayStubEarningsBreakdown) SetYtdAmount(v float64) {
	o.YtdAmount.Set(&v)
}

func (o PayStubEarningsBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["canonical_description"] = o.CanonicalDescription.Get()
	}
	if true {
		toSerialize["current_amount"] = o.CurrentAmount.Get()
	}
	if true {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["hours"] = o.Hours.Get()
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode.Get()
	}
	if true {
		toSerialize["rate"] = o.Rate.Get()
	}
	if true {
		toSerialize["unofficial_currency_code"] = o.UnofficialCurrencyCode.Get()
	}
	if true {
		toSerialize["ytd_amount"] = o.YtdAmount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PayStubEarningsBreakdown) UnmarshalJSON(bytes []byte) (err error) {
	varPayStubEarningsBreakdown := _PayStubEarningsBreakdown{}

	if err = json.Unmarshal(bytes, &varPayStubEarningsBreakdown); err == nil {
		*o = PayStubEarningsBreakdown(varPayStubEarningsBreakdown)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "canonical_description")
		delete(additionalProperties, "current_amount")
		delete(additionalProperties, "description")
		delete(additionalProperties, "hours")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "rate")
		delete(additionalProperties, "unofficial_currency_code")
		delete(additionalProperties, "ytd_amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePayStubEarningsBreakdown struct {
	value *PayStubEarningsBreakdown
	isSet bool
}

func (v NullablePayStubEarningsBreakdown) Get() *PayStubEarningsBreakdown {
	return v.value
}

func (v *NullablePayStubEarningsBreakdown) Set(val *PayStubEarningsBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullablePayStubEarningsBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullablePayStubEarningsBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayStubEarningsBreakdown(val *PayStubEarningsBreakdown) *NullablePayStubEarningsBreakdown {
	return &NullablePayStubEarningsBreakdown{value: val, isSet: true}
}

func (v NullablePayStubEarningsBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayStubEarningsBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


