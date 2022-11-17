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

// AssetReportPDFGetRequestOptions An optional object to filter or add data to `/asset_report/get` results. If provided, must be non-`null`.
type AssetReportPDFGetRequestOptions struct {
	// The maximum integer number of days of history to include in the Asset Report.
	DaysToInclude NullableInt32 `json:"days_to_include,omitempty"`
}

// NewAssetReportPDFGetRequestOptions instantiates a new AssetReportPDFGetRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportPDFGetRequestOptions() *AssetReportPDFGetRequestOptions {
	this := AssetReportPDFGetRequestOptions{}
	return &this
}

// NewAssetReportPDFGetRequestOptionsWithDefaults instantiates a new AssetReportPDFGetRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportPDFGetRequestOptionsWithDefaults() *AssetReportPDFGetRequestOptions {
	this := AssetReportPDFGetRequestOptions{}
	return &this
}

// GetDaysToInclude returns the DaysToInclude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetReportPDFGetRequestOptions) GetDaysToInclude() int32 {
	if o == nil || o.DaysToInclude.Get() == nil {
		var ret int32
		return ret
	}
	return *o.DaysToInclude.Get()
}

// GetDaysToIncludeOk returns a tuple with the DaysToInclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetReportPDFGetRequestOptions) GetDaysToIncludeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DaysToInclude.Get(), o.DaysToInclude.IsSet()
}

// HasDaysToInclude returns a boolean if a field has been set.
func (o *AssetReportPDFGetRequestOptions) HasDaysToInclude() bool {
	if o != nil && o.DaysToInclude.IsSet() {
		return true
	}

	return false
}

// SetDaysToInclude gets a reference to the given NullableInt32 and assigns it to the DaysToInclude field.
func (o *AssetReportPDFGetRequestOptions) SetDaysToInclude(v int32) {
	o.DaysToInclude.Set(&v)
}
// SetDaysToIncludeNil sets the value for DaysToInclude to be an explicit nil
func (o *AssetReportPDFGetRequestOptions) SetDaysToIncludeNil() {
	o.DaysToInclude.Set(nil)
}

// UnsetDaysToInclude ensures that no value is present for DaysToInclude, not even an explicit nil
func (o *AssetReportPDFGetRequestOptions) UnsetDaysToInclude() {
	o.DaysToInclude.Unset()
}

func (o AssetReportPDFGetRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DaysToInclude.IsSet() {
		toSerialize["days_to_include"] = o.DaysToInclude.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAssetReportPDFGetRequestOptions struct {
	value *AssetReportPDFGetRequestOptions
	isSet bool
}

func (v NullableAssetReportPDFGetRequestOptions) Get() *AssetReportPDFGetRequestOptions {
	return v.value
}

func (v *NullableAssetReportPDFGetRequestOptions) Set(val *AssetReportPDFGetRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportPDFGetRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportPDFGetRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportPDFGetRequestOptions(val *AssetReportPDFGetRequestOptions) *NullableAssetReportPDFGetRequestOptions {
	return &NullableAssetReportPDFGetRequestOptions{value: val, isSet: true}
}

func (v NullableAssetReportPDFGetRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportPDFGetRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


