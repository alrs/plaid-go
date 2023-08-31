/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.419.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// HostedLinkDeliveryMethod How Plaid should deliver the Plaid Link session to the customer. 'sms' will deliver via SMS. Must pass `user.phone_number`. 'email' will deliver via email. Must pass `user.email_address`. 
type HostedLinkDeliveryMethod string

var _ = fmt.Printf

// List of HostedLinkDeliveryMethod
const (
	HOSTEDLINKDELIVERYMETHOD_SMS HostedLinkDeliveryMethod = "sms"
	HOSTEDLINKDELIVERYMETHOD_EMAIL HostedLinkDeliveryMethod = "email"
)

var allowedHostedLinkDeliveryMethodEnumValues = []HostedLinkDeliveryMethod{
	"sms",
	"email",
}

func (v *HostedLinkDeliveryMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := HostedLinkDeliveryMethod(value)


	*v = enumTypeValue
	return nil
}

// NewHostedLinkDeliveryMethodFromValue returns a pointer to a valid HostedLinkDeliveryMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHostedLinkDeliveryMethodFromValue(v string) (*HostedLinkDeliveryMethod, error) {
	ev := HostedLinkDeliveryMethod(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HostedLinkDeliveryMethod) IsValid() bool {
	for _, existing := range allowedHostedLinkDeliveryMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostedLinkDeliveryMethod value
func (v HostedLinkDeliveryMethod) Ptr() *HostedLinkDeliveryMethod {
	return &v
}

type NullableHostedLinkDeliveryMethod struct {
	value *HostedLinkDeliveryMethod
	isSet bool
}

func (v NullableHostedLinkDeliveryMethod) Get() *HostedLinkDeliveryMethod {
	return v.value
}

func (v *NullableHostedLinkDeliveryMethod) Set(val *HostedLinkDeliveryMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableHostedLinkDeliveryMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableHostedLinkDeliveryMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostedLinkDeliveryMethod(val *HostedLinkDeliveryMethod) *NullableHostedLinkDeliveryMethod {
	return &NullableHostedLinkDeliveryMethod{value: val, isSet: true}
}

func (v NullableHostedLinkDeliveryMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostedLinkDeliveryMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
