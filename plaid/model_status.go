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

// Status Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
type Status struct {
	// Satus Code.
	StatusCode NullableString `json:"StatusCode"`
	// Status Description.
	StatusDescription NullableString `json:"StatusDescription"`
	AdditionalProperties map[string]interface{}
}

type _Status Status

// NewStatus instantiates a new Status object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatus(statusCode NullableString, statusDescription NullableString) *Status {
	this := Status{}
	this.StatusCode = statusCode
	this.StatusDescription = statusDescription
	return &this
}

// NewStatusWithDefaults instantiates a new Status object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusWithDefaults() *Status {
	this := Status{}
	return &this
}

// GetStatusCode returns the StatusCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Status) GetStatusCode() string {
	if o == nil || o.StatusCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.StatusCode.Get()
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Status) GetStatusCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StatusCode.Get(), o.StatusCode.IsSet()
}

// SetStatusCode sets field value
func (o *Status) SetStatusCode(v string) {
	o.StatusCode.Set(&v)
}

// GetStatusDescription returns the StatusDescription field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Status) GetStatusDescription() string {
	if o == nil || o.StatusDescription.Get() == nil {
		var ret string
		return ret
	}

	return *o.StatusDescription.Get()
}

// GetStatusDescriptionOk returns a tuple with the StatusDescription field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Status) GetStatusDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StatusDescription.Get(), o.StatusDescription.IsSet()
}

// SetStatusDescription sets field value
func (o *Status) SetStatusDescription(v string) {
	o.StatusDescription.Set(&v)
}

func (o Status) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["StatusCode"] = o.StatusCode.Get()
	}
	if true {
		toSerialize["StatusDescription"] = o.StatusDescription.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Status) UnmarshalJSON(bytes []byte) (err error) {
	varStatus := _Status{}

	if err = json.Unmarshal(bytes, &varStatus); err == nil {
		*o = Status(varStatus)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "StatusCode")
		delete(additionalProperties, "StatusDescription")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStatus struct {
	value *Status
	isSet bool
}

func (v NullableStatus) Get() *Status {
	return v.value
}

func (v *NullableStatus) Set(val *Status) {
	v.value = val
	v.isSet = true
}

func (v NullableStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatus(val *Status) *NullableStatus {
	return &NullableStatus{value: val, isSet: true}
}

func (v NullableStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


