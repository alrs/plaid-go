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
)

// RecurringTransactionsUpdateWebhook Fired when an Item's recurring transactions data is updated. After receipt of this webhook, the updated data can be fetched from `/transactions/recurring/get`.
type RecurringTransactionsUpdateWebhook struct {
	// `TRANSACTIONS`
	WebhookType string `json:"webhook_type"`
	// `RECURRING_TRANSACTIONS_UPDATE`
	WebhookCode string `json:"webhook_code"`
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	// A list of `account_ids` for accounts that have new or updated recurring transactions data.
	AccountIds []string `json:"account_ids"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _RecurringTransactionsUpdateWebhook RecurringTransactionsUpdateWebhook

// NewRecurringTransactionsUpdateWebhook instantiates a new RecurringTransactionsUpdateWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecurringTransactionsUpdateWebhook(webhookType string, webhookCode string, itemId string, accountIds []string, environment WebhookEnvironmentValues) *RecurringTransactionsUpdateWebhook {
	this := RecurringTransactionsUpdateWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.ItemId = itemId
	this.AccountIds = accountIds
	this.Environment = environment
	return &this
}

// NewRecurringTransactionsUpdateWebhookWithDefaults instantiates a new RecurringTransactionsUpdateWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecurringTransactionsUpdateWebhookWithDefaults() *RecurringTransactionsUpdateWebhook {
	this := RecurringTransactionsUpdateWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *RecurringTransactionsUpdateWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *RecurringTransactionsUpdateWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *RecurringTransactionsUpdateWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *RecurringTransactionsUpdateWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *RecurringTransactionsUpdateWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *RecurringTransactionsUpdateWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetItemId returns the ItemId field value
func (o *RecurringTransactionsUpdateWebhook) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *RecurringTransactionsUpdateWebhook) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *RecurringTransactionsUpdateWebhook) SetItemId(v string) {
	o.ItemId = v
}

// GetAccountIds returns the AccountIds field value
func (o *RecurringTransactionsUpdateWebhook) GetAccountIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AccountIds
}

// GetAccountIdsOk returns a tuple with the AccountIds field value
// and a boolean to check if the value has been set.
func (o *RecurringTransactionsUpdateWebhook) GetAccountIdsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountIds, true
}

// SetAccountIds sets field value
func (o *RecurringTransactionsUpdateWebhook) SetAccountIds(v []string) {
	o.AccountIds = v
}

// GetEnvironment returns the Environment field value
func (o *RecurringTransactionsUpdateWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *RecurringTransactionsUpdateWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *RecurringTransactionsUpdateWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o RecurringTransactionsUpdateWebhook) MarshalJSON() ([]byte, error) {
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
		toSerialize["account_ids"] = o.AccountIds
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecurringTransactionsUpdateWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varRecurringTransactionsUpdateWebhook := _RecurringTransactionsUpdateWebhook{}

	if err = json.Unmarshal(bytes, &varRecurringTransactionsUpdateWebhook); err == nil {
		*o = RecurringTransactionsUpdateWebhook(varRecurringTransactionsUpdateWebhook)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "account_ids")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecurringTransactionsUpdateWebhook struct {
	value *RecurringTransactionsUpdateWebhook
	isSet bool
}

func (v NullableRecurringTransactionsUpdateWebhook) Get() *RecurringTransactionsUpdateWebhook {
	return v.value
}

func (v *NullableRecurringTransactionsUpdateWebhook) Set(val *RecurringTransactionsUpdateWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurringTransactionsUpdateWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurringTransactionsUpdateWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurringTransactionsUpdateWebhook(val *RecurringTransactionsUpdateWebhook) *NullableRecurringTransactionsUpdateWebhook {
	return &NullableRecurringTransactionsUpdateWebhook{value: val, isSet: true}
}

func (v NullableRecurringTransactionsUpdateWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecurringTransactionsUpdateWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


