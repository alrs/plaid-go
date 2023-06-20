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

// InstitutionSupportedNetworks Contains the RTP network and types supported by the linked Item's institution.
type InstitutionSupportedNetworks struct {
	Rtp TransferCapabilitiesGetRTP `json:"rtp"`
	AdditionalProperties map[string]interface{}
}

type _InstitutionSupportedNetworks InstitutionSupportedNetworks

// NewInstitutionSupportedNetworks instantiates a new InstitutionSupportedNetworks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstitutionSupportedNetworks(rtp TransferCapabilitiesGetRTP) *InstitutionSupportedNetworks {
	this := InstitutionSupportedNetworks{}
	this.Rtp = rtp
	return &this
}

// NewInstitutionSupportedNetworksWithDefaults instantiates a new InstitutionSupportedNetworks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstitutionSupportedNetworksWithDefaults() *InstitutionSupportedNetworks {
	this := InstitutionSupportedNetworks{}
	return &this
}

// GetRtp returns the Rtp field value
func (o *InstitutionSupportedNetworks) GetRtp() TransferCapabilitiesGetRTP {
	if o == nil {
		var ret TransferCapabilitiesGetRTP
		return ret
	}

	return o.Rtp
}

// GetRtpOk returns a tuple with the Rtp field value
// and a boolean to check if the value has been set.
func (o *InstitutionSupportedNetworks) GetRtpOk() (*TransferCapabilitiesGetRTP, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Rtp, true
}

// SetRtp sets field value
func (o *InstitutionSupportedNetworks) SetRtp(v TransferCapabilitiesGetRTP) {
	o.Rtp = v
}

func (o InstitutionSupportedNetworks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rtp"] = o.Rtp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InstitutionSupportedNetworks) UnmarshalJSON(bytes []byte) (err error) {
	varInstitutionSupportedNetworks := _InstitutionSupportedNetworks{}

	if err = json.Unmarshal(bytes, &varInstitutionSupportedNetworks); err == nil {
		*o = InstitutionSupportedNetworks(varInstitutionSupportedNetworks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "rtp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInstitutionSupportedNetworks struct {
	value *InstitutionSupportedNetworks
	isSet bool
}

func (v NullableInstitutionSupportedNetworks) Get() *InstitutionSupportedNetworks {
	return v.value
}

func (v *NullableInstitutionSupportedNetworks) Set(val *InstitutionSupportedNetworks) {
	v.value = val
	v.isSet = true
}

func (v NullableInstitutionSupportedNetworks) IsSet() bool {
	return v.isSet
}

func (v *NullableInstitutionSupportedNetworks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstitutionSupportedNetworks(val *InstitutionSupportedNetworks) *NullableInstitutionSupportedNetworks {
	return &NullableInstitutionSupportedNetworks{value: val, isSet: true}
}

func (v NullableInstitutionSupportedNetworks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstitutionSupportedNetworks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


