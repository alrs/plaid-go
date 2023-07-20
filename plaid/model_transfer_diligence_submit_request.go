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

// TransferDiligenceSubmitRequest Defines the request schema for `/transfer/diligence/submit`
type TransferDiligenceSubmitRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Client ID of the the originator whose diligence that you want to submit.
	OriginatorClientId string `json:"originator_client_id"`
	OriginatorDiligence TransferOriginatorDiligence `json:"originator_diligence"`
}

// NewTransferDiligenceSubmitRequest instantiates a new TransferDiligenceSubmitRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferDiligenceSubmitRequest(originatorClientId string, originatorDiligence TransferOriginatorDiligence) *TransferDiligenceSubmitRequest {
	this := TransferDiligenceSubmitRequest{}
	this.OriginatorClientId = originatorClientId
	this.OriginatorDiligence = originatorDiligence
	return &this
}

// NewTransferDiligenceSubmitRequestWithDefaults instantiates a new TransferDiligenceSubmitRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferDiligenceSubmitRequestWithDefaults() *TransferDiligenceSubmitRequest {
	this := TransferDiligenceSubmitRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransferDiligenceSubmitRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferDiligenceSubmitRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransferDiligenceSubmitRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransferDiligenceSubmitRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransferDiligenceSubmitRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferDiligenceSubmitRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransferDiligenceSubmitRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransferDiligenceSubmitRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetOriginatorClientId returns the OriginatorClientId field value
func (o *TransferDiligenceSubmitRequest) GetOriginatorClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginatorClientId
}

// GetOriginatorClientIdOk returns a tuple with the OriginatorClientId field value
// and a boolean to check if the value has been set.
func (o *TransferDiligenceSubmitRequest) GetOriginatorClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OriginatorClientId, true
}

// SetOriginatorClientId sets field value
func (o *TransferDiligenceSubmitRequest) SetOriginatorClientId(v string) {
	o.OriginatorClientId = v
}

// GetOriginatorDiligence returns the OriginatorDiligence field value
func (o *TransferDiligenceSubmitRequest) GetOriginatorDiligence() TransferOriginatorDiligence {
	if o == nil {
		var ret TransferOriginatorDiligence
		return ret
	}

	return o.OriginatorDiligence
}

// GetOriginatorDiligenceOk returns a tuple with the OriginatorDiligence field value
// and a boolean to check if the value has been set.
func (o *TransferDiligenceSubmitRequest) GetOriginatorDiligenceOk() (*TransferOriginatorDiligence, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OriginatorDiligence, true
}

// SetOriginatorDiligence sets field value
func (o *TransferDiligenceSubmitRequest) SetOriginatorDiligence(v TransferOriginatorDiligence) {
	o.OriginatorDiligence = v
}

func (o TransferDiligenceSubmitRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["originator_client_id"] = o.OriginatorClientId
	}
	if true {
		toSerialize["originator_diligence"] = o.OriginatorDiligence
	}
	return json.Marshal(toSerialize)
}

type NullableTransferDiligenceSubmitRequest struct {
	value *TransferDiligenceSubmitRequest
	isSet bool
}

func (v NullableTransferDiligenceSubmitRequest) Get() *TransferDiligenceSubmitRequest {
	return v.value
}

func (v *NullableTransferDiligenceSubmitRequest) Set(val *TransferDiligenceSubmitRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferDiligenceSubmitRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferDiligenceSubmitRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferDiligenceSubmitRequest(val *TransferDiligenceSubmitRequest) *NullableTransferDiligenceSubmitRequest {
	return &NullableTransferDiligenceSubmitRequest{value: val, isSet: true}
}

func (v NullableTransferDiligenceSubmitRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferDiligenceSubmitRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


