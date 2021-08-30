/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.26.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// NetPay An object representing information about the net pay amount on the paystub.
type NetPay struct {
	DistributionDetails *[]DistributionDetails `json:"distribution_details,omitempty"`
	Total *Total `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetPay NetPay

// NewNetPay instantiates a new NetPay object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetPay() *NetPay {
	this := NetPay{}
	return &this
}

// NewNetPayWithDefaults instantiates a new NetPay object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetPayWithDefaults() *NetPay {
	this := NetPay{}
	return &this
}

// GetDistributionDetails returns the DistributionDetails field value if set, zero value otherwise.
func (o *NetPay) GetDistributionDetails() []DistributionDetails {
	if o == nil || o.DistributionDetails == nil {
		var ret []DistributionDetails
		return ret
	}
	return *o.DistributionDetails
}

// GetDistributionDetailsOk returns a tuple with the DistributionDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetPay) GetDistributionDetailsOk() (*[]DistributionDetails, bool) {
	if o == nil || o.DistributionDetails == nil {
		return nil, false
	}
	return o.DistributionDetails, true
}

// HasDistributionDetails returns a boolean if a field has been set.
func (o *NetPay) HasDistributionDetails() bool {
	if o != nil && o.DistributionDetails != nil {
		return true
	}

	return false
}

// SetDistributionDetails gets a reference to the given []DistributionDetails and assigns it to the DistributionDetails field.
func (o *NetPay) SetDistributionDetails(v []DistributionDetails) {
	o.DistributionDetails = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *NetPay) GetTotal() Total {
	if o == nil || o.Total == nil {
		var ret Total
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetPay) GetTotalOk() (*Total, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *NetPay) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given Total and assigns it to the Total field.
func (o *NetPay) SetTotal(v Total) {
	o.Total = &v
}

func (o NetPay) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DistributionDetails != nil {
		toSerialize["distribution_details"] = o.DistributionDetails
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NetPay) UnmarshalJSON(bytes []byte) (err error) {
	varNetPay := _NetPay{}

	if err = json.Unmarshal(bytes, &varNetPay); err == nil {
		*o = NetPay(varNetPay)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "distribution_details")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetPay struct {
	value *NetPay
	isSet bool
}

func (v NullableNetPay) Get() *NetPay {
	return v.value
}

func (v *NullableNetPay) Set(val *NetPay) {
	v.value = val
	v.isSet = true
}

func (v NullableNetPay) IsSet() bool {
	return v.isSet
}

func (v *NullableNetPay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetPay(val *NetPay) *NullableNetPay {
	return &NullableNetPay{value: val, isSet: true}
}

func (v NullableNetPay) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetPay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


