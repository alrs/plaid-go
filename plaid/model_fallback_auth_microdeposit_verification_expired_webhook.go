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

// FallbackAuthMicrodepositVerificationExpiredWebhook Fires when an account has an expired verification when using micro-deposits
type FallbackAuthMicrodepositVerificationExpiredWebhook struct {
	// `AUTH`
	WebhookType string `json:"webhook_type"`
	// `VERIFICATION_EXPIRED`
	WebhookCode string `json:"webhook_code"`
	// The error code associated with the webhook.
	Error NullableString `json:"error,omitempty"`
	// The external account ID associated with the micro-deposit
	AccountId string `json:"account_id"`
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _FallbackAuthMicrodepositVerificationExpiredWebhook FallbackAuthMicrodepositVerificationExpiredWebhook

// NewFallbackAuthMicrodepositVerificationExpiredWebhook instantiates a new FallbackAuthMicrodepositVerificationExpiredWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFallbackAuthMicrodepositVerificationExpiredWebhook(webhookType string, webhookCode string, accountId string, itemId string, environment WebhookEnvironmentValues) *FallbackAuthMicrodepositVerificationExpiredWebhook {
	this := FallbackAuthMicrodepositVerificationExpiredWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.AccountId = accountId
	this.ItemId = itemId
	this.Environment = environment
	return &this
}

// NewFallbackAuthMicrodepositVerificationExpiredWebhookWithDefaults instantiates a new FallbackAuthMicrodepositVerificationExpiredWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFallbackAuthMicrodepositVerificationExpiredWebhookWithDefaults() *FallbackAuthMicrodepositVerificationExpiredWebhook {
	this := FallbackAuthMicrodepositVerificationExpiredWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *FallbackAuthMicrodepositVerificationExpiredWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *FallbackAuthMicrodepositVerificationExpiredWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *FallbackAuthMicrodepositVerificationExpiredWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *FallbackAuthMicrodepositVerificationExpiredWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *FallbackAuthMicrodepositVerificationExpiredWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *FallbackAuthMicrodepositVerificationExpiredWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FallbackAuthMicrodepositVerificationExpiredWebhook) GetError() string {
	if o == nil || o.Error.Get() == nil {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FallbackAuthMicrodepositVerificationExpiredWebhook) GetErrorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *FallbackAuthMicrodepositVerificationExpiredWebhook) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *FallbackAuthMicrodepositVerificationExpiredWebhook) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *FallbackAuthMicrodepositVerificationExpiredWebhook) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *FallbackAuthMicrodepositVerificationExpiredWebhook) UnsetError() {
	o.Error.Unset()
}

// GetAccountId returns the AccountId field value
func (o *FallbackAuthMicrodepositVerificationExpiredWebhook) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *FallbackAuthMicrodepositVerificationExpiredWebhook) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *FallbackAuthMicrodepositVerificationExpiredWebhook) SetAccountId(v string) {
	o.AccountId = v
}

// GetItemId returns the ItemId field value
func (o *FallbackAuthMicrodepositVerificationExpiredWebhook) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *FallbackAuthMicrodepositVerificationExpiredWebhook) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *FallbackAuthMicrodepositVerificationExpiredWebhook) SetItemId(v string) {
	o.ItemId = v
}

// GetEnvironment returns the Environment field value
func (o *FallbackAuthMicrodepositVerificationExpiredWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *FallbackAuthMicrodepositVerificationExpiredWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *FallbackAuthMicrodepositVerificationExpiredWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o FallbackAuthMicrodepositVerificationExpiredWebhook) MarshalJSON() ([]byte, error) {
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
		toSerialize["account_id"] = o.AccountId
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

func (o *FallbackAuthMicrodepositVerificationExpiredWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varFallbackAuthMicrodepositVerificationExpiredWebhook := _FallbackAuthMicrodepositVerificationExpiredWebhook{}

	if err = json.Unmarshal(bytes, &varFallbackAuthMicrodepositVerificationExpiredWebhook); err == nil {
		*o = FallbackAuthMicrodepositVerificationExpiredWebhook(varFallbackAuthMicrodepositVerificationExpiredWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "error")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFallbackAuthMicrodepositVerificationExpiredWebhook struct {
	value *FallbackAuthMicrodepositVerificationExpiredWebhook
	isSet bool
}

func (v NullableFallbackAuthMicrodepositVerificationExpiredWebhook) Get() *FallbackAuthMicrodepositVerificationExpiredWebhook {
	return v.value
}

func (v *NullableFallbackAuthMicrodepositVerificationExpiredWebhook) Set(val *FallbackAuthMicrodepositVerificationExpiredWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableFallbackAuthMicrodepositVerificationExpiredWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableFallbackAuthMicrodepositVerificationExpiredWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFallbackAuthMicrodepositVerificationExpiredWebhook(val *FallbackAuthMicrodepositVerificationExpiredWebhook) *NullableFallbackAuthMicrodepositVerificationExpiredWebhook {
	return &NullableFallbackAuthMicrodepositVerificationExpiredWebhook{value: val, isSet: true}
}

func (v NullableFallbackAuthMicrodepositVerificationExpiredWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFallbackAuthMicrodepositVerificationExpiredWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


