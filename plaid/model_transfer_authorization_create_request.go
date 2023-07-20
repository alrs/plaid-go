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

// TransferAuthorizationCreateRequest Defines the request schema for `/transfer/authorization/create`
type TransferAuthorizationCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The Plaid `access_token` for the account that will be debited or credited. Required if not using `payment_profile_token`.
	AccessToken *string `json:"access_token,omitempty"`
	// The Plaid `account_id` corresponding to the end-user account that will be debited or credited. Required when creating a transfer using an `access_token`.
	AccountId *string `json:"account_id,omitempty"`
	// The id of the funding account to use, available in the Plaid Dashboard. This determines which of your business checking accounts will be credited or debited. Defaults to the account configured during onboarding.
	FundingAccountId NullableString `json:"funding_account_id,omitempty"`
	// The payment profile token associated with the Payment Profile that will be debited or credited. Required if not using `access_token`.
	PaymentProfileToken *string `json:"payment_profile_token,omitempty"`
	Type TransferType `json:"type"`
	Network TransferNetwork `json:"network"`
	// The amount of the transfer (decimal string with two digits of precision e.g. \"10.00\").
	Amount string `json:"amount"`
	AchClass *ACHClass `json:"ach_class,omitempty"`
	User TransferAuthorizationUserInRequest `json:"user"`
	Device *TransferAuthorizationDevice `json:"device,omitempty"`
	// Plaid's unique identifier for the origination account for this authorization. If not specified, the default account will be used.
	OriginationAccountId *string `json:"origination_account_id,omitempty"`
	// The currency of the transfer amount. The default value is \"USD\".
	IsoCurrencyCode *string `json:"iso_currency_code,omitempty"`
	// A random key provided by the client, per unique authorization, which expires after 48 hours. Maximum of 50 characters.  The API supports idempotency for safely retrying requests without accidentally performing the same operation twice. For example, if a request to create an authorization fails due to a network connection error, you can retry the request with the same idempotency key to guarantee that only a single authorization is created.  This idempotency key expires after 48 hours, after which the same key can be reused. Failure to provide this key may result in duplicate charges.  Required for guaranteed ACH customers.
	IdempotencyKey NullableString `json:"idempotency_key,omitempty"`
	// If the end user is initiating the specific transfer themselves via an interactive UI, this should be `true`; for automatic recurring payments where the end user is not actually initiating each individual transfer, it should be `false`.
	UserPresent NullableBool `json:"user_present,omitempty"`
	// If set to `false`, Plaid will not offer a `guarantee_decision` for this request (Guarantee customers only).
	WithGuarantee NullableBool `json:"with_guarantee,omitempty"`
	// The unique identifier returned by Plaid's [beacon](https://plaid.com/docs/transfer/guarantee/#using-a-beacon) when it is run on your webpage.
	BeaconSessionId NullableString `json:"beacon_session_id,omitempty"`
	// The Plaid client ID that is the originator of this transfer. Only needed if creating transfers on behalf of another client as a third-party sender (TPS).
	OriginatorClientId NullableString `json:"originator_client_id,omitempty"`
	CreditFundsSource NullableTransferCreditFundsSource `json:"credit_funds_source,omitempty"`
	// Plaid’s unique identifier for a test clock. This field may only be used when using `sandbox` environment. If provided, the `authorization` is created at the `virtual_time` on the provided `test_clock`.
	TestClockId NullableString `json:"test_clock_id,omitempty"`
}

// NewTransferAuthorizationCreateRequest instantiates a new TransferAuthorizationCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferAuthorizationCreateRequest(type_ TransferType, network TransferNetwork, amount string, user TransferAuthorizationUserInRequest) *TransferAuthorizationCreateRequest {
	this := TransferAuthorizationCreateRequest{}
	this.Type = type_
	this.Network = network
	this.Amount = amount
	this.User = user
	var withGuarantee bool = true
	this.WithGuarantee = *NewNullableBool(&withGuarantee)
	return &this
}

