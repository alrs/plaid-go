/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.391.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// RiskCheckDevice Result summary object specifying values for `device` attributes of risk check.
type RiskCheckDevice struct {
	IpProxyType NullableProxyType `json:"ip_proxy_type"`
	// Count of spam lists the IP address is associated with if known.
	IpSpamListCount NullableInt32 `json:"ip_spam_list_count"`
	// UTC offset of the timezone associated with the IP address.
	IpTimezoneOffset NullableString `json:"ip_timezone_offset"`
	AdditionalProperties map[string]interface{}
}

type _RiskCheckDevice RiskCheckDevice

// NewRiskCheckDevice instantiates a new RiskCheckDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskCheckDevice(ipProxyType NullableProxyType, ipSpamListCount NullableInt32, ipTimezoneOffset NullableString) *RiskCheckDevice {
	this := RiskCheckDevice{}
	this.IpProxyType = ipProxyType
	this.IpSpamListCount = ipSpamListCount
	this.IpTimezoneOffset = ipTimezoneOffset
	return &this
}

// NewRiskCheckDeviceWithDefaults instantiates a new RiskCheckDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskCheckDeviceWithDefaults() *RiskCheckDevice {
	this := RiskCheckDevice{}
	return &this
}

// GetIpProxyType returns the IpProxyType field value
// If the value is explicit nil, the zero value for ProxyType will be returned
func (o *RiskCheckDevice) GetIpProxyType() ProxyType {
	if o == nil || o.IpProxyType.Get() == nil {
		var ret ProxyType
		return ret
	}

	return *o.IpProxyType.Get()
}

// GetIpProxyTypeOk returns a tuple with the IpProxyType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RiskCheckDevice) GetIpProxyTypeOk() (*ProxyType, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IpProxyType.Get(), o.IpProxyType.IsSet()
}

// SetIpProxyType sets field value
func (o *RiskCheckDevice) SetIpProxyType(v ProxyType) {
	o.IpProxyType.Set(&v)
}

// GetIpSpamListCount returns the IpSpamListCount field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *RiskCheckDevice) GetIpSpamListCount() int32 {
	if o == nil || o.IpSpamListCount.Get() == nil {
		var ret int32
		return ret
	}

	return *o.IpSpamListCount.Get()
}

// GetIpSpamListCountOk returns a tuple with the IpSpamListCount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RiskCheckDevice) GetIpSpamListCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IpSpamListCount.Get(), o.IpSpamListCount.IsSet()
}

// SetIpSpamListCount sets field value
func (o *RiskCheckDevice) SetIpSpamListCount(v int32) {
	o.IpSpamListCount.Set(&v)
}

// GetIpTimezoneOffset returns the IpTimezoneOffset field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RiskCheckDevice) GetIpTimezoneOffset() string {
	if o == nil || o.IpTimezoneOffset.Get() == nil {
		var ret string
		return ret
	}

	return *o.IpTimezoneOffset.Get()
}

// GetIpTimezoneOffsetOk returns a tuple with the IpTimezoneOffset field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RiskCheckDevice) GetIpTimezoneOffsetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IpTimezoneOffset.Get(), o.IpTimezoneOffset.IsSet()
}

// SetIpTimezoneOffset sets field value
func (o *RiskCheckDevice) SetIpTimezoneOffset(v string) {
	o.IpTimezoneOffset.Set(&v)
}

func (o RiskCheckDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ip_proxy_type"] = o.IpProxyType.Get()
	}
	if true {
		toSerialize["ip_spam_list_count"] = o.IpSpamListCount.Get()
	}
	if true {
		toSerialize["ip_timezone_offset"] = o.IpTimezoneOffset.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiskCheckDevice) UnmarshalJSON(bytes []byte) (err error) {
	varRiskCheckDevice := _RiskCheckDevice{}

	if err = json.Unmarshal(bytes, &varRiskCheckDevice); err == nil {
		*o = RiskCheckDevice(varRiskCheckDevice)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ip_proxy_type")
		delete(additionalProperties, "ip_spam_list_count")
		delete(additionalProperties, "ip_timezone_offset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRiskCheckDevice struct {
	value *RiskCheckDevice
	isSet bool
}

func (v NullableRiskCheckDevice) Get() *RiskCheckDevice {
	return v.value
}

func (v *NullableRiskCheckDevice) Set(val *RiskCheckDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskCheckDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskCheckDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskCheckDevice(val *RiskCheckDevice) *NullableRiskCheckDevice {
	return &NullableRiskCheckDevice{value: val, isSet: true}
}

func (v NullableRiskCheckDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskCheckDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


