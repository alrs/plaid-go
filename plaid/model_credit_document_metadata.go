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

// CreditDocumentMetadata Object representing metadata pertaining to the document.
type CreditDocumentMetadata struct {
	// The name of the document.
	Name string `json:"name"`
	// The type of document.  `PAYSTUB`: A paystub.  `BANK_STATEMENT`: A bank statement.  `US_TAX_W2`: A W-2 wage and tax statement provided by a US employer reflecting wages earned by the employee.  `US_MILITARY_ERAS`: An electronic Retirement Account Statement (eRAS) issued by the US military.  `US_MILITARY_LES`: A Leave and Earnings Statement (LES) issued by the US military.  `US_MILITARY_CLES`: A Civilian Leave and Earnings Statment (CLES) issued by the US military.  `GIG`: Used to indicate that the income is related to gig work. Does not necessarily correspond to a specific document type.  `NONE`: Used to indicate that there is no underlying document for the data.  `UNKNOWN`: Document type could not be determined.
	DocumentType NullableString `json:"document_type"`
	// Signed URL to retrieve the underlying file.
	DownloadUrl NullableString `json:"download_url"`
	// The processing status of the document.
	Status NullableString `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _CreditDocumentMetadata CreditDocumentMetadata

// NewCreditDocumentMetadata instantiates a new CreditDocumentMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditDocumentMetadata(name string, documentType NullableString, downloadUrl NullableString, status NullableString) *CreditDocumentMetadata {
	this := CreditDocumentMetadata{}
	this.Name = name
	this.DocumentType = documentType
	this.DownloadUrl = downloadUrl
	this.Status = status
	return &this
}

// NewCreditDocumentMetadataWithDefaults instantiates a new CreditDocumentMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditDocumentMetadataWithDefaults() *CreditDocumentMetadata {
	this := CreditDocumentMetadata{}
	return &this
}

// GetName returns the Name field value
func (o *CreditDocumentMetadata) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreditDocumentMetadata) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreditDocumentMetadata) SetName(v string) {
	o.Name = v
}

// GetDocumentType returns the DocumentType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditDocumentMetadata) GetDocumentType() string {
	if o == nil || o.DocumentType.Get() == nil {
		var ret string
		return ret
	}

	return *o.DocumentType.Get()
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditDocumentMetadata) GetDocumentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DocumentType.Get(), o.DocumentType.IsSet()
}

// SetDocumentType sets field value
func (o *CreditDocumentMetadata) SetDocumentType(v string) {
	o.DocumentType.Set(&v)
}

// GetDownloadUrl returns the DownloadUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditDocumentMetadata) GetDownloadUrl() string {
	if o == nil || o.DownloadUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.DownloadUrl.Get()
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditDocumentMetadata) GetDownloadUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DownloadUrl.Get(), o.DownloadUrl.IsSet()
}

// SetDownloadUrl sets field value
func (o *CreditDocumentMetadata) SetDownloadUrl(v string) {
	o.DownloadUrl.Set(&v)
}

// GetStatus returns the Status field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditDocumentMetadata) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}

	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditDocumentMetadata) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// SetStatus sets field value
func (o *CreditDocumentMetadata) SetStatus(v string) {
	o.Status.Set(&v)
}

func (o CreditDocumentMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["document_type"] = o.DocumentType.Get()
	}
	if true {
		toSerialize["download_url"] = o.DownloadUrl.Get()
	}
	if true {
		toSerialize["status"] = o.Status.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditDocumentMetadata) UnmarshalJSON(bytes []byte) (err error) {
	varCreditDocumentMetadata := _CreditDocumentMetadata{}

	if err = json.Unmarshal(bytes, &varCreditDocumentMetadata); err == nil {
		*o = CreditDocumentMetadata(varCreditDocumentMetadata)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "document_type")
		delete(additionalProperties, "download_url")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditDocumentMetadata struct {
	value *CreditDocumentMetadata
	isSet bool
}

func (v NullableCreditDocumentMetadata) Get() *CreditDocumentMetadata {
	return v.value
}

func (v *NullableCreditDocumentMetadata) Set(val *CreditDocumentMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditDocumentMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditDocumentMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditDocumentMetadata(val *CreditDocumentMetadata) *NullableCreditDocumentMetadata {
	return &NullableCreditDocumentMetadata{value: val, isSet: true}
}

func (v NullableCreditDocumentMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditDocumentMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


