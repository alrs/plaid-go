/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.181.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// ValidationSource No documentation available.
type ValidationSource struct {
	// No documentation available.
	ValidationSourceName NullableString `json:"ValidationSourceName"`
	// No documentation available.
	ValidationSourceReferenceIdentifier NullableString `json:"ValidationSourceReferenceIdentifier"`
	AdditionalProperties map[string]interface{}
}

type _ValidationSource ValidationSource

// NewValidationSource instantiates a new ValidationSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidationSource(validationSourceName NullableString, validationSourceReferenceIdentifier NullableString) *ValidationSource {
	this := ValidationSource{}
	this.ValidationSourceName = validationSourceName
	this.ValidationSourceReferenceIdentifier = validationSourceReferenceIdentifier
	return &this
}

// NewValidationSourceWithDefaults instantiates a new ValidationSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationSourceWithDefaults() *ValidationSource {
	this := ValidationSource{}
	return &this
}

// GetValidationSourceName returns the ValidationSourceName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ValidationSource) GetValidationSourceName() string {
	if o == nil || o.ValidationSourceName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ValidationSourceName.Get()
}

// GetValidationSourceNameOk returns a tuple with the ValidationSourceName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ValidationSource) GetValidationSourceNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ValidationSourceName.Get(), o.ValidationSourceName.IsSet()
}

// SetValidationSourceName sets field value
func (o *ValidationSource) SetValidationSourceName(v string) {
	o.ValidationSourceName.Set(&v)
}

// GetValidationSourceReferenceIdentifier returns the ValidationSourceReferenceIdentifier field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ValidationSource) GetValidationSourceReferenceIdentifier() string {
	if o == nil || o.ValidationSourceReferenceIdentifier.Get() == nil {
		var ret string
		return ret
	}

	return *o.ValidationSourceReferenceIdentifier.Get()
}

// GetValidationSourceReferenceIdentifierOk returns a tuple with the ValidationSourceReferenceIdentifier field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ValidationSource) GetValidationSourceReferenceIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ValidationSourceReferenceIdentifier.Get(), o.ValidationSourceReferenceIdentifier.IsSet()
}

// SetValidationSourceReferenceIdentifier sets field value
func (o *ValidationSource) SetValidationSourceReferenceIdentifier(v string) {
	o.ValidationSourceReferenceIdentifier.Set(&v)
}

func (o ValidationSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ValidationSourceName"] = o.ValidationSourceName.Get()
	}
	if true {
		toSerialize["ValidationSourceReferenceIdentifier"] = o.ValidationSourceReferenceIdentifier.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ValidationSource) UnmarshalJSON(bytes []byte) (err error) {
	varValidationSource := _ValidationSource{}

	if err = json.Unmarshal(bytes, &varValidationSource); err == nil {
		*o = ValidationSource(varValidationSource)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ValidationSourceName")
		delete(additionalProperties, "ValidationSourceReferenceIdentifier")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableValidationSource struct {
	value *ValidationSource
	isSet bool
}

func (v NullableValidationSource) Get() *ValidationSource {
	return v.value
}

func (v *NullableValidationSource) Set(val *ValidationSource) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationSource) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationSource(val *ValidationSource) *NullableValidationSource {
	return &NullableValidationSource{value: val, isSet: true}
}

func (v NullableValidationSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


