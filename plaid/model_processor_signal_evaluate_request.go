/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.343.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// ProcessorSignalEvaluateRequest ProcessorSignalEvaluateRequest defines the request schema for `/processor/signal/evaluate`
type ProcessorSignalEvaluateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The processor token obtained from the Plaid integration partner. Processor tokens are in the format: `processor-<environment>-<identifier>`
	ProcessorToken string `json:"processor_token"`
	// The unique ID that you would like to use to refer to this transaction. For your convenience mapping your internal data, you could use your internal ID/identifier for this transaction. The max length for this field is 36 characters.
	ClientTransactionId string `json:"client_transaction_id"`
	// The transaction amount, in USD (e.g. `102.05`)
	Amount float64 `json:"amount"`
	// `true` if the end user is present while initiating the ACH transfer and the endpoint is being called; `false` otherwise (for example, when the ACH transfer is scheduled and the end user is not present, or you call this endpoint after the ACH transfer but before submitting the Nacha file for ACH processing).
	UserPresent NullableBool `json:"user_present,omitempty"`
	// A unique ID that identifies the end user in your system. This ID is used to correlate requests by a user with multiple Items. The max length for this field is 36 characters. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.
	ClientUserId *string `json:"client_user_id,omitempty"`
	// **true** if the ACH transaction is a recurring transaction; **false** otherwise 
	IsRecurring NullableBool `json:"is_recurring,omitempty"`
	// The default ACH or non-ACH payment method to complete the transaction. `SAME_DAY_ACH`: Same Day ACH by NACHA. The debit transaction is processed and settled on the same day `NEXT_DAY_ACH`: Next Day ACH settlement for debit transactions, offered by some payment processors `STANDARD_ACH`: standard ACH by NACHA `REAL_TIME_PAYMENTS`: real-time payments such as RTP and FedNow `DEBIT_CARD`: if the default payment is over debit card networks `MULTIPLE_PAYMENT_METHODS`: if there is no default debit rail or there are multiple payment methods Possible values:  `SAME_DAY_ACH`, `NEXT_DAY_ACH`, `STANDARD_ACH`, `REAL_TIME_PAYMENTS`, `DEBIT_CARD`, `MULTIPLE_PAYMENT_METHODS`
	DefaultPaymentMethod NullableString `json:"default_payment_method,omitempty"`
	User *SignalUser `json:"user,omitempty"`
	Device *SignalDevice `json:"device,omitempty"`
}

// NewProcessorSignalEvaluateRequest instantiates a new ProcessorSignalEvaluateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorSignalEvaluateRequest(processorToken string, clientTransactionId string, amount float64) *ProcessorSignalEvaluateRequest {
	this := ProcessorSignalEvaluateRequest{}
	this.ProcessorToken = processorToken
	this.ClientTransactionId = clientTransactionId
	this.Amount = amount
	return &this
}

// NewProcessorSignalEvaluateRequestWithDefaults instantiates a new ProcessorSignalEvaluateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorSignalEvaluateRequestWithDefaults() *ProcessorSignalEvaluateRequest {
	this := ProcessorSignalEvaluateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ProcessorSignalEvaluateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorSignalEvaluateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ProcessorSignalEvaluateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ProcessorSignalEvaluateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ProcessorSignalEvaluateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorSignalEvaluateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ProcessorSignalEvaluateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ProcessorSignalEvaluateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetProcessorToken returns the ProcessorToken field value
func (o *ProcessorSignalEvaluateRequest) GetProcessorToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessorToken
}

// GetProcessorTokenOk returns a tuple with the ProcessorToken field value
// and a boolean to check if the value has been set.
func (o *ProcessorSignalEvaluateRequest) GetProcessorTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProcessorToken, true
}

// SetProcessorToken sets field value
func (o *ProcessorSignalEvaluateRequest) SetProcessorToken(v string) {
	o.ProcessorToken = v
}

// GetClientTransactionId returns the ClientTransactionId field value
func (o *ProcessorSignalEvaluateRequest) GetClientTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientTransactionId
}

