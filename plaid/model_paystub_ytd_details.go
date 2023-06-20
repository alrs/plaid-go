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

// PaystubYTDDetails The amount of income earned year to date, as based on paystub data.
type PaystubYTDDetails struct {
	// Year-to-date gross earnings.
	GrossEarnings NullableFloat64 `json:"gross_earnings,omitempty"`
	// Year-to-date net (take home) earnings.
	NetEarnings NullableFloat64 `json:"net_earnings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaystubYTDDetails PaystubYTDDetails

// NewPaystubYTDDetails instantiates a new PaystubYTDDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaystubYTDDetails() *PaystubYTDDetails {
	this := PaystubYTDDetails{}
	return &this
}

// NewPaystubYTDDetailsWithDefaults instantiates a new PaystubYTDDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaystubYTDDetailsWithDefaults() *PaystubYTDDetails {
	this := PaystubYTDDetails{}
	return &this
}

// GetGrossEarnings returns the GrossEarnings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaystubYTDDetails) GetGrossEarnings() float64 {
	if o == nil || o.GrossEarnings.Get() == nil {
		var ret float64
		return ret
	}
	return *o.GrossEarnings.Get()
}

// GetGrossEarningsOk returns a tuple with the GrossEarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaystubYTDDetails) GetGrossEarningsOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GrossEarnings.Get(), o.GrossEarnings.IsSet()
}

// HasGrossEarnings returns a boolean if a field has been set.
func (o *PaystubYTDDetails) HasGrossEarnings() bool {
	if o != nil && o.GrossEarnings.IsSet() {
		return true
	}

	return false
}

// SetGrossEarnings gets a reference to the given NullableFloat64 and assigns it to the GrossEarnings field.
func (o *PaystubYTDDetails) SetGrossEarnings(v float64) {
	o.GrossEarnings.Set(&v)
}
// SetGrossEarningsNil sets the value for GrossEarnings to be an explicit nil
func (o *PaystubYTDDetails) SetGrossEarningsNil() {
	o.GrossEarnings.Set(nil)
}

// UnsetGrossEarnings ensures that no value is present for GrossEarnings, not even an explicit nil
func (o *PaystubYTDDetails) UnsetGrossEarnings() {
	o.GrossEarnings.Unset()
}

// GetNetEarnings returns the NetEarnings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaystubYTDDetails) GetNetEarnings() float64 {
	if o == nil || o.NetEarnings.Get() == nil {
		var ret float64
		return ret
	}
	return *o.NetEarnings.Get()
}

// GetNetEarningsOk returns a tuple with the NetEarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaystubYTDDetails) GetNetEarningsOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetEarnings.Get(), o.NetEarnings.IsSet()
}

// HasNetEarnings returns a boolean if a field has been set.
func (o *PaystubYTDDetails) HasNetEarnings() bool {
	if o != nil && o.NetEarnings.IsSet() {
		return true
	}

	return false
}

// SetNetEarnings gets a reference to the given NullableFloat64 and assigns it to the NetEarnings field.
func (o *PaystubYTDDetails) SetNetEarnings(v float64) {
	o.NetEarnings.Set(&v)
}
// SetNetEarningsNil sets the value for NetEarnings to be an explicit nil
func (o *PaystubYTDDetails) SetNetEarningsNil() {
	o.NetEarnings.Set(nil)
}

// UnsetNetEarnings ensures that no value is present for NetEarnings, not even an explicit nil
func (o *PaystubYTDDetails) UnsetNetEarnings() {
	o.NetEarnings.Unset()
}

func (o PaystubYTDDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GrossEarnings.IsSet() {
		toSerialize["gross_earnings"] = o.GrossEarnings.Get()
	}
	if o.NetEarnings.IsSet() {
		toSerialize["net_earnings"] = o.NetEarnings.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaystubYTDDetails) UnmarshalJSON(bytes []byte) (err error) {
	varPaystubYTDDetails := _PaystubYTDDetails{}

	if err = json.Unmarshal(bytes, &varPaystubYTDDetails); err == nil {
		*o = PaystubYTDDetails(varPaystubYTDDetails)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "gross_earnings")
		delete(additionalProperties, "net_earnings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaystubYTDDetails struct {
	value *PaystubYTDDetails
	isSet bool
}

func (v NullablePaystubYTDDetails) Get() *PaystubYTDDetails {
	return v.value
}

func (v *NullablePaystubYTDDetails) Set(val *PaystubYTDDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePaystubYTDDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePaystubYTDDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaystubYTDDetails(val *PaystubYTDDetails) *NullablePaystubYTDDetails {
	return &NullablePaystubYTDDetails{value: val, isSet: true}
}

func (v NullablePaystubYTDDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaystubYTDDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


