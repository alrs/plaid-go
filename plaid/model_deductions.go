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

// Deductions An object with the deduction information found on a paystub.
type Deductions struct {
	Subtotals *[]Total `json:"subtotals,omitempty"`
	Breakdown []DeductionsBreakdown `json:"breakdown"`
	Totals *[]Total `json:"totals,omitempty"`
	Total DeductionsTotal `json:"total"`
	AdditionalProperties map[string]interface{}
}

type _Deductions Deductions

// NewDeductions instantiates a new Deductions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeductions(breakdown []DeductionsBreakdown, total DeductionsTotal) *Deductions {
	this := Deductions{}
	this.Breakdown = breakdown
	this.Total = total
	return &this
}

// NewDeductionsWithDefaults instantiates a new Deductions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeductionsWithDefaults() *Deductions {
	this := Deductions{}
	return &this
}

// GetSubtotals returns the Subtotals field value if set, zero value otherwise.
func (o *Deductions) GetSubtotals() []Total {
	if o == nil || o.Subtotals == nil {
		var ret []Total
		return ret
	}
	return *o.Subtotals
}

// GetSubtotalsOk returns a tuple with the Subtotals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deductions) GetSubtotalsOk() (*[]Total, bool) {
	if o == nil || o.Subtotals == nil {
		return nil, false
	}
	return o.Subtotals, true
}

// HasSubtotals returns a boolean if a field has been set.
func (o *Deductions) HasSubtotals() bool {
	if o != nil && o.Subtotals != nil {
		return true
	}

	return false
}

// SetSubtotals gets a reference to the given []Total and assigns it to the Subtotals field.
func (o *Deductions) SetSubtotals(v []Total) {
	o.Subtotals = &v
}

// GetBreakdown returns the Breakdown field value
func (o *Deductions) GetBreakdown() []DeductionsBreakdown {
	if o == nil {
		var ret []DeductionsBreakdown
		return ret
	}

	return o.Breakdown
}

// GetBreakdownOk returns a tuple with the Breakdown field value
// and a boolean to check if the value has been set.
func (o *Deductions) GetBreakdownOk() (*[]DeductionsBreakdown, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Breakdown, true
}

// SetBreakdown sets field value
func (o *Deductions) SetBreakdown(v []DeductionsBreakdown) {
	o.Breakdown = v
}

// GetTotals returns the Totals field value if set, zero value otherwise.
func (o *Deductions) GetTotals() []Total {
	if o == nil || o.Totals == nil {
		var ret []Total
		return ret
	}
	return *o.Totals
}

// GetTotalsOk returns a tuple with the Totals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deductions) GetTotalsOk() (*[]Total, bool) {
	if o == nil || o.Totals == nil {
		return nil, false
	}
	return o.Totals, true
}

// HasTotals returns a boolean if a field has been set.
func (o *Deductions) HasTotals() bool {
	if o != nil && o.Totals != nil {
		return true
	}

	return false
}

// SetTotals gets a reference to the given []Total and assigns it to the Totals field.
func (o *Deductions) SetTotals(v []Total) {
	o.Totals = &v
}

// GetTotal returns the Total field value
func (o *Deductions) GetTotal() DeductionsTotal {
	if o == nil {
		var ret DeductionsTotal
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *Deductions) GetTotalOk() (*DeductionsTotal, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *Deductions) SetTotal(v DeductionsTotal) {
	o.Total = v
}

func (o Deductions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Subtotals != nil {
		toSerialize["subtotals"] = o.Subtotals
	}
	if true {
		toSerialize["breakdown"] = o.Breakdown
	}
	if o.Totals != nil {
		toSerialize["totals"] = o.Totals
	}
	if true {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Deductions) UnmarshalJSON(bytes []byte) (err error) {
	varDeductions := _Deductions{}

	if err = json.Unmarshal(bytes, &varDeductions); err == nil {
		*o = Deductions(varDeductions)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "subtotals")
		delete(additionalProperties, "breakdown")
		delete(additionalProperties, "totals")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeductions struct {
	value *Deductions
	isSet bool
}

func (v NullableDeductions) Get() *Deductions {
	return v.value
}

func (v *NullableDeductions) Set(val *Deductions) {
	v.value = val
	v.isSet = true
}

func (v NullableDeductions) IsSet() bool {
	return v.isSet
}

func (v *NullableDeductions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeductions(val *Deductions) *NullableDeductions {
	return &NullableDeductions{value: val, isSet: true}
}

func (v NullableDeductions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeductions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


