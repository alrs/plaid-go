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

// DepositSwitchAltCreateRequest DepositSwitchAltCreateRequest defines the request schema for `/deposit_switch/alt/create`
type DepositSwitchAltCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	TargetAccount DepositSwitchTargetAccount `json:"target_account"`
	TargetUser DepositSwitchTargetUser `json:"target_user"`
	Options *DepositSwitchCreateRequestOptions `json:"options,omitempty"`
	// ISO-3166-1 alpha-2 country code standard.
	CountryCode NullableString `json:"country_code,omitempty"`
}

// NewDepositSwitchAltCreateRequest instantiates a new DepositSwitchAltCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositSwitchAltCreateRequest(targetAccount DepositSwitchTargetAccount, targetUser DepositSwitchTargetUser) *DepositSwitchAltCreateRequest {
	this := DepositSwitchAltCreateRequest{}
	this.TargetAccount = targetAccount
	this.TargetUser = targetUser
	return &this
}

// NewDepositSwitchAltCreateRequestWithDefaults instantiates a new DepositSwitchAltCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositSwitchAltCreateRequestWithDefaults() *DepositSwitchAltCreateRequest {
	this := DepositSwitchAltCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *DepositSwitchAltCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositSwitchAltCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *DepositSwitchAltCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *DepositSwitchAltCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *DepositSwitchAltCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositSwitchAltCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *DepositSwitchAltCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *DepositSwitchAltCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetTargetAccount returns the TargetAccount field value
func (o *DepositSwitchAltCreateRequest) GetTargetAccount() DepositSwitchTargetAccount {
	if o == nil {
		var ret DepositSwitchTargetAccount
		return ret
	}

	return o.TargetAccount
}

// GetTargetAccountOk returns a tuple with the TargetAccount field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchAltCreateRequest) GetTargetAccountOk() (*DepositSwitchTargetAccount, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetAccount, true
}

// SetTargetAccount sets field value
func (o *DepositSwitchAltCreateRequest) SetTargetAccount(v DepositSwitchTargetAccount) {
	o.TargetAccount = v
}

// GetTargetUser returns the TargetUser field value
func (o *DepositSwitchAltCreateRequest) GetTargetUser() DepositSwitchTargetUser {
	if o == nil {
		var ret DepositSwitchTargetUser
		return ret
	}

	return o.TargetUser
}

// GetTargetUserOk returns a tuple with the TargetUser field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchAltCreateRequest) GetTargetUserOk() (*DepositSwitchTargetUser, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetUser, true
}

// SetTargetUser sets field value
func (o *DepositSwitchAltCreateRequest) SetTargetUser(v DepositSwitchTargetUser) {
	o.TargetUser = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *DepositSwitchAltCreateRequest) GetOptions() DepositSwitchCreateRequestOptions {
	if o == nil || o.Options == nil {
		var ret DepositSwitchCreateRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositSwitchAltCreateRequest) GetOptionsOk() (*DepositSwitchCreateRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *DepositSwitchAltCreateRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given DepositSwitchCreateRequestOptions and assigns it to the Options field.
func (o *DepositSwitchAltCreateRequest) SetOptions(v DepositSwitchCreateRequestOptions) {
	o.Options = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DepositSwitchAltCreateRequest) GetCountryCode() string {
	if o == nil || o.CountryCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.CountryCode.Get()
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepositSwitchAltCreateRequest) GetCountryCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CountryCode.Get(), o.CountryCode.IsSet()
}

// HasCountryCode returns a boolean if a field has been set.
func (o *DepositSwitchAltCreateRequest) HasCountryCode() bool {
	if o != nil && o.CountryCode.IsSet() {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given NullableString and assigns it to the CountryCode field.
func (o *DepositSwitchAltCreateRequest) SetCountryCode(v string) {
	o.CountryCode.Set(&v)
}
// SetCountryCodeNil sets the value for CountryCode to be an explicit nil
func (o *DepositSwitchAltCreateRequest) SetCountryCodeNil() {
	o.CountryCode.Set(nil)
}

// UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
func (o *DepositSwitchAltCreateRequest) UnsetCountryCode() {
	o.CountryCode.Unset()
}

func (o DepositSwitchAltCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["target_account"] = o.TargetAccount
	}
	if true {
		toSerialize["target_user"] = o.TargetUser
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.CountryCode.IsSet() {
		toSerialize["country_code"] = o.CountryCode.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDepositSwitchAltCreateRequest struct {
	value *DepositSwitchAltCreateRequest
	isSet bool
}

func (v NullableDepositSwitchAltCreateRequest) Get() *DepositSwitchAltCreateRequest {
	return v.value
}

func (v *NullableDepositSwitchAltCreateRequest) Set(val *DepositSwitchAltCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositSwitchAltCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositSwitchAltCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositSwitchAltCreateRequest(val *DepositSwitchAltCreateRequest) *NullableDepositSwitchAltCreateRequest {
	return &NullableDepositSwitchAltCreateRequest{value: val, isSet: true}
}

func (v NullableDepositSwitchAltCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositSwitchAltCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