// NewTransferAuthorizationCreateRequestWithDefaults instantiates a new TransferAuthorizationCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferAuthorizationCreateRequestWithDefaults() *TransferAuthorizationCreateRequest {
	this := TransferAuthorizationCreateRequest{}
	var withGuarantee bool = true
	this.WithGuarantee = *NewNullableBool(&withGuarantee)
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransferAuthorizationCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransferAuthorizationCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransferAuthorizationCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransferAuthorizationCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransferAuthorizationCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransferAuthorizationCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *TransferAuthorizationCreateRequest) GetAccessToken() string {
	if o == nil || o.AccessToken == nil {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *TransferAuthorizationCreateRequest) HasAccessToken() bool {
	if o != nil && o.AccessToken != nil {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *TransferAuthorizationCreateRequest) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *TransferAuthorizationCreateRequest) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *TransferAuthorizationCreateRequest) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *TransferAuthorizationCreateRequest) SetAccountId(v string) {
	o.AccountId = &v
}

// GetFundingAccountId returns the FundingAccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferAuthorizationCreateRequest) GetFundingAccountId() string {
	if o == nil || o.FundingAccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.FundingAccountId.Get()
}

// GetFundingAccountIdOk returns a tuple with the FundingAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferAuthorizationCreateRequest) GetFundingAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FundingAccountId.Get(), o.FundingAccountId.IsSet()
}

// HasFundingAccountId returns a boolean if a field has been set.
func (o *TransferAuthorizationCreateRequest) HasFundingAccountId() bool {
	if o != nil && o.FundingAccountId.IsSet() {
		return true
	}

	return false
}

// SetFundingAccountId gets a reference to the given NullableString and assigns it to the FundingAccountId field.
func (o *TransferAuthorizationCreateRequest) SetFundingAccountId(v string) {
	o.FundingAccountId.Set(&v)
}
// SetFundingAccountIdNil sets the value for FundingAccountId to be an explicit nil
func (o *TransferAuthorizationCreateRequest) SetFundingAccountIdNil() {
	o.FundingAccountId.Set(nil)
}

// UnsetFundingAccountId ensures that no value is present for FundingAccountId, not even an explicit nil
func (o *TransferAuthorizationCreateRequest) UnsetFundingAccountId() {
	o.FundingAccountId.Unset()
}

// GetPaymentProfileToken returns the PaymentProfileToken field value if set, zero value otherwise.
func (o *TransferAuthorizationCreateRequest) GetPaymentProfileToken() string {
	if o == nil || o.PaymentProfileToken == nil {
		var ret string
		return ret
	}
	return *o.PaymentProfileToken
}

// GetPaymentProfileTokenOk returns a tuple with the PaymentProfileToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetPaymentProfileTokenOk() (*string, bool) {
	if o == nil || o.PaymentProfileToken == nil {
		return nil, false
	}
	return o.PaymentProfileToken, true
}

// HasPaymentProfileToken returns a boolean if a field has been set.
func (o *TransferAuthorizationCreateRequest) HasPaymentProfileToken() bool {
	if o != nil && o.PaymentProfileToken != nil {
		return true
	}

	return false
}

// SetPaymentProfileToken gets a reference to the given string and assigns it to the PaymentProfileToken field.
func (o *TransferAuthorizationCreateRequest) SetPaymentProfileToken(v string) {
	o.PaymentProfileToken = &v
}

