/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.379.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// NetPay An object representing information about the net pay amount on the paystub.
type NetPay struct {
	// Raw amount of the net pay for the pay period
	CurrentAmount NullableFloat64 `json:"current_amount,omitempty"`
	// Description of the net pay
	Description NullableString `json:"description,omitempty"`
	// The ISO-4217 currency code of the net pay. Always `null` if `unofficial_currency_code` is non-null.
	IsoCurrencyCode NullableString `json:"iso_currency_code,omitempty"`
	// The unofficial currency code associated with the net pay. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.
	UnofficialCurrencyCode NullableString `json:"unofficial_currency_code,omitempty"`
	// The year-to-date amount of the net pay
	YtdAmount NullableFloat64 `json:"ytd_amount,omitempty"`
	Total *Total `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetPay NetPay

// NewNetPay instantiates a new NetPay object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetPay() *NetPay {
	this := NetPay{}
	return &this
}

// NewNetPayWithDefaults instantiates a new NetPay object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetPayWithDefaults() *NetPay {
	this := NetPay{}
	return &this
}

// GetCurrentAmount returns the CurrentAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetPay) GetCurrentAmount() float64 {
	if o == nil || o.CurrentAmount.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CurrentAmount.Get()
}

// GetCurrentAmountOk returns a tuple with the CurrentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetPay) GetCurrentAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CurrentAmount.Get(), o.CurrentAmount.IsSet()
}

// HasCurrentAmount returns a boolean if a field has been set.
func (o *NetPay) HasCurrentAmount() bool {
	if o != nil && o.CurrentAmount.IsSet() {
		return true
	}

	return false
}

// SetCurrentAmount gets a reference to the given NullableFloat64 and assigns it to the CurrentAmount field.
func (o *NetPay) SetCurrentAmount(v float64) {
	o.CurrentAmount.Set(&v)
}
// SetCurrentAmountNil sets the value for CurrentAmount to be an explicit nil
func (o *NetPay) SetCurrentAmountNil() {
	o.CurrentAmount.Set(nil)
}

// UnsetCurrentAmount ensures that no value is present for CurrentAmount, not even an explicit nil
func (o *NetPay) UnsetCurrentAmount() {
	o.CurrentAmount.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetPay) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetPay) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *NetPay) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *NetPay) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *NetPay) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *NetPay) UnsetDescription() {
	o.Description.Unset()
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetPay) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.IsoCurrencyCode.Get()
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetPay) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsoCurrencyCode.Get(), o.IsoCurrencyCode.IsSet()
}

// HasIsoCurrencyCode returns a boolean if a field has been set.
func (o *NetPay) HasIsoCurrencyCode() bool {
	if o != nil && o.IsoCurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetIsoCurrencyCode gets a reference to the given NullableString and assigns it to the IsoCurrencyCode field.
func (o *NetPay) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode.Set(&v)
}
// SetIsoCurrencyCodeNil sets the value for IsoCurrencyCode to be an explicit nil
func (o *NetPay) SetIsoCurrencyCodeNil() {
	o.IsoCurrencyCode.Set(nil)
}

// UnsetIsoCurrencyCode ensures that no value is present for IsoCurrencyCode, not even an explicit nil
func (o *NetPay) UnsetIsoCurrencyCode() {
	o.IsoCurrencyCode.Unset()
}

// GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetPay) GetUnofficialCurrencyCode() string {
	if o == nil || o.UnofficialCurrencyCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.UnofficialCurrencyCode.Get()
}

// GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetPay) GetUnofficialCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnofficialCurrencyCode.Get(), o.UnofficialCurrencyCode.IsSet()
}

// HasUnofficialCurrencyCode returns a boolean if a field has been set.
func (o *NetPay) HasUnofficialCurrencyCode() bool {
	if o != nil && o.UnofficialCurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetUnofficialCurrencyCode gets a reference to the given NullableString and assigns it to the UnofficialCurrencyCode field.
func (o *NetPay) SetUnofficialCurrencyCode(v string) {
	o.UnofficialCurrencyCode.Set(&v)
}
// SetUnofficialCurrencyCodeNil sets the value for UnofficialCurrencyCode to be an explicit nil
func (o *NetPay) SetUnofficialCurrencyCodeNil() {
	o.UnofficialCurrencyCode.Set(nil)
}

// UnsetUnofficialCurrencyCode ensures that no value is present for UnofficialCurrencyCode, not even an explicit nil
func (o *NetPay) UnsetUnofficialCurrencyCode() {
	o.UnofficialCurrencyCode.Unset()
}

// GetYtdAmount returns the YtdAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetPay) GetYtdAmount() float64 {
	if o == nil || o.YtdAmount.Get() == nil {
		var ret float64
		return ret
	}
	return *o.YtdAmount.Get()
}

// GetYtdAmountOk returns a tuple with the YtdAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetPay) GetYtdAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.YtdAmount.Get(), o.YtdAmount.IsSet()
}

// HasYtdAmount returns a boolean if a field has been set.
func (o *NetPay) HasYtdAmount() bool {
	if o != nil && o.YtdAmount.IsSet() {
		return true
	}

	return false
}

// SetYtdAmount gets a reference to the given NullableFloat64 and assigns it to the YtdAmount field.
func (o *NetPay) SetYtdAmount(v float64) {
	o.YtdAmount.Set(&v)
}
// SetYtdAmountNil sets the value for YtdAmount to be an explicit nil
func (o *NetPay) SetYtdAmountNil() {
	o.YtdAmount.Set(nil)
}

// UnsetYtdAmount ensures that no value is present for YtdAmount, not even an explicit nil
func (o *NetPay) UnsetYtdAmount() {
	o.YtdAmount.Unset()
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *NetPay) GetTotal() Total {
	if o == nil || o.Total == nil {
		var ret Total
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetPay) GetTotalOk() (*Total, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *NetPay) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given Total and assigns it to the Total field.
func (o *NetPay) SetTotal(v Total) {
	o.Total = &v
}

func (o NetPay) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentAmount.IsSet() {
		toSerialize["current_amount"] = o.CurrentAmount.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
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
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NetPay) UnmarshalJSON(bytes []byte) (err error) {
	varNetPay := _NetPay{}

	if err = json.Unmarshal(bytes, &varNetPay); err == nil {
		*o = NetPay(varNetPay)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "current_amount")
		delete(additionalProperties, "description")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "unofficial_currency_code")
		delete(additionalProperties, "ytd_amount")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetPay struct {
	value *NetPay
	isSet bool
}

func (v NullableNetPay) Get() *NetPay {
	return v.value
}

func (v *NullableNetPay) Set(val *NetPay) {
	v.value = val
	v.isSet = true
}

func (v NullableNetPay) IsSet() bool {
	return v.isSet
}

func (v *NullableNetPay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetPay(val *NetPay) *NullableNetPay {
	return &NullableNetPay{value: val, isSet: true}
}

func (v NullableNetPay) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetPay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


