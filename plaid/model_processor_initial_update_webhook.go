/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.419.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// ProcessorInitialUpdateWebhook This webhook is only sent to [Plaid processor partners](https://plaid.com/docs/auth/partnerships/).  Fired when an Item's initial transaction pull is completed. Once this webhook has been fired, transaction data for the most recent 30 days can be fetched for the Item. If [Account Select v2](https://plaid.com/docs/link/customization/#account-select) is enabled, this webhook will also be fired if account selections for the Item are updated, with `new_transactions` set to the number of net new transactions pulled after the account selection update.  This webhook is intended for use with `/processor/transactions/get`; if you are using the newer `/processor/transactions/sync` endpoint, this webhook will still be fired to maintain backwards compatibility, but it is recommended to listen for and respond to the `SYNC_UPDATES_AVAILABLE` webhook instead.
type ProcessorInitialUpdateWebhook struct {
	// `TRANSACTIONS`
	WebhookType string `json:"webhook_type"`
	// `INITIAL_UPDATE`
	WebhookCode string `json:"webhook_code"`
	// The error code associated with the webhook.
	Error NullableString `json:"error,omitempty"`
	// The number of new, unfetched transactions available.
	NewTransactions float32 `json:"new_transactions"`
	// The ID of the account.
	AccountId string `json:"account_id"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _ProcessorInitialUpdateWebhook ProcessorInitialUpdateWebhook

// NewProcessorInitialUpdateWebhook instantiates a new ProcessorInitialUpdateWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorInitialUpdateWebhook(webhookType string, webhookCode string, newTransactions float32, accountId string, environment WebhookEnvironmentValues) *ProcessorInitialUpdateWebhook {
	this := ProcessorInitialUpdateWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.NewTransactions = newTransactions
	this.AccountId = accountId
	this.Environment = environment
	return &this
}

// NewProcessorInitialUpdateWebhookWithDefaults instantiates a new ProcessorInitialUpdateWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorInitialUpdateWebhookWithDefaults() *ProcessorInitialUpdateWebhook {
	this := ProcessorInitialUpdateWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *ProcessorInitialUpdateWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *ProcessorInitialUpdateWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *ProcessorInitialUpdateWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *ProcessorInitialUpdateWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *ProcessorInitialUpdateWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *ProcessorInitialUpdateWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessorInitialUpdateWebhook) GetError() string {
	if o == nil || o.Error.Get() == nil {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessorInitialUpdateWebhook) GetErrorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *ProcessorInitialUpdateWebhook) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *ProcessorInitialUpdateWebhook) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *ProcessorInitialUpdateWebhook) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *ProcessorInitialUpdateWebhook) UnsetError() {
	o.Error.Unset()
}

// GetNewTransactions returns the NewTransactions field value
func (o *ProcessorInitialUpdateWebhook) GetNewTransactions() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NewTransactions
}

// GetNewTransactionsOk returns a tuple with the NewTransactions field value
// and a boolean to check if the value has been set.
func (o *ProcessorInitialUpdateWebhook) GetNewTransactionsOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NewTransactions, true
}

// SetNewTransactions sets field value
func (o *ProcessorInitialUpdateWebhook) SetNewTransactions(v float32) {
	o.NewTransactions = v
}

// GetAccountId returns the AccountId field value
func (o *ProcessorInitialUpdateWebhook) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *ProcessorInitialUpdateWebhook) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *ProcessorInitialUpdateWebhook) SetAccountId(v string) {
	o.AccountId = v
}

// GetEnvironment returns the Environment field value
func (o *ProcessorInitialUpdateWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *ProcessorInitialUpdateWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *ProcessorInitialUpdateWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o ProcessorInitialUpdateWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if true {
		toSerialize["new_transactions"] = o.NewTransactions
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProcessorInitialUpdateWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varProcessorInitialUpdateWebhook := _ProcessorInitialUpdateWebhook{}

	if err = json.Unmarshal(bytes, &varProcessorInitialUpdateWebhook); err == nil {
		*o = ProcessorInitialUpdateWebhook(varProcessorInitialUpdateWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "error")
		delete(additionalProperties, "new_transactions")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProcessorInitialUpdateWebhook struct {
	value *ProcessorInitialUpdateWebhook
	isSet bool
}

func (v NullableProcessorInitialUpdateWebhook) Get() *ProcessorInitialUpdateWebhook {
	return v.value
}

func (v *NullableProcessorInitialUpdateWebhook) Set(val *ProcessorInitialUpdateWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorInitialUpdateWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorInitialUpdateWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorInitialUpdateWebhook(val *ProcessorInitialUpdateWebhook) *NullableProcessorInitialUpdateWebhook {
	return &NullableProcessorInitialUpdateWebhook{value: val, isSet: true}
}

func (v NullableProcessorInitialUpdateWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorInitialUpdateWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


