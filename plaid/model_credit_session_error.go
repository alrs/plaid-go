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

// CreditSessionError The details of a Link error.
type CreditSessionError struct {
	// A broad categorization of the error.
	ErrorType *string `json:"error_type,omitempty"`
	// The particular error code.
	ErrorCode *string `json:"error_code,omitempty"`
	// A developer-friendly representation of the error code.
	ErrorMessage *string `json:"error_message,omitempty"`
	// A user-friendly representation of the error code. `null` if the error is not related to user action.
	DisplayMessage NullableString `json:"display_message,omitempty"`
}

// NewCreditSessionError instantiates a new CreditSessionError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditSessionError() *CreditSessionError {
	this := CreditSessionError{}
	return &this
}

// NewCreditSessionErrorWithDefaults instantiates a new CreditSessionError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditSessionErrorWithDefaults() *CreditSessionError {
	this := CreditSessionError{}
	return &this
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *CreditSessionError) GetErrorType() string {
	if o == nil || o.ErrorType == nil {
		var ret string
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditSessionError) GetErrorTypeOk() (*string, bool) {
	if o == nil || o.ErrorType == nil {
		return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *CreditSessionError) HasErrorType() bool {
	if o != nil && o.ErrorType != nil {
		return true
	}

	return false
}

// SetErrorType gets a reference to the given string and assigns it to the ErrorType field.
func (o *CreditSessionError) SetErrorType(v string) {
	o.ErrorType = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *CreditSessionError) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditSessionError) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *CreditSessionError) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *CreditSessionError) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *CreditSessionError) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditSessionError) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *CreditSessionError) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *CreditSessionError) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetDisplayMessage returns the DisplayMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditSessionError) GetDisplayMessage() string {
	if o == nil || o.DisplayMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayMessage.Get()
}

// GetDisplayMessageOk returns a tuple with the DisplayMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditSessionError) GetDisplayMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayMessage.Get(), o.DisplayMessage.IsSet()
}

// HasDisplayMessage returns a boolean if a field has been set.
func (o *CreditSessionError) HasDisplayMessage() bool {
	if o != nil && o.DisplayMessage.IsSet() {
		return true
	}

	return false
}

// SetDisplayMessage gets a reference to the given NullableString and assigns it to the DisplayMessage field.
func (o *CreditSessionError) SetDisplayMessage(v string) {
	o.DisplayMessage.Set(&v)
}
// SetDisplayMessageNil sets the value for DisplayMessage to be an explicit nil
func (o *CreditSessionError) SetDisplayMessageNil() {
	o.DisplayMessage.Set(nil)
}

// UnsetDisplayMessage ensures that no value is present for DisplayMessage, not even an explicit nil
func (o *CreditSessionError) UnsetDisplayMessage() {
	o.DisplayMessage.Unset()
}

func (o CreditSessionError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorType != nil {
		toSerialize["error_type"] = o.ErrorType
	}
	if o.ErrorCode != nil {
		toSerialize["error_code"] = o.ErrorCode
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.DisplayMessage.IsSet() {
		toSerialize["display_message"] = o.DisplayMessage.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCreditSessionError struct {
	value *CreditSessionError
	isSet bool
}

func (v NullableCreditSessionError) Get() *CreditSessionError {
	return v.value
}

func (v *NullableCreditSessionError) Set(val *CreditSessionError) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditSessionError) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditSessionError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditSessionError(val *CreditSessionError) *NullableCreditSessionError {
	return &NullableCreditSessionError{value: val, isSet: true}
}

func (v NullableCreditSessionError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditSessionError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


