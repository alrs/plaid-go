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

// BankTransfersEventsUpdateWebhookForAuth Fired when new ACH events are available. To begin receiving this webhook, you must first register your webhook listener endpoint via the [webhooks page in the Dashboard](https://dashboard.plaid.com/team/webhooks). The `BANK_TRANSFERS_EVENTS_UPDATE` webhook can be used to track the progress of ACH transfers used in [micro-deposit verification](/docs/auth/coverage/microdeposit-events/). Receiving this webhook indicates you should fetch the new events from `/bank_transfer/event/sync`. Note that [Transfer](https://plaid.com/docs/transfer) customers should use Transfer webhooks instead of using `BANK_TRANSFERS_EVENTS_UPDATE`; see [micro-deposit events documentation](https://plaid.com/docs/auth/coverage/microdeposit-events/) for more details.
type BankTransfersEventsUpdateWebhookForAuth struct {
	// `BANK_TRANSFERS`
	WebhookType string `json:"webhook_type"`
	// `BANK_TRANSFERS_EVENTS_UPDATE`
	WebhookCode string `json:"webhook_code"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _BankTransfersEventsUpdateWebhookForAuth BankTransfersEventsUpdateWebhookForAuth

// NewBankTransfersEventsUpdateWebhookForAuth instantiates a new BankTransfersEventsUpdateWebhookForAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransfersEventsUpdateWebhookForAuth(webhookType string, webhookCode string, environment WebhookEnvironmentValues) *BankTransfersEventsUpdateWebhookForAuth {
	this := BankTransfersEventsUpdateWebhookForAuth{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.Environment = environment
	return &this
}

// NewBankTransfersEventsUpdateWebhookForAuthWithDefaults instantiates a new BankTransfersEventsUpdateWebhookForAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransfersEventsUpdateWebhookForAuthWithDefaults() *BankTransfersEventsUpdateWebhookForAuth {
	this := BankTransfersEventsUpdateWebhookForAuth{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *BankTransfersEventsUpdateWebhookForAuth) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *BankTransfersEventsUpdateWebhookForAuth) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *BankTransfersEventsUpdateWebhookForAuth) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *BankTransfersEventsUpdateWebhookForAuth) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *BankTransfersEventsUpdateWebhookForAuth) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *BankTransfersEventsUpdateWebhookForAuth) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetEnvironment returns the Environment field value
func (o *BankTransfersEventsUpdateWebhookForAuth) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *BankTransfersEventsUpdateWebhookForAuth) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *BankTransfersEventsUpdateWebhookForAuth) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o BankTransfersEventsUpdateWebhookForAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BankTransfersEventsUpdateWebhookForAuth) UnmarshalJSON(bytes []byte) (err error) {
	varBankTransfersEventsUpdateWebhookForAuth := _BankTransfersEventsUpdateWebhookForAuth{}

	if err = json.Unmarshal(bytes, &varBankTransfersEventsUpdateWebhookForAuth); err == nil {
		*o = BankTransfersEventsUpdateWebhookForAuth(varBankTransfersEventsUpdateWebhookForAuth)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBankTransfersEventsUpdateWebhookForAuth struct {
	value *BankTransfersEventsUpdateWebhookForAuth
	isSet bool
}

func (v NullableBankTransfersEventsUpdateWebhookForAuth) Get() *BankTransfersEventsUpdateWebhookForAuth {
	return v.value
}

func (v *NullableBankTransfersEventsUpdateWebhookForAuth) Set(val *BankTransfersEventsUpdateWebhookForAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransfersEventsUpdateWebhookForAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransfersEventsUpdateWebhookForAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransfersEventsUpdateWebhookForAuth(val *BankTransfersEventsUpdateWebhookForAuth) *NullableBankTransfersEventsUpdateWebhookForAuth {
	return &NullableBankTransfersEventsUpdateWebhookForAuth{value: val, isSet: true}
}

func (v NullableBankTransfersEventsUpdateWebhookForAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransfersEventsUpdateWebhookForAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


