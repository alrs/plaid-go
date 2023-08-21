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

// TransactionsEnrichRequestOptions An optional object to be used with the request.
type TransactionsEnrichRequestOptions struct {
	// Include `legacy_category` and `legacy_category_id` in the response (in addition to the default `personal_finance_category`).  Categories are based on Plaid's legacy taxonomy. For a full list of legacy categories, see [`/categories/get`](https://plaid.com/docs/api/products/transactions/#categoriesget).
	IncludeLegacyCategory *bool `json:"include_legacy_category,omitempty"`
}

// NewTransactionsEnrichRequestOptions instantiates a new TransactionsEnrichRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsEnrichRequestOptions() *TransactionsEnrichRequestOptions {
	this := TransactionsEnrichRequestOptions{}
	var includeLegacyCategory bool = false
	this.IncludeLegacyCategory = &includeLegacyCategory
	return &this
}

// NewTransactionsEnrichRequestOptionsWithDefaults instantiates a new TransactionsEnrichRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsEnrichRequestOptionsWithDefaults() *TransactionsEnrichRequestOptions {
	this := TransactionsEnrichRequestOptions{}
	var includeLegacyCategory bool = false
	this.IncludeLegacyCategory = &includeLegacyCategory
	return &this
}

// GetIncludeLegacyCategory returns the IncludeLegacyCategory field value if set, zero value otherwise.
func (o *TransactionsEnrichRequestOptions) GetIncludeLegacyCategory() bool {
	if o == nil || o.IncludeLegacyCategory == nil {
		var ret bool
		return ret
	}
	return *o.IncludeLegacyCategory
}

// GetIncludeLegacyCategoryOk returns a tuple with the IncludeLegacyCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsEnrichRequestOptions) GetIncludeLegacyCategoryOk() (*bool, bool) {
	if o == nil || o.IncludeLegacyCategory == nil {
		return nil, false
	}
	return o.IncludeLegacyCategory, true
}

// HasIncludeLegacyCategory returns a boolean if a field has been set.
func (o *TransactionsEnrichRequestOptions) HasIncludeLegacyCategory() bool {
	if o != nil && o.IncludeLegacyCategory != nil {
		return true
	}

	return false
}

// SetIncludeLegacyCategory gets a reference to the given bool and assigns it to the IncludeLegacyCategory field.
func (o *TransactionsEnrichRequestOptions) SetIncludeLegacyCategory(v bool) {
	o.IncludeLegacyCategory = &v
}

func (o TransactionsEnrichRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IncludeLegacyCategory != nil {
		toSerialize["include_legacy_category"] = o.IncludeLegacyCategory
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionsEnrichRequestOptions struct {
	value *TransactionsEnrichRequestOptions
	isSet bool
}

func (v NullableTransactionsEnrichRequestOptions) Get() *TransactionsEnrichRequestOptions {
	return v.value
}

func (v *NullableTransactionsEnrichRequestOptions) Set(val *TransactionsEnrichRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsEnrichRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsEnrichRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsEnrichRequestOptions(val *TransactionsEnrichRequestOptions) *NullableTransactionsEnrichRequestOptions {
	return &NullableTransactionsEnrichRequestOptions{value: val, isSet: true}
}

func (v NullableTransactionsEnrichRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsEnrichRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


