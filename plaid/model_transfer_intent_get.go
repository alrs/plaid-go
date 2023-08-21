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
	"time"
)

// TransferIntentGet Represents a transfer intent within Transfer UI.
type TransferIntentGet struct {
	// Plaid's unique identifier for a transfer intent object.
	Id string `json:"id"`
	// The datetime the transfer was created. This will be of the form `2006-01-02T15:04:05Z`.
	Created time.Time `json:"created"`
	Status TransferIntentStatus `json:"status"`
	// Plaid's unique identifier for the transfer created through the UI. Returned only if the transfer was successfully created. Null value otherwise.
	TransferId NullableString `json:"transfer_id"`
	FailureReason NullableTransferIntentGetFailureReason `json:"failure_reason"`
	AuthorizationDecision NullableTransferIntentAuthorizationDecision `json:"authorization_decision"`
	AuthorizationDecisionRationale NullableTransferAuthorizationDecisionRationale `json:"authorization_decision_rationale"`
	// The Plaid `account_id` for the account that will be debited or credited. Returned only if `account_id` was set on intent creation.
	AccountId NullableString `json:"account_id,omitempty"`
	// Plaid’s unique identifier for the origination account used for the transfer.
	OriginationAccountId string `json:"origination_account_id"`
	// The id of the funding account to use, available in the Plaid Dashboard. This determines which of your business checking accounts will be credited or debited.
	FundingAccountId string `json:"funding_account_id"`
	// The amount of the transfer (decimal string with two digits of precision e.g. \"10.00\").
	Amount string `json:"amount"`
	Mode TransferIntentCreateMode `json:"mode"`
	Network *TransferIntentCreateNetwork `json:"network,omitempty"`
	AchClass *ACHClass `json:"ach_class,omitempty"`
	User TransferUserInResponse `json:"user"`
	// A description for the underlying transfer. Maximum of 8 characters.
	Description string `json:"description"`
	// The Metadata object is a mapping of client-provided string fields to any string value. The following limitations apply: The JSON values must be Strings (no nested JSON objects allowed) Only ASCII characters may be used Maximum of 50 key/value pairs Maximum key length of 40 characters Maximum value length of 500 characters 
	Metadata map[string]string `json:"metadata,omitempty"`
	// The currency of the transfer amount, e.g. \"USD\"
	IsoCurrencyCode string `json:"iso_currency_code"`
	// When `true`, the transfer requires a `GUARANTEED` decision by Plaid to proceed (Guarantee customers only).
	RequireGuarantee NullableBool `json:"require_guarantee,omitempty"`
	GuaranteeDecision NullableTransferAuthorizationGuaranteeDecision `json:"guarantee_decision"`
	GuaranteeDecisionRationale NullableTransferAuthorizationGuaranteeDecisionRationale `json:"guarantee_decision_rationale"`
	AdditionalProperties map[string]interface{}
}

type _TransferIntentGet TransferIntentGet

// NewTransferIntentGet instantiates a new TransferIntentGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferIntentGet(id string, created time.Time, status TransferIntentStatus, transferId NullableString, failureReason NullableTransferIntentGetFailureReason, authorizationDecision NullableTransferIntentAuthorizationDecision, authorizationDecisionRationale NullableTransferAuthorizationDecisionRationale, originationAccountId string, fundingAccountId string, amount string, mode TransferIntentCreateMode, user TransferUserInResponse, description string, isoCurrencyCode string, guaranteeDecision NullableTransferAuthorizationGuaranteeDecision, guaranteeDecisionRationale NullableTransferAuthorizationGuaranteeDecisionRationale) *TransferIntentGet {
	this := TransferIntentGet{}
	this.Id = id
	this.Created = created
	this.Status = status
	this.TransferId = transferId
	this.FailureReason = failureReason
	this.AuthorizationDecision = authorizationDecision
	this.AuthorizationDecisionRationale = authorizationDecisionRationale
	this.OriginationAccountId = originationAccountId
	this.FundingAccountId = fundingAccountId
	this.Amount = amount
	this.Mode = mode
	var network TransferIntentCreateNetwork = TRANSFERINTENTCREATENETWORK_SAME_DAY_ACH
	this.Network = &network
	this.User = user
	this.Description = description
	this.IsoCurrencyCode = isoCurrencyCode
	this.GuaranteeDecision = guaranteeDecision
	this.GuaranteeDecisionRationale = guaranteeDecisionRationale
	return &this
}

// NewTransferIntentGetWithDefaults instantiates a new TransferIntentGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferIntentGetWithDefaults() *TransferIntentGet {
	this := TransferIntentGet{}
	var network TransferIntentCreateNetwork = TRANSFERINTENTCREATENETWORK_SAME_DAY_ACH
	this.Network = &network
	return &this
}

