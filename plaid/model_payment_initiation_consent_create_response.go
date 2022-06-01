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
)

// PaymentInitiationConsentCreateResponse PaymentInitiationConsentCreateResponse defines the response schema for `/payment_initiation/consent/create`
type PaymentInitiationConsentCreateResponse struct {
	// A unique ID identifying the payment consent.
	ConsentId string `json:"consent_id"`
	Status PaymentInitiationConsentStatus `json:"status"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _PaymentInitiationConsentCreateResponse PaymentInitiationConsentCreateResponse

// NewPaymentInitiationConsentCreateResponse instantiates a new PaymentInitiationConsentCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationConsentCreateResponse(consentId string, status PaymentInitiationConsentStatus, requestId string) *PaymentInitiationConsentCreateResponse {
	this := PaymentInitiationConsentCreateResponse{}
	this.ConsentId = consentId
	this.Status = status
	this.RequestId = requestId
	return &this
}

// NewPaymentInitiationConsentCreateResponseWithDefaults instantiates a new PaymentInitiationConsentCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationConsentCreateResponseWithDefaults() *PaymentInitiationConsentCreateResponse {
	this := PaymentInitiationConsentCreateResponse{}
	return &this
}

// GetConsentId returns the ConsentId field value
func (o *PaymentInitiationConsentCreateResponse) GetConsentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsentId
}

// GetConsentIdOk returns a tuple with the ConsentId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsentCreateResponse) GetConsentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConsentId, true
}

// SetConsentId sets field value
func (o *PaymentInitiationConsentCreateResponse) SetConsentId(v string) {
	o.ConsentId = v
}

// GetStatus returns the Status field value
func (o *PaymentInitiationConsentCreateResponse) GetStatus() PaymentInitiationConsentStatus {
	if o == nil {
		var ret PaymentInitiationConsentStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsentCreateResponse) GetStatusOk() (*PaymentInitiationConsentStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentInitiationConsentCreateResponse) SetStatus(v PaymentInitiationConsentStatus) {
	o.Status = v
}

// GetRequestId returns the RequestId field value
func (o *PaymentInitiationConsentCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsentCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *PaymentInitiationConsentCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o PaymentInitiationConsentCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["consent_id"] = o.ConsentId
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

func (o *PaymentInitiationConsentCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentInitiationConsentCreateResponse := _PaymentInitiationConsentCreateResponse{}

	if err = json.Unmarshal(bytes, &varPaymentInitiationConsentCreateResponse); err == nil {
		*o = PaymentInitiationConsentCreateResponse(varPaymentInitiationConsentCreateResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "consent_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentInitiationConsentCreateResponse struct {
	value *PaymentInitiationConsentCreateResponse
	isSet bool
}

func (v NullablePaymentInitiationConsentCreateResponse) Get() *PaymentInitiationConsentCreateResponse {
	return v.value
}

func (v *NullablePaymentInitiationConsentCreateResponse) Set(val *PaymentInitiationConsentCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationConsentCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationConsentCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationConsentCreateResponse(val *PaymentInitiationConsentCreateResponse) *NullablePaymentInitiationConsentCreateResponse {
	return &NullablePaymentInitiationConsentCreateResponse{value: val, isSet: true}
}

func (v NullablePaymentInitiationConsentCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationConsentCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


