/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.79.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// PaymentInitiationConsentConstraints Limitations that will be applied to payments initiated using the payment consent.
type PaymentInitiationConsentConstraints struct {
	ValidDateTime *PaymentConsentValidDateTime `json:"valid_date_time,omitempty"`
	MaxPaymentAmount *PaymentConsentMaxPaymentAmount `json:"max_payment_amount,omitempty"`
	// A list of amount limitations per period of time.
	PeriodicAmounts *[]PaymentConsentPeriodicAmount `json:"periodic_amounts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaymentInitiationConsentConstraints PaymentInitiationConsentConstraints

// NewPaymentInitiationConsentConstraints instantiates a new PaymentInitiationConsentConstraints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationConsentConstraints() *PaymentInitiationConsentConstraints {
	this := PaymentInitiationConsentConstraints{}
	return &this
}

// NewPaymentInitiationConsentConstraintsWithDefaults instantiates a new PaymentInitiationConsentConstraints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationConsentConstraintsWithDefaults() *PaymentInitiationConsentConstraints {
	this := PaymentInitiationConsentConstraints{}
	return &this
}

// GetValidDateTime returns the ValidDateTime field value if set, zero value otherwise.
func (o *PaymentInitiationConsentConstraints) GetValidDateTime() PaymentConsentValidDateTime {
	if o == nil || o.ValidDateTime == nil {
		var ret PaymentConsentValidDateTime
		return ret
	}
	return *o.ValidDateTime
}

// GetValidDateTimeOk returns a tuple with the ValidDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsentConstraints) GetValidDateTimeOk() (*PaymentConsentValidDateTime, bool) {
	if o == nil || o.ValidDateTime == nil {
		return nil, false
	}
	return o.ValidDateTime, true
}

// HasValidDateTime returns a boolean if a field has been set.
func (o *PaymentInitiationConsentConstraints) HasValidDateTime() bool {
	if o != nil && o.ValidDateTime != nil {
		return true
	}

	return false
}

// SetValidDateTime gets a reference to the given PaymentConsentValidDateTime and assigns it to the ValidDateTime field.
func (o *PaymentInitiationConsentConstraints) SetValidDateTime(v PaymentConsentValidDateTime) {
	o.ValidDateTime = &v
}

// GetMaxPaymentAmount returns the MaxPaymentAmount field value if set, zero value otherwise.
func (o *PaymentInitiationConsentConstraints) GetMaxPaymentAmount() PaymentConsentMaxPaymentAmount {
	if o == nil || o.MaxPaymentAmount == nil {
		var ret PaymentConsentMaxPaymentAmount
		return ret
	}
	return *o.MaxPaymentAmount
}

// GetMaxPaymentAmountOk returns a tuple with the MaxPaymentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsentConstraints) GetMaxPaymentAmountOk() (*PaymentConsentMaxPaymentAmount, bool) {
	if o == nil || o.MaxPaymentAmount == nil {
		return nil, false
	}
	return o.MaxPaymentAmount, true
}

// HasMaxPaymentAmount returns a boolean if a field has been set.
func (o *PaymentInitiationConsentConstraints) HasMaxPaymentAmount() bool {
	if o != nil && o.MaxPaymentAmount != nil {
		return true
	}

	return false
}

// SetMaxPaymentAmount gets a reference to the given PaymentConsentMaxPaymentAmount and assigns it to the MaxPaymentAmount field.
func (o *PaymentInitiationConsentConstraints) SetMaxPaymentAmount(v PaymentConsentMaxPaymentAmount) {
	o.MaxPaymentAmount = &v
}

// GetPeriodicAmounts returns the PeriodicAmounts field value if set, zero value otherwise.
func (o *PaymentInitiationConsentConstraints) GetPeriodicAmounts() []PaymentConsentPeriodicAmount {
	if o == nil || o.PeriodicAmounts == nil {
		var ret []PaymentConsentPeriodicAmount
		return ret
	}
	return *o.PeriodicAmounts
}

// GetPeriodicAmountsOk returns a tuple with the PeriodicAmounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsentConstraints) GetPeriodicAmountsOk() (*[]PaymentConsentPeriodicAmount, bool) {
	if o == nil || o.PeriodicAmounts == nil {
		return nil, false
	}
	return o.PeriodicAmounts, true
}

// HasPeriodicAmounts returns a boolean if a field has been set.
func (o *PaymentInitiationConsentConstraints) HasPeriodicAmounts() bool {
	if o != nil && o.PeriodicAmounts != nil {
		return true
	}

	return false
}

// SetPeriodicAmounts gets a reference to the given []PaymentConsentPeriodicAmount and assigns it to the PeriodicAmounts field.
func (o *PaymentInitiationConsentConstraints) SetPeriodicAmounts(v []PaymentConsentPeriodicAmount) {
	o.PeriodicAmounts = &v
}

func (o PaymentInitiationConsentConstraints) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ValidDateTime != nil {
		toSerialize["valid_date_time"] = o.ValidDateTime
	}
	if o.MaxPaymentAmount != nil {
		toSerialize["max_payment_amount"] = o.MaxPaymentAmount
	}
	if o.PeriodicAmounts != nil {
		toSerialize["periodic_amounts"] = o.PeriodicAmounts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentInitiationConsentConstraints) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentInitiationConsentConstraints := _PaymentInitiationConsentConstraints{}

	if err = json.Unmarshal(bytes, &varPaymentInitiationConsentConstraints); err == nil {
		*o = PaymentInitiationConsentConstraints(varPaymentInitiationConsentConstraints)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "valid_date_time")
		delete(additionalProperties, "max_payment_amount")
		delete(additionalProperties, "periodic_amounts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentInitiationConsentConstraints struct {
	value *PaymentInitiationConsentConstraints
	isSet bool
}

func (v NullablePaymentInitiationConsentConstraints) Get() *PaymentInitiationConsentConstraints {
	return v.value
}

func (v *NullablePaymentInitiationConsentConstraints) Set(val *PaymentInitiationConsentConstraints) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationConsentConstraints) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationConsentConstraints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationConsentConstraints(val *PaymentInitiationConsentConstraints) *NullablePaymentInitiationConsentConstraints {
	return &NullablePaymentInitiationConsentConstraints{value: val, isSet: true}
}

func (v NullablePaymentInitiationConsentConstraints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationConsentConstraints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

