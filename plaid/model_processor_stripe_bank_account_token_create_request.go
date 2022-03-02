/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.79.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// ProcessorStripeBankAccountTokenCreateRequest ProcessorStripeBankAccountTokenCreateRequest defines the request schema for `/processor/stripe/bank_account/create`
type ProcessorStripeBankAccountTokenCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	// The `account_id` value obtained from the `onSuccess` callback in Link
	AccountId string `json:"account_id"`
}

// NewProcessorStripeBankAccountTokenCreateRequest instantiates a new ProcessorStripeBankAccountTokenCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorStripeBankAccountTokenCreateRequest(accessToken string, accountId string) *ProcessorStripeBankAccountTokenCreateRequest {
	this := ProcessorStripeBankAccountTokenCreateRequest{}
	this.AccessToken = accessToken
	this.AccountId = accountId
	return &this
}

// NewProcessorStripeBankAccountTokenCreateRequestWithDefaults instantiates a new ProcessorStripeBankAccountTokenCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorStripeBankAccountTokenCreateRequestWithDefaults() *ProcessorStripeBankAccountTokenCreateRequest {
	this := ProcessorStripeBankAccountTokenCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ProcessorStripeBankAccountTokenCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStripeBankAccountTokenCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ProcessorStripeBankAccountTokenCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ProcessorStripeBankAccountTokenCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ProcessorStripeBankAccountTokenCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStripeBankAccountTokenCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ProcessorStripeBankAccountTokenCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ProcessorStripeBankAccountTokenCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccessToken returns the AccessToken field value
func (o *ProcessorStripeBankAccountTokenCreateRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *ProcessorStripeBankAccountTokenCreateRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *ProcessorStripeBankAccountTokenCreateRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetAccountId returns the AccountId field value
func (o *ProcessorStripeBankAccountTokenCreateRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *ProcessorStripeBankAccountTokenCreateRequest) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *ProcessorStripeBankAccountTokenCreateRequest) SetAccountId(v string) {
	o.AccountId = v
}

func (o ProcessorStripeBankAccountTokenCreateRequest) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	return json.Marshal(toSerialize)
}

type NullableProcessorStripeBankAccountTokenCreateRequest struct {
	value *ProcessorStripeBankAccountTokenCreateRequest
	isSet bool
}

func (v NullableProcessorStripeBankAccountTokenCreateRequest) Get() *ProcessorStripeBankAccountTokenCreateRequest {
	return v.value
}

func (v *NullableProcessorStripeBankAccountTokenCreateRequest) Set(val *ProcessorStripeBankAccountTokenCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorStripeBankAccountTokenCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorStripeBankAccountTokenCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorStripeBankAccountTokenCreateRequest(val *ProcessorStripeBankAccountTokenCreateRequest) *NullableProcessorStripeBankAccountTokenCreateRequest {
	return &NullableProcessorStripeBankAccountTokenCreateRequest{value: val, isSet: true}
}

func (v NullableProcessorStripeBankAccountTokenCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorStripeBankAccountTokenCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


