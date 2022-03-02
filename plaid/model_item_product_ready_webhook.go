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

// ItemProductReadyWebhook Fired once Plaid calculates income from an Item.
type ItemProductReadyWebhook struct {
	// `INCOME`
	WebhookType string `json:"webhook_type"`
	// `PRODUCT_READY`
	WebhookCode string `json:"webhook_code"`
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	Error *PlaidError `json:"error,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ItemProductReadyWebhook ItemProductReadyWebhook

// NewItemProductReadyWebhook instantiates a new ItemProductReadyWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemProductReadyWebhook(webhookType string, webhookCode string, itemId string) *ItemProductReadyWebhook {
	this := ItemProductReadyWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.ItemId = itemId
	return &this
}

// NewItemProductReadyWebhookWithDefaults instantiates a new ItemProductReadyWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemProductReadyWebhookWithDefaults() *ItemProductReadyWebhook {
	this := ItemProductReadyWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *ItemProductReadyWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *ItemProductReadyWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *ItemProductReadyWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *ItemProductReadyWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *ItemProductReadyWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *ItemProductReadyWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetItemId returns the ItemId field value
func (o *ItemProductReadyWebhook) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *ItemProductReadyWebhook) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *ItemProductReadyWebhook) SetItemId(v string) {
	o.ItemId = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ItemProductReadyWebhook) GetError() PlaidError {
	if o == nil || o.Error == nil {
		var ret PlaidError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemProductReadyWebhook) GetErrorOk() (*PlaidError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ItemProductReadyWebhook) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given PlaidError and assigns it to the Error field.
func (o *ItemProductReadyWebhook) SetError(v PlaidError) {
	o.Error = &v
}

func (o ItemProductReadyWebhook) MarshalJSON() ([]byte, error) {
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
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ItemProductReadyWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varItemProductReadyWebhook := _ItemProductReadyWebhook{}

	if err = json.Unmarshal(bytes, &varItemProductReadyWebhook); err == nil {
		*o = ItemProductReadyWebhook(varItemProductReadyWebhook)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableItemProductReadyWebhook struct {
	value *ItemProductReadyWebhook
	isSet bool
}

func (v NullableItemProductReadyWebhook) Get() *ItemProductReadyWebhook {
	return v.value
}

func (v *NullableItemProductReadyWebhook) Set(val *ItemProductReadyWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableItemProductReadyWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableItemProductReadyWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemProductReadyWebhook(val *ItemProductReadyWebhook) *NullableItemProductReadyWebhook {
	return &NullableItemProductReadyWebhook{value: val, isSet: true}
}

func (v NullableItemProductReadyWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemProductReadyWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


