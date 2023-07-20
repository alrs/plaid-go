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
	"time"
)

// FDXNotification Provides the base fields of a notification. Clients will read the `type` property to determine the expected notification payload
type FDXNotification struct {
	// Id of notification
	NotificationId string `json:"notificationId"`
	Type FDXNotificationType `json:"type"`
	// ISO 8601 date-time in format 'YYYY-MM-DDThh:mm:ss.nnn[Z|[+|-]hh:mm]' according to [IETF RFC3339](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14)
	SentOn time.Time `json:"sentOn"`
	Category FDXNotificationCategory `json:"category"`
	Severity *FDXNotificationSeverity `json:"severity,omitempty"`
	Priority *FDXNotificationPriority `json:"priority,omitempty"`
	Publisher FDXParty `json:"publisher"`
	Subscriber *FDXParty `json:"subscriber,omitempty"`
	NotificationPayload FDXNotificationPayload `json:"notificationPayload"`
	Url *FDXHateoasLink `json:"url,omitempty"`
}

// NewFDXNotification instantiates a new FDXNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFDXNotification(notificationId string, type_ FDXNotificationType, sentOn time.Time, category FDXNotificationCategory, publisher FDXParty, notificationPayload FDXNotificationPayload) *FDXNotification {
	this := FDXNotification{}
	this.NotificationId = notificationId
	this.Type = type_
	this.SentOn = sentOn
	this.Category = category
	this.Publisher = publisher
	this.NotificationPayload = notificationPayload
	return &this
}

// NewFDXNotificationWithDefaults instantiates a new FDXNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFDXNotificationWithDefaults() *FDXNotification {
	this := FDXNotification{}
	return &this
}

// GetNotificationId returns the NotificationId field value
func (o *FDXNotification) GetNotificationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotificationId
}

// GetNotificationIdOk returns a tuple with the NotificationId field value
// and a boolean to check if the value has been set.
func (o *FDXNotification) GetNotificationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NotificationId, true
}

// SetNotificationId sets field value
func (o *FDXNotification) SetNotificationId(v string) {
	o.NotificationId = v
}

// GetType returns the Type field value
func (o *FDXNotification) GetType() FDXNotificationType {
	if o == nil {
		var ret FDXNotificationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FDXNotification) GetTypeOk() (*FDXNotificationType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FDXNotification) SetType(v FDXNotificationType) {
	o.Type = v
}

// GetSentOn returns the SentOn field value
func (o *FDXNotification) GetSentOn() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.SentOn
}

// GetSentOnOk returns a tuple with the SentOn field value
// and a boolean to check if the value has been set.
func (o *FDXNotification) GetSentOnOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SentOn, true
}

// SetSentOn sets field value
func (o *FDXNotification) SetSentOn(v time.Time) {
	o.SentOn = v
}

// GetCategory returns the Category field value
func (o *FDXNotification) GetCategory() FDXNotificationCategory {
	if o == nil {
		var ret FDXNotificationCategory
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *FDXNotification) GetCategoryOk() (*FDXNotificationCategory, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *FDXNotification) SetCategory(v FDXNotificationCategory) {
	o.Category = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *FDXNotification) GetSeverity() FDXNotificationSeverity {
	if o == nil || o.Severity == nil {
		var ret FDXNotificationSeverity
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FDXNotification) GetSeverityOk() (*FDXNotificationSeverity, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *FDXNotification) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given FDXNotificationSeverity and assigns it to the Severity field.
func (o *FDXNotification) SetSeverity(v FDXNotificationSeverity) {
	o.Severity = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *FDXNotification) GetPriority() FDXNotificationPriority {
	if o == nil || o.Priority == nil {
		var ret FDXNotificationPriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FDXNotification) GetPriorityOk() (*FDXNotificationPriority, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *FDXNotification) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given FDXNotificationPriority and assigns it to the Priority field.
func (o *FDXNotification) SetPriority(v FDXNotificationPriority) {
	o.Priority = &v
}

// GetPublisher returns the Publisher field value
func (o *FDXNotification) GetPublisher() FDXParty {
	if o == nil {
		var ret FDXParty
		return ret
	}

	return o.Publisher
}

// GetPublisherOk returns a tuple with the Publisher field value
// and a boolean to check if the value has been set.
func (o *FDXNotification) GetPublisherOk() (*FDXParty, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Publisher, true
}

// SetPublisher sets field value
func (o *FDXNotification) SetPublisher(v FDXParty) {
	o.Publisher = v
}

// GetSubscriber returns the Subscriber field value if set, zero value otherwise.
func (o *FDXNotification) GetSubscriber() FDXParty {
	if o == nil || o.Subscriber == nil {
		var ret FDXParty
		return ret
	}
	return *o.Subscriber
}

// GetSubscriberOk returns a tuple with the Subscriber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FDXNotification) GetSubscriberOk() (*FDXParty, bool) {
	if o == nil || o.Subscriber == nil {
		return nil, false
	}
	return o.Subscriber, true
}

// HasSubscriber returns a boolean if a field has been set.
func (o *FDXNotification) HasSubscriber() bool {
	if o != nil && o.Subscriber != nil {
		return true
	}

	return false
}

// SetSubscriber gets a reference to the given FDXParty and assigns it to the Subscriber field.
func (o *FDXNotification) SetSubscriber(v FDXParty) {
	o.Subscriber = &v
}

// GetNotificationPayload returns the NotificationPayload field value
func (o *FDXNotification) GetNotificationPayload() FDXNotificationPayload {
	if o == nil {
		var ret FDXNotificationPayload
		return ret
	}

	return o.NotificationPayload
}

// GetNotificationPayloadOk returns a tuple with the NotificationPayload field value
// and a boolean to check if the value has been set.
func (o *FDXNotification) GetNotificationPayloadOk() (*FDXNotificationPayload, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NotificationPayload, true
}

// SetNotificationPayload sets field value
func (o *FDXNotification) SetNotificationPayload(v FDXNotificationPayload) {
	o.NotificationPayload = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *FDXNotification) GetUrl() FDXHateoasLink {
	if o == nil || o.Url == nil {
		var ret FDXHateoasLink
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FDXNotification) GetUrlOk() (*FDXHateoasLink, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *FDXNotification) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given FDXHateoasLink and assigns it to the Url field.
func (o *FDXNotification) SetUrl(v FDXHateoasLink) {
	o.Url = &v
}

func (o FDXNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["notificationId"] = o.NotificationId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["sentOn"] = o.SentOn
	}
	if true {
		toSerialize["category"] = o.Category
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if true {
		toSerialize["publisher"] = o.Publisher
	}
	if o.Subscriber != nil {
		toSerialize["subscriber"] = o.Subscriber
	}
	if true {
		toSerialize["notificationPayload"] = o.NotificationPayload
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableFDXNotification struct {
	value *FDXNotification
	isSet bool
}

func (v NullableFDXNotification) Get() *FDXNotification {
	return v.value
}

func (v *NullableFDXNotification) Set(val *FDXNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableFDXNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableFDXNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFDXNotification(val *FDXNotification) *NullableFDXNotification {
	return &NullableFDXNotification{value: val, isSet: true}
}

func (v NullableFDXNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFDXNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


