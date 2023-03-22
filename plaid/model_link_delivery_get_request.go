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

// LinkDeliveryGetRequest LinkDeliveryGetRequest defines the request schema for `/link_delivery/get`
type LinkDeliveryGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The ID for the Link Delivery session from a previous invocation of `/link_delivery/create`.
	LinkDeliverySessionId string `json:"link_delivery_session_id"`
}

// NewLinkDeliveryGetRequest instantiates a new LinkDeliveryGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkDeliveryGetRequest(linkDeliverySessionId string) *LinkDeliveryGetRequest {
	this := LinkDeliveryGetRequest{}
	this.LinkDeliverySessionId = linkDeliverySessionId
	return &this
}

// NewLinkDeliveryGetRequestWithDefaults instantiates a new LinkDeliveryGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkDeliveryGetRequestWithDefaults() *LinkDeliveryGetRequest {
	this := LinkDeliveryGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *LinkDeliveryGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDeliveryGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *LinkDeliveryGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *LinkDeliveryGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *LinkDeliveryGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDeliveryGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *LinkDeliveryGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *LinkDeliveryGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetLinkDeliverySessionId returns the LinkDeliverySessionId field value
func (o *LinkDeliveryGetRequest) GetLinkDeliverySessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkDeliverySessionId
}

// GetLinkDeliverySessionIdOk returns a tuple with the LinkDeliverySessionId field value
// and a boolean to check if the value has been set.
func (o *LinkDeliveryGetRequest) GetLinkDeliverySessionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LinkDeliverySessionId, true
}

// SetLinkDeliverySessionId sets field value
func (o *LinkDeliveryGetRequest) SetLinkDeliverySessionId(v string) {
	o.LinkDeliverySessionId = v
}

func (o LinkDeliveryGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["link_delivery_session_id"] = o.LinkDeliverySessionId
	}
	return json.Marshal(toSerialize)
}

type NullableLinkDeliveryGetRequest struct {
	value *LinkDeliveryGetRequest
	isSet bool
}

func (v NullableLinkDeliveryGetRequest) Get() *LinkDeliveryGetRequest {
	return v.value
}

func (v *NullableLinkDeliveryGetRequest) Set(val *LinkDeliveryGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkDeliveryGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkDeliveryGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkDeliveryGetRequest(val *LinkDeliveryGetRequest) *NullableLinkDeliveryGetRequest {
	return &NullableLinkDeliveryGetRequest{value: val, isSet: true}
}

func (v NullableLinkDeliveryGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkDeliveryGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


