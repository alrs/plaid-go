/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.26.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// IncomeBreakdown An object representing a breakdown of the different income types on the paystub.
type IncomeBreakdown struct {
	// The type of income. Possible values include:   `\"regular\"`: regular income   `\"overtime\"`: overtime income    `\"bonus\"`: bonus income
	Type NullableString `json:"type"`
	// The hourly rate at which the income is paid.
	Rate NullableFloat32 `json:"rate"`
	// The number of hours logged for this income for this pay period.
	Hours NullableFloat32 `json:"hours"`
	// The total pay for this pay period.
	Total NullableFloat32 `json:"total"`
	AdditionalProperties map[string]interface{}
}

type _IncomeBreakdown IncomeBreakdown

// NewIncomeBreakdown instantiates a new IncomeBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeBreakdown(type_ NullableString, rate NullableFloat32, hours NullableFloat32, total NullableFloat32) *IncomeBreakdown {
	this := IncomeBreakdown{}
	this.Type = type_
	this.Rate = rate
	this.Hours = hours
	this.Total = total
	return &this
}

// NewIncomeBreakdownWithDefaults instantiates a new IncomeBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeBreakdownWithDefaults() *IncomeBreakdown {
	this := IncomeBreakdown{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IncomeBreakdown) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeBreakdown) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *IncomeBreakdown) SetType(v string) {
	o.Type.Set(&v)
}

// GetRate returns the Rate field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *IncomeBreakdown) GetRate() float32 {
	if o == nil || o.Rate.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Rate.Get()
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeBreakdown) GetRateOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Rate.Get(), o.Rate.IsSet()
}

// SetRate sets field value
func (o *IncomeBreakdown) SetRate(v float32) {
	o.Rate.Set(&v)
}

// GetHours returns the Hours field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *IncomeBreakdown) GetHours() float32 {
	if o == nil || o.Hours.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Hours.Get()
}

// GetHoursOk returns a tuple with the Hours field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeBreakdown) GetHoursOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Hours.Get(), o.Hours.IsSet()
}

// SetHours sets field value
func (o *IncomeBreakdown) SetHours(v float32) {
	o.Hours.Set(&v)
}

// GetTotal returns the Total field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *IncomeBreakdown) GetTotal() float32 {
	if o == nil || o.Total.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Total.Get()
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeBreakdown) GetTotalOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Total.Get(), o.Total.IsSet()
}

// SetTotal sets field value
func (o *IncomeBreakdown) SetTotal(v float32) {
	o.Total.Set(&v)
}

func (o IncomeBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type.Get()
	}
	if true {
		toSerialize["rate"] = o.Rate.Get()
	}
	if true {
		toSerialize["hours"] = o.Hours.Get()
	}
	if true {
		toSerialize["total"] = o.Total.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IncomeBreakdown) UnmarshalJSON(bytes []byte) (err error) {
	varIncomeBreakdown := _IncomeBreakdown{}

	if err = json.Unmarshal(bytes, &varIncomeBreakdown); err == nil {
		*o = IncomeBreakdown(varIncomeBreakdown)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "rate")
		delete(additionalProperties, "hours")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIncomeBreakdown struct {
	value *IncomeBreakdown
	isSet bool
}

func (v NullableIncomeBreakdown) Get() *IncomeBreakdown {
	return v.value
}

func (v *NullableIncomeBreakdown) Set(val *IncomeBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeBreakdown(val *IncomeBreakdown) *NullableIncomeBreakdown {
	return &NullableIncomeBreakdown{value: val, isSet: true}
}

func (v NullableIncomeBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