// GetType returns the Type field value
func (o *TransferAuthorizationCreateRequest) GetType() TransferType {
	if o == nil {
		var ret TransferType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetTypeOk() (*TransferType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransferAuthorizationCreateRequest) SetType(v TransferType) {
	o.Type = v
}

// GetNetwork returns the Network field value
func (o *TransferAuthorizationCreateRequest) GetNetwork() TransferNetwork {
	if o == nil {
		var ret TransferNetwork
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetNetworkOk() (*TransferNetwork, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *TransferAuthorizationCreateRequest) SetNetwork(v TransferNetwork) {
	o.Network = v
}

// GetAmount returns the Amount field value
func (o *TransferAuthorizationCreateRequest) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransferAuthorizationCreateRequest) SetAmount(v string) {
	o.Amount = v
}

// GetAchClass returns the AchClass field value if set, zero value otherwise.
func (o *TransferAuthorizationCreateRequest) GetAchClass() ACHClass {
	if o == nil || o.AchClass == nil {
		var ret ACHClass
		return ret
	}
	return *o.AchClass
}

// GetAchClassOk returns a tuple with the AchClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetAchClassOk() (*ACHClass, bool) {
	if o == nil || o.AchClass == nil {
		return nil, false
	}
	return o.AchClass, true
}

// HasAchClass returns a boolean if a field has been set.
func (o *TransferAuthorizationCreateRequest) HasAchClass() bool {
	if o != nil && o.AchClass != nil {
		return true
	}

	return false
}

// SetAchClass gets a reference to the given ACHClass and assigns it to the AchClass field.
func (o *TransferAuthorizationCreateRequest) SetAchClass(v ACHClass) {
	o.AchClass = &v
}

// GetUser returns the User field value
func (o *TransferAuthorizationCreateRequest) GetUser() TransferAuthorizationUserInRequest {
	if o == nil {
		var ret TransferAuthorizationUserInRequest
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetUserOk() (*TransferAuthorizationUserInRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *TransferAuthorizationCreateRequest) SetUser(v TransferAuthorizationUserInRequest) {
	o.User = v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *TransferAuthorizationCreateRequest) GetDevice() TransferAuthorizationDevice {
	if o == nil || o.Device == nil {
		var ret TransferAuthorizationDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetDeviceOk() (*TransferAuthorizationDevice, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *TransferAuthorizationCreateRequest) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given TransferAuthorizationDevice and assigns it to the Device field.
func (o *TransferAuthorizationCreateRequest) SetDevice(v TransferAuthorizationDevice) {
	o.Device = &v
}

// GetOriginationAccountId returns the OriginationAccountId field value if set, zero value otherwise.
func (o *TransferAuthorizationCreateRequest) GetOriginationAccountId() string {
	if o == nil || o.OriginationAccountId == nil {
		var ret string
		return ret
	}
	return *o.OriginationAccountId
}

// GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetOriginationAccountIdOk() (*string, bool) {
	if o == nil || o.OriginationAccountId == nil {
		return nil, false
	}
	return o.OriginationAccountId, true
}

// HasOriginationAccountId returns a boolean if a field has been set.
func (o *TransferAuthorizationCreateRequest) HasOriginationAccountId() bool {
	if o != nil && o.OriginationAccountId != nil {
		return true
	}

	return false
}

// SetOriginationAccountId gets a reference to the given string and assigns it to the OriginationAccountId field.
func (o *TransferAuthorizationCreateRequest) SetOriginationAccountId(v string) {
	o.OriginationAccountId = &v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value if set, zero value otherwise.
func (o *TransferAuthorizationCreateRequest) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil || o.IsoCurrencyCode == nil {
		return nil, false
	}
	return o.IsoCurrencyCode, true
}

// HasIsoCurrencyCode returns a boolean if a field has been set.
func (o *TransferAuthorizationCreateRequest) HasIsoCurrencyCode() bool {
	if o != nil && o.IsoCurrencyCode != nil {
		return true
	}

	return false
}

// SetIsoCurrencyCode gets a reference to the given string and assigns it to the IsoCurrencyCode field.
func (o *TransferAuthorizationCreateRequest) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = &v
}

// GetIdempotencyKey returns the IdempotencyKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferAuthorizationCreateRequest) GetIdempotencyKey() string {
	if o == nil || o.IdempotencyKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.IdempotencyKey.Get()
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferAuthorizationCreateRequest) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IdempotencyKey.Get(), o.IdempotencyKey.IsSet()
}

// HasIdempotencyKey returns a boolean if a field has been set.
func (o *TransferAuthorizationCreateRequest) HasIdempotencyKey() bool {
	if o != nil && o.IdempotencyKey.IsSet() {
		return true
	}

	return false
}

// SetIdempotencyKey gets a reference to the given NullableString and assigns it to the IdempotencyKey field.
func (o *TransferAuthorizationCreateRequest) SetIdempotencyKey(v string) {
	o.IdempotencyKey.Set(&v)
}
// SetIdempotencyKeyNil sets the value for IdempotencyKey to be an explicit nil
func (o *TransferAuthorizationCreateRequest) SetIdempotencyKeyNil() {
	o.IdempotencyKey.Set(nil)
}

// UnsetIdempotencyKey ensures that no value is present for IdempotencyKey, not even an explicit nil
func (o *TransferAuthorizationCreateRequest) UnsetIdempotencyKey() {
	o.IdempotencyKey.Unset()
}

// GetUserPresent returns the UserPresent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferAuthorizationCreateRequest) GetUserPresent() bool {
	if o == nil || o.UserPresent.Get() == nil {
		var ret bool
		return ret
	}
	return *o.UserPresent.Get()
}

// GetUserPresentOk returns a tuple with the UserPresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferAuthorizationCreateRequest) GetUserPresentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserPresent.Get(), o.UserPresent.IsSet()
}

