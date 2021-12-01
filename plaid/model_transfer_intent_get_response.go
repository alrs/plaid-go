/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.54.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"time"
)

// TransferIntentGetResponse Defines the response schema for `/transfer/intent/get`
type TransferIntentGetResponse struct {
	// Plaid's unique identifier for a transfer intent object.
	Id string `json:"id"`
	// The datetime the transfer was created. This will be of the form `2006-01-02T15:04:05Z`.
	Created time.Time `json:"created"`
	// The status of the transfer intent.
	Status string `json:"status"`
	// Plaid's unique identifier for the transfer created through the UI. Returned only if the transfer was successfully created. Null value otherwise.
	TransferId NullableString `json:"transfer_id"`
	FailureReason NullableTransferIntentGetFailureReason `json:"failure_reason"`
	//  A decision regarding the proposed transfer.  `APPROVED` – The proposed transfer has received the end user's consent and has been approved for processing. Plaid has also reviewed the proposed transfer and has approved it for processing.   `PERMITTED` – Plaid was unable to fetch the information required to approve or decline the proposed transfer. You may proceed with the transfer, but further review is recommended. Plaid is awaiting further instructions from the client.  `DECLINED` – Plaid reviewed the proposed transfer and declined processing. Refer to the `code` field in the `decision_rationale` object for details. Null value otherwise.
	AuthorizationDecision NullableString `json:"authorization_decision"`
	AuthorizationDecisionRationale NullableTransferAuthorizationDecisionRationale `json:"authorization_decision_rationale"`
	// The Plaid `account_id` for the account that will be debited or credited.
	AccountId string `json:"account_id"`
	// Plaid’s unique identifier for the origination account used for the transfer.
	OriginationAccountId string `json:"origination_account_id"`
	// The amount of the transfer (decimal string with two digits of precision e.g. “10.00”).
	Amount string `json:"amount"`
	Mode TransferIntentCreateMode `json:"mode"`
	AchClass ACHClass `json:"ach_class"`
	User TransferUserInResponse `json:"user"`
	// A description for the underlying transfer. Maximum of 8 characters.
	Description string `json:"description"`
	// The Metadata object is a mapping of client-provided string fields to any string value. The following limitations apply: - The JSON values must be Strings (no nested JSON objects allowed) - Only ASCII characters may be used - Maximum of 50 key/value pairs - Maximum key length of 40 characters - Maximum value length of 500 characters 
	Metadata map[string]string `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransferIntentGetResponse TransferIntentGetResponse

// NewTransferIntentGetResponse instantiates a new TransferIntentGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferIntentGetResponse(id string, created time.Time, status string, transferId NullableString, failureReason NullableTransferIntentGetFailureReason, authorizationDecision NullableString, authorizationDecisionRationale NullableTransferAuthorizationDecisionRationale, accountId string, originationAccountId string, amount string, mode TransferIntentCreateMode, achClass ACHClass, user TransferUserInResponse, description string) *TransferIntentGetResponse {
	this := TransferIntentGetResponse{}
	this.Id = id
	this.Created = created
	this.Status = status
	this.TransferId = transferId
	this.FailureReason = failureReason
	this.AuthorizationDecision = authorizationDecision
	this.AuthorizationDecisionRationale = authorizationDecisionRationale
	this.AccountId = accountId
	this.OriginationAccountId = originationAccountId
	this.Amount = amount
	this.Mode = mode
	this.AchClass = achClass
	this.User = user
	this.Description = description
	return &this
}

// NewTransferIntentGetResponseWithDefaults instantiates a new TransferIntentGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferIntentGetResponseWithDefaults() *TransferIntentGetResponse {
	this := TransferIntentGetResponse{}
	return &this
}

// GetId returns the Id field value
func (o *TransferIntentGetResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGetResponse) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TransferIntentGetResponse) SetId(v string) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *TransferIntentGetResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGetResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *TransferIntentGetResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetStatus returns the Status field value
func (o *TransferIntentGetResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGetResponse) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TransferIntentGetResponse) SetStatus(v string) {
	o.Status = v
}

// GetTransferId returns the TransferId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransferIntentGetResponse) GetTransferId() string {
	if o == nil || o.TransferId.Get() == nil {
		var ret string
		return ret
	}

	return *o.TransferId.Get()
}

// GetTransferIdOk returns a tuple with the TransferId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferIntentGetResponse) GetTransferIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TransferId.Get(), o.TransferId.IsSet()
}

// SetTransferId sets field value
func (o *TransferIntentGetResponse) SetTransferId(v string) {
	o.TransferId.Set(&v)
}

// GetFailureReason returns the FailureReason field value
// If the value is explicit nil, the zero value for TransferIntentGetFailureReason will be returned
func (o *TransferIntentGetResponse) GetFailureReason() TransferIntentGetFailureReason {
	if o == nil || o.FailureReason.Get() == nil {
		var ret TransferIntentGetFailureReason
		return ret
	}

	return *o.FailureReason.Get()
}

// GetFailureReasonOk returns a tuple with the FailureReason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferIntentGetResponse) GetFailureReasonOk() (*TransferIntentGetFailureReason, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FailureReason.Get(), o.FailureReason.IsSet()
}

// SetFailureReason sets field value
func (o *TransferIntentGetResponse) SetFailureReason(v TransferIntentGetFailureReason) {
	o.FailureReason.Set(&v)
}

// GetAuthorizationDecision returns the AuthorizationDecision field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransferIntentGetResponse) GetAuthorizationDecision() string {
	if o == nil || o.AuthorizationDecision.Get() == nil {
		var ret string
		return ret
	}

	return *o.AuthorizationDecision.Get()
}

// GetAuthorizationDecisionOk returns a tuple with the AuthorizationDecision field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferIntentGetResponse) GetAuthorizationDecisionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AuthorizationDecision.Get(), o.AuthorizationDecision.IsSet()
}

// SetAuthorizationDecision sets field value
func (o *TransferIntentGetResponse) SetAuthorizationDecision(v string) {
	o.AuthorizationDecision.Set(&v)
}

// GetAuthorizationDecisionRationale returns the AuthorizationDecisionRationale field value
// If the value is explicit nil, the zero value for TransferAuthorizationDecisionRationale will be returned
func (o *TransferIntentGetResponse) GetAuthorizationDecisionRationale() TransferAuthorizationDecisionRationale {
	if o == nil || o.AuthorizationDecisionRationale.Get() == nil {
		var ret TransferAuthorizationDecisionRationale
		return ret
	}

	return *o.AuthorizationDecisionRationale.Get()
}

// GetAuthorizationDecisionRationaleOk returns a tuple with the AuthorizationDecisionRationale field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferIntentGetResponse) GetAuthorizationDecisionRationaleOk() (*TransferAuthorizationDecisionRationale, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AuthorizationDecisionRationale.Get(), o.AuthorizationDecisionRationale.IsSet()
}

// SetAuthorizationDecisionRationale sets field value
func (o *TransferIntentGetResponse) SetAuthorizationDecisionRationale(v TransferAuthorizationDecisionRationale) {
	o.AuthorizationDecisionRationale.Set(&v)
}

// GetAccountId returns the AccountId field value
func (o *TransferIntentGetResponse) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGetResponse) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *TransferIntentGetResponse) SetAccountId(v string) {
	o.AccountId = v
}

// GetOriginationAccountId returns the OriginationAccountId field value
func (o *TransferIntentGetResponse) GetOriginationAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginationAccountId
}

// GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGetResponse) GetOriginationAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OriginationAccountId, true
}

// SetOriginationAccountId sets field value
func (o *TransferIntentGetResponse) SetOriginationAccountId(v string) {
	o.OriginationAccountId = v
}

// GetAmount returns the Amount field value
func (o *TransferIntentGetResponse) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGetResponse) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransferIntentGetResponse) SetAmount(v string) {
	o.Amount = v
}

// GetMode returns the Mode field value
func (o *TransferIntentGetResponse) GetMode() TransferIntentCreateMode {
	if o == nil {
		var ret TransferIntentCreateMode
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGetResponse) GetModeOk() (*TransferIntentCreateMode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *TransferIntentGetResponse) SetMode(v TransferIntentCreateMode) {
	o.Mode = v
}

// GetAchClass returns the AchClass field value
func (o *TransferIntentGetResponse) GetAchClass() ACHClass {
	if o == nil {
		var ret ACHClass
		return ret
	}

	return o.AchClass
}

// GetAchClassOk returns a tuple with the AchClass field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGetResponse) GetAchClassOk() (*ACHClass, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AchClass, true
}

// SetAchClass sets field value
func (o *TransferIntentGetResponse) SetAchClass(v ACHClass) {
	o.AchClass = v
}

// GetUser returns the User field value
func (o *TransferIntentGetResponse) GetUser() TransferUserInResponse {
	if o == nil {
		var ret TransferUserInResponse
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGetResponse) GetUserOk() (*TransferUserInResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *TransferIntentGetResponse) SetUser(v TransferUserInResponse) {
	o.User = v
}

// GetDescription returns the Description field value
func (o *TransferIntentGetResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGetResponse) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TransferIntentGetResponse) SetDescription(v string) {
	o.Description = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferIntentGetResponse) GetMetadata() map[string]string {
	if o == nil  {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferIntentGetResponse) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TransferIntentGetResponse) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *TransferIntentGetResponse) SetMetadata(v map[string]string) {
	o.Metadata = v
}

func (o TransferIntentGetResponse) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["origination_account_id"] = o.OriginationAccountId
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if true {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferIntentGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferIntentGetResponse := _TransferIntentGetResponse{}

	if err = json.Unmarshal(bytes, &varTransferIntentGetResponse); err == nil {
		*o = TransferIntentGetResponse(varTransferIntentGetResponse)
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
		delete(additionalProperties, "amount")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "ach_class")
		delete(additionalProperties, "user")
		delete(additionalProperties, "description")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferIntentGetResponse struct {
	value *TransferIntentGetResponse
	isSet bool
}

func (v NullableTransferIntentGetResponse) Get() *TransferIntentGetResponse {
	return v.value
}

func (v *NullableTransferIntentGetResponse) Set(val *TransferIntentGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferIntentGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferIntentGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferIntentGetResponse(val *TransferIntentGetResponse) *NullableTransferIntentGetResponse {
	return &NullableTransferIntentGetResponse{value: val, isSet: true}
}

func (v NullableTransferIntentGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferIntentGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

