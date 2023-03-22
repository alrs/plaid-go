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

// CreditFreddieMacVerificationOfAssetResponseVOE25 Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
type CreditFreddieMacVerificationOfAssetResponseVOE25 struct {
	ASSETS CreditFreddieMacAssetsVOE25 `json:"ASSETS"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacVerificationOfAssetResponseVOE25 CreditFreddieMacVerificationOfAssetResponseVOE25

// NewCreditFreddieMacVerificationOfAssetResponseVOE25 instantiates a new CreditFreddieMacVerificationOfAssetResponseVOE25 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacVerificationOfAssetResponseVOE25(aSSETS CreditFreddieMacAssetsVOE25) *CreditFreddieMacVerificationOfAssetResponseVOE25 {
	this := CreditFreddieMacVerificationOfAssetResponseVOE25{}
	this.ASSETS = aSSETS
	return &this
}

// NewCreditFreddieMacVerificationOfAssetResponseVOE25WithDefaults instantiates a new CreditFreddieMacVerificationOfAssetResponseVOE25 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacVerificationOfAssetResponseVOE25WithDefaults() *CreditFreddieMacVerificationOfAssetResponseVOE25 {
	this := CreditFreddieMacVerificationOfAssetResponseVOE25{}
	return &this
}

// GetASSETS returns the ASSETS field value
func (o *CreditFreddieMacVerificationOfAssetResponseVOE25) GetASSETS() CreditFreddieMacAssetsVOE25 {
	if o == nil {
		var ret CreditFreddieMacAssetsVOE25
		return ret
	}

	return o.ASSETS
}

// GetASSETSOk returns a tuple with the ASSETS field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacVerificationOfAssetResponseVOE25) GetASSETSOk() (*CreditFreddieMacAssetsVOE25, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASSETS, true
}

// SetASSETS sets field value
func (o *CreditFreddieMacVerificationOfAssetResponseVOE25) SetASSETS(v CreditFreddieMacAssetsVOE25) {
	o.ASSETS = v
}

func (o CreditFreddieMacVerificationOfAssetResponseVOE25) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ASSETS"] = o.ASSETS
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacVerificationOfAssetResponseVOE25) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacVerificationOfAssetResponseVOE25 := _CreditFreddieMacVerificationOfAssetResponseVOE25{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacVerificationOfAssetResponseVOE25); err == nil {
		*o = CreditFreddieMacVerificationOfAssetResponseVOE25(varCreditFreddieMacVerificationOfAssetResponseVOE25)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ASSETS")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieMacVerificationOfAssetResponseVOE25 struct {
	value *CreditFreddieMacVerificationOfAssetResponseVOE25
	isSet bool
}

func (v NullableCreditFreddieMacVerificationOfAssetResponseVOE25) Get() *CreditFreddieMacVerificationOfAssetResponseVOE25 {
	return v.value
}

func (v *NullableCreditFreddieMacVerificationOfAssetResponseVOE25) Set(val *CreditFreddieMacVerificationOfAssetResponseVOE25) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacVerificationOfAssetResponseVOE25) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacVerificationOfAssetResponseVOE25) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacVerificationOfAssetResponseVOE25(val *CreditFreddieMacVerificationOfAssetResponseVOE25) *NullableCreditFreddieMacVerificationOfAssetResponseVOE25 {
	return &NullableCreditFreddieMacVerificationOfAssetResponseVOE25{value: val, isSet: true}
}

func (v NullableCreditFreddieMacVerificationOfAssetResponseVOE25) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacVerificationOfAssetResponseVOE25) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


