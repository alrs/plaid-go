/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.205.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"time"
)

// LinkDeliveryGetResponse LinkDeliveryGetRequest defines the response schema for `/link_delivery/get`
type LinkDeliveryGetResponse struct {
	Status LinkDeliverySessionStatus `json:"status"`
	// Timestamp in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (`YYYY-MM-DDTHH:mm:ssZ`) indicating the time the given Link Delivery Session was created at
	CreatedAt time.Time `json:"created_at"`
	// The public tokens returned by the Link session upon completion
	PublicTokens *[]string `json:"public_tokens,omitempty"`
	// Timestamp in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (`YYYY-MM-DDTHH:mm:ssZ`) indicating the time the given Link Delivery Session was completed at
	CompletedAt NullableTime `json:"completed_at,omitempty"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _LinkDeliveryGetResponse LinkDeliveryGetResponse

// NewLinkDeliveryGetResponse instantiates a new LinkDeliveryGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkDeliveryGetResponse(status LinkDeliverySessionStatus, createdAt time.Time, requestId string) *LinkDeliveryGetResponse {
	this := LinkDeliveryGetResponse{}
	this.Status = status
	this.CreatedAt = createdAt
	this.RequestId = requestId
	return &this
}

// NewLinkDeliveryGetResponseWithDefaults instantiates a new LinkDeliveryGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkDeliveryGetResponseWithDefaults() *LinkDeliveryGetResponse {
	this := LinkDeliveryGetResponse{}
	return &this
}

// GetStatus returns the Status field value
func (o *LinkDeliveryGetResponse) GetStatus() LinkDeliverySessionStatus {
	if o == nil {
		var ret LinkDeliverySessionStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LinkDeliveryGetResponse) GetStatusOk() (*LinkDeliverySessionStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *LinkDeliveryGetResponse) SetStatus(v LinkDeliverySessionStatus) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *LinkDeliveryGetResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LinkDeliveryGetResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *LinkDeliveryGetResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetPublicTokens returns the PublicTokens field value if set, zero value otherwise.
func (o *LinkDeliveryGetResponse) GetPublicTokens() []string {
	if o == nil || o.PublicTokens == nil {
		var ret []string
		return ret
	}
	return *o.PublicTokens
}

// GetPublicTokensOk returns a tuple with the PublicTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDeliveryGetResponse) GetPublicTokensOk() (*[]string, bool) {
	if o == nil || o.PublicTokens == nil {
		return nil, false
	}
	return o.PublicTokens, true
}

// HasPublicTokens returns a boolean if a field has been set.
func (o *LinkDeliveryGetResponse) HasPublicTokens() bool {
	if o != nil && o.PublicTokens != nil {
		return true
	}

	return false
}

// SetPublicTokens gets a reference to the given []string and assigns it to the PublicTokens field.
func (o *LinkDeliveryGetResponse) SetPublicTokens(v []string) {
	o.PublicTokens = &v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LinkDeliveryGetResponse) GetCompletedAt() time.Time {
	if o == nil || o.CompletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkDeliveryGetResponse) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *LinkDeliveryGetResponse) HasCompletedAt() bool {
	if o != nil && o.CompletedAt.IsSet() {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given NullableTime and assigns it to the CompletedAt field.
func (o *LinkDeliveryGetResponse) SetCompletedAt(v time.Time) {
	o.CompletedAt.Set(&v)
}
// SetCompletedAtNil sets the value for CompletedAt to be an explicit nil
func (o *LinkDeliveryGetResponse) SetCompletedAtNil() {
	o.CompletedAt.Set(nil)
}

// UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
func (o *LinkDeliveryGetResponse) UnsetCompletedAt() {
	o.CompletedAt.Unset()
}

// GetRequestId returns the RequestId field value
func (o *LinkDeliveryGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *LinkDeliveryGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *LinkDeliveryGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o LinkDeliveryGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.PublicTokens != nil {
		toSerialize["public_tokens"] = o.PublicTokens
	}
	if o.CompletedAt.IsSet() {
		toSerialize["completed_at"] = o.CompletedAt.Get()
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkDeliveryGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varLinkDeliveryGetResponse := _LinkDeliveryGetResponse{}

	if err = json.Unmarshal(bytes, &varLinkDeliveryGetResponse); err == nil {
		*o = LinkDeliveryGetResponse(varLinkDeliveryGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "public_tokens")
		delete(additionalProperties, "completed_at")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkDeliveryGetResponse struct {
	value *LinkDeliveryGetResponse
	isSet bool
}

func (v NullableLinkDeliveryGetResponse) Get() *LinkDeliveryGetResponse {
	return v.value
}

func (v *NullableLinkDeliveryGetResponse) Set(val *LinkDeliveryGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkDeliveryGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkDeliveryGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkDeliveryGetResponse(val *LinkDeliveryGetResponse) *NullableLinkDeliveryGetResponse {
	return &NullableLinkDeliveryGetResponse{value: val, isSet: true}
}

func (v NullableLinkDeliveryGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkDeliveryGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


