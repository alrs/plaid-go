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

// AssetReportRelayCreateResponse AssetReportRelayCreateResponse defines the response schema for `/asset_report/relay/create`
type AssetReportRelayCreateResponse struct {
	// A token that can be shared with a third party to allow them to access the Asset Report. This token should be stored securely.
	AssetRelayToken string `json:"asset_relay_token"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _AssetReportRelayCreateResponse AssetReportRelayCreateResponse

// NewAssetReportRelayCreateResponse instantiates a new AssetReportRelayCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportRelayCreateResponse(assetRelayToken string, requestId string) *AssetReportRelayCreateResponse {
	this := AssetReportRelayCreateResponse{}
	this.AssetRelayToken = assetRelayToken
	this.RequestId = requestId
	return &this
}

// NewAssetReportRelayCreateResponseWithDefaults instantiates a new AssetReportRelayCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportRelayCreateResponseWithDefaults() *AssetReportRelayCreateResponse {
	this := AssetReportRelayCreateResponse{}
	return &this
}

// GetAssetRelayToken returns the AssetRelayToken field value
func (o *AssetReportRelayCreateResponse) GetAssetRelayToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetRelayToken
}

// GetAssetRelayTokenOk returns a tuple with the AssetRelayToken field value
// and a boolean to check if the value has been set.
func (o *AssetReportRelayCreateResponse) GetAssetRelayTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetRelayToken, true
}

// SetAssetRelayToken sets field value
func (o *AssetReportRelayCreateResponse) SetAssetRelayToken(v string) {
	o.AssetRelayToken = v
}

// GetRequestId returns the RequestId field value
func (o *AssetReportRelayCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *AssetReportRelayCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *AssetReportRelayCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o AssetReportRelayCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["asset_relay_token"] = o.AssetRelayToken
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetReportRelayCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAssetReportRelayCreateResponse := _AssetReportRelayCreateResponse{}

	if err = json.Unmarshal(bytes, &varAssetReportRelayCreateResponse); err == nil {
		*o = AssetReportRelayCreateResponse(varAssetReportRelayCreateResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "asset_relay_token")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetReportRelayCreateResponse struct {
	value *AssetReportRelayCreateResponse
	isSet bool
}

func (v NullableAssetReportRelayCreateResponse) Get() *AssetReportRelayCreateResponse {
	return v.value
}

func (v *NullableAssetReportRelayCreateResponse) Set(val *AssetReportRelayCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportRelayCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportRelayCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportRelayCreateResponse(val *AssetReportRelayCreateResponse) *NullableAssetReportRelayCreateResponse {
	return &NullableAssetReportRelayCreateResponse{value: val, isSet: true}
}

func (v NullableAssetReportRelayCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportRelayCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


