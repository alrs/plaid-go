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

// IncomeVerificationPrecheckEmployer Information about the end user's employer
type IncomeVerificationPrecheckEmployer struct {
	// The employer's name
	Name NullableString `json:"name,omitempty"`
	Address NullableIncomeVerificationPrecheckEmployerAddress `json:"address,omitempty"`
	// The employer's tax id
	TaxId NullableString `json:"tax_id,omitempty"`
	// The URL for the employer's public website
	Url NullableString `json:"url,omitempty"`
}

// NewIncomeVerificationPrecheckEmployer instantiates a new IncomeVerificationPrecheckEmployer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeVerificationPrecheckEmployer() *IncomeVerificationPrecheckEmployer {
	this := IncomeVerificationPrecheckEmployer{}
	return &this
}

// NewIncomeVerificationPrecheckEmployerWithDefaults instantiates a new IncomeVerificationPrecheckEmployer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeVerificationPrecheckEmployerWithDefaults() *IncomeVerificationPrecheckEmployer {
	this := IncomeVerificationPrecheckEmployer{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeVerificationPrecheckEmployer) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeVerificationPrecheckEmployer) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckEmployer) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *IncomeVerificationPrecheckEmployer) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *IncomeVerificationPrecheckEmployer) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *IncomeVerificationPrecheckEmployer) UnsetName() {
	o.Name.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeVerificationPrecheckEmployer) GetAddress() IncomeVerificationPrecheckEmployerAddress {
	if o == nil || o.Address.Get() == nil {
		var ret IncomeVerificationPrecheckEmployerAddress
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeVerificationPrecheckEmployer) GetAddressOk() (*IncomeVerificationPrecheckEmployerAddress, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckEmployer) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableIncomeVerificationPrecheckEmployerAddress and assigns it to the Address field.
func (o *IncomeVerificationPrecheckEmployer) SetAddress(v IncomeVerificationPrecheckEmployerAddress) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *IncomeVerificationPrecheckEmployer) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *IncomeVerificationPrecheckEmployer) UnsetAddress() {
	o.Address.Unset()
}

// GetTaxId returns the TaxId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeVerificationPrecheckEmployer) GetTaxId() string {
	if o == nil || o.TaxId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TaxId.Get()
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeVerificationPrecheckEmployer) GetTaxIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TaxId.Get(), o.TaxId.IsSet()
}

// HasTaxId returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckEmployer) HasTaxId() bool {
	if o != nil && o.TaxId.IsSet() {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given NullableString and assigns it to the TaxId field.
func (o *IncomeVerificationPrecheckEmployer) SetTaxId(v string) {
	o.TaxId.Set(&v)
}
// SetTaxIdNil sets the value for TaxId to be an explicit nil
func (o *IncomeVerificationPrecheckEmployer) SetTaxIdNil() {
	o.TaxId.Set(nil)
}

// UnsetTaxId ensures that no value is present for TaxId, not even an explicit nil
func (o *IncomeVerificationPrecheckEmployer) UnsetTaxId() {
	o.TaxId.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeVerificationPrecheckEmployer) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeVerificationPrecheckEmployer) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckEmployer) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *IncomeVerificationPrecheckEmployer) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *IncomeVerificationPrecheckEmployer) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *IncomeVerificationPrecheckEmployer) UnsetUrl() {
	o.Url.Unset()
}

func (o IncomeVerificationPrecheckEmployer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.TaxId.IsSet() {
		toSerialize["tax_id"] = o.TaxId.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIncomeVerificationPrecheckEmployer struct {
	value *IncomeVerificationPrecheckEmployer
	isSet bool
}

func (v NullableIncomeVerificationPrecheckEmployer) Get() *IncomeVerificationPrecheckEmployer {
	return v.value
}

func (v *NullableIncomeVerificationPrecheckEmployer) Set(val *IncomeVerificationPrecheckEmployer) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationPrecheckEmployer) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationPrecheckEmployer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationPrecheckEmployer(val *IncomeVerificationPrecheckEmployer) *NullableIncomeVerificationPrecheckEmployer {
	return &NullableIncomeVerificationPrecheckEmployer{value: val, isSet: true}
}

func (v NullableIncomeVerificationPrecheckEmployer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationPrecheckEmployer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


