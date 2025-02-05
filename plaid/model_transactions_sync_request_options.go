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

// TransactionsSyncRequestOptions An optional object to be used with the request. If specified, `options` must not be `null`.
type TransactionsSyncRequestOptions struct {
	// Include the raw unparsed transaction description from the financial institution. This field is disabled by default. If you need this information in addition to the parsed data provided, contact your Plaid Account Manager or submit a [Support request](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/product-functionality).
	IncludeOriginalDescription NullableBool `json:"include_original_description,omitempty"`
	// Include the [`personal_finance_category`](https://plaid.com/docs/api/products/transactions/#transactions-sync-response-added-personal-finance-category) object in the response.  All implementations are encouraged to use set this field to `true` and to use the `personal_finance_category` instead of `category` for more meaningful and accurate categorization.  See the [`taxonomy csv file`](https://plaid.com/documents/transactions-personal-finance-category-taxonomy.csv) for a full list of personal finance categories.
	IncludePersonalFinanceCategory *bool `json:"include_personal_finance_category,omitempty"`
	// Include counterparties and extra merchant fields in the transaction. This field is disabled by default. If you need this information in addition to the parsed data provided, contact your Plaid Account Manager.
	IncludeLogoAndCounterpartyBeta *bool `json:"include_logo_and_counterparty_beta,omitempty"`
}

// NewTransactionsSyncRequestOptions instantiates a new TransactionsSyncRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsSyncRequestOptions() *TransactionsSyncRequestOptions {
	this := TransactionsSyncRequestOptions{}
	var includeOriginalDescription bool = false
	this.IncludeOriginalDescription = *NewNullableBool(&includeOriginalDescription)
	var includePersonalFinanceCategory bool = false
	this.IncludePersonalFinanceCategory = &includePersonalFinanceCategory
	var includeLogoAndCounterpartyBeta bool = false
	this.IncludeLogoAndCounterpartyBeta = &includeLogoAndCounterpartyBeta
	return &this
}

// NewTransactionsSyncRequestOptionsWithDefaults instantiates a new TransactionsSyncRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsSyncRequestOptionsWithDefaults() *TransactionsSyncRequestOptions {
	this := TransactionsSyncRequestOptions{}
	var includeOriginalDescription bool = false
	this.IncludeOriginalDescription = *NewNullableBool(&includeOriginalDescription)
	var includePersonalFinanceCategory bool = false
	this.IncludePersonalFinanceCategory = &includePersonalFinanceCategory
	var includeLogoAndCounterpartyBeta bool = false
	this.IncludeLogoAndCounterpartyBeta = &includeLogoAndCounterpartyBeta
	return &this
}

// GetIncludeOriginalDescription returns the IncludeOriginalDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionsSyncRequestOptions) GetIncludeOriginalDescription() bool {
	if o == nil || o.IncludeOriginalDescription.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IncludeOriginalDescription.Get()
}

// GetIncludeOriginalDescriptionOk returns a tuple with the IncludeOriginalDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionsSyncRequestOptions) GetIncludeOriginalDescriptionOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IncludeOriginalDescription.Get(), o.IncludeOriginalDescription.IsSet()
}

// HasIncludeOriginalDescription returns a boolean if a field has been set.
func (o *TransactionsSyncRequestOptions) HasIncludeOriginalDescription() bool {
	if o != nil && o.IncludeOriginalDescription.IsSet() {
		return true
	}

	return false
}

// SetIncludeOriginalDescription gets a reference to the given NullableBool and assigns it to the IncludeOriginalDescription field.
func (o *TransactionsSyncRequestOptions) SetIncludeOriginalDescription(v bool) {
	o.IncludeOriginalDescription.Set(&v)
}
// SetIncludeOriginalDescriptionNil sets the value for IncludeOriginalDescription to be an explicit nil
func (o *TransactionsSyncRequestOptions) SetIncludeOriginalDescriptionNil() {
	o.IncludeOriginalDescription.Set(nil)
}

