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

// LiabilitiesObject An object containing liability accounts
type LiabilitiesObject struct {
	// The credit accounts returned.
	Credit []CreditCardLiability `json:"credit"`
	// The mortgage accounts returned.
	Mortgage []MortgageLiability `json:"mortgage"`
	// The student loan accounts returned.
	Student []StudentLoan `json:"student"`
	AdditionalProperties map[string]interface{}
}

type _LiabilitiesObject LiabilitiesObject

// NewLiabilitiesObject instantiates a new LiabilitiesObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiabilitiesObject(credit []CreditCardLiability, mortgage []MortgageLiability, student []StudentLoan) *LiabilitiesObject {
	this := LiabilitiesObject{}
	this.Credit = credit
	this.Mortgage = mortgage
	this.Student = student
	return &this
}

// NewLiabilitiesObjectWithDefaults instantiates a new LiabilitiesObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiabilitiesObjectWithDefaults() *LiabilitiesObject {
	this := LiabilitiesObject{}
	return &this
}

// GetCredit returns the Credit field value
// If the value is explicit nil, the zero value for []CreditCardLiability will be returned
func (o *LiabilitiesObject) GetCredit() []CreditCardLiability {
	if o == nil {
		var ret []CreditCardLiability
		return ret
	}

	return o.Credit
}

// GetCreditOk returns a tuple with the Credit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiabilitiesObject) GetCreditOk() (*[]CreditCardLiability, bool) {
	if o == nil || o.Credit == nil {
		return nil, false
	}
	return &o.Credit, true
}

// SetCredit sets field value
func (o *LiabilitiesObject) SetCredit(v []CreditCardLiability) {
	o.Credit = v
}

// GetMortgage returns the Mortgage field value
// If the value is explicit nil, the zero value for []MortgageLiability will be returned
func (o *LiabilitiesObject) GetMortgage() []MortgageLiability {
	if o == nil {
		var ret []MortgageLiability
		return ret
	}

	return o.Mortgage
}

// GetMortgageOk returns a tuple with the Mortgage field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiabilitiesObject) GetMortgageOk() (*[]MortgageLiability, bool) {
	if o == nil || o.Mortgage == nil {
		return nil, false
	}
	return &o.Mortgage, true
}

// SetMortgage sets field value
func (o *LiabilitiesObject) SetMortgage(v []MortgageLiability) {
	o.Mortgage = v
}

// GetStudent returns the Student field value
// If the value is explicit nil, the zero value for []StudentLoan will be returned
func (o *LiabilitiesObject) GetStudent() []StudentLoan {
	if o == nil {
		var ret []StudentLoan
		return ret
	}

	return o.Student
}

// GetStudentOk returns a tuple with the Student field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiabilitiesObject) GetStudentOk() (*[]StudentLoan, bool) {
	if o == nil || o.Student == nil {
		return nil, false
	}
	return &o.Student, true
}

// SetStudent sets field value
func (o *LiabilitiesObject) SetStudent(v []StudentLoan) {
	o.Student = v
}

func (o LiabilitiesObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Credit != nil {
		toSerialize["credit"] = o.Credit
	}
	if o.Mortgage != nil {
		toSerialize["mortgage"] = o.Mortgage
	}
	if o.Student != nil {
		toSerialize["student"] = o.Student
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LiabilitiesObject) UnmarshalJSON(bytes []byte) (err error) {
	varLiabilitiesObject := _LiabilitiesObject{}

	if err = json.Unmarshal(bytes, &varLiabilitiesObject); err == nil {
		*o = LiabilitiesObject(varLiabilitiesObject)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "credit")
		delete(additionalProperties, "mortgage")
		delete(additionalProperties, "student")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLiabilitiesObject struct {
	value *LiabilitiesObject
	isSet bool
}

func (v NullableLiabilitiesObject) Get() *LiabilitiesObject {
	return v.value
}

func (v *NullableLiabilitiesObject) Set(val *LiabilitiesObject) {
	v.value = val
	v.isSet = true
}

func (v NullableLiabilitiesObject) IsSet() bool {
	return v.isSet
}

func (v *NullableLiabilitiesObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiabilitiesObject(val *LiabilitiesObject) *NullableLiabilitiesObject {
	return &NullableLiabilitiesObject{value: val, isSet: true}
}

func (v NullableLiabilitiesObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiabilitiesObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


