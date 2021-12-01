/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.54.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// IncomeVerificationRefreshResponse IncomeVerificationRequestResponse defines the response schema for `/income/verification/refresh`
type IncomeVerificationRefreshResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	VerificationRefreshStatus VerificationRefreshStatus `json:"verification_refresh_status"`
}

// NewIncomeVerificationRefreshResponse instantiates a new IncomeVerificationRefreshResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeVerificationRefreshResponse(requestId string, verificationRefreshStatus VerificationRefreshStatus) *IncomeVerificationRefreshResponse {
	this := IncomeVerificationRefreshResponse{}
	this.RequestId = requestId
	this.VerificationRefreshStatus = verificationRefreshStatus
	return &this
}

// NewIncomeVerificationRefreshResponseWithDefaults instantiates a new IncomeVerificationRefreshResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeVerificationRefreshResponseWithDefaults() *IncomeVerificationRefreshResponse {
	this := IncomeVerificationRefreshResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *IncomeVerificationRefreshResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *IncomeVerificationRefreshResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *IncomeVerificationRefreshResponse) SetRequestId(v string) {
	o.RequestId = v
}

// GetVerificationRefreshStatus returns the VerificationRefreshStatus field value
func (o *IncomeVerificationRefreshResponse) GetVerificationRefreshStatus() VerificationRefreshStatus {
	if o == nil {
		var ret VerificationRefreshStatus
		return ret
	}

	return o.VerificationRefreshStatus
}

// GetVerificationRefreshStatusOk returns a tuple with the VerificationRefreshStatus field value
// and a boolean to check if the value has been set.
func (o *IncomeVerificationRefreshResponse) GetVerificationRefreshStatusOk() (*VerificationRefreshStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerificationRefreshStatus, true
}

// SetVerificationRefreshStatus sets field value
func (o *IncomeVerificationRefreshResponse) SetVerificationRefreshStatus(v VerificationRefreshStatus) {
	o.VerificationRefreshStatus = v
}

func (o IncomeVerificationRefreshResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}
	if true {
		toSerialize["verification_refresh_status"] = o.VerificationRefreshStatus
	}
	return json.Marshal(toSerialize)
}

type NullableIncomeVerificationRefreshResponse struct {
	value *IncomeVerificationRefreshResponse
	isSet bool
}

func (v NullableIncomeVerificationRefreshResponse) Get() *IncomeVerificationRefreshResponse {
	return v.value
}

func (v *NullableIncomeVerificationRefreshResponse) Set(val *IncomeVerificationRefreshResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationRefreshResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationRefreshResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationRefreshResponse(val *IncomeVerificationRefreshResponse) *NullableIncomeVerificationRefreshResponse {
	return &NullableIncomeVerificationRefreshResponse{value: val, isSet: true}
}

func (v NullableIncomeVerificationRefreshResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationRefreshResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

