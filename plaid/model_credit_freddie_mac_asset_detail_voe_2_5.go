/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.379.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// CreditFreddieMacAssetDetailVOE25 Details about an asset.
type CreditFreddieMacAssetDetailVOE25 struct {
	// A vendor created unique Identifier.
	AssetUniqueIdentifier string `json:"AssetUniqueIdentifier"`
	// A unique alphanumeric string identifying an asset.
	AssetAccountIdentifier string `json:"AssetAccountIdentifier"`
	// Account Report As of Date / Create Date. Format YYYY-MM-DD
	AssetAsOfDate string `json:"AssetAsOfDate"`
	// A text description that further defines the Asset. This could be used to describe the shares associated with the stocks, bonds or mutual funds, retirement funds or business owned that the borrower has disclosed (named) as an asset.
	AssetDescription NullableString `json:"AssetDescription"`
	AssetType AssetType `json:"AssetType"`
	// Additional Asset Decription some examples are Investment Tax-Deferred , Loan, 401K, 403B, Checking, Money Market, Credit Card,ROTH,529,Biller,ROLLOVER,CD,Savings,Investment Taxable, IRA, Mortgage, Line Of Credit.
	AssetTypeAdditionalDescription NullableString `json:"AssetTypeAdditionalDescription"`
	// The Number of days requested made to the Financial Institution. Example When looking for 3 months of data from the FI, pass in 90 days.
	AssetDaysRequestedCount int32 `json:"AssetDaysRequestedCount"`
	// Ownership type of the asset account.
	AssetOwnershipType NullableString `json:"AssetOwnershipType"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacAssetDetailVOE25 CreditFreddieMacAssetDetailVOE25

// NewCreditFreddieMacAssetDetailVOE25 instantiates a new CreditFreddieMacAssetDetailVOE25 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacAssetDetailVOE25(assetUniqueIdentifier string, assetAccountIdentifier string, assetAsOfDate string, assetDescription NullableString, assetType AssetType, assetTypeAdditionalDescription NullableString, assetDaysRequestedCount int32, assetOwnershipType NullableString) *CreditFreddieMacAssetDetailVOE25 {
	this := CreditFreddieMacAssetDetailVOE25{}
	this.AssetUniqueIdentifier = assetUniqueIdentifier
	this.AssetAccountIdentifier = assetAccountIdentifier
	this.AssetAsOfDate = assetAsOfDate
	this.AssetDescription = assetDescription
	this.AssetType = assetType
	this.AssetTypeAdditionalDescription = assetTypeAdditionalDescription
	this.AssetDaysRequestedCount = assetDaysRequestedCount
	this.AssetOwnershipType = assetOwnershipType
	return &this
}

// NewCreditFreddieMacAssetDetailVOE25WithDefaults instantiates a new CreditFreddieMacAssetDetailVOE25 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacAssetDetailVOE25WithDefaults() *CreditFreddieMacAssetDetailVOE25 {
	this := CreditFreddieMacAssetDetailVOE25{}
	return &this
}

// GetAssetUniqueIdentifier returns the AssetUniqueIdentifier field value
func (o *CreditFreddieMacAssetDetailVOE25) GetAssetUniqueIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetUniqueIdentifier
}

// GetAssetUniqueIdentifierOk returns a tuple with the AssetUniqueIdentifier field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetDetailVOE25) GetAssetUniqueIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetUniqueIdentifier, true
}

// SetAssetUniqueIdentifier sets field value
func (o *CreditFreddieMacAssetDetailVOE25) SetAssetUniqueIdentifier(v string) {
	o.AssetUniqueIdentifier = v
}

// GetAssetAccountIdentifier returns the AssetAccountIdentifier field value
func (o *CreditFreddieMacAssetDetailVOE25) GetAssetAccountIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetAccountIdentifier
}

// GetAssetAccountIdentifierOk returns a tuple with the AssetAccountIdentifier field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetDetailVOE25) GetAssetAccountIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetAccountIdentifier, true
}

// SetAssetAccountIdentifier sets field value
func (o *CreditFreddieMacAssetDetailVOE25) SetAssetAccountIdentifier(v string) {
	o.AssetAccountIdentifier = v
}

// GetAssetAsOfDate returns the AssetAsOfDate field value
func (o *CreditFreddieMacAssetDetailVOE25) GetAssetAsOfDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetAsOfDate
}

// GetAssetAsOfDateOk returns a tuple with the AssetAsOfDate field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetDetailVOE25) GetAssetAsOfDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetAsOfDate, true
}

// SetAssetAsOfDate sets field value
func (o *CreditFreddieMacAssetDetailVOE25) SetAssetAsOfDate(v string) {
	o.AssetAsOfDate = v
}

// GetAssetDescription returns the AssetDescription field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditFreddieMacAssetDetailVOE25) GetAssetDescription() string {
	if o == nil || o.AssetDescription.Get() == nil {
		var ret string
		return ret
	}

	return *o.AssetDescription.Get()
}

// GetAssetDescriptionOk returns a tuple with the AssetDescription field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditFreddieMacAssetDetailVOE25) GetAssetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AssetDescription.Get(), o.AssetDescription.IsSet()
}

// SetAssetDescription sets field value
func (o *CreditFreddieMacAssetDetailVOE25) SetAssetDescription(v string) {
	o.AssetDescription.Set(&v)
}

// GetAssetType returns the AssetType field value
func (o *CreditFreddieMacAssetDetailVOE25) GetAssetType() AssetType {
	if o == nil {
		var ret AssetType
		return ret
	}

	return o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetDetailVOE25) GetAssetTypeOk() (*AssetType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetType, true
}

// SetAssetType sets field value
func (o *CreditFreddieMacAssetDetailVOE25) SetAssetType(v AssetType) {
	o.AssetType = v
}

// GetAssetTypeAdditionalDescription returns the AssetTypeAdditionalDescription field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditFreddieMacAssetDetailVOE25) GetAssetTypeAdditionalDescription() string {
	if o == nil || o.AssetTypeAdditionalDescription.Get() == nil {
		var ret string
		return ret
	}

	return *o.AssetTypeAdditionalDescription.Get()
}

// GetAssetTypeAdditionalDescriptionOk returns a tuple with the AssetTypeAdditionalDescription field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditFreddieMacAssetDetailVOE25) GetAssetTypeAdditionalDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AssetTypeAdditionalDescription.Get(), o.AssetTypeAdditionalDescription.IsSet()
}

// SetAssetTypeAdditionalDescription sets field value
func (o *CreditFreddieMacAssetDetailVOE25) SetAssetTypeAdditionalDescription(v string) {
	o.AssetTypeAdditionalDescription.Set(&v)
}

// GetAssetDaysRequestedCount returns the AssetDaysRequestedCount field value
func (o *CreditFreddieMacAssetDetailVOE25) GetAssetDaysRequestedCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AssetDaysRequestedCount
}

// GetAssetDaysRequestedCountOk returns a tuple with the AssetDaysRequestedCount field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetDetailVOE25) GetAssetDaysRequestedCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetDaysRequestedCount, true
}

// SetAssetDaysRequestedCount sets field value
func (o *CreditFreddieMacAssetDetailVOE25) SetAssetDaysRequestedCount(v int32) {
	o.AssetDaysRequestedCount = v
}

// GetAssetOwnershipType returns the AssetOwnershipType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditFreddieMacAssetDetailVOE25) GetAssetOwnershipType() string {
	if o == nil || o.AssetOwnershipType.Get() == nil {
		var ret string
		return ret
	}

	return *o.AssetOwnershipType.Get()
}

// GetAssetOwnershipTypeOk returns a tuple with the AssetOwnershipType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditFreddieMacAssetDetailVOE25) GetAssetOwnershipTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AssetOwnershipType.Get(), o.AssetOwnershipType.IsSet()
}

// SetAssetOwnershipType sets field value
func (o *CreditFreddieMacAssetDetailVOE25) SetAssetOwnershipType(v string) {
	o.AssetOwnershipType.Set(&v)
}

func (o CreditFreddieMacAssetDetailVOE25) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["AssetUniqueIdentifier"] = o.AssetUniqueIdentifier
	}
	if true {
		toSerialize["AssetAccountIdentifier"] = o.AssetAccountIdentifier
	}
	if true {
		toSerialize["AssetAsOfDate"] = o.AssetAsOfDate
	}
	if true {
		toSerialize["AssetDescription"] = o.AssetDescription.Get()
	}
	if true {
		toSerialize["AssetType"] = o.AssetType
	}
	if true {
		toSerialize["AssetTypeAdditionalDescription"] = o.AssetTypeAdditionalDescription.Get()
	}
	if true {
		toSerialize["AssetDaysRequestedCount"] = o.AssetDaysRequestedCount
	}
	if true {
		toSerialize["AssetOwnershipType"] = o.AssetOwnershipType.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacAssetDetailVOE25) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacAssetDetailVOE25 := _CreditFreddieMacAssetDetailVOE25{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacAssetDetailVOE25); err == nil {
		*o = CreditFreddieMacAssetDetailVOE25(varCreditFreddieMacAssetDetailVOE25)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "AssetUniqueIdentifier")
		delete(additionalProperties, "AssetAccountIdentifier")
		delete(additionalProperties, "AssetAsOfDate")
		delete(additionalProperties, "AssetDescription")
		delete(additionalProperties, "AssetType")
		delete(additionalProperties, "AssetTypeAdditionalDescription")
		delete(additionalProperties, "AssetDaysRequestedCount")
		delete(additionalProperties, "AssetOwnershipType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieMacAssetDetailVOE25 struct {
	value *CreditFreddieMacAssetDetailVOE25
	isSet bool
}

func (v NullableCreditFreddieMacAssetDetailVOE25) Get() *CreditFreddieMacAssetDetailVOE25 {
	return v.value
}

func (v *NullableCreditFreddieMacAssetDetailVOE25) Set(val *CreditFreddieMacAssetDetailVOE25) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacAssetDetailVOE25) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacAssetDetailVOE25) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacAssetDetailVOE25(val *CreditFreddieMacAssetDetailVOE25) *NullableCreditFreddieMacAssetDetailVOE25 {
	return &NullableCreditFreddieMacAssetDetailVOE25{value: val, isSet: true}
}

func (v NullableCreditFreddieMacAssetDetailVOE25) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacAssetDetailVOE25) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


