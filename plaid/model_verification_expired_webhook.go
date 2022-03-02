/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.79.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// VerificationExpiredWebhook Fired when an Item was not verified via automated micro-deposits after seven days since the automated micro-deposit was made.
type VerificationExpiredWebhook struct {
	// `AUTH`
	WebhookType string `json:"webhook_type"`
	// `VERIFICATION_EXPIRED`
	WebhookCode string `json:"webhook_code"`
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	// The `account_id` of the account associated with the webhook
	AccountId string `json:"account_id"`
	AdditionalProperties map[string]interface{}
}

type _VerificationExpiredWebhook VerificationExpiredWebhook

// NewVerificationExpiredWebhook instantiates a new VerificationExpiredWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationExpiredWebhook(webhookType string, webhookCode string, itemId string, accountId string) *VerificationExpiredWebhook {
	this := VerificationExpiredWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.ItemId = itemId
	this.AccountId = accountId
	return &this
}

// NewVerificationExpiredWebhookWithDefaults instantiates a new VerificationExpiredWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationExpiredWebhookWithDefaults() *VerificationExpiredWebhook {
	this := VerificationExpiredWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *VerificationExpiredWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *VerificationExpiredWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *VerificationExpiredWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *VerificationExpiredWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *VerificationExpiredWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *VerificationExpiredWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetItemId returns the ItemId field value
func (o *VerificationExpiredWebhook) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *VerificationExpiredWebhook) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *VerificationExpiredWebhook) SetItemId(v string) {
	o.ItemId = v
}

// GetAccountId returns the AccountId field value
func (o *VerificationExpiredWebhook) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *VerificationExpiredWebhook) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *VerificationExpiredWebhook) SetAccountId(v string) {
	o.AccountId = v
}

func (o VerificationExpiredWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["item_id"] = o.ItemId
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VerificationExpiredWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varVerificationExpiredWebhook := _VerificationExpiredWebhook{}

	if err = json.Unmarshal(bytes, &varVerificationExpiredWebhook); err == nil {
		*o = VerificationExpiredWebhook(varVerificationExpiredWebhook)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "account_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVerificationExpiredWebhook struct {
	value *VerificationExpiredWebhook
	isSet bool
}

func (v NullableVerificationExpiredWebhook) Get() *VerificationExpiredWebhook {
	return v.value
}

func (v *NullableVerificationExpiredWebhook) Set(val *VerificationExpiredWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationExpiredWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationExpiredWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationExpiredWebhook(val *VerificationExpiredWebhook) *NullableVerificationExpiredWebhook {
	return &NullableVerificationExpiredWebhook{value: val, isSet: true}
}

func (v NullableVerificationExpiredWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationExpiredWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


