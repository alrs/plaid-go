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

// W2StateAndLocalWages struct for W2StateAndLocalWages
type W2StateAndLocalWages struct {
	// State associated with the wage.
	State NullableString `json:"state,omitempty"`
	// State identification number of the employer.
	EmployerStateIdNumber NullableString `json:"employer_state_id_number,omitempty"`
	// Wages and tips from the specified state.
	StateWagesTips NullableString `json:"state_wages_tips,omitempty"`
	// Income tax from the specified state.
	StateIncomeTax NullableString `json:"state_income_tax,omitempty"`
	// Wages and tips from the locality.
	LocalWagesTips NullableString `json:"local_wages_tips,omitempty"`
	// Income tax from the locality.
	LocalIncomeTax NullableString `json:"local_income_tax,omitempty"`
	// Name of the locality.
	LocalityName NullableString `json:"locality_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _W2StateAndLocalWages W2StateAndLocalWages

// NewW2StateAndLocalWages instantiates a new W2StateAndLocalWages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewW2StateAndLocalWages() *W2StateAndLocalWages {
	this := W2StateAndLocalWages{}
	return &this
}

// NewW2StateAndLocalWagesWithDefaults instantiates a new W2StateAndLocalWages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewW2StateAndLocalWagesWithDefaults() *W2StateAndLocalWages {
	this := W2StateAndLocalWages{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2StateAndLocalWages) GetState() string {
	if o == nil || o.State.Get() == nil {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2StateAndLocalWages) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *W2StateAndLocalWages) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *W2StateAndLocalWages) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *W2StateAndLocalWages) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *W2StateAndLocalWages) UnsetState() {
	o.State.Unset()
}

// GetEmployerStateIdNumber returns the EmployerStateIdNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2StateAndLocalWages) GetEmployerStateIdNumber() string {
	if o == nil || o.EmployerStateIdNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.EmployerStateIdNumber.Get()
}

// GetEmployerStateIdNumberOk returns a tuple with the EmployerStateIdNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2StateAndLocalWages) GetEmployerStateIdNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmployerStateIdNumber.Get(), o.EmployerStateIdNumber.IsSet()
}

// HasEmployerStateIdNumber returns a boolean if a field has been set.
func (o *W2StateAndLocalWages) HasEmployerStateIdNumber() bool {
	if o != nil && o.EmployerStateIdNumber.IsSet() {
		return true
	}

	return false
}

// SetEmployerStateIdNumber gets a reference to the given NullableString and assigns it to the EmployerStateIdNumber field.
func (o *W2StateAndLocalWages) SetEmployerStateIdNumber(v string) {
	o.EmployerStateIdNumber.Set(&v)
}
// SetEmployerStateIdNumberNil sets the value for EmployerStateIdNumber to be an explicit nil
func (o *W2StateAndLocalWages) SetEmployerStateIdNumberNil() {
	o.EmployerStateIdNumber.Set(nil)
}

// UnsetEmployerStateIdNumber ensures that no value is present for EmployerStateIdNumber, not even an explicit nil
func (o *W2StateAndLocalWages) UnsetEmployerStateIdNumber() {
	o.EmployerStateIdNumber.Unset()
}

// GetStateWagesTips returns the StateWagesTips field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2StateAndLocalWages) GetStateWagesTips() string {
	if o == nil || o.StateWagesTips.Get() == nil {
		var ret string
		return ret
	}
	return *o.StateWagesTips.Get()
}

// GetStateWagesTipsOk returns a tuple with the StateWagesTips field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2StateAndLocalWages) GetStateWagesTipsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StateWagesTips.Get(), o.StateWagesTips.IsSet()
}

// HasStateWagesTips returns a boolean if a field has been set.
func (o *W2StateAndLocalWages) HasStateWagesTips() bool {
	if o != nil && o.StateWagesTips.IsSet() {
		return true
	}

	return false
}

// SetStateWagesTips gets a reference to the given NullableString and assigns it to the StateWagesTips field.
func (o *W2StateAndLocalWages) SetStateWagesTips(v string) {
	o.StateWagesTips.Set(&v)
}
// SetStateWagesTipsNil sets the value for StateWagesTips to be an explicit nil
func (o *W2StateAndLocalWages) SetStateWagesTipsNil() {
	o.StateWagesTips.Set(nil)
}

// UnsetStateWagesTips ensures that no value is present for StateWagesTips, not even an explicit nil
func (o *W2StateAndLocalWages) UnsetStateWagesTips() {
	o.StateWagesTips.Unset()
}

// GetStateIncomeTax returns the StateIncomeTax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2StateAndLocalWages) GetStateIncomeTax() string {
	if o == nil || o.StateIncomeTax.Get() == nil {
		var ret string
		return ret
	}
	return *o.StateIncomeTax.Get()
}

// GetStateIncomeTaxOk returns a tuple with the StateIncomeTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2StateAndLocalWages) GetStateIncomeTaxOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StateIncomeTax.Get(), o.StateIncomeTax.IsSet()
}

// HasStateIncomeTax returns a boolean if a field has been set.
func (o *W2StateAndLocalWages) HasStateIncomeTax() bool {
	if o != nil && o.StateIncomeTax.IsSet() {
		return true
	}

	return false
}

// SetStateIncomeTax gets a reference to the given NullableString and assigns it to the StateIncomeTax field.
func (o *W2StateAndLocalWages) SetStateIncomeTax(v string) {
	o.StateIncomeTax.Set(&v)
}
// SetStateIncomeTaxNil sets the value for StateIncomeTax to be an explicit nil
func (o *W2StateAndLocalWages) SetStateIncomeTaxNil() {
	o.StateIncomeTax.Set(nil)
}

// UnsetStateIncomeTax ensures that no value is present for StateIncomeTax, not even an explicit nil
func (o *W2StateAndLocalWages) UnsetStateIncomeTax() {
	o.StateIncomeTax.Unset()
}

// GetLocalWagesTips returns the LocalWagesTips field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2StateAndLocalWages) GetLocalWagesTips() string {
	if o == nil || o.LocalWagesTips.Get() == nil {
		var ret string
		return ret
	}
	return *o.LocalWagesTips.Get()
}

// GetLocalWagesTipsOk returns a tuple with the LocalWagesTips field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2StateAndLocalWages) GetLocalWagesTipsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LocalWagesTips.Get(), o.LocalWagesTips.IsSet()
}

// HasLocalWagesTips returns a boolean if a field has been set.
func (o *W2StateAndLocalWages) HasLocalWagesTips() bool {
	if o != nil && o.LocalWagesTips.IsSet() {
		return true
	}

	return false
}

// SetLocalWagesTips gets a reference to the given NullableString and assigns it to the LocalWagesTips field.
func (o *W2StateAndLocalWages) SetLocalWagesTips(v string) {
	o.LocalWagesTips.Set(&v)
}
// SetLocalWagesTipsNil sets the value for LocalWagesTips to be an explicit nil
func (o *W2StateAndLocalWages) SetLocalWagesTipsNil() {
	o.LocalWagesTips.Set(nil)
}

// UnsetLocalWagesTips ensures that no value is present for LocalWagesTips, not even an explicit nil
func (o *W2StateAndLocalWages) UnsetLocalWagesTips() {
	o.LocalWagesTips.Unset()
}

// GetLocalIncomeTax returns the LocalIncomeTax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2StateAndLocalWages) GetLocalIncomeTax() string {
	if o == nil || o.LocalIncomeTax.Get() == nil {
		var ret string
		return ret
	}
	return *o.LocalIncomeTax.Get()
}

// GetLocalIncomeTaxOk returns a tuple with the LocalIncomeTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2StateAndLocalWages) GetLocalIncomeTaxOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LocalIncomeTax.Get(), o.LocalIncomeTax.IsSet()
}

// HasLocalIncomeTax returns a boolean if a field has been set.
func (o *W2StateAndLocalWages) HasLocalIncomeTax() bool {
	if o != nil && o.LocalIncomeTax.IsSet() {
		return true
	}

	return false
}

// SetLocalIncomeTax gets a reference to the given NullableString and assigns it to the LocalIncomeTax field.
func (o *W2StateAndLocalWages) SetLocalIncomeTax(v string) {
	o.LocalIncomeTax.Set(&v)
}
// SetLocalIncomeTaxNil sets the value for LocalIncomeTax to be an explicit nil
func (o *W2StateAndLocalWages) SetLocalIncomeTaxNil() {
	o.LocalIncomeTax.Set(nil)
}

// UnsetLocalIncomeTax ensures that no value is present for LocalIncomeTax, not even an explicit nil
func (o *W2StateAndLocalWages) UnsetLocalIncomeTax() {
	o.LocalIncomeTax.Unset()
}

// GetLocalityName returns the LocalityName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2StateAndLocalWages) GetLocalityName() string {
	if o == nil || o.LocalityName.Get() == nil {
		var ret string
		return ret
	}
	return *o.LocalityName.Get()
}

// GetLocalityNameOk returns a tuple with the LocalityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2StateAndLocalWages) GetLocalityNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LocalityName.Get(), o.LocalityName.IsSet()
}

// HasLocalityName returns a boolean if a field has been set.
func (o *W2StateAndLocalWages) HasLocalityName() bool {
	if o != nil && o.LocalityName.IsSet() {
		return true
	}

	return false
}

// SetLocalityName gets a reference to the given NullableString and assigns it to the LocalityName field.
func (o *W2StateAndLocalWages) SetLocalityName(v string) {
	o.LocalityName.Set(&v)
}
// SetLocalityNameNil sets the value for LocalityName to be an explicit nil
func (o *W2StateAndLocalWages) SetLocalityNameNil() {
	o.LocalityName.Set(nil)
}

// UnsetLocalityName ensures that no value is present for LocalityName, not even an explicit nil
func (o *W2StateAndLocalWages) UnsetLocalityName() {
	o.LocalityName.Unset()
}

func (o W2StateAndLocalWages) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.EmployerStateIdNumber.IsSet() {
		toSerialize["employer_state_id_number"] = o.EmployerStateIdNumber.Get()
	}
	if o.StateWagesTips.IsSet() {
		toSerialize["state_wages_tips"] = o.StateWagesTips.Get()
	}
	if o.StateIncomeTax.IsSet() {
		toSerialize["state_income_tax"] = o.StateIncomeTax.Get()
	}
	if o.LocalWagesTips.IsSet() {
		toSerialize["local_wages_tips"] = o.LocalWagesTips.Get()
	}
	if o.LocalIncomeTax.IsSet() {
		toSerialize["local_income_tax"] = o.LocalIncomeTax.Get()
	}
	if o.LocalityName.IsSet() {
		toSerialize["locality_name"] = o.LocalityName.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *W2StateAndLocalWages) UnmarshalJSON(bytes []byte) (err error) {
	varW2StateAndLocalWages := _W2StateAndLocalWages{}

	if err = json.Unmarshal(bytes, &varW2StateAndLocalWages); err == nil {
		*o = W2StateAndLocalWages(varW2StateAndLocalWages)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "state")
		delete(additionalProperties, "employer_state_id_number")
		delete(additionalProperties, "state_wages_tips")
		delete(additionalProperties, "state_income_tax")
		delete(additionalProperties, "local_wages_tips")
		delete(additionalProperties, "local_income_tax")
		delete(additionalProperties, "locality_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableW2StateAndLocalWages struct {
	value *W2StateAndLocalWages
	isSet bool
}

func (v NullableW2StateAndLocalWages) Get() *W2StateAndLocalWages {
	return v.value
}

func (v *NullableW2StateAndLocalWages) Set(val *W2StateAndLocalWages) {
	v.value = val
	v.isSet = true
}

func (v NullableW2StateAndLocalWages) IsSet() bool {
	return v.isSet
}

func (v *NullableW2StateAndLocalWages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableW2StateAndLocalWages(val *W2StateAndLocalWages) *NullableW2StateAndLocalWages {
	return &NullableW2StateAndLocalWages{value: val, isSet: true}
}

func (v NullableW2StateAndLocalWages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableW2StateAndLocalWages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


