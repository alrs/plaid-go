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

// DefaultUpdateWebhook Fired when new transaction data is available for an Item. Plaid will typically check for new transaction data several times a day. 
type DefaultUpdateWebhook struct {
	// `TRANSACTIONS`
	WebhookType string `json:"webhook_type"`
	// `DEFAULT_UPDATE`
	WebhookCode string `json:"webhook_code"`
	Error *PlaidError `json:"error,omitempty"`
	// The number of new transactions detected since the last time this webhook was fired.
	NewTransactions float32 `json:"new_transactions"`
	// The `item_id` of the Item the webhook relates to.
	ItemId string `json:"item_id"`
	AdditionalProperties map[string]interface{}
}

type _DefaultUpdateWebhook DefaultUpdateWebhook

// NewDefaultUpdateWebhook instantiates a new DefaultUpdateWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultUpdateWebhook(webhookType string, webhookCode string, newTransactions float32, itemId string) *DefaultUpdateWebhook {
	this := DefaultUpdateWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.NewTransactions = newTransactions
	this.ItemId = itemId
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

// GetError returns the Error field value if set, zero value otherwise.
func (o *DefaultUpdateWebhook) GetError() PlaidError {
	if o == nil || o.Error == nil {
		var ret PlaidError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultUpdateWebhook) GetErrorOk() (*PlaidError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *DefaultUpdateWebhook) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given PlaidError and assigns it to the Error field.
func (o *DefaultUpdateWebhook) SetError(v PlaidError) {
	o.Error = &v
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

func (o DefaultUpdateWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if true {
		toSerialize["new_transactions"] = o.NewTransactions
	}
	if true {
		toSerialize["item_id"] = o.ItemId
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
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "error")
		delete(additionalProperties, "new_transactions")
		delete(additionalProperties, "item_id")
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