// GetId returns the Id field value
func (o *TransferIntentGet) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGet) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TransferIntentGet) SetId(v string) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *TransferIntentGet) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGet) GetCreatedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *TransferIntentGet) SetCreated(v time.Time) {
	o.Created = v
}

// GetStatus returns the Status field value
func (o *TransferIntentGet) GetStatus() TransferIntentStatus {
	if o == nil {
		var ret TransferIntentStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGet) GetStatusOk() (*TransferIntentStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TransferIntentGet) SetStatus(v TransferIntentStatus) {
	o.Status = v
}

// GetTransferId returns the TransferId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransferIntentGet) GetTransferId() string {
	if o == nil || o.TransferId.Get() == nil {
		var ret string
		return ret
	}

	return *o.TransferId.Get()
}

// GetTransferIdOk returns a tuple with the TransferId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferIntentGet) GetTransferIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TransferId.Get(), o.TransferId.IsSet()
}

// SetTransferId sets field value
func (o *TransferIntentGet) SetTransferId(v string) {
	o.TransferId.Set(&v)
}

// GetFailureReason returns the FailureReason field value
// If the value is explicit nil, the zero value for TransferIntentGetFailureReason will be returned
func (o *TransferIntentGet) GetFailureReason() TransferIntentGetFailureReason {
	if o == nil || o.FailureReason.Get() == nil {
		var ret TransferIntentGetFailureReason
		return ret
	}

	return *o.FailureReason.Get()
}

// GetFailureReasonOk returns a tuple with the FailureReason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferIntentGet) GetFailureReasonOk() (*TransferIntentGetFailureReason, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FailureReason.Get(), o.FailureReason.IsSet()
}

// SetFailureReason sets field value
func (o *TransferIntentGet) SetFailureReason(v TransferIntentGetFailureReason) {
	o.FailureReason.Set(&v)
}

// GetAuthorizationDecision returns the AuthorizationDecision field value
// If the value is explicit nil, the zero value for TransferIntentAuthorizationDecision will be returned
func (o *TransferIntentGet) GetAuthorizationDecision() TransferIntentAuthorizationDecision {
	if o == nil || o.AuthorizationDecision.Get() == nil {
		var ret TransferIntentAuthorizationDecision
		return ret
	}

	return *o.AuthorizationDecision.Get()
}

// GetAuthorizationDecisionOk returns a tuple with the AuthorizationDecision field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferIntentGet) GetAuthorizationDecisionOk() (*TransferIntentAuthorizationDecision, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AuthorizationDecision.Get(), o.AuthorizationDecision.IsSet()
}

// SetAuthorizationDecision sets field value
func (o *TransferIntentGet) SetAuthorizationDecision(v TransferIntentAuthorizationDecision) {
	o.AuthorizationDecision.Set(&v)
}

// GetAuthorizationDecisionRationale returns the AuthorizationDecisionRationale field value
// If the value is explicit nil, the zero value for TransferAuthorizationDecisionRationale will be returned
func (o *TransferIntentGet) GetAuthorizationDecisionRationale() TransferAuthorizationDecisionRationale {
	if o == nil || o.AuthorizationDecisionRationale.Get() == nil {
		var ret TransferAuthorizationDecisionRationale
		return ret
	}

	return *o.AuthorizationDecisionRationale.Get()
}

// GetAuthorizationDecisionRationaleOk returns a tuple with the AuthorizationDecisionRationale field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferIntentGet) GetAuthorizationDecisionRationaleOk() (*TransferAuthorizationDecisionRationale, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AuthorizationDecisionRationale.Get(), o.AuthorizationDecisionRationale.IsSet()
}

// SetAuthorizationDecisionRationale sets field value
func (o *TransferIntentGet) SetAuthorizationDecisionRationale(v TransferAuthorizationDecisionRationale) {
	o.AuthorizationDecisionRationale.Set(&v)
}

// GetAccountId returns the AccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferIntentGet) GetAccountId() string {
	if o == nil || o.AccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountId.Get()
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferIntentGet) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountId.Get(), o.AccountId.IsSet()
}

// HasAccountId returns a boolean if a field has been set.
func (o *TransferIntentGet) HasAccountId() bool {
	if o != nil && o.AccountId.IsSet() {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given NullableString and assigns it to the AccountId field.
func (o *TransferIntentGet) SetAccountId(v string) {
	o.AccountId.Set(&v)
}
// SetAccountIdNil sets the value for AccountId to be an explicit nil
func (o *TransferIntentGet) SetAccountIdNil() {
	o.AccountId.Set(nil)
}

// UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
func (o *TransferIntentGet) UnsetAccountId() {
	o.AccountId.Unset()
}

// GetOriginationAccountId returns the OriginationAccountId field value
func (o *TransferIntentGet) GetOriginationAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginationAccountId
}

// GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGet) GetOriginationAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OriginationAccountId, true
}

