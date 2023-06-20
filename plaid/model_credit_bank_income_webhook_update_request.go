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

// CreditBankIncomeWebhookUpdateRequest CreditBankIncomeWebhookUpdateRequest defines the request schema for `/credit/bank_income/webhook/update`.
type CreditBankIncomeWebhookUpdateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The user token associated with the User data is being requested for.
	UserToken string `json:"user_token"`
	// Whether the user should be enabled for proactive webhook notifications when their income changes
	EnableWebhooks bool `json:"enable_webhooks"`
}

// NewCreditBankIncomeWebhookUpdateRequest instantiates a new CreditBankIncomeWebhookUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankIncomeWebhookUpdateRequest(userToken string, enableWebhooks bool) *CreditBankIncomeWebhookUpdateRequest {
	this := CreditBankIncomeWebhookUpdateRequest{}
	this.UserToken = userToken
	this.EnableWebhooks = enableWebhooks
	return &this
}

// NewCreditBankIncomeWebhookUpdateRequestWithDefaults instantiates a new CreditBankIncomeWebhookUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankIncomeWebhookUpdateRequestWithDefaults() *CreditBankIncomeWebhookUpdateRequest {
	this := CreditBankIncomeWebhookUpdateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CreditBankIncomeWebhookUpdateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeWebhookUpdateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CreditBankIncomeWebhookUpdateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CreditBankIncomeWebhookUpdateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *CreditBankIncomeWebhookUpdateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeWebhookUpdateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *CreditBankIncomeWebhookUpdateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *CreditBankIncomeWebhookUpdateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetUserToken returns the UserToken field value
func (o *CreditBankIncomeWebhookUpdateRequest) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeWebhookUpdateRequest) GetUserTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value
func (o *CreditBankIncomeWebhookUpdateRequest) SetUserToken(v string) {
	o.UserToken = v
}

// GetEnableWebhooks returns the EnableWebhooks field value
func (o *CreditBankIncomeWebhookUpdateRequest) GetEnableWebhooks() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnableWebhooks
}

// GetEnableWebhooksOk returns a tuple with the EnableWebhooks field value
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeWebhookUpdateRequest) GetEnableWebhooksOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EnableWebhooks, true
}

// SetEnableWebhooks sets field value
func (o *CreditBankIncomeWebhookUpdateRequest) SetEnableWebhooks(v bool) {
	o.EnableWebhooks = v
}

func (o CreditBankIncomeWebhookUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["user_token"] = o.UserToken
	}
	if true {
		toSerialize["enable_webhooks"] = o.EnableWebhooks
	}
	return json.Marshal(toSerialize)
}

type NullableCreditBankIncomeWebhookUpdateRequest struct {
	value *CreditBankIncomeWebhookUpdateRequest
	isSet bool
}

func (v NullableCreditBankIncomeWebhookUpdateRequest) Get() *CreditBankIncomeWebhookUpdateRequest {
	return v.value
}

func (v *NullableCreditBankIncomeWebhookUpdateRequest) Set(val *CreditBankIncomeWebhookUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankIncomeWebhookUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankIncomeWebhookUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankIncomeWebhookUpdateRequest(val *CreditBankIncomeWebhookUpdateRequest) *NullableCreditBankIncomeWebhookUpdateRequest {
	return &NullableCreditBankIncomeWebhookUpdateRequest{value: val, isSet: true}
}

func (v NullableCreditBankIncomeWebhookUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankIncomeWebhookUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

