/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.379.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// HoldingsDefaultUpdateWebhook Fired when new or updated holdings have been detected on an investment account. The webhook typically fires in response to any newly added holdings or price changes to existing holdings, most commonly after market close.
type HoldingsDefaultUpdateWebhook struct {
	// `HOLDINGS`
	WebhookType string `json:"webhook_type"`
	// `DEFAULT_UPDATE`
	WebhookCode string `json:"webhook_code"`
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	Error NullablePlaidError `json:"error,omitempty"`
	// The number of new holdings reported since the last time this webhook was fired.
	NewHoldings float32 `json:"new_holdings"`
	// The number of updated holdings reported since the last time this webhook was fired.
	UpdatedHoldings float32 `json:"updated_holdings"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _HoldingsDefaultUpdateWebhook HoldingsDefaultUpdateWebhook

// NewHoldingsDefaultUpdateWebhook instantiates a new HoldingsDefaultUpdateWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHoldingsDefaultUpdateWebhook(webhookType string, webhookCode string, itemId string, newHoldings float32, updatedHoldings float32, environment WebhookEnvironmentValues) *HoldingsDefaultUpdateWebhook {
	this := HoldingsDefaultUpdateWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.ItemId = itemId
	this.NewHoldings = newHoldings
	this.UpdatedHoldings = updatedHoldings
	this.Environment = environment
	return &this
}

// NewHoldingsDefaultUpdateWebhookWithDefaults instantiates a new HoldingsDefaultUpdateWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHoldingsDefaultUpdateWebhookWithDefaults() *HoldingsDefaultUpdateWebhook {
	this := HoldingsDefaultUpdateWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *HoldingsDefaultUpdateWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *HoldingsDefaultUpdateWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *HoldingsDefaultUpdateWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *HoldingsDefaultUpdateWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *HoldingsDefaultUpdateWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *HoldingsDefaultUpdateWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetItemId returns the ItemId field value
func (o *HoldingsDefaultUpdateWebhook) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *HoldingsDefaultUpdateWebhook) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *HoldingsDefaultUpdateWebhook) SetItemId(v string) {
	o.ItemId = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HoldingsDefaultUpdateWebhook) GetError() PlaidError {
	if o == nil || o.Error.Get() == nil {
		var ret PlaidError
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HoldingsDefaultUpdateWebhook) GetErrorOk() (*PlaidError, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *HoldingsDefaultUpdateWebhook) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullablePlaidError and assigns it to the Error field.
func (o *HoldingsDefaultUpdateWebhook) SetError(v PlaidError) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *HoldingsDefaultUpdateWebhook) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *HoldingsDefaultUpdateWebhook) UnsetError() {
	o.Error.Unset()
}

// GetNewHoldings returns the NewHoldings field value
func (o *HoldingsDefaultUpdateWebhook) GetNewHoldings() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NewHoldings
}

// GetNewHoldingsOk returns a tuple with the NewHoldings field value
// and a boolean to check if the value has been set.
func (o *HoldingsDefaultUpdateWebhook) GetNewHoldingsOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NewHoldings, true
}

// SetNewHoldings sets field value
func (o *HoldingsDefaultUpdateWebhook) SetNewHoldings(v float32) {
	o.NewHoldings = v
}

// GetUpdatedHoldings returns the UpdatedHoldings field value
func (o *HoldingsDefaultUpdateWebhook) GetUpdatedHoldings() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UpdatedHoldings
}

// GetUpdatedHoldingsOk returns a tuple with the UpdatedHoldings field value
// and a boolean to check if the value has been set.
func (o *HoldingsDefaultUpdateWebhook) GetUpdatedHoldingsOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdatedHoldings, true
}

// SetUpdatedHoldings sets field value
func (o *HoldingsDefaultUpdateWebhook) SetUpdatedHoldings(v float32) {
	o.UpdatedHoldings = v
}

// GetEnvironment returns the Environment field value
func (o *HoldingsDefaultUpdateWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *HoldingsDefaultUpdateWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *HoldingsDefaultUpdateWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o HoldingsDefaultUpdateWebhook) MarshalJSON() ([]byte, error) {
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
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if true {
		toSerialize["new_holdings"] = o.NewHoldings
	}
	if true {
		toSerialize["updated_holdings"] = o.UpdatedHoldings
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HoldingsDefaultUpdateWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varHoldingsDefaultUpdateWebhook := _HoldingsDefaultUpdateWebhook{}

	if err = json.Unmarshal(bytes, &varHoldingsDefaultUpdateWebhook); err == nil {
		*o = HoldingsDefaultUpdateWebhook(varHoldingsDefaultUpdateWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "error")
		delete(additionalProperties, "new_holdings")
		delete(additionalProperties, "updated_holdings")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHoldingsDefaultUpdateWebhook struct {
	value *HoldingsDefaultUpdateWebhook
	isSet bool
}

func (v NullableHoldingsDefaultUpdateWebhook) Get() *HoldingsDefaultUpdateWebhook {
	return v.value
}

func (v *NullableHoldingsDefaultUpdateWebhook) Set(val *HoldingsDefaultUpdateWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableHoldingsDefaultUpdateWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableHoldingsDefaultUpdateWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHoldingsDefaultUpdateWebhook(val *HoldingsDefaultUpdateWebhook) *NullableHoldingsDefaultUpdateWebhook {
	return &NullableHoldingsDefaultUpdateWebhook{value: val, isSet: true}
}

func (v NullableHoldingsDefaultUpdateWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHoldingsDefaultUpdateWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


