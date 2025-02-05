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

// CreditFreddieMacAssetTransactionDetailVOE25 Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
type CreditFreddieMacAssetTransactionDetailVOE25 struct {
	// A vendor created unique Identifier.
	AssetTransactionUniqueIdentifier string `json:"AssetTransactionUniqueIdentifier"`
	// Asset Transaction Date.
	AssetTransactionDate string `json:"AssetTransactionDate"`
	// Asset Transaction Post Date.
	AssetTransactionPostDate string `json:"AssetTransactionPostDate"`
	AssetTransactionType AssetTransactionType `json:"AssetTransactionType"`
	// Populate with who did the transaction.
	AssetTransactionPaidByName NullableString `json:"AssetTransactionPaidByName"`
	// Populate with for whom the transaction is done.
	AssetTransactionPaidToName NullableString `json:"AssetTransactionPaidToName"`
	// FI Provided - examples are atm, cash, check, credit, debit, deposit, directDebit, directDeposit, dividend, fee, interest, other, payment, pointOfSale, repeatPayment, serviceCharge, transfer.
	AssetTransactionTypeAdditionalDescription NullableString `json:"AssetTransactionTypeAdditionalDescription"`
	AssetTransactionCategoryType NullableAssetTransactionCategoryType `json:"AssetTransactionCategoryType"`
	// FI provided Transaction Identifier.
	FinancialInstitutionTransactionIdentifier NullableString `json:"FinancialInstitutionTransactionIdentifier"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacAssetTransactionDetailVOE25 CreditFreddieMacAssetTransactionDetailVOE25

// NewCreditFreddieMacAssetTransactionDetailVOE25 instantiates a new CreditFreddieMacAssetTransactionDetailVOE25 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacAssetTransactionDetailVOE25(assetTransactionUniqueIdentifier string, assetTransactionDate string, assetTransactionPostDate string, assetTransactionType AssetTransactionType, assetTransactionPaidByName NullableString, assetTransactionPaidToName NullableString, assetTransactionTypeAdditionalDescription NullableString, assetTransactionCategoryType NullableAssetTransactionCategoryType, financialInstitutionTransactionIdentifier NullableString) *CreditFreddieMacAssetTransactionDetailVOE25 {
	this := CreditFreddieMacAssetTransactionDetailVOE25{}
	this.AssetTransactionUniqueIdentifier = assetTransactionUniqueIdentifier
	this.AssetTransactionDate = assetTransactionDate
	this.AssetTransactionPostDate = assetTransactionPostDate
	this.AssetTransactionType = assetTransactionType
	this.AssetTransactionPaidByName = assetTransactionPaidByName
	this.AssetTransactionPaidToName = assetTransactionPaidToName
	this.AssetTransactionTypeAdditionalDescription = assetTransactionTypeAdditionalDescription
	this.AssetTransactionCategoryType = assetTransactionCategoryType
	this.FinancialInstitutionTransactionIdentifier = financialInstitutionTransactionIdentifier
	return &this
}

// NewCreditFreddieMacAssetTransactionDetailVOE25WithDefaults instantiates a new CreditFreddieMacAssetTransactionDetailVOE25 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacAssetTransactionDetailVOE25WithDefaults() *CreditFreddieMacAssetTransactionDetailVOE25 {
	this := CreditFreddieMacAssetTransactionDetailVOE25{}
	return &this
}

// GetAssetTransactionUniqueIdentifier returns the AssetTransactionUniqueIdentifier field value
func (o *CreditFreddieMacAssetTransactionDetailVOE25) GetAssetTransactionUniqueIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetTransactionUniqueIdentifier
}

// GetAssetTransactionUniqueIdentifierOk returns a tuple with the AssetTransactionUniqueIdentifier field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetTransactionDetailVOE25) GetAssetTransactionUniqueIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetTransactionUniqueIdentifier, true
}

// SetAssetTransactionUniqueIdentifier sets field value
func (o *CreditFreddieMacAssetTransactionDetailVOE25) SetAssetTransactionUniqueIdentifier(v string) {
	o.AssetTransactionUniqueIdentifier = v
}

// GetAssetTransactionDate returns the AssetTransactionDate field value
func (o *CreditFreddieMacAssetTransactionDetailVOE25) GetAssetTransactionDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetTransactionDate
}

// GetAssetTransactionDateOk returns a tuple with the AssetTransactionDate field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetTransactionDetailVOE25) GetAssetTransactionDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetTransactionDate, true
}

// SetAssetTransactionDate sets field value
func (o *CreditFreddieMacAssetTransactionDetailVOE25) SetAssetTransactionDate(v string) {
	o.AssetTransactionDate = v
}

// GetAssetTransactionPostDate returns the AssetTransactionPostDate field value
func (o *CreditFreddieMacAssetTransactionDetailVOE25) GetAssetTransactionPostDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetTransactionPostDate
}

// GetAssetTransactionPostDateOk returns a tuple with the AssetTransactionPostDate field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetTransactionDetailVOE25) GetAssetTransactionPostDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetTransactionPostDate, true
}

// SetAssetTransactionPostDate sets field value
func (o *CreditFreddieMacAssetTransactionDetailVOE25) SetAssetTransactionPostDate(v string) {
	o.AssetTransactionPostDate = v
}

// GetAssetTransactionType returns the AssetTransactionType field value
func (o *CreditFreddieMacAssetTransactionDetailVOE25) GetAssetTransactionType() AssetTransactionType {
	if o == nil {
		var ret AssetTransactionType
		return ret
	}

	return o.AssetTransactionType
}

// GetAssetTransactionTypeOk returns a tuple with the AssetTransactionType field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetTransactionDetailVOE25) GetAssetTransactionTypeOk() (*AssetTransactionType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetTransactionType, true
}

// SetAssetTransactionType sets field value
func (o *CreditFreddieMacAssetTransactionDetailVOE25) SetAssetTransactionType(v AssetTransactionType) {
	o.AssetTransactionType = v
}

// GetAssetTransactionPaidByName returns the AssetTransactionPaidByName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditFreddieMacAssetTransactionDetailVOE25) GetAssetTransactionPaidByName() string {
	if o == nil || o.AssetTransactionPaidByName.Get() == nil {
		var ret string
		return ret
	}

	return *o.AssetTransactionPaidByName.Get()
}

// GetAssetTransactionPaidByNameOk returns a tuple with the AssetTransactionPaidByName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditFreddieMacAssetTransactionDetailVOE25) GetAssetTransactionPaidByNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AssetTransactionPaidByName.Get(), o.AssetTransactionPaidByName.IsSet()
}

// SetAssetTransactionPaidByName sets field value
func (o *CreditFreddieMacAssetTransactionDetailVOE25) SetAssetTransactionPaidByName(v string) {
	o.AssetTransactionPaidByName.Set(&v)
}

// GetAssetTransactionPaidToName returns the AssetTransactionPaidToName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditFreddieMacAssetTransactionDetailVOE25) GetAssetTransactionPaidToName() string {
	if o == nil || o.AssetTransactionPaidToName.Get() == nil {
		var ret string
		return ret
	}

	return *o.AssetTransactionPaidToName.Get()
}

// GetAssetTransactionPaidToNameOk returns a tuple with the AssetTransactionPaidToName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditFreddieMacAssetTransactionDetailVOE25) GetAssetTransactionPaidToNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AssetTransactionPaidToName.Get(), o.AssetTransactionPaidToName.IsSet()
}

// SetAssetTransactionPaidToName sets field value
func (o *CreditFreddieMacAssetTransactionDetailVOE25) SetAssetTransactionPaidToName(v string) {
	o.AssetTransactionPaidToName.Set(&v)
}

// GetAssetTransactionTypeAdditionalDescription returns the AssetTransactionTypeAdditionalDescription field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditFreddieMacAssetTransactionDetailVOE25) GetAssetTransactionTypeAdditionalDescription() string {
	if o == nil || o.AssetTransactionTypeAdditionalDescription.Get() == nil {
		var ret string
		return ret
	}

	return *o.AssetTransactionTypeAdditionalDescription.Get()
}

// GetAssetTransactionTypeAdditionalDescriptionOk returns a tuple with the AssetTransactionTypeAdditionalDescription field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditFreddieMacAssetTransactionDetailVOE25) GetAssetTransactionTypeAdditionalDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AssetTransactionTypeAdditionalDescription.Get(), o.AssetTransactionTypeAdditionalDescription.IsSet()
}

// SetAssetTransactionTypeAdditionalDescription sets field value
func (o *CreditFreddieMacAssetTransactionDetailVOE25) SetAssetTransactionTypeAdditionalDescription(v string) {
	o.AssetTransactionTypeAdditionalDescription.Set(&v)
}

// GetAssetTransactionCategoryType returns the AssetTransactionCategoryType field value
// If the value is explicit nil, the zero value for AssetTransactionCategoryType will be returned
func (o *CreditFreddieMacAssetTransactionDetailVOE25) GetAssetTransactionCategoryType() AssetTransactionCategoryType {
	if o == nil || o.AssetTransactionCategoryType.Get() == nil {
		var ret AssetTransactionCategoryType
		return ret
	}

	return *o.AssetTransactionCategoryType.Get()
}

// GetAssetTransactionCategoryTypeOk returns a tuple with the AssetTransactionCategoryType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditFreddieMacAssetTransactionDetailVOE25) GetAssetTransactionCategoryTypeOk() (*AssetTransactionCategoryType, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AssetTransactionCategoryType.Get(), o.AssetTransactionCategoryType.IsSet()
}

// SetAssetTransactionCategoryType sets field value
func (o *CreditFreddieMacAssetTransactionDetailVOE25) SetAssetTransactionCategoryType(v AssetTransactionCategoryType) {
	o.AssetTransactionCategoryType.Set(&v)
}

// GetFinancialInstitutionTransactionIdentifier returns the FinancialInstitutionTransactionIdentifier field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditFreddieMacAssetTransactionDetailVOE25) GetFinancialInstitutionTransactionIdentifier() string {
	if o == nil || o.FinancialInstitutionTransactionIdentifier.Get() == nil {
		var ret string
		return ret
	}

	return *o.FinancialInstitutionTransactionIdentifier.Get()
}

// GetFinancialInstitutionTransactionIdentifierOk returns a tuple with the FinancialInstitutionTransactionIdentifier field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditFreddieMacAssetTransactionDetailVOE25) GetFinancialInstitutionTransactionIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FinancialInstitutionTransactionIdentifier.Get(), o.FinancialInstitutionTransactionIdentifier.IsSet()
}

// SetFinancialInstitutionTransactionIdentifier sets field value
func (o *CreditFreddieMacAssetTransactionDetailVOE25) SetFinancialInstitutionTransactionIdentifier(v string) {
	o.FinancialInstitutionTransactionIdentifier.Set(&v)
}

func (o CreditFreddieMacAssetTransactionDetailVOE25) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["AssetTransactionUniqueIdentifier"] = o.AssetTransactionUniqueIdentifier
	}
	if true {
		toSerialize["AssetTransactionDate"] = o.AssetTransactionDate
	}
	if true {
		toSerialize["AssetTransactionPostDate"] = o.AssetTransactionPostDate
	}
	if true {
		toSerialize["AssetTransactionType"] = o.AssetTransactionType
	}
	if true {
		toSerialize["AssetTransactionPaidByName"] = o.AssetTransactionPaidByName.Get()
	}
	if true {
		toSerialize["AssetTransactionPaidToName"] = o.AssetTransactionPaidToName.Get()
	}
	if true {
		toSerialize["AssetTransactionTypeAdditionalDescription"] = o.AssetTransactionTypeAdditionalDescription.Get()
	}
	if true {
		toSerialize["AssetTransactionCategoryType"] = o.AssetTransactionCategoryType.Get()
	}
	if true {
		toSerialize["FinancialInstitutionTransactionIdentifier"] = o.FinancialInstitutionTransactionIdentifier.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacAssetTransactionDetailVOE25) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacAssetTransactionDetailVOE25 := _CreditFreddieMacAssetTransactionDetailVOE25{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacAssetTransactionDetailVOE25); err == nil {
		*o = CreditFreddieMacAssetTransactionDetailVOE25(varCreditFreddieMacAssetTransactionDetailVOE25)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "AssetTransactionUniqueIdentifier")
		delete(additionalProperties, "AssetTransactionDate")
		delete(additionalProperties, "AssetTransactionPostDate")
		delete(additionalProperties, "AssetTransactionType")
		delete(additionalProperties, "AssetTransactionPaidByName")
		delete(additionalProperties, "AssetTransactionPaidToName")
		delete(additionalProperties, "AssetTransactionTypeAdditionalDescription")
		delete(additionalProperties, "AssetTransactionCategoryType")
		delete(additionalProperties, "FinancialInstitutionTransactionIdentifier")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieMacAssetTransactionDetailVOE25 struct {
	value *CreditFreddieMacAssetTransactionDetailVOE25
	isSet bool
}

func (v NullableCreditFreddieMacAssetTransactionDetailVOE25) Get() *CreditFreddieMacAssetTransactionDetailVOE25 {
	return v.value
}

func (v *NullableCreditFreddieMacAssetTransactionDetailVOE25) Set(val *CreditFreddieMacAssetTransactionDetailVOE25) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacAssetTransactionDetailVOE25) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacAssetTransactionDetailVOE25) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacAssetTransactionDetailVOE25(val *CreditFreddieMacAssetTransactionDetailVOE25) *NullableCreditFreddieMacAssetTransactionDetailVOE25 {
	return &NullableCreditFreddieMacAssetTransactionDetailVOE25{value: val, isSet: true}
}

func (v NullableCreditFreddieMacAssetTransactionDetailVOE25) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacAssetTransactionDetailVOE25) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