// HasUserPresent returns a boolean if a field has been set.
func (o *TransferAuthorizationCreateRequest) HasUserPresent() bool {
	if o != nil && o.UserPresent.IsSet() {
		return true
	}

	return false
}

// SetUserPresent gets a reference to the given NullableBool and assigns it to the UserPresent field.
func (o *TransferAuthorizationCreateRequest) SetUserPresent(v bool) {
	o.UserPresent.Set(&v)
}
// SetUserPresentNil sets the value for UserPresent to be an explicit nil
func (o *TransferAuthorizationCreateRequest) SetUserPresentNil() {
	o.UserPresent.Set(nil)
}

// UnsetUserPresent ensures that no value is present for UserPresent, not even an explicit nil
func (o *TransferAuthorizationCreateRequest) UnsetUserPresent() {
	o.UserPresent.Unset()
}

// GetWithGuarantee returns the WithGuarantee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferAuthorizationCreateRequest) GetWithGuarantee() bool {
	if o == nil || o.WithGuarantee.Get() == nil {
		var ret bool
		return ret
	}
	return *o.WithGuarantee.Get()
}

// GetWithGuaranteeOk returns a tuple with the WithGuarantee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferAuthorizationCreateRequest) GetWithGuaranteeOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WithGuarantee.Get(), o.WithGuarantee.IsSet()
}

// HasWithGuarantee returns a boolean if a field has been set.
func (o *TransferAuthorizationCreateRequest) HasWithGuarantee() bool {
	if o != nil && o.WithGuarantee.IsSet() {
		return true
	}

	return false
}

// SetWithGuarantee gets a reference to the given NullableBool and assigns it to the WithGuarantee field.
func (o *TransferAuthorizationCreateRequest) SetWithGuarantee(v bool) {
	o.WithGuarantee.Set(&v)
}
// SetWithGuaranteeNil sets the value for WithGuarantee to be an explicit nil
func (o *TransferAuthorizationCreateRequest) SetWithGuaranteeNil() {
	o.WithGuarantee.Set(nil)
}

// UnsetWithGuarantee ensures that no value is present for WithGuarantee, not even an explicit nil
func (o *TransferAuthorizationCreateRequest) UnsetWithGuarantee() {
	o.WithGuarantee.Unset()
}

