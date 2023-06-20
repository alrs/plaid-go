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

// WatchlistScreeningIndividualReviewCreateRequest Request input for creating a screening review
type WatchlistScreeningIndividualReviewCreateRequest struct {
	// Hits to mark as a true positive after thorough manual review. These hits will never recur or be updated once dismissed. In most cases, confirmed hits indicate that the customer should be rejected.
	ConfirmedHits []string `json:"confirmed_hits"`
	// Hits to mark as a false positive after thorough manual review. These hits will never recur or be updated once dismissed.
	DismissedHits []string `json:"dismissed_hits"`
	// A comment submitted by a team member as part of reviewing a watchlist screening.
	Comment NullableString `json:"comment,omitempty"`
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// ID of the associated screening.
	WatchlistScreeningId string `json:"watchlist_screening_id"`
}

// NewWatchlistScreeningIndividualReviewCreateRequest instantiates a new WatchlistScreeningIndividualReviewCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningIndividualReviewCreateRequest(confirmedHits []string, dismissedHits []string, watchlistScreeningId string) *WatchlistScreeningIndividualReviewCreateRequest {
	this := WatchlistScreeningIndividualReviewCreateRequest{}
	this.ConfirmedHits = confirmedHits
	this.DismissedHits = dismissedHits
	this.WatchlistScreeningId = watchlistScreeningId
	return &this
}

// NewWatchlistScreeningIndividualReviewCreateRequestWithDefaults instantiates a new WatchlistScreeningIndividualReviewCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningIndividualReviewCreateRequestWithDefaults() *WatchlistScreeningIndividualReviewCreateRequest {
	this := WatchlistScreeningIndividualReviewCreateRequest{}
	return &this
}

// GetConfirmedHits returns the ConfirmedHits field value
func (o *WatchlistScreeningIndividualReviewCreateRequest) GetConfirmedHits() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ConfirmedHits
}

// GetConfirmedHitsOk returns a tuple with the ConfirmedHits field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualReviewCreateRequest) GetConfirmedHitsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfirmedHits, true
}

// SetConfirmedHits sets field value
func (o *WatchlistScreeningIndividualReviewCreateRequest) SetConfirmedHits(v []string) {
	o.ConfirmedHits = v
}

// GetDismissedHits returns the DismissedHits field value
func (o *WatchlistScreeningIndividualReviewCreateRequest) GetDismissedHits() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DismissedHits
}

// GetDismissedHitsOk returns a tuple with the DismissedHits field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualReviewCreateRequest) GetDismissedHitsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DismissedHits, true
}

// SetDismissedHits sets field value
func (o *WatchlistScreeningIndividualReviewCreateRequest) SetDismissedHits(v []string) {
	o.DismissedHits = v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WatchlistScreeningIndividualReviewCreateRequest) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningIndividualReviewCreateRequest) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *WatchlistScreeningIndividualReviewCreateRequest) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *WatchlistScreeningIndividualReviewCreateRequest) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *WatchlistScreeningIndividualReviewCreateRequest) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *WatchlistScreeningIndividualReviewCreateRequest) UnsetComment() {
	o.Comment.Unset()
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *WatchlistScreeningIndividualReviewCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualReviewCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *WatchlistScreeningIndividualReviewCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *WatchlistScreeningIndividualReviewCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *WatchlistScreeningIndividualReviewCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualReviewCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *WatchlistScreeningIndividualReviewCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *WatchlistScreeningIndividualReviewCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetWatchlistScreeningId returns the WatchlistScreeningId field value
func (o *WatchlistScreeningIndividualReviewCreateRequest) GetWatchlistScreeningId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WatchlistScreeningId
}

// GetWatchlistScreeningIdOk returns a tuple with the WatchlistScreeningId field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualReviewCreateRequest) GetWatchlistScreeningIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WatchlistScreeningId, true
}

// SetWatchlistScreeningId sets field value
func (o *WatchlistScreeningIndividualReviewCreateRequest) SetWatchlistScreeningId(v string) {
	o.WatchlistScreeningId = v
}

func (o WatchlistScreeningIndividualReviewCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["confirmed_hits"] = o.ConfirmedHits
	}
	if true {
		toSerialize["dismissed_hits"] = o.DismissedHits
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["watchlist_screening_id"] = o.WatchlistScreeningId
	}
	return json.Marshal(toSerialize)
}

type NullableWatchlistScreeningIndividualReviewCreateRequest struct {
	value *WatchlistScreeningIndividualReviewCreateRequest
	isSet bool
}

func (v NullableWatchlistScreeningIndividualReviewCreateRequest) Get() *WatchlistScreeningIndividualReviewCreateRequest {
	return v.value
}

func (v *NullableWatchlistScreeningIndividualReviewCreateRequest) Set(val *WatchlistScreeningIndividualReviewCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningIndividualReviewCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningIndividualReviewCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningIndividualReviewCreateRequest(val *WatchlistScreeningIndividualReviewCreateRequest) *NullableWatchlistScreeningIndividualReviewCreateRequest {
	return &NullableWatchlistScreeningIndividualReviewCreateRequest{value: val, isSet: true}
}

func (v NullableWatchlistScreeningIndividualReviewCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningIndividualReviewCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


