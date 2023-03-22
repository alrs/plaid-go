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

// CreditBankIncomeRefreshResponse CreditBankIncomeRefreshResponse defines the response schema for `/credit/bank_income/refresh`.
type CreditBankIncomeRefreshResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _CreditBankIncomeRefreshResponse CreditBankIncomeRefreshResponse

// NewCreditBankIncomeRefreshResponse instantiates a new CreditBankIncomeRefreshResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankIncomeRefreshResponse(requestId string) *CreditBankIncomeRefreshResponse {
	this := CreditBankIncomeRefreshResponse{}
	this.RequestId = requestId
	return &this
}

// NewCreditBankIncomeRefreshResponseWithDefaults instantiates a new CreditBankIncomeRefreshResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankIncomeRefreshResponseWithDefaults() *CreditBankIncomeRefreshResponse {
	this := CreditBankIncomeRefreshResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *CreditBankIncomeRefreshResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeRefreshResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *CreditBankIncomeRefreshResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o CreditBankIncomeRefreshResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditBankIncomeRefreshResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCreditBankIncomeRefreshResponse := _CreditBankIncomeRefreshResponse{}

	if err = json.Unmarshal(bytes, &varCreditBankIncomeRefreshResponse); err == nil {
		*o = CreditBankIncomeRefreshResponse(varCreditBankIncomeRefreshResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditBankIncomeRefreshResponse struct {
	value *CreditBankIncomeRefreshResponse
	isSet bool
}

func (v NullableCreditBankIncomeRefreshResponse) Get() *CreditBankIncomeRefreshResponse {
	return v.value
}

func (v *NullableCreditBankIncomeRefreshResponse) Set(val *CreditBankIncomeRefreshResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankIncomeRefreshResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankIncomeRefreshResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankIncomeRefreshResponse(val *CreditBankIncomeRefreshResponse) *NullableCreditBankIncomeRefreshResponse {
	return &NullableCreditBankIncomeRefreshResponse{value: val, isSet: true}
}

func (v NullableCreditBankIncomeRefreshResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankIncomeRefreshResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


