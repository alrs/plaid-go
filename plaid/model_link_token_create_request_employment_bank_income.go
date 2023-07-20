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

// LinkTokenCreateRequestEmploymentBankIncome Specifies options for initializing Link for use with Bank Employment. This field is required if `employment` is included in the `products` array and `bank` is specified in `employment_source_types`.
type LinkTokenCreateRequestEmploymentBankIncome struct {
	// The number of days of data to request for the Bank Employment product.
	DaysRequested int32 `json:"days_requested"`
}

// NewLinkTokenCreateRequestEmploymentBankIncome instantiates a new LinkTokenCreateRequestEmploymentBankIncome object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenCreateRequestEmploymentBankIncome(daysRequested int32) *LinkTokenCreateRequestEmploymentBankIncome {
	this := LinkTokenCreateRequestEmploymentBankIncome{}
	this.DaysRequested = daysRequested
	return &this
}

// NewLinkTokenCreateRequestEmploymentBankIncomeWithDefaults instantiates a new LinkTokenCreateRequestEmploymentBankIncome object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenCreateRequestEmploymentBankIncomeWithDefaults() *LinkTokenCreateRequestEmploymentBankIncome {
	this := LinkTokenCreateRequestEmploymentBankIncome{}
	return &this
}

// GetDaysRequested returns the DaysRequested field value
func (o *LinkTokenCreateRequestEmploymentBankIncome) GetDaysRequested() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DaysRequested
}

// GetDaysRequestedOk returns a tuple with the DaysRequested field value
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestEmploymentBankIncome) GetDaysRequestedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DaysRequested, true
}

// SetDaysRequested sets field value
func (o *LinkTokenCreateRequestEmploymentBankIncome) SetDaysRequested(v int32) {
	o.DaysRequested = v
}

func (o LinkTokenCreateRequestEmploymentBankIncome) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["days_requested"] = o.DaysRequested
	}
	return json.Marshal(toSerialize)
}

type NullableLinkTokenCreateRequestEmploymentBankIncome struct {
	value *LinkTokenCreateRequestEmploymentBankIncome
	isSet bool
}

func (v NullableLinkTokenCreateRequestEmploymentBankIncome) Get() *LinkTokenCreateRequestEmploymentBankIncome {
	return v.value
}

func (v *NullableLinkTokenCreateRequestEmploymentBankIncome) Set(val *LinkTokenCreateRequestEmploymentBankIncome) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenCreateRequestEmploymentBankIncome) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenCreateRequestEmploymentBankIncome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenCreateRequestEmploymentBankIncome(val *LinkTokenCreateRequestEmploymentBankIncome) *NullableLinkTokenCreateRequestEmploymentBankIncome {
	return &NullableLinkTokenCreateRequestEmploymentBankIncome{value: val, isSet: true}
}

func (v NullableLinkTokenCreateRequestEmploymentBankIncome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenCreateRequestEmploymentBankIncome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


