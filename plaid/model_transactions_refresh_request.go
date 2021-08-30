/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.26.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TransactionsRefreshRequest TransactionsRefreshRequest defines the request schema for `/transactions/refresh`
type TransactionsRefreshRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
}

// NewTransactionsRefreshRequest instantiates a new TransactionsRefreshRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsRefreshRequest(accessToken string) *TransactionsRefreshRequest {
	this := TransactionsRefreshRequest{}
	this.AccessToken = accessToken
	return &this
}

// NewTransactionsRefreshRequestWithDefaults instantiates a new TransactionsRefreshRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsRefreshRequestWithDefaults() *TransactionsRefreshRequest {
	this := TransactionsRefreshRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransactionsRefreshRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsRefreshRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransactionsRefreshRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransactionsRefreshRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetAccessToken returns the AccessToken field value
func (o *TransactionsRefreshRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *TransactionsRefreshRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *TransactionsRefreshRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransactionsRefreshRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsRefreshRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransactionsRefreshRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransactionsRefreshRequest) SetSecret(v string) {
	o.Secret = &v
}

func (o TransactionsRefreshRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionsRefreshRequest struct {
	value *TransactionsRefreshRequest
	isSet bool
}

func (v NullableTransactionsRefreshRequest) Get() *TransactionsRefreshRequest {
	return v.value
}

func (v *NullableTransactionsRefreshRequest) Set(val *TransactionsRefreshRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsRefreshRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsRefreshRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsRefreshRequest(val *TransactionsRefreshRequest) *NullableTransactionsRefreshRequest {
	return &NullableTransactionsRefreshRequest{value: val, isSet: true}
}

func (v NullableTransactionsRefreshRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsRefreshRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


