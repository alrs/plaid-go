/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.79.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// WalletTransactionsListRequest WalletTransactionsListRequest defines the request schema for `/wallet/transactions/list`
type WalletTransactionsListRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The ID of the e-wallet to fetch transactions from
	WalletId string `json:"wallet_id"`
	// A base64 value representing the latest transaction that has already been requested. Set this to `next_cursor` received from the previous `/wallet/transactions/list` request. If provided, the response will only contain transactions created before that transaction. If omitted, the response will contain transactions starting from the most recent, and in descending order by the `created_at` time.
	Cursor *string `json:"cursor,omitempty"`
	// The number of transactions to fetch
	Count *int32 `json:"count,omitempty"`
}

// NewWalletTransactionsListRequest instantiates a new WalletTransactionsListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletTransactionsListRequest(walletId string) *WalletTransactionsListRequest {
	this := WalletTransactionsListRequest{}
	this.WalletId = walletId
	var count int32 = 10
	this.Count = &count
	return &this
}

// NewWalletTransactionsListRequestWithDefaults instantiates a new WalletTransactionsListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletTransactionsListRequestWithDefaults() *WalletTransactionsListRequest {
	this := WalletTransactionsListRequest{}
	var count int32 = 10
	this.Count = &count
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *WalletTransactionsListRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletTransactionsListRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *WalletTransactionsListRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *WalletTransactionsListRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *WalletTransactionsListRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletTransactionsListRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *WalletTransactionsListRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *WalletTransactionsListRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetWalletId returns the WalletId field value
func (o *WalletTransactionsListRequest) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionsListRequest) GetWalletIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *WalletTransactionsListRequest) SetWalletId(v string) {
	o.WalletId = v
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *WalletTransactionsListRequest) GetCursor() string {
	if o == nil || o.Cursor == nil {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletTransactionsListRequest) GetCursorOk() (*string, bool) {
	if o == nil || o.Cursor == nil {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *WalletTransactionsListRequest) HasCursor() bool {
	if o != nil && o.Cursor != nil {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *WalletTransactionsListRequest) SetCursor(v string) {
	o.Cursor = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *WalletTransactionsListRequest) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletTransactionsListRequest) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *WalletTransactionsListRequest) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *WalletTransactionsListRequest) SetCount(v int32) {
	o.Count = &v
}

func (o WalletTransactionsListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["wallet_id"] = o.WalletId
	}
	if o.Cursor != nil {
		toSerialize["cursor"] = o.Cursor
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableWalletTransactionsListRequest struct {
	value *WalletTransactionsListRequest
	isSet bool
}

func (v NullableWalletTransactionsListRequest) Get() *WalletTransactionsListRequest {
	return v.value
}

func (v *NullableWalletTransactionsListRequest) Set(val *WalletTransactionsListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletTransactionsListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletTransactionsListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletTransactionsListRequest(val *WalletTransactionsListRequest) *NullableWalletTransactionsListRequest {
	return &NullableWalletTransactionsListRequest{value: val, isSet: true}
}

func (v NullableWalletTransactionsListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletTransactionsListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


