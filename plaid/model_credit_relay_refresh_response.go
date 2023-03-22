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

// CreditRelayRefreshResponse CreditRelayRefreshResponse defines the response schema for `/credit/relay/refresh`
type CreditRelayRefreshResponse struct {
	RelayToken string `json:"relay_token"`
	// A unique ID identifying an Asset Report. Like all Plaid identifiers, this ID is case sensitive.
	AssetReportId *string `json:"asset_report_id,omitempty"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _CreditRelayRefreshResponse CreditRelayRefreshResponse

// NewCreditRelayRefreshResponse instantiates a new CreditRelayRefreshResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditRelayRefreshResponse(relayToken string, requestId string) *CreditRelayRefreshResponse {
	this := CreditRelayRefreshResponse{}
	this.RelayToken = relayToken
	this.RequestId = requestId
	return &this
}

// NewCreditRelayRefreshResponseWithDefaults instantiates a new CreditRelayRefreshResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditRelayRefreshResponseWithDefaults() *CreditRelayRefreshResponse {
	this := CreditRelayRefreshResponse{}
	return &this
}

// GetRelayToken returns the RelayToken field value
func (o *CreditRelayRefreshResponse) GetRelayToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelayToken
}

// GetRelayTokenOk returns a tuple with the RelayToken field value
// and a boolean to check if the value has been set.
func (o *CreditRelayRefreshResponse) GetRelayTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RelayToken, true
}

// SetRelayToken sets field value
func (o *CreditRelayRefreshResponse) SetRelayToken(v string) {
	o.RelayToken = v
}

// GetAssetReportId returns the AssetReportId field value if set, zero value otherwise.
func (o *CreditRelayRefreshResponse) GetAssetReportId() string {
	if o == nil || o.AssetReportId == nil {
		var ret string
		return ret
	}
	return *o.AssetReportId
}

// GetAssetReportIdOk returns a tuple with the AssetReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditRelayRefreshResponse) GetAssetReportIdOk() (*string, bool) {
	if o == nil || o.AssetReportId == nil {
		return nil, false
	}
	return o.AssetReportId, true
}

// HasAssetReportId returns a boolean if a field has been set.
func (o *CreditRelayRefreshResponse) HasAssetReportId() bool {
	if o != nil && o.AssetReportId != nil {
		return true
	}

	return false
}

// SetAssetReportId gets a reference to the given string and assigns it to the AssetReportId field.
func (o *CreditRelayRefreshResponse) SetAssetReportId(v string) {
	o.AssetReportId = &v
}

// GetRequestId returns the RequestId field value
func (o *CreditRelayRefreshResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *CreditRelayRefreshResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *CreditRelayRefreshResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o CreditRelayRefreshResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["relay_token"] = o.RelayToken
	}
	if o.AssetReportId != nil {
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

func (o *CreditRelayRefreshResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCreditRelayRefreshResponse := _CreditRelayRefreshResponse{}

	if err = json.Unmarshal(bytes, &varCreditRelayRefreshResponse); err == nil {
		*o = CreditRelayRefreshResponse(varCreditRelayRefreshResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "relay_token")
		delete(additionalProperties, "asset_report_id")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditRelayRefreshResponse struct {
	value *CreditRelayRefreshResponse
	isSet bool
}

func (v NullableCreditRelayRefreshResponse) Get() *CreditRelayRefreshResponse {
	return v.value
}

func (v *NullableCreditRelayRefreshResponse) Set(val *CreditRelayRefreshResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditRelayRefreshResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditRelayRefreshResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditRelayRefreshResponse(val *CreditRelayRefreshResponse) *NullableCreditRelayRefreshResponse {
	return &NullableCreditRelayRefreshResponse{value: val, isSet: true}
}

func (v NullableCreditRelayRefreshResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditRelayRefreshResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


