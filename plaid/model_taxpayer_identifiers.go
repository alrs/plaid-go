/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.205.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TaxpayerIdentifiers The collection of TAXPAYER_IDENTIFICATION elements
type TaxpayerIdentifiers struct {
	TAXPAYER_IDENTIFIER TaxpayerIdentifier `json:"TAXPAYER_IDENTIFIER"`
	AdditionalProperties map[string]interface{}
}

type _TaxpayerIdentifiers TaxpayerIdentifiers

// NewTaxpayerIdentifiers instantiates a new TaxpayerIdentifiers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxpayerIdentifiers(tAXPAYERIDENTIFIER TaxpayerIdentifier) *TaxpayerIdentifiers {
	this := TaxpayerIdentifiers{}
	this.TAXPAYER_IDENTIFIER = tAXPAYERIDENTIFIER
	return &this
}

// NewTaxpayerIdentifiersWithDefaults instantiates a new TaxpayerIdentifiers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxpayerIdentifiersWithDefaults() *TaxpayerIdentifiers {
	this := TaxpayerIdentifiers{}
	return &this
}

// GetTAXPAYER_IDENTIFIER returns the TAXPAYER_IDENTIFIER field value
func (o *TaxpayerIdentifiers) GetTAXPAYER_IDENTIFIER() TaxpayerIdentifier {
	if o == nil {
		var ret TaxpayerIdentifier
		return ret
	}

	return o.TAXPAYER_IDENTIFIER
}

// GetTAXPAYER_IDENTIFIEROk returns a tuple with the TAXPAYER_IDENTIFIER field value
// and a boolean to check if the value has been set.
func (o *TaxpayerIdentifiers) GetTAXPAYER_IDENTIFIEROk() (*TaxpayerIdentifier, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TAXPAYER_IDENTIFIER, true
}

// SetTAXPAYER_IDENTIFIER sets field value
func (o *TaxpayerIdentifiers) SetTAXPAYER_IDENTIFIER(v TaxpayerIdentifier) {
	o.TAXPAYER_IDENTIFIER = v
}

func (o TaxpayerIdentifiers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["TAXPAYER_IDENTIFIER"] = o.TAXPAYER_IDENTIFIER
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TaxpayerIdentifiers) UnmarshalJSON(bytes []byte) (err error) {
	varTaxpayerIdentifiers := _TaxpayerIdentifiers{}

	if err = json.Unmarshal(bytes, &varTaxpayerIdentifiers); err == nil {
		*o = TaxpayerIdentifiers(varTaxpayerIdentifiers)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "TAXPAYER_IDENTIFIER")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTaxpayerIdentifiers struct {
	value *TaxpayerIdentifiers
	isSet bool
}

func (v NullableTaxpayerIdentifiers) Get() *TaxpayerIdentifiers {
	return v.value
}

func (v *NullableTaxpayerIdentifiers) Set(val *TaxpayerIdentifiers) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxpayerIdentifiers) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxpayerIdentifiers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxpayerIdentifiers(val *TaxpayerIdentifiers) *NullableTaxpayerIdentifiers {
	return &NullableTaxpayerIdentifiers{value: val, isSet: true}
}

func (v NullableTaxpayerIdentifiers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxpayerIdentifiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


