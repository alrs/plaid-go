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

// SelfieAnalysis High level descriptions of how the associated selfie was processed. If a selfie fails verification, the details in the `analysis` object should help clarify why the selfie was rejected.
type SelfieAnalysis struct {
	DocumentComparison SelfieAnalysisDocumentComparison `json:"document_comparison"`
	AdditionalProperties map[string]interface{}
}

type _SelfieAnalysis SelfieAnalysis

// NewSelfieAnalysis instantiates a new SelfieAnalysis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfieAnalysis(documentComparison SelfieAnalysisDocumentComparison) *SelfieAnalysis {
	this := SelfieAnalysis{}
	this.DocumentComparison = documentComparison
	return &this
}

// NewSelfieAnalysisWithDefaults instantiates a new SelfieAnalysis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfieAnalysisWithDefaults() *SelfieAnalysis {
	this := SelfieAnalysis{}
	return &this
}

// GetDocumentComparison returns the DocumentComparison field value
func (o *SelfieAnalysis) GetDocumentComparison() SelfieAnalysisDocumentComparison {
	if o == nil {
		var ret SelfieAnalysisDocumentComparison
		return ret
	}

	return o.DocumentComparison
}

// GetDocumentComparisonOk returns a tuple with the DocumentComparison field value
// and a boolean to check if the value has been set.
func (o *SelfieAnalysis) GetDocumentComparisonOk() (*SelfieAnalysisDocumentComparison, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocumentComparison, true
}

// SetDocumentComparison sets field value
func (o *SelfieAnalysis) SetDocumentComparison(v SelfieAnalysisDocumentComparison) {
	o.DocumentComparison = v
}

func (o SelfieAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["document_comparison"] = o.DocumentComparison
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SelfieAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	varSelfieAnalysis := _SelfieAnalysis{}

	if err = json.Unmarshal(bytes, &varSelfieAnalysis); err == nil {
		*o = SelfieAnalysis(varSelfieAnalysis)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "document_comparison")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSelfieAnalysis struct {
	value *SelfieAnalysis
	isSet bool
}

func (v NullableSelfieAnalysis) Get() *SelfieAnalysis {
	return v.value
}

func (v *NullableSelfieAnalysis) Set(val *SelfieAnalysis) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfieAnalysis) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfieAnalysis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfieAnalysis(val *SelfieAnalysis) *NullableSelfieAnalysis {
	return &NullableSelfieAnalysis{value: val, isSet: true}
}

func (v NullableSelfieAnalysis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfieAnalysis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

