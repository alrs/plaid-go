/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.121.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// UpdateIndividualScreeningRequest Request input for editing an individual watchlist screening
type UpdateIndividualScreeningRequest struct {
	// ID of the associated screening.
	WatchlistScreeningId string `json:"watchlist_screening_id"`
	SearchTerms NullableUpdateIndividualScreeningRequestSearchTerms `json:"search_terms,omitempty"`
	Assignee NullableString `json:"assignee,omitempty"`
	Status NullableWatchlistScreeningStatus `json:"status,omitempty"`
	ClientUserId NullableString `json:"client_user_id,omitempty"`
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// A list of fields to reset back to null
	ResetFields []UpdateIndividualScreeningRequestResettableField `json:"reset_fields,omitempty"`
}

// NewUpdateIndividualScreeningRequest instantiates a new UpdateIndividualScreeningRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateIndividualScreeningRequest(watchlistScreeningId string) *UpdateIndividualScreeningRequest {
	this := UpdateIndividualScreeningRequest{}
	this.WatchlistScreeningId = watchlistScreeningId
	return &this
}

// NewUpdateIndividualScreeningRequestWithDefaults instantiates a new UpdateIndividualScreeningRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateIndividualScreeningRequestWithDefaults() *UpdateIndividualScreeningRequest {
	this := UpdateIndividualScreeningRequest{}
	return &this
}

// GetWatchlistScreeningId returns the WatchlistScreeningId field value
func (o *UpdateIndividualScreeningRequest) GetWatchlistScreeningId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WatchlistScreeningId
}

// GetWatchlistScreeningIdOk returns a tuple with the WatchlistScreeningId field value
// and a boolean to check if the value has been set.
func (o *UpdateIndividualScreeningRequest) GetWatchlistScreeningIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WatchlistScreeningId, true
}

// SetWatchlistScreeningId sets field value
func (o *UpdateIndividualScreeningRequest) SetWatchlistScreeningId(v string) {
	o.WatchlistScreeningId = v
}

// GetSearchTerms returns the SearchTerms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateIndividualScreeningRequest) GetSearchTerms() UpdateIndividualScreeningRequestSearchTerms {
	if o == nil || o.SearchTerms.Get() == nil {
		var ret UpdateIndividualScreeningRequestSearchTerms
		return ret
	}
	return *o.SearchTerms.Get()
}

// GetSearchTermsOk returns a tuple with the SearchTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateIndividualScreeningRequest) GetSearchTermsOk() (*UpdateIndividualScreeningRequestSearchTerms, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SearchTerms.Get(), o.SearchTerms.IsSet()
}

// HasSearchTerms returns a boolean if a field has been set.
func (o *UpdateIndividualScreeningRequest) HasSearchTerms() bool {
	if o != nil && o.SearchTerms.IsSet() {
		return true
	}

	return false
}

// SetSearchTerms gets a reference to the given NullableUpdateIndividualScreeningRequestSearchTerms and assigns it to the SearchTerms field.
func (o *UpdateIndividualScreeningRequest) SetSearchTerms(v UpdateIndividualScreeningRequestSearchTerms) {
	o.SearchTerms.Set(&v)
}
// SetSearchTermsNil sets the value for SearchTerms to be an explicit nil
func (o *UpdateIndividualScreeningRequest) SetSearchTermsNil() {
	o.SearchTerms.Set(nil)
}

// UnsetSearchTerms ensures that no value is present for SearchTerms, not even an explicit nil
func (o *UpdateIndividualScreeningRequest) UnsetSearchTerms() {
	o.SearchTerms.Unset()
}

// GetAssignee returns the Assignee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateIndividualScreeningRequest) GetAssignee() string {
	if o == nil || o.Assignee.Get() == nil {
		var ret string
		return ret
	}
	return *o.Assignee.Get()
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateIndividualScreeningRequest) GetAssigneeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Assignee.Get(), o.Assignee.IsSet()
}

