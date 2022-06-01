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

// WalletTransactionsListResponse WalletTransactionsListResponse defines the response schema for `/wallet/transactions/list`
type WalletTransactionsListResponse struct {
	// An array of transactions of an e-wallet, associated with the given `wallet_id`
	Transactions []WalletTransaction `json:"transactions"`
	// Cursor used for fetching transactions created before the latest transaction provided in this response
	NextCursor *string `json:"next_cursor,omitempty"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _WalletTransactionsListResponse WalletTransactionsListResponse

// NewWalletTransactionsListResponse instantiates a new WalletTransactionsListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletTransactionsListResponse(transactions []WalletTransaction, requestId string) *WalletTransactionsListResponse {
	this := WalletTransactionsListResponse{}
	this.Transactions = transactions
	this.RequestId = requestId
	return &this
}

// NewWalletTransactionsListResponseWithDefaults instantiates a new WalletTransactionsListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletTransactionsListResponseWithDefaults() *WalletTransactionsListResponse {
	this := WalletTransactionsListResponse{}
	return &this
}

// GetTransactions returns the Transactions field value
func (o *WalletTransactionsListResponse) GetTransactions() []WalletTransaction {
	if o == nil {
		var ret []WalletTransaction
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionsListResponse) GetTransactionsOk() (*[]WalletTransaction, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Transactions, true
}

// SetTransactions sets field value
func (o *WalletTransactionsListResponse) SetTransactions(v []WalletTransaction) {
	o.Transactions = v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *WalletTransactionsListResponse) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletTransactionsListResponse) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *WalletTransactionsListResponse) HasNextCursor() bool {
	if o != nil && o.NextCursor != nil {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *WalletTransactionsListResponse) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetRequestId returns the RequestId field value
func (o *WalletTransactionsListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionsListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *WalletTransactionsListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o WalletTransactionsListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transactions"] = o.Transactions
	}
	if o.NextCursor != nil {
		toSerialize["next_cursor"] = o.NextCursor
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WalletTransactionsListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWalletTransactionsListResponse := _WalletTransactionsListResponse{}

	if err = json.Unmarshal(bytes, &varWalletTransactionsListResponse); err == nil {
		*o = WalletTransactionsListResponse(varWalletTransactionsListResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "transactions")
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWalletTransactionsListResponse struct {
	value *WalletTransactionsListResponse
	isSet bool
}

func (v NullableWalletTransactionsListResponse) Get() *WalletTransactionsListResponse {
	return v.value
}

func (v *NullableWalletTransactionsListResponse) Set(val *WalletTransactionsListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletTransactionsListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletTransactionsListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletTransactionsListResponse(val *WalletTransactionsListResponse) *NullableWalletTransactionsListResponse {
	return &NullableWalletTransactionsListResponse{value: val, isSet: true}
}

func (v NullableWalletTransactionsListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletTransactionsListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


