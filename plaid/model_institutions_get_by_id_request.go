/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.64.14
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// InstitutionsGetByIdRequest InstitutionsGetByIdRequest defines the request schema for `/institutions/get_by_id`
type InstitutionsGetByIdRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The ID of the institution to get details about
	InstitutionId string `json:"institution_id"`
	// Specify an array of Plaid-supported country codes this institution supports, using the ISO-3166-1 alpha-2 country code standard. In API versions 2019-05-29 and earlier, the `country_codes` parameter is an optional parameter within the `options` object and will default to `[US]` if it is not supplied. 
	CountryCodes []CountryCode `json:"country_codes"`
	Options *InstitutionsGetByIdRequestOptions `json:"options,omitempty"`
}

// NewInstitutionsGetByIdRequest instantiates a new InstitutionsGetByIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstitutionsGetByIdRequest(institutionId string, countryCodes []CountryCode) *InstitutionsGetByIdRequest {
	this := InstitutionsGetByIdRequest{}
	this.InstitutionId = institutionId
	this.CountryCodes = countryCodes
	return &this
}

// NewInstitutionsGetByIdRequestWithDefaults instantiates a new InstitutionsGetByIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstitutionsGetByIdRequestWithDefaults() *InstitutionsGetByIdRequest {
	this := InstitutionsGetByIdRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *InstitutionsGetByIdRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsGetByIdRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *InstitutionsGetByIdRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *InstitutionsGetByIdRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *InstitutionsGetByIdRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsGetByIdRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *InstitutionsGetByIdRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *InstitutionsGetByIdRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetInstitutionId returns the InstitutionId field value
func (o *InstitutionsGetByIdRequest) GetInstitutionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value
// and a boolean to check if the value has been set.
func (o *InstitutionsGetByIdRequest) GetInstitutionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstitutionId, true
}

// SetInstitutionId sets field value
func (o *InstitutionsGetByIdRequest) SetInstitutionId(v string) {
	o.InstitutionId = v
}

// GetCountryCodes returns the CountryCodes field value
func (o *InstitutionsGetByIdRequest) GetCountryCodes() []CountryCode {
	if o == nil {
		var ret []CountryCode
		return ret
	}

	return o.CountryCodes
}

// GetCountryCodesOk returns a tuple with the CountryCodes field value
// and a boolean to check if the value has been set.
func (o *InstitutionsGetByIdRequest) GetCountryCodesOk() (*[]CountryCode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CountryCodes, true
}

// SetCountryCodes sets field value
func (o *InstitutionsGetByIdRequest) SetCountryCodes(v []CountryCode) {
	o.CountryCodes = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *InstitutionsGetByIdRequest) GetOptions() InstitutionsGetByIdRequestOptions {
	if o == nil || o.Options == nil {
		var ret InstitutionsGetByIdRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsGetByIdRequest) GetOptionsOk() (*InstitutionsGetByIdRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *InstitutionsGetByIdRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given InstitutionsGetByIdRequestOptions and assigns it to the Options field.
func (o *InstitutionsGetByIdRequest) SetOptions(v InstitutionsGetByIdRequestOptions) {
	o.Options = &v
}

func (o InstitutionsGetByIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["institution_id"] = o.InstitutionId
	}
	if true {
		toSerialize["country_codes"] = o.CountryCodes
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableInstitutionsGetByIdRequest struct {
	value *InstitutionsGetByIdRequest
	isSet bool
}

func (v NullableInstitutionsGetByIdRequest) Get() *InstitutionsGetByIdRequest {
	return v.value
}

func (v *NullableInstitutionsGetByIdRequest) Set(val *InstitutionsGetByIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInstitutionsGetByIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInstitutionsGetByIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstitutionsGetByIdRequest(val *InstitutionsGetByIdRequest) *NullableInstitutionsGetByIdRequest {
	return &NullableInstitutionsGetByIdRequest{value: val, isSet: true}
}

func (v NullableInstitutionsGetByIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstitutionsGetByIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


