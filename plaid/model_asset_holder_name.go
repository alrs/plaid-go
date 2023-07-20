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

// AssetHolderName Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
type AssetHolderName struct {
	// The unparsed name of either an individual or a legal entity.
	FullName string `json:"FullName"`
	AdditionalProperties map[string]interface{}
}

type _AssetHolderName AssetHolderName

// NewAssetHolderName instantiates a new AssetHolderName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetHolderName(fullName string) *AssetHolderName {
	this := AssetHolderName{}
	this.FullName = fullName
	return &this
}

// NewAssetHolderNameWithDefaults instantiates a new AssetHolderName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetHolderNameWithDefaults() *AssetHolderName {
	this := AssetHolderName{}
	return &this
}

// GetFullName returns the FullName field value
func (o *AssetHolderName) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *AssetHolderName) GetFullNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *AssetHolderName) SetFullName(v string) {
	o.FullName = v
}

func (o AssetHolderName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["FullName"] = o.FullName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetHolderName) UnmarshalJSON(bytes []byte) (err error) {
	varAssetHolderName := _AssetHolderName{}

	if err = json.Unmarshal(bytes, &varAssetHolderName); err == nil {
		*o = AssetHolderName(varAssetHolderName)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "FullName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetHolderName struct {
	value *AssetHolderName
	isSet bool
}

func (v NullableAssetHolderName) Get() *AssetHolderName {
	return v.value
}

func (v *NullableAssetHolderName) Set(val *AssetHolderName) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetHolderName) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetHolderName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetHolderName(val *AssetHolderName) *NullableAssetHolderName {
	return &NullableAssetHolderName{value: val, isSet: true}
}

func (v NullableAssetHolderName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetHolderName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


