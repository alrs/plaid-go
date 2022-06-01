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
	"fmt"
)

// WatchlistScreeningStatus A status enum indicating whether a screening is still pending review, has been rejected, or has been cleared.
type WatchlistScreeningStatus string

// List of WatchlistScreeningStatus
const (
	WATCHLISTSCREENINGSTATUS_REJECTED WatchlistScreeningStatus = "rejected"
	WATCHLISTSCREENINGSTATUS_PENDING_REVIEW WatchlistScreeningStatus = "pending_review"
	WATCHLISTSCREENINGSTATUS_CLEARED WatchlistScreeningStatus = "cleared"
)

var allowedWatchlistScreeningStatusEnumValues = []WatchlistScreeningStatus{
	"rejected",
	"pending_review",
	"cleared",
}

func (v *WatchlistScreeningStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WatchlistScreeningStatus(value)
	for _, existing := range allowedWatchlistScreeningStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WatchlistScreeningStatus", value)
}

// NewWatchlistScreeningStatusFromValue returns a pointer to a valid WatchlistScreeningStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWatchlistScreeningStatusFromValue(v string) (*WatchlistScreeningStatus, error) {
	ev := WatchlistScreeningStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WatchlistScreeningStatus: valid values are %v", v, allowedWatchlistScreeningStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WatchlistScreeningStatus) IsValid() bool {
	for _, existing := range allowedWatchlistScreeningStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WatchlistScreeningStatus value
func (v WatchlistScreeningStatus) Ptr() *WatchlistScreeningStatus {
	return &v
}

type NullableWatchlistScreeningStatus struct {
	value *WatchlistScreeningStatus
	isSet bool
}

func (v NullableWatchlistScreeningStatus) Get() *WatchlistScreeningStatus {
	return v.value
}

func (v *NullableWatchlistScreeningStatus) Set(val *WatchlistScreeningStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningStatus(val *WatchlistScreeningStatus) *NullableWatchlistScreeningStatus {
	return &NullableWatchlistScreeningStatus{value: val, isSet: true}
}

func (v NullableWatchlistScreeningStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

