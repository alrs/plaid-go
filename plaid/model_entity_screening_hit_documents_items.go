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

// EntityScreeningHitDocumentsItems Analyzed documents for the associated hit
type EntityScreeningHitDocumentsItems struct {
	Analysis *MatchSummary `json:"analysis,omitempty"`
	Data *EntityDocument `json:"data,omitempty"`
}

// NewEntityScreeningHitDocumentsItems instantiates a new EntityScreeningHitDocumentsItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityScreeningHitDocumentsItems() *EntityScreeningHitDocumentsItems {
	this := EntityScreeningHitDocumentsItems{}
	return &this
}

// NewEntityScreeningHitDocumentsItemsWithDefaults instantiates a new EntityScreeningHitDocumentsItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityScreeningHitDocumentsItemsWithDefaults() *EntityScreeningHitDocumentsItems {
	this := EntityScreeningHitDocumentsItems{}
	return &this
}

// GetAnalysis returns the Analysis field value if set, zero value otherwise.
func (o *EntityScreeningHitDocumentsItems) GetAnalysis() MatchSummary {
	if o == nil || o.Analysis == nil {
		var ret MatchSummary
		return ret
	}
	return *o.Analysis
}

// GetAnalysisOk returns a tuple with the Analysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitDocumentsItems) GetAnalysisOk() (*MatchSummary, bool) {
	if o == nil || o.Analysis == nil {
		return nil, false
	}
	return o.Analysis, true
}

// HasAnalysis returns a boolean if a field has been set.
func (o *EntityScreeningHitDocumentsItems) HasAnalysis() bool {
	if o != nil && o.Analysis != nil {
		return true
	}

	return false
}

// SetAnalysis gets a reference to the given MatchSummary and assigns it to the Analysis field.
func (o *EntityScreeningHitDocumentsItems) SetAnalysis(v MatchSummary) {
	o.Analysis = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EntityScreeningHitDocumentsItems) GetData() EntityDocument {
	if o == nil || o.Data == nil {
		var ret EntityDocument
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitDocumentsItems) GetDataOk() (*EntityDocument, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EntityScreeningHitDocumentsItems) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given EntityDocument and assigns it to the Data field.
func (o *EntityScreeningHitDocumentsItems) SetData(v EntityDocument) {
	o.Data = &v
}

func (o EntityScreeningHitDocumentsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Analysis != nil {
		toSerialize["analysis"] = o.Analysis
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableEntityScreeningHitDocumentsItems struct {
	value *EntityScreeningHitDocumentsItems
	isSet bool
}

func (v NullableEntityScreeningHitDocumentsItems) Get() *EntityScreeningHitDocumentsItems {
	return v.value
}

func (v *NullableEntityScreeningHitDocumentsItems) Set(val *EntityScreeningHitDocumentsItems) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityScreeningHitDocumentsItems) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityScreeningHitDocumentsItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityScreeningHitDocumentsItems(val *EntityScreeningHitDocumentsItems) *NullableEntityScreeningHitDocumentsItems {
	return &NullableEntityScreeningHitDocumentsItems{value: val, isSet: true}
}

func (v NullableEntityScreeningHitDocumentsItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityScreeningHitDocumentsItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


