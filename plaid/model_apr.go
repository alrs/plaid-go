/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.64.14
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// APR Information about the APR on the account.
type APR struct {
	// Annual Percentage Rate applied. 
	AprPercentage float32 `json:"apr_percentage"`
	// The type of balance to which the APR applies.
	AprType string `json:"apr_type"`
	// Amount of money that is subjected to the APR if a balance was carried beyond payment due date. How it is calculated can vary by card issuer. It is often calculated as an average daily balance.
	BalanceSubjectToApr NullableFloat32 `json:"balance_subject_to_apr"`
	// Amount of money charged due to interest from last statement.
	InterestChargeAmount NullableFloat32 `json:"interest_charge_amount"`
	AdditionalProperties map[string]interface{}
}

type _APR APR

// NewAPR instantiates a new APR object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPR(aprPercentage float32, aprType string, balanceSubjectToApr NullableFloat32, interestChargeAmount NullableFloat32) *APR {
	this := APR{}
	this.AprPercentage = aprPercentage
	this.AprType = aprType
	this.BalanceSubjectToApr = balanceSubjectToApr
	this.InterestChargeAmount = interestChargeAmount
	return &this
}

// NewAPRWithDefaults instantiates a new APR object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPRWithDefaults() *APR {
	this := APR{}
	return &this
}

// GetAprPercentage returns the AprPercentage field value
func (o *APR) GetAprPercentage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AprPercentage
}

// GetAprPercentageOk returns a tuple with the AprPercentage field value
// and a boolean to check if the value has been set.
func (o *APR) GetAprPercentageOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AprPercentage, true
}

// SetAprPercentage sets field value
func (o *APR) SetAprPercentage(v float32) {
	o.AprPercentage = v
}

// GetAprType returns the AprType field value
func (o *APR) GetAprType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AprType
}

// GetAprTypeOk returns a tuple with the AprType field value
// and a boolean to check if the value has been set.
func (o *APR) GetAprTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AprType, true
}

// SetAprType sets field value
func (o *APR) SetAprType(v string) {
	o.AprType = v
}

// GetBalanceSubjectToApr returns the BalanceSubjectToApr field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *APR) GetBalanceSubjectToApr() float32 {
	if o == nil || o.BalanceSubjectToApr.Get() == nil {
		var ret float32
		return ret
	}

	return *o.BalanceSubjectToApr.Get()
}

// GetBalanceSubjectToAprOk returns a tuple with the BalanceSubjectToApr field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *APR) GetBalanceSubjectToAprOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BalanceSubjectToApr.Get(), o.BalanceSubjectToApr.IsSet()
}

// SetBalanceSubjectToApr sets field value
func (o *APR) SetBalanceSubjectToApr(v float32) {
	o.BalanceSubjectToApr.Set(&v)
}

// GetInterestChargeAmount returns the InterestChargeAmount field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *APR) GetInterestChargeAmount() float32 {
	if o == nil || o.InterestChargeAmount.Get() == nil {
		var ret float32
		return ret
	}

	return *o.InterestChargeAmount.Get()
}

// GetInterestChargeAmountOk returns a tuple with the InterestChargeAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *APR) GetInterestChargeAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InterestChargeAmount.Get(), o.InterestChargeAmount.IsSet()
}

// SetInterestChargeAmount sets field value
func (o *APR) SetInterestChargeAmount(v float32) {
	o.InterestChargeAmount.Set(&v)
}

func (o APR) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["apr_percentage"] = o.AprPercentage
	}
	if true {
		toSerialize["apr_type"] = o.AprType
	}
	if true {
		toSerialize["balance_subject_to_apr"] = o.BalanceSubjectToApr.Get()
	}
	if true {
		toSerialize["interest_charge_amount"] = o.InterestChargeAmount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *APR) UnmarshalJSON(bytes []byte) (err error) {
	varAPR := _APR{}

	if err = json.Unmarshal(bytes, &varAPR); err == nil {
		*o = APR(varAPR)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "apr_percentage")
		delete(additionalProperties, "apr_type")
		delete(additionalProperties, "balance_subject_to_apr")
		delete(additionalProperties, "interest_charge_amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAPR struct {
	value *APR
	isSet bool
}

func (v NullableAPR) Get() *APR {
	return v.value
}

func (v *NullableAPR) Set(val *APR) {
	v.value = val
	v.isSet = true
}

func (v NullableAPR) IsSet() bool {
	return v.isSet
}

func (v *NullableAPR) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPR(val *APR) *NullableAPR {
	return &NullableAPR{value: val, isSet: true}
}

func (v NullableAPR) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPR) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


