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
	"time"
)

// PaymentStatusUpdateWebhook Fired when the status of a payment has changed.
type PaymentStatusUpdateWebhook struct {
	// `PAYMENT_INITIATION`
	WebhookType string `json:"webhook_type"`
	// `PAYMENT_STATUS_UPDATE`
	WebhookCode string `json:"webhook_code"`
	// The `payment_id` for the payment being updated
	PaymentId string `json:"payment_id"`
	NewPaymentStatus PaymentInitiationPaymentStatus `json:"new_payment_status"`
	OldPaymentStatus PaymentInitiationPaymentStatus `json:"old_payment_status"`
	// The original value of the reference when creating the payment.
	OriginalReference NullableString `json:"original_reference"`
	// The value of the reference sent to the bank after adjustment to pass bank validation rules.
	AdjustedReference NullableString `json:"adjusted_reference,omitempty"`
	// The original value of the `start_date` provided during the creation of a standing order. If the payment is not a standing order, this field will be `null`.
	OriginalStartDate NullableString `json:"original_start_date"`
	// The start date sent to the bank after adjusting for holidays or weekends.  Will be provided in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). If the start date did not require adjustment, or if the payment is not a standing order, this field will be `null`.
	AdjustedStartDate NullableString `json:"adjusted_start_date"`
	// The timestamp of the update, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format, e.g. `\"2017-09-14T14:42:19.350Z\"`
	Timestamp time.Time `json:"timestamp"`
	Error *PlaidError `json:"error,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaymentStatusUpdateWebhook PaymentStatusUpdateWebhook

// NewPaymentStatusUpdateWebhook instantiates a new PaymentStatusUpdateWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentStatusUpdateWebhook(webhookType string, webhookCode string, paymentId string, newPaymentStatus PaymentInitiationPaymentStatus, oldPaymentStatus PaymentInitiationPaymentStatus, originalReference NullableString, originalStartDate NullableString, adjustedStartDate NullableString, timestamp time.Time) *PaymentStatusUpdateWebhook {
	this := PaymentStatusUpdateWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.PaymentId = paymentId
	this.NewPaymentStatus = newPaymentStatus
	this.OldPaymentStatus = oldPaymentStatus
	this.OriginalReference = originalReference
	this.OriginalStartDate = originalStartDate
	this.AdjustedStartDate = adjustedStartDate
	this.Timestamp = timestamp
	return &this
}

// NewPaymentStatusUpdateWebhookWithDefaults instantiates a new PaymentStatusUpdateWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentStatusUpdateWebhookWithDefaults() *PaymentStatusUpdateWebhook {
	this := PaymentStatusUpdateWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *PaymentStatusUpdateWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *PaymentStatusUpdateWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *PaymentStatusUpdateWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *PaymentStatusUpdateWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *PaymentStatusUpdateWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *PaymentStatusUpdateWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetPaymentId returns the PaymentId field value
func (o *PaymentStatusUpdateWebhook) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *PaymentStatusUpdateWebhook) GetPaymentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *PaymentStatusUpdateWebhook) SetPaymentId(v string) {
	o.PaymentId = v
}

// GetNewPaymentStatus returns the NewPaymentStatus field value
func (o *PaymentStatusUpdateWebhook) GetNewPaymentStatus() PaymentInitiationPaymentStatus {
	if o == nil {
		var ret PaymentInitiationPaymentStatus
		return ret
	}

	return o.NewPaymentStatus
}

// GetNewPaymentStatusOk returns a tuple with the NewPaymentStatus field value
// and a boolean to check if the value has been set.
func (o *PaymentStatusUpdateWebhook) GetNewPaymentStatusOk() (*PaymentInitiationPaymentStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NewPaymentStatus, true
}

// SetNewPaymentStatus sets field value
func (o *PaymentStatusUpdateWebhook) SetNewPaymentStatus(v PaymentInitiationPaymentStatus) {
	o.NewPaymentStatus = v
}

// GetOldPaymentStatus returns the OldPaymentStatus field value
func (o *PaymentStatusUpdateWebhook) GetOldPaymentStatus() PaymentInitiationPaymentStatus {
	if o == nil {
		var ret PaymentInitiationPaymentStatus
		return ret
	}

	return o.OldPaymentStatus
}

// GetOldPaymentStatusOk returns a tuple with the OldPaymentStatus field value
// and a boolean to check if the value has been set.
func (o *PaymentStatusUpdateWebhook) GetOldPaymentStatusOk() (*PaymentInitiationPaymentStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OldPaymentStatus, true
}

// SetOldPaymentStatus sets field value
func (o *PaymentStatusUpdateWebhook) SetOldPaymentStatus(v PaymentInitiationPaymentStatus) {
	o.OldPaymentStatus = v
}

// GetOriginalReference returns the OriginalReference field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentStatusUpdateWebhook) GetOriginalReference() string {
	if o == nil || o.OriginalReference.Get() == nil {
		var ret string
		return ret
	}

	return *o.OriginalReference.Get()
}

// GetOriginalReferenceOk returns a tuple with the OriginalReference field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentStatusUpdateWebhook) GetOriginalReferenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginalReference.Get(), o.OriginalReference.IsSet()
}

// SetOriginalReference sets field value
func (o *PaymentStatusUpdateWebhook) SetOriginalReference(v string) {
	o.OriginalReference.Set(&v)
}

// GetAdjustedReference returns the AdjustedReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentStatusUpdateWebhook) GetAdjustedReference() string {
	if o == nil || o.AdjustedReference.Get() == nil {
		var ret string
		return ret
	}
	return *o.AdjustedReference.Get()
}

// GetAdjustedReferenceOk returns a tuple with the AdjustedReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentStatusUpdateWebhook) GetAdjustedReferenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AdjustedReference.Get(), o.AdjustedReference.IsSet()
}

// HasAdjustedReference returns a boolean if a field has been set.
func (o *PaymentStatusUpdateWebhook) HasAdjustedReference() bool {
	if o != nil && o.AdjustedReference.IsSet() {
		return true
	}

	return false
}

// SetAdjustedReference gets a reference to the given NullableString and assigns it to the AdjustedReference field.
func (o *PaymentStatusUpdateWebhook) SetAdjustedReference(v string) {
	o.AdjustedReference.Set(&v)
}
// SetAdjustedReferenceNil sets the value for AdjustedReference to be an explicit nil
func (o *PaymentStatusUpdateWebhook) SetAdjustedReferenceNil() {
	o.AdjustedReference.Set(nil)
}

// UnsetAdjustedReference ensures that no value is present for AdjustedReference, not even an explicit nil
func (o *PaymentStatusUpdateWebhook) UnsetAdjustedReference() {
	o.AdjustedReference.Unset()
}

// GetOriginalStartDate returns the OriginalStartDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentStatusUpdateWebhook) GetOriginalStartDate() string {
	if o == nil || o.OriginalStartDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.OriginalStartDate.Get()
}

// GetOriginalStartDateOk returns a tuple with the OriginalStartDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentStatusUpdateWebhook) GetOriginalStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginalStartDate.Get(), o.OriginalStartDate.IsSet()
}

// SetOriginalStartDate sets field value
func (o *PaymentStatusUpdateWebhook) SetOriginalStartDate(v string) {
	o.OriginalStartDate.Set(&v)
}

// GetAdjustedStartDate returns the AdjustedStartDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentStatusUpdateWebhook) GetAdjustedStartDate() string {
	if o == nil || o.AdjustedStartDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.AdjustedStartDate.Get()
}

// GetAdjustedStartDateOk returns a tuple with the AdjustedStartDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentStatusUpdateWebhook) GetAdjustedStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AdjustedStartDate.Get(), o.AdjustedStartDate.IsSet()
}

// SetAdjustedStartDate sets field value
func (o *PaymentStatusUpdateWebhook) SetAdjustedStartDate(v string) {
	o.AdjustedStartDate.Set(&v)
}

// GetTimestamp returns the Timestamp field value
func (o *PaymentStatusUpdateWebhook) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *PaymentStatusUpdateWebhook) GetTimestampOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *PaymentStatusUpdateWebhook) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *PaymentStatusUpdateWebhook) GetError() PlaidError {
	if o == nil || o.Error == nil {
		var ret PlaidError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentStatusUpdateWebhook) GetErrorOk() (*PlaidError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *PaymentStatusUpdateWebhook) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given PlaidError and assigns it to the Error field.
func (o *PaymentStatusUpdateWebhook) SetError(v PlaidError) {
	o.Error = &v
}

func (o PaymentStatusUpdateWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["payment_id"] = o.PaymentId
	}
	if true {
		toSerialize["new_payment_status"] = o.NewPaymentStatus
	}
	if true {
		toSerialize["old_payment_status"] = o.OldPaymentStatus
	}
	if true {
		toSerialize["original_reference"] = o.OriginalReference.Get()
	}
	if o.AdjustedReference.IsSet() {
		toSerialize["adjusted_reference"] = o.AdjustedReference.Get()
	}
	if true {
		toSerialize["original_start_date"] = o.OriginalStartDate.Get()
	}
	if true {
		toSerialize["adjusted_start_date"] = o.AdjustedStartDate.Get()
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentStatusUpdateWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentStatusUpdateWebhook := _PaymentStatusUpdateWebhook{}

	if err = json.Unmarshal(bytes, &varPaymentStatusUpdateWebhook); err == nil {
		*o = PaymentStatusUpdateWebhook(varPaymentStatusUpdateWebhook)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "payment_id")
		delete(additionalProperties, "new_payment_status")
		delete(additionalProperties, "old_payment_status")
		delete(additionalProperties, "original_reference")
		delete(additionalProperties, "adjusted_reference")
		delete(additionalProperties, "original_start_date")
		delete(additionalProperties, "adjusted_start_date")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentStatusUpdateWebhook struct {
	value *PaymentStatusUpdateWebhook
	isSet bool
}

func (v NullablePaymentStatusUpdateWebhook) Get() *PaymentStatusUpdateWebhook {
	return v.value
}

func (v *NullablePaymentStatusUpdateWebhook) Set(val *PaymentStatusUpdateWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentStatusUpdateWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentStatusUpdateWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentStatusUpdateWebhook(val *PaymentStatusUpdateWebhook) *NullablePaymentStatusUpdateWebhook {
	return &NullablePaymentStatusUpdateWebhook{value: val, isSet: true}
}

func (v NullablePaymentStatusUpdateWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentStatusUpdateWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


