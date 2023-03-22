/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.343.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// WatchlistScreeningHitStatus The current state of review. All watchlist screening hits begin in a `pending_review` state but can be changed by creating a review. When a hit is in the `pending_review` state, it will always show the latest version of the watchlist data Plaid has available and be compared against the latest customer information saved in the watchlist screening. Once a hit has been marked as `confirmed` or `dismissed` it will no longer be updated so that the state is as it was when the review was first conducted.
type WatchlistScreeningHitStatus string

var _ = fmt.Printf

// List of WatchlistScreeningHitStatus
const (
	WATCHLISTSCREENINGHITSTATUS_CONFIRMED WatchlistScreeningHitStatus = "confirmed"
	WATCHLISTSCREENINGHITSTATUS_PENDING_REVIEW WatchlistScreeningHitStatus = "pending_review"
	WATCHLISTSCREENINGHITSTATUS_DISMISSED WatchlistScreeningHitStatus = "dismissed"
)

var allowedWatchlistScreeningHitStatusEnumValues = []WatchlistScreeningHitStatus{
	"confirmed",
	"pending_review",
	"dismissed",
}

func (v *WatchlistScreeningHitStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := WatchlistScreeningHitStatus(value)


	*v = enumTypeValue
	return nil
}

// NewWatchlistScreeningHitStatusFromValue returns a pointer to a valid WatchlistScreeningHitStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWatchlistScreeningHitStatusFromValue(v string) (*WatchlistScreeningHitStatus, error) {
	ev := WatchlistScreeningHitStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WatchlistScreeningHitStatus) IsValid() bool {
	for _, existing := range allowedWatchlistScreeningHitStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WatchlistScreeningHitStatus value
func (v WatchlistScreeningHitStatus) Ptr() *WatchlistScreeningHitStatus {
	return &v
}

type NullableWatchlistScreeningHitStatus struct {
	value *WatchlistScreeningHitStatus
	isSet bool
}

func (v NullableWatchlistScreeningHitStatus) Get() *WatchlistScreeningHitStatus {
	return v.value
}

func (v *NullableWatchlistScreeningHitStatus) Set(val *WatchlistScreeningHitStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningHitStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningHitStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningHitStatus(val *WatchlistScreeningHitStatus) *NullableWatchlistScreeningHitStatus {
	return &NullableWatchlistScreeningHitStatus{value: val, isSet: true}
}

func (v NullableWatchlistScreeningHitStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningHitStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

