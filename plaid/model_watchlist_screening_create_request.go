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

// WatchlistScreeningCreateRequest Request input for creating an individual watchlist screening
type WatchlistScreeningCreateRequest struct {
	SearchTerms WatchlistScreeningRequestSearchTerms `json:"search_terms"`
	ClientUserId NullableString `json:"client_user_id,omitempty"`
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
}

// NewWatchlistScreeningCreateRequest instantiates a new WatchlistScreeningCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningCreateRequest(searchTerms WatchlistScreeningRequestSearchTerms) *WatchlistScreeningCreateRequest {
	this := WatchlistScreeningCreateRequest{}
	this.SearchTerms = searchTerms
	return &this
}

// NewWatchlistScreeningCreateRequestWithDefaults instantiates a new WatchlistScreeningCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningCreateRequestWithDefaults() *WatchlistScreeningCreateRequest {
	this := WatchlistScreeningCreateRequest{}
	return &this
}

// GetSearchTerms returns the SearchTerms field value
func (o *WatchlistScreeningCreateRequest) GetSearchTerms() WatchlistScreeningRequestSearchTerms {
	if o == nil {
		var ret WatchlistScreeningRequestSearchTerms
		return ret
	}

	return o.SearchTerms
}

// GetSearchTermsOk returns a tuple with the SearchTerms field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningCreateRequest) GetSearchTermsOk() (*WatchlistScreeningRequestSearchTerms, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SearchTerms, true
}

// SetSearchTerms sets field value
func (o *WatchlistScreeningCreateRequest) SetSearchTerms(v WatchlistScreeningRequestSearchTerms) {
	o.SearchTerms = v
}

// GetClientUserId returns the ClientUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WatchlistScreeningCreateRequest) GetClientUserId() string {
	if o == nil || o.ClientUserId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientUserId.Get()
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningCreateRequest) GetClientUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientUserId.Get(), o.ClientUserId.IsSet()
}

// HasClientUserId returns a boolean if a field has been set.
func (o *WatchlistScreeningCreateRequest) HasClientUserId() bool {
	if o != nil && o.ClientUserId.IsSet() {
		return true
	}

	return false
}

// SetClientUserId gets a reference to the given NullableString and assigns it to the ClientUserId field.
func (o *WatchlistScreeningCreateRequest) SetClientUserId(v string) {
	o.ClientUserId.Set(&v)
}
// SetClientUserIdNil sets the value for ClientUserId to be an explicit nil
func (o *WatchlistScreeningCreateRequest) SetClientUserIdNil() {
	o.ClientUserId.Set(nil)
}

// UnsetClientUserId ensures that no value is present for ClientUserId, not even an explicit nil
func (o *WatchlistScreeningCreateRequest) UnsetClientUserId() {
	o.ClientUserId.Unset()
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *WatchlistScreeningCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *WatchlistScreeningCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *WatchlistScreeningCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *WatchlistScreeningCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *WatchlistScreeningCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *WatchlistScreeningCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

func (o WatchlistScreeningCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["search_terms"] = o.SearchTerms
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
	return json.Marshal(toSerialize)
}

type NullableWatchlistScreeningCreateRequest struct {
	value *WatchlistScreeningCreateRequest
	isSet bool
}

func (v NullableWatchlistScreeningCreateRequest) Get() *WatchlistScreeningCreateRequest {
	return v.value
}

func (v *NullableWatchlistScreeningCreateRequest) Set(val *WatchlistScreeningCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningCreateRequest(val *WatchlistScreeningCreateRequest) *NullableWatchlistScreeningCreateRequest {
	return &NullableWatchlistScreeningCreateRequest{value: val, isSet: true}
}

func (v NullableWatchlistScreeningCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


