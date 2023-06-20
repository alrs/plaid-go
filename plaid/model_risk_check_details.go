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

// RiskCheckDetails Additional information for the `risk_check` step.
type RiskCheckDetails struct {
	Status IdentityVerificationStepStatus `json:"status"`
	Behavior NullableRiskCheckBehavior `json:"behavior"`
	Email NullableRiskCheckEmail `json:"email"`
	Phone NullableRiskCheckPhone `json:"phone"`
	// Array of result summary objects specifying values for `device` attributes of risk check.
	Devices []RiskCheckDevice `json:"devices"`
	AdditionalProperties map[string]interface{}
}

type _RiskCheckDetails RiskCheckDetails

// NewRiskCheckDetails instantiates a new RiskCheckDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskCheckDetails(status IdentityVerificationStepStatus, behavior NullableRiskCheckBehavior, email NullableRiskCheckEmail, phone NullableRiskCheckPhone, devices []RiskCheckDevice) *RiskCheckDetails {
	this := RiskCheckDetails{}
	this.Status = status
	this.Behavior = behavior
	this.Email = email
	this.Phone = phone
	this.Devices = devices
	return &this
}

// NewRiskCheckDetailsWithDefaults instantiates a new RiskCheckDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskCheckDetailsWithDefaults() *RiskCheckDetails {
	this := RiskCheckDetails{}
	return &this
}

// GetStatus returns the Status field value
func (o *RiskCheckDetails) GetStatus() IdentityVerificationStepStatus {
	if o == nil {
		var ret IdentityVerificationStepStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RiskCheckDetails) GetStatusOk() (*IdentityVerificationStepStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RiskCheckDetails) SetStatus(v IdentityVerificationStepStatus) {
	o.Status = v
}

// GetBehavior returns the Behavior field value
// If the value is explicit nil, the zero value for RiskCheckBehavior will be returned
func (o *RiskCheckDetails) GetBehavior() RiskCheckBehavior {
	if o == nil || o.Behavior.Get() == nil {
		var ret RiskCheckBehavior
		return ret
	}

	return *o.Behavior.Get()
}

// GetBehaviorOk returns a tuple with the Behavior field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RiskCheckDetails) GetBehaviorOk() (*RiskCheckBehavior, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Behavior.Get(), o.Behavior.IsSet()
}

// SetBehavior sets field value
func (o *RiskCheckDetails) SetBehavior(v RiskCheckBehavior) {
	o.Behavior.Set(&v)
}

// GetEmail returns the Email field value
// If the value is explicit nil, the zero value for RiskCheckEmail will be returned
func (o *RiskCheckDetails) GetEmail() RiskCheckEmail {
	if o == nil || o.Email.Get() == nil {
		var ret RiskCheckEmail
		return ret
	}

	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RiskCheckDetails) GetEmailOk() (*RiskCheckEmail, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// SetEmail sets field value
func (o *RiskCheckDetails) SetEmail(v RiskCheckEmail) {
	o.Email.Set(&v)
}

// GetPhone returns the Phone field value
// If the value is explicit nil, the zero value for RiskCheckPhone will be returned
func (o *RiskCheckDetails) GetPhone() RiskCheckPhone {
	if o == nil || o.Phone.Get() == nil {
		var ret RiskCheckPhone
		return ret
	}

	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RiskCheckDetails) GetPhoneOk() (*RiskCheckPhone, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// SetPhone sets field value
func (o *RiskCheckDetails) SetPhone(v RiskCheckPhone) {
	o.Phone.Set(&v)
}

// GetDevices returns the Devices field value
func (o *RiskCheckDetails) GetDevices() []RiskCheckDevice {
	if o == nil {
		var ret []RiskCheckDevice
		return ret
	}

	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value
// and a boolean to check if the value has been set.
func (o *RiskCheckDetails) GetDevicesOk() (*[]RiskCheckDevice, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Devices, true
}

// SetDevices sets field value
func (o *RiskCheckDetails) SetDevices(v []RiskCheckDevice) {
	o.Devices = v
}

func (o RiskCheckDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["behavior"] = o.Behavior.Get()
	}
	if true {
		toSerialize["email"] = o.Email.Get()
	}
	if true {
		toSerialize["phone"] = o.Phone.Get()
	}
	if true {
		toSerialize["devices"] = o.Devices
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiskCheckDetails) UnmarshalJSON(bytes []byte) (err error) {
	varRiskCheckDetails := _RiskCheckDetails{}

	if err = json.Unmarshal(bytes, &varRiskCheckDetails); err == nil {
		*o = RiskCheckDetails(varRiskCheckDetails)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "behavior")
		delete(additionalProperties, "email")
		delete(additionalProperties, "phone")
		delete(additionalProperties, "devices")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRiskCheckDetails struct {
	value *RiskCheckDetails
	isSet bool
}

func (v NullableRiskCheckDetails) Get() *RiskCheckDetails {
	return v.value
}

func (v *NullableRiskCheckDetails) Set(val *RiskCheckDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskCheckDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskCheckDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskCheckDetails(val *RiskCheckDetails) *NullableRiskCheckDetails {
	return &NullableRiskCheckDetails{value: val, isSet: true}
}

func (v NullableRiskCheckDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskCheckDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


