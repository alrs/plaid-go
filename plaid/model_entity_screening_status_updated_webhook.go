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

// EntityScreeningStatusUpdatedWebhook Fired when an entity screening status has changed, which can occur manually via the dashboard or during ongoing monitoring.
type EntityScreeningStatusUpdatedWebhook struct {
	// `ENTITY_SCREENING`
	WebhookType string `json:"webhook_type"`
	// `STATUS_UPDATED`
	WebhookCode string `json:"webhook_code"`
	// The ID of the associated screening.
	ScreeningId string `json:"screening_id"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _EntityScreeningStatusUpdatedWebhook EntityScreeningStatusUpdatedWebhook

// NewEntityScreeningStatusUpdatedWebhook instantiates a new EntityScreeningStatusUpdatedWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityScreeningStatusUpdatedWebhook(webhookType string, webhookCode string, screeningId string, environment WebhookEnvironmentValues) *EntityScreeningStatusUpdatedWebhook {
	this := EntityScreeningStatusUpdatedWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.ScreeningId = screeningId
	this.Environment = environment
	return &this
}

// NewEntityScreeningStatusUpdatedWebhookWithDefaults instantiates a new EntityScreeningStatusUpdatedWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityScreeningStatusUpdatedWebhookWithDefaults() *EntityScreeningStatusUpdatedWebhook {
	this := EntityScreeningStatusUpdatedWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *EntityScreeningStatusUpdatedWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *EntityScreeningStatusUpdatedWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *EntityScreeningStatusUpdatedWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *EntityScreeningStatusUpdatedWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *EntityScreeningStatusUpdatedWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *EntityScreeningStatusUpdatedWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetScreeningId returns the ScreeningId field value
func (o *EntityScreeningStatusUpdatedWebhook) GetScreeningId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScreeningId
}

// GetScreeningIdOk returns a tuple with the ScreeningId field value
// and a boolean to check if the value has been set.
func (o *EntityScreeningStatusUpdatedWebhook) GetScreeningIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScreeningId, true
}

// SetScreeningId sets field value
func (o *EntityScreeningStatusUpdatedWebhook) SetScreeningId(v string) {
	o.ScreeningId = v
}

// GetEnvironment returns the Environment field value
func (o *EntityScreeningStatusUpdatedWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *EntityScreeningStatusUpdatedWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *EntityScreeningStatusUpdatedWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o EntityScreeningStatusUpdatedWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["screening_id"] = o.ScreeningId
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntityScreeningStatusUpdatedWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varEntityScreeningStatusUpdatedWebhook := _EntityScreeningStatusUpdatedWebhook{}

	if err = json.Unmarshal(bytes, &varEntityScreeningStatusUpdatedWebhook); err == nil {
		*o = EntityScreeningStatusUpdatedWebhook(varEntityScreeningStatusUpdatedWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "screening_id")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntityScreeningStatusUpdatedWebhook struct {
	value *EntityScreeningStatusUpdatedWebhook
	isSet bool
}

func (v NullableEntityScreeningStatusUpdatedWebhook) Get() *EntityScreeningStatusUpdatedWebhook {
	return v.value
}

func (v *NullableEntityScreeningStatusUpdatedWebhook) Set(val *EntityScreeningStatusUpdatedWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityScreeningStatusUpdatedWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityScreeningStatusUpdatedWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityScreeningStatusUpdatedWebhook(val *EntityScreeningStatusUpdatedWebhook) *NullableEntityScreeningStatusUpdatedWebhook {
	return &NullableEntityScreeningStatusUpdatedWebhook{value: val, isSet: true}
}

func (v NullableEntityScreeningStatusUpdatedWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityScreeningStatusUpdatedWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


