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

// AccountAssetsAllOf struct for AccountAssetsAllOf
type AccountAssetsAllOf struct {
	// The duration of transaction history available for this Item, typically defined as the time since the date of the earliest transaction in that account. Only returned by Assets endpoints.
	DaysAvailable *float32 `json:"days_available,omitempty"`
	// Transaction history associated with the account. Only returned by Assets endpoints. Transaction history returned by endpoints such as `/transactions/get` or `/investments/transactions/get` will be returned in the top-level `transactions` field instead.
	Transactions *[]AssetReportTransaction `json:"transactions,omitempty"`
	// Data returned by the financial institution about the account owner or owners. Only returned by Identity or Assets endpoints. For business accounts, the name reported may be either the name of the individual or the name of the business, depending on the institution. Multiple owners on a single account will be represented in the same `owner` object, not in multiple owner objects within the array. In API versions 2018-05-22 and earlier, the `owners` object is not returned, and instead identity information is returned in the top level `identity` object. For more details, see [Plaid API versioning](https://plaid.com/docs/api/versioning/#version-2019-05-29)
	Owners *[]Owner `json:"owners,omitempty"`
	OwnershipType NullableOwnershipType `json:"ownership_type,omitempty"`
	// Calculated data about the historical balances on the account. Only returned by Assets endpoints and currently not supported by `brokerage` or `investment` accounts.
	HistoricalBalances *[]HistoricalBalance `json:"historical_balances,omitempty"`
}

// NewAccountAssetsAllOf instantiates a new AccountAssetsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountAssetsAllOf() *AccountAssetsAllOf {
	this := AccountAssetsAllOf{}
	return &this
}

// NewAccountAssetsAllOfWithDefaults instantiates a new AccountAssetsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountAssetsAllOfWithDefaults() *AccountAssetsAllOf {
	this := AccountAssetsAllOf{}
	return &this
}

// GetDaysAvailable returns the DaysAvailable field value if set, zero value otherwise.
func (o *AccountAssetsAllOf) GetDaysAvailable() float32 {
	if o == nil || o.DaysAvailable == nil {
		var ret float32
		return ret
	}
	return *o.DaysAvailable
}

// GetDaysAvailableOk returns a tuple with the DaysAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAssetsAllOf) GetDaysAvailableOk() (*float32, bool) {
	if o == nil || o.DaysAvailable == nil {
		return nil, false
	}
	return o.DaysAvailable, true
}

// HasDaysAvailable returns a boolean if a field has been set.
func (o *AccountAssetsAllOf) HasDaysAvailable() bool {
	if o != nil && o.DaysAvailable != nil {
		return true
	}

	return false
}

// SetDaysAvailable gets a reference to the given float32 and assigns it to the DaysAvailable field.
func (o *AccountAssetsAllOf) SetDaysAvailable(v float32) {
	o.DaysAvailable = &v
}

// GetTransactions returns the Transactions field value if set, zero value otherwise.
func (o *AccountAssetsAllOf) GetTransactions() []AssetReportTransaction {
	if o == nil || o.Transactions == nil {
		var ret []AssetReportTransaction
		return ret
	}
	return *o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAssetsAllOf) GetTransactionsOk() (*[]AssetReportTransaction, bool) {
	if o == nil || o.Transactions == nil {
		return nil, false
	}
	return o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *AccountAssetsAllOf) HasTransactions() bool {
	if o != nil && o.Transactions != nil {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given []AssetReportTransaction and assigns it to the Transactions field.
func (o *AccountAssetsAllOf) SetTransactions(v []AssetReportTransaction) {
	o.Transactions = &v
}

// GetOwners returns the Owners field value if set, zero value otherwise.
func (o *AccountAssetsAllOf) GetOwners() []Owner {
	if o == nil || o.Owners == nil {
		var ret []Owner
		return ret
	}
	return *o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAssetsAllOf) GetOwnersOk() (*[]Owner, bool) {
	if o == nil || o.Owners == nil {
		return nil, false
	}
	return o.Owners, true
}

// HasOwners returns a boolean if a field has been set.
func (o *AccountAssetsAllOf) HasOwners() bool {
	if o != nil && o.Owners != nil {
		return true
	}

	return false
}

// SetOwners gets a reference to the given []Owner and assigns it to the Owners field.
func (o *AccountAssetsAllOf) SetOwners(v []Owner) {
	o.Owners = &v
}

// GetOwnershipType returns the OwnershipType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountAssetsAllOf) GetOwnershipType() OwnershipType {
	if o == nil || o.OwnershipType.Get() == nil {
		var ret OwnershipType
		return ret
	}
	return *o.OwnershipType.Get()
}

// GetOwnershipTypeOk returns a tuple with the OwnershipType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountAssetsAllOf) GetOwnershipTypeOk() (*OwnershipType, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OwnershipType.Get(), o.OwnershipType.IsSet()
}