// GetClientTransactionIdOk returns a tuple with the ClientTransactionId field value
// and a boolean to check if the value has been set.
func (o *ProcessorSignalEvaluateRequest) GetClientTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientTransactionId, true
}

// SetClientTransactionId sets field value
func (o *ProcessorSignalEvaluateRequest) SetClientTransactionId(v string) {
	o.ClientTransactionId = v
}

// GetAmount returns the Amount field value
func (o *ProcessorSignalEvaluateRequest) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ProcessorSignalEvaluateRequest) GetAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ProcessorSignalEvaluateRequest) SetAmount(v float64) {
	o.Amount = v
}

// GetUserPresent returns the UserPresent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessorSignalEvaluateRequest) GetUserPresent() bool {
	if o == nil || o.UserPresent.Get() == nil {
		var ret bool
		return ret
	}
	return *o.UserPresent.Get()
}

// GetUserPresentOk returns a tuple with the UserPresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessorSignalEvaluateRequest) GetUserPresentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserPresent.Get(), o.UserPresent.IsSet()
}

// HasUserPresent returns a boolean if a field has been set.
func (o *ProcessorSignalEvaluateRequest) HasUserPresent() bool {
	if o != nil && o.UserPresent.IsSet() {
		return true
	}

	return false
}

// SetUserPresent gets a reference to the given NullableBool and assigns it to the UserPresent field.
func (o *ProcessorSignalEvaluateRequest) SetUserPresent(v bool) {
	o.UserPresent.Set(&v)
}
// SetUserPresentNil sets the value for UserPresent to be an explicit nil
func (o *ProcessorSignalEvaluateRequest) SetUserPresentNil() {
	o.UserPresent.Set(nil)
}

// UnsetUserPresent ensures that no value is present for UserPresent, not even an explicit nil
func (o *ProcessorSignalEvaluateRequest) UnsetUserPresent() {
	o.UserPresent.Unset()
}

// GetClientUserId returns the ClientUserId field value if set, zero value otherwise.
func (o *ProcessorSignalEvaluateRequest) GetClientUserId() string {
	if o == nil || o.ClientUserId == nil {
		var ret string
		return ret
	}
	return *o.ClientUserId
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorSignalEvaluateRequest) GetClientUserIdOk() (*string, bool) {
	if o == nil || o.ClientUserId == nil {
		return nil, false
	}
	return o.ClientUserId, true
}

// HasClientUserId returns a boolean if a field has been set.
func (o *ProcessorSignalEvaluateRequest) HasClientUserId() bool {
	if o != nil && o.ClientUserId != nil {
		return true
	}

	return false
}

// SetClientUserId gets a reference to the given string and assigns it to the ClientUserId field.
func (o *ProcessorSignalEvaluateRequest) SetClientUserId(v string) {
	o.ClientUserId = &v
}

// GetIsRecurring returns the IsRecurring field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessorSignalEvaluateRequest) GetIsRecurring() bool {
	if o == nil || o.IsRecurring.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsRecurring.Get()
}

// GetIsRecurringOk returns a tuple with the IsRecurring field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessorSignalEvaluateRequest) GetIsRecurringOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsRecurring.Get(), o.IsRecurring.IsSet()
}

// HasIsRecurring returns a boolean if a field has been set.
func (o *ProcessorSignalEvaluateRequest) HasIsRecurring() bool {
	if o != nil && o.IsRecurring.IsSet() {
		return true
	}

	return false
}

// SetIsRecurring gets a reference to the given NullableBool and assigns it to the IsRecurring field.
func (o *ProcessorSignalEvaluateRequest) SetIsRecurring(v bool) {
	o.IsRecurring.Set(&v)
}
// SetIsRecurringNil sets the value for IsRecurring to be an explicit nil
func (o *ProcessorSignalEvaluateRequest) SetIsRecurringNil() {
	o.IsRecurring.Set(nil)
}

