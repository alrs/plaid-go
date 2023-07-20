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

// WatchlistScreeningSearchTerms Search terms for creating an individual watchlist screening
type WatchlistScreeningSearchTerms struct {
	// ID of the associated program.
	WatchlistProgramId string `json:"watchlist_program_id"`
	// The legal name of the individual being screened.
	LegalName string `json:"legal_name"`
	// A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).
	DateOfBirth NullableString `json:"date_of_birth"`
	// The numeric or alphanumeric identifier associated with this document.
	DocumentNumber NullableString `json:"document_number"`
	// Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.
	Country NullableString `json:"country"`
	// The current version of the search terms. Starts at `1` and increments with each edit to `search_terms`.
	Version int32 `json:"version"`
	AdditionalProperties map[string]interface{}
}

type _WatchlistScreeningSearchTerms WatchlistScreeningSearchTerms

// NewWatchlistScreeningSearchTerms instantiates a new WatchlistScreeningSearchTerms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningSearchTerms(watchlistProgramId string, legalName string, dateOfBirth NullableString, documentNumber NullableString, country NullableString, version int32) *WatchlistScreeningSearchTerms {
	this := WatchlistScreeningSearchTerms{}
	this.WatchlistProgramId = watchlistProgramId
	this.LegalName = legalName
	this.DateOfBirth = dateOfBirth
	this.DocumentNumber = documentNumber
	this.Country = country
	this.Version = version
	return &this
}

// NewWatchlistScreeningSearchTermsWithDefaults instantiates a new WatchlistScreeningSearchTerms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningSearchTermsWithDefaults() *WatchlistScreeningSearchTerms {
	this := WatchlistScreeningSearchTerms{}
	return &this
}

// GetWatchlistProgramId returns the WatchlistProgramId field value
func (o *WatchlistScreeningSearchTerms) GetWatchlistProgramId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WatchlistProgramId
}

// GetWatchlistProgramIdOk returns a tuple with the WatchlistProgramId field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningSearchTerms) GetWatchlistProgramIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WatchlistProgramId, true
}

// SetWatchlistProgramId sets field value
func (o *WatchlistScreeningSearchTerms) SetWatchlistProgramId(v string) {
	o.WatchlistProgramId = v
}

// GetLegalName returns the LegalName field value
func (o *WatchlistScreeningSearchTerms) GetLegalName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LegalName
}

// GetLegalNameOk returns a tuple with the LegalName field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningSearchTerms) GetLegalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LegalName, true
}

// SetLegalName sets field value
func (o *WatchlistScreeningSearchTerms) SetLegalName(v string) {
	o.LegalName = v
}

// GetDateOfBirth returns the DateOfBirth field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchlistScreeningSearchTerms) GetDateOfBirth() string {
	if o == nil || o.DateOfBirth.Get() == nil {
		var ret string
		return ret
	}

	return *o.DateOfBirth.Get()
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningSearchTerms) GetDateOfBirthOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DateOfBirth.Get(), o.DateOfBirth.IsSet()
}

// SetDateOfBirth sets field value
func (o *WatchlistScreeningSearchTerms) SetDateOfBirth(v string) {
	o.DateOfBirth.Set(&v)
}

// GetDocumentNumber returns the DocumentNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchlistScreeningSearchTerms) GetDocumentNumber() string {
	if o == nil || o.DocumentNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.DocumentNumber.Get()
}

// GetDocumentNumberOk returns a tuple with the DocumentNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningSearchTerms) GetDocumentNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DocumentNumber.Get(), o.DocumentNumber.IsSet()
}

// SetDocumentNumber sets field value
func (o *WatchlistScreeningSearchTerms) SetDocumentNumber(v string) {
	o.DocumentNumber.Set(&v)
}

// GetCountry returns the Country field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchlistScreeningSearchTerms) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}

	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningSearchTerms) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// SetCountry sets field value
func (o *WatchlistScreeningSearchTerms) SetCountry(v string) {
	o.Country.Set(&v)
}

// GetVersion returns the Version field value
func (o *WatchlistScreeningSearchTerms) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningSearchTerms) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *WatchlistScreeningSearchTerms) SetVersion(v int32) {
	o.Version = v
}

func (o WatchlistScreeningSearchTerms) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["watchlist_program_id"] = o.WatchlistProgramId
	}
	if true {
		toSerialize["legal_name"] = o.LegalName
	}
	if true {
		toSerialize["date_of_birth"] = o.DateOfBirth.Get()
	}
	if true {
		toSerialize["document_number"] = o.DocumentNumber.Get()
	}
	if true {
		toSerialize["country"] = o.Country.Get()
	}
	if true {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WatchlistScreeningSearchTerms) UnmarshalJSON(bytes []byte) (err error) {
	varWatchlistScreeningSearchTerms := _WatchlistScreeningSearchTerms{}

	if err = json.Unmarshal(bytes, &varWatchlistScreeningSearchTerms); err == nil {
		*o = WatchlistScreeningSearchTerms(varWatchlistScreeningSearchTerms)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "watchlist_program_id")
		delete(additionalProperties, "legal_name")
		delete(additionalProperties, "date_of_birth")
		delete(additionalProperties, "document_number")
		delete(additionalProperties, "country")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWatchlistScreeningSearchTerms struct {
	value *WatchlistScreeningSearchTerms
	isSet bool
}

func (v NullableWatchlistScreeningSearchTerms) Get() *WatchlistScreeningSearchTerms {
	return v.value
}

func (v *NullableWatchlistScreeningSearchTerms) Set(val *WatchlistScreeningSearchTerms) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningSearchTerms) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningSearchTerms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningSearchTerms(val *WatchlistScreeningSearchTerms) *NullableWatchlistScreeningSearchTerms {
	return &NullableWatchlistScreeningSearchTerms{value: val, isSet: true}
}

func (v NullableWatchlistScreeningSearchTerms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningSearchTerms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


