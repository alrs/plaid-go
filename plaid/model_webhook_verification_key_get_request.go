/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.64.14
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// WebhookVerificationKeyGetRequest WebhookVerificationKeyGetRequest defines the request schema for `/webhook_verification_key/get`
type WebhookVerificationKeyGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The key ID ( `kid` ) from the JWT header.
	KeyId string `json:"key_id"`
}

// NewWebhookVerificationKeyGetRequest instantiates a new WebhookVerificationKeyGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookVerificationKeyGetRequest(keyId string) *WebhookVerificationKeyGetRequest {
	this := WebhookVerificationKeyGetRequest{}
	this.KeyId = keyId
	return &this
}

// NewWebhookVerificationKeyGetRequestWithDefaults instantiates a new WebhookVerificationKeyGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookVerificationKeyGetRequestWithDefaults() *WebhookVerificationKeyGetRequest {
	this := WebhookVerificationKeyGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *WebhookVerificationKeyGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookVerificationKeyGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *WebhookVerificationKeyGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *WebhookVerificationKeyGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *WebhookVerificationKeyGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookVerificationKeyGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *WebhookVerificationKeyGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *WebhookVerificationKeyGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetKeyId returns the KeyId field value
func (o *WebhookVerificationKeyGetRequest) GetKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value
// and a boolean to check if the value has been set.
func (o *WebhookVerificationKeyGetRequest) GetKeyIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.KeyId, true
}

// SetKeyId sets field value
func (o *WebhookVerificationKeyGetRequest) SetKeyId(v string) {
	o.KeyId = v
}

func (o WebhookVerificationKeyGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["key_id"] = o.KeyId
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookVerificationKeyGetRequest struct {
	value *WebhookVerificationKeyGetRequest
	isSet bool
}

func (v NullableWebhookVerificationKeyGetRequest) Get() *WebhookVerificationKeyGetRequest {
	return v.value
}

func (v *NullableWebhookVerificationKeyGetRequest) Set(val *WebhookVerificationKeyGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookVerificationKeyGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookVerificationKeyGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookVerificationKeyGetRequest(val *WebhookVerificationKeyGetRequest) *NullableWebhookVerificationKeyGetRequest {
	return &NullableWebhookVerificationKeyGetRequest{value: val, isSet: true}
}

func (v NullableWebhookVerificationKeyGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookVerificationKeyGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


