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

// Services A collection of objects that describe requests and responses for services.
type Services struct {
	SERVICE Service `json:"SERVICE"`
	AdditionalProperties map[string]interface{}
}

type _Services Services

// NewServices instantiates a new Services object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServices(sERVICE Service) *Services {
	this := Services{}
	this.SERVICE = sERVICE
	return &this
}

// NewServicesWithDefaults instantiates a new Services object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesWithDefaults() *Services {
	this := Services{}
	return &this
}

// GetSERVICE returns the SERVICE field value
func (o *Services) GetSERVICE() Service {
	if o == nil {
		var ret Service
		return ret
	}

	return o.SERVICE
}

// GetSERVICEOk returns a tuple with the SERVICE field value
// and a boolean to check if the value has been set.
func (o *Services) GetSERVICEOk() (*Service, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SERVICE, true
}

// SetSERVICE sets field value
func (o *Services) SetSERVICE(v Service) {
	o.SERVICE = v
}

func (o Services) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["SERVICE"] = o.SERVICE
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Services) UnmarshalJSON(bytes []byte) (err error) {
	varServices := _Services{}

	if err = json.Unmarshal(bytes, &varServices); err == nil {
		*o = Services(varServices)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "SERVICE")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServices struct {
	value *Services
	isSet bool
}

func (v NullableServices) Get() *Services {
	return v.value
}

func (v *NullableServices) Set(val *Services) {
	v.value = val
	v.isSet = true
}

func (v NullableServices) IsSet() bool {
	return v.isSet
}

func (v *NullableServices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServices(val *Services) *NullableServices {
	return &NullableServices{value: val, isSet: true}
}

func (v NullableServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


