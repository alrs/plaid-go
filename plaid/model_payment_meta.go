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

// PaymentMeta Transaction information specific to inter-bank transfers. If the transaction was not an inter-bank transfer, all fields will be `null`.  If the `transactions` object was returned by a Transactions endpoint such as `/transactions/get`, the `payment_meta` key will always appear, but no data elements are guaranteed. If the `transactions` object was returned by an Assets endpoint such as `/asset_report/get/` or `/asset_report/pdf/get`, this field will only appear in an Asset Report with Insights.
type PaymentMeta struct {
	// The transaction reference number supplied by the financial institution.
	ReferenceNumber NullableString `json:"reference_number"`
	// The ACH PPD ID for the payer.
	PpdId NullableString `json:"ppd_id"`
	// For transfers, the party that is receiving the transaction.
	Payee NullableString `json:"payee"`
	// The party initiating a wire transfer. Will be `null` if the transaction is not a wire transfer.
	ByOrderOf NullableString `json:"by_order_of"`
	// For transfers, the party that is paying the transaction.
	Payer NullableString `json:"payer"`
	// The type of transfer, e.g. 'ACH'
	PaymentMethod NullableString `json:"payment_method"`
	// The name of the payment processor
	PaymentProcessor NullableString `json:"payment_processor"`
	// The payer-supplied description of the transfer.
	Reason NullableString `json:"reason"`
	AdditionalProperties map[string]interface{}
}

type _PaymentMeta PaymentMeta

// NewPaymentMeta instantiates a new PaymentMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMeta(referenceNumber NullableString, ppdId NullableString, payee NullableString, byOrderOf NullableString, payer NullableString, paymentMethod NullableString, paymentProcessor NullableString, reason NullableString) *PaymentMeta {
	this := PaymentMeta{}
	this.ReferenceNumber = referenceNumber
	this.PpdId = ppdId
	this.Payee = payee
	this.ByOrderOf = byOrderOf
	this.Payer = payer
	this.PaymentMethod = paymentMethod
	this.PaymentProcessor = paymentProcessor
	this.Reason = reason
	return &this
}

// NewPaymentMetaWithDefaults instantiates a new PaymentMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMetaWithDefaults() *PaymentMeta {
	this := PaymentMeta{}
	return &this
}

// GetReferenceNumber returns the ReferenceNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentMeta) GetReferenceNumber() string {
	if o == nil || o.ReferenceNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.ReferenceNumber.Get()
}

// GetReferenceNumberOk returns a tuple with the ReferenceNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMeta) GetReferenceNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReferenceNumber.Get(), o.ReferenceNumber.IsSet()
}

// SetReferenceNumber sets field value
func (o *PaymentMeta) SetReferenceNumber(v string) {
	o.ReferenceNumber.Set(&v)
}

// GetPpdId returns the PpdId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentMeta) GetPpdId() string {
	if o == nil || o.PpdId.Get() == nil {
		var ret string
		return ret
	}

	return *o.PpdId.Get()
}

// GetPpdIdOk returns a tuple with the PpdId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMeta) GetPpdIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PpdId.Get(), o.PpdId.IsSet()
}

// SetPpdId sets field value
func (o *PaymentMeta) SetPpdId(v string) {
	o.PpdId.Set(&v)
}

// GetPayee returns the Payee field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentMeta) GetPayee() string {
	if o == nil || o.Payee.Get() == nil {
		var ret string
		return ret
	}

	return *o.Payee.Get()
}

// GetPayeeOk returns a tuple with the Payee field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMeta) GetPayeeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Payee.Get(), o.Payee.IsSet()
}

// SetPayee sets field value
func (o *PaymentMeta) SetPayee(v string) {
	o.Payee.Set(&v)
}

// GetByOrderOf returns the ByOrderOf field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentMeta) GetByOrderOf() string {
	if o == nil || o.ByOrderOf.Get() == nil {
		var ret string
		return ret
	}

	return *o.ByOrderOf.Get()
}

