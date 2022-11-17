/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.205.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// SandboxBankTransferSimulateRequest Defines the request schema for `/sandbox/bank_transfer/simulate`
type SandboxBankTransferSimulateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Plaid’s unique identifier for a bank transfer.
	BankTransferId string `json:"bank_transfer_id"`
	// The asynchronous event to be simulated. May be: `posted`, `failed`, or `reversed`.  An error will be returned if the event type is incompatible with the current transfer status. Compatible status --> event type transitions include:  `pending` --> `failed`  `pending` --> `posted`  `posted` --> `reversed` 
	EventType string `json:"event_type"`
	FailureReason NullableBankTransferFailure `json:"failure_reason,omitempty"`
}

// NewSandboxBankTransferSimulateRequest instantiates a new SandboxBankTransferSimulateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxBankTransferSimulateRequest(bankTransferId string, eventType string) *SandboxBankTransferSimulateRequest {
	this := SandboxBankTransferSimulateRequest{}
	this.BankTransferId = bankTransferId
	this.EventType = eventType
	return &this
}

// NewSandboxBankTransferSimulateRequestWithDefaults instantiates a new SandboxBankTransferSimulateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxBankTransferSimulateRequestWithDefaults() *SandboxBankTransferSimulateRequest {
	this := SandboxBankTransferSimulateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SandboxBankTransferSimulateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxBankTransferSimulateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SandboxBankTransferSimulateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SandboxBankTransferSimulateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *SandboxBankTransferSimulateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxBankTransferSimulateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *SandboxBankTransferSimulateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *SandboxBankTransferSimulateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetBankTransferId returns the BankTransferId field value
func (o *SandboxBankTransferSimulateRequest) GetBankTransferId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankTransferId
}

// GetBankTransferIdOk returns a tuple with the BankTransferId field value
// and a boolean to check if the value has been set.
func (o *SandboxBankTransferSimulateRequest) GetBankTransferIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankTransferId, true
}

// SetBankTransferId sets field value
func (o *SandboxBankTransferSimulateRequest) SetBankTransferId(v string) {
	o.BankTransferId = v
}

// GetEventType returns the EventType field value
func (o *SandboxBankTransferSimulateRequest) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *SandboxBankTransferSimulateRequest) GetEventTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *SandboxBankTransferSimulateRequest) SetEventType(v string) {
	o.EventType = v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SandboxBankTransferSimulateRequest) GetFailureReason() BankTransferFailure {
	if o == nil || o.FailureReason.Get() == nil {
		var ret BankTransferFailure
		return ret
	}
	return *o.FailureReason.Get()
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SandboxBankTransferSimulateRequest) GetFailureReasonOk() (*BankTransferFailure, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FailureReason.Get(), o.FailureReason.IsSet()
}

// HasFailureReason returns a boolean if a field has been set.
func (o *SandboxBankTransferSimulateRequest) HasFailureReason() bool {
	if o != nil && o.FailureReason.IsSet() {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given NullableBankTransferFailure and assigns it to the FailureReason field.
func (o *SandboxBankTransferSimulateRequest) SetFailureReason(v BankTransferFailure) {
	o.FailureReason.Set(&v)
}
// SetFailureReasonNil sets the value for FailureReason to be an explicit nil
func (o *SandboxBankTransferSimulateRequest) SetFailureReasonNil() {
	o.FailureReason.Set(nil)
}

// UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
func (o *SandboxBankTransferSimulateRequest) UnsetFailureReason() {
	o.FailureReason.Unset()
}

func (o SandboxBankTransferSimulateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["bank_transfer_id"] = o.BankTransferId
	}
	if true {
		toSerialize["event_type"] = o.EventType
	}
	if o.FailureReason.IsSet() {
		toSerialize["failure_reason"] = o.FailureReason.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxBankTransferSimulateRequest struct {
	value *SandboxBankTransferSimulateRequest
	isSet bool
}

func (v NullableSandboxBankTransferSimulateRequest) Get() *SandboxBankTransferSimulateRequest {
	return v.value
}

func (v *NullableSandboxBankTransferSimulateRequest) Set(val *SandboxBankTransferSimulateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxBankTransferSimulateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxBankTransferSimulateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxBankTransferSimulateRequest(val *SandboxBankTransferSimulateRequest) *NullableSandboxBankTransferSimulateRequest {
	return &NullableSandboxBankTransferSimulateRequest{value: val, isSet: true}
}

func (v NullableSandboxBankTransferSimulateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxBankTransferSimulateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


