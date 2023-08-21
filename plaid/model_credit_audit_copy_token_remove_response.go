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

// CreditAuditCopyTokenRemoveResponse CreditAuditCopyTokenRemoveResponse defines the response schema for `/credit/audit_copy_token/remove`
type CreditAuditCopyTokenRemoveResponse struct {
	// `true` if the Audit Copy was successfully removed.
	Removed bool `json:"removed"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _CreditAuditCopyTokenRemoveResponse CreditAuditCopyTokenRemoveResponse

// NewCreditAuditCopyTokenRemoveResponse instantiates a new CreditAuditCopyTokenRemoveResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditAuditCopyTokenRemoveResponse(removed bool, requestId string) *CreditAuditCopyTokenRemoveResponse {
	this := CreditAuditCopyTokenRemoveResponse{}
	this.Removed = removed
	this.RequestId = requestId
	return &this
}

// NewCreditAuditCopyTokenRemoveResponseWithDefaults instantiates a new CreditAuditCopyTokenRemoveResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditAuditCopyTokenRemoveResponseWithDefaults() *CreditAuditCopyTokenRemoveResponse {
	this := CreditAuditCopyTokenRemoveResponse{}
	return &this
}

// GetRemoved returns the Removed field value
func (o *CreditAuditCopyTokenRemoveResponse) GetRemoved() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Removed
}

// GetRemovedOk returns a tuple with the Removed field value
// and a boolean to check if the value has been set.
func (o *CreditAuditCopyTokenRemoveResponse) GetRemovedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Removed, true
}

// SetRemoved sets field value
func (o *CreditAuditCopyTokenRemoveResponse) SetRemoved(v bool) {
	o.Removed = v
}

// GetRequestId returns the RequestId field value
func (o *CreditAuditCopyTokenRemoveResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *CreditAuditCopyTokenRemoveResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *CreditAuditCopyTokenRemoveResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o CreditAuditCopyTokenRemoveResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["removed"] = o.Removed
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditAuditCopyTokenRemoveResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCreditAuditCopyTokenRemoveResponse := _CreditAuditCopyTokenRemoveResponse{}

	if err = json.Unmarshal(bytes, &varCreditAuditCopyTokenRemoveResponse); err == nil {
		*o = CreditAuditCopyTokenRemoveResponse(varCreditAuditCopyTokenRemoveResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "removed")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditAuditCopyTokenRemoveResponse struct {
	value *CreditAuditCopyTokenRemoveResponse
	isSet bool
}

func (v NullableCreditAuditCopyTokenRemoveResponse) Get() *CreditAuditCopyTokenRemoveResponse {
	return v.value
}

func (v *NullableCreditAuditCopyTokenRemoveResponse) Set(val *CreditAuditCopyTokenRemoveResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditAuditCopyTokenRemoveResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditAuditCopyTokenRemoveResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditAuditCopyTokenRemoveResponse(val *CreditAuditCopyTokenRemoveResponse) *NullableCreditAuditCopyTokenRemoveResponse {
	return &NullableCreditAuditCopyTokenRemoveResponse{value: val, isSet: true}
}

func (v NullableCreditAuditCopyTokenRemoveResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditAuditCopyTokenRemoveResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


