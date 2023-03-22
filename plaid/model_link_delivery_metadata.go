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

// LinkDeliveryMetadata Information related to the related to the delivery of the link session to users
type LinkDeliveryMetadata struct {
	CommunicationMethod *LinkDeliveryWebhookCommunicationMethod `json:"communication_method,omitempty"`
	DeliveryStatus *LinkDeliveryWebhookDeliveryStatus `json:"delivery_status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinkDeliveryMetadata LinkDeliveryMetadata

// NewLinkDeliveryMetadata instantiates a new LinkDeliveryMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkDeliveryMetadata() *LinkDeliveryMetadata {
	this := LinkDeliveryMetadata{}
	return &this
}

// NewLinkDeliveryMetadataWithDefaults instantiates a new LinkDeliveryMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkDeliveryMetadataWithDefaults() *LinkDeliveryMetadata {
	this := LinkDeliveryMetadata{}
	return &this
}

// GetCommunicationMethod returns the CommunicationMethod field value if set, zero value otherwise.
func (o *LinkDeliveryMetadata) GetCommunicationMethod() LinkDeliveryWebhookCommunicationMethod {
	if o == nil || o.CommunicationMethod == nil {
		var ret LinkDeliveryWebhookCommunicationMethod
		return ret
	}
	return *o.CommunicationMethod
}

// GetCommunicationMethodOk returns a tuple with the CommunicationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDeliveryMetadata) GetCommunicationMethodOk() (*LinkDeliveryWebhookCommunicationMethod, bool) {
	if o == nil || o.CommunicationMethod == nil {
		return nil, false
	}
	return o.CommunicationMethod, true
}

// HasCommunicationMethod returns a boolean if a field has been set.
func (o *LinkDeliveryMetadata) HasCommunicationMethod() bool {
	if o != nil && o.CommunicationMethod != nil {
		return true
	}

	return false
}

// SetCommunicationMethod gets a reference to the given LinkDeliveryWebhookCommunicationMethod and assigns it to the CommunicationMethod field.
func (o *LinkDeliveryMetadata) SetCommunicationMethod(v LinkDeliveryWebhookCommunicationMethod) {
	o.CommunicationMethod = &v
}

// GetDeliveryStatus returns the DeliveryStatus field value if set, zero value otherwise.
func (o *LinkDeliveryMetadata) GetDeliveryStatus() LinkDeliveryWebhookDeliveryStatus {
	if o == nil || o.DeliveryStatus == nil {
		var ret LinkDeliveryWebhookDeliveryStatus
		return ret
	}
	return *o.DeliveryStatus
}

// GetDeliveryStatusOk returns a tuple with the DeliveryStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDeliveryMetadata) GetDeliveryStatusOk() (*LinkDeliveryWebhookDeliveryStatus, bool) {
	if o == nil || o.DeliveryStatus == nil {
		return nil, false
	}
	return o.DeliveryStatus, true
}

// HasDeliveryStatus returns a boolean if a field has been set.
func (o *LinkDeliveryMetadata) HasDeliveryStatus() bool {
	if o != nil && o.DeliveryStatus != nil {
		return true
	}

	return false
}

// SetDeliveryStatus gets a reference to the given LinkDeliveryWebhookDeliveryStatus and assigns it to the DeliveryStatus field.
func (o *LinkDeliveryMetadata) SetDeliveryStatus(v LinkDeliveryWebhookDeliveryStatus) {
	o.DeliveryStatus = &v
}

func (o LinkDeliveryMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CommunicationMethod != nil {
		toSerialize["communication_method"] = o.CommunicationMethod
	}
	if o.DeliveryStatus != nil {
		toSerialize["delivery_status"] = o.DeliveryStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkDeliveryMetadata) UnmarshalJSON(bytes []byte) (err error) {
	varLinkDeliveryMetadata := _LinkDeliveryMetadata{}

	if err = json.Unmarshal(bytes, &varLinkDeliveryMetadata); err == nil {
		*o = LinkDeliveryMetadata(varLinkDeliveryMetadata)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "communication_method")
		delete(additionalProperties, "delivery_status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkDeliveryMetadata struct {
	value *LinkDeliveryMetadata
	isSet bool
}

func (v NullableLinkDeliveryMetadata) Get() *LinkDeliveryMetadata {
	return v.value
}

func (v *NullableLinkDeliveryMetadata) Set(val *LinkDeliveryMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkDeliveryMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkDeliveryMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkDeliveryMetadata(val *LinkDeliveryMetadata) *NullableLinkDeliveryMetadata {
	return &NullableLinkDeliveryMetadata{value: val, isSet: true}
}

func (v NullableLinkDeliveryMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkDeliveryMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


