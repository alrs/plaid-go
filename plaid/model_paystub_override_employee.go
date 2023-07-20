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

// PaystubOverrideEmployee The employee on the paystub.
type PaystubOverrideEmployee struct {
	// The name of the employee.
	Name *string `json:"name,omitempty"`
	Address *PaystubOverrideEmployeeAddress `json:"address,omitempty"`
}

// NewPaystubOverrideEmployee instantiates a new PaystubOverrideEmployee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaystubOverrideEmployee() *PaystubOverrideEmployee {
	this := PaystubOverrideEmployee{}
	return &this
}

// NewPaystubOverrideEmployeeWithDefaults instantiates a new PaystubOverrideEmployee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaystubOverrideEmployeeWithDefaults() *PaystubOverrideEmployee {
	this := PaystubOverrideEmployee{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PaystubOverrideEmployee) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaystubOverrideEmployee) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PaystubOverrideEmployee) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PaystubOverrideEmployee) SetName(v string) {
	o.Name = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *PaystubOverrideEmployee) GetAddress() PaystubOverrideEmployeeAddress {
	if o == nil || o.Address == nil {
		var ret PaystubOverrideEmployeeAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaystubOverrideEmployee) GetAddressOk() (*PaystubOverrideEmployeeAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *PaystubOverrideEmployee) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given PaystubOverrideEmployeeAddress and assigns it to the Address field.
func (o *PaystubOverrideEmployee) SetAddress(v PaystubOverrideEmployeeAddress) {
	o.Address = &v
}

func (o PaystubOverrideEmployee) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullablePaystubOverrideEmployee struct {
	value *PaystubOverrideEmployee
	isSet bool
}

func (v NullablePaystubOverrideEmployee) Get() *PaystubOverrideEmployee {
	return v.value
}

func (v *NullablePaystubOverrideEmployee) Set(val *PaystubOverrideEmployee) {
	v.value = val
	v.isSet = true
}

func (v NullablePaystubOverrideEmployee) IsSet() bool {
	return v.isSet
}

func (v *NullablePaystubOverrideEmployee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaystubOverrideEmployee(val *PaystubOverrideEmployee) *NullablePaystubOverrideEmployee {
	return &NullablePaystubOverrideEmployee{value: val, isSet: true}
}

func (v NullablePaystubOverrideEmployee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaystubOverrideEmployee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


