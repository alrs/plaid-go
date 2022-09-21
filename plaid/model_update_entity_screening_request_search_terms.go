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

// UpdateEntityScreeningRequestSearchTerms Search terms for editing an entity watchlist screening
type UpdateEntityScreeningRequestSearchTerms struct {
	// ID of the associated entity program.
	EntityWatchlistProgramId string `json:"entity_watchlist_program_id"`
	// The name of the organization being screened.
	LegalName *string `json:"legal_name,omitempty"`
	// The numeric or alphanumeric identifier associated with this document.
	DocumentNumber *string `json:"document_number,omitempty"`
	// A valid email address.
	EmailAddress *string `json:"email_address,omitempty"`
	// Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.
	Country *string `json:"country,omitempty"`
	// A phone number in E.164 format.
	PhoneNumber *string `json:"phone_number,omitempty"`
	// An 'http' or 'https' URL (must begin with either of those).
	Url *string `json:"url,omitempty"`
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId string `json:"client_id"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret string `json:"secret"`
}

// NewUpdateEntityScreeningRequestSearchTerms instantiates a new UpdateEntityScreeningRequestSearchTerms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEntityScreeningRequestSearchTerms(entityWatchlistProgramId string, clientId string, secret string) *UpdateEntityScreeningRequestSearchTerms {
	this := UpdateEntityScreeningRequestSearchTerms{}
	this.EntityWatchlistProgramId = entityWatchlistProgramId
	this.ClientId = clientId
	this.Secret = secret
	return &this
}

// NewUpdateEntityScreeningRequestSearchTermsWithDefaults instantiates a new UpdateEntityScreeningRequestSearchTerms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEntityScreeningRequestSearchTermsWithDefaults() *UpdateEntityScreeningRequestSearchTerms {
	this := UpdateEntityScreeningRequestSearchTerms{}
	return &this
}

// GetEntityWatchlistProgramId returns the EntityWatchlistProgramId field value
func (o *UpdateEntityScreeningRequestSearchTerms) GetEntityWatchlistProgramId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityWatchlistProgramId
}

// GetEntityWatchlistProgramIdOk returns a tuple with the EntityWatchlistProgramId field value
// and a boolean to check if the value has been set.
func (o *UpdateEntityScreeningRequestSearchTerms) GetEntityWatchlistProgramIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EntityWatchlistProgramId, true
}

// SetEntityWatchlistProgramId sets field value
func (o *UpdateEntityScreeningRequestSearchTerms) SetEntityWatchlistProgramId(v string) {
	o.EntityWatchlistProgramId = v
}

// GetLegalName returns the LegalName field value if set, zero value otherwise.
func (o *UpdateEntityScreeningRequestSearchTerms) GetLegalName() string {
	if o == nil || o.LegalName == nil {
		var ret string
		return ret
	}
	return *o.LegalName
}

// GetLegalNameOk returns a tuple with the LegalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEntityScreeningRequestSearchTerms) GetLegalNameOk() (*string, bool) {
	if o == nil || o.LegalName == nil {
		return nil, false
	}
	return o.LegalName, true
}

// HasLegalName returns a boolean if a field has been set.
func (o *UpdateEntityScreeningRequestSearchTerms) HasLegalName() bool {
	if o != nil && o.LegalName != nil {
		return true
	}

	return false
}

// SetLegalName gets a reference to the given string and assigns it to the LegalName field.
func (o *UpdateEntityScreeningRequestSearchTerms) SetLegalName(v string) {
	o.LegalName = &v
}

// GetDocumentNumber returns the DocumentNumber field value if set, zero value otherwise.
func (o *UpdateEntityScreeningRequestSearchTerms) GetDocumentNumber() string {
	if o == nil || o.DocumentNumber == nil {
		var ret string
		return ret
	}
	return *o.DocumentNumber
}

// GetDocumentNumberOk returns a tuple with the DocumentNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEntityScreeningRequestSearchTerms) GetDocumentNumberOk() (*string, bool) {
	if o == nil || o.DocumentNumber == nil {
		return nil, false
	}
	return o.DocumentNumber, true
}

// HasDocumentNumber returns a boolean if a field has been set.
func (o *UpdateEntityScreeningRequestSearchTerms) HasDocumentNumber() bool {
	if o != nil && o.DocumentNumber != nil {
		return true
	}

	return false
}

// SetDocumentNumber gets a reference to the given string and assigns it to the DocumentNumber field.
func (o *UpdateEntityScreeningRequestSearchTerms) SetDocumentNumber(v string) {
	o.DocumentNumber = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *UpdateEntityScreeningRequestSearchTerms) GetEmailAddress() string {
	if o == nil || o.EmailAddress == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEntityScreeningRequestSearchTerms) GetEmailAddressOk() (*string, bool) {
	if o == nil || o.EmailAddress == nil {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *UpdateEntityScreeningRequestSearchTerms) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *UpdateEntityScreeningRequestSearchTerms) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *UpdateEntityScreeningRequestSearchTerms) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEntityScreeningRequestSearchTerms) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *UpdateEntityScreeningRequestSearchTerms) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *UpdateEntityScreeningRequestSearchTerms) SetCountry(v string) {
	o.Country = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *UpdateEntityScreeningRequestSearchTerms) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEntityScreeningRequestSearchTerms) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *UpdateEntityScreeningRequestSearchTerms) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *UpdateEntityScreeningRequestSearchTerms) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UpdateEntityScreeningRequestSearchTerms) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEntityScreeningRequestSearchTerms) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UpdateEntityScreeningRequestSearchTerms) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UpdateEntityScreeningRequestSearchTerms) SetUrl(v string) {
	o.Url = &v
}

// GetClientId returns the ClientId field value
func (o *UpdateEntityScreeningRequestSearchTerms) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *UpdateEntityScreeningRequestSearchTerms) GetClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *UpdateEntityScreeningRequestSearchTerms) SetClientId(v string) {
	o.ClientId = v
}

// GetSecret returns the Secret field value
func (o *UpdateEntityScreeningRequestSearchTerms) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *UpdateEntityScreeningRequestSearchTerms) GetSecretOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *UpdateEntityScreeningRequestSearchTerms) SetSecret(v string) {
	o.Secret = v
}

func (o UpdateEntityScreeningRequestSearchTerms) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entity_watchlist_program_id"] = o.EntityWatchlistProgramId
	}
	if o.LegalName != nil {
		toSerialize["legal_name"] = o.LegalName
	}
	if o.DocumentNumber != nil {
		toSerialize["document_number"] = o.DocumentNumber
	}
	if o.EmailAddress != nil {
		toSerialize["email_address"] = o.EmailAddress
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.PhoneNumber != nil {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateEntityScreeningRequestSearchTerms struct {
	value *UpdateEntityScreeningRequestSearchTerms
	isSet bool
}

func (v NullableUpdateEntityScreeningRequestSearchTerms) Get() *UpdateEntityScreeningRequestSearchTerms {
	return v.value
}

func (v *NullableUpdateEntityScreeningRequestSearchTerms) Set(val *UpdateEntityScreeningRequestSearchTerms) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEntityScreeningRequestSearchTerms) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEntityScreeningRequestSearchTerms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEntityScreeningRequestSearchTerms(val *UpdateEntityScreeningRequestSearchTerms) *NullableUpdateEntityScreeningRequestSearchTerms {
	return &NullableUpdateEntityScreeningRequestSearchTerms{value: val, isSet: true}
}

func (v NullableUpdateEntityScreeningRequestSearchTerms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEntityScreeningRequestSearchTerms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


