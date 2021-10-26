/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.40.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// EarningsTotal An object representing both the current pay period and year to date amount for an earning category.
type EarningsTotal struct {
	// Commonly used term to describe the line item.
	CanonicalDescription NullableString `json:"canonical_description,omitempty"`
	// Total amount of the earnings for this pay period
	CurrentAmount NullableFloat32 `json:"current_amount,omitempty"`
	CurrentPay *Pay `json:"current_pay,omitempty"`
	YtdPay *Pay `json:"ytd_pay,omitempty"`
	// Total number of hours worked for this pay period
	Hours NullableFloat32 `json:"hours,omitempty"`
	// The ISO-4217 currency code of the line item. Always `null` if `unofficial_currency_code` is non-null.
	IsoCurrencyCode NullableString `json:"iso_currency_code,omitempty"`
	// The unofficial currency code associated with the security. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.
	UnofficialCurrencyCode NullableString `json:"unofficial_currency_code,omitempty"`
	// The total year-to-date amount of the earnings
	YtdAmount NullableFloat32 `json:"ytd_amount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EarningsTotal EarningsTotal

// NewEarningsTotal instantiates a new EarningsTotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEarningsTotal() *EarningsTotal {
	this := EarningsTotal{}
	return &this
}

// NewEarningsTotalWithDefaults instantiates a new EarningsTotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEarningsTotalWithDefaults() *EarningsTotal {
	this := EarningsTotal{}
	return &this
}

// GetCanonicalDescription returns the CanonicalDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EarningsTotal) GetCanonicalDescription() string {
	if o == nil || o.CanonicalDescription.Get() == nil {
		var ret string
		return ret
	}
	return *o.CanonicalDescription.Get()
}

// GetCanonicalDescriptionOk returns a tuple with the CanonicalDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EarningsTotal) GetCanonicalDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CanonicalDescription.Get(), o.CanonicalDescription.IsSet()
}

// HasCanonicalDescription returns a boolean if a field has been set.
func (o *EarningsTotal) HasCanonicalDescription() bool {
	if o != nil && o.CanonicalDescription.IsSet() {
		return true
	}

	return false
}

// SetCanonicalDescription gets a reference to the given NullableString and assigns it to the CanonicalDescription field.
func (o *EarningsTotal) SetCanonicalDescription(v string) {
	o.CanonicalDescription.Set(&v)
}
// SetCanonicalDescriptionNil sets the value for CanonicalDescription to be an explicit nil
func (o *EarningsTotal) SetCanonicalDescriptionNil() {
	o.CanonicalDescription.Set(nil)
}

// UnsetCanonicalDescription ensures that no value is present for CanonicalDescription, not even an explicit nil
func (o *EarningsTotal) UnsetCanonicalDescription() {
	o.CanonicalDescription.Unset()
}

// GetCurrentAmount returns the CurrentAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EarningsTotal) GetCurrentAmount() float32 {
	if o == nil || o.CurrentAmount.Get() == nil {
		var ret float32
		return ret
	}
	return *o.CurrentAmount.Get()
}

// GetCurrentAmountOk returns a tuple with the CurrentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EarningsTotal) GetCurrentAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CurrentAmount.Get(), o.CurrentAmount.IsSet()
}

// HasCurrentAmount returns a boolean if a field has been set.
func (o *EarningsTotal) HasCurrentAmount() bool {
	if o != nil && o.CurrentAmount.IsSet() {
		return true
	}

	return false
}

// SetCurrentAmount gets a reference to the given NullableFloat32 and assigns it to the CurrentAmount field.
func (o *EarningsTotal) SetCurrentAmount(v float32) {
	o.CurrentAmount.Set(&v)
}
// SetCurrentAmountNil sets the value for CurrentAmount to be an explicit nil
func (o *EarningsTotal) SetCurrentAmountNil() {
	o.CurrentAmount.Set(nil)
}

// UnsetCurrentAmount ensures that no value is present for CurrentAmount, not even an explicit nil
func (o *EarningsTotal) UnsetCurrentAmount() {
	o.CurrentAmount.Unset()
}

// GetCurrentPay returns the CurrentPay field value if set, zero value otherwise.
func (o *EarningsTotal) GetCurrentPay() Pay {
	if o == nil || o.CurrentPay == nil {
		var ret Pay
		return ret
	}
	return *o.CurrentPay
}

// GetCurrentPayOk returns a tuple with the CurrentPay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsTotal) GetCurrentPayOk() (*Pay, bool) {
	if o == nil || o.CurrentPay == nil {
		return nil, false
	}
	return o.CurrentPay, true
}

// HasCurrentPay returns a boolean if a field has been set.
func (o *EarningsTotal) HasCurrentPay() bool {
	if o != nil && o.CurrentPay != nil {
		return true
	}

	return false
}

// SetCurrentPay gets a reference to the given Pay and assigns it to the CurrentPay field.
func (o *EarningsTotal) SetCurrentPay(v Pay) {
	o.CurrentPay = &v
}

// GetYtdPay returns the YtdPay field value if set, zero value otherwise.
func (o *EarningsTotal) GetYtdPay() Pay {
	if o == nil || o.YtdPay == nil {
		var ret Pay
		return ret
	}
	return *o.YtdPay
}

// GetYtdPayOk returns a tuple with the YtdPay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsTotal) GetYtdPayOk() (*Pay, bool) {
	if o == nil || o.YtdPay == nil {
		return nil, false
	}
	return o.YtdPay, true
}

// HasYtdPay returns a boolean if a field has been set.
func (o *EarningsTotal) HasYtdPay() bool {
	if o != nil && o.YtdPay != nil {
		return true
	}

	return false
}

// SetYtdPay gets a reference to the given Pay and assigns it to the YtdPay field.
func (o *EarningsTotal) SetYtdPay(v Pay) {
	o.YtdPay = &v
}

// GetHours returns the Hours field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EarningsTotal) GetHours() float32 {
	if o == nil || o.Hours.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Hours.Get()
}

// GetHoursOk returns a tuple with the Hours field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EarningsTotal) GetHoursOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Hours.Get(), o.Hours.IsSet()
}

// HasHours returns a boolean if a field has been set.
func (o *EarningsTotal) HasHours() bool {
	if o != nil && o.Hours.IsSet() {
		return true
	}

	return false
}

// SetHours gets a reference to the given NullableFloat32 and assigns it to the Hours field.
func (o *EarningsTotal) SetHours(v float32) {
	o.Hours.Set(&v)
}
// SetHoursNil sets the value for Hours to be an explicit nil
func (o *EarningsTotal) SetHoursNil() {
	o.Hours.Set(nil)
}

// UnsetHours ensures that no value is present for Hours, not even an explicit nil
func (o *EarningsTotal) UnsetHours() {
	o.Hours.Unset()
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EarningsTotal) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.IsoCurrencyCode.Get()
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EarningsTotal) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsoCurrencyCode.Get(), o.IsoCurrencyCode.IsSet()
}

// HasIsoCurrencyCode returns a boolean if a field has been set.
func (o *EarningsTotal) HasIsoCurrencyCode() bool {
	if o != nil && o.IsoCurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetIsoCurrencyCode gets a reference to the given NullableString and assigns it to the IsoCurrencyCode field.
func (o *EarningsTotal) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode.Set(&v)
}
// SetIsoCurrencyCodeNil sets the value for IsoCurrencyCode to be an explicit nil
func (o *EarningsTotal) SetIsoCurrencyCodeNil() {
	o.IsoCurrencyCode.Set(nil)
}

// UnsetIsoCurrencyCode ensures that no value is present for IsoCurrencyCode, not even an explicit nil
func (o *EarningsTotal) UnsetIsoCurrencyCode() {
	o.IsoCurrencyCode.Unset()
}

// GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EarningsTotal) GetUnofficialCurrencyCode() string {
	if o == nil || o.UnofficialCurrencyCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.UnofficialCurrencyCode.Get()
}

// GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EarningsTotal) GetUnofficialCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnofficialCurrencyCode.Get(), o.UnofficialCurrencyCode.IsSet()
}

// HasUnofficialCurrencyCode returns a boolean if a field has been set.
func (o *EarningsTotal) HasUnofficialCurrencyCode() bool {
	if o != nil && o.UnofficialCurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetUnofficialCurrencyCode gets a reference to the given NullableString and assigns it to the UnofficialCurrencyCode field.
func (o *EarningsTotal) SetUnofficialCurrencyCode(v string) {
	o.UnofficialCurrencyCode.Set(&v)
}
// SetUnofficialCurrencyCodeNil sets the value for UnofficialCurrencyCode to be an explicit nil
func (o *EarningsTotal) SetUnofficialCurrencyCodeNil() {
	o.UnofficialCurrencyCode.Set(nil)
}

// UnsetUnofficialCurrencyCode ensures that no value is present for UnofficialCurrencyCode, not even an explicit nil
func (o *EarningsTotal) UnsetUnofficialCurrencyCode() {
	o.UnofficialCurrencyCode.Unset()
}

// GetYtdAmount returns the YtdAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EarningsTotal) GetYtdAmount() float32 {
	if o == nil || o.YtdAmount.Get() == nil {
		var ret float32
		return ret
	}
	return *o.YtdAmount.Get()
}

// GetYtdAmountOk returns a tuple with the YtdAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EarningsTotal) GetYtdAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.YtdAmount.Get(), o.YtdAmount.IsSet()
}

// HasYtdAmount returns a boolean if a field has been set.
func (o *EarningsTotal) HasYtdAmount() bool {
	if o != nil && o.YtdAmount.IsSet() {
		return true
	}

	return false
}

// SetYtdAmount gets a reference to the given NullableFloat32 and assigns it to the YtdAmount field.
func (o *EarningsTotal) SetYtdAmount(v float32) {
	o.YtdAmount.Set(&v)
}
// SetYtdAmountNil sets the value for YtdAmount to be an explicit nil
func (o *EarningsTotal) SetYtdAmountNil() {
	o.YtdAmount.Set(nil)
}

// UnsetYtdAmount ensures that no value is present for YtdAmount, not even an explicit nil
func (o *EarningsTotal) UnsetYtdAmount() {
	o.YtdAmount.Unset()
}

func (o EarningsTotal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanonicalDescription.IsSet() {
		toSerialize["canonical_description"] = o.CanonicalDescription.Get()
	}
	if o.CurrentAmount.IsSet() {
		toSerialize["current_amount"] = o.CurrentAmount.Get()
	}
	if o.CurrentPay != nil {
		toSerialize["current_pay"] = o.CurrentPay
	}
	if o.YtdPay != nil {
		toSerialize["ytd_pay"] = o.YtdPay
	}
	if o.Hours.IsSet() {
		toSerialize["hours"] = o.Hours.Get()
	}
	if o.IsoCurrencyCode.IsSet() {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode.Get()
	}
	if o.UnofficialCurrencyCode.IsSet() {
		toSerialize["unofficial_currency_code"] = o.UnofficialCurrencyCode.Get()
	}
	if o.YtdAmount.IsSet() {
		toSerialize["ytd_amount"] = o.YtdAmount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EarningsTotal) UnmarshalJSON(bytes []byte) (err error) {
	varEarningsTotal := _EarningsTotal{}

	if err = json.Unmarshal(bytes, &varEarningsTotal); err == nil {
		*o = EarningsTotal(varEarningsTotal)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "canonical_description")
		delete(additionalProperties, "current_amount")
		delete(additionalProperties, "current_pay")
		delete(additionalProperties, "ytd_pay")
		delete(additionalProperties, "hours")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "unofficial_currency_code")
		delete(additionalProperties, "ytd_amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEarningsTotal struct {
	value *EarningsTotal
	isSet bool
}

func (v NullableEarningsTotal) Get() *EarningsTotal {
	return v.value
}

func (v *NullableEarningsTotal) Set(val *EarningsTotal) {
	v.value = val
	v.isSet = true
}

func (v NullableEarningsTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableEarningsTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEarningsTotal(val *EarningsTotal) *NullableEarningsTotal {
	return &NullableEarningsTotal{value: val, isSet: true}
}

func (v NullableEarningsTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEarningsTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


