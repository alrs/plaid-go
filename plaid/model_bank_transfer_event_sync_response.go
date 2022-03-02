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

// BankTransferEventSyncResponse Defines the response schema for `/bank_transfer/event/sync`
type BankTransferEventSyncResponse struct {
	BankTransferEvents []BankTransferEvent `json:"bank_transfer_events"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _BankTransferEventSyncResponse BankTransferEventSyncResponse

// NewBankTransferEventSyncResponse instantiates a new BankTransferEventSyncResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferEventSyncResponse(bankTransferEvents []BankTransferEvent, requestId string) *BankTransferEventSyncResponse {
	this := BankTransferEventSyncResponse{}
	this.BankTransferEvents = bankTransferEvents
	this.RequestId = requestId
	return &this
}

// NewBankTransferEventSyncResponseWithDefaults instantiates a new BankTransferEventSyncResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferEventSyncResponseWithDefaults() *BankTransferEventSyncResponse {
	this := BankTransferEventSyncResponse{}
	return &this
}

// GetBankTransferEvents returns the BankTransferEvents field value
func (o *BankTransferEventSyncResponse) GetBankTransferEvents() []BankTransferEvent {
	if o == nil {
		var ret []BankTransferEvent
		return ret
	}

	return o.BankTransferEvents
}

// GetBankTransferEventsOk returns a tuple with the BankTransferEvents field value
// and a boolean to check if the value has been set.
func (o *BankTransferEventSyncResponse) GetBankTransferEventsOk() (*[]BankTransferEvent, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankTransferEvents, true
}

// SetBankTransferEvents sets field value
func (o *BankTransferEventSyncResponse) SetBankTransferEvents(v []BankTransferEvent) {
	o.BankTransferEvents = v
}

// GetRequestId returns the RequestId field value
func (o *BankTransferEventSyncResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *BankTransferEventSyncResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *BankTransferEventSyncResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o BankTransferEventSyncResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bank_transfer_events"] = o.BankTransferEvents
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BankTransferEventSyncResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBankTransferEventSyncResponse := _BankTransferEventSyncResponse{}

	if err = json.Unmarshal(bytes, &varBankTransferEventSyncResponse); err == nil {
		*o = BankTransferEventSyncResponse(varBankTransferEventSyncResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "bank_transfer_events")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBankTransferEventSyncResponse struct {
	value *BankTransferEventSyncResponse
	isSet bool
}

func (v NullableBankTransferEventSyncResponse) Get() *BankTransferEventSyncResponse {
	return v.value
}

func (v *NullableBankTransferEventSyncResponse) Set(val *BankTransferEventSyncResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferEventSyncResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferEventSyncResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferEventSyncResponse(val *BankTransferEventSyncResponse) *NullableBankTransferEventSyncResponse {
	return &NullableBankTransferEventSyncResponse{value: val, isSet: true}
}

func (v NullableBankTransferEventSyncResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferEventSyncResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