// UnsetIsRecurring ensures that no value is present for IsRecurring, not even an explicit nil
func (o *ProcessorSignalEvaluateRequest) UnsetIsRecurring() {
	o.IsRecurring.Unset()
}

// GetDefaultPaymentMethod returns the DefaultPaymentMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessorSignalEvaluateRequest) GetDefaultPaymentMethod() string {
	if o == nil || o.DefaultPaymentMethod.Get() == nil {
		var ret string
		return ret
	}
	return *o.DefaultPaymentMethod.Get()
}

// GetDefaultPaymentMethodOk returns a tuple with the DefaultPaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessorSignalEvaluateRequest) GetDefaultPaymentMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DefaultPaymentMethod.Get(), o.DefaultPaymentMethod.IsSet()
}

// HasDefaultPaymentMethod returns a boolean if a field has been set.
func (o *ProcessorSignalEvaluateRequest) HasDefaultPaymentMethod() bool {
	if o != nil && o.DefaultPaymentMethod.IsSet() {
		return true
	}

	return false
}

// SetDefaultPaymentMethod gets a reference to the given NullableString and assigns it to the DefaultPaymentMethod field.
func (o *ProcessorSignalEvaluateRequest) SetDefaultPaymentMethod(v string) {
	o.DefaultPaymentMethod.Set(&v)
}
// SetDefaultPaymentMethodNil sets the value for DefaultPaymentMethod to be an explicit nil
func (o *ProcessorSignalEvaluateRequest) SetDefaultPaymentMethodNil() {
	o.DefaultPaymentMethod.Set(nil)
}

// UnsetDefaultPaymentMethod ensures that no value is present for DefaultPaymentMethod, not even an explicit nil
func (o *ProcessorSignalEvaluateRequest) UnsetDefaultPaymentMethod() {
	o.DefaultPaymentMethod.Unset()
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ProcessorSignalEvaluateRequest) GetUser() SignalUser {
	if o == nil || o.User == nil {
		var ret SignalUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorSignalEvaluateRequest) GetUserOk() (*SignalUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ProcessorSignalEvaluateRequest) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given SignalUser and assigns it to the User field.
func (o *ProcessorSignalEvaluateRequest) SetUser(v SignalUser) {
	o.User = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *ProcessorSignalEvaluateRequest) GetDevice() SignalDevice {
	if o == nil || o.Device == nil {
		var ret SignalDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorSignalEvaluateRequest) GetDeviceOk() (*SignalDevice, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *ProcessorSignalEvaluateRequest) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given SignalDevice and assigns it to the Device field.
func (o *ProcessorSignalEvaluateRequest) SetDevice(v SignalDevice) {
	o.Device = &v
}

func (o ProcessorSignalEvaluateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["processor_token"] = o.ProcessorToken
	}
	if true {
		toSerialize["client_transaction_id"] = o.ClientTransactionId
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.UserPresent.IsSet() {
		toSerialize["user_present"] = o.UserPresent.Get()
	}
	if o.ClientUserId != nil {
		toSerialize["client_user_id"] = o.ClientUserId
	}
	if o.IsRecurring.IsSet() {
		toSerialize["is_recurring"] = o.IsRecurring.Get()
	}
	if o.DefaultPaymentMethod.IsSet() {
		toSerialize["default_payment_method"] = o.DefaultPaymentMethod.Get()
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	return json.Marshal(toSerialize)
}

type NullableProcessorSignalEvaluateRequest struct {
	value *ProcessorSignalEvaluateRequest
	isSet bool
}

func (v NullableProcessorSignalEvaluateRequest) Get() *ProcessorSignalEvaluateRequest {
	return v.value
}

func (v *NullableProcessorSignalEvaluateRequest) Set(val *ProcessorSignalEvaluateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorSignalEvaluateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorSignalEvaluateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorSignalEvaluateRequest(val *ProcessorSignalEvaluateRequest) *NullableProcessorSignalEvaluateRequest {
	return &NullableProcessorSignalEvaluateRequest{value: val, isSet: true}
}

func (v NullableProcessorSignalEvaluateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorSignalEvaluateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


