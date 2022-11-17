/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.205.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// CreditAuditCopyTokenCreateRequest CreditAuditCopyTokenCreateRequest defines the request schema for `/credit/audit_copy_token/create`
type CreditAuditCopyTokenCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// List of report tokens; can include both Asset Report tokens and Income Report tokens.
	ReportTokens []string `json:"report_tokens"`
}

// NewCreditAuditCopyTokenCreateRequest instantiates a new CreditAuditCopyTokenCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditAuditCopyTokenCreateRequest(reportTokens []string) *CreditAuditCopyTokenCreateRequest {
	this := CreditAuditCopyTokenCreateRequest{}
	this.ReportTokens = reportTokens
	return &this
}

// NewCreditAuditCopyTokenCreateRequestWithDefaults instantiates a new CreditAuditCopyTokenCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditAuditCopyTokenCreateRequestWithDefaults() *CreditAuditCopyTokenCreateRequest {
	this := CreditAuditCopyTokenCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CreditAuditCopyTokenCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditAuditCopyTokenCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CreditAuditCopyTokenCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CreditAuditCopyTokenCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *CreditAuditCopyTokenCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditAuditCopyTokenCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *CreditAuditCopyTokenCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *CreditAuditCopyTokenCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetReportTokens returns the ReportTokens field value
func (o *CreditAuditCopyTokenCreateRequest) GetReportTokens() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ReportTokens
}

// GetReportTokensOk returns a tuple with the ReportTokens field value
// and a boolean to check if the value has been set.
func (o *CreditAuditCopyTokenCreateRequest) GetReportTokensOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReportTokens, true
}

// SetReportTokens sets field value
func (o *CreditAuditCopyTokenCreateRequest) SetReportTokens(v []string) {
	o.ReportTokens = v
}

func (o CreditAuditCopyTokenCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["report_tokens"] = o.ReportTokens
	}
	return json.Marshal(toSerialize)
}

type NullableCreditAuditCopyTokenCreateRequest struct {
	value *CreditAuditCopyTokenCreateRequest
	isSet bool
}

func (v NullableCreditAuditCopyTokenCreateRequest) Get() *CreditAuditCopyTokenCreateRequest {
	return v.value
}

func (v *NullableCreditAuditCopyTokenCreateRequest) Set(val *CreditAuditCopyTokenCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditAuditCopyTokenCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditAuditCopyTokenCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditAuditCopyTokenCreateRequest(val *CreditAuditCopyTokenCreateRequest) *NullableCreditAuditCopyTokenCreateRequest {
	return &NullableCreditAuditCopyTokenCreateRequest{value: val, isSet: true}
}

func (v NullableCreditAuditCopyTokenCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditAuditCopyTokenCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


