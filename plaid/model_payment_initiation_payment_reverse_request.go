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

// PaymentInitiationPaymentReverseRequest PaymentInitiationPaymentReverseRequest defines the request schema for `/payment_initiation/payment/reverse`
type PaymentInitiationPaymentReverseRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The ID of the payment to reverse
	PaymentId string `json:"payment_id"`
	// A random key provided by the client, per unique wallet transaction. Maximum of 128 characters.  The API supports idempotency for safely retrying requests without accidentally performing the same operation twice. If a request to execute a wallet transaction fails due to a network connection error, then after a minimum delay of one minute, you can retry the request with the same idempotency key to guarantee that only a single wallet transaction is created. If the request was successfully processed, it will prevent any transaction that uses the same idempotency key, and was received within 24 hours of the first request, from being processed.
	IdempotencyKey string `json:"idempotency_key"`
	// A reference for the refund. This must be an alphanumeric string with 6 to 18 characters and must not contain any special characters or spaces.
	Reference string `json:"reference"`
	Amount *PaymentAmountToRefund `json:"amount,omitempty"`
}

// NewPaymentInitiationPaymentReverseRequest instantiates a new PaymentInitiationPaymentReverseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationPaymentReverseRequest(paymentId string, idempotencyKey string, reference string) *PaymentInitiationPaymentReverseRequest {
	this := PaymentInitiationPaymentReverseRequest{}
	this.PaymentId = paymentId
	this.IdempotencyKey = idempotencyKey
	this.Reference = reference
	return &this
}

// NewPaymentInitiationPaymentReverseRequestWithDefaults instantiates a new PaymentInitiationPaymentReverseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationPaymentReverseRequestWithDefaults() *PaymentInitiationPaymentReverseRequest {
	this := PaymentInitiationPaymentReverseRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PaymentInitiationPaymentReverseRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentReverseRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentReverseRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PaymentInitiationPaymentReverseRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *PaymentInitiationPaymentReverseRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentReverseRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentReverseRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *PaymentInitiationPaymentReverseRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetPaymentId returns the PaymentId field value
func (o *PaymentInitiationPaymentReverseRequest) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentReverseRequest) GetPaymentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *PaymentInitiationPaymentReverseRequest) SetPaymentId(v string) {
	o.PaymentId = v
}

// GetIdempotencyKey returns the IdempotencyKey field value
func (o *PaymentInitiationPaymentReverseRequest) GetIdempotencyKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentReverseRequest) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IdempotencyKey, true
}

// SetIdempotencyKey sets field value
func (o *PaymentInitiationPaymentReverseRequest) SetIdempotencyKey(v string) {
	o.IdempotencyKey = v
}

// GetReference returns the Reference field value
func (o *PaymentInitiationPaymentReverseRequest) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentReverseRequest) GetReferenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *PaymentInitiationPaymentReverseRequest) SetReference(v string) {
	o.Reference = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PaymentInitiationPaymentReverseRequest) GetAmount() PaymentAmountToRefund {
	if o == nil || o.Amount == nil {
		var ret PaymentAmountToRefund
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentReverseRequest) GetAmountOk() (*PaymentAmountToRefund, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentReverseRequest) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given PaymentAmountToRefund and assigns it to the Amount field.
func (o *PaymentInitiationPaymentReverseRequest) SetAmount(v PaymentAmountToRefund) {
	o.Amount = &v
}

func (o PaymentInitiationPaymentReverseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["payment_id"] = o.PaymentId
	}
	if true {
		toSerialize["idempotency_key"] = o.IdempotencyKey
	}
	if true {
		toSerialize["reference"] = o.Reference
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentInitiationPaymentReverseRequest struct {
	value *PaymentInitiationPaymentReverseRequest
	isSet bool
}

func (v NullablePaymentInitiationPaymentReverseRequest) Get() *PaymentInitiationPaymentReverseRequest {
	return v.value
}

func (v *NullablePaymentInitiationPaymentReverseRequest) Set(val *PaymentInitiationPaymentReverseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationPaymentReverseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationPaymentReverseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationPaymentReverseRequest(val *PaymentInitiationPaymentReverseRequest) *NullablePaymentInitiationPaymentReverseRequest {
	return &NullablePaymentInitiationPaymentReverseRequest{value: val, isSet: true}
}

func (v NullablePaymentInitiationPaymentReverseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationPaymentReverseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


