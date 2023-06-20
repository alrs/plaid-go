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

// EmploymentDetails An object representing employment details found on a paystub.
type EmploymentDetails struct {
	AnnualSalary *Pay `json:"annual_salary,omitempty"`
	// Date on which the employee was hired, in the YYYY-MM-DD format.
	HireDate NullableString `json:"hire_date,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmploymentDetails EmploymentDetails

// NewEmploymentDetails instantiates a new EmploymentDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmploymentDetails() *EmploymentDetails {
	this := EmploymentDetails{}
	return &this
}

// NewEmploymentDetailsWithDefaults instantiates a new EmploymentDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmploymentDetailsWithDefaults() *EmploymentDetails {
	this := EmploymentDetails{}
	return &this
}

// GetAnnualSalary returns the AnnualSalary field value if set, zero value otherwise.
func (o *EmploymentDetails) GetAnnualSalary() Pay {
	if o == nil || o.AnnualSalary == nil {
		var ret Pay
		return ret
	}
	return *o.AnnualSalary
}

// GetAnnualSalaryOk returns a tuple with the AnnualSalary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmploymentDetails) GetAnnualSalaryOk() (*Pay, bool) {
	if o == nil || o.AnnualSalary == nil {
		return nil, false
	}
	return o.AnnualSalary, true
}

// HasAnnualSalary returns a boolean if a field has been set.
func (o *EmploymentDetails) HasAnnualSalary() bool {
	if o != nil && o.AnnualSalary != nil {
		return true
	}

	return false
}

// SetAnnualSalary gets a reference to the given Pay and assigns it to the AnnualSalary field.
func (o *EmploymentDetails) SetAnnualSalary(v Pay) {
	o.AnnualSalary = &v
}

// GetHireDate returns the HireDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmploymentDetails) GetHireDate() string {
	if o == nil || o.HireDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.HireDate.Get()
}

// GetHireDateOk returns a tuple with the HireDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmploymentDetails) GetHireDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HireDate.Get(), o.HireDate.IsSet()
}

// HasHireDate returns a boolean if a field has been set.
func (o *EmploymentDetails) HasHireDate() bool {
	if o != nil && o.HireDate.IsSet() {
		return true
	}

	return false
}

// SetHireDate gets a reference to the given NullableString and assigns it to the HireDate field.
func (o *EmploymentDetails) SetHireDate(v string) {
	o.HireDate.Set(&v)
}
// SetHireDateNil sets the value for HireDate to be an explicit nil
func (o *EmploymentDetails) SetHireDateNil() {
	o.HireDate.Set(nil)
}

// UnsetHireDate ensures that no value is present for HireDate, not even an explicit nil
func (o *EmploymentDetails) UnsetHireDate() {
	o.HireDate.Unset()
}

func (o EmploymentDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AnnualSalary != nil {
		toSerialize["annual_salary"] = o.AnnualSalary
	}
	if o.HireDate.IsSet() {
		toSerialize["hire_date"] = o.HireDate.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EmploymentDetails) UnmarshalJSON(bytes []byte) (err error) {
	varEmploymentDetails := _EmploymentDetails{}

	if err = json.Unmarshal(bytes, &varEmploymentDetails); err == nil {
		*o = EmploymentDetails(varEmploymentDetails)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "annual_salary")
		delete(additionalProperties, "hire_date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmploymentDetails struct {
	value *EmploymentDetails
	isSet bool
}

func (v NullableEmploymentDetails) Get() *EmploymentDetails {
	return v.value
}

func (v *NullableEmploymentDetails) Set(val *EmploymentDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableEmploymentDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableEmploymentDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmploymentDetails(val *EmploymentDetails) *NullableEmploymentDetails {
	return &NullableEmploymentDetails{value: val, isSet: true}
}

func (v NullableEmploymentDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmploymentDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


