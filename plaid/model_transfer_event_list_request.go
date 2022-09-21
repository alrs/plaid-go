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
	"time"
)

// TransferEventListRequest Defines the request schema for `/transfer/event/list`
type TransferEventListRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The start datetime of transfers to list. This should be in RFC 3339 format (i.e. `2019-12-06T22:35:49Z`)
	StartDate NullableTime `json:"start_date,omitempty"`
	// The end datetime of transfers to list. This should be in RFC 3339 format (i.e. `2019-12-06T22:35:49Z`)
	EndDate NullableTime `json:"end_date,omitempty"`
	// Plaid’s unique identifier for a transfer.
	TransferId NullableString `json:"transfer_id,omitempty"`
	// The account ID to get events for all transactions to/from an account.
	AccountId NullableString `json:"account_id,omitempty"`
	TransferType NullableTransferEventListTransferType `json:"transfer_type,omitempty"`
	// Filter events by event type.
	EventTypes *[]TransferEventType `json:"event_types,omitempty"`
	// Plaid’s unique identifier for a sweep.
	SweepId *string `json:"sweep_id,omitempty"`
	// The maximum number of transfer events to return. If the number of events matching the above parameters is greater than `count`, the most recent events will be returned.
	Count NullableInt32 `json:"count,omitempty"`
	// The offset into the list of transfer events. When `count`=25 and `offset`=0, the first 25 events will be returned. When `count`=25 and `offset`=25, the next 25 events will be returned.
	Offset NullableInt32 `json:"offset,omitempty"`
	// The origination account ID to get events for transfers from a specific origination account.
	OriginationAccountId NullableString `json:"origination_account_id,omitempty"`
}

// NewTransferEventListRequest instantiates a new TransferEventListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferEventListRequest() *TransferEventListRequest {
	this := TransferEventListRequest{}
	var count int32 = 25
	this.Count = *NewNullableInt32(&count)
	var offset int32 = 0
	this.Offset = *NewNullableInt32(&offset)
	return &this
}

// NewTransferEventListRequestWithDefaults instantiates a new TransferEventListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferEventListRequestWithDefaults() *TransferEventListRequest {
	this := TransferEventListRequest{}
	var count int32 = 25
	this.Count = *NewNullableInt32(&count)
	var offset int32 = 0
	this.Offset = *NewNullableInt32(&offset)
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransferEventListRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEventListRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransferEventListRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransferEventListRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransferEventListRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEventListRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransferEventListRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransferEventListRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferEventListRequest) GetStartDate() time.Time {
	if o == nil || o.StartDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferEventListRequest) GetStartDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *TransferEventListRequest) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableTime and assigns it to the StartDate field.
func (o *TransferEventListRequest) SetStartDate(v time.Time) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *TransferEventListRequest) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *TransferEventListRequest) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferEventListRequest) GetEndDate() time.Time {
	if o == nil || o.EndDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferEventListRequest) GetEndDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *TransferEventListRequest) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableTime and assigns it to the EndDate field.
func (o *TransferEventListRequest) SetEndDate(v time.Time) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *TransferEventListRequest) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *TransferEventListRequest) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetTransferId returns the TransferId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferEventListRequest) GetTransferId() string {
	if o == nil || o.TransferId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TransferId.Get()
}

// GetTransferIdOk returns a tuple with the TransferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferEventListRequest) GetTransferIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TransferId.Get(), o.TransferId.IsSet()
}

// HasTransferId returns a boolean if a field has been set.
func (o *TransferEventListRequest) HasTransferId() bool {
	if o != nil && o.TransferId.IsSet() {
		return true
	}

	return false
}

// SetTransferId gets a reference to the given NullableString and assigns it to the TransferId field.
func (o *TransferEventListRequest) SetTransferId(v string) {
	o.TransferId.Set(&v)
}
// SetTransferIdNil sets the value for TransferId to be an explicit nil
func (o *TransferEventListRequest) SetTransferIdNil() {
	o.TransferId.Set(nil)
}

// UnsetTransferId ensures that no value is present for TransferId, not even an explicit nil
func (o *TransferEventListRequest) UnsetTransferId() {
	o.TransferId.Unset()
}

