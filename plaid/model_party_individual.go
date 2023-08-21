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

// PartyIndividual Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
type PartyIndividual struct {
	NAME IndividualName `json:"NAME"`
	AdditionalProperties map[string]interface{}
}

type _PartyIndividual PartyIndividual

// NewPartyIndividual instantiates a new PartyIndividual object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartyIndividual(nAME IndividualName) *PartyIndividual {
	this := PartyIndividual{}
	this.NAME = nAME
	return &this
}

// NewPartyIndividualWithDefaults instantiates a new PartyIndividual object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartyIndividualWithDefaults() *PartyIndividual {
	this := PartyIndividual{}
	return &this
}

// GetNAME returns the NAME field value
func (o *PartyIndividual) GetNAME() IndividualName {
	if o == nil {
		var ret IndividualName
		return ret
	}

	return o.NAME
}

// GetNAMEOk returns a tuple with the NAME field value
// and a boolean to check if the value has been set.
func (o *PartyIndividual) GetNAMEOk() (*IndividualName, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NAME, true
}

// SetNAME sets field value
func (o *PartyIndividual) SetNAME(v IndividualName) {
	o.NAME = v
}

func (o PartyIndividual) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["NAME"] = o.NAME
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PartyIndividual) UnmarshalJSON(bytes []byte) (err error) {
	varPartyIndividual := _PartyIndividual{}

	if err = json.Unmarshal(bytes, &varPartyIndividual); err == nil {
		*o = PartyIndividual(varPartyIndividual)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "NAME")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePartyIndividual struct {
	value *PartyIndividual
	isSet bool
}

func (v NullablePartyIndividual) Get() *PartyIndividual {
	return v.value
}

func (v *NullablePartyIndividual) Set(val *PartyIndividual) {
	v.value = val
	v.isSet = true
}

func (v NullablePartyIndividual) IsSet() bool {
	return v.isSet
}

func (v *NullablePartyIndividual) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartyIndividual(val *PartyIndividual) *NullablePartyIndividual {
	return &NullablePartyIndividual{value: val, isSet: true}
}

func (v NullablePartyIndividual) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartyIndividual) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


