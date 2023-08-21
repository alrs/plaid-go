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

// AssetsErrorWebhook Fired when Asset Report generation has failed. The resulting `error` will have an `error_type` of `ASSET_REPORT_ERROR`.
type AssetsErrorWebhook struct {
	// `ASSETS`
	WebhookType string `json:"webhook_type"`
	// `ERROR`
	WebhookCode string `json:"webhook_code"`
	Error NullablePlaidError `json:"error"`
	// The ID associated with the Asset Report.
	AssetReportId string `json:"asset_report_id"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _AssetsErrorWebhook AssetsErrorWebhook

// NewAssetsErrorWebhook instantiates a new AssetsErrorWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsErrorWebhook(webhookType string, webhookCode string, error_ NullablePlaidError, assetReportId string, environment WebhookEnvironmentValues) *AssetsErrorWebhook {
	this := AssetsErrorWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.Error = error_
	this.AssetReportId = assetReportId
	this.Environment = environment
	return &this
}

// NewAssetsErrorWebhookWithDefaults instantiates a new AssetsErrorWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsErrorWebhookWithDefaults() *AssetsErrorWebhook {
	this := AssetsErrorWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *AssetsErrorWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *AssetsErrorWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *AssetsErrorWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *AssetsErrorWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *AssetsErrorWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *AssetsErrorWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetError returns the Error field value
// If the value is explicit nil, the zero value for PlaidError will be returned
func (o *AssetsErrorWebhook) GetError() PlaidError {
	if o == nil || o.Error.Get() == nil {
		var ret PlaidError
		return ret
	}

	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetsErrorWebhook) GetErrorOk() (*PlaidError, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// SetError sets field value
func (o *AssetsErrorWebhook) SetError(v PlaidError) {
	o.Error.Set(&v)
}

// GetAssetReportId returns the AssetReportId field value
func (o *AssetsErrorWebhook) GetAssetReportId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetReportId
}

// GetAssetReportIdOk returns a tuple with the AssetReportId field value
// and a boolean to check if the value has been set.
func (o *AssetsErrorWebhook) GetAssetReportIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetReportId, true
}

// SetAssetReportId sets field value
func (o *AssetsErrorWebhook) SetAssetReportId(v string) {
	o.AssetReportId = v
}

// GetEnvironment returns the Environment field value
func (o *AssetsErrorWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *AssetsErrorWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *AssetsErrorWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o AssetsErrorWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["error"] = o.Error.Get()
	}
	if true {
		toSerialize["asset_report_id"] = o.AssetReportId
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetsErrorWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varAssetsErrorWebhook := _AssetsErrorWebhook{}

	if err = json.Unmarshal(bytes, &varAssetsErrorWebhook); err == nil {
		*o = AssetsErrorWebhook(varAssetsErrorWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "error")
		delete(additionalProperties, "asset_report_id")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetsErrorWebhook struct {
	value *AssetsErrorWebhook
	isSet bool
}

func (v NullableAssetsErrorWebhook) Get() *AssetsErrorWebhook {
	return v.value
}

func (v *NullableAssetsErrorWebhook) Set(val *AssetsErrorWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsErrorWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsErrorWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsErrorWebhook(val *AssetsErrorWebhook) *NullableAssetsErrorWebhook {
	return &NullableAssetsErrorWebhook{value: val, isSet: true}
}

func (v NullableAssetsErrorWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsErrorWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


