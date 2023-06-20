/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.379.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// BankIncomeRefreshUpdateWebhook Fired when a change to the user's income is detected. You should call `/credit/bank_income/refresh` to get updated income data for the user
type BankIncomeRefreshUpdateWebhook struct {
	// `INCOME`
	WebhookType string `json:"webhook_type"`
	// `BANK_INCOME_REFRESH_UPDATE`
	WebhookCode string `json:"webhook_code"`
	// The `user_id` corresponding to the user the webhook has fired for.
	UserId string `json:"user_id"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _BankIncomeRefreshUpdateWebhook BankIncomeRefreshUpdateWebhook

// NewBankIncomeRefreshUpdateWebhook instantiates a new BankIncomeRefreshUpdateWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankIncomeRefreshUpdateWebhook(webhookType string, webhookCode string, userId string, environment WebhookEnvironmentValues) *BankIncomeRefreshUpdateWebhook {
	this := BankIncomeRefreshUpdateWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.UserId = userId
	this.Environment = environment
	return &this
}

// NewBankIncomeRefreshUpdateWebhookWithDefaults instantiates a new BankIncomeRefreshUpdateWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankIncomeRefreshUpdateWebhookWithDefaults() *BankIncomeRefreshUpdateWebhook {
	this := BankIncomeRefreshUpdateWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *BankIncomeRefreshUpdateWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *BankIncomeRefreshUpdateWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *BankIncomeRefreshUpdateWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *BankIncomeRefreshUpdateWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *BankIncomeRefreshUpdateWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *BankIncomeRefreshUpdateWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetUserId returns the UserId field value
func (o *BankIncomeRefreshUpdateWebhook) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *BankIncomeRefreshUpdateWebhook) GetUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *BankIncomeRefreshUpdateWebhook) SetUserId(v string) {
	o.UserId = v
}

// GetEnvironment returns the Environment field value
func (o *BankIncomeRefreshUpdateWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *BankIncomeRefreshUpdateWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *BankIncomeRefreshUpdateWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o BankIncomeRefreshUpdateWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["user_id"] = o.UserId
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BankIncomeRefreshUpdateWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varBankIncomeRefreshUpdateWebhook := _BankIncomeRefreshUpdateWebhook{}

	if err = json.Unmarshal(bytes, &varBankIncomeRefreshUpdateWebhook); err == nil {
		*o = BankIncomeRefreshUpdateWebhook(varBankIncomeRefreshUpdateWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBankIncomeRefreshUpdateWebhook struct {
	value *BankIncomeRefreshUpdateWebhook
	isSet bool
}

func (v NullableBankIncomeRefreshUpdateWebhook) Get() *BankIncomeRefreshUpdateWebhook {
	return v.value
}

func (v *NullableBankIncomeRefreshUpdateWebhook) Set(val *BankIncomeRefreshUpdateWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableBankIncomeRefreshUpdateWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableBankIncomeRefreshUpdateWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankIncomeRefreshUpdateWebhook(val *BankIncomeRefreshUpdateWebhook) *NullableBankIncomeRefreshUpdateWebhook {
	return &NullableBankIncomeRefreshUpdateWebhook{value: val, isSet: true}
}

func (v NullableBankIncomeRefreshUpdateWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankIncomeRefreshUpdateWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

