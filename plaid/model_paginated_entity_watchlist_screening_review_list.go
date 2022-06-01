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

// PaginatedEntityWatchlistScreeningReviewList Paginated list of entity watchlist screening reviews
type PaginatedEntityWatchlistScreeningReviewList struct {
	// List of entity watchlist screening reviews
	EntityWatchlistScreeningReviews []EntityWatchlistScreeningReview `json:"entity_watchlist_screening_reviews"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	// An identifier that determines which page of results you receive.
	NextCursor NullableString `json:"next_cursor"`
}

// NewPaginatedEntityWatchlistScreeningReviewList instantiates a new PaginatedEntityWatchlistScreeningReviewList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedEntityWatchlistScreeningReviewList(entityWatchlistScreeningReviews []EntityWatchlistScreeningReview, requestId string, nextCursor NullableString) *PaginatedEntityWatchlistScreeningReviewList {
	this := PaginatedEntityWatchlistScreeningReviewList{}
	this.EntityWatchlistScreeningReviews = entityWatchlistScreeningReviews
	this.RequestId = requestId
	this.NextCursor = nextCursor
	return &this
}

// NewPaginatedEntityWatchlistScreeningReviewListWithDefaults instantiates a new PaginatedEntityWatchlistScreeningReviewList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedEntityWatchlistScreeningReviewListWithDefaults() *PaginatedEntityWatchlistScreeningReviewList {
	this := PaginatedEntityWatchlistScreeningReviewList{}
	return &this
}

// GetEntityWatchlistScreeningReviews returns the EntityWatchlistScreeningReviews field value
func (o *PaginatedEntityWatchlistScreeningReviewList) GetEntityWatchlistScreeningReviews() []EntityWatchlistScreeningReview {
	if o == nil {
		var ret []EntityWatchlistScreeningReview
		return ret
	}

	return o.EntityWatchlistScreeningReviews
}

// GetEntityWatchlistScreeningReviewsOk returns a tuple with the EntityWatchlistScreeningReviews field value
// and a boolean to check if the value has been set.
func (o *PaginatedEntityWatchlistScreeningReviewList) GetEntityWatchlistScreeningReviewsOk() (*[]EntityWatchlistScreeningReview, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EntityWatchlistScreeningReviews, true
}

// SetEntityWatchlistScreeningReviews sets field value
func (o *PaginatedEntityWatchlistScreeningReviewList) SetEntityWatchlistScreeningReviews(v []EntityWatchlistScreeningReview) {
	o.EntityWatchlistScreeningReviews = v
}

// GetRequestId returns the RequestId field value
func (o *PaginatedEntityWatchlistScreeningReviewList) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *PaginatedEntityWatchlistScreeningReviewList) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *PaginatedEntityWatchlistScreeningReviewList) SetRequestId(v string) {
	o.RequestId = v
}

// GetNextCursor returns the NextCursor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaginatedEntityWatchlistScreeningReviewList) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedEntityWatchlistScreeningReviewList) GetNextCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// SetNextCursor sets field value
func (o *PaginatedEntityWatchlistScreeningReviewList) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}

func (o PaginatedEntityWatchlistScreeningReviewList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entity_watchlist_screening_reviews"] = o.EntityWatchlistScreeningReviews
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}
	if true {
		toSerialize["next_cursor"] = o.NextCursor.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedEntityWatchlistScreeningReviewList struct {
	value *PaginatedEntityWatchlistScreeningReviewList
	isSet bool
}

func (v NullablePaginatedEntityWatchlistScreeningReviewList) Get() *PaginatedEntityWatchlistScreeningReviewList {
	return v.value
}

func (v *NullablePaginatedEntityWatchlistScreeningReviewList) Set(val *PaginatedEntityWatchlistScreeningReviewList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedEntityWatchlistScreeningReviewList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedEntityWatchlistScreeningReviewList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedEntityWatchlistScreeningReviewList(val *PaginatedEntityWatchlistScreeningReviewList) *NullablePaginatedEntityWatchlistScreeningReviewList {
	return &NullablePaginatedEntityWatchlistScreeningReviewList{value: val, isSet: true}
}

func (v NullablePaginatedEntityWatchlistScreeningReviewList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedEntityWatchlistScreeningReviewList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


