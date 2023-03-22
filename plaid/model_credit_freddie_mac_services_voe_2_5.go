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

// CreditFreddieMacServicesVOE25 A collection of objects that describe requests and responses for services.
type CreditFreddieMacServicesVOE25 struct {
	SERVICE CreditFreddieMacServiceVOE25 `json:"SERVICE"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacServicesVOE25 CreditFreddieMacServicesVOE25

// NewCreditFreddieMacServicesVOE25 instantiates a new CreditFreddieMacServicesVOE25 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacServicesVOE25(sERVICE CreditFreddieMacServiceVOE25) *CreditFreddieMacServicesVOE25 {
	this := CreditFreddieMacServicesVOE25{}
	this.SERVICE = sERVICE
	return &this
}

// NewCreditFreddieMacServicesVOE25WithDefaults instantiates a new CreditFreddieMacServicesVOE25 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacServicesVOE25WithDefaults() *CreditFreddieMacServicesVOE25 {
	this := CreditFreddieMacServicesVOE25{}
	return &this
}

// GetSERVICE returns the SERVICE field value
func (o *CreditFreddieMacServicesVOE25) GetSERVICE() CreditFreddieMacServiceVOE25 {
	if o == nil {
		var ret CreditFreddieMacServiceVOE25
		return ret
	}

	return o.SERVICE
}

// GetSERVICEOk returns a tuple with the SERVICE field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacServicesVOE25) GetSERVICEOk() (*CreditFreddieMacServiceVOE25, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SERVICE, true
}

// SetSERVICE sets field value
func (o *CreditFreddieMacServicesVOE25) SetSERVICE(v CreditFreddieMacServiceVOE25) {
	o.SERVICE = v
}

func (o CreditFreddieMacServicesVOE25) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["SERVICE"] = o.SERVICE
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacServicesVOE25) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacServicesVOE25 := _CreditFreddieMacServicesVOE25{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacServicesVOE25); err == nil {
		*o = CreditFreddieMacServicesVOE25(varCreditFreddieMacServicesVOE25)
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

type NullableCreditFreddieMacServicesVOE25 struct {
	value *CreditFreddieMacServicesVOE25
	isSet bool
}

func (v NullableCreditFreddieMacServicesVOE25) Get() *CreditFreddieMacServicesVOE25 {
	return v.value
}

func (v *NullableCreditFreddieMacServicesVOE25) Set(val *CreditFreddieMacServicesVOE25) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacServicesVOE25) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacServicesVOE25) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacServicesVOE25(val *CreditFreddieMacServicesVOE25) *NullableCreditFreddieMacServicesVOE25 {
	return &NullableCreditFreddieMacServicesVOE25{value: val, isSet: true}
}

func (v NullableCreditFreddieMacServicesVOE25) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacServicesVOE25) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


