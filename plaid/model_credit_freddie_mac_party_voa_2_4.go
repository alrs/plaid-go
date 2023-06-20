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

// CreditFreddieMacPartyVOA24 A collection of information about a single party to a transaction. Included direct participants like the borrower and seller as well as indirect participants such as the flood certificate provider.
type CreditFreddieMacPartyVOA24 struct {
	INDIVIDUAL CreditFreddieMacPartyIndividualVOA24 `json:"INDIVIDUAL"`
	ROLES Roles `json:"ROLES"`
	TAXPAYER_IDENTIFIERS TaxpayerIdentifiers `json:"TAXPAYER_IDENTIFIERS"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacPartyVOA24 CreditFreddieMacPartyVOA24

// NewCreditFreddieMacPartyVOA24 instantiates a new CreditFreddieMacPartyVOA24 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacPartyVOA24(iNDIVIDUAL CreditFreddieMacPartyIndividualVOA24, rOLES Roles, tAXPAYERIDENTIFIERS TaxpayerIdentifiers) *CreditFreddieMacPartyVOA24 {
	this := CreditFreddieMacPartyVOA24{}
	this.INDIVIDUAL = iNDIVIDUAL
	this.ROLES = rOLES
	this.TAXPAYER_IDENTIFIERS = tAXPAYERIDENTIFIERS
	return &this
}

// NewCreditFreddieMacPartyVOA24WithDefaults instantiates a new CreditFreddieMacPartyVOA24 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacPartyVOA24WithDefaults() *CreditFreddieMacPartyVOA24 {
	this := CreditFreddieMacPartyVOA24{}
	return &this
}

// GetINDIVIDUAL returns the INDIVIDUAL field value
func (o *CreditFreddieMacPartyVOA24) GetINDIVIDUAL() CreditFreddieMacPartyIndividualVOA24 {
	if o == nil {
		var ret CreditFreddieMacPartyIndividualVOA24
		return ret
	}

	return o.INDIVIDUAL
}

// GetINDIVIDUALOk returns a tuple with the INDIVIDUAL field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacPartyVOA24) GetINDIVIDUALOk() (*CreditFreddieMacPartyIndividualVOA24, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.INDIVIDUAL, true
}

// SetINDIVIDUAL sets field value
func (o *CreditFreddieMacPartyVOA24) SetINDIVIDUAL(v CreditFreddieMacPartyIndividualVOA24) {
	o.INDIVIDUAL = v
}

// GetROLES returns the ROLES field value
func (o *CreditFreddieMacPartyVOA24) GetROLES() Roles {
	if o == nil {
		var ret Roles
		return ret
	}

	return o.ROLES
}

// GetROLESOk returns a tuple with the ROLES field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacPartyVOA24) GetROLESOk() (*Roles, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ROLES, true
}

// SetROLES sets field value
func (o *CreditFreddieMacPartyVOA24) SetROLES(v Roles) {
	o.ROLES = v
}

// GetTAXPAYER_IDENTIFIERS returns the TAXPAYER_IDENTIFIERS field value
func (o *CreditFreddieMacPartyVOA24) GetTAXPAYER_IDENTIFIERS() TaxpayerIdentifiers {
	if o == nil {
		var ret TaxpayerIdentifiers
		return ret
	}

	return o.TAXPAYER_IDENTIFIERS
}

// GetTAXPAYER_IDENTIFIERSOk returns a tuple with the TAXPAYER_IDENTIFIERS field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacPartyVOA24) GetTAXPAYER_IDENTIFIERSOk() (*TaxpayerIdentifiers, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TAXPAYER_IDENTIFIERS, true
}

// SetTAXPAYER_IDENTIFIERS sets field value
func (o *CreditFreddieMacPartyVOA24) SetTAXPAYER_IDENTIFIERS(v TaxpayerIdentifiers) {
	o.TAXPAYER_IDENTIFIERS = v
}

func (o CreditFreddieMacPartyVOA24) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["INDIVIDUAL"] = o.INDIVIDUAL
	}
	if true {
		toSerialize["ROLES"] = o.ROLES
	}
	if true {
		toSerialize["TAXPAYER_IDENTIFIERS"] = o.TAXPAYER_IDENTIFIERS
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacPartyVOA24) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacPartyVOA24 := _CreditFreddieMacPartyVOA24{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacPartyVOA24); err == nil {
		*o = CreditFreddieMacPartyVOA24(varCreditFreddieMacPartyVOA24)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "INDIVIDUAL")
		delete(additionalProperties, "ROLES")
		delete(additionalProperties, "TAXPAYER_IDENTIFIERS")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieMacPartyVOA24 struct {
	value *CreditFreddieMacPartyVOA24
	isSet bool
}

func (v NullableCreditFreddieMacPartyVOA24) Get() *CreditFreddieMacPartyVOA24 {
	return v.value
}

func (v *NullableCreditFreddieMacPartyVOA24) Set(val *CreditFreddieMacPartyVOA24) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacPartyVOA24) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacPartyVOA24) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacPartyVOA24(val *CreditFreddieMacPartyVOA24) *NullableCreditFreddieMacPartyVOA24 {
	return &NullableCreditFreddieMacPartyVOA24{value: val, isSet: true}
}

func (v NullableCreditFreddieMacPartyVOA24) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacPartyVOA24) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


