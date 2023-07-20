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

// TaxpayerIdentifier Information about the Taxpayer identification values assigned to the individual or legal entity.Information about the Taxpayer identification values assigned to the individual or legal entity.
type TaxpayerIdentifier struct {
	TaxpayerIdentifierType NullableTaxpayerIdentifierType `json:"TaxpayerIdentifierType"`
	// The value of the taxpayer identifier as assigned by the IRS to the individual or legal entity.
	TaxpayerIdentifierValue NullableString `json:"TaxpayerIdentifierValue"`
	AdditionalProperties map[string]interface{}
}

type _TaxpayerIdentifier TaxpayerIdentifier

// NewTaxpayerIdentifier instantiates a new TaxpayerIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxpayerIdentifier(taxpayerIdentifierType NullableTaxpayerIdentifierType, taxpayerIdentifierValue NullableString) *TaxpayerIdentifier {
	this := TaxpayerIdentifier{}
	this.TaxpayerIdentifierType = taxpayerIdentifierType
	this.TaxpayerIdentifierValue = taxpayerIdentifierValue
	return &this
}

// NewTaxpayerIdentifierWithDefaults instantiates a new TaxpayerIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxpayerIdentifierWithDefaults() *TaxpayerIdentifier {
	this := TaxpayerIdentifier{}
	return &this
}

// GetTaxpayerIdentifierType returns the TaxpayerIdentifierType field value
// If the value is explicit nil, the zero value for TaxpayerIdentifierType will be returned
func (o *TaxpayerIdentifier) GetTaxpayerIdentifierType() TaxpayerIdentifierType {
	if o == nil || o.TaxpayerIdentifierType.Get() == nil {
		var ret TaxpayerIdentifierType
		return ret
	}

	return *o.TaxpayerIdentifierType.Get()
}

// GetTaxpayerIdentifierTypeOk returns a tuple with the TaxpayerIdentifierType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxpayerIdentifier) GetTaxpayerIdentifierTypeOk() (*TaxpayerIdentifierType, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TaxpayerIdentifierType.Get(), o.TaxpayerIdentifierType.IsSet()
}

// SetTaxpayerIdentifierType sets field value
func (o *TaxpayerIdentifier) SetTaxpayerIdentifierType(v TaxpayerIdentifierType) {
	o.TaxpayerIdentifierType.Set(&v)
}

// GetTaxpayerIdentifierValue returns the TaxpayerIdentifierValue field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TaxpayerIdentifier) GetTaxpayerIdentifierValue() string {
	if o == nil || o.TaxpayerIdentifierValue.Get() == nil {
		var ret string
		return ret
	}

	return *o.TaxpayerIdentifierValue.Get()
}

// GetTaxpayerIdentifierValueOk returns a tuple with the TaxpayerIdentifierValue field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxpayerIdentifier) GetTaxpayerIdentifierValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TaxpayerIdentifierValue.Get(), o.TaxpayerIdentifierValue.IsSet()
}

// SetTaxpayerIdentifierValue sets field value
func (o *TaxpayerIdentifier) SetTaxpayerIdentifierValue(v string) {
	o.TaxpayerIdentifierValue.Set(&v)
}

func (o TaxpayerIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["TaxpayerIdentifierType"] = o.TaxpayerIdentifierType.Get()
	}
	if true {
		toSerialize["TaxpayerIdentifierValue"] = o.TaxpayerIdentifierValue.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TaxpayerIdentifier) UnmarshalJSON(bytes []byte) (err error) {
	varTaxpayerIdentifier := _TaxpayerIdentifier{}

	if err = json.Unmarshal(bytes, &varTaxpayerIdentifier); err == nil {
		*o = TaxpayerIdentifier(varTaxpayerIdentifier)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "TaxpayerIdentifierType")
		delete(additionalProperties, "TaxpayerIdentifierValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTaxpayerIdentifier struct {
	value *TaxpayerIdentifier
	isSet bool
}

func (v NullableTaxpayerIdentifier) Get() *TaxpayerIdentifier {
	return v.value
}

func (v *NullableTaxpayerIdentifier) Set(val *TaxpayerIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxpayerIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxpayerIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxpayerIdentifier(val *TaxpayerIdentifier) *NullableTaxpayerIdentifier {
	return &NullableTaxpayerIdentifier{value: val, isSet: true}
}

func (v NullableTaxpayerIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxpayerIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


