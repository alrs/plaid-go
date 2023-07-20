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
)

// BankTransferListResponse Defines the response schema for `/bank_transfer/list`
type BankTransferListResponse struct {
	BankTransfers []BankTransfer `json:"bank_transfers"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _BankTransferListResponse BankTransferListResponse

// NewBankTransferListResponse instantiates a new BankTransferListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferListResponse(bankTransfers []BankTransfer, requestId string) *BankTransferListResponse {
	this := BankTransferListResponse{}
	this.BankTransfers = bankTransfers
	this.RequestId = requestId
	return &this
}

// NewBankTransferListResponseWithDefaults instantiates a new BankTransferListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferListResponseWithDefaults() *BankTransferListResponse {
	this := BankTransferListResponse{}
	return &this
}

// GetBankTransfers returns the BankTransfers field value
func (o *BankTransferListResponse) GetBankTransfers() []BankTransfer {
	if o == nil {
		var ret []BankTransfer
		return ret
	}

	return o.BankTransfers
}

// GetBankTransfersOk returns a tuple with the BankTransfers field value
// and a boolean to check if the value has been set.
func (o *BankTransferListResponse) GetBankTransfersOk() (*[]BankTransfer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankTransfers, true
}

// SetBankTransfers sets field value
func (o *BankTransferListResponse) SetBankTransfers(v []BankTransfer) {
	o.BankTransfers = v
}

// GetRequestId returns the RequestId field value
func (o *BankTransferListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *BankTransferListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *BankTransferListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o BankTransferListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bank_transfers"] = o.BankTransfers
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BankTransferListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBankTransferListResponse := _BankTransferListResponse{}

	if err = json.Unmarshal(bytes, &varBankTransferListResponse); err == nil {
		*o = BankTransferListResponse(varBankTransferListResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "bank_transfers")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBankTransferListResponse struct {
	value *BankTransferListResponse
	isSet bool
}

func (v NullableBankTransferListResponse) Get() *BankTransferListResponse {
	return v.value
}

func (v *NullableBankTransferListResponse) Set(val *BankTransferListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferListResponse(val *BankTransferListResponse) *NullableBankTransferListResponse {
	return &NullableBankTransferListResponse{value: val, isSet: true}
}

func (v NullableBankTransferListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


