/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.146.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// WatchlistScreeningDocument An official document, usually issued by a governing body or institution, with an associated identifier.
type WatchlistScreeningDocument struct {
	Type WatchlistScreeningDocumentType `json:"type"`
	// The numeric or alphanumeric identifier associated with this document.
	Number string `json:"number"`
}

// NewWatchlistScreeningDocument instantiates a new WatchlistScreeningDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningDocument(type_ WatchlistScreeningDocumentType, number string) *WatchlistScreeningDocument {
	this := WatchlistScreeningDocument{}
	this.Type = type_
	this.Number = number
	return &this
}

// NewWatchlistScreeningDocumentWithDefaults instantiates a new WatchlistScreeningDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningDocumentWithDefaults() *WatchlistScreeningDocument {
	this := WatchlistScreeningDocument{}
	return &this
}

// GetType returns the Type field value
func (o *WatchlistScreeningDocument) GetType() WatchlistScreeningDocumentType {
	if o == nil {
		var ret WatchlistScreeningDocumentType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningDocument) GetTypeOk() (*WatchlistScreeningDocumentType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WatchlistScreeningDocument) SetType(v WatchlistScreeningDocumentType) {
	o.Type = v
}

// GetNumber returns the Number field value
func (o *WatchlistScreeningDocument) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningDocument) GetNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *WatchlistScreeningDocument) SetNumber(v string) {
	o.Number = v
}

func (o WatchlistScreeningDocument) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["number"] = o.Number
	}
	return json.Marshal(toSerialize)
}

type NullableWatchlistScreeningDocument struct {
	value *WatchlistScreeningDocument
	isSet bool
}

func (v NullableWatchlistScreeningDocument) Get() *WatchlistScreeningDocument {
	return v.value
}

func (v *NullableWatchlistScreeningDocument) Set(val *WatchlistScreeningDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningDocument(val *WatchlistScreeningDocument) *NullableWatchlistScreeningDocument {
	return &NullableWatchlistScreeningDocument{value: val, isSet: true}
}

func (v NullableWatchlistScreeningDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

