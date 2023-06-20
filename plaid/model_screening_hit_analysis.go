/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.379.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// ScreeningHitAnalysis Analysis information describing why a screening hit matched the provided user information
type ScreeningHitAnalysis struct {
	DatesOfBirth *MatchSummaryCode `json:"dates_of_birth,omitempty"`
	Documents *MatchSummaryCode `json:"documents,omitempty"`
	Locations *MatchSummaryCode `json:"locations,omitempty"`
	Names *MatchSummaryCode `json:"names,omitempty"`
	// The version of the screening's `search_terms` that were compared when the screening hit was added. screening hits are immutable once they have been reviewed. If changes are detected due to updates to the screening's `search_terms`, the associated program, or the list's source data prior to review, the screening hit will be updated to reflect those changes.
	SearchTermsVersion int32 `json:"search_terms_version"`
	AdditionalProperties map[string]interface{}
}

type _ScreeningHitAnalysis ScreeningHitAnalysis

// NewScreeningHitAnalysis instantiates a new ScreeningHitAnalysis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScreeningHitAnalysis(searchTermsVersion int32) *ScreeningHitAnalysis {
	this := ScreeningHitAnalysis{}
	this.SearchTermsVersion = searchTermsVersion
	return &this
}

// NewScreeningHitAnalysisWithDefaults instantiates a new ScreeningHitAnalysis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScreeningHitAnalysisWithDefaults() *ScreeningHitAnalysis {
	this := ScreeningHitAnalysis{}
	return &this
}

// GetDatesOfBirth returns the DatesOfBirth field value if set, zero value otherwise.
func (o *ScreeningHitAnalysis) GetDatesOfBirth() MatchSummaryCode {
	if o == nil || o.DatesOfBirth == nil {
		var ret MatchSummaryCode
		return ret
	}
	return *o.DatesOfBirth
}

// GetDatesOfBirthOk returns a tuple with the DatesOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreeningHitAnalysis) GetDatesOfBirthOk() (*MatchSummaryCode, bool) {
	if o == nil || o.DatesOfBirth == nil {
		return nil, false
	}
	return o.DatesOfBirth, true
}

// HasDatesOfBirth returns a boolean if a field has been set.
func (o *ScreeningHitAnalysis) HasDatesOfBirth() bool {
	if o != nil && o.DatesOfBirth != nil {
		return true
	}

	return false
}

// SetDatesOfBirth gets a reference to the given MatchSummaryCode and assigns it to the DatesOfBirth field.
func (o *ScreeningHitAnalysis) SetDatesOfBirth(v MatchSummaryCode) {
	o.DatesOfBirth = &v
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *ScreeningHitAnalysis) GetDocuments() MatchSummaryCode {
	if o == nil || o.Documents == nil {
		var ret MatchSummaryCode
		return ret
	}
	return *o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreeningHitAnalysis) GetDocumentsOk() (*MatchSummaryCode, bool) {
	if o == nil || o.Documents == nil {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *ScreeningHitAnalysis) HasDocuments() bool {
	if o != nil && o.Documents != nil {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given MatchSummaryCode and assigns it to the Documents field.
func (o *ScreeningHitAnalysis) SetDocuments(v MatchSummaryCode) {
	o.Documents = &v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *ScreeningHitAnalysis) GetLocations() MatchSummaryCode {
	if o == nil || o.Locations == nil {
		var ret MatchSummaryCode
		return ret
	}
	return *o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreeningHitAnalysis) GetLocationsOk() (*MatchSummaryCode, bool) {
	if o == nil || o.Locations == nil {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *ScreeningHitAnalysis) HasLocations() bool {
	if o != nil && o.Locations != nil {
		return true
	}

	return false
}

// SetLocations gets a reference to the given MatchSummaryCode and assigns it to the Locations field.
func (o *ScreeningHitAnalysis) SetLocations(v MatchSummaryCode) {
	o.Locations = &v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *ScreeningHitAnalysis) GetNames() MatchSummaryCode {
	if o == nil || o.Names == nil {
		var ret MatchSummaryCode
		return ret
	}
	return *o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreeningHitAnalysis) GetNamesOk() (*MatchSummaryCode, bool) {
	if o == nil || o.Names == nil {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *ScreeningHitAnalysis) HasNames() bool {
	if o != nil && o.Names != nil {
		return true
	}

	return false
}

// SetNames gets a reference to the given MatchSummaryCode and assigns it to the Names field.
func (o *ScreeningHitAnalysis) SetNames(v MatchSummaryCode) {
	o.Names = &v
}

// GetSearchTermsVersion returns the SearchTermsVersion field value
func (o *ScreeningHitAnalysis) GetSearchTermsVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SearchTermsVersion
}

// GetSearchTermsVersionOk returns a tuple with the SearchTermsVersion field value
// and a boolean to check if the value has been set.
func (o *ScreeningHitAnalysis) GetSearchTermsVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SearchTermsVersion, true
}

// SetSearchTermsVersion sets field value
func (o *ScreeningHitAnalysis) SetSearchTermsVersion(v int32) {
	o.SearchTermsVersion = v
}

func (o ScreeningHitAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DatesOfBirth != nil {
		toSerialize["dates_of_birth"] = o.DatesOfBirth
	}
	if o.Documents != nil {
		toSerialize["documents"] = o.Documents
	}
	if o.Locations != nil {
		toSerialize["locations"] = o.Locations
	}
	if o.Names != nil {
		toSerialize["names"] = o.Names
	}
	if true {
		toSerialize["search_terms_version"] = o.SearchTermsVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ScreeningHitAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	varScreeningHitAnalysis := _ScreeningHitAnalysis{}

	if err = json.Unmarshal(bytes, &varScreeningHitAnalysis); err == nil {
		*o = ScreeningHitAnalysis(varScreeningHitAnalysis)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dates_of_birth")
		delete(additionalProperties, "documents")
		delete(additionalProperties, "locations")
		delete(additionalProperties, "names")
		delete(additionalProperties, "search_terms_version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScreeningHitAnalysis struct {
	value *ScreeningHitAnalysis
	isSet bool
}

func (v NullableScreeningHitAnalysis) Get() *ScreeningHitAnalysis {
	return v.value
}

func (v *NullableScreeningHitAnalysis) Set(val *ScreeningHitAnalysis) {
	v.value = val
	v.isSet = true
}

func (v NullableScreeningHitAnalysis) IsSet() bool {
	return v.isSet
}

func (v *NullableScreeningHitAnalysis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScreeningHitAnalysis(val *ScreeningHitAnalysis) *NullableScreeningHitAnalysis {
	return &NullableScreeningHitAnalysis{value: val, isSet: true}
}

func (v NullableScreeningHitAnalysis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScreeningHitAnalysis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


