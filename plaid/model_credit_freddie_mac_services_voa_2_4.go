/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.419.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// CreditFreddieMacServicesVOA24 A collection of objects that describe requests and responses for services.
type CreditFreddieMacServicesVOA24 struct {
	SERVICE CreditFreddieMacServiceVOA24 `json:"SERVICE"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacServicesVOA24 CreditFreddieMacServicesVOA24

// NewCreditFreddieMacServicesVOA24 instantiates a new CreditFreddieMacServicesVOA24 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacServicesVOA24(sERVICE CreditFreddieMacServiceVOA24) *CreditFreddieMacServicesVOA24 {
	this := CreditFreddieMacServicesVOA24{}
	this.SERVICE = sERVICE
	return &this
}

// NewCreditFreddieMacServicesVOA24WithDefaults instantiates a new CreditFreddieMacServicesVOA24 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacServicesVOA24WithDefaults() *CreditFreddieMacServicesVOA24 {
	this := CreditFreddieMacServicesVOA24{}
	return &this
}

// GetSERVICE returns the SERVICE field value
func (o *CreditFreddieMacServicesVOA24) GetSERVICE() CreditFreddieMacServiceVOA24 {
	if o == nil {
		var ret CreditFreddieMacServiceVOA24
		return ret
	}

	return o.SERVICE
}

// GetSERVICEOk returns a tuple with the SERVICE field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacServicesVOA24) GetSERVICEOk() (*CreditFreddieMacServiceVOA24, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SERVICE, true
}

// SetSERVICE sets field value
func (o *CreditFreddieMacServicesVOA24) SetSERVICE(v CreditFreddieMacServiceVOA24) {
	o.SERVICE = v
}

func (o CreditFreddieMacServicesVOA24) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["SERVICE"] = o.SERVICE
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacServicesVOA24) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacServicesVOA24 := _CreditFreddieMacServicesVOA24{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacServicesVOA24); err == nil {
		*o = CreditFreddieMacServicesVOA24(varCreditFreddieMacServicesVOA24)
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

type NullableCreditFreddieMacServicesVOA24 struct {
	value *CreditFreddieMacServicesVOA24
	isSet bool
}

func (v NullableCreditFreddieMacServicesVOA24) Get() *CreditFreddieMacServicesVOA24 {
	return v.value
}

func (v *NullableCreditFreddieMacServicesVOA24) Set(val *CreditFreddieMacServicesVOA24) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacServicesVOA24) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacServicesVOA24) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacServicesVOA24(val *CreditFreddieMacServicesVOA24) *NullableCreditFreddieMacServicesVOA24 {
	return &NullableCreditFreddieMacServicesVOA24{value: val, isSet: true}
}

func (v NullableCreditFreddieMacServicesVOA24) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacServicesVOA24) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


