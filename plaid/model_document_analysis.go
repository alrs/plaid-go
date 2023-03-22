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

// DocumentAnalysis High level descriptions of how the associated document was processed. If a document fails verification, the details in the `analysis` object should help clarify why the document was rejected.
type DocumentAnalysis struct {
	Authenticity DocumentAuthenticityMatchCode `json:"authenticity"`
	ImageQuality ImageQuality `json:"image_quality"`
	ExtractedData NullablePhysicalDocumentExtractedDataAnalysis `json:"extracted_data"`
	AdditionalProperties map[string]interface{}
}

type _DocumentAnalysis DocumentAnalysis

// NewDocumentAnalysis instantiates a new DocumentAnalysis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentAnalysis(authenticity DocumentAuthenticityMatchCode, imageQuality ImageQuality, extractedData NullablePhysicalDocumentExtractedDataAnalysis) *DocumentAnalysis {
	this := DocumentAnalysis{}
	this.Authenticity = authenticity
	this.ImageQuality = imageQuality
	this.ExtractedData = extractedData
	return &this
}

// NewDocumentAnalysisWithDefaults instantiates a new DocumentAnalysis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentAnalysisWithDefaults() *DocumentAnalysis {
	this := DocumentAnalysis{}
	return &this
}

// GetAuthenticity returns the Authenticity field value
func (o *DocumentAnalysis) GetAuthenticity() DocumentAuthenticityMatchCode {
	if o == nil {
		var ret DocumentAuthenticityMatchCode
		return ret
	}

	return o.Authenticity
}

// GetAuthenticityOk returns a tuple with the Authenticity field value
// and a boolean to check if the value has been set.
func (o *DocumentAnalysis) GetAuthenticityOk() (*DocumentAuthenticityMatchCode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Authenticity, true
}

// SetAuthenticity sets field value
func (o *DocumentAnalysis) SetAuthenticity(v DocumentAuthenticityMatchCode) {
	o.Authenticity = v
}

// GetImageQuality returns the ImageQuality field value
func (o *DocumentAnalysis) GetImageQuality() ImageQuality {
	if o == nil {
		var ret ImageQuality
		return ret
	}

	return o.ImageQuality
}

// GetImageQualityOk returns a tuple with the ImageQuality field value
// and a boolean to check if the value has been set.
func (o *DocumentAnalysis) GetImageQualityOk() (*ImageQuality, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImageQuality, true
}

// SetImageQuality sets field value
func (o *DocumentAnalysis) SetImageQuality(v ImageQuality) {
	o.ImageQuality = v
}

// GetExtractedData returns the ExtractedData field value
// If the value is explicit nil, the zero value for PhysicalDocumentExtractedDataAnalysis will be returned
func (o *DocumentAnalysis) GetExtractedData() PhysicalDocumentExtractedDataAnalysis {
	if o == nil || o.ExtractedData.Get() == nil {
		var ret PhysicalDocumentExtractedDataAnalysis
		return ret
	}

	return *o.ExtractedData.Get()
}

// GetExtractedDataOk returns a tuple with the ExtractedData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentAnalysis) GetExtractedDataOk() (*PhysicalDocumentExtractedDataAnalysis, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExtractedData.Get(), o.ExtractedData.IsSet()
}

// SetExtractedData sets field value
func (o *DocumentAnalysis) SetExtractedData(v PhysicalDocumentExtractedDataAnalysis) {
	o.ExtractedData.Set(&v)
}

func (o DocumentAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["authenticity"] = o.Authenticity
	}
	if true {
		toSerialize["image_quality"] = o.ImageQuality
	}
	if true {
		toSerialize["extracted_data"] = o.ExtractedData.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DocumentAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	varDocumentAnalysis := _DocumentAnalysis{}

	if err = json.Unmarshal(bytes, &varDocumentAnalysis); err == nil {
		*o = DocumentAnalysis(varDocumentAnalysis)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "authenticity")
		delete(additionalProperties, "image_quality")
		delete(additionalProperties, "extracted_data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDocumentAnalysis struct {
	value *DocumentAnalysis
	isSet bool
}

func (v NullableDocumentAnalysis) Get() *DocumentAnalysis {
	return v.value
}

func (v *NullableDocumentAnalysis) Set(val *DocumentAnalysis) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentAnalysis) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentAnalysis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentAnalysis(val *DocumentAnalysis) *NullableDocumentAnalysis {
	return &NullableDocumentAnalysis{value: val, isSet: true}
}

func (v NullableDocumentAnalysis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentAnalysis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


