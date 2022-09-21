/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.181.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// SandboxOauthSelectAccountsRequest Defines the request schema for `sandbox/oauth/select_accounts`
type SandboxOauthSelectAccountsRequest struct {
	OauthStateId string `json:"oauth_state_id"`
	Accounts []string `json:"accounts"`
}

// NewSandboxOauthSelectAccountsRequest instantiates a new SandboxOauthSelectAccountsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxOauthSelectAccountsRequest(oauthStateId string, accounts []string) *SandboxOauthSelectAccountsRequest {
	this := SandboxOauthSelectAccountsRequest{}
	this.OauthStateId = oauthStateId
	this.Accounts = accounts
	return &this
}

// NewSandboxOauthSelectAccountsRequestWithDefaults instantiates a new SandboxOauthSelectAccountsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxOauthSelectAccountsRequestWithDefaults() *SandboxOauthSelectAccountsRequest {
	this := SandboxOauthSelectAccountsRequest{}
	return &this
}

// GetOauthStateId returns the OauthStateId field value
func (o *SandboxOauthSelectAccountsRequest) GetOauthStateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OauthStateId
}

// GetOauthStateIdOk returns a tuple with the OauthStateId field value
// and a boolean to check if the value has been set.
func (o *SandboxOauthSelectAccountsRequest) GetOauthStateIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OauthStateId, true
}

// SetOauthStateId sets field value
func (o *SandboxOauthSelectAccountsRequest) SetOauthStateId(v string) {
	o.OauthStateId = v
}

// GetAccounts returns the Accounts field value
func (o *SandboxOauthSelectAccountsRequest) GetAccounts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *SandboxOauthSelectAccountsRequest) GetAccountsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Accounts, true
}

// SetAccounts sets field value
func (o *SandboxOauthSelectAccountsRequest) SetAccounts(v []string) {
	o.Accounts = v
}

func (o SandboxOauthSelectAccountsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["oauth_state_id"] = o.OauthStateId
	}
	if true {
		toSerialize["accounts"] = o.Accounts
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxOauthSelectAccountsRequest struct {
	value *SandboxOauthSelectAccountsRequest
	isSet bool
}

func (v NullableSandboxOauthSelectAccountsRequest) Get() *SandboxOauthSelectAccountsRequest {
	return v.value
}

func (v *NullableSandboxOauthSelectAccountsRequest) Set(val *SandboxOauthSelectAccountsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxOauthSelectAccountsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxOauthSelectAccountsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxOauthSelectAccountsRequest(val *SandboxOauthSelectAccountsRequest) *NullableSandboxOauthSelectAccountsRequest {
	return &NullableSandboxOauthSelectAccountsRequest{value: val, isSet: true}
}

func (v NullableSandboxOauthSelectAccountsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxOauthSelectAccountsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


