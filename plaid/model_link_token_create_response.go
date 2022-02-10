/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.64.14
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"time"
)

// LinkTokenCreateResponse LinkTokenCreateResponse defines the response schema for `/link/token/create`
type LinkTokenCreateResponse struct {
	// A `link_token`, which can be supplied to Link in order to initialize it and receive a `public_token`, which can be exchanged for an `access_token`.
	LinkToken string `json:"link_token"`
	// The expiration date for the `link_token`, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format. A `link_token` created to generate a `public_token` that will be exchanged for a new `access_token` expires after 4 hours. A `link_token` created for an existing Item (such as when updating an existing `access_token` by launching Link in update mode) expires after 30 minutes.
	Expiration time.Time `json:"expiration"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _LinkTokenCreateResponse LinkTokenCreateResponse

// NewLinkTokenCreateResponse instantiates a new LinkTokenCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenCreateResponse(linkToken string, expiration time.Time, requestId string) *LinkTokenCreateResponse {
	this := LinkTokenCreateResponse{}
	this.LinkToken = linkToken
	this.Expiration = expiration
	this.RequestId = requestId
	return &this
}

// NewLinkTokenCreateResponseWithDefaults instantiates a new LinkTokenCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenCreateResponseWithDefaults() *LinkTokenCreateResponse {
	this := LinkTokenCreateResponse{}
	return &this
}

// GetLinkToken returns the LinkToken field value
func (o *LinkTokenCreateResponse) GetLinkToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkToken
}

// GetLinkTokenOk returns a tuple with the LinkToken field value
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateResponse) GetLinkTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LinkToken, true
}

// SetLinkToken sets field value
func (o *LinkTokenCreateResponse) SetLinkToken(v string) {
	o.LinkToken = v
}

// GetExpiration returns the Expiration field value
func (o *LinkTokenCreateResponse) GetExpiration() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateResponse) GetExpirationOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Expiration, true
}

// SetExpiration sets field value
func (o *LinkTokenCreateResponse) SetExpiration(v time.Time) {
	o.Expiration = v
}

// GetRequestId returns the RequestId field value
func (o *LinkTokenCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *LinkTokenCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o LinkTokenCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["link_token"] = o.LinkToken
	}
	if true {
		toSerialize["expiration"] = o.Expiration
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkTokenCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varLinkTokenCreateResponse := _LinkTokenCreateResponse{}

	if err = json.Unmarshal(bytes, &varLinkTokenCreateResponse); err == nil {
		*o = LinkTokenCreateResponse(varLinkTokenCreateResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "link_token")
		delete(additionalProperties, "expiration")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkTokenCreateResponse struct {
	value *LinkTokenCreateResponse
	isSet bool
}

func (v NullableLinkTokenCreateResponse) Get() *LinkTokenCreateResponse {
	return v.value
}

func (v *NullableLinkTokenCreateResponse) Set(val *LinkTokenCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenCreateResponse(val *LinkTokenCreateResponse) *NullableLinkTokenCreateResponse {
	return &NullableLinkTokenCreateResponse{value: val, isSet: true}
}

func (v NullableLinkTokenCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


