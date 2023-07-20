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

// PaymentInitiationPaymentGetResponse PaymentInitiationPaymentGetResponse defines the response schema for `/payment_initation/payment/get`
type PaymentInitiationPaymentGetResponse struct {
	// The ID of the payment. Like all Plaid identifiers, the `payment_id` is case sensitive.
	PaymentId string `json:"payment_id"`
	Amount PaymentAmount `json:"amount"`
	Status PaymentInitiationPaymentStatus `json:"status"`
	// The ID of the recipient
	RecipientId string `json:"recipient_id"`
	// A reference for the payment.
	Reference string `json:"reference"`
	// The value of the reference sent to the bank after adjustment to pass bank validation rules.
	AdjustedReference NullableString `json:"adjusted_reference,omitempty"`
	// The date and time of the last time the `status` was updated, in IS0 8601 format
	LastStatusUpdate time.Time `json:"last_status_update"`
	Schedule NullableExternalPaymentScheduleGet `json:"schedule,omitempty"`
	RefundDetails NullableExternalPaymentRefundDetails `json:"refund_details,omitempty"`
	Bacs NullableSenderBACSNullable `json:"bacs"`
	// The International Bank Account Number (IBAN) for the sender, if specified in the `/payment_initiation/payment/create` call.
	Iban NullableString `json:"iban"`
	// Refund IDs associated with the payment.
	RefundIds []string `json:"refund_ids,omitempty"`
	AmountRefunded *PaymentAmountRefunded `json:"amount_refunded,omitempty"`
	// The EMI (E-Money Institution) wallet that this payment is associated with, if any. This wallet is used as an intermediary account to enable Plaid to reconcile the settlement of funds for Payment Initiation requests.
	WalletId NullableString `json:"wallet_id,omitempty"`
	Scheme NullablePaymentScheme `json:"scheme,omitempty"`
	AdjustedScheme NullablePaymentScheme `json:"adjusted_scheme,omitempty"`
	// The payment consent ID that this payment was initiated with. Is present only when payment was initiated using the payment consent.
	ConsentId NullableString `json:"consent_id,omitempty"`
	// The transaction ID that this payment is associated with, if any. This is present only when a payment was initiated using virtual accounts.
	TransactionId NullableString `json:"transaction_id,omitempty"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _PaymentInitiationPaymentGetResponse PaymentInitiationPaymentGetResponse

// NewPaymentInitiationPaymentGetResponse instantiates a new PaymentInitiationPaymentGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationPaymentGetResponse(paymentId string, amount PaymentAmount, status PaymentInitiationPaymentStatus, recipientId string, reference string, lastStatusUpdate time.Time, bacs NullableSenderBACSNullable, iban NullableString, requestId string) *PaymentInitiationPaymentGetResponse {
	this := PaymentInitiationPaymentGetResponse{}
	this.PaymentId = paymentId
	this.Amount = amount
	this.Status = status
	this.RecipientId = recipientId
	this.Reference = reference
	this.LastStatusUpdate = lastStatusUpdate
	this.Bacs = bacs
	this.Iban = iban
	this.RequestId = requestId
	return &this
}

// NewPaymentInitiationPaymentGetResponseWithDefaults instantiates a new PaymentInitiationPaymentGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationPaymentGetResponseWithDefaults() *PaymentInitiationPaymentGetResponse {
	this := PaymentInitiationPaymentGetResponse{}
	return &this
}

// GetPaymentId returns the PaymentId field value
func (o *PaymentInitiationPaymentGetResponse) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentGetResponse) GetPaymentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *PaymentInitiationPaymentGetResponse) SetPaymentId(v string) {
	o.PaymentId = v
}

// GetAmount returns the Amount field value
func (o *PaymentInitiationPaymentGetResponse) GetAmount() PaymentAmount {
	if o == nil {
		var ret PaymentAmount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentGetResponse) GetAmountOk() (*PaymentAmount, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PaymentInitiationPaymentGetResponse) SetAmount(v PaymentAmount) {
	o.Amount = v
}

// GetStatus returns the Status field value
func (o *PaymentInitiationPaymentGetResponse) GetStatus() PaymentInitiationPaymentStatus {
	if o == nil {
		var ret PaymentInitiationPaymentStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentGetResponse) GetStatusOk() (*PaymentInitiationPaymentStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentInitiationPaymentGetResponse) SetStatus(v PaymentInitiationPaymentStatus) {
	o.Status = v
}

// GetRecipientId returns the RecipientId field value
func (o *PaymentInitiationPaymentGetResponse) GetRecipientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientId
}

// GetRecipientIdOk returns a tuple with the RecipientId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentGetResponse) GetRecipientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecipientId, true
}

// SetRecipientId sets field value
func (o *PaymentInitiationPaymentGetResponse) SetRecipientId(v string) {
	o.RecipientId = v
}

// GetReference returns the Reference field value
func (o *PaymentInitiationPaymentGetResponse) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentGetResponse) GetReferenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *PaymentInitiationPaymentGetResponse) SetReference(v string) {
	o.Reference = v
}

// GetAdjustedReference returns the AdjustedReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationPaymentGetResponse) GetAdjustedReference() string {
	if o == nil || o.AdjustedReference.Get() == nil {
		var ret string
		return ret
	}
	return *o.AdjustedReference.Get()
}

// GetAdjustedReferenceOk returns a tuple with the AdjustedReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationPaymentGetResponse) GetAdjustedReferenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AdjustedReference.Get(), o.AdjustedReference.IsSet()
}

// HasAdjustedReference returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentGetResponse) HasAdjustedReference() bool {
	if o != nil && o.AdjustedReference.IsSet() {
		return true
	}

	return false
}

// SetAdjustedReference gets a reference to the given NullableString and assigns it to the AdjustedReference field.
func (o *PaymentInitiationPaymentGetResponse) SetAdjustedReference(v string) {
	o.AdjustedReference.Set(&v)
}
// SetAdjustedReferenceNil sets the value for AdjustedReference to be an explicit nil
func (o *PaymentInitiationPaymentGetResponse) SetAdjustedReferenceNil() {
	o.AdjustedReference.Set(nil)
}

// UnsetAdjustedReference ensures that no value is present for AdjustedReference, not even an explicit nil
func (o *PaymentInitiationPaymentGetResponse) UnsetAdjustedReference() {
	o.AdjustedReference.Unset()
}

// GetLastStatusUpdate returns the LastStatusUpdate field value
func (o *PaymentInitiationPaymentGetResponse) GetLastStatusUpdate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastStatusUpdate
}

// GetLastStatusUpdateOk returns a tuple with the LastStatusUpdate field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentGetResponse) GetLastStatusUpdateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastStatusUpdate, true
}

// SetLastStatusUpdate sets field value
func (o *PaymentInitiationPaymentGetResponse) SetLastStatusUpdate(v time.Time) {
	o.LastStatusUpdate = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationPaymentGetResponse) GetSchedule() ExternalPaymentScheduleGet {
	if o == nil || o.Schedule.Get() == nil {
		var ret ExternalPaymentScheduleGet
		return ret
	}
	return *o.Schedule.Get()
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationPaymentGetResponse) GetScheduleOk() (*ExternalPaymentScheduleGet, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Schedule.Get(), o.Schedule.IsSet()
}

// HasSchedule returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentGetResponse) HasSchedule() bool {
	if o != nil && o.Schedule.IsSet() {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given NullableExternalPaymentScheduleGet and assigns it to the Schedule field.
func (o *PaymentInitiationPaymentGetResponse) SetSchedule(v ExternalPaymentScheduleGet) {
	o.Schedule.Set(&v)
}
// SetScheduleNil sets the value for Schedule to be an explicit nil
func (o *PaymentInitiationPaymentGetResponse) SetScheduleNil() {
	o.Schedule.Set(nil)
}

// UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
func (o *PaymentInitiationPaymentGetResponse) UnsetSchedule() {
	o.Schedule.Unset()
}

// GetRefundDetails returns the RefundDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationPaymentGetResponse) GetRefundDetails() ExternalPaymentRefundDetails {
	if o == nil || o.RefundDetails.Get() == nil {
		var ret ExternalPaymentRefundDetails
		return ret
	}
	return *o.RefundDetails.Get()
}

// GetRefundDetailsOk returns a tuple with the RefundDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationPaymentGetResponse) GetRefundDetailsOk() (*ExternalPaymentRefundDetails, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefundDetails.Get(), o.RefundDetails.IsSet()
}

// HasRefundDetails returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentGetResponse) HasRefundDetails() bool {
	if o != nil && o.RefundDetails.IsSet() {
		return true
	}

	return false
}

// SetRefundDetails gets a reference to the given NullableExternalPaymentRefundDetails and assigns it to the RefundDetails field.
func (o *PaymentInitiationPaymentGetResponse) SetRefundDetails(v ExternalPaymentRefundDetails) {
	o.RefundDetails.Set(&v)
}
// SetRefundDetailsNil sets the value for RefundDetails to be an explicit nil
func (o *PaymentInitiationPaymentGetResponse) SetRefundDetailsNil() {
	o.RefundDetails.Set(nil)
}

// UnsetRefundDetails ensures that no value is present for RefundDetails, not even an explicit nil
func (o *PaymentInitiationPaymentGetResponse) UnsetRefundDetails() {
	o.RefundDetails.Unset()
}

// GetBacs returns the Bacs field value
// If the value is explicit nil, the zero value for SenderBACSNullable will be returned
func (o *PaymentInitiationPaymentGetResponse) GetBacs() SenderBACSNullable {
	if o == nil || o.Bacs.Get() == nil {
		var ret SenderBACSNullable
		return ret
	}

	return *o.Bacs.Get()
}

// GetBacsOk returns a tuple with the Bacs field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationPaymentGetResponse) GetBacsOk() (*SenderBACSNullable, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Bacs.Get(), o.Bacs.IsSet()
}

// SetBacs sets field value
func (o *PaymentInitiationPaymentGetResponse) SetBacs(v SenderBACSNullable) {
	o.Bacs.Set(&v)
}

// GetIban returns the Iban field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentInitiationPaymentGetResponse) GetIban() string {
	if o == nil || o.Iban.Get() == nil {
		var ret string
		return ret
	}

	return *o.Iban.Get()
}

// GetIbanOk returns a tuple with the Iban field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationPaymentGetResponse) GetIbanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Iban.Get(), o.Iban.IsSet()
}

// SetIban sets field value
func (o *PaymentInitiationPaymentGetResponse) SetIban(v string) {
	o.Iban.Set(&v)
}

// GetRefundIds returns the RefundIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationPaymentGetResponse) GetRefundIds() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.RefundIds
}

// GetRefundIdsOk returns a tuple with the RefundIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationPaymentGetResponse) GetRefundIdsOk() (*[]string, bool) {
	if o == nil || o.RefundIds == nil {
		return nil, false
	}
	return &o.RefundIds, true
}

// HasRefundIds returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentGetResponse) HasRefundIds() bool {
	if o != nil && o.RefundIds != nil {
		return true
	}

	return false
}

// SetRefundIds gets a reference to the given []string and assigns it to the RefundIds field.
func (o *PaymentInitiationPaymentGetResponse) SetRefundIds(v []string) {
	o.RefundIds = v
}

// GetAmountRefunded returns the AmountRefunded field value if set, zero value otherwise.
func (o *PaymentInitiationPaymentGetResponse) GetAmountRefunded() PaymentAmountRefunded {
	if o == nil || o.AmountRefunded == nil {
		var ret PaymentAmountRefunded
		return ret
	}
	return *o.AmountRefunded
}

// GetAmountRefundedOk returns a tuple with the AmountRefunded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentGetResponse) GetAmountRefundedOk() (*PaymentAmountRefunded, bool) {
	if o == nil || o.AmountRefunded == nil {
		return nil, false
	}
	return o.AmountRefunded, true
}

// HasAmountRefunded returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentGetResponse) HasAmountRefunded() bool {
	if o != nil && o.AmountRefunded != nil {
		return true
	}

	return false
}

// SetAmountRefunded gets a reference to the given PaymentAmountRefunded and assigns it to the AmountRefunded field.
func (o *PaymentInitiationPaymentGetResponse) SetAmountRefunded(v PaymentAmountRefunded) {
	o.AmountRefunded = &v
}

// GetWalletId returns the WalletId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationPaymentGetResponse) GetWalletId() string {
	if o == nil || o.WalletId.Get() == nil {
		var ret string
		return ret
	}
	return *o.WalletId.Get()
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationPaymentGetResponse) GetWalletIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WalletId.Get(), o.WalletId.IsSet()
}

// HasWalletId returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentGetResponse) HasWalletId() bool {
	if o != nil && o.WalletId.IsSet() {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given NullableString and assigns it to the WalletId field.
func (o *PaymentInitiationPaymentGetResponse) SetWalletId(v string) {
	o.WalletId.Set(&v)
}
// SetWalletIdNil sets the value for WalletId to be an explicit nil
func (o *PaymentInitiationPaymentGetResponse) SetWalletIdNil() {
	o.WalletId.Set(nil)
}

// UnsetWalletId ensures that no value is present for WalletId, not even an explicit nil
func (o *PaymentInitiationPaymentGetResponse) UnsetWalletId() {
	o.WalletId.Unset()
}

// GetScheme returns the Scheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationPaymentGetResponse) GetScheme() PaymentScheme {
	if o == nil || o.Scheme.Get() == nil {
		var ret PaymentScheme
		return ret
	}
	return *o.Scheme.Get()
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationPaymentGetResponse) GetSchemeOk() (*PaymentScheme, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Scheme.Get(), o.Scheme.IsSet()
}

// HasScheme returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentGetResponse) HasScheme() bool {
	if o != nil && o.Scheme.IsSet() {
		return true
	}

	return false
}

// SetScheme gets a reference to the given NullablePaymentScheme and assigns it to the Scheme field.
func (o *PaymentInitiationPaymentGetResponse) SetScheme(v PaymentScheme) {
	o.Scheme.Set(&v)
}
// SetSchemeNil sets the value for Scheme to be an explicit nil
func (o *PaymentInitiationPaymentGetResponse) SetSchemeNil() {
	o.Scheme.Set(nil)
}

// UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
func (o *PaymentInitiationPaymentGetResponse) UnsetScheme() {
	o.Scheme.Unset()
}

// GetAdjustedScheme returns the AdjustedScheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationPaymentGetResponse) GetAdjustedScheme() PaymentScheme {
	if o == nil || o.AdjustedScheme.Get() == nil {
		var ret PaymentScheme
		return ret
	}
	return *o.AdjustedScheme.Get()
}

// GetAdjustedSchemeOk returns a tuple with the AdjustedScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationPaymentGetResponse) GetAdjustedSchemeOk() (*PaymentScheme, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AdjustedScheme.Get(), o.AdjustedScheme.IsSet()
}

// HasAdjustedScheme returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentGetResponse) HasAdjustedScheme() bool {
	if o != nil && o.AdjustedScheme.IsSet() {
		return true
	}

	return false
}

// SetAdjustedScheme gets a reference to the given NullablePaymentScheme and assigns it to the AdjustedScheme field.
func (o *PaymentInitiationPaymentGetResponse) SetAdjustedScheme(v PaymentScheme) {
	o.AdjustedScheme.Set(&v)
}
// SetAdjustedSchemeNil sets the value for AdjustedScheme to be an explicit nil
func (o *PaymentInitiationPaymentGetResponse) SetAdjustedSchemeNil() {
	o.AdjustedScheme.Set(nil)
}

// UnsetAdjustedScheme ensures that no value is present for AdjustedScheme, not even an explicit nil
func (o *PaymentInitiationPaymentGetResponse) UnsetAdjustedScheme() {
	o.AdjustedScheme.Unset()
}

// GetConsentId returns the ConsentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationPaymentGetResponse) GetConsentId() string {
	if o == nil || o.ConsentId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConsentId.Get()
}

// GetConsentIdOk returns a tuple with the ConsentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationPaymentGetResponse) GetConsentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConsentId.Get(), o.ConsentId.IsSet()
}

// HasConsentId returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentGetResponse) HasConsentId() bool {
	if o != nil && o.ConsentId.IsSet() {
		return true
	}

	return false
}

// SetConsentId gets a reference to the given NullableString and assigns it to the ConsentId field.
func (o *PaymentInitiationPaymentGetResponse) SetConsentId(v string) {
	o.ConsentId.Set(&v)
}
// SetConsentIdNil sets the value for ConsentId to be an explicit nil
func (o *PaymentInitiationPaymentGetResponse) SetConsentIdNil() {
	o.ConsentId.Set(nil)
}

// UnsetConsentId ensures that no value is present for ConsentId, not even an explicit nil
func (o *PaymentInitiationPaymentGetResponse) UnsetConsentId() {
	o.ConsentId.Unset()
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationPaymentGetResponse) GetTransactionId() string {
	if o == nil || o.TransactionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TransactionId.Get()
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationPaymentGetResponse) GetTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TransactionId.Get(), o.TransactionId.IsSet()
}

// HasTransactionId returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentGetResponse) HasTransactionId() bool {
	if o != nil && o.TransactionId.IsSet() {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given NullableString and assigns it to the TransactionId field.
func (o *PaymentInitiationPaymentGetResponse) SetTransactionId(v string) {
	o.TransactionId.Set(&v)
}
// SetTransactionIdNil sets the value for TransactionId to be an explicit nil
func (o *PaymentInitiationPaymentGetResponse) SetTransactionIdNil() {
	o.TransactionId.Set(nil)
}

// UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
func (o *PaymentInitiationPaymentGetResponse) UnsetTransactionId() {
	o.TransactionId.Unset()
}

// GetRequestId returns the RequestId field value
func (o *PaymentInitiationPaymentGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *PaymentInitiationPaymentGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o PaymentInitiationPaymentGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["payment_id"] = o.PaymentId
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["recipient_id"] = o.RecipientId
	}
	if true {
		toSerialize["reference"] = o.Reference
	}
	if o.AdjustedReference.IsSet() {
		toSerialize["adjusted_reference"] = o.AdjustedReference.Get()
	}
	if true {
		toSerialize["last_status_update"] = o.LastStatusUpdate
	}
	if o.Schedule.IsSet() {
		toSerialize["schedule"] = o.Schedule.Get()
	}
	if o.RefundDetails.IsSet() {
		toSerialize["refund_details"] = o.RefundDetails.Get()
	}
	if true {
		toSerialize["bacs"] = o.Bacs.Get()
	}
	if true {
		toSerialize["iban"] = o.Iban.Get()
	}
	if o.RefundIds != nil {
		toSerialize["refund_ids"] = o.RefundIds
	}
	if o.AmountRefunded != nil {
		toSerialize["amount_refunded"] = o.AmountRefunded
	}
	if o.WalletId.IsSet() {
		toSerialize["wallet_id"] = o.WalletId.Get()
	}
	if o.Scheme.IsSet() {
		toSerialize["scheme"] = o.Scheme.Get()
	}
	if o.AdjustedScheme.IsSet() {
		toSerialize["adjusted_scheme"] = o.AdjustedScheme.Get()
	}
	if o.ConsentId.IsSet() {
		toSerialize["consent_id"] = o.ConsentId.Get()
	}
	if o.TransactionId.IsSet() {
		toSerialize["transaction_id"] = o.TransactionId.Get()
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentInitiationPaymentGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentInitiationPaymentGetResponse := _PaymentInitiationPaymentGetResponse{}

	if err = json.Unmarshal(bytes, &varPaymentInitiationPaymentGetResponse); err == nil {
		*o = PaymentInitiationPaymentGetResponse(varPaymentInitiationPaymentGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "payment_id")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "status")
		delete(additionalProperties, "recipient_id")
		delete(additionalProperties, "reference")
		delete(additionalProperties, "adjusted_reference")
		delete(additionalProperties, "last_status_update")
		delete(additionalProperties, "schedule")
		delete(additionalProperties, "refund_details")
		delete(additionalProperties, "bacs")
		delete(additionalProperties, "iban")
		delete(additionalProperties, "refund_ids")
		delete(additionalProperties, "amount_refunded")
		delete(additionalProperties, "wallet_id")
		delete(additionalProperties, "scheme")
		delete(additionalProperties, "adjusted_scheme")
		delete(additionalProperties, "consent_id")
		delete(additionalProperties, "transaction_id")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentInitiationPaymentGetResponse struct {
	value *PaymentInitiationPaymentGetResponse
	isSet bool
}

func (v NullablePaymentInitiationPaymentGetResponse) Get() *PaymentInitiationPaymentGetResponse {
	return v.value
}

func (v *NullablePaymentInitiationPaymentGetResponse) Set(val *PaymentInitiationPaymentGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationPaymentGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationPaymentGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationPaymentGetResponse(val *PaymentInitiationPaymentGetResponse) *NullablePaymentInitiationPaymentGetResponse {
	return &NullablePaymentInitiationPaymentGetResponse{value: val, isSet: true}
}

func (v NullablePaymentInitiationPaymentGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationPaymentGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


