/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.419.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// TransferIntentCreateNetwork The network or rails used for the transfer. Defaults to `same-day-ach`.  For transfers submitted as either `ach` or `same-day-ach` the cutoff for same-day is 3:30 PM Eastern Time and the cutoff for next-day transfers is 5:30 PM Eastern Time. It is recommended to submit a transfer at least 15 minutes before the cutoff time in order to ensure that it will be processed before the cutoff. Any transfer that is indicated as `same-day-ach` and that misses the same-day cutoff, but is submitted in time for the next-day cutoff, will be sent over next-day rails and will not incur same-day charges. Note that both legs of the transfer will be downgraded if applicable.
type TransferIntentCreateNetwork string

var _ = fmt.Printf

// List of TransferIntentCreateNetwork
const (
	TRANSFERINTENTCREATENETWORK_ACH TransferIntentCreateNetwork = "ach"
	TRANSFERINTENTCREATENETWORK_SAME_DAY_ACH TransferIntentCreateNetwork = "same-day-ach"
)

var allowedTransferIntentCreateNetworkEnumValues = []TransferIntentCreateNetwork{
	"ach",
	"same-day-ach",
}

func (v *TransferIntentCreateNetwork) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := TransferIntentCreateNetwork(value)


	*v = enumTypeValue
	return nil
}

// NewTransferIntentCreateNetworkFromValue returns a pointer to a valid TransferIntentCreateNetwork
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferIntentCreateNetworkFromValue(v string) (*TransferIntentCreateNetwork, error) {
	ev := TransferIntentCreateNetwork(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferIntentCreateNetwork) IsValid() bool {
	for _, existing := range allowedTransferIntentCreateNetworkEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferIntentCreateNetwork value
func (v TransferIntentCreateNetwork) Ptr() *TransferIntentCreateNetwork {
	return &v
}

type NullableTransferIntentCreateNetwork struct {
	value *TransferIntentCreateNetwork
	isSet bool
}

func (v NullableTransferIntentCreateNetwork) Get() *TransferIntentCreateNetwork {
	return v.value
}

func (v *NullableTransferIntentCreateNetwork) Set(val *TransferIntentCreateNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferIntentCreateNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferIntentCreateNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferIntentCreateNetwork(val *TransferIntentCreateNetwork) *NullableTransferIntentCreateNetwork {
	return &NullableTransferIntentCreateNetwork{value: val, isSet: true}
}

func (v NullableTransferIntentCreateNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferIntentCreateNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

