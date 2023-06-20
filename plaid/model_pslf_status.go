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

// PSLFStatus Information about the student's eligibility in the Public Service Loan Forgiveness program. This is only returned if the institution is FedLoan (`ins_116527`). 
type PSLFStatus struct {
	// The estimated date borrower will have completed 120 qualifying monthly payments. Returned in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).
	EstimatedEligibilityDate NullableString `json:"estimated_eligibility_date"`
	// The number of qualifying payments that have been made.
	PaymentsMade NullableInt32 `json:"payments_made"`
	// The number of qualifying payments remaining.
	PaymentsRemaining NullableInt32 `json:"payments_remaining"`
	AdditionalProperties map[string]interface{}
}

type _PSLFStatus PSLFStatus

// NewPSLFStatus instantiates a new PSLFStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPSLFStatus(estimatedEligibilityDate NullableString, paymentsMade NullableInt32, paymentsRemaining NullableInt32) *PSLFStatus {
	this := PSLFStatus{}
	this.EstimatedEligibilityDate = estimatedEligibilityDate
	this.PaymentsMade = paymentsMade
	this.PaymentsRemaining = paymentsRemaining
	return &this
}

// NewPSLFStatusWithDefaults instantiates a new PSLFStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPSLFStatusWithDefaults() *PSLFStatus {
	this := PSLFStatus{}
	return &this
}

// GetEstimatedEligibilityDate returns the EstimatedEligibilityDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PSLFStatus) GetEstimatedEligibilityDate() string {
	if o == nil || o.EstimatedEligibilityDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.EstimatedEligibilityDate.Get()
}

// GetEstimatedEligibilityDateOk returns a tuple with the EstimatedEligibilityDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PSLFStatus) GetEstimatedEligibilityDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EstimatedEligibilityDate.Get(), o.EstimatedEligibilityDate.IsSet()
}

// SetEstimatedEligibilityDate sets field value
func (o *PSLFStatus) SetEstimatedEligibilityDate(v string) {
	o.EstimatedEligibilityDate.Set(&v)
}

// GetPaymentsMade returns the PaymentsMade field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PSLFStatus) GetPaymentsMade() int32 {
	if o == nil || o.PaymentsMade.Get() == nil {
		var ret int32
		return ret
	}

	return *o.PaymentsMade.Get()
}

// GetPaymentsMadeOk returns a tuple with the PaymentsMade field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PSLFStatus) GetPaymentsMadeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaymentsMade.Get(), o.PaymentsMade.IsSet()
}

// SetPaymentsMade sets field value
func (o *PSLFStatus) SetPaymentsMade(v int32) {
	o.PaymentsMade.Set(&v)
}

// GetPaymentsRemaining returns the PaymentsRemaining field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PSLFStatus) GetPaymentsRemaining() int32 {
	if o == nil || o.PaymentsRemaining.Get() == nil {
		var ret int32
		return ret
	}

	return *o.PaymentsRemaining.Get()
}

// GetPaymentsRemainingOk returns a tuple with the PaymentsRemaining field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PSLFStatus) GetPaymentsRemainingOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaymentsRemaining.Get(), o.PaymentsRemaining.IsSet()
}

// SetPaymentsRemaining sets field value
func (o *PSLFStatus) SetPaymentsRemaining(v int32) {
	o.PaymentsRemaining.Set(&v)
}

func (o PSLFStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["estimated_eligibility_date"] = o.EstimatedEligibilityDate.Get()
	}
	if true {
		toSerialize["payments_made"] = o.PaymentsMade.Get()
	}
	if true {
		toSerialize["payments_remaining"] = o.PaymentsRemaining.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PSLFStatus) UnmarshalJSON(bytes []byte) (err error) {
	varPSLFStatus := _PSLFStatus{}

	if err = json.Unmarshal(bytes, &varPSLFStatus); err == nil {
		*o = PSLFStatus(varPSLFStatus)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "estimated_eligibility_date")
		delete(additionalProperties, "payments_made")
		delete(additionalProperties, "payments_remaining")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePSLFStatus struct {
	value *PSLFStatus
	isSet bool
}

func (v NullablePSLFStatus) Get() *PSLFStatus {
	return v.value
}

func (v *NullablePSLFStatus) Set(val *PSLFStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePSLFStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePSLFStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePSLFStatus(val *PSLFStatus) *NullablePSLFStatus {
	return &NullablePSLFStatus{value: val, isSet: true}
}

func (v NullablePSLFStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePSLFStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


