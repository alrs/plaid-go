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

// CreditFreddieMacVerificationOfAssetsDealVOA24 An object representing an Asset Report with Freddie Mac schema.
type CreditFreddieMacVerificationOfAssetsDealVOA24 struct {
	LOANS CreditFreddieMacLoansVOA24 `json:"LOANS"`
	PARTIES CreditFreddieMacPartiesVOA24 `json:"PARTIES"`
	SERVICES CreditFreddieMacServicesVOA24 `json:"SERVICES"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacVerificationOfAssetsDealVOA24 CreditFreddieMacVerificationOfAssetsDealVOA24

// NewCreditFreddieMacVerificationOfAssetsDealVOA24 instantiates a new CreditFreddieMacVerificationOfAssetsDealVOA24 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacVerificationOfAssetsDealVOA24(lOANS CreditFreddieMacLoansVOA24, pARTIES CreditFreddieMacPartiesVOA24, sERVICES CreditFreddieMacServicesVOA24) *CreditFreddieMacVerificationOfAssetsDealVOA24 {
	this := CreditFreddieMacVerificationOfAssetsDealVOA24{}
	this.LOANS = lOANS
	this.PARTIES = pARTIES
	this.SERVICES = sERVICES
	return &this
}

// NewCreditFreddieMacVerificationOfAssetsDealVOA24WithDefaults instantiates a new CreditFreddieMacVerificationOfAssetsDealVOA24 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacVerificationOfAssetsDealVOA24WithDefaults() *CreditFreddieMacVerificationOfAssetsDealVOA24 {
	this := CreditFreddieMacVerificationOfAssetsDealVOA24{}
	return &this
}

// GetLOANS returns the LOANS field value
func (o *CreditFreddieMacVerificationOfAssetsDealVOA24) GetLOANS() CreditFreddieMacLoansVOA24 {
	if o == nil {
		var ret CreditFreddieMacLoansVOA24
		return ret
	}

	return o.LOANS
}

// GetLOANSOk returns a tuple with the LOANS field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacVerificationOfAssetsDealVOA24) GetLOANSOk() (*CreditFreddieMacLoansVOA24, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LOANS, true
}

// SetLOANS sets field value
func (o *CreditFreddieMacVerificationOfAssetsDealVOA24) SetLOANS(v CreditFreddieMacLoansVOA24) {
	o.LOANS = v
}

// GetPARTIES returns the PARTIES field value
func (o *CreditFreddieMacVerificationOfAssetsDealVOA24) GetPARTIES() CreditFreddieMacPartiesVOA24 {
	if o == nil {
		var ret CreditFreddieMacPartiesVOA24
		return ret
	}

	return o.PARTIES
}

// GetPARTIESOk returns a tuple with the PARTIES field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacVerificationOfAssetsDealVOA24) GetPARTIESOk() (*CreditFreddieMacPartiesVOA24, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PARTIES, true
}

// SetPARTIES sets field value
func (o *CreditFreddieMacVerificationOfAssetsDealVOA24) SetPARTIES(v CreditFreddieMacPartiesVOA24) {
	o.PARTIES = v
}

// GetSERVICES returns the SERVICES field value
func (o *CreditFreddieMacVerificationOfAssetsDealVOA24) GetSERVICES() CreditFreddieMacServicesVOA24 {
	if o == nil {
		var ret CreditFreddieMacServicesVOA24
		return ret
	}

	return o.SERVICES
}

// GetSERVICESOk returns a tuple with the SERVICES field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacVerificationOfAssetsDealVOA24) GetSERVICESOk() (*CreditFreddieMacServicesVOA24, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SERVICES, true
}

// SetSERVICES sets field value
func (o *CreditFreddieMacVerificationOfAssetsDealVOA24) SetSERVICES(v CreditFreddieMacServicesVOA24) {
	o.SERVICES = v
}

func (o CreditFreddieMacVerificationOfAssetsDealVOA24) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["LOANS"] = o.LOANS
	}
	if true {
		toSerialize["PARTIES"] = o.PARTIES
	}
	if true {
		toSerialize["SERVICES"] = o.SERVICES
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacVerificationOfAssetsDealVOA24) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacVerificationOfAssetsDealVOA24 := _CreditFreddieMacVerificationOfAssetsDealVOA24{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacVerificationOfAssetsDealVOA24); err == nil {
		*o = CreditFreddieMacVerificationOfAssetsDealVOA24(varCreditFreddieMacVerificationOfAssetsDealVOA24)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "LOANS")
		delete(additionalProperties, "PARTIES")
		delete(additionalProperties, "SERVICES")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieMacVerificationOfAssetsDealVOA24 struct {
	value *CreditFreddieMacVerificationOfAssetsDealVOA24
	isSet bool
}

func (v NullableCreditFreddieMacVerificationOfAssetsDealVOA24) Get() *CreditFreddieMacVerificationOfAssetsDealVOA24 {
	return v.value
}

func (v *NullableCreditFreddieMacVerificationOfAssetsDealVOA24) Set(val *CreditFreddieMacVerificationOfAssetsDealVOA24) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacVerificationOfAssetsDealVOA24) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacVerificationOfAssetsDealVOA24) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacVerificationOfAssetsDealVOA24(val *CreditFreddieMacVerificationOfAssetsDealVOA24) *NullableCreditFreddieMacVerificationOfAssetsDealVOA24 {
	return &NullableCreditFreddieMacVerificationOfAssetsDealVOA24{value: val, isSet: true}
}

func (v NullableCreditFreddieMacVerificationOfAssetsDealVOA24) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacVerificationOfAssetsDealVOA24) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