// GetAccountId returns the AccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferEventListRequest) GetAccountId() string {
	if o == nil || o.AccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountId.Get()
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferEventListRequest) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountId.Get(), o.AccountId.IsSet()
}

// HasAccountId returns a boolean if a field has been set.
func (o *TransferEventListRequest) HasAccountId() bool {
	if o != nil && o.AccountId.IsSet() {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given NullableString and assigns it to the AccountId field.
func (o *TransferEventListRequest) SetAccountId(v string) {
	o.AccountId.Set(&v)
}
// SetAccountIdNil sets the value for AccountId to be an explicit nil
func (o *TransferEventListRequest) SetAccountIdNil() {
	o.AccountId.Set(nil)
}

// UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
func (o *TransferEventListRequest) UnsetAccountId() {
	o.AccountId.Unset()
}

// GetTransferType returns the TransferType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferEventListRequest) GetTransferType() TransferEventListTransferType {
	if o == nil || o.TransferType.Get() == nil {
		var ret TransferEventListTransferType
		return ret
	}
	return *o.TransferType.Get()
}

// GetTransferTypeOk returns a tuple with the TransferType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferEventListRequest) GetTransferTypeOk() (*TransferEventListTransferType, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TransferType.Get(), o.TransferType.IsSet()
}

// HasTransferType returns a boolean if a field has been set.
func (o *TransferEventListRequest) HasTransferType() bool {
	if o != nil && o.TransferType.IsSet() {
		return true
	}

	return false
}

// SetTransferType gets a reference to the given NullableTransferEventListTransferType and assigns it to the TransferType field.
func (o *TransferEventListRequest) SetTransferType(v TransferEventListTransferType) {
	o.TransferType.Set(&v)
}
// SetTransferTypeNil sets the value for TransferType to be an explicit nil
func (o *TransferEventListRequest) SetTransferTypeNil() {
	o.TransferType.Set(nil)
}

// UnsetTransferType ensures that no value is present for TransferType, not even an explicit nil
func (o *TransferEventListRequest) UnsetTransferType() {
	o.TransferType.Unset()
}

// GetEventTypes returns the EventTypes field value if set, zero value otherwise.
func (o *TransferEventListRequest) GetEventTypes() []TransferEventType {
	if o == nil || o.EventTypes == nil {
		var ret []TransferEventType
		return ret
	}
	return *o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEventListRequest) GetEventTypesOk() (*[]TransferEventType, bool) {
	if o == nil || o.EventTypes == nil {
		return nil, false
	}
	return o.EventTypes, true
}

// HasEventTypes returns a boolean if a field has been set.
func (o *TransferEventListRequest) HasEventTypes() bool {
	if o != nil && o.EventTypes != nil {
		return true
	}

	return false
}

// SetEventTypes gets a reference to the given []TransferEventType and assigns it to the EventTypes field.
func (o *TransferEventListRequest) SetEventTypes(v []TransferEventType) {
	o.EventTypes = &v
}

// GetSweepId returns the SweepId field value if set, zero value otherwise.
func (o *TransferEventListRequest) GetSweepId() string {
	if o == nil || o.SweepId == nil {
		var ret string
		return ret
	}
	return *o.SweepId
}

// GetSweepIdOk returns a tuple with the SweepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEventListRequest) GetSweepIdOk() (*string, bool) {
	if o == nil || o.SweepId == nil {
		return nil, false
	}
	return o.SweepId, true
}

// HasSweepId returns a boolean if a field has been set.
func (o *TransferEventListRequest) HasSweepId() bool {
	if o != nil && o.SweepId != nil {
		return true
	}

	return false
}

// SetSweepId gets a reference to the given string and assigns it to the SweepId field.
func (o *TransferEventListRequest) SetSweepId(v string) {
	o.SweepId = &v
}

// GetCount returns the Count field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferEventListRequest) GetCount() int32 {
	if o == nil || o.Count.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Count.Get()
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferEventListRequest) GetCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Count.Get(), o.Count.IsSet()
}

// HasCount returns a boolean if a field has been set.
func (o *TransferEventListRequest) HasCount() bool {
	if o != nil && o.Count.IsSet() {
		return true
	}

	return false
}

