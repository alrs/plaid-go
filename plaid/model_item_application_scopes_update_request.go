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

// ItemApplicationScopesUpdateRequest ItemApplicationScopesUpdateRequest defines the request schema for `/item/application/scopes/update`
type ItemApplicationScopesUpdateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	// This field will map to the application ID that is returned from /item/applications/list, or provided to the institution in an oauth redirect.
	ApplicationId string `json:"application_id"`
	Scopes Scopes `json:"scopes"`
	// When scopes are updated during enrollment, this field must be populated with the state sent to the partner in the OAuth Login URI. This field is required when the context is `ENROLLMENT`.
	State *string `json:"state,omitempty"`
	Context ScopesContext `json:"context"`
}

// NewItemApplicationScopesUpdateRequest instantiates a new ItemApplicationScopesUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemApplicationScopesUpdateRequest(accessToken string, applicationId string, scopes Scopes, context ScopesContext) *ItemApplicationScopesUpdateRequest {
	this := ItemApplicationScopesUpdateRequest{}
	this.AccessToken = accessToken
	this.ApplicationId = applicationId
	this.Scopes = scopes
	this.Context = context
	return &this
}

// NewItemApplicationScopesUpdateRequestWithDefaults instantiates a new ItemApplicationScopesUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemApplicationScopesUpdateRequestWithDefaults() *ItemApplicationScopesUpdateRequest {
	this := ItemApplicationScopesUpdateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ItemApplicationScopesUpdateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemApplicationScopesUpdateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ItemApplicationScopesUpdateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ItemApplicationScopesUpdateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ItemApplicationScopesUpdateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemApplicationScopesUpdateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ItemApplicationScopesUpdateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ItemApplicationScopesUpdateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccessToken returns the AccessToken field value
func (o *ItemApplicationScopesUpdateRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *ItemApplicationScopesUpdateRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *ItemApplicationScopesUpdateRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetApplicationId returns the ApplicationId field value
func (o *ItemApplicationScopesUpdateRequest) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *ItemApplicationScopesUpdateRequest) GetApplicationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *ItemApplicationScopesUpdateRequest) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetScopes returns the Scopes field value
func (o *ItemApplicationScopesUpdateRequest) GetScopes() Scopes {
	if o == nil {
		var ret Scopes
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *ItemApplicationScopesUpdateRequest) GetScopesOk() (*Scopes, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Scopes, true
}

// SetScopes sets field value
func (o *ItemApplicationScopesUpdateRequest) SetScopes(v Scopes) {
	o.Scopes = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ItemApplicationScopesUpdateRequest) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemApplicationScopesUpdateRequest) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ItemApplicationScopesUpdateRequest) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ItemApplicationScopesUpdateRequest) SetState(v string) {
	o.State = &v
}

// GetContext returns the Context field value
func (o *ItemApplicationScopesUpdateRequest) GetContext() ScopesContext {
	if o == nil {
		var ret ScopesContext
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *ItemApplicationScopesUpdateRequest) GetContextOk() (*ScopesContext, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *ItemApplicationScopesUpdateRequest) SetContext(v ScopesContext) {
	o.Context = v
}

func (o ItemApplicationScopesUpdateRequest) MarshalJSON() ([]byte, error) {
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
		toSerialize["application_id"] = o.ApplicationId
	}
	if true {
		toSerialize["scopes"] = o.Scopes
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["context"] = o.Context
	}
	return json.Marshal(toSerialize)
}

type NullableItemApplicationScopesUpdateRequest struct {
	value *ItemApplicationScopesUpdateRequest
	isSet bool
}

func (v NullableItemApplicationScopesUpdateRequest) Get() *ItemApplicationScopesUpdateRequest {
	return v.value
}

func (v *NullableItemApplicationScopesUpdateRequest) Set(val *ItemApplicationScopesUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableItemApplicationScopesUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableItemApplicationScopesUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemApplicationScopesUpdateRequest(val *ItemApplicationScopesUpdateRequest) *NullableItemApplicationScopesUpdateRequest {
	return &NullableItemApplicationScopesUpdateRequest{value: val, isSet: true}
}

func (v NullableItemApplicationScopesUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemApplicationScopesUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


