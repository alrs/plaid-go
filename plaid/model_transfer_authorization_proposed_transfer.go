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

// TransferAuthorizationProposedTransfer Details regarding the proposed transfer.
type TransferAuthorizationProposedTransfer struct {
	AchClass *ACHClass `json:"ach_class,omitempty"`
	// The Plaid `account_id` for the account that will be debited or credited.
	AccountId *string `json:"account_id,omitempty"`
	Type TransferType `json:"type"`
	User TransferUserInResponse `json:"user"`
	// The amount of the transfer (decimal string with two digits of precision e.g. \"10.00\").
	Amount string `json:"amount"`
	// The network or rails used for the transfer.
	Network string `json:"network"`
	// Plaid's unique identifier for the origination account that was used for this transfer.
	OriginationAccountId string `json:"origination_account_id"`
	// The currency of the transfer amount. The default value is \"USD\".
	IsoCurrencyCode string `json:"iso_currency_code"`
	// The Plaid client ID that is the originator of this transfer. Only present if created on behalf of another client as a third-party sender (TPS).
	OriginatorClientId NullableString `json:"originator_client_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferAuthorizationProposedTransfer TransferAuthorizationProposedTransfer

// NewTransferAuthorizationProposedTransfer instantiates a new TransferAuthorizationProposedTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferAuthorizationProposedTransfer(type_ TransferType, user TransferUserInResponse, amount string, network string, originationAccountId string, isoCurrencyCode string, originatorClientId NullableString) *TransferAuthorizationProposedTransfer {
	this := TransferAuthorizationProposedTransfer{}
	this.Type = type_
	this.User = user
	this.Amount = amount
	this.Network = network
	this.OriginationAccountId = originationAccountId
	this.IsoCurrencyCode = isoCurrencyCode
	this.OriginatorClientId = originatorClientId
	return &this
}

// NewTransferAuthorizationProposedTransferWithDefaults instantiates a new TransferAuthorizationProposedTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferAuthorizationProposedTransferWithDefaults() *TransferAuthorizationProposedTransfer {
	this := TransferAuthorizationProposedTransfer{}
	return &this
}

// GetAchClass returns the AchClass field value if set, zero value otherwise.
func (o *TransferAuthorizationProposedTransfer) GetAchClass() ACHClass {
	if o == nil || o.AchClass == nil {
		var ret ACHClass
		return ret
	}
	return *o.AchClass
}

// GetAchClassOk returns a tuple with the AchClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationProposedTransfer) GetAchClassOk() (*ACHClass, bool) {
	if o == nil || o.AchClass == nil {
		return nil, false
	}
	return o.AchClass, true
}

// HasAchClass returns a boolean if a field has been set.
func (o *TransferAuthorizationProposedTransfer) HasAchClass() bool {
	if o != nil && o.AchClass != nil {
		return true
	}

	return false
}

// SetAchClass gets a reference to the given ACHClass and assigns it to the AchClass field.
func (o *TransferAuthorizationProposedTransfer) SetAchClass(v ACHClass) {
	o.AchClass = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *TransferAuthorizationProposedTransfer) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationProposedTransfer) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *TransferAuthorizationProposedTransfer) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *TransferAuthorizationProposedTransfer) SetAccountId(v string) {
	o.AccountId = &v
}

// GetType returns the Type field value
func (o *TransferAuthorizationProposedTransfer) GetType() TransferType {
	if o == nil {
		var ret TransferType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationProposedTransfer) GetTypeOk() (*TransferType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransferAuthorizationProposedTransfer) SetType(v TransferType) {
	o.Type = v
}

// GetUser returns the User field value
func (o *TransferAuthorizationProposedTransfer) GetUser() TransferUserInResponse {
	if o == nil {
		var ret TransferUserInResponse
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationProposedTransfer) GetUserOk() (*TransferUserInResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *TransferAuthorizationProposedTransfer) SetUser(v TransferUserInResponse) {
	o.User = v
}

// GetAmount returns the Amount field value
func (o *TransferAuthorizationProposedTransfer) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationProposedTransfer) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransferAuthorizationProposedTransfer) SetAmount(v string) {
	o.Amount = v
}

// GetNetwork returns the Network field value
func (o *TransferAuthorizationProposedTransfer) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationProposedTransfer) GetNetworkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *TransferAuthorizationProposedTransfer) SetNetwork(v string) {
	o.Network = v
}

// GetOriginationAccountId returns the OriginationAccountId field value
func (o *TransferAuthorizationProposedTransfer) GetOriginationAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginationAccountId
}

// GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationProposedTransfer) GetOriginationAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OriginationAccountId, true
}

// SetOriginationAccountId sets field value
func (o *TransferAuthorizationProposedTransfer) SetOriginationAccountId(v string) {
	o.OriginationAccountId = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
func (o *TransferAuthorizationProposedTransfer) GetIsoCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationProposedTransfer) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsoCurrencyCode, true
}

// SetIsoCurrencyCode sets field value
func (o *TransferAuthorizationProposedTransfer) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = v
}

// GetOriginatorClientId returns the OriginatorClientId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransferAuthorizationProposedTransfer) GetOriginatorClientId() string {
	if o == nil || o.OriginatorClientId.Get() == nil {
		var ret string
		return ret
	}

	return *o.OriginatorClientId.Get()
}

// GetOriginatorClientIdOk returns a tuple with the OriginatorClientId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferAuthorizationProposedTransfer) GetOriginatorClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginatorClientId.Get(), o.OriginatorClientId.IsSet()
}

// SetOriginatorClientId sets field value
func (o *TransferAuthorizationProposedTransfer) SetOriginatorClientId(v string) {
	o.OriginatorClientId.Set(&v)
}

func (o TransferAuthorizationProposedTransfer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AchClass != nil {
		toSerialize["ach_class"] = o.AchClass
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["origination_account_id"] = o.OriginationAccountId
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}
	if true {
		toSerialize["originator_client_id"] = o.OriginatorClientId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferAuthorizationProposedTransfer) UnmarshalJSON(bytes []byte) (err error) {
	varTransferAuthorizationProposedTransfer := _TransferAuthorizationProposedTransfer{}

	if err = json.Unmarshal(bytes, &varTransferAuthorizationProposedTransfer); err == nil {
		*o = TransferAuthorizationProposedTransfer(varTransferAuthorizationProposedTransfer)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ach_class")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "user")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "network")
		delete(additionalProperties, "origination_account_id")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "originator_client_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferAuthorizationProposedTransfer struct {
	value *TransferAuthorizationProposedTransfer
	isSet bool
}

func (v NullableTransferAuthorizationProposedTransfer) Get() *TransferAuthorizationProposedTransfer {
	return v.value
}

func (v *NullableTransferAuthorizationProposedTransfer) Set(val *TransferAuthorizationProposedTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferAuthorizationProposedTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferAuthorizationProposedTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferAuthorizationProposedTransfer(val *TransferAuthorizationProposedTransfer) *NullableTransferAuthorizationProposedTransfer {
	return &NullableTransferAuthorizationProposedTransfer{value: val, isSet: true}
}

func (v NullableTransferAuthorizationProposedTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferAuthorizationProposedTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


