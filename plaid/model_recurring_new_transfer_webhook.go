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
)

// RecurringNewTransferWebhook Fired when a new transfer of a recurring transfer is originated.
type RecurringNewTransferWebhook struct {
	// `TRANSFER`
	WebhookType string `json:"webhook_type"`
	// `RECURRING_NEW_TRANSFER`
	WebhookCode string `json:"webhook_code"`
	// Plaid’s unique identifier for a recurring transfer.
	RecurringTransferId string `json:"recurring_transfer_id"`
	// Plaid’s unique identifier for a transfer.
	TransferId string `json:"transfer_id"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _RecurringNewTransferWebhook RecurringNewTransferWebhook

// NewRecurringNewTransferWebhook instantiates a new RecurringNewTransferWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecurringNewTransferWebhook(webhookType string, webhookCode string, recurringTransferId string, transferId string, environment WebhookEnvironmentValues) *RecurringNewTransferWebhook {
	this := RecurringNewTransferWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.RecurringTransferId = recurringTransferId
	this.TransferId = transferId
	this.Environment = environment
	return &this
}

// NewRecurringNewTransferWebhookWithDefaults instantiates a new RecurringNewTransferWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecurringNewTransferWebhookWithDefaults() *RecurringNewTransferWebhook {
	this := RecurringNewTransferWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *RecurringNewTransferWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *RecurringNewTransferWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *RecurringNewTransferWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *RecurringNewTransferWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *RecurringNewTransferWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *RecurringNewTransferWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetRecurringTransferId returns the RecurringTransferId field value
func (o *RecurringNewTransferWebhook) GetRecurringTransferId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecurringTransferId
}

// GetRecurringTransferIdOk returns a tuple with the RecurringTransferId field value
// and a boolean to check if the value has been set.
func (o *RecurringNewTransferWebhook) GetRecurringTransferIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecurringTransferId, true
}

// SetRecurringTransferId sets field value
func (o *RecurringNewTransferWebhook) SetRecurringTransferId(v string) {
	o.RecurringTransferId = v
}

// GetTransferId returns the TransferId field value
func (o *RecurringNewTransferWebhook) GetTransferId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransferId
}

// GetTransferIdOk returns a tuple with the TransferId field value
// and a boolean to check if the value has been set.
func (o *RecurringNewTransferWebhook) GetTransferIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransferId, true
}

// SetTransferId sets field value
func (o *RecurringNewTransferWebhook) SetTransferId(v string) {
	o.TransferId = v
}

// GetEnvironment returns the Environment field value
func (o *RecurringNewTransferWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *RecurringNewTransferWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *RecurringNewTransferWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o RecurringNewTransferWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["recurring_transfer_id"] = o.RecurringTransferId
	}
	if true {
		toSerialize["transfer_id"] = o.TransferId
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecurringNewTransferWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varRecurringNewTransferWebhook := _RecurringNewTransferWebhook{}

	if err = json.Unmarshal(bytes, &varRecurringNewTransferWebhook); err == nil {
		*o = RecurringNewTransferWebhook(varRecurringNewTransferWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "recurring_transfer_id")
		delete(additionalProperties, "transfer_id")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecurringNewTransferWebhook struct {
	value *RecurringNewTransferWebhook
	isSet bool
}

func (v NullableRecurringNewTransferWebhook) Get() *RecurringNewTransferWebhook {
	return v.value
}

func (v *NullableRecurringNewTransferWebhook) Set(val *RecurringNewTransferWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurringNewTransferWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurringNewTransferWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurringNewTransferWebhook(val *RecurringNewTransferWebhook) *NullableRecurringNewTransferWebhook {
	return &NullableRecurringNewTransferWebhook{value: val, isSet: true}
}

func (v NullableRecurringNewTransferWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecurringNewTransferWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


