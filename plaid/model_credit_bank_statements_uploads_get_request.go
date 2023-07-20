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

// CreditBankStatementsUploadsGetRequest CreditBankStatementsUploadsGetRequest defines the request schema for `/credit/bank_statements/uploads/get`
type CreditBankStatementsUploadsGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The user token associated with the User data is being requested for.
	UserToken string `json:"user_token"`
}

// NewCreditBankStatementsUploadsGetRequest instantiates a new CreditBankStatementsUploadsGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankStatementsUploadsGetRequest(userToken string) *CreditBankStatementsUploadsGetRequest {
	this := CreditBankStatementsUploadsGetRequest{}
	this.UserToken = userToken
	return &this
}

// NewCreditBankStatementsUploadsGetRequestWithDefaults instantiates a new CreditBankStatementsUploadsGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankStatementsUploadsGetRequestWithDefaults() *CreditBankStatementsUploadsGetRequest {
	this := CreditBankStatementsUploadsGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CreditBankStatementsUploadsGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankStatementsUploadsGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CreditBankStatementsUploadsGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CreditBankStatementsUploadsGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *CreditBankStatementsUploadsGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankStatementsUploadsGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *CreditBankStatementsUploadsGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *CreditBankStatementsUploadsGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetUserToken returns the UserToken field value
func (o *CreditBankStatementsUploadsGetRequest) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *CreditBankStatementsUploadsGetRequest) GetUserTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value
func (o *CreditBankStatementsUploadsGetRequest) SetUserToken(v string) {
	o.UserToken = v
}

func (o CreditBankStatementsUploadsGetRequest) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableCreditBankStatementsUploadsGetRequest struct {
	value *CreditBankStatementsUploadsGetRequest
	isSet bool
}

func (v NullableCreditBankStatementsUploadsGetRequest) Get() *CreditBankStatementsUploadsGetRequest {
	return v.value
}

func (v *NullableCreditBankStatementsUploadsGetRequest) Set(val *CreditBankStatementsUploadsGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankStatementsUploadsGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankStatementsUploadsGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankStatementsUploadsGetRequest(val *CreditBankStatementsUploadsGetRequest) *NullableCreditBankStatementsUploadsGetRequest {
	return &NullableCreditBankStatementsUploadsGetRequest{value: val, isSet: true}
}

func (v NullableCreditBankStatementsUploadsGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankStatementsUploadsGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


