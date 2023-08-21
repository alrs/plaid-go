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

// CreditPayrollIncomeRefreshRequest CreditPayrollIncomeRefreshRequest defines the request schema for `/credit/payroll_income/refresh`
type CreditPayrollIncomeRefreshRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The user token associated with the User data is being requested for.
	UserToken string `json:"user_token"`
	Options *CreditPayrollIncomeRefreshRequestOptions `json:"options,omitempty"`
}

// NewCreditPayrollIncomeRefreshRequest instantiates a new CreditPayrollIncomeRefreshRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditPayrollIncomeRefreshRequest(userToken string) *CreditPayrollIncomeRefreshRequest {
	this := CreditPayrollIncomeRefreshRequest{}
	this.UserToken = userToken
	return &this
}

// NewCreditPayrollIncomeRefreshRequestWithDefaults instantiates a new CreditPayrollIncomeRefreshRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditPayrollIncomeRefreshRequestWithDefaults() *CreditPayrollIncomeRefreshRequest {
	this := CreditPayrollIncomeRefreshRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CreditPayrollIncomeRefreshRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditPayrollIncomeRefreshRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CreditPayrollIncomeRefreshRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CreditPayrollIncomeRefreshRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *CreditPayrollIncomeRefreshRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditPayrollIncomeRefreshRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *CreditPayrollIncomeRefreshRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *CreditPayrollIncomeRefreshRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetUserToken returns the UserToken field value
func (o *CreditPayrollIncomeRefreshRequest) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *CreditPayrollIncomeRefreshRequest) GetUserTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value
func (o *CreditPayrollIncomeRefreshRequest) SetUserToken(v string) {
	o.UserToken = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *CreditPayrollIncomeRefreshRequest) GetOptions() CreditPayrollIncomeRefreshRequestOptions {
	if o == nil || o.Options == nil {
		var ret CreditPayrollIncomeRefreshRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditPayrollIncomeRefreshRequest) GetOptionsOk() (*CreditPayrollIncomeRefreshRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *CreditPayrollIncomeRefreshRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given CreditPayrollIncomeRefreshRequestOptions and assigns it to the Options field.
func (o *CreditPayrollIncomeRefreshRequest) SetOptions(v CreditPayrollIncomeRefreshRequestOptions) {
	o.Options = &v
}

func (o CreditPayrollIncomeRefreshRequest) MarshalJSON() ([]byte, error) {
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
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableCreditPayrollIncomeRefreshRequest struct {
	value *CreditPayrollIncomeRefreshRequest
	isSet bool
}

func (v NullableCreditPayrollIncomeRefreshRequest) Get() *CreditPayrollIncomeRefreshRequest {
	return v.value
}

func (v *NullableCreditPayrollIncomeRefreshRequest) Set(val *CreditPayrollIncomeRefreshRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditPayrollIncomeRefreshRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditPayrollIncomeRefreshRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditPayrollIncomeRefreshRequest(val *CreditPayrollIncomeRefreshRequest) *NullableCreditPayrollIncomeRefreshRequest {
	return &NullableCreditPayrollIncomeRefreshRequest{value: val, isSet: true}
}

func (v NullableCreditPayrollIncomeRefreshRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditPayrollIncomeRefreshRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


