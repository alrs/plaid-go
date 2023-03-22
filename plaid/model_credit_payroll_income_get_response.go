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

// CreditPayrollIncomeGetResponse Defines the response body for `/credit/payroll_income/get`.
type CreditPayrollIncomeGetResponse struct {
	// Array of payroll items.
	Items []PayrollItem `json:"items"`
	Error NullablePlaidError `json:"error,omitempty"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _CreditPayrollIncomeGetResponse CreditPayrollIncomeGetResponse

// NewCreditPayrollIncomeGetResponse instantiates a new CreditPayrollIncomeGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditPayrollIncomeGetResponse(items []PayrollItem, requestId string) *CreditPayrollIncomeGetResponse {
	this := CreditPayrollIncomeGetResponse{}
	this.Items = items
	this.RequestId = requestId
	return &this
}

// NewCreditPayrollIncomeGetResponseWithDefaults instantiates a new CreditPayrollIncomeGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditPayrollIncomeGetResponseWithDefaults() *CreditPayrollIncomeGetResponse {
	this := CreditPayrollIncomeGetResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *CreditPayrollIncomeGetResponse) GetItems() []PayrollItem {
	if o == nil {
		var ret []PayrollItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CreditPayrollIncomeGetResponse) GetItemsOk() (*[]PayrollItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *CreditPayrollIncomeGetResponse) SetItems(v []PayrollItem) {
	o.Items = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditPayrollIncomeGetResponse) GetError() PlaidError {
	if o == nil || o.Error.Get() == nil {
		var ret PlaidError
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditPayrollIncomeGetResponse) GetErrorOk() (*PlaidError, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *CreditPayrollIncomeGetResponse) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullablePlaidError and assigns it to the Error field.
func (o *CreditPayrollIncomeGetResponse) SetError(v PlaidError) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *CreditPayrollIncomeGetResponse) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *CreditPayrollIncomeGetResponse) UnsetError() {
	o.Error.Unset()
}

// GetRequestId returns the RequestId field value
func (o *CreditPayrollIncomeGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *CreditPayrollIncomeGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *CreditPayrollIncomeGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o CreditPayrollIncomeGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditPayrollIncomeGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCreditPayrollIncomeGetResponse := _CreditPayrollIncomeGetResponse{}

	if err = json.Unmarshal(bytes, &varCreditPayrollIncomeGetResponse); err == nil {
		*o = CreditPayrollIncomeGetResponse(varCreditPayrollIncomeGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		delete(additionalProperties, "error")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditPayrollIncomeGetResponse struct {
	value *CreditPayrollIncomeGetResponse
	isSet bool
}

func (v NullableCreditPayrollIncomeGetResponse) Get() *CreditPayrollIncomeGetResponse {
	return v.value
}

func (v *NullableCreditPayrollIncomeGetResponse) Set(val *CreditPayrollIncomeGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditPayrollIncomeGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditPayrollIncomeGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditPayrollIncomeGetResponse(val *CreditPayrollIncomeGetResponse) *NullableCreditPayrollIncomeGetResponse {
	return &NullableCreditPayrollIncomeGetResponse{value: val, isSet: true}
}

func (v NullableCreditPayrollIncomeGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditPayrollIncomeGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


