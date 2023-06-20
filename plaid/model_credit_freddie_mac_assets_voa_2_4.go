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

// CreditFreddieMacAssetsVOA24 Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
type CreditFreddieMacAssetsVOA24 struct {
	// Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
	ASSET []CreditFreddieMacAssetVOA24 `json:"ASSET"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacAssetsVOA24 CreditFreddieMacAssetsVOA24

// NewCreditFreddieMacAssetsVOA24 instantiates a new CreditFreddieMacAssetsVOA24 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacAssetsVOA24(aSSET []CreditFreddieMacAssetVOA24) *CreditFreddieMacAssetsVOA24 {
	this := CreditFreddieMacAssetsVOA24{}
	this.ASSET = aSSET
	return &this
}

// NewCreditFreddieMacAssetsVOA24WithDefaults instantiates a new CreditFreddieMacAssetsVOA24 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacAssetsVOA24WithDefaults() *CreditFreddieMacAssetsVOA24 {
	this := CreditFreddieMacAssetsVOA24{}
	return &this
}

// GetASSET returns the ASSET field value
func (o *CreditFreddieMacAssetsVOA24) GetASSET() []CreditFreddieMacAssetVOA24 {
	if o == nil {
		var ret []CreditFreddieMacAssetVOA24
		return ret
	}

	return o.ASSET
}

// GetASSETOk returns a tuple with the ASSET field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetsVOA24) GetASSETOk() (*[]CreditFreddieMacAssetVOA24, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASSET, true
}

// SetASSET sets field value
func (o *CreditFreddieMacAssetsVOA24) SetASSET(v []CreditFreddieMacAssetVOA24) {
	o.ASSET = v
}

func (o CreditFreddieMacAssetsVOA24) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ASSET"] = o.ASSET
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacAssetsVOA24) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacAssetsVOA24 := _CreditFreddieMacAssetsVOA24{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacAssetsVOA24); err == nil {
		*o = CreditFreddieMacAssetsVOA24(varCreditFreddieMacAssetsVOA24)
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

type NullableCreditFreddieMacAssetsVOA24 struct {
	value *CreditFreddieMacAssetsVOA24
	isSet bool
}

func (v NullableCreditFreddieMacAssetsVOA24) Get() *CreditFreddieMacAssetsVOA24 {
	return v.value
}

func (v *NullableCreditFreddieMacAssetsVOA24) Set(val *CreditFreddieMacAssetsVOA24) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacAssetsVOA24) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacAssetsVOA24) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacAssetsVOA24(val *CreditFreddieMacAssetsVOA24) *NullableCreditFreddieMacAssetsVOA24 {
	return &NullableCreditFreddieMacAssetsVOA24{value: val, isSet: true}
}

func (v NullableCreditFreddieMacAssetsVOA24) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacAssetsVOA24) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