// SetOriginationAccountId sets field value
func (o *TransferIntentGet) SetOriginationAccountId(v string) {
	o.OriginationAccountId = v
}

// GetFundingAccountId returns the FundingAccountId field value
func (o *TransferIntentGet) GetFundingAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FundingAccountId
}

// GetFundingAccountIdOk returns a tuple with the FundingAccountId field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGet) GetFundingAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FundingAccountId, true
}

// SetFundingAccountId sets field value
func (o *TransferIntentGet) SetFundingAccountId(v string) {
	o.FundingAccountId = v
}

// GetAmount returns the Amount field value
func (o *TransferIntentGet) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGet) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransferIntentGet) SetAmount(v string) {
	o.Amount = v
}

// GetMode returns the Mode field value
func (o *TransferIntentGet) GetMode() TransferIntentCreateMode {
	if o == nil {
		var ret TransferIntentCreateMode
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGet) GetModeOk() (*TransferIntentCreateMode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *TransferIntentGet) SetMode(v TransferIntentCreateMode) {
	o.Mode = v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *TransferIntentGet) GetNetwork() TransferIntentCreateNetwork {
	if o == nil || o.Network == nil {
		var ret TransferIntentCreateNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferIntentGet) GetNetworkOk() (*TransferIntentCreateNetwork, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *TransferIntentGet) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given TransferIntentCreateNetwork and assigns it to the Network field.
func (o *TransferIntentGet) SetNetwork(v TransferIntentCreateNetwork) {
	o.Network = &v
}

// GetAchClass returns the AchClass field value if set, zero value otherwise.
func (o *TransferIntentGet) GetAchClass() ACHClass {
	if o == nil || o.AchClass == nil {
		var ret ACHClass
		return ret
	}
	return *o.AchClass
}

// GetAchClassOk returns a tuple with the AchClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferIntentGet) GetAchClassOk() (*ACHClass, bool) {
	if o == nil || o.AchClass == nil {
		return nil, false
	}
	return o.AchClass, true
}

// HasAchClass returns a boolean if a field has been set.
func (o *TransferIntentGet) HasAchClass() bool {
	if o != nil && o.AchClass != nil {
		return true
	}

	return false
}

// SetAchClass gets a reference to the given ACHClass and assigns it to the AchClass field.
func (o *TransferIntentGet) SetAchClass(v ACHClass) {
	o.AchClass = &v
}

// GetUser returns the User field value
func (o *TransferIntentGet) GetUser() TransferUserInResponse {
	if o == nil {
		var ret TransferUserInResponse
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGet) GetUserOk() (*TransferUserInResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *TransferIntentGet) SetUser(v TransferUserInResponse) {
	o.User = v
}

// GetDescription returns the Description field value
func (o *TransferIntentGet) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGet) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TransferIntentGet) SetDescription(v string) {
	o.Description = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferIntentGet) GetMetadata() map[string]string {
	if o == nil  {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferIntentGet) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TransferIntentGet) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *TransferIntentGet) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
func (o *TransferIntentGet) GetIsoCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGet) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsoCurrencyCode, true
}

// SetIsoCurrencyCode sets field value
func (o *TransferIntentGet) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = v
}

// GetRequireGuarantee returns the RequireGuarantee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferIntentGet) GetRequireGuarantee() bool {
	if o == nil || o.RequireGuarantee.Get() == nil {
		var ret bool
		return ret
	}
	return *o.RequireGuarantee.Get()
}

// GetRequireGuaranteeOk returns a tuple with the RequireGuarantee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferIntentGet) GetRequireGuaranteeOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RequireGuarantee.Get(), o.RequireGuarantee.IsSet()
}

// HasRequireGuarantee returns a boolean if a field has been set.
func (o *TransferIntentGet) HasRequireGuarantee() bool {
	if o != nil && o.RequireGuarantee.IsSet() {
		return true
	}

	return false
}

// SetRequireGuarantee gets a reference to the given NullableBool and assigns it to the RequireGuarantee field.
func (o *TransferIntentGet) SetRequireGuarantee(v bool) {
	o.RequireGuarantee.Set(&v)
}
// SetRequireGuaranteeNil sets the value for RequireGuarantee to be an explicit nil
func (o *TransferIntentGet) SetRequireGuaranteeNil() {
	o.RequireGuarantee.Set(nil)
}

// UnsetRequireGuarantee ensures that no value is present for RequireGuarantee, not even an explicit nil
func (o *TransferIntentGet) UnsetRequireGuarantee() {
	o.RequireGuarantee.Unset()
}

