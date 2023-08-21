/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.419.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// EntityScreeningHitNamesItems Analyzed names for the associated hit
type EntityScreeningHitNamesItems struct {
	Analysis *MatchSummary `json:"analysis,omitempty"`
	Data *EntityScreeningHitNames `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntityScreeningHitNamesItems EntityScreeningHitNamesItems

// NewEntityScreeningHitNamesItems instantiates a new EntityScreeningHitNamesItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityScreeningHitNamesItems() *EntityScreeningHitNamesItems {
	this := EntityScreeningHitNamesItems{}
	return &this
}

// NewEntityScreeningHitNamesItemsWithDefaults instantiates a new EntityScreeningHitNamesItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityScreeningHitNamesItemsWithDefaults() *EntityScreeningHitNamesItems {
	this := EntityScreeningHitNamesItems{}
	return &this
}

// GetAnalysis returns the Analysis field value if set, zero value otherwise.
func (o *EntityScreeningHitNamesItems) GetAnalysis() MatchSummary {
	if o == nil || o.Analysis == nil {
		var ret MatchSummary
		return ret
	}
	return *o.Analysis
}

// GetAnalysisOk returns a tuple with the Analysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitNamesItems) GetAnalysisOk() (*MatchSummary, bool) {
	if o == nil || o.Analysis == nil {
		return nil, false
	}
	return o.Analysis, true
}

// HasAnalysis returns a boolean if a field has been set.
func (o *EntityScreeningHitNamesItems) HasAnalysis() bool {
	if o != nil && o.Analysis != nil {
		return true
	}

	return false
}

// SetAnalysis gets a reference to the given MatchSummary and assigns it to the Analysis field.
func (o *EntityScreeningHitNamesItems) SetAnalysis(v MatchSummary) {
	o.Analysis = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EntityScreeningHitNamesItems) GetData() EntityScreeningHitNames {
	if o == nil || o.Data == nil {
		var ret EntityScreeningHitNames
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitNamesItems) GetDataOk() (*EntityScreeningHitNames, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EntityScreeningHitNamesItems) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given EntityScreeningHitNames and assigns it to the Data field.
func (o *EntityScreeningHitNamesItems) SetData(v EntityScreeningHitNames) {
	o.Data = &v
}

func (o EntityScreeningHitNamesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Analysis != nil {
		toSerialize["analysis"] = o.Analysis
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntityScreeningHitNamesItems) UnmarshalJSON(bytes []byte) (err error) {
	varEntityScreeningHitNamesItems := _EntityScreeningHitNamesItems{}

	if err = json.Unmarshal(bytes, &varEntityScreeningHitNamesItems); err == nil {
		*o = EntityScreeningHitNamesItems(varEntityScreeningHitNamesItems)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "analysis")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntityScreeningHitNamesItems struct {
	value *EntityScreeningHitNamesItems
	isSet bool
}

func (v NullableEntityScreeningHitNamesItems) Get() *EntityScreeningHitNamesItems {
	return v.value
}

func (v *NullableEntityScreeningHitNamesItems) Set(val *EntityScreeningHitNamesItems) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityScreeningHitNamesItems) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityScreeningHitNamesItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityScreeningHitNamesItems(val *EntityScreeningHitNamesItems) *NullableEntityScreeningHitNamesItems {
	return &NullableEntityScreeningHitNamesItems{value: val, isSet: true}
}

func (v NullableEntityScreeningHitNamesItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityScreeningHitNamesItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


