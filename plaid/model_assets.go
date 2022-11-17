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

// Assets No documentation available
type Assets struct {
	// No documentation available
	ASSET []Asset `json:"ASSET"`
	AdditionalProperties map[string]interface{}
}

type _Assets Assets

// NewAssets instantiates a new Assets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssets(aSSET []Asset) *Assets {
	this := Assets{}
	this.ASSET = aSSET
	return &this
}

// NewAssetsWithDefaults instantiates a new Assets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsWithDefaults() *Assets {
	this := Assets{}
	return &this
}

// GetASSET returns the ASSET field value
func (o *Assets) GetASSET() []Asset {
	if o == nil {
		var ret []Asset
		return ret
	}

	return o.ASSET
}

// GetASSETOk returns a tuple with the ASSET field value
// and a boolean to check if the value has been set.
func (o *Assets) GetASSETOk() (*[]Asset, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASSET, true
}

// SetASSET sets field value
func (o *Assets) SetASSET(v []Asset) {
	o.ASSET = v
}

func (o Assets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ASSET"] = o.ASSET
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Assets) UnmarshalJSON(bytes []byte) (err error) {
	varAssets := _Assets{}

	if err = json.Unmarshal(bytes, &varAssets); err == nil {
		*o = Assets(varAssets)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ASSET")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssets struct {
	value *Assets
	isSet bool
}

func (v NullableAssets) Get() *Assets {
	return v.value
}

func (v *NullableAssets) Set(val *Assets) {
	v.value = val
	v.isSet = true
}

func (v NullableAssets) IsSet() bool {
	return v.isSet
}

func (v *NullableAssets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssets(val *Assets) *NullableAssets {
	return &NullableAssets{value: val, isSet: true}
}

func (v NullableAssets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


