/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.379.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// PaymentInitiationConsentPaymentExecuteResponse PaymentInitiationConsentPaymentExecuteResponse defines the response schema for `/payment_initiation/consent/payment/execute`
type PaymentInitiationConsentPaymentExecuteResponse struct {
	// A unique ID identifying the payment
	PaymentId string `json:"payment_id"`
	Status PaymentInitiationPaymentStatus `json:"status"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _PaymentInitiationConsentPaymentExecuteResponse PaymentInitiationConsentPaymentExecuteResponse

// NewPaymentInitiationConsentPaymentExecuteResponse instantiates a new PaymentInitiationConsentPaymentExecuteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationConsentPaymentExecuteResponse(paymentId string, status PaymentInitiationPaymentStatus, requestId string) *PaymentInitiationConsentPaymentExecuteResponse {
	this := PaymentInitiationConsentPaymentExecuteResponse{}
	this.PaymentId = paymentId
	this.Status = status
	this.RequestId = requestId
	return &this
}

// NewPaymentInitiationConsentPaymentExecuteResponseWithDefaults instantiates a new PaymentInitiationConsentPaymentExecuteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationConsentPaymentExecuteResponseWithDefaults() *PaymentInitiationConsentPaymentExecuteResponse {
	this := PaymentInitiationConsentPaymentExecuteResponse{}
	return &this
}

// GetPaymentId returns the PaymentId field value
func (o *PaymentInitiationConsentPaymentExecuteResponse) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsentPaymentExecuteResponse) GetPaymentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *PaymentInitiationConsentPaymentExecuteResponse) SetPaymentId(v string) {
	o.PaymentId = v
}

// GetStatus returns the Status field value
func (o *PaymentInitiationConsentPaymentExecuteResponse) GetStatus() PaymentInitiationPaymentStatus {
	if o == nil {
		var ret PaymentInitiationPaymentStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsentPaymentExecuteResponse) GetStatusOk() (*PaymentInitiationPaymentStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentInitiationConsentPaymentExecuteResponse) SetStatus(v PaymentInitiationPaymentStatus) {
	o.Status = v
}

// GetRequestId returns the RequestId field value
func (o *PaymentInitiationConsentPaymentExecuteResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsentPaymentExecuteResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *PaymentInitiationConsentPaymentExecuteResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o PaymentInitiationConsentPaymentExecuteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["payment_id"] = o.PaymentId
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

func (o *PaymentInitiationConsentPaymentExecuteResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentInitiationConsentPaymentExecuteResponse := _PaymentInitiationConsentPaymentExecuteResponse{}

	if err = json.Unmarshal(bytes, &varPaymentInitiationConsentPaymentExecuteResponse); err == nil {
		*o = PaymentInitiationConsentPaymentExecuteResponse(varPaymentInitiationConsentPaymentExecuteResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "payment_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentInitiationConsentPaymentExecuteResponse struct {
	value *PaymentInitiationConsentPaymentExecuteResponse
	isSet bool
}

func (v NullablePaymentInitiationConsentPaymentExecuteResponse) Get() *PaymentInitiationConsentPaymentExecuteResponse {
	return v.value
}

func (v *NullablePaymentInitiationConsentPaymentExecuteResponse) Set(val *PaymentInitiationConsentPaymentExecuteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationConsentPaymentExecuteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationConsentPaymentExecuteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationConsentPaymentExecuteResponse(val *PaymentInitiationConsentPaymentExecuteResponse) *NullablePaymentInitiationConsentPaymentExecuteResponse {
	return &NullablePaymentInitiationConsentPaymentExecuteResponse{value: val, isSet: true}
}

func (v NullablePaymentInitiationConsentPaymentExecuteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationConsentPaymentExecuteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