// GetBeaconSessionId returns the BeaconSessionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferAuthorizationCreateRequest) GetBeaconSessionId() string {
	if o == nil || o.BeaconSessionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.BeaconSessionId.Get()
}

// GetBeaconSessionIdOk returns a tuple with the BeaconSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferAuthorizationCreateRequest) GetBeaconSessionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BeaconSessionId.Get(), o.BeaconSessionId.IsSet()
}

// HasBeaconSessionId returns a boolean if a field has been set.
func (o *TransferAuthorizationCreateRequest) HasBeaconSessionId() bool {
	if o != nil && o.BeaconSessionId.IsSet() {
		return true
	}

	return false
}

// SetBeaconSessionId gets a reference to the given NullableString and assigns it to the BeaconSessionId field.
func (o *TransferAuthorizationCreateRequest) SetBeaconSessionId(v string) {
	o.BeaconSessionId.Set(&v)
}
// SetBeaconSessionIdNil sets the value for BeaconSessionId to be an explicit nil
func (o *TransferAuthorizationCreateRequest) SetBeaconSessionIdNil() {
	o.BeaconSessionId.Set(nil)
}

// UnsetBeaconSessionId ensures that no value is present for BeaconSessionId, not even an explicit nil
func (o *TransferAuthorizationCreateRequest) UnsetBeaconSessionId() {
	o.BeaconSessionId.Unset()
}

// GetOriginatorClientId returns the OriginatorClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferAuthorizationCreateRequest) GetOriginatorClientId() string {
	if o == nil || o.OriginatorClientId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OriginatorClientId.Get()
}

// GetOriginatorClientIdOk returns a tuple with the OriginatorClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferAuthorizationCreateRequest) GetOriginatorClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginatorClientId.Get(), o.OriginatorClientId.IsSet()
}

// HasOriginatorClientId returns a boolean if a field has been set.
func (o *TransferAuthorizationCreateRequest) HasOriginatorClientId() bool {
	if o != nil && o.OriginatorClientId.IsSet() {
		return true
	}

	return false
}

// SetOriginatorClientId gets a reference to the given NullableString and assigns it to the OriginatorClientId field.
func (o *TransferAuthorizationCreateRequest) SetOriginatorClientId(v string) {
	o.OriginatorClientId.Set(&v)
}
// SetOriginatorClientIdNil sets the value for OriginatorClientId to be an explicit nil
func (o *TransferAuthorizationCreateRequest) SetOriginatorClientIdNil() {
	o.OriginatorClientId.Set(nil)
}

// UnsetOriginatorClientId ensures that no value is present for OriginatorClientId, not even an explicit nil
func (o *TransferAuthorizationCreateRequest) UnsetOriginatorClientId() {
	o.OriginatorClientId.Unset()
}

// GetCreditFundsSource returns the CreditFundsSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferAuthorizationCreateRequest) GetCreditFundsSource() TransferCreditFundsSource {
	if o == nil || o.CreditFundsSource.Get() == nil {
		var ret TransferCreditFundsSource
		return ret
	}
	return *o.CreditFundsSource.Get()
}

// GetCreditFundsSourceOk returns a tuple with the CreditFundsSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferAuthorizationCreateRequest) GetCreditFundsSourceOk() (*TransferCreditFundsSource, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreditFundsSource.Get(), o.CreditFundsSource.IsSet()
}

// HasCreditFundsSource returns a boolean if a field has been set.
func (o *TransferAuthorizationCreateRequest) HasCreditFundsSource() bool {
	if o != nil && o.CreditFundsSource.IsSet() {
		return true
	}

	return false
}

// SetCreditFundsSource gets a reference to the given NullableTransferCreditFundsSource and assigns it to the CreditFundsSource field.
func (o *TransferAuthorizationCreateRequest) SetCreditFundsSource(v TransferCreditFundsSource) {
	o.CreditFundsSource.Set(&v)
}
// SetCreditFundsSourceNil sets the value for CreditFundsSource to be an explicit nil
func (o *TransferAuthorizationCreateRequest) SetCreditFundsSourceNil() {
	o.CreditFundsSource.Set(nil)
}

