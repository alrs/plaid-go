/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.343.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// PartnerCustomerOAuthInstitutionsGetRequest Request schema for `/partner/customer/oauth_institutions/get`.
type PartnerCustomerOAuthInstitutionsGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	EndCustomerClientId string `json:"end_customer_client_id"`
}

// NewPartnerCustomerOAuthInstitutionsGetRequest instantiates a new PartnerCustomerOAuthInstitutionsGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerCustomerOAuthInstitutionsGetRequest(endCustomerClientId string) *PartnerCustomerOAuthInstitutionsGetRequest {
	this := PartnerCustomerOAuthInstitutionsGetRequest{}
	this.EndCustomerClientId = endCustomerClientId
	return &this
}

// NewPartnerCustomerOAuthInstitutionsGetRequestWithDefaults instantiates a new PartnerCustomerOAuthInstitutionsGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerCustomerOAuthInstitutionsGetRequestWithDefaults() *PartnerCustomerOAuthInstitutionsGetRequest {
	this := PartnerCustomerOAuthInstitutionsGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PartnerCustomerOAuthInstitutionsGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCustomerOAuthInstitutionsGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PartnerCustomerOAuthInstitutionsGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PartnerCustomerOAuthInstitutionsGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *PartnerCustomerOAuthInstitutionsGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCustomerOAuthInstitutionsGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *PartnerCustomerOAuthInstitutionsGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *PartnerCustomerOAuthInstitutionsGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetEndCustomerClientId returns the EndCustomerClientId field value
func (o *PartnerCustomerOAuthInstitutionsGetRequest) GetEndCustomerClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndCustomerClientId
}

// GetEndCustomerClientIdOk returns a tuple with the EndCustomerClientId field value
// and a boolean to check if the value has been set.
func (o *PartnerCustomerOAuthInstitutionsGetRequest) GetEndCustomerClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndCustomerClientId, true
}

// SetEndCustomerClientId sets field value
func (o *PartnerCustomerOAuthInstitutionsGetRequest) SetEndCustomerClientId(v string) {
	o.EndCustomerClientId = v
}

func (o PartnerCustomerOAuthInstitutionsGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["end_customer_client_id"] = o.EndCustomerClientId
	}
	return json.Marshal(toSerialize)
}

type NullablePartnerCustomerOAuthInstitutionsGetRequest struct {
	value *PartnerCustomerOAuthInstitutionsGetRequest
	isSet bool
}

func (v NullablePartnerCustomerOAuthInstitutionsGetRequest) Get() *PartnerCustomerOAuthInstitutionsGetRequest {
	return v.value
}

func (v *NullablePartnerCustomerOAuthInstitutionsGetRequest) Set(val *PartnerCustomerOAuthInstitutionsGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerCustomerOAuthInstitutionsGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerCustomerOAuthInstitutionsGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerCustomerOAuthInstitutionsGetRequest(val *PartnerCustomerOAuthInstitutionsGetRequest) *NullablePartnerCustomerOAuthInstitutionsGetRequest {
	return &NullablePartnerCustomerOAuthInstitutionsGetRequest{value: val, isSet: true}
}

func (v NullablePartnerCustomerOAuthInstitutionsGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerCustomerOAuthInstitutionsGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


