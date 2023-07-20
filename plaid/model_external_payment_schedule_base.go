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

// ExternalPaymentScheduleBase The schedule that the payment will be executed on. If a schedule is provided, the payment is automatically set up as a standing order. If no schedule is specified, the payment will be executed only once.
type ExternalPaymentScheduleBase struct {
	Interval *PaymentScheduleInterval `json:"interval,omitempty"`
	// The day of the interval on which to schedule the payment.  If the payment interval is weekly, `interval_execution_day` should be an integer from 1 (Monday) to 7 (Sunday).  If the payment interval is monthly, `interval_execution_day` should be an integer indicating which day of the month to make the payment on. Integers from 1 to 28 can be used to make a payment on that day of the month. Negative integers from -1 to -5 can be used to make a payment relative to the end of the month. To make a payment on the last day of the month, use -1; to make the payment on the second-to-last day, use -2, and so on.
	IntervalExecutionDay *int32 `json:"interval_execution_day,omitempty"`
	// A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). Standing order payments will begin on the first `interval_execution_day` on or after the `start_date`.  If the first `interval_execution_day` on or after the start date is also the same day that `/payment_initiation/payment/create` was called, the bank *may* make the first payment on that day, but it is not guaranteed to do so.
	StartDate *string `json:"start_date,omitempty"`
	// A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). Standing order payments will end on the last `interval_execution_day` on or before the `end_date`. If the only `interval_execution_day` between the start date and the end date (inclusive) is also the same day that `/payment_initiation/payment/create` was called, the bank *may* make a payment on that day, but it is not guaranteed to do so.
	EndDate NullableString `json:"end_date,omitempty"`
	// The start date sent to the bank after adjusting for holidays or weekends.  Will be provided in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). If the start date did not require adjustment, this field will be `null`.
	AdjustedStartDate NullableString `json:"adjusted_start_date,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExternalPaymentScheduleBase ExternalPaymentScheduleBase

// NewExternalPaymentScheduleBase instantiates a new ExternalPaymentScheduleBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPaymentScheduleBase() *ExternalPaymentScheduleBase {
	this := ExternalPaymentScheduleBase{}
	return &this
}

// NewExternalPaymentScheduleBaseWithDefaults instantiates a new ExternalPaymentScheduleBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPaymentScheduleBaseWithDefaults() *ExternalPaymentScheduleBase {
	this := ExternalPaymentScheduleBase{}
	return &this
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *ExternalPaymentScheduleBase) GetInterval() PaymentScheduleInterval {
	if o == nil || o.Interval == nil {
		var ret PaymentScheduleInterval
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPaymentScheduleBase) GetIntervalOk() (*PaymentScheduleInterval, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *ExternalPaymentScheduleBase) HasInterval() bool {
	if o != nil && o.Interval != nil {
		return true
	}

	return false
}

// SetInterval gets a reference to the given PaymentScheduleInterval and assigns it to the Interval field.
func (o *ExternalPaymentScheduleBase) SetInterval(v PaymentScheduleInterval) {
	o.Interval = &v
}

// GetIntervalExecutionDay returns the IntervalExecutionDay field value if set, zero value otherwise.
func (o *ExternalPaymentScheduleBase) GetIntervalExecutionDay() int32 {
	if o == nil || o.IntervalExecutionDay == nil {
		var ret int32
		return ret
	}
	return *o.IntervalExecutionDay
}

// GetIntervalExecutionDayOk returns a tuple with the IntervalExecutionDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPaymentScheduleBase) GetIntervalExecutionDayOk() (*int32, bool) {
	if o == nil || o.IntervalExecutionDay == nil {
		return nil, false
	}
	return o.IntervalExecutionDay, true
}

// HasIntervalExecutionDay returns a boolean if a field has been set.
func (o *ExternalPaymentScheduleBase) HasIntervalExecutionDay() bool {
	if o != nil && o.IntervalExecutionDay != nil {
		return true
	}

	return false
}

// SetIntervalExecutionDay gets a reference to the given int32 and assigns it to the IntervalExecutionDay field.
func (o *ExternalPaymentScheduleBase) SetIntervalExecutionDay(v int32) {
	o.IntervalExecutionDay = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ExternalPaymentScheduleBase) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPaymentScheduleBase) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ExternalPaymentScheduleBase) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *ExternalPaymentScheduleBase) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalPaymentScheduleBase) GetEndDate() string {
	if o == nil || o.EndDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalPaymentScheduleBase) GetEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *ExternalPaymentScheduleBase) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableString and assigns it to the EndDate field.
func (o *ExternalPaymentScheduleBase) SetEndDate(v string) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *ExternalPaymentScheduleBase) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *ExternalPaymentScheduleBase) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetAdjustedStartDate returns the AdjustedStartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalPaymentScheduleBase) GetAdjustedStartDate() string {
	if o == nil || o.AdjustedStartDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.AdjustedStartDate.Get()
}

// GetAdjustedStartDateOk returns a tuple with the AdjustedStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalPaymentScheduleBase) GetAdjustedStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AdjustedStartDate.Get(), o.AdjustedStartDate.IsSet()
}

// HasAdjustedStartDate returns a boolean if a field has been set.
func (o *ExternalPaymentScheduleBase) HasAdjustedStartDate() bool {
	if o != nil && o.AdjustedStartDate.IsSet() {
		return true
	}

	return false
}

// SetAdjustedStartDate gets a reference to the given NullableString and assigns it to the AdjustedStartDate field.
func (o *ExternalPaymentScheduleBase) SetAdjustedStartDate(v string) {
	o.AdjustedStartDate.Set(&v)
}
// SetAdjustedStartDateNil sets the value for AdjustedStartDate to be an explicit nil
func (o *ExternalPaymentScheduleBase) SetAdjustedStartDateNil() {
	o.AdjustedStartDate.Set(nil)
}

// UnsetAdjustedStartDate ensures that no value is present for AdjustedStartDate, not even an explicit nil
func (o *ExternalPaymentScheduleBase) UnsetAdjustedStartDate() {
	o.AdjustedStartDate.Unset()
}

func (o ExternalPaymentScheduleBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	if o.IntervalExecutionDay != nil {
		toSerialize["interval_execution_day"] = o.IntervalExecutionDay
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	if o.EndDate.IsSet() {
		toSerialize["end_date"] = o.EndDate.Get()
	}
	if o.AdjustedStartDate.IsSet() {
		toSerialize["adjusted_start_date"] = o.AdjustedStartDate.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ExternalPaymentScheduleBase) UnmarshalJSON(bytes []byte) (err error) {
	varExternalPaymentScheduleBase := _ExternalPaymentScheduleBase{}

	if err = json.Unmarshal(bytes, &varExternalPaymentScheduleBase); err == nil {
		*o = ExternalPaymentScheduleBase(varExternalPaymentScheduleBase)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "interval")
		delete(additionalProperties, "interval_execution_day")
		delete(additionalProperties, "start_date")
		delete(additionalProperties, "end_date")
		delete(additionalProperties, "adjusted_start_date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExternalPaymentScheduleBase struct {
	value *ExternalPaymentScheduleBase
	isSet bool
}

func (v NullableExternalPaymentScheduleBase) Get() *ExternalPaymentScheduleBase {
	return v.value
}

func (v *NullableExternalPaymentScheduleBase) Set(val *ExternalPaymentScheduleBase) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPaymentScheduleBase) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPaymentScheduleBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPaymentScheduleBase(val *ExternalPaymentScheduleBase) *NullableExternalPaymentScheduleBase {
	return &NullableExternalPaymentScheduleBase{value: val, isSet: true}
}

func (v NullableExternalPaymentScheduleBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPaymentScheduleBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


