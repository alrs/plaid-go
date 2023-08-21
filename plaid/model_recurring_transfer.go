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

// RecurringTransfer Represents a recurring transfer within the Transfers API.
type RecurringTransfer struct {
	// Plaid’s unique identifier for a recurring transfer.
	RecurringTransferId string `json:"recurring_transfer_id"`
	// The datetime when this transfer was created. This will be of the form `2006-01-02T15:04:05Z`
	Created time.Time `json:"created"`
	// A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).  The next transfer origination date after bank holiday adjustment.
	NextOriginationDate NullableString `json:"next_origination_date"`
	// Plaid’s unique identifier for a test clock.
	TestClockId NullableString `json:"test_clock_id,omitempty"`
	Type TransferType `json:"type"`
	// The amount of the transfer (decimal string with two digits of precision e.g. \"10.00\").
	Amount string `json:"amount"`
	Status TransferRecurringStatus `json:"status"`
	AchClass *ACHClass `json:"ach_class,omitempty"`
	Network TransferNetwork `json:"network"`
	// Plaid’s unique identifier for the origination account that was used for this transfer.
	OriginationAccountId string `json:"origination_account_id"`
	// The Plaid `account_id` corresponding to the end-user account that will be debited or credited.
	AccountId string `json:"account_id"`
	// The id of the funding account to use, available in the Plaid Dashboard. This determines which of your business checking accounts will be credited or debited.
	FundingAccountId string `json:"funding_account_id"`
	// The currency of the transfer amount, e.g. \"USD\"
	IsoCurrencyCode string `json:"iso_currency_code"`
	// The description of the recurring transfer.
	Description string `json:"description"`
	TransferIds []string `json:"transfer_ids"`
	User TransferUserInResponse `json:"user"`
	Schedule TransferRecurringSchedule `json:"schedule"`
	AdditionalProperties map[string]interface{}
}

type _RecurringTransfer RecurringTransfer

// NewRecurringTransfer instantiates a new RecurringTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecurringTransfer(recurringTransferId string, created time.Time, nextOriginationDate NullableString, type_ TransferType, amount string, status TransferRecurringStatus, network TransferNetwork, originationAccountId string, accountId string, fundingAccountId string, isoCurrencyCode string, description string, transferIds []string, user TransferUserInResponse, schedule TransferRecurringSchedule) *RecurringTransfer {
	this := RecurringTransfer{}
	this.RecurringTransferId = recurringTransferId
	this.Created = created
	this.NextOriginationDate = nextOriginationDate
	this.Type = type_
	this.Amount = amount
	this.Status = status
	this.Network = network
	this.OriginationAccountId = originationAccountId
	this.AccountId = accountId
	this.FundingAccountId = fundingAccountId
	this.IsoCurrencyCode = isoCurrencyCode
	this.Description = description
	this.TransferIds = transferIds
	this.User = user
	this.Schedule = schedule
	return &this
}

// NewRecurringTransferWithDefaults instantiates a new RecurringTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecurringTransferWithDefaults() *RecurringTransfer {
	this := RecurringTransfer{}
	return &this
}

// GetRecurringTransferId returns the RecurringTransferId field value
func (o *RecurringTransfer) GetRecurringTransferId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecurringTransferId
}

// GetRecurringTransferIdOk returns a tuple with the RecurringTransferId field value
// and a boolean to check if the value has been set.
func (o *RecurringTransfer) GetRecurringTransferIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecurringTransferId, true
}

// SetRecurringTransferId sets field value
func (o *RecurringTransfer) SetRecurringTransferId(v string) {
	o.RecurringTransferId = v
}

// GetCreated returns the Created field value
func (o *RecurringTransfer) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *RecurringTransfer) GetCreatedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *RecurringTransfer) SetCreated(v time.Time) {
	o.Created = v
}

// GetNextOriginationDate returns the NextOriginationDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RecurringTransfer) GetNextOriginationDate() string {
	if o == nil || o.NextOriginationDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextOriginationDate.Get()
}

