/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.121.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// AssetReportRelayRefreshResponse AssetReportRelayRefreshResponse defines the response schema for `/asset_report/relay/refresh`
type AssetReportRelayRefreshResponse struct {
	AssetRelayToken string `json:"asset_relay_token"`
	// A unique ID identifying an Asset Report. Like all Plaid identifiers, this ID is case sensitive.
	AssetReportId string `json:"asset_report_id"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _AssetReportRelayRefreshResponse AssetReportRelayRefreshResponse

// NewAssetReportRelayRefreshResponse instantiates a new AssetReportRelayRefreshResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportRelayRefreshResponse(assetRelayToken string, assetReportId string, requestId string) *AssetReportRelayRefreshResponse {
	this := AssetReportRelayRefreshResponse{}
	this.AssetRelayToken = assetRelayToken
	this.AssetReportId = assetReportId
	this.RequestId = requestId
	return &this
}

// NewAssetReportRelayRefreshResponseWithDefaults instantiates a new AssetReportRelayRefreshResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportRelayRefreshResponseWithDefaults() *AssetReportRelayRefreshResponse {
	this := AssetReportRelayRefreshResponse{}
	return &this
}

// GetAssetRelayToken returns the AssetRelayToken field value
func (o *AssetReportRelayRefreshResponse) GetAssetRelayToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetRelayToken
}

// GetAssetRelayTokenOk returns a tuple with the AssetRelayToken field value
// and a boolean to check if the value has been set.
func (o *AssetReportRelayRefreshResponse) GetAssetRelayTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetRelayToken, true
}

// SetAssetRelayToken sets field value
func (o *AssetReportRelayRefreshResponse) SetAssetRelayToken(v string) {
	o.AssetRelayToken = v
}

// GetAssetReportId returns the AssetReportId field value
func (o *AssetReportRelayRefreshResponse) GetAssetReportId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetReportId
}

// GetAssetReportIdOk returns a tuple with the AssetReportId field value
// and a boolean to check if the value has been set.
func (o *AssetReportRelayRefreshResponse) GetAssetReportIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetReportId, true
}

// SetAssetReportId sets field value
func (o *AssetReportRelayRefreshResponse) SetAssetReportId(v string) {
	o.AssetReportId = v
}

// GetRequestId returns the RequestId field value
func (o *AssetReportRelayRefreshResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *AssetReportRelayRefreshResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *AssetReportRelayRefreshResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o AssetReportRelayRefreshResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["asset_relay_token"] = o.AssetRelayToken
	}
	if true {
		toSerialize["asset_report_id"] = o.AssetReportId
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetReportRelayRefreshResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAssetReportRelayRefreshResponse := _AssetReportRelayRefreshResponse{}

	if err = json.Unmarshal(bytes, &varAssetReportRelayRefreshResponse); err == nil {
		*o = AssetReportRelayRefreshResponse(varAssetReportRelayRefreshResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "asset_relay_token")
		delete(additionalProperties, "asset_report_id")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetReportRelayRefreshResponse struct {
	value *AssetReportRelayRefreshResponse
	isSet bool
}

func (v NullableAssetReportRelayRefreshResponse) Get() *AssetReportRelayRefreshResponse {
	return v.value
}

func (v *NullableAssetReportRelayRefreshResponse) Set(val *AssetReportRelayRefreshResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportRelayRefreshResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportRelayRefreshResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportRelayRefreshResponse(val *AssetReportRelayRefreshResponse) *NullableAssetReportRelayRefreshResponse {
	return &NullableAssetReportRelayRefreshResponse{value: val, isSet: true}
}

func (v NullableAssetReportRelayRefreshResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportRelayRefreshResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


