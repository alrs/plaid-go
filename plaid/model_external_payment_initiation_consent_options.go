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

// ExternalPaymentInitiationConsentOptions Additional payment consent options
type ExternalPaymentInitiationConsentOptions struct {
	// When `true`, Plaid will attempt to request refund details from the payee's financial institution.  Support varies between financial institutions and will not always be available.  If refund details could be retrieved, they will be available in the `/payment_initiation/payment/get` response.
	RequestRefundDetails NullableBool `json:"request_refund_details,omitempty"`
	// The International Bank Account Number (IBAN) for the payer's account. Where possible, the end user will be able to set up payment consent using only the specified bank account if provided.
	Iban NullableString `json:"iban,omitempty"`
	Bacs NullablePaymentInitiationOptionalRestrictionBacs `json:"bacs,omitempty"`
}

// NewExternalPaymentInitiationConsentOptions instantiates a new ExternalPaymentInitiationConsentOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPaymentInitiationConsentOptions() *ExternalPaymentInitiationConsentOptions {
	this := ExternalPaymentInitiationConsentOptions{}
	return &this
}

// NewExternalPaymentInitiationConsentOptionsWithDefaults instantiates a new ExternalPaymentInitiationConsentOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPaymentInitiationConsentOptionsWithDefaults() *ExternalPaymentInitiationConsentOptions {
	this := ExternalPaymentInitiationConsentOptions{}
	return &this
}

// GetRequestRefundDetails returns the RequestRefundDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalPaymentInitiationConsentOptions) GetRequestRefundDetails() bool {
	if o == nil || o.RequestRefundDetails.Get() == nil {
		var ret bool
		return ret
	}
	return *o.RequestRefundDetails.Get()
}

// GetRequestRefundDetailsOk returns a tuple with the RequestRefundDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalPaymentInitiationConsentOptions) GetRequestRefundDetailsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RequestRefundDetails.Get(), o.RequestRefundDetails.IsSet()
}

// HasRequestRefundDetails returns a boolean if a field has been set.
func (o *ExternalPaymentInitiationConsentOptions) HasRequestRefundDetails() bool {
	if o != nil && o.RequestRefundDetails.IsSet() {
		return true
	}

	return false
}

// SetRequestRefundDetails gets a reference to the given NullableBool and assigns it to the RequestRefundDetails field.
func (o *ExternalPaymentInitiationConsentOptions) SetRequestRefundDetails(v bool) {
	o.RequestRefundDetails.Set(&v)
}
// SetRequestRefundDetailsNil sets the value for RequestRefundDetails to be an explicit nil
func (o *ExternalPaymentInitiationConsentOptions) SetRequestRefundDetailsNil() {
	o.RequestRefundDetails.Set(nil)
}

// UnsetRequestRefundDetails ensures that no value is present for RequestRefundDetails, not even an explicit nil
func (o *ExternalPaymentInitiationConsentOptions) UnsetRequestRefundDetails() {
	o.RequestRefundDetails.Unset()
}

// GetIban returns the Iban field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalPaymentInitiationConsentOptions) GetIban() string {
	if o == nil || o.Iban.Get() == nil {
		var ret string
		return ret
	}
	return *o.Iban.Get()
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalPaymentInitiationConsentOptions) GetIbanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Iban.Get(), o.Iban.IsSet()
}

// HasIban returns a boolean if a field has been set.
func (o *ExternalPaymentInitiationConsentOptions) HasIban() bool {
	if o != nil && o.Iban.IsSet() {
		return true
	}

	return false
}

// SetIban gets a reference to the given NullableString and assigns it to the Iban field.
func (o *ExternalPaymentInitiationConsentOptions) SetIban(v string) {
	o.Iban.Set(&v)
}
// SetIbanNil sets the value for Iban to be an explicit nil
func (o *ExternalPaymentInitiationConsentOptions) SetIbanNil() {
	o.Iban.Set(nil)
}

// UnsetIban ensures that no value is present for Iban, not even an explicit nil
func (o *ExternalPaymentInitiationConsentOptions) UnsetIban() {
	o.Iban.Unset()
}

// GetBacs returns the Bacs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalPaymentInitiationConsentOptions) GetBacs() PaymentInitiationOptionalRestrictionBacs {
	if o == nil || o.Bacs.Get() == nil {
		var ret PaymentInitiationOptionalRestrictionBacs
		return ret
	}
	return *o.Bacs.Get()
}

// GetBacsOk returns a tuple with the Bacs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalPaymentInitiationConsentOptions) GetBacsOk() (*PaymentInitiationOptionalRestrictionBacs, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Bacs.Get(), o.Bacs.IsSet()
}

// HasBacs returns a boolean if a field has been set.
func (o *ExternalPaymentInitiationConsentOptions) HasBacs() bool {
	if o != nil && o.Bacs.IsSet() {
		return true
	}

	return false
}

// SetBacs gets a reference to the given NullablePaymentInitiationOptionalRestrictionBacs and assigns it to the Bacs field.
func (o *ExternalPaymentInitiationConsentOptions) SetBacs(v PaymentInitiationOptionalRestrictionBacs) {
	o.Bacs.Set(&v)
}
// SetBacsNil sets the value for Bacs to be an explicit nil
func (o *ExternalPaymentInitiationConsentOptions) SetBacsNil() {
	o.Bacs.Set(nil)
}

// UnsetBacs ensures that no value is present for Bacs, not even an explicit nil
func (o *ExternalPaymentInitiationConsentOptions) UnsetBacs() {
	o.Bacs.Unset()
}

func (o ExternalPaymentInitiationConsentOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestRefundDetails.IsSet() {
		toSerialize["request_refund_details"] = o.RequestRefundDetails.Get()
	}
	if o.Iban.IsSet() {
		toSerialize["iban"] = o.Iban.Get()
	}
	if o.Bacs.IsSet() {
		toSerialize["bacs"] = o.Bacs.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableExternalPaymentInitiationConsentOptions struct {
	value *ExternalPaymentInitiationConsentOptions
	isSet bool
}

func (v NullableExternalPaymentInitiationConsentOptions) Get() *ExternalPaymentInitiationConsentOptions {
	return v.value
}

func (v *NullableExternalPaymentInitiationConsentOptions) Set(val *ExternalPaymentInitiationConsentOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPaymentInitiationConsentOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPaymentInitiationConsentOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPaymentInitiationConsentOptions(val *ExternalPaymentInitiationConsentOptions) *NullableExternalPaymentInitiationConsentOptions {
	return &NullableExternalPaymentInitiationConsentOptions{value: val, isSet: true}
}

func (v NullableExternalPaymentInitiationConsentOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPaymentInitiationConsentOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


