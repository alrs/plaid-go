/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.64.14
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// TransferNetwork The network or rails used for the transfer. Valid options are `ach` or `same-day-ach`.
type TransferNetwork string

// List of TransferNetwork
const (
	TRANSFERNETWORK_ACH TransferNetwork = "ach"
	TRANSFERNETWORK_SAME_DAY_ACH TransferNetwork = "same-day-ach"
)

var allowedTransferNetworkEnumValues = []TransferNetwork{
	"ach",
	"same-day-ach",
}

func (v *TransferNetwork) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransferNetwork(value)
	for _, existing := range allowedTransferNetworkEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransferNetwork", value)
}

// NewTransferNetworkFromValue returns a pointer to a valid TransferNetwork
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferNetworkFromValue(v string) (*TransferNetwork, error) {
	ev := TransferNetwork(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransferNetwork: valid values are %v", v, allowedTransferNetworkEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferNetwork) IsValid() bool {
	for _, existing := range allowedTransferNetworkEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferNetwork value
func (v TransferNetwork) Ptr() *TransferNetwork {
	return &v
}

type NullableTransferNetwork struct {
	value *TransferNetwork
	isSet bool
}

func (v NullableTransferNetwork) Get() *TransferNetwork {
	return v.value
}

func (v *NullableTransferNetwork) Set(val *TransferNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferNetwork(val *TransferNetwork) *NullableTransferNetwork {
	return &NullableTransferNetwork{value: val, isSet: true}
}

func (v NullableTransferNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

