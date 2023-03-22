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

// SyncUpdatesAvailableWebhook Fired when an Item's transactions change. This can be due to any event resulting in new changes, such as an initial 30-day transactions fetch upon the initialization of an Item with transactions, the backfill of historical transactions that occurs shortly after, or when changes are populated from a regularly-scheduled transactions update job. It is recommended to listen for the `SYNC_UPDATES_AVAILABLE` webhook when using the `/transactions/sync` endpoint. Note that when using `/transactions/sync` the older webhooks `INITIAL_UPDATE`, `HISTORICAL_UPDATE`, `DEFAULT_UPDATE`, and `TRANSACTIONS_REMOVED`, which are intended for use with `/transactions/get`, will also continue to be sent in order to maintain backwards compatibility. It is not necessary to listen for and respond to those webhooks when using `/transactions/sync`.  After receipt of this webhook, the new changes can be fetched for the Item from `/transactions/sync`.  Note that to receive this webhook for an Item, `/transactions/sync` must have been called at least once on that Item. This means that, unlike the `INITIAL_UPDATE` and `HISTORICAL_UPDATE` webhooks, it will not fire immediately upon Item creation. If `/transactions/sync` is called on an Item that was *not* initialized with Transactions, the webhook will fire twice: once the first 30 days of transactions data has been fetched, and a second time when all available historical transactions data has been fetched.  This webhook will typically not fire in the Sandbox environment, due to the lack of dynamic transactions data. To test this webhook in Sandbox, call `/sandbox/item/fire_webhook`.
type SyncUpdatesAvailableWebhook struct {
	// `TRANSACTIONS`
	WebhookType string `json:"webhook_type"`
	// `SYNC_UPDATES_AVAILABLE`
	WebhookCode string `json:"webhook_code"`
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	// Indicates if initial pull information is available.
	InitialUpdateComplete bool `json:"initial_update_complete"`
	// Indicates if historical pull information is available.
	HistoricalUpdateComplete bool `json:"historical_update_complete"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _SyncUpdatesAvailableWebhook SyncUpdatesAvailableWebhook

// NewSyncUpdatesAvailableWebhook instantiates a new SyncUpdatesAvailableWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncUpdatesAvailableWebhook(webhookType string, webhookCode string, itemId string, initialUpdateComplete bool, historicalUpdateComplete bool, environment WebhookEnvironmentValues) *SyncUpdatesAvailableWebhook {
	this := SyncUpdatesAvailableWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.ItemId = itemId
	this.InitialUpdateComplete = initialUpdateComplete
	this.HistoricalUpdateComplete = historicalUpdateComplete
	this.Environment = environment
	return &this
}

// NewSyncUpdatesAvailableWebhookWithDefaults instantiates a new SyncUpdatesAvailableWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncUpdatesAvailableWebhookWithDefaults() *SyncUpdatesAvailableWebhook {
	this := SyncUpdatesAvailableWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *SyncUpdatesAvailableWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *SyncUpdatesAvailableWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *SyncUpdatesAvailableWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *SyncUpdatesAvailableWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *SyncUpdatesAvailableWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *SyncUpdatesAvailableWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetItemId returns the ItemId field value
func (o *SyncUpdatesAvailableWebhook) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *SyncUpdatesAvailableWebhook) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *SyncUpdatesAvailableWebhook) SetItemId(v string) {
	o.ItemId = v
}

// GetInitialUpdateComplete returns the InitialUpdateComplete field value
func (o *SyncUpdatesAvailableWebhook) GetInitialUpdateComplete() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InitialUpdateComplete
}

// GetInitialUpdateCompleteOk returns a tuple with the InitialUpdateComplete field value
// and a boolean to check if the value has been set.
func (o *SyncUpdatesAvailableWebhook) GetInitialUpdateCompleteOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InitialUpdateComplete, true
}

// SetInitialUpdateComplete sets field value
func (o *SyncUpdatesAvailableWebhook) SetInitialUpdateComplete(v bool) {
	o.InitialUpdateComplete = v
}

// GetHistoricalUpdateComplete returns the HistoricalUpdateComplete field value
func (o *SyncUpdatesAvailableWebhook) GetHistoricalUpdateComplete() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HistoricalUpdateComplete
}

// GetHistoricalUpdateCompleteOk returns a tuple with the HistoricalUpdateComplete field value
// and a boolean to check if the value has been set.
func (o *SyncUpdatesAvailableWebhook) GetHistoricalUpdateCompleteOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HistoricalUpdateComplete, true
}

// SetHistoricalUpdateComplete sets field value
func (o *SyncUpdatesAvailableWebhook) SetHistoricalUpdateComplete(v bool) {
	o.HistoricalUpdateComplete = v
}

// GetEnvironment returns the Environment field value
func (o *SyncUpdatesAvailableWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *SyncUpdatesAvailableWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *SyncUpdatesAvailableWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o SyncUpdatesAvailableWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["item_id"] = o.ItemId
	}
	if true {
		toSerialize["initial_update_complete"] = o.InitialUpdateComplete
	}
	if true {
		toSerialize["historical_update_complete"] = o.HistoricalUpdateComplete
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SyncUpdatesAvailableWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varSyncUpdatesAvailableWebhook := _SyncUpdatesAvailableWebhook{}

	if err = json.Unmarshal(bytes, &varSyncUpdatesAvailableWebhook); err == nil {
		*o = SyncUpdatesAvailableWebhook(varSyncUpdatesAvailableWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "initial_update_complete")
		delete(additionalProperties, "historical_update_complete")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSyncUpdatesAvailableWebhook struct {
	value *SyncUpdatesAvailableWebhook
	isSet bool
}

func (v NullableSyncUpdatesAvailableWebhook) Get() *SyncUpdatesAvailableWebhook {
	return v.value
}

func (v *NullableSyncUpdatesAvailableWebhook) Set(val *SyncUpdatesAvailableWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncUpdatesAvailableWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncUpdatesAvailableWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncUpdatesAvailableWebhook(val *SyncUpdatesAvailableWebhook) *NullableSyncUpdatesAvailableWebhook {
	return &NullableSyncUpdatesAvailableWebhook{value: val, isSet: true}
}

func (v NullableSyncUpdatesAvailableWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncUpdatesAvailableWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