// UnsetIncludeOriginalDescription ensures that no value is present for IncludeOriginalDescription, not even an explicit nil
func (o *TransactionsSyncRequestOptions) UnsetIncludeOriginalDescription() {
	o.IncludeOriginalDescription.Unset()
}

// GetIncludePersonalFinanceCategory returns the IncludePersonalFinanceCategory field value if set, zero value otherwise.
func (o *TransactionsSyncRequestOptions) GetIncludePersonalFinanceCategory() bool {
	if o == nil || o.IncludePersonalFinanceCategory == nil {
		var ret bool
		return ret
	}
	return *o.IncludePersonalFinanceCategory
}

// GetIncludePersonalFinanceCategoryOk returns a tuple with the IncludePersonalFinanceCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsSyncRequestOptions) GetIncludePersonalFinanceCategoryOk() (*bool, bool) {
	if o == nil || o.IncludePersonalFinanceCategory == nil {
		return nil, false
	}
	return o.IncludePersonalFinanceCategory, true
}

// HasIncludePersonalFinanceCategory returns a boolean if a field has been set.
func (o *TransactionsSyncRequestOptions) HasIncludePersonalFinanceCategory() bool {
	if o != nil && o.IncludePersonalFinanceCategory != nil {
		return true
	}

	return false
}

// SetIncludePersonalFinanceCategory gets a reference to the given bool and assigns it to the IncludePersonalFinanceCategory field.
func (o *TransactionsSyncRequestOptions) SetIncludePersonalFinanceCategory(v bool) {
	o.IncludePersonalFinanceCategory = &v
}

// GetIncludeLogoAndCounterpartyBeta returns the IncludeLogoAndCounterpartyBeta field value if set, zero value otherwise.
func (o *TransactionsSyncRequestOptions) GetIncludeLogoAndCounterpartyBeta() bool {
	if o == nil || o.IncludeLogoAndCounterpartyBeta == nil {
		var ret bool
		return ret
	}
	return *o.IncludeLogoAndCounterpartyBeta
}

// GetIncludeLogoAndCounterpartyBetaOk returns a tuple with the IncludeLogoAndCounterpartyBeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsSyncRequestOptions) GetIncludeLogoAndCounterpartyBetaOk() (*bool, bool) {
	if o == nil || o.IncludeLogoAndCounterpartyBeta == nil {
		return nil, false
	}
	return o.IncludeLogoAndCounterpartyBeta, true
}

// HasIncludeLogoAndCounterpartyBeta returns a boolean if a field has been set.
func (o *TransactionsSyncRequestOptions) HasIncludeLogoAndCounterpartyBeta() bool {
	if o != nil && o.IncludeLogoAndCounterpartyBeta != nil {
		return true
	}

	return false
}

// SetIncludeLogoAndCounterpartyBeta gets a reference to the given bool and assigns it to the IncludeLogoAndCounterpartyBeta field.
func (o *TransactionsSyncRequestOptions) SetIncludeLogoAndCounterpartyBeta(v bool) {
	o.IncludeLogoAndCounterpartyBeta = &v
}

func (o TransactionsSyncRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IncludeOriginalDescription.IsSet() {
		toSerialize["include_original_description"] = o.IncludeOriginalDescription.Get()
	}
	if o.IncludePersonalFinanceCategory != nil {
		toSerialize["include_personal_finance_category"] = o.IncludePersonalFinanceCategory
	}
	if o.IncludeLogoAndCounterpartyBeta != nil {
		toSerialize["include_logo_and_counterparty_beta"] = o.IncludeLogoAndCounterpartyBeta
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionsSyncRequestOptions struct {
	value *TransactionsSyncRequestOptions
	isSet bool
}

func (v NullableTransactionsSyncRequestOptions) Get() *TransactionsSyncRequestOptions {
	return v.value
}

func (v *NullableTransactionsSyncRequestOptions) Set(val *TransactionsSyncRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsSyncRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsSyncRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsSyncRequestOptions(val *TransactionsSyncRequestOptions) *NullableTransactionsSyncRequestOptions {
	return &NullableTransactionsSyncRequestOptions{value: val, isSet: true}
}

func (v NullableTransactionsSyncRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsSyncRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


