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

// PaymentProfileCreateResponse PaymentProfileCreateResponse defines the response schema for `/payment_profile/create`
type PaymentProfileCreateResponse struct {
	// A payment profile token associated with the Payment Profile data that is being requested.
	PaymentProfileToken string `json:"payment_profile_token"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _PaymentProfileCreateResponse PaymentProfileCreateResponse

// NewPaymentProfileCreateResponse instantiates a new PaymentProfileCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentProfileCreateResponse(paymentProfileToken string, requestId string) *PaymentProfileCreateResponse {
	this := PaymentProfileCreateResponse{}
	this.PaymentProfileToken = paymentProfileToken
	this.RequestId = requestId
	return &this
}

// NewPaymentProfileCreateResponseWithDefaults instantiates a new PaymentProfileCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentProfileCreateResponseWithDefaults() *PaymentProfileCreateResponse {
	this := PaymentProfileCreateResponse{}
	return &this
}

// GetPaymentProfileToken returns the PaymentProfileToken field value
func (o *PaymentProfileCreateResponse) GetPaymentProfileToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentProfileToken
}

// GetPaymentProfileTokenOk returns a tuple with the PaymentProfileToken field value
// and a boolean to check if the value has been set.
func (o *PaymentProfileCreateResponse) GetPaymentProfileTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentProfileToken, true
}

// SetPaymentProfileToken sets field value
func (o *PaymentProfileCreateResponse) SetPaymentProfileToken(v string) {
	o.PaymentProfileToken = v
}

// GetRequestId returns the RequestId field value
func (o *PaymentProfileCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *PaymentProfileCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *PaymentProfileCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o PaymentProfileCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["payment_profile_token"] = o.PaymentProfileToken
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentProfileCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentProfileCreateResponse := _PaymentProfileCreateResponse{}

	if err = json.Unmarshal(bytes, &varPaymentProfileCreateResponse); err == nil {
		*o = PaymentProfileCreateResponse(varPaymentProfileCreateResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "payment_profile_token")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentProfileCreateResponse struct {
	value *PaymentProfileCreateResponse
	isSet bool
}

func (v NullablePaymentProfileCreateResponse) Get() *PaymentProfileCreateResponse {
	return v.value
}

func (v *NullablePaymentProfileCreateResponse) Set(val *PaymentProfileCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentProfileCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentProfileCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentProfileCreateResponse(val *PaymentProfileCreateResponse) *NullablePaymentProfileCreateResponse {
	return &NullablePaymentProfileCreateResponse{value: val, isSet: true}
}

func (v NullablePaymentProfileCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentProfileCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