// UnsetCreditFundsSource ensures that no value is present for CreditFundsSource, not even an explicit nil
func (o *TransferAuthorizationCreateRequest) UnsetCreditFundsSource() {
	o.CreditFundsSource.Unset()
}

// GetTestClockId returns the TestClockId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferAuthorizationCreateRequest) GetTestClockId() string {
	if o == nil || o.TestClockId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TestClockId.Get()
}

// GetTestClockIdOk returns a tuple with the TestClockId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferAuthorizationCreateRequest) GetTestClockIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TestClockId.Get(), o.TestClockId.IsSet()
}

// HasTestClockId returns a boolean if a field has been set.
func (o *TransferAuthorizationCreateRequest) HasTestClockId() bool {
	if o != nil && o.TestClockId.IsSet() {
		return true
	}

	return false
}

// SetTestClockId gets a reference to the given NullableString and assigns it to the TestClockId field.
func (o *TransferAuthorizationCreateRequest) SetTestClockId(v string) {
	o.TestClockId.Set(&v)
}
// SetTestClockIdNil sets the value for TestClockId to be an explicit nil
func (o *TransferAuthorizationCreateRequest) SetTestClockIdNil() {
	o.TestClockId.Set(nil)
}

// UnsetTestClockId ensures that no value is present for TestClockId, not even an explicit nil
func (o *TransferAuthorizationCreateRequest) UnsetTestClockId() {
	o.TestClockId.Unset()
}

func (o TransferAuthorizationCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.AccessToken != nil {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.FundingAccountId.IsSet() {
		toSerialize["funding_account_id"] = o.FundingAccountId.Get()
	}
	if o.PaymentProfileToken != nil {
		toSerialize["payment_profile_token"] = o.PaymentProfileToken
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
	if o.AchClass != nil {
		toSerialize["ach_class"] = o.AchClass
	}
	if true {
		toSerialize["user"] = o.User
	}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.OriginationAccountId != nil {
		toSerialize["origination_account_id"] = o.OriginationAccountId
	}
	if o.IsoCurrencyCode != nil {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}
	if o.IdempotencyKey.IsSet() {
		toSerialize["idempotency_key"] = o.IdempotencyKey.Get()
	}
	if o.UserPresent.IsSet() {
		toSerialize["user_present"] = o.UserPresent.Get()
	}
	if o.WithGuarantee.IsSet() {
		toSerialize["with_guarantee"] = o.WithGuarantee.Get()
	}
	if o.BeaconSessionId.IsSet() {
		toSerialize["beacon_session_id"] = o.BeaconSessionId.Get()
	}
	if o.OriginatorClientId.IsSet() {
		toSerialize["originator_client_id"] = o.OriginatorClientId.Get()
	}
	if o.CreditFundsSource.IsSet() {
		toSerialize["credit_funds_source"] = o.CreditFundsSource.Get()
	}
	if o.TestClockId.IsSet() {
		toSerialize["test_clock_id"] = o.TestClockId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTransferAuthorizationCreateRequest struct {
	value *TransferAuthorizationCreateRequest
	isSet bool
}

func (v NullableTransferAuthorizationCreateRequest) Get() *TransferAuthorizationCreateRequest {
	return v.value
}

func (v *NullableTransferAuthorizationCreateRequest) Set(val *TransferAuthorizationCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferAuthorizationCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferAuthorizationCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferAuthorizationCreateRequest(val *TransferAuthorizationCreateRequest) *NullableTransferAuthorizationCreateRequest {
	return &NullableTransferAuthorizationCreateRequest{value: val, isSet: true}
}

func (v NullableTransferAuthorizationCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferAuthorizationCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


