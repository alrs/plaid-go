/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.181.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// BankTransferCreateRequest Defines the request schema for `/bank_transfer/create`
type BankTransferCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// A random key provided by the client, per unique bank transfer. Maximum of 50 characters.  The API supports idempotency for safely retrying requests without accidentally performing the same operation twice. For example, if a request to create a bank transfer fails due to a network connection error, you can retry the request with the same idempotency key to guarantee that only a single bank transfer is created.
	IdempotencyKey string `json:"idempotency_key"`
	// The Plaid `access_token` for the account that will be debited or credited.
	AccessToken string `json:"access_token"`
	// The Plaid `account_id` for the account that will be debited or credited.
	AccountId string `json:"account_id"`
	Type BankTransferType `json:"type"`
	Network BankTransferNetwork `json:"network"`
	// The amount of the bank transfer (decimal string with two digits of precision e.g. \"10.00\").
	Amount string `json:"amount"`
	// The currency of the transfer amount – should be set to \"USD\".
	IsoCurrencyCode string `json:"iso_currency_code"`
	// The transfer description. Maximum of 10 characters.
	Description string `json:"description"`
	AchClass *ACHClass `json:"ach_class,omitempty"`
	User BankTransferUser `json:"user"`
	// An arbitrary string provided by the client for storage with the bank transfer. May be up to 100 characters.
	CustomTag NullableString `json:"custom_tag,omitempty"`
	// The Metadata object is a mapping of client-provided string fields to any string value. The following limitations apply: - The JSON values must be Strings (no nested JSON objects allowed) - Only ASCII characters may be used - Maximum of 50 key/value pairs - Maximum key length of 40 characters - Maximum value length of 500 characters 
	Metadata map[string]string `json:"metadata,omitempty"`
	// Plaid’s unique identifier for the origination account for this transfer. If you have more than one origination account, this value must be specified. Otherwise, this field should be left blank.
	OriginationAccountId NullableString `json:"origination_account_id,omitempty"`
}

// NewBankTransferCreateRequest instantiates a new BankTransferCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferCreateRequest(idempotencyKey string, accessToken string, accountId string, type_ BankTransferType, network BankTransferNetwork, amount string, isoCurrencyCode string, description string, user BankTransferUser) *BankTransferCreateRequest {
	this := BankTransferCreateRequest{}
	this.IdempotencyKey = idempotencyKey
	this.AccessToken = accessToken
	this.AccountId = accountId
	this.Type = type_
	this.Network = network
	this.Amount = amount
	this.IsoCurrencyCode = isoCurrencyCode
	this.Description = description
	this.User = user
	return &this
}

// NewBankTransferCreateRequestWithDefaults instantiates a new BankTransferCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferCreateRequestWithDefaults() *BankTransferCreateRequest {
	this := BankTransferCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *BankTransferCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *BankTransferCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *BankTransferCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *BankTransferCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *BankTransferCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *BankTransferCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetIdempotencyKey returns the IdempotencyKey field value
func (o *BankTransferCreateRequest) GetIdempotencyKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value
// and a boolean to check if the value has been set.
func (o *BankTransferCreateRequest) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IdempotencyKey, true
}

// SetIdempotencyKey sets field value
func (o *BankTransferCreateRequest) SetIdempotencyKey(v string) {
	o.IdempotencyKey = v
}

// GetAccessToken returns the AccessToken field value
func (o *BankTransferCreateRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *BankTransferCreateRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *BankTransferCreateRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetAccountId returns the AccountId field value
func (o *BankTransferCreateRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *BankTransferCreateRequest) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *BankTransferCreateRequest) SetAccountId(v string) {
	o.AccountId = v
}

// GetType returns the Type field value
func (o *BankTransferCreateRequest) GetType() BankTransferType {
	if o == nil {
		var ret BankTransferType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BankTransferCreateRequest) GetTypeOk() (*BankTransferType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BankTransferCreateRequest) SetType(v BankTransferType) {
	o.Type = v
}

// GetNetwork returns the Network field value
func (o *BankTransferCreateRequest) GetNetwork() BankTransferNetwork {
	if o == nil {
		var ret BankTransferNetwork
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *BankTransferCreateRequest) GetNetworkOk() (*BankTransferNetwork, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *BankTransferCreateRequest) SetNetwork(v BankTransferNetwork) {
	o.Network = v
}

// GetAmount returns the Amount field value
func (o *BankTransferCreateRequest) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *BankTransferCreateRequest) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *BankTransferCreateRequest) SetAmount(v string) {
	o.Amount = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
func (o *BankTransferCreateRequest) GetIsoCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *BankTransferCreateRequest) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsoCurrencyCode, true
}

// SetIsoCurrencyCode sets field value
func (o *BankTransferCreateRequest) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = v
}

