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

// IdentityVerificationRetryRequestStepsObject Instructions for the `custom` retry strategy specifying which steps should be required or skipped.   Note:   This field must be provided when the retry strategy is `custom` and must be omitted otherwise.  Custom retries override settings in your Plaid Template. For example, if your Plaid Template has `verify_sms` disabled, a custom retry with `verify_sms` enabled will still require the step.  The `selfie_check` step is currently not supported on the sandbox server. Sandbox requests will silently disable the `selfie_check` step when provided.
type IdentityVerificationRetryRequestStepsObject struct {
	// A boolean field specifying whether the new session should require or skip the `verify_sms` step.
	VerifySms bool `json:"verify_sms"`
	// A boolean field specifying whether the new session should require or skip the `kyc_check` step.
	KycCheck bool `json:"kyc_check"`
	// A boolean field specifying whether the new session should require or skip the `documentary_verification` step.
	DocumentaryVerification bool `json:"documentary_verification"`
	// A boolean field specifying whether the new session should require or skip the `selfie_check` step.
	SelfieCheck bool `json:"selfie_check"`
}

// NewIdentityVerificationRetryRequestStepsObject instantiates a new IdentityVerificationRetryRequestStepsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityVerificationRetryRequestStepsObject(verifySms bool, kycCheck bool, documentaryVerification bool, selfieCheck bool) *IdentityVerificationRetryRequestStepsObject {
	this := IdentityVerificationRetryRequestStepsObject{}
	this.VerifySms = verifySms
	this.KycCheck = kycCheck
	this.DocumentaryVerification = documentaryVerification
	this.SelfieCheck = selfieCheck
	return &this
}

// NewIdentityVerificationRetryRequestStepsObjectWithDefaults instantiates a new IdentityVerificationRetryRequestStepsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityVerificationRetryRequestStepsObjectWithDefaults() *IdentityVerificationRetryRequestStepsObject {
	this := IdentityVerificationRetryRequestStepsObject{}
	return &this
}

// GetVerifySms returns the VerifySms field value
func (o *IdentityVerificationRetryRequestStepsObject) GetVerifySms() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.VerifySms
}

// GetVerifySmsOk returns a tuple with the VerifySms field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationRetryRequestStepsObject) GetVerifySmsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerifySms, true
}

// SetVerifySms sets field value
func (o *IdentityVerificationRetryRequestStepsObject) SetVerifySms(v bool) {
	o.VerifySms = v
}

// GetKycCheck returns the KycCheck field value
func (o *IdentityVerificationRetryRequestStepsObject) GetKycCheck() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.KycCheck
}

// GetKycCheckOk returns a tuple with the KycCheck field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationRetryRequestStepsObject) GetKycCheckOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.KycCheck, true
}

// SetKycCheck sets field value
func (o *IdentityVerificationRetryRequestStepsObject) SetKycCheck(v bool) {
	o.KycCheck = v
}

// GetDocumentaryVerification returns the DocumentaryVerification field value
func (o *IdentityVerificationRetryRequestStepsObject) GetDocumentaryVerification() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DocumentaryVerification
}

// GetDocumentaryVerificationOk returns a tuple with the DocumentaryVerification field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationRetryRequestStepsObject) GetDocumentaryVerificationOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocumentaryVerification, true
}

// SetDocumentaryVerification sets field value
func (o *IdentityVerificationRetryRequestStepsObject) SetDocumentaryVerification(v bool) {
	o.DocumentaryVerification = v
}

// GetSelfieCheck returns the SelfieCheck field value
func (o *IdentityVerificationRetryRequestStepsObject) GetSelfieCheck() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SelfieCheck
}

// GetSelfieCheckOk returns a tuple with the SelfieCheck field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationRetryRequestStepsObject) GetSelfieCheckOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SelfieCheck, true
}

// SetSelfieCheck sets field value
func (o *IdentityVerificationRetryRequestStepsObject) SetSelfieCheck(v bool) {
	o.SelfieCheck = v
}

func (o IdentityVerificationRetryRequestStepsObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["verify_sms"] = o.VerifySms
	}
	if true {
		toSerialize["kyc_check"] = o.KycCheck
	}
	if true {
		toSerialize["documentary_verification"] = o.DocumentaryVerification
	}
	if true {
		toSerialize["selfie_check"] = o.SelfieCheck
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityVerificationRetryRequestStepsObject struct {
	value *IdentityVerificationRetryRequestStepsObject
	isSet bool
}

func (v NullableIdentityVerificationRetryRequestStepsObject) Get() *IdentityVerificationRetryRequestStepsObject {
	return v.value
}

func (v *NullableIdentityVerificationRetryRequestStepsObject) Set(val *IdentityVerificationRetryRequestStepsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityVerificationRetryRequestStepsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityVerificationRetryRequestStepsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityVerificationRetryRequestStepsObject(val *IdentityVerificationRetryRequestStepsObject) *NullableIdentityVerificationRetryRequestStepsObject {
	return &NullableIdentityVerificationRetryRequestStepsObject{value: val, isSet: true}
}

func (v NullableIdentityVerificationRetryRequestStepsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityVerificationRetryRequestStepsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


