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
	"fmt"
)

// LinkDeliveryDeliveryMethod The delivery method to be used to deliver the Link Delivery URL.  `SMS`: The URL will be delivered through SMS  `EMAIL`: The URL will be delivered through email
type LinkDeliveryDeliveryMethod string

// List of LinkDeliveryDeliveryMethod
const (
	LINKDELIVERYDELIVERYMETHOD_SMS LinkDeliveryDeliveryMethod = "SMS"
	LINKDELIVERYDELIVERYMETHOD_EMAIL LinkDeliveryDeliveryMethod = "EMAIL"
)

var allowedLinkDeliveryDeliveryMethodEnumValues = []LinkDeliveryDeliveryMethod{
	"SMS",
	"EMAIL",
}

func (v *LinkDeliveryDeliveryMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LinkDeliveryDeliveryMethod(value)
	for _, existing := range allowedLinkDeliveryDeliveryMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LinkDeliveryDeliveryMethod", value)
}

// NewLinkDeliveryDeliveryMethodFromValue returns a pointer to a valid LinkDeliveryDeliveryMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLinkDeliveryDeliveryMethodFromValue(v string) (*LinkDeliveryDeliveryMethod, error) {
	ev := LinkDeliveryDeliveryMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LinkDeliveryDeliveryMethod: valid values are %v", v, allowedLinkDeliveryDeliveryMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LinkDeliveryDeliveryMethod) IsValid() bool {
	for _, existing := range allowedLinkDeliveryDeliveryMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LinkDeliveryDeliveryMethod value
func (v LinkDeliveryDeliveryMethod) Ptr() *LinkDeliveryDeliveryMethod {
	return &v
}

type NullableLinkDeliveryDeliveryMethod struct {
	value *LinkDeliveryDeliveryMethod
	isSet bool
}

func (v NullableLinkDeliveryDeliveryMethod) Get() *LinkDeliveryDeliveryMethod {
	return v.value
}

func (v *NullableLinkDeliveryDeliveryMethod) Set(val *LinkDeliveryDeliveryMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkDeliveryDeliveryMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkDeliveryDeliveryMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkDeliveryDeliveryMethod(val *LinkDeliveryDeliveryMethod) *NullableLinkDeliveryDeliveryMethod {
	return &NullableLinkDeliveryDeliveryMethod{value: val, isSet: true}
}

func (v NullableLinkDeliveryDeliveryMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkDeliveryDeliveryMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

