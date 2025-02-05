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

// CreditFreddieMacAssetVOA24 Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
type CreditFreddieMacAssetVOA24 struct {
	ASSET_DETAIL AssetDetail `json:"ASSET_DETAIL"`
	ASSET_OWNERS AssetOwners `json:"ASSET_OWNERS"`
	ASSET_HOLDER AssetHolder `json:"ASSET_HOLDER"`
	ASSET_TRANSACTIONS CreditFreddieMacAssetTransactionsVOA24 `json:"ASSET_TRANSACTIONS"`
	VALIDATION_SOURCES ValidationSources `json:"VALIDATION_SOURCES"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacAssetVOA24 CreditFreddieMacAssetVOA24

// NewCreditFreddieMacAssetVOA24 instantiates a new CreditFreddieMacAssetVOA24 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacAssetVOA24(aSSETDETAIL AssetDetail, aSSETOWNERS AssetOwners, aSSETHOLDER AssetHolder, aSSETTRANSACTIONS CreditFreddieMacAssetTransactionsVOA24, vALIDATIONSOURCES ValidationSources) *CreditFreddieMacAssetVOA24 {
	this := CreditFreddieMacAssetVOA24{}
	this.ASSET_DETAIL = aSSETDETAIL
	this.ASSET_OWNERS = aSSETOWNERS
	this.ASSET_HOLDER = aSSETHOLDER
	this.ASSET_TRANSACTIONS = aSSETTRANSACTIONS
	this.VALIDATION_SOURCES = vALIDATIONSOURCES
	return &this
}

// NewCreditFreddieMacAssetVOA24WithDefaults instantiates a new CreditFreddieMacAssetVOA24 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacAssetVOA24WithDefaults() *CreditFreddieMacAssetVOA24 {
	this := CreditFreddieMacAssetVOA24{}
	return &this
}

// GetASSET_DETAIL returns the ASSET_DETAIL field value
func (o *CreditFreddieMacAssetVOA24) GetASSET_DETAIL() AssetDetail {
	if o == nil {
		var ret AssetDetail
		return ret
	}

	return o.ASSET_DETAIL
}

// GetASSET_DETAILOk returns a tuple with the ASSET_DETAIL field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetVOA24) GetASSET_DETAILOk() (*AssetDetail, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASSET_DETAIL, true
}

// SetASSET_DETAIL sets field value
func (o *CreditFreddieMacAssetVOA24) SetASSET_DETAIL(v AssetDetail) {
	o.ASSET_DETAIL = v
}

// GetASSET_OWNERS returns the ASSET_OWNERS field value
func (o *CreditFreddieMacAssetVOA24) GetASSET_OWNERS() AssetOwners {
	if o == nil {
		var ret AssetOwners
		return ret
	}

	return o.ASSET_OWNERS
}

// GetASSET_OWNERSOk returns a tuple with the ASSET_OWNERS field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetVOA24) GetASSET_OWNERSOk() (*AssetOwners, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASSET_OWNERS, true
}

// SetASSET_OWNERS sets field value
func (o *CreditFreddieMacAssetVOA24) SetASSET_OWNERS(v AssetOwners) {
	o.ASSET_OWNERS = v
}

// GetASSET_HOLDER returns the ASSET_HOLDER field value
func (o *CreditFreddieMacAssetVOA24) GetASSET_HOLDER() AssetHolder {
	if o == nil {
		var ret AssetHolder
		return ret
	}

	return o.ASSET_HOLDER
}

// GetASSET_HOLDEROk returns a tuple with the ASSET_HOLDER field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetVOA24) GetASSET_HOLDEROk() (*AssetHolder, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASSET_HOLDER, true
}

// SetASSET_HOLDER sets field value
func (o *CreditFreddieMacAssetVOA24) SetASSET_HOLDER(v AssetHolder) {
	o.ASSET_HOLDER = v
}

// GetASSET_TRANSACTIONS returns the ASSET_TRANSACTIONS field value
func (o *CreditFreddieMacAssetVOA24) GetASSET_TRANSACTIONS() CreditFreddieMacAssetTransactionsVOA24 {
	if o == nil {
		var ret CreditFreddieMacAssetTransactionsVOA24
		return ret
	}

	return o.ASSET_TRANSACTIONS
}

// GetASSET_TRANSACTIONSOk returns a tuple with the ASSET_TRANSACTIONS field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetVOA24) GetASSET_TRANSACTIONSOk() (*CreditFreddieMacAssetTransactionsVOA24, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASSET_TRANSACTIONS, true
}

// SetASSET_TRANSACTIONS sets field value
func (o *CreditFreddieMacAssetVOA24) SetASSET_TRANSACTIONS(v CreditFreddieMacAssetTransactionsVOA24) {
	o.ASSET_TRANSACTIONS = v
}

// GetVALIDATION_SOURCES returns the VALIDATION_SOURCES field value
func (o *CreditFreddieMacAssetVOA24) GetVALIDATION_SOURCES() ValidationSources {
	if o == nil {
		var ret ValidationSources
		return ret
	}

	return o.VALIDATION_SOURCES
}

// GetVALIDATION_SOURCESOk returns a tuple with the VALIDATION_SOURCES field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetVOA24) GetVALIDATION_SOURCESOk() (*ValidationSources, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VALIDATION_SOURCES, true
}

// SetVALIDATION_SOURCES sets field value
func (o *CreditFreddieMacAssetVOA24) SetVALIDATION_SOURCES(v ValidationSources) {
	o.VALIDATION_SOURCES = v
}

func (o CreditFreddieMacAssetVOA24) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ASSET_DETAIL"] = o.ASSET_DETAIL
	}
	if true {
		toSerialize["ASSET_OWNERS"] = o.ASSET_OWNERS
	}
	if true {
		toSerialize["ASSET_HOLDER"] = o.ASSET_HOLDER
	}
	if true {
		toSerialize["ASSET_TRANSACTIONS"] = o.ASSET_TRANSACTIONS
	}
	if true {
		toSerialize["VALIDATION_SOURCES"] = o.VALIDATION_SOURCES
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacAssetVOA24) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacAssetVOA24 := _CreditFreddieMacAssetVOA24{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacAssetVOA24); err == nil {
		*o = CreditFreddieMacAssetVOA24(varCreditFreddieMacAssetVOA24)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ASSET_DETAIL")
		delete(additionalProperties, "ASSET_OWNERS")
		delete(additionalProperties, "ASSET_HOLDER")
		delete(additionalProperties, "ASSET_TRANSACTIONS")
		delete(additionalProperties, "VALIDATION_SOURCES")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieMacAssetVOA24 struct {
	value *CreditFreddieMacAssetVOA24
	isSet bool
}

func (v NullableCreditFreddieMacAssetVOA24) Get() *CreditFreddieMacAssetVOA24 {
	return v.value
}

func (v *NullableCreditFreddieMacAssetVOA24) Set(val *CreditFreddieMacAssetVOA24) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacAssetVOA24) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacAssetVOA24) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacAssetVOA24(val *CreditFreddieMacAssetVOA24) *NullableCreditFreddieMacAssetVOA24 {
	return &NullableCreditFreddieMacAssetVOA24{value: val, isSet: true}
}

func (v NullableCreditFreddieMacAssetVOA24) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacAssetVOA24) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


