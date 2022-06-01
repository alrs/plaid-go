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

// PhysicalDocumentExtractedDataAnalysis Analysis of the data extracted from the submitted document.
type PhysicalDocumentExtractedDataAnalysis struct {
	Name DocumentNameMatchCode `json:"name"`
	DateOfBirth DocumentDateOfBirthMatchCode `json:"date_of_birth"`
	ExpirationDate ExpirationDate `json:"expiration_date"`
	IssuingCountry IssuingCountry `json:"issuing_country"`
}

// NewPhysicalDocumentExtractedDataAnalysis instantiates a new PhysicalDocumentExtractedDataAnalysis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhysicalDocumentExtractedDataAnalysis(name DocumentNameMatchCode, dateOfBirth DocumentDateOfBirthMatchCode, expirationDate ExpirationDate, issuingCountry IssuingCountry) *PhysicalDocumentExtractedDataAnalysis {
	this := PhysicalDocumentExtractedDataAnalysis{}
	this.Name = name
	this.DateOfBirth = dateOfBirth
	this.ExpirationDate = expirationDate
	this.IssuingCountry = issuingCountry
	return &this
}

// NewPhysicalDocumentExtractedDataAnalysisWithDefaults instantiates a new PhysicalDocumentExtractedDataAnalysis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhysicalDocumentExtractedDataAnalysisWithDefaults() *PhysicalDocumentExtractedDataAnalysis {
	this := PhysicalDocumentExtractedDataAnalysis{}
	return &this
}

// GetName returns the Name field value
func (o *PhysicalDocumentExtractedDataAnalysis) GetName() DocumentNameMatchCode {
	if o == nil {
		var ret DocumentNameMatchCode
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PhysicalDocumentExtractedDataAnalysis) GetNameOk() (*DocumentNameMatchCode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PhysicalDocumentExtractedDataAnalysis) SetName(v DocumentNameMatchCode) {
	o.Name = v
}

// GetDateOfBirth returns the DateOfBirth field value
func (o *PhysicalDocumentExtractedDataAnalysis) GetDateOfBirth() DocumentDateOfBirthMatchCode {
	if o == nil {
		var ret DocumentDateOfBirthMatchCode
		return ret
	}

	return o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value
// and a boolean to check if the value has been set.
func (o *PhysicalDocumentExtractedDataAnalysis) GetDateOfBirthOk() (*DocumentDateOfBirthMatchCode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DateOfBirth, true
}

// SetDateOfBirth sets field value
func (o *PhysicalDocumentExtractedDataAnalysis) SetDateOfBirth(v DocumentDateOfBirthMatchCode) {
	o.DateOfBirth = v
}

// GetExpirationDate returns the ExpirationDate field value
func (o *PhysicalDocumentExtractedDataAnalysis) GetExpirationDate() ExpirationDate {
	if o == nil {
		var ret ExpirationDate
		return ret
	}

	return o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value
// and a boolean to check if the value has been set.
func (o *PhysicalDocumentExtractedDataAnalysis) GetExpirationDateOk() (*ExpirationDate, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExpirationDate, true
}

// SetExpirationDate sets field value
func (o *PhysicalDocumentExtractedDataAnalysis) SetExpirationDate(v ExpirationDate) {
	o.ExpirationDate = v
}

// GetIssuingCountry returns the IssuingCountry field value
func (o *PhysicalDocumentExtractedDataAnalysis) GetIssuingCountry() IssuingCountry {
	if o == nil {
		var ret IssuingCountry
		return ret
	}

	return o.IssuingCountry
}

// GetIssuingCountryOk returns a tuple with the IssuingCountry field value
// and a boolean to check if the value has been set.
func (o *PhysicalDocumentExtractedDataAnalysis) GetIssuingCountryOk() (*IssuingCountry, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IssuingCountry, true
}

// SetIssuingCountry sets field value
func (o *PhysicalDocumentExtractedDataAnalysis) SetIssuingCountry(v IssuingCountry) {
	o.IssuingCountry = v
}

func (o PhysicalDocumentExtractedDataAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["date_of_birth"] = o.DateOfBirth
	}
	if true {
		toSerialize["expiration_date"] = o.ExpirationDate
	}
	if true {
		toSerialize["issuing_country"] = o.IssuingCountry
	}
	return json.Marshal(toSerialize)
}

type NullablePhysicalDocumentExtractedDataAnalysis struct {
	value *PhysicalDocumentExtractedDataAnalysis
	isSet bool
}

func (v NullablePhysicalDocumentExtractedDataAnalysis) Get() *PhysicalDocumentExtractedDataAnalysis {
	return v.value
}

func (v *NullablePhysicalDocumentExtractedDataAnalysis) Set(val *PhysicalDocumentExtractedDataAnalysis) {
	v.value = val
	v.isSet = true
}

func (v NullablePhysicalDocumentExtractedDataAnalysis) IsSet() bool {
	return v.isSet
}

func (v *NullablePhysicalDocumentExtractedDataAnalysis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhysicalDocumentExtractedDataAnalysis(val *PhysicalDocumentExtractedDataAnalysis) *NullablePhysicalDocumentExtractedDataAnalysis {
	return &NullablePhysicalDocumentExtractedDataAnalysis{value: val, isSet: true}
}

func (v NullablePhysicalDocumentExtractedDataAnalysis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhysicalDocumentExtractedDataAnalysis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


