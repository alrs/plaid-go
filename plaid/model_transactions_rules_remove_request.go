/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.121.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TransactionsRulesRemoveRequest TransactionsRulesRemoveRequest defines the request schema for `/beta/transactions/rules/v1/remove`
type TransactionsRulesRemoveRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId string `json:"client_id"`
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret string `json:"secret"`
	// A rule's unique identifier
	RuleId string `json:"rule_id"`
}

// NewTransactionsRulesRemoveRequest instantiates a new TransactionsRulesRemoveRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsRulesRemoveRequest(clientId string, accessToken string, secret string, ruleId string) *TransactionsRulesRemoveRequest {
	this := TransactionsRulesRemoveRequest{}
	this.ClientId = clientId
	this.AccessToken = accessToken
	this.Secret = secret
	this.RuleId = ruleId
	return &this
}

// NewTransactionsRulesRemoveRequestWithDefaults instantiates a new TransactionsRulesRemoveRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsRulesRemoveRequestWithDefaults() *TransactionsRulesRemoveRequest {
	this := TransactionsRulesRemoveRequest{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *TransactionsRulesRemoveRequest) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *TransactionsRulesRemoveRequest) GetClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *TransactionsRulesRemoveRequest) SetClientId(v string) {
	o.ClientId = v
}

// GetAccessToken returns the AccessToken field value
func (o *TransactionsRulesRemoveRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *TransactionsRulesRemoveRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *TransactionsRulesRemoveRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetSecret returns the Secret field value
func (o *TransactionsRulesRemoveRequest) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *TransactionsRulesRemoveRequest) GetSecretOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *TransactionsRulesRemoveRequest) SetSecret(v string) {
	o.Secret = v
}

// GetRuleId returns the RuleId field value
func (o *TransactionsRulesRemoveRequest) GetRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
func (o *TransactionsRulesRemoveRequest) GetRuleIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RuleId, true
}

// SetRuleId sets field value
func (o *TransactionsRulesRemoveRequest) SetRuleId(v string) {
	o.RuleId = v
}

func (o TransactionsRulesRemoveRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if true {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["rule_id"] = o.RuleId
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionsRulesRemoveRequest struct {
	value *TransactionsRulesRemoveRequest
	isSet bool
}

func (v NullableTransactionsRulesRemoveRequest) Get() *TransactionsRulesRemoveRequest {
	return v.value
}

func (v *NullableTransactionsRulesRemoveRequest) Set(val *TransactionsRulesRemoveRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsRulesRemoveRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsRulesRemoveRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsRulesRemoveRequest(val *TransactionsRulesRemoveRequest) *NullableTransactionsRulesRemoveRequest {
	return &NullableTransactionsRulesRemoveRequest{value: val, isSet: true}
}

func (v NullableTransactionsRulesRemoveRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsRulesRemoveRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


