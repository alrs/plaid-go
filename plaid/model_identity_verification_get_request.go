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

// IdentityVerificationGetRequest Request input for fetching an identity verification
type IdentityVerificationGetRequest struct {
	// ID of the associated Identity Verification attempt.
	IdentityVerificationId string `json:"identity_verification_id"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
}

// NewIdentityVerificationGetRequest instantiates a new IdentityVerificationGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityVerificationGetRequest(identityVerificationId string) *IdentityVerificationGetRequest {
	this := IdentityVerificationGetRequest{}
	this.IdentityVerificationId = identityVerificationId
	return &this
}

// NewIdentityVerificationGetRequestWithDefaults instantiates a new IdentityVerificationGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityVerificationGetRequestWithDefaults() *IdentityVerificationGetRequest {
	this := IdentityVerificationGetRequest{}
	return &this
}

// GetIdentityVerificationId returns the IdentityVerificationId field value
func (o *IdentityVerificationGetRequest) GetIdentityVerificationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityVerificationId
}

// GetIdentityVerificationIdOk returns a tuple with the IdentityVerificationId field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationGetRequest) GetIdentityVerificationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IdentityVerificationId, true
}

// SetIdentityVerificationId sets field value
func (o *IdentityVerificationGetRequest) SetIdentityVerificationId(v string) {
	o.IdentityVerificationId = v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *IdentityVerificationGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityVerificationGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *IdentityVerificationGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *IdentityVerificationGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *IdentityVerificationGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityVerificationGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *IdentityVerificationGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *IdentityVerificationGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

func (o IdentityVerificationGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["identity_verification_id"] = o.IdentityVerificationId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityVerificationGetRequest struct {
	value *IdentityVerificationGetRequest
	isSet bool
}

func (v NullableIdentityVerificationGetRequest) Get() *IdentityVerificationGetRequest {
	return v.value
}

func (v *NullableIdentityVerificationGetRequest) Set(val *IdentityVerificationGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityVerificationGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityVerificationGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityVerificationGetRequest(val *IdentityVerificationGetRequest) *NullableIdentityVerificationGetRequest {
	return &NullableIdentityVerificationGetRequest{value: val, isSet: true}
}

func (v NullableIdentityVerificationGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityVerificationGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


