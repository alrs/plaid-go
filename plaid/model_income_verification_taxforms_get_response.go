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

// IncomeVerificationTaxformsGetResponse IncomeVerificationTaxformsGetResponse defines the response schema for `/income/verification/taxforms/get`
type IncomeVerificationTaxformsGetResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId *string `json:"request_id,omitempty"`
	DocumentMetadata []DocumentMetadata `json:"document_metadata"`
	// A list of forms.
	Taxforms []Taxform `json:"taxforms"`
	Error NullablePlaidError `json:"error,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IncomeVerificationTaxformsGetResponse IncomeVerificationTaxformsGetResponse

// NewIncomeVerificationTaxformsGetResponse instantiates a new IncomeVerificationTaxformsGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeVerificationTaxformsGetResponse(documentMetadata []DocumentMetadata, taxforms []Taxform) *IncomeVerificationTaxformsGetResponse {
	this := IncomeVerificationTaxformsGetResponse{}
	this.DocumentMetadata = documentMetadata
	this.Taxforms = taxforms
	return &this
}

// NewIncomeVerificationTaxformsGetResponseWithDefaults instantiates a new IncomeVerificationTaxformsGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeVerificationTaxformsGetResponseWithDefaults() *IncomeVerificationTaxformsGetResponse {
	this := IncomeVerificationTaxformsGetResponse{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *IncomeVerificationTaxformsGetResponse) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationTaxformsGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *IncomeVerificationTaxformsGetResponse) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *IncomeVerificationTaxformsGetResponse) SetRequestId(v string) {
	o.RequestId = &v
}

// GetDocumentMetadata returns the DocumentMetadata field value
func (o *IncomeVerificationTaxformsGetResponse) GetDocumentMetadata() []DocumentMetadata {
	if o == nil {
		var ret []DocumentMetadata
		return ret
	}

	return o.DocumentMetadata
}

// GetDocumentMetadataOk returns a tuple with the DocumentMetadata field value
// and a boolean to check if the value has been set.
func (o *IncomeVerificationTaxformsGetResponse) GetDocumentMetadataOk() (*[]DocumentMetadata, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocumentMetadata, true
}

// SetDocumentMetadata sets field value
func (o *IncomeVerificationTaxformsGetResponse) SetDocumentMetadata(v []DocumentMetadata) {
	o.DocumentMetadata = v
}

// GetTaxforms returns the Taxforms field value
func (o *IncomeVerificationTaxformsGetResponse) GetTaxforms() []Taxform {
	if o == nil {
		var ret []Taxform
		return ret
	}

	return o.Taxforms
}

// GetTaxformsOk returns a tuple with the Taxforms field value
// and a boolean to check if the value has been set.
func (o *IncomeVerificationTaxformsGetResponse) GetTaxformsOk() (*[]Taxform, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Taxforms, true
}

// SetTaxforms sets field value
func (o *IncomeVerificationTaxformsGetResponse) SetTaxforms(v []Taxform) {
	o.Taxforms = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeVerificationTaxformsGetResponse) GetError() PlaidError {
	if o == nil || o.Error.Get() == nil {
		var ret PlaidError
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeVerificationTaxformsGetResponse) GetErrorOk() (*PlaidError, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *IncomeVerificationTaxformsGetResponse) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullablePlaidError and assigns it to the Error field.
func (o *IncomeVerificationTaxformsGetResponse) SetError(v PlaidError) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *IncomeVerificationTaxformsGetResponse) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *IncomeVerificationTaxformsGetResponse) UnsetError() {
	o.Error.Unset()
}

func (o IncomeVerificationTaxformsGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}
	if true {
		toSerialize["document_metadata"] = o.DocumentMetadata
	}
	if true {
		toSerialize["taxforms"] = o.Taxforms
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IncomeVerificationTaxformsGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varIncomeVerificationTaxformsGetResponse := _IncomeVerificationTaxformsGetResponse{}

	if err = json.Unmarshal(bytes, &varIncomeVerificationTaxformsGetResponse); err == nil {
		*o = IncomeVerificationTaxformsGetResponse(varIncomeVerificationTaxformsGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		delete(additionalProperties, "document_metadata")
		delete(additionalProperties, "taxforms")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIncomeVerificationTaxformsGetResponse struct {
	value *IncomeVerificationTaxformsGetResponse
	isSet bool
}

func (v NullableIncomeVerificationTaxformsGetResponse) Get() *IncomeVerificationTaxformsGetResponse {
	return v.value
}

func (v *NullableIncomeVerificationTaxformsGetResponse) Set(val *IncomeVerificationTaxformsGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationTaxformsGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationTaxformsGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationTaxformsGetResponse(val *IncomeVerificationTaxformsGetResponse) *NullableIncomeVerificationTaxformsGetResponse {
	return &NullableIncomeVerificationTaxformsGetResponse{value: val, isSet: true}
}

func (v NullableIncomeVerificationTaxformsGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationTaxformsGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


