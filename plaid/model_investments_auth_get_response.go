/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.391.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// InvestmentsAuthGetResponse InvestmentsAuthGetResponse defines the response schema for `/investments/auth/get`
type InvestmentsAuthGetResponse struct {
	// The accounts for which data is being retrieved
	Accounts []AccountBase `json:"accounts"`
	// The holdings belonging to investment accounts associated with the Item. Details of the securities in the holdings are provided in the `securities` field. 
	Holdings []Holding `json:"holdings"`
	// Objects describing the securities held in the accounts associated with the Item. 
	Securities []Security `json:"securities"`
	// Information about the account owners for the accounts associated with the Item. 
	Owners []InvestmentsAuthOwner `json:"owners"`
	Numbers InvestmentsAuthGetNumbers `json:"numbers"`
	Item Item `json:"item"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _InvestmentsAuthGetResponse InvestmentsAuthGetResponse

// NewInvestmentsAuthGetResponse instantiates a new InvestmentsAuthGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvestmentsAuthGetResponse(accounts []AccountBase, holdings []Holding, securities []Security, owners []InvestmentsAuthOwner, numbers InvestmentsAuthGetNumbers, item Item, requestId string) *InvestmentsAuthGetResponse {
	this := InvestmentsAuthGetResponse{}
	this.Accounts = accounts
	this.Holdings = holdings
	this.Securities = securities
	this.Owners = owners
	this.Numbers = numbers
	this.Item = item
	this.RequestId = requestId
	return &this
}

// NewInvestmentsAuthGetResponseWithDefaults instantiates a new InvestmentsAuthGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvestmentsAuthGetResponseWithDefaults() *InvestmentsAuthGetResponse {
	this := InvestmentsAuthGetResponse{}
	return &this
}

// GetAccounts returns the Accounts field value
func (o *InvestmentsAuthGetResponse) GetAccounts() []AccountBase {
	if o == nil {
		var ret []AccountBase
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *InvestmentsAuthGetResponse) GetAccountsOk() (*[]AccountBase, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Accounts, true
}

// SetAccounts sets field value
func (o *InvestmentsAuthGetResponse) SetAccounts(v []AccountBase) {
	o.Accounts = v
}

// GetHoldings returns the Holdings field value
func (o *InvestmentsAuthGetResponse) GetHoldings() []Holding {
	if o == nil {
		var ret []Holding
		return ret
	}

	return o.Holdings
}

// GetHoldingsOk returns a tuple with the Holdings field value
// and a boolean to check if the value has been set.
func (o *InvestmentsAuthGetResponse) GetHoldingsOk() (*[]Holding, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Holdings, true
}

// SetHoldings sets field value
func (o *InvestmentsAuthGetResponse) SetHoldings(v []Holding) {
	o.Holdings = v
}

// GetSecurities returns the Securities field value
func (o *InvestmentsAuthGetResponse) GetSecurities() []Security {
	if o == nil {
		var ret []Security
		return ret
	}

	return o.Securities
}

// GetSecuritiesOk returns a tuple with the Securities field value
// and a boolean to check if the value has been set.
func (o *InvestmentsAuthGetResponse) GetSecuritiesOk() (*[]Security, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Securities, true
}

// SetSecurities sets field value
func (o *InvestmentsAuthGetResponse) SetSecurities(v []Security) {
	o.Securities = v
}

// GetOwners returns the Owners field value
func (o *InvestmentsAuthGetResponse) GetOwners() []InvestmentsAuthOwner {
	if o == nil {
		var ret []InvestmentsAuthOwner
		return ret
	}

	return o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value
// and a boolean to check if the value has been set.
func (o *InvestmentsAuthGetResponse) GetOwnersOk() (*[]InvestmentsAuthOwner, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Owners, true
}

// SetOwners sets field value
func (o *InvestmentsAuthGetResponse) SetOwners(v []InvestmentsAuthOwner) {
	o.Owners = v
}

// GetNumbers returns the Numbers field value
func (o *InvestmentsAuthGetResponse) GetNumbers() InvestmentsAuthGetNumbers {
	if o == nil {
		var ret InvestmentsAuthGetNumbers
		return ret
	}

	return o.Numbers
}

// GetNumbersOk returns a tuple with the Numbers field value
// and a boolean to check if the value has been set.
func (o *InvestmentsAuthGetResponse) GetNumbersOk() (*InvestmentsAuthGetNumbers, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Numbers, true
}

// SetNumbers sets field value
func (o *InvestmentsAuthGetResponse) SetNumbers(v InvestmentsAuthGetNumbers) {
	o.Numbers = v
}

// GetItem returns the Item field value
func (o *InvestmentsAuthGetResponse) GetItem() Item {
	if o == nil {
		var ret Item
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *InvestmentsAuthGetResponse) GetItemOk() (*Item, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *InvestmentsAuthGetResponse) SetItem(v Item) {
	o.Item = v
}

// GetRequestId returns the RequestId field value
func (o *InvestmentsAuthGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *InvestmentsAuthGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *InvestmentsAuthGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o InvestmentsAuthGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accounts"] = o.Accounts
	}
	if true {
		toSerialize["holdings"] = o.Holdings
	}
	if true {
		toSerialize["securities"] = o.Securities
	}
	if true {
		toSerialize["owners"] = o.Owners
	}
	if true {
		toSerialize["numbers"] = o.Numbers
	}
	if true {
		toSerialize["item"] = o.Item
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InvestmentsAuthGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varInvestmentsAuthGetResponse := _InvestmentsAuthGetResponse{}

	if err = json.Unmarshal(bytes, &varInvestmentsAuthGetResponse); err == nil {
		*o = InvestmentsAuthGetResponse(varInvestmentsAuthGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accounts")
		delete(additionalProperties, "holdings")
		delete(additionalProperties, "securities")
		delete(additionalProperties, "owners")
		delete(additionalProperties, "numbers")
		delete(additionalProperties, "item")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvestmentsAuthGetResponse struct {
	value *InvestmentsAuthGetResponse
	isSet bool
}

func (v NullableInvestmentsAuthGetResponse) Get() *InvestmentsAuthGetResponse {
	return v.value
}

func (v *NullableInvestmentsAuthGetResponse) Set(val *InvestmentsAuthGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInvestmentsAuthGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInvestmentsAuthGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvestmentsAuthGetResponse(val *InvestmentsAuthGetResponse) *NullableInvestmentsAuthGetResponse {
	return &NullableInvestmentsAuthGetResponse{value: val, isSet: true}
}

func (v NullableInvestmentsAuthGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvestmentsAuthGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


