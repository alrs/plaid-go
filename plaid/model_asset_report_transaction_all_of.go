/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.40.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// AssetReportTransactionAllOf struct for AssetReportTransactionAllOf
type AssetReportTransactionAllOf struct {
	// The date on which the transaction took place, in IS0 8601 format.
	DateTransacted NullableString `json:"date_transacted,omitempty"`
}

// NewAssetReportTransactionAllOf instantiates a new AssetReportTransactionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportTransactionAllOf() *AssetReportTransactionAllOf {
	this := AssetReportTransactionAllOf{}
	return &this
}

// NewAssetReportTransactionAllOfWithDefaults instantiates a new AssetReportTransactionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportTransactionAllOfWithDefaults() *AssetReportTransactionAllOf {
	this := AssetReportTransactionAllOf{}
	return &this
}

// GetDateTransacted returns the DateTransacted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetReportTransactionAllOf) GetDateTransacted() string {
	if o == nil || o.DateTransacted.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateTransacted.Get()
}

// GetDateTransactedOk returns a tuple with the DateTransacted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetReportTransactionAllOf) GetDateTransactedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DateTransacted.Get(), o.DateTransacted.IsSet()
}

// HasDateTransacted returns a boolean if a field has been set.
func (o *AssetReportTransactionAllOf) HasDateTransacted() bool {
	if o != nil && o.DateTransacted.IsSet() {
		return true
	}

	return false
}

// SetDateTransacted gets a reference to the given NullableString and assigns it to the DateTransacted field.
func (o *AssetReportTransactionAllOf) SetDateTransacted(v string) {
	o.DateTransacted.Set(&v)
}
// SetDateTransactedNil sets the value for DateTransacted to be an explicit nil
func (o *AssetReportTransactionAllOf) SetDateTransactedNil() {
	o.DateTransacted.Set(nil)
}

// UnsetDateTransacted ensures that no value is present for DateTransacted, not even an explicit nil
func (o *AssetReportTransactionAllOf) UnsetDateTransacted() {
	o.DateTransacted.Unset()
}

func (o AssetReportTransactionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DateTransacted.IsSet() {
		toSerialize["date_transacted"] = o.DateTransacted.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAssetReportTransactionAllOf struct {
	value *AssetReportTransactionAllOf
	isSet bool
}

func (v NullableAssetReportTransactionAllOf) Get() *AssetReportTransactionAllOf {
	return v.value
}

func (v *NullableAssetReportTransactionAllOf) Set(val *AssetReportTransactionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportTransactionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportTransactionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportTransactionAllOf(val *AssetReportTransactionAllOf) *NullableAssetReportTransactionAllOf {
	return &NullableAssetReportTransactionAllOf{value: val, isSet: true}
}

func (v NullableAssetReportTransactionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportTransactionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