// GetNextOriginationDateOk returns a tuple with the NextOriginationDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecurringTransfer) GetNextOriginationDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextOriginationDate.Get(), o.NextOriginationDate.IsSet()
}

// SetNextOriginationDate sets field value
func (o *RecurringTransfer) SetNextOriginationDate(v string) {
	o.NextOriginationDate.Set(&v)
}

// GetTestClockId returns the TestClockId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecurringTransfer) GetTestClockId() string {
	if o == nil || o.TestClockId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TestClockId.Get()
}

// GetTestClockIdOk returns a tuple with the TestClockId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecurringTransfer) GetTestClockIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TestClockId.Get(), o.TestClockId.IsSet()
}

// HasTestClockId returns a boolean if a field has been set.
func (o *RecurringTransfer) HasTestClockId() bool {
	if o != nil && o.TestClockId.IsSet() {
		return true
	}

	return false
}

// SetTestClockId gets a reference to the given NullableString and assigns it to the TestClockId field.
func (o *RecurringTransfer) SetTestClockId(v string) {
	o.TestClockId.Set(&v)
}
// SetTestClockIdNil sets the value for TestClockId to be an explicit nil
func (o *RecurringTransfer) SetTestClockIdNil() {
	o.TestClockId.Set(nil)
}

// UnsetTestClockId ensures that no value is present for TestClockId, not even an explicit nil
func (o *RecurringTransfer) UnsetTestClockId() {
	o.TestClockId.Unset()
}

// GetType returns the Type field value
func (o *RecurringTransfer) GetType() TransferType {
	if o == nil {
		var ret TransferType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RecurringTransfer) GetTypeOk() (*TransferType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RecurringTransfer) SetType(v TransferType) {
	o.Type = v
}

// GetAmount returns the Amount field value
func (o *RecurringTransfer) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *RecurringTransfer) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *RecurringTransfer) SetAmount(v string) {
	o.Amount = v
}

// GetStatus returns the Status field value
func (o *RecurringTransfer) GetStatus() TransferRecurringStatus {
	if o == nil {
		var ret TransferRecurringStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RecurringTransfer) GetStatusOk() (*TransferRecurringStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RecurringTransfer) SetStatus(v TransferRecurringStatus) {
	o.Status = v
}

// GetAchClass returns the AchClass field value if set, zero value otherwise.
func (o *RecurringTransfer) GetAchClass() ACHClass {
	if o == nil || o.AchClass == nil {
		var ret ACHClass
		return ret
	}
	return *o.AchClass
}

// GetAchClassOk returns a tuple with the AchClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringTransfer) GetAchClassOk() (*ACHClass, bool) {
	if o == nil || o.AchClass == nil {
		return nil, false
	}
	return o.AchClass, true
}

// HasAchClass returns a boolean if a field has been set.
func (o *RecurringTransfer) HasAchClass() bool {
	if o != nil && o.AchClass != nil {
		return true
	}

	return false
}

// SetAchClass gets a reference to the given ACHClass and assigns it to the AchClass field.
func (o *RecurringTransfer) SetAchClass(v ACHClass) {
	o.AchClass = &v
}

// GetNetwork returns the Network field value
func (o *RecurringTransfer) GetNetwork() TransferNetwork {
	if o == nil {
		var ret TransferNetwork
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *RecurringTransfer) GetNetworkOk() (*TransferNetwork, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *RecurringTransfer) SetNetwork(v TransferNetwork) {
	o.Network = v
}

// GetOriginationAccountId returns the OriginationAccountId field value
func (o *RecurringTransfer) GetOriginationAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginationAccountId
}

// GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field value
// and a boolean to check if the value has been set.
func (o *RecurringTransfer) GetOriginationAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OriginationAccountId, true
}

// SetOriginationAccountId sets field value
func (o *RecurringTransfer) SetOriginationAccountId(v string) {
	o.OriginationAccountId = v
}

