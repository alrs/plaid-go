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

// AutomaticallyVerifiedWebhook Fired when an Item is verified via automated micro-deposits. We recommend communicating to your users when this event is received to notify them that their account is verified and ready for use.
type AutomaticallyVerifiedWebhook struct {
	// `AUTH`
	WebhookType string `json:"webhook_type"`
	// `AUTOMATICALLY_VERIFIED`
	WebhookCode string `json:"webhook_code"`
	// The `account_id` of the account associated with the webhook
	AccountId string `json:"account_id"`
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	AdditionalProperties map[string]interface{}
}

type _AutomaticallyVerifiedWebhook AutomaticallyVerifiedWebhook

// NewAutomaticallyVerifiedWebhook instantiates a new AutomaticallyVerifiedWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomaticallyVerifiedWebhook(webhookType string, webhookCode string, accountId string, itemId string) *AutomaticallyVerifiedWebhook {
	this := AutomaticallyVerifiedWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.AccountId = accountId
	this.ItemId = itemId
	return &this
}

// NewAutomaticallyVerifiedWebhookWithDefaults instantiates a new AutomaticallyVerifiedWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomaticallyVerifiedWebhookWithDefaults() *AutomaticallyVerifiedWebhook {
	this := AutomaticallyVerifiedWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *AutomaticallyVerifiedWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *AutomaticallyVerifiedWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *AutomaticallyVerifiedWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *AutomaticallyVerifiedWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *AutomaticallyVerifiedWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *AutomaticallyVerifiedWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetAccountId returns the AccountId field value
func (o *AutomaticallyVerifiedWebhook) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AutomaticallyVerifiedWebhook) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AutomaticallyVerifiedWebhook) SetAccountId(v string) {
	o.AccountId = v
}

// GetItemId returns the ItemId field value
func (o *AutomaticallyVerifiedWebhook) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *AutomaticallyVerifiedWebhook) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *AutomaticallyVerifiedWebhook) SetItemId(v string) {
	o.ItemId = v
}

func (o AutomaticallyVerifiedWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["item_id"] = o.ItemId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AutomaticallyVerifiedWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varAutomaticallyVerifiedWebhook := _AutomaticallyVerifiedWebhook{}

	if err = json.Unmarshal(bytes, &varAutomaticallyVerifiedWebhook); err == nil {
		*o = AutomaticallyVerifiedWebhook(varAutomaticallyVerifiedWebhook)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "item_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAutomaticallyVerifiedWebhook struct {
	value *AutomaticallyVerifiedWebhook
	isSet bool
}

func (v NullableAutomaticallyVerifiedWebhook) Get() *AutomaticallyVerifiedWebhook {
	return v.value
}

func (v *NullableAutomaticallyVerifiedWebhook) Set(val *AutomaticallyVerifiedWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableAutomaticallyVerifiedWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableAutomaticallyVerifiedWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutomaticallyVerifiedWebhook(val *AutomaticallyVerifiedWebhook) *NullableAutomaticallyVerifiedWebhook {
	return &NullableAutomaticallyVerifiedWebhook{value: val, isSet: true}
}

func (v NullableAutomaticallyVerifiedWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutomaticallyVerifiedWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


