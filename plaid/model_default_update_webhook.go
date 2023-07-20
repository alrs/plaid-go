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

// DefaultUpdateWebhook Fired when new transaction data is available for an Item. Plaid will typically check for new transaction data several times a day.  This webhook is intended for use with `/transactions/get`; if you are using the newer `/transactions/sync` endpoint, this webhook will still be fired to maintain backwards compatibility, but it is recommended to listen for and respond to the `SYNC_UPDATES_AVAILABLE` webhook instead. 
type DefaultUpdateWebhook struct {
	// `TRANSACTIONS`
	WebhookType string `json:"webhook_type"`
	// `DEFAULT_UPDATE`
	WebhookCode string `json:"webhook_code"`
	Error NullablePlaidError `json:"error,omitempty"`
	// The number of new transactions detected since the last time this webhook was fired.
	NewTransactions float32 `json:"new_transactions"`
	// The `item_id` of the Item the webhook relates to.
	ItemId string `json:"item_id"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _DefaultUpdateWebhook DefaultUpdateWebhook

// NewDefaultUpdateWebhook instantiates a new DefaultUpdateWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultUpdateWebhook(webhookType string, webhookCode string, newTransactions float32, itemId string, environment WebhookEnvironmentValues) *DefaultUpdateWebhook {
	this := DefaultUpdateWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.NewTransactions = newTransactions
	this.ItemId = itemId
	this.Environment = environment
	return &this
}

// NewDefaultUpdateWebhookWithDefaults instantiates a new DefaultUpdateWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultUpdateWebhookWithDefaults() *DefaultUpdateWebhook {
	this := DefaultUpdateWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *DefaultUpdateWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *DefaultUpdateWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *DefaultUpdateWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *DefaultUpdateWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *DefaultUpdateWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *DefaultUpdateWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultUpdateWebhook) GetError() PlaidError {
	if o == nil || o.Error.Get() == nil {
		var ret PlaidError
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultUpdateWebhook) GetErrorOk() (*PlaidError, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *DefaultUpdateWebhook) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullablePlaidError and assigns it to the Error field.
func (o *DefaultUpdateWebhook) SetError(v PlaidError) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *DefaultUpdateWebhook) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *DefaultUpdateWebhook) UnsetError() {
	o.Error.Unset()
}

// GetNewTransactions returns the NewTransactions field value
func (o *DefaultUpdateWebhook) GetNewTransactions() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NewTransactions
}

// GetNewTransactionsOk returns a tuple with the NewTransactions field value
// and a boolean to check if the value has been set.
func (o *DefaultUpdateWebhook) GetNewTransactionsOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NewTransactions, true
}

// SetNewTransactions sets field value
func (o *DefaultUpdateWebhook) SetNewTransactions(v float32) {
	o.NewTransactions = v
}

// GetItemId returns the ItemId field value
func (o *DefaultUpdateWebhook) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *DefaultUpdateWebhook) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *DefaultUpdateWebhook) SetItemId(v string) {
	o.ItemId = v
}

// GetEnvironment returns the Environment field value
func (o *DefaultUpdateWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *DefaultUpdateWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *DefaultUpdateWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o DefaultUpdateWebhook) MarshalJSON() ([]byte, error) {
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
		toSerialize["item_id"] = o.ItemId
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DefaultUpdateWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varDefaultUpdateWebhook := _DefaultUpdateWebhook{}

	if err = json.Unmarshal(bytes, &varDefaultUpdateWebhook); err == nil {
		*o = DefaultUpdateWebhook(varDefaultUpdateWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "error")
		delete(additionalProperties, "new_transactions")
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDefaultUpdateWebhook struct {
	value *DefaultUpdateWebhook
	isSet bool
}

func (v NullableDefaultUpdateWebhook) Get() *DefaultUpdateWebhook {
	return v.value
}

func (v *NullableDefaultUpdateWebhook) Set(val *DefaultUpdateWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultUpdateWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultUpdateWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultUpdateWebhook(val *DefaultUpdateWebhook) *NullableDefaultUpdateWebhook {
	return &NullableDefaultUpdateWebhook{value: val, isSet: true}
}

func (v NullableDefaultUpdateWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultUpdateWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


