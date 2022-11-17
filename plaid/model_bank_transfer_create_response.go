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
)

// BankTransferCreateResponse Defines the response schema for `/bank_transfer/create`
type BankTransferCreateResponse struct {
	BankTransfer BankTransfer `json:"bank_transfer"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _BankTransferCreateResponse BankTransferCreateResponse

// NewBankTransferCreateResponse instantiates a new BankTransferCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferCreateResponse(bankTransfer BankTransfer, requestId string) *BankTransferCreateResponse {
	this := BankTransferCreateResponse{}
	this.BankTransfer = bankTransfer
	this.RequestId = requestId
	return &this
}

// NewBankTransferCreateResponseWithDefaults instantiates a new BankTransferCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferCreateResponseWithDefaults() *BankTransferCreateResponse {
	this := BankTransferCreateResponse{}
	return &this
}

// GetBankTransfer returns the BankTransfer field value
func (o *BankTransferCreateResponse) GetBankTransfer() BankTransfer {
	if o == nil {
		var ret BankTransfer
		return ret
	}

	return o.BankTransfer
}

// GetBankTransferOk returns a tuple with the BankTransfer field value
// and a boolean to check if the value has been set.
func (o *BankTransferCreateResponse) GetBankTransferOk() (*BankTransfer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankTransfer, true
}

// SetBankTransfer sets field value
func (o *BankTransferCreateResponse) SetBankTransfer(v BankTransfer) {
	o.BankTransfer = v
}

// GetRequestId returns the RequestId field value
func (o *BankTransferCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *BankTransferCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *BankTransferCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o BankTransferCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bank_transfer"] = o.BankTransfer
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BankTransferCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBankTransferCreateResponse := _BankTransferCreateResponse{}

	if err = json.Unmarshal(bytes, &varBankTransferCreateResponse); err == nil {
		*o = BankTransferCreateResponse(varBankTransferCreateResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "bank_transfer")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBankTransferCreateResponse struct {
	value *BankTransferCreateResponse
	isSet bool
}

func (v NullableBankTransferCreateResponse) Get() *BankTransferCreateResponse {
	return v.value
}

func (v *NullableBankTransferCreateResponse) Set(val *BankTransferCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferCreateResponse(val *BankTransferCreateResponse) *NullableBankTransferCreateResponse {
	return &NullableBankTransferCreateResponse{value: val, isSet: true}
}

func (v NullableBankTransferCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


