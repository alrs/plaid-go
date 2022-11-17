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

// ValidationSources No documentation available.
type ValidationSources struct {
	// No documentation available.
	VALIDATION_SOURCE []ValidationSource `json:"VALIDATION_SOURCE"`
	AdditionalProperties map[string]interface{}
}

type _ValidationSources ValidationSources

// NewValidationSources instantiates a new ValidationSources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidationSources(vALIDATIONSOURCE []ValidationSource) *ValidationSources {
	this := ValidationSources{}
	this.VALIDATION_SOURCE = vALIDATIONSOURCE
	return &this
}

// NewValidationSourcesWithDefaults instantiates a new ValidationSources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationSourcesWithDefaults() *ValidationSources {
	this := ValidationSources{}
	return &this
}

// GetVALIDATION_SOURCE returns the VALIDATION_SOURCE field value
func (o *ValidationSources) GetVALIDATION_SOURCE() []ValidationSource {
	if o == nil {
		var ret []ValidationSource
		return ret
	}

	return o.VALIDATION_SOURCE
}

// GetVALIDATION_SOURCEOk returns a tuple with the VALIDATION_SOURCE field value
// and a boolean to check if the value has been set.
func (o *ValidationSources) GetVALIDATION_SOURCEOk() (*[]ValidationSource, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VALIDATION_SOURCE, true
}

// SetVALIDATION_SOURCE sets field value
func (o *ValidationSources) SetVALIDATION_SOURCE(v []ValidationSource) {
	o.VALIDATION_SOURCE = v
}

func (o ValidationSources) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["VALIDATION_SOURCE"] = o.VALIDATION_SOURCE
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ValidationSources) UnmarshalJSON(bytes []byte) (err error) {
	varValidationSources := _ValidationSources{}

	if err = json.Unmarshal(bytes, &varValidationSources); err == nil {
		*o = ValidationSources(varValidationSources)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "VALIDATION_SOURCE")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableValidationSources struct {
	value *ValidationSources
	isSet bool
}

func (v NullableValidationSources) Get() *ValidationSources {
	return v.value
}

func (v *NullableValidationSources) Set(val *ValidationSources) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationSources) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationSources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationSources(val *ValidationSources) *NullableValidationSources {
	return &NullableValidationSources{value: val, isSet: true}
}

func (v NullableValidationSources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationSources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


