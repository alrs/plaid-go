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

// AssetTransaction An object representing...
type AssetTransaction struct {
	ASSET_TRANSACTION_DETAIL AssetTransactionDetail `json:"ASSET_TRANSACTION_DETAIL"`
	// Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
	ASSET_TRANSACTION_DESCRIPTON []AssetTransactionDescription `json:"ASSET_TRANSACTION_DESCRIPTON"`
	AdditionalProperties map[string]interface{}
}

type _AssetTransaction AssetTransaction

// NewAssetTransaction instantiates a new AssetTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetTransaction(aSSETTRANSACTIONDETAIL AssetTransactionDetail, aSSETTRANSACTIONDESCRIPTON []AssetTransactionDescription) *AssetTransaction {
	this := AssetTransaction{}
	this.ASSET_TRANSACTION_DETAIL = aSSETTRANSACTIONDETAIL
	this.ASSET_TRANSACTION_DESCRIPTON = aSSETTRANSACTIONDESCRIPTON
	return &this
}

// NewAssetTransactionWithDefaults instantiates a new AssetTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetTransactionWithDefaults() *AssetTransaction {
	this := AssetTransaction{}
	return &this
}

// GetASSET_TRANSACTION_DETAIL returns the ASSET_TRANSACTION_DETAIL field value
func (o *AssetTransaction) GetASSET_TRANSACTION_DETAIL() AssetTransactionDetail {
	if o == nil {
		var ret AssetTransactionDetail
		return ret
	}

	return o.ASSET_TRANSACTION_DETAIL
}

// GetASSET_TRANSACTION_DETAILOk returns a tuple with the ASSET_TRANSACTION_DETAIL field value
// and a boolean to check if the value has been set.
func (o *AssetTransaction) GetASSET_TRANSACTION_DETAILOk() (*AssetTransactionDetail, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASSET_TRANSACTION_DETAIL, true
}

// SetASSET_TRANSACTION_DETAIL sets field value
func (o *AssetTransaction) SetASSET_TRANSACTION_DETAIL(v AssetTransactionDetail) {
	o.ASSET_TRANSACTION_DETAIL = v
}

// GetASSET_TRANSACTION_DESCRIPTON returns the ASSET_TRANSACTION_DESCRIPTON field value
func (o *AssetTransaction) GetASSET_TRANSACTION_DESCRIPTON() []AssetTransactionDescription {
	if o == nil {
		var ret []AssetTransactionDescription
		return ret
	}

	return o.ASSET_TRANSACTION_DESCRIPTON
}

// GetASSET_TRANSACTION_DESCRIPTONOk returns a tuple with the ASSET_TRANSACTION_DESCRIPTON field value
// and a boolean to check if the value has been set.
func (o *AssetTransaction) GetASSET_TRANSACTION_DESCRIPTONOk() (*[]AssetTransactionDescription, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASSET_TRANSACTION_DESCRIPTON, true
}

// SetASSET_TRANSACTION_DESCRIPTON sets field value
func (o *AssetTransaction) SetASSET_TRANSACTION_DESCRIPTON(v []AssetTransactionDescription) {
	o.ASSET_TRANSACTION_DESCRIPTON = v
}

func (o AssetTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ASSET_TRANSACTION_DETAIL"] = o.ASSET_TRANSACTION_DETAIL
	}
	if true {
		toSerialize["ASSET_TRANSACTION_DESCRIPTON"] = o.ASSET_TRANSACTION_DESCRIPTON
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetTransaction) UnmarshalJSON(bytes []byte) (err error) {
	varAssetTransaction := _AssetTransaction{}

	if err = json.Unmarshal(bytes, &varAssetTransaction); err == nil {
		*o = AssetTransaction(varAssetTransaction)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ASSET_TRANSACTION_DETAIL")
		delete(additionalProperties, "ASSET_TRANSACTION_DESCRIPTON")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetTransaction struct {
	value *AssetTransaction
	isSet bool
}

func (v NullableAssetTransaction) Get() *AssetTransaction {
	return v.value
}

func (v *NullableAssetTransaction) Set(val *AssetTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetTransaction(val *AssetTransaction) *NullableAssetTransaction {
	return &NullableAssetTransaction{value: val, isSet: true}
}

func (v NullableAssetTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


