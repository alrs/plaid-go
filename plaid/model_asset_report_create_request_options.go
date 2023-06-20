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

// AssetReportCreateRequestOptions An optional object to filter `/asset_report/create` results. If provided, must be non-`null`. The optional `user` object is required for the report to be eligible for Fannie Mae's Day 1 Certainty program.
type AssetReportCreateRequestOptions struct {
	// Client-generated identifier, which can be used by lenders to track loan applications.
	ClientReportId NullableString `json:"client_report_id,omitempty"`
	// URL to which Plaid will send Assets webhooks, for example when the requested Asset Report is ready.
	Webhook NullableString `json:"webhook,omitempty"`
	// true to return balance and identity earlier as a fast report. Defaults to false if omitted.
	IncludeFastReport NullableBool `json:"include_fast_report,omitempty"`
	// Additional information that can be included in the asset report. Possible values: `\"investments\"`
	Products *[]string `json:"products,omitempty"`
	// Use this field to request a `fast_asset` report. When Fast Assets is requested, Plaid will create two versions of the Asset Report: first, the Fast Asset Report, which will contain only current identity and balance information, and later, the Full Asset Report, which will also contain historical balance information and transaction data. A `PRODUCT_READY` webhook will be fired for each Asset Report when it is ready, and the `report_type` field will indicate whether the webhook is firing for the `full` or `fast` Asset Report. To retrieve the Fast Asset Report, call `/asset_report/get` with `fast_report` set to `true`. There is no additional charge for using Fast Assets.
	AddOns *[]AssetReportAddOns `json:"add_ons,omitempty"`
	User *AssetReportUser `json:"user,omitempty"`
}

// NewAssetReportCreateRequestOptions instantiates a new AssetReportCreateRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportCreateRequestOptions() *AssetReportCreateRequestOptions {
	this := AssetReportCreateRequestOptions{}
	return &this
}

// NewAssetReportCreateRequestOptionsWithDefaults instantiates a new AssetReportCreateRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportCreateRequestOptionsWithDefaults() *AssetReportCreateRequestOptions {
	this := AssetReportCreateRequestOptions{}
	return &this
}

// GetClientReportId returns the ClientReportId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetReportCreateRequestOptions) GetClientReportId() string {
	if o == nil || o.ClientReportId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientReportId.Get()
}

// GetClientReportIdOk returns a tuple with the ClientReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetReportCreateRequestOptions) GetClientReportIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientReportId.Get(), o.ClientReportId.IsSet()
}

// HasClientReportId returns a boolean if a field has been set.
func (o *AssetReportCreateRequestOptions) HasClientReportId() bool {
	if o != nil && o.ClientReportId.IsSet() {
		return true
	}

	return false
}

// SetClientReportId gets a reference to the given NullableString and assigns it to the ClientReportId field.
func (o *AssetReportCreateRequestOptions) SetClientReportId(v string) {
	o.ClientReportId.Set(&v)
}
// SetClientReportIdNil sets the value for ClientReportId to be an explicit nil
func (o *AssetReportCreateRequestOptions) SetClientReportIdNil() {
	o.ClientReportId.Set(nil)
}

// UnsetClientReportId ensures that no value is present for ClientReportId, not even an explicit nil
func (o *AssetReportCreateRequestOptions) UnsetClientReportId() {
	o.ClientReportId.Unset()
}

// GetWebhook returns the Webhook field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetReportCreateRequestOptions) GetWebhook() string {
	if o == nil || o.Webhook.Get() == nil {
		var ret string
		return ret
	}
	return *o.Webhook.Get()
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetReportCreateRequestOptions) GetWebhookOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Webhook.Get(), o.Webhook.IsSet()
}

// HasWebhook returns a boolean if a field has been set.
func (o *AssetReportCreateRequestOptions) HasWebhook() bool {
	if o != nil && o.Webhook.IsSet() {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given NullableString and assigns it to the Webhook field.
func (o *AssetReportCreateRequestOptions) SetWebhook(v string) {
	o.Webhook.Set(&v)
}
// SetWebhookNil sets the value for Webhook to be an explicit nil
func (o *AssetReportCreateRequestOptions) SetWebhookNil() {
	o.Webhook.Set(nil)
}

// UnsetWebhook ensures that no value is present for Webhook, not even an explicit nil
func (o *AssetReportCreateRequestOptions) UnsetWebhook() {
	o.Webhook.Unset()
}

// GetIncludeFastReport returns the IncludeFastReport field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetReportCreateRequestOptions) GetIncludeFastReport() bool {
	if o == nil || o.IncludeFastReport.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IncludeFastReport.Get()
}

// GetIncludeFastReportOk returns a tuple with the IncludeFastReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetReportCreateRequestOptions) GetIncludeFastReportOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IncludeFastReport.Get(), o.IncludeFastReport.IsSet()
}