// GetGuaranteeDecision returns the GuaranteeDecision field value
// If the value is explicit nil, the zero value for TransferAuthorizationGuaranteeDecision will be returned
func (o *TransferIntentGet) GetGuaranteeDecision() TransferAuthorizationGuaranteeDecision {
	if o == nil || o.GuaranteeDecision.Get() == nil {
		var ret TransferAuthorizationGuaranteeDecision
		return ret
	}

	return *o.GuaranteeDecision.Get()
}

// GetGuaranteeDecisionOk returns a tuple with the GuaranteeDecision field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferIntentGet) GetGuaranteeDecisionOk() (*TransferAuthorizationGuaranteeDecision, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GuaranteeDecision.Get(), o.GuaranteeDecision.IsSet()
}

// SetGuaranteeDecision sets field value
func (o *TransferIntentGet) SetGuaranteeDecision(v TransferAuthorizationGuaranteeDecision) {
	o.GuaranteeDecision.Set(&v)
}

// GetGuaranteeDecisionRationale returns the GuaranteeDecisionRationale field value
// If the value is explicit nil, the zero value for TransferAuthorizationGuaranteeDecisionRationale will be returned
func (o *TransferIntentGet) GetGuaranteeDecisionRationale() TransferAuthorizationGuaranteeDecisionRationale {
	if o == nil || o.GuaranteeDecisionRationale.Get() == nil {
		var ret TransferAuthorizationGuaranteeDecisionRationale
		return ret
	}

	return *o.GuaranteeDecisionRationale.Get()
}

// GetGuaranteeDecisionRationaleOk returns a tuple with the GuaranteeDecisionRationale field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferIntentGet) GetGuaranteeDecisionRationaleOk() (*TransferAuthorizationGuaranteeDecisionRationale, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GuaranteeDecisionRationale.Get(), o.GuaranteeDecisionRationale.IsSet()
}

// SetGuaranteeDecisionRationale sets field value
func (o *TransferIntentGet) SetGuaranteeDecisionRationale(v TransferAuthorizationGuaranteeDecisionRationale) {
	o.GuaranteeDecisionRationale.Set(&v)
}

func (o TransferIntentGet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["transfer_id"] = o.TransferId.Get()
	}
	if true {
		toSerialize["failure_reason"] = o.FailureReason.Get()
	}
	if true {
		toSerialize["authorization_decision"] = o.AuthorizationDecision.Get()
	}
	if true {
		toSerialize["authorization_decision_rationale"] = o.AuthorizationDecisionRationale.Get()
	}
	if o.AccountId.IsSet() {
		toSerialize["account_id"] = o.AccountId.Get()
	}
	if true {
		toSerialize["origination_account_id"] = o.OriginationAccountId
	}
	if true {
		toSerialize["funding_account_id"] = o.FundingAccountId
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.AchClass != nil {
		toSerialize["ach_class"] = o.AchClass
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}
	if o.RequireGuarantee.IsSet() {
		toSerialize["require_guarantee"] = o.RequireGuarantee.Get()
	}
	if true {
		toSerialize["guarantee_decision"] = o.GuaranteeDecision.Get()
	}
	if true {
		toSerialize["guarantee_decision_rationale"] = o.GuaranteeDecisionRationale.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferIntentGet) UnmarshalJSON(bytes []byte) (err error) {
	varTransferIntentGet := _TransferIntentGet{}

	if err = json.Unmarshal(bytes, &varTransferIntentGet); err == nil {
		*o = TransferIntentGet(varTransferIntentGet)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created")
		delete(additionalProperties, "status")
		delete(additionalProperties, "transfer_id")
		delete(additionalProperties, "failure_reason")
		delete(additionalProperties, "authorization_decision")
		delete(additionalProperties, "authorization_decision_rationale")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "origination_account_id")
		delete(additionalProperties, "funding_account_id")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "network")
		delete(additionalProperties, "ach_class")
		delete(additionalProperties, "user")
		delete(additionalProperties, "description")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "require_guarantee")
		delete(additionalProperties, "guarantee_decision")
		delete(additionalProperties, "guarantee_decision_rationale")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferIntentGet struct {
	value *TransferIntentGet
	isSet bool
}

func (v NullableTransferIntentGet) Get() *TransferIntentGet {
	return v.value
}

func (v *NullableTransferIntentGet) Set(val *TransferIntentGet) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferIntentGet) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferIntentGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferIntentGet(val *TransferIntentGet) *NullableTransferIntentGet {
	return &NullableTransferIntentGet{value: val, isSet: true}
}

func (v NullableTransferIntentGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferIntentGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


