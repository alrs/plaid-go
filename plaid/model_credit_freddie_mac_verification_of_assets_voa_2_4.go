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

// CreditFreddieMacVerificationOfAssetsVOA24 Verification of Assets Report
type CreditFreddieMacVerificationOfAssetsVOA24 struct {
	DEAL CreditFreddieMacVerificationOfAssetsDealVOA24 `json:"DEAL"`
	// The Verification Of Assets (VOA) schema version.
	SchemaVersion float32 `json:"SchemaVersion"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacVerificationOfAssetsVOA24 CreditFreddieMacVerificationOfAssetsVOA24

// NewCreditFreddieMacVerificationOfAssetsVOA24 instantiates a new CreditFreddieMacVerificationOfAssetsVOA24 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacVerificationOfAssetsVOA24(dEAL CreditFreddieMacVerificationOfAssetsDealVOA24, schemaVersion float32) *CreditFreddieMacVerificationOfAssetsVOA24 {
	this := CreditFreddieMacVerificationOfAssetsVOA24{}
	this.DEAL = dEAL
	this.SchemaVersion = schemaVersion
	return &this
}

// NewCreditFreddieMacVerificationOfAssetsVOA24WithDefaults instantiates a new CreditFreddieMacVerificationOfAssetsVOA24 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacVerificationOfAssetsVOA24WithDefaults() *CreditFreddieMacVerificationOfAssetsVOA24 {
	this := CreditFreddieMacVerificationOfAssetsVOA24{}
	return &this
}

// GetDEAL returns the DEAL field value
func (o *CreditFreddieMacVerificationOfAssetsVOA24) GetDEAL() CreditFreddieMacVerificationOfAssetsDealVOA24 {
	if o == nil {
		var ret CreditFreddieMacVerificationOfAssetsDealVOA24
		return ret
	}

	return o.DEAL
}

// GetDEALOk returns a tuple with the DEAL field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacVerificationOfAssetsVOA24) GetDEALOk() (*CreditFreddieMacVerificationOfAssetsDealVOA24, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DEAL, true
}

// SetDEAL sets field value
func (o *CreditFreddieMacVerificationOfAssetsVOA24) SetDEAL(v CreditFreddieMacVerificationOfAssetsDealVOA24) {
	o.DEAL = v
}

// GetSchemaVersion returns the SchemaVersion field value
func (o *CreditFreddieMacVerificationOfAssetsVOA24) GetSchemaVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacVerificationOfAssetsVOA24) GetSchemaVersionOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SchemaVersion, true
}

// SetSchemaVersion sets field value
func (o *CreditFreddieMacVerificationOfAssetsVOA24) SetSchemaVersion(v float32) {
	o.SchemaVersion = v
}

func (o CreditFreddieMacVerificationOfAssetsVOA24) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["DEAL"] = o.DEAL
	}
	if true {
		toSerialize["SchemaVersion"] = o.SchemaVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacVerificationOfAssetsVOA24) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacVerificationOfAssetsVOA24 := _CreditFreddieMacVerificationOfAssetsVOA24{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacVerificationOfAssetsVOA24); err == nil {
		*o = CreditFreddieMacVerificationOfAssetsVOA24(varCreditFreddieMacVerificationOfAssetsVOA24)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "DEAL")
		delete(additionalProperties, "SchemaVersion")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieMacVerificationOfAssetsVOA24 struct {
	value *CreditFreddieMacVerificationOfAssetsVOA24
	isSet bool
}

func (v NullableCreditFreddieMacVerificationOfAssetsVOA24) Get() *CreditFreddieMacVerificationOfAssetsVOA24 {
	return v.value
}

func (v *NullableCreditFreddieMacVerificationOfAssetsVOA24) Set(val *CreditFreddieMacVerificationOfAssetsVOA24) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacVerificationOfAssetsVOA24) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacVerificationOfAssetsVOA24) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacVerificationOfAssetsVOA24(val *CreditFreddieMacVerificationOfAssetsVOA24) *NullableCreditFreddieMacVerificationOfAssetsVOA24 {
	return &NullableCreditFreddieMacVerificationOfAssetsVOA24{value: val, isSet: true}
}

func (v NullableCreditFreddieMacVerificationOfAssetsVOA24) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacVerificationOfAssetsVOA24) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


