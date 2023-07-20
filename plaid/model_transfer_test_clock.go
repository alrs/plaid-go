/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.391.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"time"
)

// TransferTestClock Defines the test clock for a transfer.
type TransferTestClock struct {
	// Plaid’s unique identifier for a test clock.
	TestClockId string `json:"test_clock_id"`
	// The virtual timestamp on the test clock. This will be of the form `2006-01-02T15:04:05Z`.
	VirtualTime time.Time `json:"virtual_time"`
	AdditionalProperties map[string]interface{}
}

type _TransferTestClock TransferTestClock

// NewTransferTestClock instantiates a new TransferTestClock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferTestClock(testClockId string, virtualTime time.Time) *TransferTestClock {
	this := TransferTestClock{}
	this.TestClockId = testClockId
	this.VirtualTime = virtualTime
	return &this
}

// NewTransferTestClockWithDefaults instantiates a new TransferTestClock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferTestClockWithDefaults() *TransferTestClock {
	this := TransferTestClock{}
	return &this
}

// GetTestClockId returns the TestClockId field value
func (o *TransferTestClock) GetTestClockId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestClockId
}

// GetTestClockIdOk returns a tuple with the TestClockId field value
// and a boolean to check if the value has been set.
func (o *TransferTestClock) GetTestClockIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TestClockId, true
}

// SetTestClockId sets field value
func (o *TransferTestClock) SetTestClockId(v string) {
	o.TestClockId = v
}

// GetVirtualTime returns the VirtualTime field value
func (o *TransferTestClock) GetVirtualTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.VirtualTime
}

// GetVirtualTimeOk returns a tuple with the VirtualTime field value
// and a boolean to check if the value has been set.
func (o *TransferTestClock) GetVirtualTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VirtualTime, true
}

// SetVirtualTime sets field value
func (o *TransferTestClock) SetVirtualTime(v time.Time) {
	o.VirtualTime = v
}

func (o TransferTestClock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["test_clock_id"] = o.TestClockId
	}
	if true {
		toSerialize["virtual_time"] = o.VirtualTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferTestClock) UnmarshalJSON(bytes []byte) (err error) {
	varTransferTestClock := _TransferTestClock{}

	if err = json.Unmarshal(bytes, &varTransferTestClock); err == nil {
		*o = TransferTestClock(varTransferTestClock)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "test_clock_id")
		delete(additionalProperties, "virtual_time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferTestClock struct {
	value *TransferTestClock
	isSet bool
}

func (v NullableTransferTestClock) Get() *TransferTestClock {
	return v.value
}

func (v *NullableTransferTestClock) Set(val *TransferTestClock) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferTestClock) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferTestClock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferTestClock(val *TransferTestClock) *NullableTransferTestClock {
	return &NullableTransferTestClock{value: val, isSet: true}
}

func (v NullableTransferTestClock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferTestClock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


