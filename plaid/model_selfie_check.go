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

// SelfieCheck Additional information for the `selfie_check` step. This field will be `null` unless `steps.selfie_check` has reached a terminal state of either `success` or `failed`.
type SelfieCheck struct {
	Status SelfieCheckStatus `json:"status"`
	// An array of selfies submitted to the `selfie_check` step. Each entry represents one user submission.
	Selfies []SelfieCheckSelfie `json:"selfies"`
	AdditionalProperties map[string]interface{}
}

type _SelfieCheck SelfieCheck

// NewSelfieCheck instantiates a new SelfieCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfieCheck(status SelfieCheckStatus, selfies []SelfieCheckSelfie) *SelfieCheck {
	this := SelfieCheck{}
	this.Status = status
	this.Selfies = selfies
	return &this
}

// NewSelfieCheckWithDefaults instantiates a new SelfieCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfieCheckWithDefaults() *SelfieCheck {
	this := SelfieCheck{}
	return &this
}

// GetStatus returns the Status field value
func (o *SelfieCheck) GetStatus() SelfieCheckStatus {
	if o == nil {
		var ret SelfieCheckStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SelfieCheck) GetStatusOk() (*SelfieCheckStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SelfieCheck) SetStatus(v SelfieCheckStatus) {
	o.Status = v
}

// GetSelfies returns the Selfies field value
func (o *SelfieCheck) GetSelfies() []SelfieCheckSelfie {
	if o == nil {
		var ret []SelfieCheckSelfie
		return ret
	}

	return o.Selfies
}

// GetSelfiesOk returns a tuple with the Selfies field value
// and a boolean to check if the value has been set.
func (o *SelfieCheck) GetSelfiesOk() (*[]SelfieCheckSelfie, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Selfies, true
}

// SetSelfies sets field value
func (o *SelfieCheck) SetSelfies(v []SelfieCheckSelfie) {
	o.Selfies = v
}

func (o SelfieCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["selfies"] = o.Selfies
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SelfieCheck) UnmarshalJSON(bytes []byte) (err error) {
	varSelfieCheck := _SelfieCheck{}

	if err = json.Unmarshal(bytes, &varSelfieCheck); err == nil {
		*o = SelfieCheck(varSelfieCheck)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "selfies")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSelfieCheck struct {
	value *SelfieCheck
	isSet bool
}

func (v NullableSelfieCheck) Get() *SelfieCheck {
	return v.value
}

func (v *NullableSelfieCheck) Set(val *SelfieCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfieCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfieCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfieCheck(val *SelfieCheck) *NullableSelfieCheck {
	return &NullableSelfieCheck{value: val, isSet: true}
}

func (v NullableSelfieCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfieCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


