/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.419.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TransferRecurringGetRequest Defines the request schema for `/transfer/recurring/get`
type TransferRecurringGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Plaid’s unique identifier for a recurring transfer.
	RecurringTransferId string `json:"recurring_transfer_id"`
}

// NewTransferRecurringGetRequest instantiates a new TransferRecurringGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRecurringGetRequest(recurringTransferId string) *TransferRecurringGetRequest {
	this := TransferRecurringGetRequest{}
	this.RecurringTransferId = recurringTransferId
	return &this
}

// NewTransferRecurringGetRequestWithDefaults instantiates a new TransferRecurringGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRecurringGetRequestWithDefaults() *TransferRecurringGetRequest {
	this := TransferRecurringGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransferRecurringGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRecurringGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransferRecurringGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransferRecurringGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransferRecurringGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRecurringGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransferRecurringGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransferRecurringGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetRecurringTransferId returns the RecurringTransferId field value
func (o *TransferRecurringGetRequest) GetRecurringTransferId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecurringTransferId
}

// GetRecurringTransferIdOk returns a tuple with the RecurringTransferId field value
// and a boolean to check if the value has been set.
func (o *TransferRecurringGetRequest) GetRecurringTransferIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecurringTransferId, true
}

// SetRecurringTransferId sets field value
func (o *TransferRecurringGetRequest) SetRecurringTransferId(v string) {
	o.RecurringTransferId = v
}

func (o TransferRecurringGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["recurring_transfer_id"] = o.RecurringTransferId
	}
	return json.Marshal(toSerialize)
}

type NullableTransferRecurringGetRequest struct {
	value *TransferRecurringGetRequest
	isSet bool
}

func (v NullableTransferRecurringGetRequest) Get() *TransferRecurringGetRequest {
	return v.value
}

func (v *NullableTransferRecurringGetRequest) Set(val *TransferRecurringGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRecurringGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRecurringGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRecurringGetRequest(val *TransferRecurringGetRequest) *NullableTransferRecurringGetRequest {
	return &NullableTransferRecurringGetRequest{value: val, isSet: true}
}

func (v NullableTransferRecurringGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRecurringGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


