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

// LoanIdentifiers Collection of current and previous identifiers for this loan.
type LoanIdentifiers struct {
	LOAN_IDENTIFIER LoanIdentifier `json:"LOAN_IDENTIFIER"`
	AdditionalProperties map[string]interface{}
}

type _LoanIdentifiers LoanIdentifiers

// NewLoanIdentifiers instantiates a new LoanIdentifiers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoanIdentifiers(lOANIDENTIFIER LoanIdentifier) *LoanIdentifiers {
	this := LoanIdentifiers{}
	this.LOAN_IDENTIFIER = lOANIDENTIFIER
	return &this
}

// NewLoanIdentifiersWithDefaults instantiates a new LoanIdentifiers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoanIdentifiersWithDefaults() *LoanIdentifiers {
	this := LoanIdentifiers{}
	return &this
}

// GetLOAN_IDENTIFIER returns the LOAN_IDENTIFIER field value
func (o *LoanIdentifiers) GetLOAN_IDENTIFIER() LoanIdentifier {
	if o == nil {
		var ret LoanIdentifier
		return ret
	}

	return o.LOAN_IDENTIFIER
}

// GetLOAN_IDENTIFIEROk returns a tuple with the LOAN_IDENTIFIER field value
// and a boolean to check if the value has been set.
func (o *LoanIdentifiers) GetLOAN_IDENTIFIEROk() (*LoanIdentifier, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LOAN_IDENTIFIER, true
}

// SetLOAN_IDENTIFIER sets field value
func (o *LoanIdentifiers) SetLOAN_IDENTIFIER(v LoanIdentifier) {
	o.LOAN_IDENTIFIER = v
}

func (o LoanIdentifiers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["LOAN_IDENTIFIER"] = o.LOAN_IDENTIFIER
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LoanIdentifiers) UnmarshalJSON(bytes []byte) (err error) {
	varLoanIdentifiers := _LoanIdentifiers{}

	if err = json.Unmarshal(bytes, &varLoanIdentifiers); err == nil {
		*o = LoanIdentifiers(varLoanIdentifiers)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "LOAN_IDENTIFIER")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoanIdentifiers struct {
	value *LoanIdentifiers
	isSet bool
}

func (v NullableLoanIdentifiers) Get() *LoanIdentifiers {
	return v.value
}

func (v *NullableLoanIdentifiers) Set(val *LoanIdentifiers) {
	v.value = val
	v.isSet = true
}

func (v NullableLoanIdentifiers) IsSet() bool {
	return v.isSet
}

func (v *NullableLoanIdentifiers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoanIdentifiers(val *LoanIdentifiers) *NullableLoanIdentifiers {
	return &NullableLoanIdentifiers{value: val, isSet: true}
}

func (v NullableLoanIdentifiers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoanIdentifiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


