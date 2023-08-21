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

// SandboxItemFireWebhookRequest SandboxItemFireWebhookRequest defines the request schema for `/sandbox/item/fire_webhook`
type SandboxItemFireWebhookRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	WebhookType *WebhookType `json:"webhook_type,omitempty"`
	// The webhook codes that can be fired by this test endpoint.
	WebhookCode string `json:"webhook_code"`
}

// NewSandboxItemFireWebhookRequest instantiates a new SandboxItemFireWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxItemFireWebhookRequest(accessToken string, webhookCode string) *SandboxItemFireWebhookRequest {
	this := SandboxItemFireWebhookRequest{}
	this.AccessToken = accessToken
	this.WebhookCode = webhookCode
	return &this
}

// NewSandboxItemFireWebhookRequestWithDefaults instantiates a new SandboxItemFireWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxItemFireWebhookRequestWithDefaults() *SandboxItemFireWebhookRequest {
	this := SandboxItemFireWebhookRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SandboxItemFireWebhookRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxItemFireWebhookRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SandboxItemFireWebhookRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SandboxItemFireWebhookRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *SandboxItemFireWebhookRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxItemFireWebhookRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *SandboxItemFireWebhookRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *SandboxItemFireWebhookRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccessToken returns the AccessToken field value
func (o *SandboxItemFireWebhookRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *SandboxItemFireWebhookRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *SandboxItemFireWebhookRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetWebhookType returns the WebhookType field value if set, zero value otherwise.
func (o *SandboxItemFireWebhookRequest) GetWebhookType() WebhookType {
	if o == nil || o.WebhookType == nil {
		var ret WebhookType
		return ret
	}
	return *o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxItemFireWebhookRequest) GetWebhookTypeOk() (*WebhookType, bool) {
	if o == nil || o.WebhookType == nil {
		return nil, false
	}
	return o.WebhookType, true
}

// HasWebhookType returns a boolean if a field has been set.
func (o *SandboxItemFireWebhookRequest) HasWebhookType() bool {
	if o != nil && o.WebhookType != nil {
		return true
	}

	return false
}

// SetWebhookType gets a reference to the given WebhookType and assigns it to the WebhookType field.
func (o *SandboxItemFireWebhookRequest) SetWebhookType(v WebhookType) {
	o.WebhookType = &v
}

// GetWebhookCode returns the WebhookCode field value
func (o *SandboxItemFireWebhookRequest) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *SandboxItemFireWebhookRequest) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *SandboxItemFireWebhookRequest) SetWebhookCode(v string) {
	o.WebhookCode = v
}

func (o SandboxItemFireWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.WebhookType != nil {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxItemFireWebhookRequest struct {
	value *SandboxItemFireWebhookRequest
	isSet bool
}

func (v NullableSandboxItemFireWebhookRequest) Get() *SandboxItemFireWebhookRequest {
	return v.value
}

func (v *NullableSandboxItemFireWebhookRequest) Set(val *SandboxItemFireWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxItemFireWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxItemFireWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxItemFireWebhookRequest(val *SandboxItemFireWebhookRequest) *NullableSandboxItemFireWebhookRequest {
	return &NullableSandboxItemFireWebhookRequest{value: val, isSet: true}
}

func (v NullableSandboxItemFireWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxItemFireWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


