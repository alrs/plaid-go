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
)

// PaymentInitiationPaymentReverseResponse PaymentInitiationPaymentReverseResponse defines the response schema for `/payment_initation/payment/reverse`
type PaymentInitiationPaymentReverseResponse struct {
	// A unique ID identifying the refund
	RefundId string `json:"refund_id"`
	Status WalletTransactionStatus `json:"status"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _PaymentInitiationPaymentReverseResponse PaymentInitiationPaymentReverseResponse

// NewPaymentInitiationPaymentReverseResponse instantiates a new PaymentInitiationPaymentReverseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationPaymentReverseResponse(refundId string, status WalletTransactionStatus, requestId string) *PaymentInitiationPaymentReverseResponse {
	this := PaymentInitiationPaymentReverseResponse{}
	this.RefundId = refundId
	this.Status = status
	this.RequestId = requestId
	return &this
}

// NewPaymentInitiationPaymentReverseResponseWithDefaults instantiates a new PaymentInitiationPaymentReverseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationPaymentReverseResponseWithDefaults() *PaymentInitiationPaymentReverseResponse {
	this := PaymentInitiationPaymentReverseResponse{}
	return &this
}

// GetRefundId returns the RefundId field value
func (o *PaymentInitiationPaymentReverseResponse) GetRefundId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefundId
}

// GetRefundIdOk returns a tuple with the RefundId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentReverseResponse) GetRefundIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RefundId, true
}

// SetRefundId sets field value
func (o *PaymentInitiationPaymentReverseResponse) SetRefundId(v string) {
	o.RefundId = v
}

// GetStatus returns the Status field value
func (o *PaymentInitiationPaymentReverseResponse) GetStatus() WalletTransactionStatus {
	if o == nil {
		var ret WalletTransactionStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentReverseResponse) GetStatusOk() (*WalletTransactionStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentInitiationPaymentReverseResponse) SetStatus(v WalletTransactionStatus) {
	o.Status = v
}

// GetRequestId returns the RequestId field value
func (o *PaymentInitiationPaymentReverseResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentReverseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *PaymentInitiationPaymentReverseResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o PaymentInitiationPaymentReverseResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["refund_id"] = o.RefundId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentInitiationPaymentReverseResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentInitiationPaymentReverseResponse := _PaymentInitiationPaymentReverseResponse{}

	if err = json.Unmarshal(bytes, &varPaymentInitiationPaymentReverseResponse); err == nil {
		*o = PaymentInitiationPaymentReverseResponse(varPaymentInitiationPaymentReverseResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "refund_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentInitiationPaymentReverseResponse struct {
	value *PaymentInitiationPaymentReverseResponse
	isSet bool
}

func (v NullablePaymentInitiationPaymentReverseResponse) Get() *PaymentInitiationPaymentReverseResponse {
	return v.value
}

func (v *NullablePaymentInitiationPaymentReverseResponse) Set(val *PaymentInitiationPaymentReverseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationPaymentReverseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationPaymentReverseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationPaymentReverseResponse(val *PaymentInitiationPaymentReverseResponse) *NullablePaymentInitiationPaymentReverseResponse {
	return &NullablePaymentInitiationPaymentReverseResponse{value: val, isSet: true}
}

func (v NullablePaymentInitiationPaymentReverseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationPaymentReverseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


