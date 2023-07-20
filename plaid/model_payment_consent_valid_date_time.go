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
	"time"
)

// PaymentConsentValidDateTime Life span for the payment consent. After the `to` date the payment consent expires and can no longer be used for payment initiation.
type PaymentConsentValidDateTime struct {
	// The date and time from which the consent should be active, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.
	From NullableTime `json:"from,omitempty"`
	// The date and time at which the consent expires, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.
	To NullableTime `json:"to,omitempty"`
}

// NewPaymentConsentValidDateTime instantiates a new PaymentConsentValidDateTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentConsentValidDateTime() *PaymentConsentValidDateTime {
	this := PaymentConsentValidDateTime{}
	return &this
}

// NewPaymentConsentValidDateTimeWithDefaults instantiates a new PaymentConsentValidDateTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentConsentValidDateTimeWithDefaults() *PaymentConsentValidDateTime {
	this := PaymentConsentValidDateTime{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentConsentValidDateTime) GetFrom() time.Time {
	if o == nil || o.From.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.From.Get()
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentConsentValidDateTime) GetFromOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.From.Get(), o.From.IsSet()
}

// HasFrom returns a boolean if a field has been set.
func (o *PaymentConsentValidDateTime) HasFrom() bool {
	if o != nil && o.From.IsSet() {
		return true
	}

	return false
}

// SetFrom gets a reference to the given NullableTime and assigns it to the From field.
func (o *PaymentConsentValidDateTime) SetFrom(v time.Time) {
	o.From.Set(&v)
}
// SetFromNil sets the value for From to be an explicit nil
func (o *PaymentConsentValidDateTime) SetFromNil() {
	o.From.Set(nil)
}

// UnsetFrom ensures that no value is present for From, not even an explicit nil
func (o *PaymentConsentValidDateTime) UnsetFrom() {
	o.From.Unset()
}

// GetTo returns the To field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentConsentValidDateTime) GetTo() time.Time {
	if o == nil || o.To.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.To.Get()
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentConsentValidDateTime) GetToOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.To.Get(), o.To.IsSet()
}

// HasTo returns a boolean if a field has been set.
func (o *PaymentConsentValidDateTime) HasTo() bool {
	if o != nil && o.To.IsSet() {
		return true
	}

	return false
}

// SetTo gets a reference to the given NullableTime and assigns it to the To field.
func (o *PaymentConsentValidDateTime) SetTo(v time.Time) {
	o.To.Set(&v)
}
// SetToNil sets the value for To to be an explicit nil
func (o *PaymentConsentValidDateTime) SetToNil() {
	o.To.Set(nil)
}

// UnsetTo ensures that no value is present for To, not even an explicit nil
func (o *PaymentConsentValidDateTime) UnsetTo() {
	o.To.Unset()
}

func (o PaymentConsentValidDateTime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.From.IsSet() {
		toSerialize["from"] = o.From.Get()
	}
	if o.To.IsSet() {
		toSerialize["to"] = o.To.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentConsentValidDateTime struct {
	value *PaymentConsentValidDateTime
	isSet bool
}

func (v NullablePaymentConsentValidDateTime) Get() *PaymentConsentValidDateTime {
	return v.value
}

func (v *NullablePaymentConsentValidDateTime) Set(val *PaymentConsentValidDateTime) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentConsentValidDateTime) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentConsentValidDateTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentConsentValidDateTime(val *PaymentConsentValidDateTime) *NullablePaymentConsentValidDateTime {
	return &NullablePaymentConsentValidDateTime{value: val, isSet: true}
}

func (v NullablePaymentConsentValidDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentConsentValidDateTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