// HasOwnershipType returns a boolean if a field has been set.
func (o *AccountAssetsAllOf) HasOwnershipType() bool {
	if o != nil && o.OwnershipType.IsSet() {
		return true
	}

	return false
}

// SetOwnershipType gets a reference to the given NullableOwnershipType and assigns it to the OwnershipType field.
func (o *AccountAssetsAllOf) SetOwnershipType(v OwnershipType) {
	o.OwnershipType.Set(&v)
}
// SetOwnershipTypeNil sets the value for OwnershipType to be an explicit nil
func (o *AccountAssetsAllOf) SetOwnershipTypeNil() {
	o.OwnershipType.Set(nil)
}

// UnsetOwnershipType ensures that no value is present for OwnershipType, not even an explicit nil
func (o *AccountAssetsAllOf) UnsetOwnershipType() {
	o.OwnershipType.Unset()
}

// GetHistoricalBalances returns the HistoricalBalances field value if set, zero value otherwise.
func (o *AccountAssetsAllOf) GetHistoricalBalances() []HistoricalBalance {
	if o == nil || o.HistoricalBalances == nil {
		var ret []HistoricalBalance
		return ret
	}
	return *o.HistoricalBalances
}

// GetHistoricalBalancesOk returns a tuple with the HistoricalBalances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAssetsAllOf) GetHistoricalBalancesOk() (*[]HistoricalBalance, bool) {
	if o == nil || o.HistoricalBalances == nil {
		return nil, false
	}
	return o.HistoricalBalances, true
}

// HasHistoricalBalances returns a boolean if a field has been set.
func (o *AccountAssetsAllOf) HasHistoricalBalances() bool {
	if o != nil && o.HistoricalBalances != nil {
		return true
	}

	return false
}

// SetHistoricalBalances gets a reference to the given []HistoricalBalance and assigns it to the HistoricalBalances field.
func (o *AccountAssetsAllOf) SetHistoricalBalances(v []HistoricalBalance) {
	o.HistoricalBalances = &v
}

func (o AccountAssetsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DaysAvailable != nil {
		toSerialize["days_available"] = o.DaysAvailable
	}
	if o.Transactions != nil {
		toSerialize["transactions"] = o.Transactions
	}
	if o.Owners != nil {
		toSerialize["owners"] = o.Owners
	}
	if o.OwnershipType.IsSet() {
		toSerialize["ownership_type"] = o.OwnershipType.Get()
	}
	if o.HistoricalBalances != nil {
		toSerialize["historical_balances"] = o.HistoricalBalances
	}
	return json.Marshal(toSerialize)
}

type NullableAccountAssetsAllOf struct {
	value *AccountAssetsAllOf
	isSet bool
}

func (v NullableAccountAssetsAllOf) Get() *AccountAssetsAllOf {
	return v.value
}

func (v *NullableAccountAssetsAllOf) Set(val *AccountAssetsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAssetsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAssetsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAssetsAllOf(val *AccountAssetsAllOf) *NullableAccountAssetsAllOf {
	return &NullableAccountAssetsAllOf{value: val, isSet: true}
}

func (v NullableAccountAssetsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAssetsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