// HasIncludeFastReport returns a boolean if a field has been set.
func (o *AssetReportCreateRequestOptions) HasIncludeFastReport() bool {
	if o != nil && o.IncludeFastReport.IsSet() {
		return true
	}

	return false
}

// SetIncludeFastReport gets a reference to the given NullableBool and assigns it to the IncludeFastReport field.
func (o *AssetReportCreateRequestOptions) SetIncludeFastReport(v bool) {
	o.IncludeFastReport.Set(&v)
}
// SetIncludeFastReportNil sets the value for IncludeFastReport to be an explicit nil
func (o *AssetReportCreateRequestOptions) SetIncludeFastReportNil() {
	o.IncludeFastReport.Set(nil)
}

// UnsetIncludeFastReport ensures that no value is present for IncludeFastReport, not even an explicit nil
func (o *AssetReportCreateRequestOptions) UnsetIncludeFastReport() {
	o.IncludeFastReport.Unset()
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *AssetReportCreateRequestOptions) GetProducts() []string {
	if o == nil || o.Products == nil {
		var ret []string
		return ret
	}
	return *o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportCreateRequestOptions) GetProductsOk() (*[]string, bool) {
	if o == nil || o.Products == nil {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *AssetReportCreateRequestOptions) HasProducts() bool {
	if o != nil && o.Products != nil {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []string and assigns it to the Products field.
func (o *AssetReportCreateRequestOptions) SetProducts(v []string) {
	o.Products = &v
}

// GetAddOns returns the AddOns field value if set, zero value otherwise.
func (o *AssetReportCreateRequestOptions) GetAddOns() []AssetReportAddOns {
	if o == nil || o.AddOns == nil {
		var ret []AssetReportAddOns
		return ret
	}
	return *o.AddOns
}

// GetAddOnsOk returns a tuple with the AddOns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportCreateRequestOptions) GetAddOnsOk() (*[]AssetReportAddOns, bool) {
	if o == nil || o.AddOns == nil {
		return nil, false
	}
	return o.AddOns, true
}

// HasAddOns returns a boolean if a field has been set.
func (o *AssetReportCreateRequestOptions) HasAddOns() bool {
	if o != nil && o.AddOns != nil {
		return true
	}

	return false
}

// SetAddOns gets a reference to the given []AssetReportAddOns and assigns it to the AddOns field.
func (o *AssetReportCreateRequestOptions) SetAddOns(v []AssetReportAddOns) {
	o.AddOns = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *AssetReportCreateRequestOptions) GetUser() AssetReportUser {
	if o == nil || o.User == nil {
		var ret AssetReportUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportCreateRequestOptions) GetUserOk() (*AssetReportUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *AssetReportCreateRequestOptions) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given AssetReportUser and assigns it to the User field.
func (o *AssetReportCreateRequestOptions) SetUser(v AssetReportUser) {
	o.User = &v
}

func (o AssetReportCreateRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientReportId.IsSet() {
		toSerialize["client_report_id"] = o.ClientReportId.Get()
	}
	if o.Webhook.IsSet() {
		toSerialize["webhook"] = o.Webhook.Get()
	}
	if o.IncludeFastReport.IsSet() {
		toSerialize["include_fast_report"] = o.IncludeFastReport.Get()
	}
	if o.Products != nil {
		toSerialize["products"] = o.Products
	}
	if o.AddOns != nil {
		toSerialize["add_ons"] = o.AddOns
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableAssetReportCreateRequestOptions struct {
	value *AssetReportCreateRequestOptions
	isSet bool
}

func (v NullableAssetReportCreateRequestOptions) Get() *AssetReportCreateRequestOptions {
	return v.value
}

func (v *NullableAssetReportCreateRequestOptions) Set(val *AssetReportCreateRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportCreateRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportCreateRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportCreateRequestOptions(val *AssetReportCreateRequestOptions) *NullableAssetReportCreateRequestOptions {
	return &NullableAssetReportCreateRequestOptions{value: val, isSet: true}
}

func (v NullableAssetReportCreateRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportCreateRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


