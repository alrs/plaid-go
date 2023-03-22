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

// CreditAuditCopyTokenRemoveRequest CreditAuditCopyTokenRemoveRequest defines the request schema for `/credit/audit_copy_token/remove`
type CreditAuditCopyTokenRemoveRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The `audit_copy_token` granting access to the Audit Copy you would like to revoke.
	AuditCopyToken string `json:"audit_copy_token"`
}

// NewCreditAuditCopyTokenRemoveRequest instantiates a new CreditAuditCopyTokenRemoveRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditAuditCopyTokenRemoveRequest(auditCopyToken string) *CreditAuditCopyTokenRemoveRequest {
	this := CreditAuditCopyTokenRemoveRequest{}
	this.AuditCopyToken = auditCopyToken
	return &this
}

// NewCreditAuditCopyTokenRemoveRequestWithDefaults instantiates a new CreditAuditCopyTokenRemoveRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditAuditCopyTokenRemoveRequestWithDefaults() *CreditAuditCopyTokenRemoveRequest {
	this := CreditAuditCopyTokenRemoveRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CreditAuditCopyTokenRemoveRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditAuditCopyTokenRemoveRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CreditAuditCopyTokenRemoveRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CreditAuditCopyTokenRemoveRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *CreditAuditCopyTokenRemoveRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditAuditCopyTokenRemoveRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *CreditAuditCopyTokenRemoveRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *CreditAuditCopyTokenRemoveRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAuditCopyToken returns the AuditCopyToken field value
func (o *CreditAuditCopyTokenRemoveRequest) GetAuditCopyToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuditCopyToken
}

// GetAuditCopyTokenOk returns a tuple with the AuditCopyToken field value
// and a boolean to check if the value has been set.
func (o *CreditAuditCopyTokenRemoveRequest) GetAuditCopyTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuditCopyToken, true
}

// SetAuditCopyToken sets field value
func (o *CreditAuditCopyTokenRemoveRequest) SetAuditCopyToken(v string) {
	o.AuditCopyToken = v
}

func (o CreditAuditCopyTokenRemoveRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["audit_copy_token"] = o.AuditCopyToken
	}
	return json.Marshal(toSerialize)
}

type NullableCreditAuditCopyTokenRemoveRequest struct {
	value *CreditAuditCopyTokenRemoveRequest
	isSet bool
}

func (v NullableCreditAuditCopyTokenRemoveRequest) Get() *CreditAuditCopyTokenRemoveRequest {
	return v.value
}

func (v *NullableCreditAuditCopyTokenRemoveRequest) Set(val *CreditAuditCopyTokenRemoveRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditAuditCopyTokenRemoveRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditAuditCopyTokenRemoveRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditAuditCopyTokenRemoveRequest(val *CreditAuditCopyTokenRemoveRequest) *NullableCreditAuditCopyTokenRemoveRequest {
	return &NullableCreditAuditCopyTokenRemoveRequest{value: val, isSet: true}
}

func (v NullableCreditAuditCopyTokenRemoveRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditAuditCopyTokenRemoveRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


