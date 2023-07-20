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

// LinkDeliveryCommunicationMethod The communication method containing both the type and address to send the URL.
type LinkDeliveryCommunicationMethod struct {
	Method *LinkDeliveryDeliveryMethod `json:"method,omitempty"`
	// The phone number / email address that Hosted Link sessions are delivered to. Phone numbers must be in E.164 format.
	Address *string `json:"address,omitempty"`
}

// NewLinkDeliveryCommunicationMethod instantiates a new LinkDeliveryCommunicationMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkDeliveryCommunicationMethod() *LinkDeliveryCommunicationMethod {
	this := LinkDeliveryCommunicationMethod{}
	return &this
}

// NewLinkDeliveryCommunicationMethodWithDefaults instantiates a new LinkDeliveryCommunicationMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkDeliveryCommunicationMethodWithDefaults() *LinkDeliveryCommunicationMethod {
	this := LinkDeliveryCommunicationMethod{}
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *LinkDeliveryCommunicationMethod) GetMethod() LinkDeliveryDeliveryMethod {
	if o == nil || o.Method == nil {
		var ret LinkDeliveryDeliveryMethod
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDeliveryCommunicationMethod) GetMethodOk() (*LinkDeliveryDeliveryMethod, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *LinkDeliveryCommunicationMethod) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given LinkDeliveryDeliveryMethod and assigns it to the Method field.
func (o *LinkDeliveryCommunicationMethod) SetMethod(v LinkDeliveryDeliveryMethod) {
	o.Method = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *LinkDeliveryCommunicationMethod) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDeliveryCommunicationMethod) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *LinkDeliveryCommunicationMethod) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *LinkDeliveryCommunicationMethod) SetAddress(v string) {
	o.Address = &v
}

func (o LinkDeliveryCommunicationMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableLinkDeliveryCommunicationMethod struct {
	value *LinkDeliveryCommunicationMethod
	isSet bool
}

func (v NullableLinkDeliveryCommunicationMethod) Get() *LinkDeliveryCommunicationMethod {
	return v.value
}

func (v *NullableLinkDeliveryCommunicationMethod) Set(val *LinkDeliveryCommunicationMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkDeliveryCommunicationMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkDeliveryCommunicationMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkDeliveryCommunicationMethod(val *LinkDeliveryCommunicationMethod) *NullableLinkDeliveryCommunicationMethod {
	return &NullableLinkDeliveryCommunicationMethod{value: val, isSet: true}
}

func (v NullableLinkDeliveryCommunicationMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkDeliveryCommunicationMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


