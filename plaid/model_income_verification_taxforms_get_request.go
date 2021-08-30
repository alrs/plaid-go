/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.26.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// IncomeVerificationTaxformsGetRequest IncomeVerificationTaxformsGetRequest defines the request schema for `/income/verification/taxforms/get`
type IncomeVerificationTaxformsGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The ID of the verification.
	IncomeVerificationId NullableString `json:"income_verification_id,omitempty"`
	// The access token associated with the Item data is being requested for.
	AccessToken NullableString `json:"access_token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IncomeVerificationTaxformsGetRequest IncomeVerificationTaxformsGetRequest

// NewIncomeVerificationTaxformsGetRequest instantiates a new IncomeVerificationTaxformsGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeVerificationTaxformsGetRequest() *IncomeVerificationTaxformsGetRequest {
	this := IncomeVerificationTaxformsGetRequest{}
	return &this
}

// NewIncomeVerificationTaxformsGetRequestWithDefaults instantiates a new IncomeVerificationTaxformsGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeVerificationTaxformsGetRequestWithDefaults() *IncomeVerificationTaxformsGetRequest {
	this := IncomeVerificationTaxformsGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *IncomeVerificationTaxformsGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationTaxformsGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *IncomeVerificationTaxformsGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *IncomeVerificationTaxformsGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *IncomeVerificationTaxformsGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationTaxformsGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *IncomeVerificationTaxformsGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *IncomeVerificationTaxformsGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetIncomeVerificationId returns the IncomeVerificationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeVerificationTaxformsGetRequest) GetIncomeVerificationId() string {
	if o == nil || o.IncomeVerificationId.Get() == nil {
		var ret string
		return ret
	}
	return *o.IncomeVerificationId.Get()
}

// GetIncomeVerificationIdOk returns a tuple with the IncomeVerificationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeVerificationTaxformsGetRequest) GetIncomeVerificationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IncomeVerificationId.Get(), o.IncomeVerificationId.IsSet()
}

// HasIncomeVerificationId returns a boolean if a field has been set.
func (o *IncomeVerificationTaxformsGetRequest) HasIncomeVerificationId() bool {
	if o != nil && o.IncomeVerificationId.IsSet() {
		return true
	}

	return false
}

// SetIncomeVerificationId gets a reference to the given NullableString and assigns it to the IncomeVerificationId field.
func (o *IncomeVerificationTaxformsGetRequest) SetIncomeVerificationId(v string) {
	o.IncomeVerificationId.Set(&v)
}
// SetIncomeVerificationIdNil sets the value for IncomeVerificationId to be an explicit nil
func (o *IncomeVerificationTaxformsGetRequest) SetIncomeVerificationIdNil() {
	o.IncomeVerificationId.Set(nil)
}

// UnsetIncomeVerificationId ensures that no value is present for IncomeVerificationId, not even an explicit nil
func (o *IncomeVerificationTaxformsGetRequest) UnsetIncomeVerificationId() {
	o.IncomeVerificationId.Unset()
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeVerificationTaxformsGetRequest) GetAccessToken() string {
	if o == nil || o.AccessToken.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeVerificationTaxformsGetRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a field has been set.
func (o *IncomeVerificationTaxformsGetRequest) HasAccessToken() bool {
	if o != nil && o.AccessToken.IsSet() {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given NullableString and assigns it to the AccessToken field.
func (o *IncomeVerificationTaxformsGetRequest) SetAccessToken(v string) {
	o.AccessToken.Set(&v)
}
// SetAccessTokenNil sets the value for AccessToken to be an explicit nil
func (o *IncomeVerificationTaxformsGetRequest) SetAccessTokenNil() {
	o.AccessToken.Set(nil)
}

// UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
func (o *IncomeVerificationTaxformsGetRequest) UnsetAccessToken() {
	o.AccessToken.Unset()
}

func (o IncomeVerificationTaxformsGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.IncomeVerificationId.IsSet() {
		toSerialize["income_verification_id"] = o.IncomeVerificationId.Get()
	}
	if o.AccessToken.IsSet() {
		toSerialize["access_token"] = o.AccessToken.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IncomeVerificationTaxformsGetRequest) UnmarshalJSON(bytes []byte) (err error) {
	varIncomeVerificationTaxformsGetRequest := _IncomeVerificationTaxformsGetRequest{}

	if err = json.Unmarshal(bytes, &varIncomeVerificationTaxformsGetRequest); err == nil {
		*o = IncomeVerificationTaxformsGetRequest(varIncomeVerificationTaxformsGetRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "secret")
		delete(additionalProperties, "income_verification_id")
		delete(additionalProperties, "access_token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIncomeVerificationTaxformsGetRequest struct {
	value *IncomeVerificationTaxformsGetRequest
	isSet bool
}

func (v NullableIncomeVerificationTaxformsGetRequest) Get() *IncomeVerificationTaxformsGetRequest {
	return v.value
}

func (v *NullableIncomeVerificationTaxformsGetRequest) Set(val *IncomeVerificationTaxformsGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationTaxformsGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationTaxformsGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationTaxformsGetRequest(val *IncomeVerificationTaxformsGetRequest) *NullableIncomeVerificationTaxformsGetRequest {
	return &NullableIncomeVerificationTaxformsGetRequest{value: val, isSet: true}
}

func (v NullableIncomeVerificationTaxformsGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationTaxformsGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