// GetByOrderOfOk returns a tuple with the ByOrderOf field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMeta) GetByOrderOfOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ByOrderOf.Get(), o.ByOrderOf.IsSet()
}

// SetByOrderOf sets field value
func (o *PaymentMeta) SetByOrderOf(v string) {
	o.ByOrderOf.Set(&v)
}

// GetPayer returns the Payer field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentMeta) GetPayer() string {
	if o == nil || o.Payer.Get() == nil {
		var ret string
		return ret
	}

	return *o.Payer.Get()
}

// GetPayerOk returns a tuple with the Payer field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMeta) GetPayerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Payer.Get(), o.Payer.IsSet()
}

// SetPayer sets field value
func (o *PaymentMeta) SetPayer(v string) {
	o.Payer.Set(&v)
}

// GetPaymentMethod returns the PaymentMethod field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentMeta) GetPaymentMethod() string {
	if o == nil || o.PaymentMethod.Get() == nil {
		var ret string
		return ret
	}

	return *o.PaymentMethod.Get()
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMeta) GetPaymentMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaymentMethod.Get(), o.PaymentMethod.IsSet()
}

// SetPaymentMethod sets field value
func (o *PaymentMeta) SetPaymentMethod(v string) {
	o.PaymentMethod.Set(&v)
}

// GetPaymentProcessor returns the PaymentProcessor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentMeta) GetPaymentProcessor() string {
	if o == nil || o.PaymentProcessor.Get() == nil {
		var ret string
		return ret
	}

	return *o.PaymentProcessor.Get()
}

// GetPaymentProcessorOk returns a tuple with the PaymentProcessor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMeta) GetPaymentProcessorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaymentProcessor.Get(), o.PaymentProcessor.IsSet()
}

// SetPaymentProcessor sets field value
func (o *PaymentMeta) SetPaymentProcessor(v string) {
	o.PaymentProcessor.Set(&v)
}

// GetReason returns the Reason field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentMeta) GetReason() string {
	if o == nil || o.Reason.Get() == nil {
		var ret string
		return ret
	}

	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMeta) GetReasonOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// SetReason sets field value
func (o *PaymentMeta) SetReason(v string) {
	o.Reason.Set(&v)
}

func (o PaymentMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reference_number"] = o.ReferenceNumber.Get()
	}
	if true {
		toSerialize["ppd_id"] = o.PpdId.Get()
	}
	if true {
		toSerialize["payee"] = o.Payee.Get()
	}
	if true {
		toSerialize["by_order_of"] = o.ByOrderOf.Get()
	}
	if true {
		toSerialize["payer"] = o.Payer.Get()
	}
	if true {
		toSerialize["payment_method"] = o.PaymentMethod.Get()
	}
	if true {
		toSerialize["payment_processor"] = o.PaymentProcessor.Get()
	}
	if true {
		toSerialize["reason"] = o.Reason.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentMeta) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentMeta := _PaymentMeta{}

	if err = json.Unmarshal(bytes, &varPaymentMeta); err == nil {
		*o = PaymentMeta(varPaymentMeta)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "reference_number")
		delete(additionalProperties, "ppd_id")
		delete(additionalProperties, "payee")
		delete(additionalProperties, "by_order_of")
		delete(additionalProperties, "payer")
		delete(additionalProperties, "payment_method")
		delete(additionalProperties, "payment_processor")
		delete(additionalProperties, "reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentMeta struct {
	value *PaymentMeta
	isSet bool
}

func (v NullablePaymentMeta) Get() *PaymentMeta {
	return v.value
}

func (v *NullablePaymentMeta) Set(val *PaymentMeta) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMeta) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMeta(val *PaymentMeta) *NullablePaymentMeta {
	return &NullablePaymentMeta{value: val, isSet: true}
}

func (v NullablePaymentMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