// GetDescription returns the Description field value
func (o *BankTransferCreateRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *BankTransferCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *BankTransferCreateRequest) SetDescription(v string) {
	o.Description = v
}

// GetAchClass returns the AchClass field value if set, zero value otherwise.
func (o *BankTransferCreateRequest) GetAchClass() ACHClass {
	if o == nil || o.AchClass == nil {
		var ret ACHClass
		return ret
	}
	return *o.AchClass
}

// GetAchClassOk returns a tuple with the AchClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferCreateRequest) GetAchClassOk() (*ACHClass, bool) {
	if o == nil || o.AchClass == nil {
		return nil, false
	}
	return o.AchClass, true
}

// HasAchClass returns a boolean if a field has been set.
func (o *BankTransferCreateRequest) HasAchClass() bool {
	if o != nil && o.AchClass != nil {
		return true
	}

	return false
}

// SetAchClass gets a reference to the given ACHClass and assigns it to the AchClass field.
func (o *BankTransferCreateRequest) SetAchClass(v ACHClass) {
	o.AchClass = &v
}

// GetUser returns the User field value
func (o *BankTransferCreateRequest) GetUser() BankTransferUser {
	if o == nil {
		var ret BankTransferUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *BankTransferCreateRequest) GetUserOk() (*BankTransferUser, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *BankTransferCreateRequest) SetUser(v BankTransferUser) {
	o.User = v
}

// GetCustomTag returns the CustomTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankTransferCreateRequest) GetCustomTag() string {
	if o == nil || o.CustomTag.Get() == nil {
		var ret string
		return ret
	}
	return *o.CustomTag.Get()
}

// GetCustomTagOk returns a tuple with the CustomTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferCreateRequest) GetCustomTagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CustomTag.Get(), o.CustomTag.IsSet()
}

// HasCustomTag returns a boolean if a field has been set.
func (o *BankTransferCreateRequest) HasCustomTag() bool {
	if o != nil && o.CustomTag.IsSet() {
		return true
	}

	return false
}

// SetCustomTag gets a reference to the given NullableString and assigns it to the CustomTag field.
func (o *BankTransferCreateRequest) SetCustomTag(v string) {
	o.CustomTag.Set(&v)
}
// SetCustomTagNil sets the value for CustomTag to be an explicit nil
func (o *BankTransferCreateRequest) SetCustomTagNil() {
	o.CustomTag.Set(nil)
}

// UnsetCustomTag ensures that no value is present for CustomTag, not even an explicit nil
func (o *BankTransferCreateRequest) UnsetCustomTag() {
	o.CustomTag.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankTransferCreateRequest) GetMetadata() map[string]string {
	if o == nil  {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferCreateRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *BankTransferCreateRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *BankTransferCreateRequest) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetOriginationAccountId returns the OriginationAccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankTransferCreateRequest) GetOriginationAccountId() string {
	if o == nil || o.OriginationAccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OriginationAccountId.Get()
}

// GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferCreateRequest) GetOriginationAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginationAccountId.Get(), o.OriginationAccountId.IsSet()
}

// HasOriginationAccountId returns a boolean if a field has been set.
func (o *BankTransferCreateRequest) HasOriginationAccountId() bool {
	if o != nil && o.OriginationAccountId.IsSet() {
		return true
	}

	return false
}

// SetOriginationAccountId gets a reference to the given NullableString and assigns it to the OriginationAccountId field.
func (o *BankTransferCreateRequest) SetOriginationAccountId(v string) {
	o.OriginationAccountId.Set(&v)
}
// SetOriginationAccountIdNil sets the value for OriginationAccountId to be an explicit nil
func (o *BankTransferCreateRequest) SetOriginationAccountIdNil() {
	o.OriginationAccountId.Set(nil)
}

// UnsetOriginationAccountId ensures that no value is present for OriginationAccountId, not even an explicit nil
func (o *BankTransferCreateRequest) UnsetOriginationAccountId() {
	o.OriginationAccountId.Unset()
}

func (o BankTransferCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["idempotency_key"] = o.IdempotencyKey
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.AchClass != nil {
		toSerialize["ach_class"] = o.AchClass
	}
	if true {
		toSerialize["user"] = o.User
	}
	if o.CustomTag.IsSet() {
		toSerialize["custom_tag"] = o.CustomTag.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.OriginationAccountId.IsSet() {
		toSerialize["origination_account_id"] = o.OriginationAccountId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableBankTransferCreateRequest struct {
	value *BankTransferCreateRequest
	isSet bool
}

func (v NullableBankTransferCreateRequest) Get() *BankTransferCreateRequest {
	return v.value
}

func (v *NullableBankTransferCreateRequest) Set(val *BankTransferCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferCreateRequest(val *BankTransferCreateRequest) *NullableBankTransferCreateRequest {
	return &NullableBankTransferCreateRequest{value: val, isSet: true}
}

func (v NullableBankTransferCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


