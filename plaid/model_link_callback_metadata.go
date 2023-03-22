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

// LinkCallbackMetadata Information related to the callback from the hosted Link session.
type LinkCallbackMetadata struct {
	CallbackType *LinkDeliveryWebhookCallbackType `json:"callback_type,omitempty"`
	EventName *LinkEventName `json:"event_name,omitempty"`
	// Indicates where in the flow the Link user exited
	Status *string `json:"status,omitempty"`
	// A unique identifier associated with a user's actions and events through the Link flow. Include this identifier when opening a support ticket for faster turnaround.
	LinkSessionId *string `json:"link_session_id,omitempty"`
	// The request ID for the last request made by Link. This can be shared with Plaid Support to expedite investigation.
	RequestId *string `json:"request_id,omitempty"`
	Institution *LinkDeliveryInstitution `json:"institution,omitempty"`
	// A list of accounts attached to the connected Item. If Account Select is enabled via the developer dashboard, accounts will only include selected accounts.
	Accounts *[]LinkDeliveryAccount `json:"accounts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinkCallbackMetadata LinkCallbackMetadata

// NewLinkCallbackMetadata instantiates a new LinkCallbackMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkCallbackMetadata() *LinkCallbackMetadata {
	this := LinkCallbackMetadata{}
	return &this
}

// NewLinkCallbackMetadataWithDefaults instantiates a new LinkCallbackMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkCallbackMetadataWithDefaults() *LinkCallbackMetadata {
	this := LinkCallbackMetadata{}
	return &this
}

// GetCallbackType returns the CallbackType field value if set, zero value otherwise.
func (o *LinkCallbackMetadata) GetCallbackType() LinkDeliveryWebhookCallbackType {
	if o == nil || o.CallbackType == nil {
		var ret LinkDeliveryWebhookCallbackType
		return ret
	}
	return *o.CallbackType
}

// GetCallbackTypeOk returns a tuple with the CallbackType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkCallbackMetadata) GetCallbackTypeOk() (*LinkDeliveryWebhookCallbackType, bool) {
	if o == nil || o.CallbackType == nil {
		return nil, false
	}
	return o.CallbackType, true
}

// HasCallbackType returns a boolean if a field has been set.
func (o *LinkCallbackMetadata) HasCallbackType() bool {
	if o != nil && o.CallbackType != nil {
		return true
	}

	return false
}

// SetCallbackType gets a reference to the given LinkDeliveryWebhookCallbackType and assigns it to the CallbackType field.
func (o *LinkCallbackMetadata) SetCallbackType(v LinkDeliveryWebhookCallbackType) {
	o.CallbackType = &v
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *LinkCallbackMetadata) GetEventName() LinkEventName {
	if o == nil || o.EventName == nil {
		var ret LinkEventName
		return ret
	}
	return *o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkCallbackMetadata) GetEventNameOk() (*LinkEventName, bool) {
	if o == nil || o.EventName == nil {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *LinkCallbackMetadata) HasEventName() bool {
	if o != nil && o.EventName != nil {
		return true
	}

	return false
}

// SetEventName gets a reference to the given LinkEventName and assigns it to the EventName field.
func (o *LinkCallbackMetadata) SetEventName(v LinkEventName) {
	o.EventName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LinkCallbackMetadata) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkCallbackMetadata) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LinkCallbackMetadata) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LinkCallbackMetadata) SetStatus(v string) {
	o.Status = &v
}

// GetLinkSessionId returns the LinkSessionId field value if set, zero value otherwise.
func (o *LinkCallbackMetadata) GetLinkSessionId() string {
	if o == nil || o.LinkSessionId == nil {
		var ret string
		return ret
	}
	return *o.LinkSessionId
}

// GetLinkSessionIdOk returns a tuple with the LinkSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkCallbackMetadata) GetLinkSessionIdOk() (*string, bool) {
	if o == nil || o.LinkSessionId == nil {
		return nil, false
	}
	return o.LinkSessionId, true
}

// HasLinkSessionId returns a boolean if a field has been set.
func (o *LinkCallbackMetadata) HasLinkSessionId() bool {
	if o != nil && o.LinkSessionId != nil {
		return true
	}

	return false
}

// SetLinkSessionId gets a reference to the given string and assigns it to the LinkSessionId field.
func (o *LinkCallbackMetadata) SetLinkSessionId(v string) {
	o.LinkSessionId = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *LinkCallbackMetadata) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkCallbackMetadata) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *LinkCallbackMetadata) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *LinkCallbackMetadata) SetRequestId(v string) {
	o.RequestId = &v
}

// GetInstitution returns the Institution field value if set, zero value otherwise.
func (o *LinkCallbackMetadata) GetInstitution() LinkDeliveryInstitution {
	if o == nil || o.Institution == nil {
		var ret LinkDeliveryInstitution
		return ret
	}
	return *o.Institution
}

// GetInstitutionOk returns a tuple with the Institution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkCallbackMetadata) GetInstitutionOk() (*LinkDeliveryInstitution, bool) {
	if o == nil || o.Institution == nil {
		return nil, false
	}
	return o.Institution, true
}

// HasInstitution returns a boolean if a field has been set.
func (o *LinkCallbackMetadata) HasInstitution() bool {
	if o != nil && o.Institution != nil {
		return true
	}

	return false
}

// SetInstitution gets a reference to the given LinkDeliveryInstitution and assigns it to the Institution field.
func (o *LinkCallbackMetadata) SetInstitution(v LinkDeliveryInstitution) {
	o.Institution = &v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *LinkCallbackMetadata) GetAccounts() []LinkDeliveryAccount {
	if o == nil || o.Accounts == nil {
		var ret []LinkDeliveryAccount
		return ret
	}
	return *o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkCallbackMetadata) GetAccountsOk() (*[]LinkDeliveryAccount, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *LinkCallbackMetadata) HasAccounts() bool {
	if o != nil && o.Accounts != nil {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []LinkDeliveryAccount and assigns it to the Accounts field.
func (o *LinkCallbackMetadata) SetAccounts(v []LinkDeliveryAccount) {
	o.Accounts = &v
}

func (o LinkCallbackMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CallbackType != nil {
		toSerialize["callback_type"] = o.CallbackType
	}
	if o.EventName != nil {
		toSerialize["event_name"] = o.EventName
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.LinkSessionId != nil {
		toSerialize["link_session_id"] = o.LinkSessionId
	}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}
	if o.Institution != nil {
		toSerialize["institution"] = o.Institution
	}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkCallbackMetadata) UnmarshalJSON(bytes []byte) (err error) {
	varLinkCallbackMetadata := _LinkCallbackMetadata{}

	if err = json.Unmarshal(bytes, &varLinkCallbackMetadata); err == nil {
		*o = LinkCallbackMetadata(varLinkCallbackMetadata)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "callback_type")
		delete(additionalProperties, "event_name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "link_session_id")
		delete(additionalProperties, "request_id")
		delete(additionalProperties, "institution")
		delete(additionalProperties, "accounts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkCallbackMetadata struct {
	value *LinkCallbackMetadata
	isSet bool
}

func (v NullableLinkCallbackMetadata) Get() *LinkCallbackMetadata {
	return v.value
}

func (v *NullableLinkCallbackMetadata) Set(val *LinkCallbackMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkCallbackMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkCallbackMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkCallbackMetadata(val *LinkCallbackMetadata) *NullableLinkCallbackMetadata {
	return &NullableLinkCallbackMetadata{value: val, isSet: true}
}

func (v NullableLinkCallbackMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkCallbackMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