// SetCount gets a reference to the given NullableInt32 and assigns it to the Count field.
func (o *TransferEventListRequest) SetCount(v int32) {
	o.Count.Set(&v)
}
// SetCountNil sets the value for Count to be an explicit nil
func (o *TransferEventListRequest) SetCountNil() {
	o.Count.Set(nil)
}

// UnsetCount ensures that no value is present for Count, not even an explicit nil
func (o *TransferEventListRequest) UnsetCount() {
	o.Count.Unset()
}

// GetOffset returns the Offset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferEventListRequest) GetOffset() int32 {
	if o == nil || o.Offset.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Offset.Get()
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferEventListRequest) GetOffsetOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Offset.Get(), o.Offset.IsSet()
}

// HasOffset returns a boolean if a field has been set.
func (o *TransferEventListRequest) HasOffset() bool {
	if o != nil && o.Offset.IsSet() {
		return true
	}

	return false
}

// SetOffset gets a reference to the given NullableInt32 and assigns it to the Offset field.
func (o *TransferEventListRequest) SetOffset(v int32) {
	o.Offset.Set(&v)
}
// SetOffsetNil sets the value for Offset to be an explicit nil
func (o *TransferEventListRequest) SetOffsetNil() {
	o.Offset.Set(nil)
}

// UnsetOffset ensures that no value is present for Offset, not even an explicit nil
func (o *TransferEventListRequest) UnsetOffset() {
	o.Offset.Unset()
}

// GetOriginationAccountId returns the OriginationAccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferEventListRequest) GetOriginationAccountId() string {
	if o == nil || o.OriginationAccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OriginationAccountId.Get()
}

// GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferEventListRequest) GetOriginationAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginationAccountId.Get(), o.OriginationAccountId.IsSet()
}

// HasOriginationAccountId returns a boolean if a field has been set.
func (o *TransferEventListRequest) HasOriginationAccountId() bool {
	if o != nil && o.OriginationAccountId.IsSet() {
		return true
	}

	return false
}

// SetOriginationAccountId gets a reference to the given NullableString and assigns it to the OriginationAccountId field.
func (o *TransferEventListRequest) SetOriginationAccountId(v string) {
	o.OriginationAccountId.Set(&v)
}
// SetOriginationAccountIdNil sets the value for OriginationAccountId to be an explicit nil
func (o *TransferEventListRequest) SetOriginationAccountIdNil() {
	o.OriginationAccountId.Set(nil)
}

// UnsetOriginationAccountId ensures that no value is present for OriginationAccountId, not even an explicit nil
func (o *TransferEventListRequest) UnsetOriginationAccountId() {
	o.OriginationAccountId.Unset()
}

func (o TransferEventListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.StartDate.IsSet() {
		toSerialize["start_date"] = o.StartDate.Get()
	}
	if o.EndDate.IsSet() {
		toSerialize["end_date"] = o.EndDate.Get()
	}
	if o.TransferId.IsSet() {
		toSerialize["transfer_id"] = o.TransferId.Get()
	}
	if o.AccountId.IsSet() {
		toSerialize["account_id"] = o.AccountId.Get()
	}
	if o.TransferType.IsSet() {
		toSerialize["transfer_type"] = o.TransferType.Get()
	}
	if o.EventTypes != nil {
		toSerialize["event_types"] = o.EventTypes
	}
	if o.SweepId != nil {
		toSerialize["sweep_id"] = o.SweepId
	}
	if o.Count.IsSet() {
		toSerialize["count"] = o.Count.Get()
	}
	if o.Offset.IsSet() {
		toSerialize["offset"] = o.Offset.Get()
	}
	if o.OriginationAccountId.IsSet() {
		toSerialize["origination_account_id"] = o.OriginationAccountId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTransferEventListRequest struct {
	value *TransferEventListRequest
	isSet bool
}

func (v NullableTransferEventListRequest) Get() *TransferEventListRequest {
	return v.value
}

func (v *NullableTransferEventListRequest) Set(val *TransferEventListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferEventListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferEventListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferEventListRequest(val *TransferEventListRequest) *NullableTransferEventListRequest {
	return &NullableTransferEventListRequest{value: val, isSet: true}
}

func (v NullableTransferEventListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferEventListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


