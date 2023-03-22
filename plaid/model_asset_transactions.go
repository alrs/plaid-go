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

// AssetTransactions Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
type AssetTransactions struct {
	ASSET_TRANSACTION []AssetTransaction `json:"ASSET_TRANSACTION"`
	AdditionalProperties map[string]interface{}
}

type _AssetTransactions AssetTransactions

// NewAssetTransactions instantiates a new AssetTransactions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetTransactions(aSSETTRANSACTION []AssetTransaction) *AssetTransactions {
	this := AssetTransactions{}
	this.ASSET_TRANSACTION = aSSETTRANSACTION
	return &this
}

// NewAssetTransactionsWithDefaults instantiates a new AssetTransactions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetTransactionsWithDefaults() *AssetTransactions {
	this := AssetTransactions{}
	return &this
}

// GetASSET_TRANSACTION returns the ASSET_TRANSACTION field value
func (o *AssetTransactions) GetASSET_TRANSACTION() []AssetTransaction {
	if o == nil {
		var ret []AssetTransaction
		return ret
	}

	return o.ASSET_TRANSACTION
}

// GetASSET_TRANSACTIONOk returns a tuple with the ASSET_TRANSACTION field value
// and a boolean to check if the value has been set.
func (o *AssetTransactions) GetASSET_TRANSACTIONOk() (*[]AssetTransaction, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASSET_TRANSACTION, true
}

// SetASSET_TRANSACTION sets field value
func (o *AssetTransactions) SetASSET_TRANSACTION(v []AssetTransaction) {
	o.ASSET_TRANSACTION = v
}

func (o AssetTransactions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ASSET_TRANSACTION"] = o.ASSET_TRANSACTION
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetTransactions) UnmarshalJSON(bytes []byte) (err error) {
	varAssetTransactions := _AssetTransactions{}

	if err = json.Unmarshal(bytes, &varAssetTransactions); err == nil {
		*o = AssetTransactions(varAssetTransactions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ASSET_TRANSACTION")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetTransactions struct {
	value *AssetTransactions
	isSet bool
}

func (v NullableAssetTransactions) Get() *AssetTransactions {
	return v.value
}

func (v *NullableAssetTransactions) Set(val *AssetTransactions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetTransactions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetTransactions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetTransactions(val *AssetTransactions) *NullableAssetTransactions {
	return &NullableAssetTransactions{value: val, isSet: true}
}

func (v NullableAssetTransactions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetTransactions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


