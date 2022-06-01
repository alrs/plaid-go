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

// SandboxBankTransferFireWebhookResponse Defines the response schema for `/sandbox/bank_transfer/fire_webhook`
type SandboxBankTransferFireWebhookResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _SandboxBankTransferFireWebhookResponse SandboxBankTransferFireWebhookResponse

// NewSandboxBankTransferFireWebhookResponse instantiates a new SandboxBankTransferFireWebhookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxBankTransferFireWebhookResponse(requestId string) *SandboxBankTransferFireWebhookResponse {
	this := SandboxBankTransferFireWebhookResponse{}
	this.RequestId = requestId
	return &this
}

// NewSandboxBankTransferFireWebhookResponseWithDefaults instantiates a new SandboxBankTransferFireWebhookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxBankTransferFireWebhookResponseWithDefaults() *SandboxBankTransferFireWebhookResponse {
	this := SandboxBankTransferFireWebhookResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *SandboxBankTransferFireWebhookResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *SandboxBankTransferFireWebhookResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *SandboxBankTransferFireWebhookResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o SandboxBankTransferFireWebhookResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SandboxBankTransferFireWebhookResponse) UnmarshalJSON(bytes []byte) (err error) {
	varSandboxBankTransferFireWebhookResponse := _SandboxBankTransferFireWebhookResponse{}

	if err = json.Unmarshal(bytes, &varSandboxBankTransferFireWebhookResponse); err == nil {
		*o = SandboxBankTransferFireWebhookResponse(varSandboxBankTransferFireWebhookResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSandboxBankTransferFireWebhookResponse struct {
	value *SandboxBankTransferFireWebhookResponse
	isSet bool
}

func (v NullableSandboxBankTransferFireWebhookResponse) Get() *SandboxBankTransferFireWebhookResponse {
	return v.value
}

func (v *NullableSandboxBankTransferFireWebhookResponse) Set(val *SandboxBankTransferFireWebhookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxBankTransferFireWebhookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxBankTransferFireWebhookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxBankTransferFireWebhookResponse(val *SandboxBankTransferFireWebhookResponse) *NullableSandboxBankTransferFireWebhookResponse {
	return &NullableSandboxBankTransferFireWebhookResponse{value: val, isSet: true}
}

func (v NullableSandboxBankTransferFireWebhookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxBankTransferFireWebhookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


