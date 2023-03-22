/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.343.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// PhoneNumber A phone number
type PhoneNumber struct {
	// The phone number.
	Data string `json:"data"`
	// When `true`, identifies the phone number as the primary number on an account.
	Primary bool `json:"primary"`
	// The type of phone number.
	Type string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _PhoneNumber PhoneNumber

// NewPhoneNumber instantiates a new PhoneNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhoneNumber(data string, primary bool, type_ string) *PhoneNumber {
	this := PhoneNumber{}
	this.Data = data
	this.Primary = primary
	this.Type = type_
	return &this
}

// NewPhoneNumberWithDefaults instantiates a new PhoneNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhoneNumberWithDefaults() *PhoneNumber {
	this := PhoneNumber{}
	return &this
}

// GetData returns the Data field value
func (o *PhoneNumber) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PhoneNumber) GetDataOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PhoneNumber) SetData(v string) {
	o.Data = v
}

// GetPrimary returns the Primary field value
func (o *PhoneNumber) GetPrimary() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value
// and a boolean to check if the value has been set.
func (o *PhoneNumber) GetPrimaryOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Primary, true
}

// SetPrimary sets field value
func (o *PhoneNumber) SetPrimary(v bool) {
	o.Primary = v
}

// GetType returns the Type field value
func (o *PhoneNumber) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PhoneNumber) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PhoneNumber) SetType(v string) {
	o.Type = v
}

func (o PhoneNumber) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["primary"] = o.Primary
	}
	if true {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PhoneNumber) UnmarshalJSON(bytes []byte) (err error) {
	varPhoneNumber := _PhoneNumber{}

	if err = json.Unmarshal(bytes, &varPhoneNumber); err == nil {
		*o = PhoneNumber(varPhoneNumber)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "primary")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePhoneNumber struct {
	value *PhoneNumber
	isSet bool
}

func (v NullablePhoneNumber) Get() *PhoneNumber {
	return v.value
}

func (v *NullablePhoneNumber) Set(val *PhoneNumber) {
	v.value = val
	v.isSet = true
}

func (v NullablePhoneNumber) IsSet() bool {
	return v.isSet
}

func (v *NullablePhoneNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhoneNumber(val *PhoneNumber) *NullablePhoneNumber {
	return &NullablePhoneNumber{value: val, isSet: true}
}

func (v NullablePhoneNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhoneNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


