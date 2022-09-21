/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.181.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// Error We use standard HTTP response codes for success and failure notifications, and our errors are further classified by `error_type`. In general, 200 HTTP codes correspond to success, 40X codes are for developer- or user-related failures, and 50X codes are for Plaid-related issues.  Error fields will be `null` if no error has occurred.
type Error struct {
	// A broad categorization of the error. Safe for programmatic use.
	ErrorType string `json:"error_type"`
	// The particular error code. Safe for programmatic use.
	ErrorCode string `json:"error_code"`
	// A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.
	ErrorMessage string `json:"error_message"`
	// A user-friendly representation of the error code. `null` if the error is not related to user action.  This may change over time and is not safe for programmatic use.
	DisplayMessage NullableString `json:"display_message"`
	// A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks.
	RequestId *string `json:"request_id,omitempty"`
	// In the Assets product, a request can pertain to more than one Item. If an error is returned for such a request, `causes` will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  `causes` will only be provided for the `error_type` `ASSET_REPORT_ERROR`. `causes` will also not be populated inside an error nested within a `warning` object.
	Causes *[]interface{} `json:"causes,omitempty"`
	// The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook.
	Status NullableFloat32 `json:"status,omitempty"`
	// The URL of a Plaid documentation page with more information about the error
	DocumentationUrl *string `json:"documentation_url,omitempty"`
	// Suggested steps for resolving the error
	SuggestedAction NullableString `json:"suggested_action,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Error Error

// NewError instantiates a new Error object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewError(errorType string, errorCode string, errorMessage string, displayMessage NullableString) *Error {
	this := Error{}
	this.ErrorType = errorType
	this.ErrorCode = errorCode
	this.ErrorMessage = errorMessage
	this.DisplayMessage = displayMessage
	return &this
}

// NewErrorWithDefaults instantiates a new Error object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorWithDefaults() *Error {
	this := Error{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *Error) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *Error) GetErrorTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *Error) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorCode returns the ErrorCode field value
func (o *Error) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *Error) GetErrorCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *Error) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetErrorMessage returns the ErrorMessage field value
func (o *Error) GetErrorMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value
// and a boolean to check if the value has been set.
func (o *Error) GetErrorMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorMessage, true
}

// SetErrorMessage sets field value
func (o *Error) SetErrorMessage(v string) {
	o.ErrorMessage = v
}

// GetDisplayMessage returns the DisplayMessage field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Error) GetDisplayMessage() string {
	if o == nil || o.DisplayMessage.Get() == nil {
		var ret string
		return ret
	}

	return *o.DisplayMessage.Get()
}

// GetDisplayMessageOk returns a tuple with the DisplayMessage field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Error) GetDisplayMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayMessage.Get(), o.DisplayMessage.IsSet()
}

// SetDisplayMessage sets field value
func (o *Error) SetDisplayMessage(v string) {
	o.DisplayMessage.Set(&v)
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *Error) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *Error) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *Error) SetRequestId(v string) {
	o.RequestId = &v
}

// GetCauses returns the Causes field value if set, zero value otherwise.
func (o *Error) GetCauses() []interface{} {
	if o == nil || o.Causes == nil {
		var ret []interface{}
		return ret
	}
	return *o.Causes
}

// GetCausesOk returns a tuple with the Causes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetCausesOk() (*[]interface{}, bool) {
	if o == nil || o.Causes == nil {
		return nil, false
	}
	return o.Causes, true
}

// HasCauses returns a boolean if a field has been set.
func (o *Error) HasCauses() bool {
	if o != nil && o.Causes != nil {
		return true
	}

	return false
}

// SetCauses gets a reference to the given []interface{} and assigns it to the Causes field.
func (o *Error) SetCauses(v []interface{}) {
	o.Causes = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Error) GetStatus() float32 {
	if o == nil || o.Status.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Error) GetStatusOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *Error) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableFloat32 and assigns it to the Status field.
func (o *Error) SetStatus(v float32) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *Error) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *Error) UnsetStatus() {
	o.Status.Unset()
}

// GetDocumentationUrl returns the DocumentationUrl field value if set, zero value otherwise.
func (o *Error) GetDocumentationUrl() string {
	if o == nil || o.DocumentationUrl == nil {
		var ret string
		return ret
	}
	return *o.DocumentationUrl
}

// GetDocumentationUrlOk returns a tuple with the DocumentationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetDocumentationUrlOk() (*string, bool) {
	if o == nil || o.DocumentationUrl == nil {
		return nil, false
	}
	return o.DocumentationUrl, true
}

// HasDocumentationUrl returns a boolean if a field has been set.
func (o *Error) HasDocumentationUrl() bool {
	if o != nil && o.DocumentationUrl != nil {
		return true
	}

	return false
}

// SetDocumentationUrl gets a reference to the given string and assigns it to the DocumentationUrl field.
func (o *Error) SetDocumentationUrl(v string) {
	o.DocumentationUrl = &v
}

// GetSuggestedAction returns the SuggestedAction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Error) GetSuggestedAction() string {
	if o == nil || o.SuggestedAction.Get() == nil {
		var ret string
		return ret
	}
	return *o.SuggestedAction.Get()
}

// GetSuggestedActionOk returns a tuple with the SuggestedAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Error) GetSuggestedActionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SuggestedAction.Get(), o.SuggestedAction.IsSet()
}

// HasSuggestedAction returns a boolean if a field has been set.
func (o *Error) HasSuggestedAction() bool {
	if o != nil && o.SuggestedAction.IsSet() {
		return true
	}

	return false
}

// SetSuggestedAction gets a reference to the given NullableString and assigns it to the SuggestedAction field.
func (o *Error) SetSuggestedAction(v string) {
	o.SuggestedAction.Set(&v)
}
// SetSuggestedActionNil sets the value for SuggestedAction to be an explicit nil
func (o *Error) SetSuggestedActionNil() {
	o.SuggestedAction.Set(nil)
}

// UnsetSuggestedAction ensures that no value is present for SuggestedAction, not even an explicit nil
func (o *Error) UnsetSuggestedAction() {
	o.SuggestedAction.Unset()
}

func (o Error) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["error_type"] = o.ErrorType
	}
	if true {
		toSerialize["error_code"] = o.ErrorCode
	}
	if true {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if true {
		toSerialize["display_message"] = o.DisplayMessage.Get()
	}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}
	if o.Causes != nil {
		toSerialize["causes"] = o.Causes
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.DocumentationUrl != nil {
		toSerialize["documentation_url"] = o.DocumentationUrl
	}
	if o.SuggestedAction.IsSet() {
		toSerialize["suggested_action"] = o.SuggestedAction.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Error) UnmarshalJSON(bytes []byte) (err error) {
	varError := _Error{}

	if err = json.Unmarshal(bytes, &varError); err == nil {
		*o = Error(varError)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error_type")
		delete(additionalProperties, "error_code")
		delete(additionalProperties, "error_message")
		delete(additionalProperties, "display_message")
		delete(additionalProperties, "request_id")
		delete(additionalProperties, "causes")
		delete(additionalProperties, "status")
		delete(additionalProperties, "documentation_url")
		delete(additionalProperties, "suggested_action")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableError struct {
	value *Error
	isSet bool
}

func (v NullableError) Get() *Error {
	return v.value
}

func (v *NullableError) Set(val *Error) {
	v.value = val
	v.isSet = true
}

func (v NullableError) IsSet() bool {
	return v.isSet
}

func (v *NullableError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableError(val *Error) *NullableError {
	return &NullableError{value: val, isSet: true}
}

func (v NullableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


