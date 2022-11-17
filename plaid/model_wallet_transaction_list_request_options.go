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
	"time"
)

// WalletTransactionListRequestOptions Additional wallet transaction options
type WalletTransactionListRequestOptions struct {
	// Timestamp in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DDThh:mm:ssZ) for filtering transactions, inclusive of the provided date.
	StartTime *time.Time `json:"start_time,omitempty"`
	// Timestamp in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DDThh:mm:ssZ) for filtering transactions, inclusive of the provided date.
	EndTime *time.Time `json:"end_time,omitempty"`
}

// NewWalletTransactionListRequestOptions instantiates a new WalletTransactionListRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletTransactionListRequestOptions() *WalletTransactionListRequestOptions {
	this := WalletTransactionListRequestOptions{}
	return &this
}

// NewWalletTransactionListRequestOptionsWithDefaults instantiates a new WalletTransactionListRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletTransactionListRequestOptionsWithDefaults() *WalletTransactionListRequestOptions {
	this := WalletTransactionListRequestOptions{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *WalletTransactionListRequestOptions) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletTransactionListRequestOptions) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *WalletTransactionListRequestOptions) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *WalletTransactionListRequestOptions) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *WalletTransactionListRequestOptions) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletTransactionListRequestOptions) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *WalletTransactionListRequestOptions) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *WalletTransactionListRequestOptions) SetEndTime(v time.Time) {
	o.EndTime = &v
}

func (o WalletTransactionListRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}
	if o.EndTime != nil {
		toSerialize["end_time"] = o.EndTime
	}
	return json.Marshal(toSerialize)
}

type NullableWalletTransactionListRequestOptions struct {
	value *WalletTransactionListRequestOptions
	isSet bool
}

func (v NullableWalletTransactionListRequestOptions) Get() *WalletTransactionListRequestOptions {
	return v.value
}

func (v *NullableWalletTransactionListRequestOptions) Set(val *WalletTransactionListRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletTransactionListRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletTransactionListRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletTransactionListRequestOptions(val *WalletTransactionListRequestOptions) *NullableWalletTransactionListRequestOptions {
	return &NullableWalletTransactionListRequestOptions{value: val, isSet: true}
}

func (v NullableWalletTransactionListRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletTransactionListRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


