/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.205.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// Statuses A collection of STATUS containers.
type Statuses struct {
	STATUS Status `json:"STATUS"`
	AdditionalProperties map[string]interface{}
}

type _Statuses Statuses

// NewStatuses instantiates a new Statuses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatuses(sTATUS Status) *Statuses {
	this := Statuses{}
	this.STATUS = sTATUS
	return &this
}

// NewStatusesWithDefaults instantiates a new Statuses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusesWithDefaults() *Statuses {
	this := Statuses{}
	return &this
}

// GetSTATUS returns the STATUS field value
func (o *Statuses) GetSTATUS() Status {
	if o == nil {
		var ret Status
		return ret
	}

	return o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value
// and a boolean to check if the value has been set.
func (o *Statuses) GetSTATUSOk() (*Status, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.STATUS, true
}

// SetSTATUS sets field value
func (o *Statuses) SetSTATUS(v Status) {
	o.STATUS = v
}

func (o Statuses) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["STATUS"] = o.STATUS
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Statuses) UnmarshalJSON(bytes []byte) (err error) {
	varStatuses := _Statuses{}

	if err = json.Unmarshal(bytes, &varStatuses); err == nil {
		*o = Statuses(varStatuses)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "STATUS")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStatuses struct {
	value *Statuses
	isSet bool
}

func (v NullableStatuses) Get() *Statuses {
	return v.value
}

func (v *NullableStatuses) Set(val *Statuses) {
	v.value = val
	v.isSet = true
}

func (v NullableStatuses) IsSet() bool {
	return v.isSet
}

func (v *NullableStatuses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatuses(val *Statuses) *NullableStatuses {
	return &NullableStatuses{value: val, isSet: true}
}

func (v NullableStatuses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatuses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


