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

// SignalScores Risk scoring details broken down by risk category.
type SignalScores struct {
	CustomerInitiatedReturnRisk *CustomerInitiatedReturnRisk `json:"customer_initiated_return_risk,omitempty"`
	BankInitiatedReturnRisk *BankInitiatedReturnRisk `json:"bank_initiated_return_risk,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SignalScores SignalScores

// NewSignalScores instantiates a new SignalScores object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignalScores() *SignalScores {
	this := SignalScores{}
	return &this
}

// NewSignalScoresWithDefaults instantiates a new SignalScores object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignalScoresWithDefaults() *SignalScores {
	this := SignalScores{}
	return &this
}

// GetCustomerInitiatedReturnRisk returns the CustomerInitiatedReturnRisk field value if set, zero value otherwise.
func (o *SignalScores) GetCustomerInitiatedReturnRisk() CustomerInitiatedReturnRisk {
	if o == nil || o.CustomerInitiatedReturnRisk == nil {
		var ret CustomerInitiatedReturnRisk
		return ret
	}
	return *o.CustomerInitiatedReturnRisk
}

// GetCustomerInitiatedReturnRiskOk returns a tuple with the CustomerInitiatedReturnRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalScores) GetCustomerInitiatedReturnRiskOk() (*CustomerInitiatedReturnRisk, bool) {
	if o == nil || o.CustomerInitiatedReturnRisk == nil {
		return nil, false
	}
	return o.CustomerInitiatedReturnRisk, true
}

// HasCustomerInitiatedReturnRisk returns a boolean if a field has been set.
func (o *SignalScores) HasCustomerInitiatedReturnRisk() bool {
	if o != nil && o.CustomerInitiatedReturnRisk != nil {
		return true
	}

	return false
}

// SetCustomerInitiatedReturnRisk gets a reference to the given CustomerInitiatedReturnRisk and assigns it to the CustomerInitiatedReturnRisk field.
func (o *SignalScores) SetCustomerInitiatedReturnRisk(v CustomerInitiatedReturnRisk) {
	o.CustomerInitiatedReturnRisk = &v
}

// GetBankInitiatedReturnRisk returns the BankInitiatedReturnRisk field value if set, zero value otherwise.
func (o *SignalScores) GetBankInitiatedReturnRisk() BankInitiatedReturnRisk {
	if o == nil || o.BankInitiatedReturnRisk == nil {
		var ret BankInitiatedReturnRisk
		return ret
	}
	return *o.BankInitiatedReturnRisk
}

// GetBankInitiatedReturnRiskOk returns a tuple with the BankInitiatedReturnRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalScores) GetBankInitiatedReturnRiskOk() (*BankInitiatedReturnRisk, bool) {
	if o == nil || o.BankInitiatedReturnRisk == nil {
		return nil, false
	}
	return o.BankInitiatedReturnRisk, true
}

// HasBankInitiatedReturnRisk returns a boolean if a field has been set.
func (o *SignalScores) HasBankInitiatedReturnRisk() bool {
	if o != nil && o.BankInitiatedReturnRisk != nil {
		return true
	}

	return false
}

// SetBankInitiatedReturnRisk gets a reference to the given BankInitiatedReturnRisk and assigns it to the BankInitiatedReturnRisk field.
func (o *SignalScores) SetBankInitiatedReturnRisk(v BankInitiatedReturnRisk) {
	o.BankInitiatedReturnRisk = &v
}

func (o SignalScores) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomerInitiatedReturnRisk != nil {
		toSerialize["customer_initiated_return_risk"] = o.CustomerInitiatedReturnRisk
	}
	if o.BankInitiatedReturnRisk != nil {
		toSerialize["bank_initiated_return_risk"] = o.BankInitiatedReturnRisk
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SignalScores) UnmarshalJSON(bytes []byte) (err error) {
	varSignalScores := _SignalScores{}

	if err = json.Unmarshal(bytes, &varSignalScores); err == nil {
		*o = SignalScores(varSignalScores)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "customer_initiated_return_risk")
		delete(additionalProperties, "bank_initiated_return_risk")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSignalScores struct {
	value *SignalScores
	isSet bool
}

func (v NullableSignalScores) Get() *SignalScores {
	return v.value
}

func (v *NullableSignalScores) Set(val *SignalScores) {
	v.value = val
	v.isSet = true
}

func (v NullableSignalScores) IsSet() bool {
	return v.isSet
}

func (v *NullableSignalScores) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignalScores(val *SignalScores) *NullableSignalScores {
	return &NullableSignalScores{value: val, isSet: true}
}

func (v NullableSignalScores) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignalScores) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


