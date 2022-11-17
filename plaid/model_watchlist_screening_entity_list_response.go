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

// WatchlistScreeningEntityListResponse Paginated list of entity watchlist screenings
type WatchlistScreeningEntityListResponse struct {
	// List of entity watchlist screening
	EntityWatchlistScreenings []EntityWatchlistScreening `json:"entity_watchlist_screenings"`
	// An identifier that determines which page of results you receive.
	NextCursor NullableString `json:"next_cursor"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _WatchlistScreeningEntityListResponse WatchlistScreeningEntityListResponse

// NewWatchlistScreeningEntityListResponse instantiates a new WatchlistScreeningEntityListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningEntityListResponse(entityWatchlistScreenings []EntityWatchlistScreening, nextCursor NullableString, requestId string) *WatchlistScreeningEntityListResponse {
	this := WatchlistScreeningEntityListResponse{}
	this.EntityWatchlistScreenings = entityWatchlistScreenings
	this.NextCursor = nextCursor
	this.RequestId = requestId
	return &this
}

// NewWatchlistScreeningEntityListResponseWithDefaults instantiates a new WatchlistScreeningEntityListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningEntityListResponseWithDefaults() *WatchlistScreeningEntityListResponse {
	this := WatchlistScreeningEntityListResponse{}
	return &this
}

// GetEntityWatchlistScreenings returns the EntityWatchlistScreenings field value
func (o *WatchlistScreeningEntityListResponse) GetEntityWatchlistScreenings() []EntityWatchlistScreening {
	if o == nil {
		var ret []EntityWatchlistScreening
		return ret
	}

	return o.EntityWatchlistScreenings
}

// GetEntityWatchlistScreeningsOk returns a tuple with the EntityWatchlistScreenings field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningEntityListResponse) GetEntityWatchlistScreeningsOk() (*[]EntityWatchlistScreening, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EntityWatchlistScreenings, true
}

// SetEntityWatchlistScreenings sets field value
func (o *WatchlistScreeningEntityListResponse) SetEntityWatchlistScreenings(v []EntityWatchlistScreening) {
	o.EntityWatchlistScreenings = v
}

// GetNextCursor returns the NextCursor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchlistScreeningEntityListResponse) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningEntityListResponse) GetNextCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// SetNextCursor sets field value
func (o *WatchlistScreeningEntityListResponse) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}

// GetRequestId returns the RequestId field value
func (o *WatchlistScreeningEntityListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningEntityListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *WatchlistScreeningEntityListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o WatchlistScreeningEntityListResponse) MarshalJSON() ([]byte, error) {
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

func (o *WatchlistScreeningEntityListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWatchlistScreeningEntityListResponse := _WatchlistScreeningEntityListResponse{}

	if err = json.Unmarshal(bytes, &varWatchlistScreeningEntityListResponse); err == nil {
		*o = WatchlistScreeningEntityListResponse(varWatchlistScreeningEntityListResponse)
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

type NullableWatchlistScreeningEntityListResponse struct {
	value *WatchlistScreeningEntityListResponse
	isSet bool
}

func (v NullableWatchlistScreeningEntityListResponse) Get() *WatchlistScreeningEntityListResponse {
	return v.value
}

func (v *NullableWatchlistScreeningEntityListResponse) Set(val *WatchlistScreeningEntityListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningEntityListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningEntityListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningEntityListResponse(val *WatchlistScreeningEntityListResponse) *NullableWatchlistScreeningEntityListResponse {
	return &NullableWatchlistScreeningEntityListResponse{value: val, isSet: true}
}

func (v NullableWatchlistScreeningEntityListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningEntityListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