// GetAccountId returns the AccountId field value
func (o *RecurringTransfer) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *RecurringTransfer) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *RecurringTransfer) SetAccountId(v string) {
	o.AccountId = v
}

// GetFundingAccountId returns the FundingAccountId field value
func (o *RecurringTransfer) GetFundingAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FundingAccountId
}

// GetFundingAccountIdOk returns a tuple with the FundingAccountId field value
// and a boolean to check if the value has been set.
func (o *RecurringTransfer) GetFundingAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FundingAccountId, true
}

// SetFundingAccountId sets field value
func (o *RecurringTransfer) SetFundingAccountId(v string) {
	o.FundingAccountId = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
func (o *RecurringTransfer) GetIsoCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *RecurringTransfer) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsoCurrencyCode, true
}

// SetIsoCurrencyCode sets field value
func (o *RecurringTransfer) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = v
}

// GetDescription returns the Description field value
func (o *RecurringTransfer) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RecurringTransfer) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *RecurringTransfer) SetDescription(v string) {
	o.Description = v
}

// GetTransferIds returns the TransferIds field value
func (o *RecurringTransfer) GetTransferIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TransferIds
}

// GetTransferIdsOk returns a tuple with the TransferIds field value
// and a boolean to check if the value has been set.
func (o *RecurringTransfer) GetTransferIdsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransferIds, true
}

// SetTransferIds sets field value
func (o *RecurringTransfer) SetTransferIds(v []string) {
	o.TransferIds = v
}

// GetUser returns the User field value
func (o *RecurringTransfer) GetUser() TransferUserInResponse {
	if o == nil {
		var ret TransferUserInResponse
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *RecurringTransfer) GetUserOk() (*TransferUserInResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *RecurringTransfer) SetUser(v TransferUserInResponse) {
	o.User = v
}

// GetSchedule returns the Schedule field value
func (o *RecurringTransfer) GetSchedule() TransferRecurringSchedule {
	if o == nil {
		var ret TransferRecurringSchedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *RecurringTransfer) GetScheduleOk() (*TransferRecurringSchedule, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *RecurringTransfer) SetSchedule(v TransferRecurringSchedule) {
	o.Schedule = v
}

func (o RecurringTransfer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recurring_transfer_id"] = o.RecurringTransferId
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["next_origination_date"] = o.NextOriginationDate.Get()
	}
	if o.TestClockId.IsSet() {
		toSerialize["test_clock_id"] = o.TestClockId.Get()
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.AchClass != nil {
		toSerialize["ach_class"] = o.AchClass
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["origination_account_id"] = o.OriginationAccountId
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["funding_account_id"] = o.FundingAccountId
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["transfer_ids"] = o.TransferIds
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["schedule"] = o.Schedule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecurringTransfer) UnmarshalJSON(bytes []byte) (err error) {
	varRecurringTransfer := _RecurringTransfer{}

	if err = json.Unmarshal(bytes, &varRecurringTransfer); err == nil {
		*o = RecurringTransfer(varRecurringTransfer)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "recurring_transfer_id")
		delete(additionalProperties, "created")
		delete(additionalProperties, "next_origination_date")
		delete(additionalProperties, "test_clock_id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "status")
		delete(additionalProperties, "ach_class")
		delete(additionalProperties, "network")
		delete(additionalProperties, "origination_account_id")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "funding_account_id")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "description")
		delete(additionalProperties, "transfer_ids")
		delete(additionalProperties, "user")
		delete(additionalProperties, "schedule")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecurringTransfer struct {
	value *RecurringTransfer
	isSet bool
}

func (v NullableRecurringTransfer) Get() *RecurringTransfer {
	return v.value
}

func (v *NullableRecurringTransfer) Set(val *RecurringTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurringTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurringTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurringTransfer(val *RecurringTransfer) *NullableRecurringTransfer {
	return &NullableRecurringTransfer{value: val, isSet: true}
}

func (v NullableRecurringTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecurringTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


