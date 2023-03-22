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

// AssetReportFilterRequest AssetReportFilterRequest defines the request schema for `/asset_report/filter`
type AssetReportFilterRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// A token that can be provided to endpoints such as `/asset_report/get` or `/asset_report/pdf/get` to fetch or update an Asset Report.
	AssetReportToken string `json:"asset_report_token"`
	// The accounts to exclude from the Asset Report, identified by `account_id`.
	AccountIdsToExclude []string `json:"account_ids_to_exclude"`
}

// NewAssetReportFilterRequest instantiates a new AssetReportFilterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportFilterRequest(assetReportToken string, accountIdsToExclude []string) *AssetReportFilterRequest {
	this := AssetReportFilterRequest{}
	this.AssetReportToken = assetReportToken
	this.AccountIdsToExclude = accountIdsToExclude
	return &this
}

// NewAssetReportFilterRequestWithDefaults instantiates a new AssetReportFilterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportFilterRequestWithDefaults() *AssetReportFilterRequest {
	this := AssetReportFilterRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AssetReportFilterRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportFilterRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AssetReportFilterRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AssetReportFilterRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *AssetReportFilterRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportFilterRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *AssetReportFilterRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *AssetReportFilterRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAssetReportToken returns the AssetReportToken field value
func (o *AssetReportFilterRequest) GetAssetReportToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetReportToken
}

// GetAssetReportTokenOk returns a tuple with the AssetReportToken field value
// and a boolean to check if the value has been set.
func (o *AssetReportFilterRequest) GetAssetReportTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetReportToken, true
}

// SetAssetReportToken sets field value
func (o *AssetReportFilterRequest) SetAssetReportToken(v string) {
	o.AssetReportToken = v
}

// GetAccountIdsToExclude returns the AccountIdsToExclude field value
func (o *AssetReportFilterRequest) GetAccountIdsToExclude() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AccountIdsToExclude
}

// GetAccountIdsToExcludeOk returns a tuple with the AccountIdsToExclude field value
// and a boolean to check if the value has been set.
func (o *AssetReportFilterRequest) GetAccountIdsToExcludeOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountIdsToExclude, true
}

// SetAccountIdsToExclude sets field value
func (o *AssetReportFilterRequest) SetAccountIdsToExclude(v []string) {
	o.AccountIdsToExclude = v
}

func (o AssetReportFilterRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["asset_report_token"] = o.AssetReportToken
	}
	if true {
		toSerialize["account_ids_to_exclude"] = o.AccountIdsToExclude
	}
	return json.Marshal(toSerialize)
}

type NullableAssetReportFilterRequest struct {
	value *AssetReportFilterRequest
	isSet bool
}

func (v NullableAssetReportFilterRequest) Get() *AssetReportFilterRequest {
	return v.value
}

func (v *NullableAssetReportFilterRequest) Set(val *AssetReportFilterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportFilterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportFilterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportFilterRequest(val *AssetReportFilterRequest) *NullableAssetReportFilterRequest {
	return &NullableAssetReportFilterRequest{value: val, isSet: true}
}

func (v NullableAssetReportFilterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportFilterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


