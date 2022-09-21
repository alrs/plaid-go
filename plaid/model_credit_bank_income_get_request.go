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

// CreditBankIncomeGetRequest CreditBankIncomeGetRequest defines the request schema for `/credit/bank_income/get`.
type CreditBankIncomeGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The user token associated with the User data is being requested for.
	UserToken *string `json:"user_token,omitempty"`
	Options *CreditBankIncomeGetRequestOptions `json:"options,omitempty"`
}

// NewCreditBankIncomeGetRequest instantiates a new CreditBankIncomeGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankIncomeGetRequest() *CreditBankIncomeGetRequest {
	this := CreditBankIncomeGetRequest{}
	return &this
}

// NewCreditBankIncomeGetRequestWithDefaults instantiates a new CreditBankIncomeGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankIncomeGetRequestWithDefaults() *CreditBankIncomeGetRequest {
	this := CreditBankIncomeGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CreditBankIncomeGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CreditBankIncomeGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CreditBankIncomeGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *CreditBankIncomeGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *CreditBankIncomeGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *CreditBankIncomeGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetUserToken returns the UserToken field value if set, zero value otherwise.
func (o *CreditBankIncomeGetRequest) GetUserToken() string {
	if o == nil || o.UserToken == nil {
		var ret string
		return ret
	}
	return *o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeGetRequest) GetUserTokenOk() (*string, bool) {
	if o == nil || o.UserToken == nil {
		return nil, false
	}
	return o.UserToken, true
}

// HasUserToken returns a boolean if a field has been set.
func (o *CreditBankIncomeGetRequest) HasUserToken() bool {
	if o != nil && o.UserToken != nil {
		return true
	}

	return false
}

// SetUserToken gets a reference to the given string and assigns it to the UserToken field.
func (o *CreditBankIncomeGetRequest) SetUserToken(v string) {
	o.UserToken = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *CreditBankIncomeGetRequest) GetOptions() CreditBankIncomeGetRequestOptions {
	if o == nil || o.Options == nil {
		var ret CreditBankIncomeGetRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeGetRequest) GetOptionsOk() (*CreditBankIncomeGetRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *CreditBankIncomeGetRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given CreditBankIncomeGetRequestOptions and assigns it to the Options field.
func (o *CreditBankIncomeGetRequest) SetOptions(v CreditBankIncomeGetRequestOptions) {
	o.Options = &v
}

func (o CreditBankIncomeGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.UserToken != nil {
		toSerialize["user_token"] = o.UserToken
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableCreditBankIncomeGetRequest struct {
	value *CreditBankIncomeGetRequest
	isSet bool
}

func (v NullableCreditBankIncomeGetRequest) Get() *CreditBankIncomeGetRequest {
	return v.value
}

func (v *NullableCreditBankIncomeGetRequest) Set(val *CreditBankIncomeGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankIncomeGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankIncomeGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankIncomeGetRequest(val *CreditBankIncomeGetRequest) *NullableCreditBankIncomeGetRequest {
	return &NullableCreditBankIncomeGetRequest{value: val, isSet: true}
}

func (v NullableCreditBankIncomeGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankIncomeGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


