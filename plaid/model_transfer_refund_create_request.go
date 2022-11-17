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

// TransferRefundCreateRequest Defines the request schema for `/transfer/refund/create`
type TransferRefundCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The ID of the transfer to refund.
	TransferId string `json:"transfer_id"`
	// The amount of the refund (decimal string with two digits of precision e.g. \"10.00\").
	Amount string `json:"amount"`
	// A random key provided by the client, per unique refund. Maximum of 50 characters.  The API supports idempotency for safely retrying requests without accidentally performing the same operation twice. For example, if a request to create a refund fails due to a network connection error, you can retry the request with the same idempotency key to guarantee that only a single refund is created.
	IdempotencyKey string `json:"idempotency_key"`
}

// NewTransferRefundCreateRequest instantiates a new TransferRefundCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRefundCreateRequest(transferId string, amount string, idempotencyKey string) *TransferRefundCreateRequest {
	this := TransferRefundCreateRequest{}
	this.TransferId = transferId
	this.Amount = amount
	this.IdempotencyKey = idempotencyKey
	return &this
}

// NewTransferRefundCreateRequestWithDefaults instantiates a new TransferRefundCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRefundCreateRequestWithDefaults() *TransferRefundCreateRequest {
	this := TransferRefundCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransferRefundCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRefundCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransferRefundCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransferRefundCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransferRefundCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRefundCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransferRefundCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransferRefundCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetTransferId returns the TransferId field value
func (o *TransferRefundCreateRequest) GetTransferId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransferId
}

// GetTransferIdOk returns a tuple with the TransferId field value
// and a boolean to check if the value has been set.
func (o *TransferRefundCreateRequest) GetTransferIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransferId, true
}

// SetTransferId sets field value
func (o *TransferRefundCreateRequest) SetTransferId(v string) {
	o.TransferId = v
}

// GetAmount returns the Amount field value
func (o *TransferRefundCreateRequest) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransferRefundCreateRequest) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransferRefundCreateRequest) SetAmount(v string) {
	o.Amount = v
}

// GetIdempotencyKey returns the IdempotencyKey field value
func (o *TransferRefundCreateRequest) GetIdempotencyKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value
// and a boolean to check if the value has been set.
func (o *TransferRefundCreateRequest) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IdempotencyKey, true
}

// SetIdempotencyKey sets field value
func (o *TransferRefundCreateRequest) SetIdempotencyKey(v string) {
	o.IdempotencyKey = v
}

func (o TransferRefundCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["transfer_id"] = o.TransferId
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["idempotency_key"] = o.IdempotencyKey
	}
	return json.Marshal(toSerialize)
}

type NullableTransferRefundCreateRequest struct {
	value *TransferRefundCreateRequest
	isSet bool
}

func (v NullableTransferRefundCreateRequest) Get() *TransferRefundCreateRequest {
	return v.value
}

func (v *NullableTransferRefundCreateRequest) Set(val *TransferRefundCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRefundCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRefundCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRefundCreateRequest(val *TransferRefundCreateRequest) *NullableTransferRefundCreateRequest {
	return &NullableTransferRefundCreateRequest{value: val, isSet: true}
}

func (v NullableTransferRefundCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRefundCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


