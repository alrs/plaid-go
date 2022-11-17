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

// WatchlistScreeningEntityHistoryListResponse Paginated list of entity watchlist screenings
type WatchlistScreeningEntityHistoryListResponse struct {
	// List of entity watchlist screening
	EntityWatchlistScreenings []EntityWatchlistScreening `json:"entity_watchlist_screenings"`
	// An identifier that determines which page of results you receive.
	NextCursor NullableString `json:"next_cursor"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _WatchlistScreeningEntityHistoryListResponse WatchlistScreeningEntityHistoryListResponse

// NewWatchlistScreeningEntityHistoryListResponse instantiates a new WatchlistScreeningEntityHistoryListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningEntityHistoryListResponse(entityWatchlistScreenings []EntityWatchlistScreening, nextCursor NullableString, requestId string) *WatchlistScreeningEntityHistoryListResponse {
	this := WatchlistScreeningEntityHistoryListResponse{}
	this.EntityWatchlistScreenings = entityWatchlistScreenings
	this.NextCursor = nextCursor
	this.RequestId = requestId
	return &this
}

// NewWatchlistScreeningEntityHistoryListResponseWithDefaults instantiates a new WatchlistScreeningEntityHistoryListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningEntityHistoryListResponseWithDefaults() *WatchlistScreeningEntityHistoryListResponse {
	this := WatchlistScreeningEntityHistoryListResponse{}
	return &this
}

// GetEntityWatchlistScreenings returns the EntityWatchlistScreenings field value
func (o *WatchlistScreeningEntityHistoryListResponse) GetEntityWatchlistScreenings() []EntityWatchlistScreening {
	if o == nil {
		var ret []EntityWatchlistScreening
		return ret
	}

	return o.EntityWatchlistScreenings
}

// GetEntityWatchlistScreeningsOk returns a tuple with the EntityWatchlistScreenings field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningEntityHistoryListResponse) GetEntityWatchlistScreeningsOk() (*[]EntityWatchlistScreening, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EntityWatchlistScreenings, true
}

// SetEntityWatchlistScreenings sets field value
func (o *WatchlistScreeningEntityHistoryListResponse) SetEntityWatchlistScreenings(v []EntityWatchlistScreening) {
	o.EntityWatchlistScreenings = v
}

// GetNextCursor returns the NextCursor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchlistScreeningEntityHistoryListResponse) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningEntityHistoryListResponse) GetNextCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// SetNextCursor sets field value
func (o *WatchlistScreeningEntityHistoryListResponse) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}

// GetRequestId returns the RequestId field value
func (o *WatchlistScreeningEntityHistoryListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningEntityHistoryListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *WatchlistScreeningEntityHistoryListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o WatchlistScreeningEntityHistoryListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entity_watchlist_screenings"] = o.EntityWatchlistScreenings
	}
	if true {
		toSerialize["next_cursor"] = o.NextCursor.Get()
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WatchlistScreeningEntityHistoryListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWatchlistScreeningEntityHistoryListResponse := _WatchlistScreeningEntityHistoryListResponse{}

	if err = json.Unmarshal(bytes, &varWatchlistScreeningEntityHistoryListResponse); err == nil {
		*o = WatchlistScreeningEntityHistoryListResponse(varWatchlistScreeningEntityHistoryListResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "entity_watchlist_screenings")
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWatchlistScreeningEntityHistoryListResponse struct {
	value *WatchlistScreeningEntityHistoryListResponse
	isSet bool
}

func (v NullableWatchlistScreeningEntityHistoryListResponse) Get() *WatchlistScreeningEntityHistoryListResponse {
	return v.value
}

func (v *NullableWatchlistScreeningEntityHistoryListResponse) Set(val *WatchlistScreeningEntityHistoryListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningEntityHistoryListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningEntityHistoryListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningEntityHistoryListResponse(val *WatchlistScreeningEntityHistoryListResponse) *NullableWatchlistScreeningEntityHistoryListResponse {
	return &NullableWatchlistScreeningEntityHistoryListResponse{value: val, isSet: true}
}

func (v NullableWatchlistScreeningEntityHistoryListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningEntityHistoryListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


