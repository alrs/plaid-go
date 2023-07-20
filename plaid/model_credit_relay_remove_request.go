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

// CreditRelayRemoveRequest CreditRelayRemoveRequest defines the request schema for `/credit/relay/remove`
type CreditRelayRemoveRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The `relay_token` you would like to revoke.
	RelayToken string `json:"relay_token"`
}

// NewCreditRelayRemoveRequest instantiates a new CreditRelayRemoveRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditRelayRemoveRequest(relayToken string) *CreditRelayRemoveRequest {
	this := CreditRelayRemoveRequest{}
	this.RelayToken = relayToken
	return &this
}

// NewCreditRelayRemoveRequestWithDefaults instantiates a new CreditRelayRemoveRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditRelayRemoveRequestWithDefaults() *CreditRelayRemoveRequest {
	this := CreditRelayRemoveRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CreditRelayRemoveRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditRelayRemoveRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CreditRelayRemoveRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CreditRelayRemoveRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *CreditRelayRemoveRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditRelayRemoveRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *CreditRelayRemoveRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *CreditRelayRemoveRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetRelayToken returns the RelayToken field value
func (o *CreditRelayRemoveRequest) GetRelayToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelayToken
}

// GetRelayTokenOk returns a tuple with the RelayToken field value
// and a boolean to check if the value has been set.
func (o *CreditRelayRemoveRequest) GetRelayTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RelayToken, true
}

// SetRelayToken sets field value
func (o *CreditRelayRemoveRequest) SetRelayToken(v string) {
	o.RelayToken = v
}

func (o CreditRelayRemoveRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["relay_token"] = o.RelayToken
	}
	return json.Marshal(toSerialize)
}

type NullableCreditRelayRemoveRequest struct {
	value *CreditRelayRemoveRequest
	isSet bool
}

func (v NullableCreditRelayRemoveRequest) Get() *CreditRelayRemoveRequest {
	return v.value
}

func (v *NullableCreditRelayRemoveRequest) Set(val *CreditRelayRemoveRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditRelayRemoveRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditRelayRemoveRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditRelayRemoveRequest(val *CreditRelayRemoveRequest) *NullableCreditRelayRemoveRequest {
	return &NullableCreditRelayRemoveRequest{value: val, isSet: true}
}

func (v NullableCreditRelayRemoveRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditRelayRemoveRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