// HasAssignee returns a boolean if a field has been set.
func (o *UpdateIndividualScreeningRequest) HasAssignee() bool {
	if o != nil && o.Assignee.IsSet() {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given NullableString and assigns it to the Assignee field.
func (o *UpdateIndividualScreeningRequest) SetAssignee(v string) {
	o.Assignee.Set(&v)
}
// SetAssigneeNil sets the value for Assignee to be an explicit nil
func (o *UpdateIndividualScreeningRequest) SetAssigneeNil() {
	o.Assignee.Set(nil)
}

// UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
func (o *UpdateIndividualScreeningRequest) UnsetAssignee() {
	o.Assignee.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateIndividualScreeningRequest) GetStatus() WatchlistScreeningStatus {
	if o == nil || o.Status.Get() == nil {
		var ret WatchlistScreeningStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateIndividualScreeningRequest) GetStatusOk() (*WatchlistScreeningStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateIndividualScreeningRequest) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableWatchlistScreeningStatus and assigns it to the Status field.
func (o *UpdateIndividualScreeningRequest) SetStatus(v WatchlistScreeningStatus) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *UpdateIndividualScreeningRequest) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *UpdateIndividualScreeningRequest) UnsetStatus() {
	o.Status.Unset()
}

// GetClientUserId returns the ClientUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateIndividualScreeningRequest) GetClientUserId() string {
	if o == nil || o.ClientUserId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientUserId.Get()
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateIndividualScreeningRequest) GetClientUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientUserId.Get(), o.ClientUserId.IsSet()
}

// HasClientUserId returns a boolean if a field has been set.
func (o *UpdateIndividualScreeningRequest) HasClientUserId() bool {
	if o != nil && o.ClientUserId.IsSet() {
		return true
	}

	return false
}

// SetClientUserId gets a reference to the given NullableString and assigns it to the ClientUserId field.
func (o *UpdateIndividualScreeningRequest) SetClientUserId(v string) {
	o.ClientUserId.Set(&v)
}
// SetClientUserIdNil sets the value for ClientUserId to be an explicit nil
func (o *UpdateIndividualScreeningRequest) SetClientUserIdNil() {
	o.ClientUserId.Set(nil)
}

// UnsetClientUserId ensures that no value is present for ClientUserId, not even an explicit nil
func (o *UpdateIndividualScreeningRequest) UnsetClientUserId() {
	o.ClientUserId.Unset()
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *UpdateIndividualScreeningRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIndividualScreeningRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *UpdateIndividualScreeningRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *UpdateIndividualScreeningRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *UpdateIndividualScreeningRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIndividualScreeningRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *UpdateIndividualScreeningRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *UpdateIndividualScreeningRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetResetFields returns the ResetFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateIndividualScreeningRequest) GetResetFields() []UpdateIndividualScreeningRequestResettableField {
	if o == nil  {
		var ret []UpdateIndividualScreeningRequestResettableField
		return ret
	}
	return o.ResetFields
}

// GetResetFieldsOk returns a tuple with the ResetFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateIndividualScreeningRequest) GetResetFieldsOk() (*[]UpdateIndividualScreeningRequestResettableField, bool) {
	if o == nil || o.ResetFields == nil {
		return nil, false
	}
	return &o.ResetFields, true
}

// HasResetFields returns a boolean if a field has been set.
func (o *UpdateIndividualScreeningRequest) HasResetFields() bool {
	if o != nil && o.ResetFields != nil {
		return true
	}

	return false
}

// SetResetFields gets a reference to the given []UpdateIndividualScreeningRequestResettableField and assigns it to the ResetFields field.
func (o *UpdateIndividualScreeningRequest) SetResetFields(v []UpdateIndividualScreeningRequestResettableField) {
	o.ResetFields = v
}

func (o UpdateIndividualScreeningRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["watchlist_screening_id"] = o.WatchlistScreeningId
	}
	if o.SearchTerms.IsSet() {
		toSerialize["search_terms"] = o.SearchTerms.Get()
	}
	if o.Assignee.IsSet() {
		toSerialize["assignee"] = o.Assignee.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.ClientUserId.IsSet() {
		toSerialize["client_user_id"] = o.ClientUserId.Get()
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.ResetFields != nil {
		toSerialize["reset_fields"] = o.ResetFields
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateIndividualScreeningRequest struct {
	value *UpdateIndividualScreeningRequest
	isSet bool
}

func (v NullableUpdateIndividualScreeningRequest) Get() *UpdateIndividualScreeningRequest {
	return v.value
}

func (v *NullableUpdateIndividualScreeningRequest) Set(val *UpdateIndividualScreeningRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateIndividualScreeningRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateIndividualScreeningRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateIndividualScreeningRequest(val *UpdateIndividualScreeningRequest) *NullableUpdateIndividualScreeningRequest {
	return &NullableUpdateIndividualScreeningRequest{value: val, isSet: true}
}

func (v NullableUpdateIndividualScreeningRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateIndividualScreeningRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


