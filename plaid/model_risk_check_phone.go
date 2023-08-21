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

// RiskCheckPhone Result summary object specifying values for `phone` attributes of risk check.
type RiskCheckPhone struct {
	// A list of online services where this phone number has been detected to have accounts or other activity.
	LinkedServices []RiskCheckLinkedService `json:"linked_services"`
	AdditionalProperties map[string]interface{}
}

type _RiskCheckPhone RiskCheckPhone

// NewRiskCheckPhone instantiates a new RiskCheckPhone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskCheckPhone(linkedServices []RiskCheckLinkedService) *RiskCheckPhone {
	this := RiskCheckPhone{}
	this.LinkedServices = linkedServices
	return &this
}

// NewRiskCheckPhoneWithDefaults instantiates a new RiskCheckPhone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskCheckPhoneWithDefaults() *RiskCheckPhone {
	this := RiskCheckPhone{}
	return &this
}

// GetLinkedServices returns the LinkedServices field value
func (o *RiskCheckPhone) GetLinkedServices() []RiskCheckLinkedService {
	if o == nil {
		var ret []RiskCheckLinkedService
		return ret
	}

	return o.LinkedServices
}

// GetLinkedServicesOk returns a tuple with the LinkedServices field value
// and a boolean to check if the value has been set.
func (o *RiskCheckPhone) GetLinkedServicesOk() (*[]RiskCheckLinkedService, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LinkedServices, true
}

// SetLinkedServices sets field value
func (o *RiskCheckPhone) SetLinkedServices(v []RiskCheckLinkedService) {
	o.LinkedServices = v
}

func (o RiskCheckPhone) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["linked_services"] = o.LinkedServices
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiskCheckPhone) UnmarshalJSON(bytes []byte) (err error) {
	varRiskCheckPhone := _RiskCheckPhone{}

	if err = json.Unmarshal(bytes, &varRiskCheckPhone); err == nil {
		*o = RiskCheckPhone(varRiskCheckPhone)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "linked_services")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRiskCheckPhone struct {
	value *RiskCheckPhone
	isSet bool
}

func (v NullableRiskCheckPhone) Get() *RiskCheckPhone {
	return v.value
}

func (v *NullableRiskCheckPhone) Set(val *RiskCheckPhone) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskCheckPhone) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskCheckPhone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskCheckPhone(val *RiskCheckPhone) *NullableRiskCheckPhone {
	return &NullableRiskCheckPhone{value: val, isSet: true}
}

func (v NullableRiskCheckPhone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskCheckPhone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


