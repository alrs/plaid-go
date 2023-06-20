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

// WatchlistScreeningEntityProgramListResponse Paginated list of entity watchlist screening programs
type WatchlistScreeningEntityProgramListResponse struct {
	// List of entity watchlist screening programs
	EntityWatchlistPrograms []EntityWatchlistProgram `json:"entity_watchlist_programs"`
	// An identifier that determines which page of results you receive.
	NextCursor NullableString `json:"next_cursor"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _WatchlistScreeningEntityProgramListResponse WatchlistScreeningEntityProgramListResponse

// NewWatchlistScreeningEntityProgramListResponse instantiates a new WatchlistScreeningEntityProgramListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningEntityProgramListResponse(entityWatchlistPrograms []EntityWatchlistProgram, nextCursor NullableString, requestId string) *WatchlistScreeningEntityProgramListResponse {
	this := WatchlistScreeningEntityProgramListResponse{}
	this.EntityWatchlistPrograms = entityWatchlistPrograms
	this.NextCursor = nextCursor
	this.RequestId = requestId
	return &this
}

// NewWatchlistScreeningEntityProgramListResponseWithDefaults instantiates a new WatchlistScreeningEntityProgramListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningEntityProgramListResponseWithDefaults() *WatchlistScreeningEntityProgramListResponse {
	this := WatchlistScreeningEntityProgramListResponse{}
	return &this
}

// GetEntityWatchlistPrograms returns the EntityWatchlistPrograms field value
func (o *WatchlistScreeningEntityProgramListResponse) GetEntityWatchlistPrograms() []EntityWatchlistProgram {
	if o == nil {
		var ret []EntityWatchlistProgram
		return ret
	}

	return o.EntityWatchlistPrograms
}

// GetEntityWatchlistProgramsOk returns a tuple with the EntityWatchlistPrograms field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningEntityProgramListResponse) GetEntityWatchlistProgramsOk() (*[]EntityWatchlistProgram, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EntityWatchlistPrograms, true
}

// SetEntityWatchlistPrograms sets field value
func (o *WatchlistScreeningEntityProgramListResponse) SetEntityWatchlistPrograms(v []EntityWatchlistProgram) {
	o.EntityWatchlistPrograms = v
}

// GetNextCursor returns the NextCursor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchlistScreeningEntityProgramListResponse) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningEntityProgramListResponse) GetNextCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// SetNextCursor sets field value
func (o *WatchlistScreeningEntityProgramListResponse) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}

// GetRequestId returns the RequestId field value
func (o *WatchlistScreeningEntityProgramListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningEntityProgramListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *WatchlistScreeningEntityProgramListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o WatchlistScreeningEntityProgramListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entity_watchlist_programs"] = o.EntityWatchlistPrograms
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

func (o *WatchlistScreeningEntityProgramListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWatchlistScreeningEntityProgramListResponse := _WatchlistScreeningEntityProgramListResponse{}

	if err = json.Unmarshal(bytes, &varWatchlistScreeningEntityProgramListResponse); err == nil {
		*o = WatchlistScreeningEntityProgramListResponse(varWatchlistScreeningEntityProgramListResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "entity_watchlist_programs")
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWatchlistScreeningEntityProgramListResponse struct {
	value *WatchlistScreeningEntityProgramListResponse
	isSet bool
}

func (v NullableWatchlistScreeningEntityProgramListResponse) Get() *WatchlistScreeningEntityProgramListResponse {
	return v.value
}

func (v *NullableWatchlistScreeningEntityProgramListResponse) Set(val *WatchlistScreeningEntityProgramListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningEntityProgramListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningEntityProgramListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningEntityProgramListResponse(val *WatchlistScreeningEntityProgramListResponse) *NullableWatchlistScreeningEntityProgramListResponse {
	return &NullableWatchlistScreeningEntityProgramListResponse{value: val, isSet: true}
}

func (v NullableWatchlistScreeningEntityProgramListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningEntityProgramListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


