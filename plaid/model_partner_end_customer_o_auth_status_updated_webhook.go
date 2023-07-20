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

// PartnerEndCustomerOAuthStatusUpdatedWebhook The webhook of type `PARTNER` and code `END_CUSTOMER_OAUTH_STATUS_UPDATED` will be fired when a partner's end customer has an update on their OAuth registration status with an institution.
type PartnerEndCustomerOAuthStatusUpdatedWebhook struct {
	// `PARTNER`
	WebhookType string `json:"webhook_type"`
	// `END_CUSTOMER_OAUTH_STATUS_UPDATED`
	WebhookCode string `json:"webhook_code"`
	// The client ID of the end customer
	EndCustomerClientId string `json:"end_customer_client_id"`
	Environment WebhookEnvironmentValues `json:"environment"`
	// The institution ID
	InstitutionId string `json:"institution_id"`
	// The institution name
	InstitutionName string `json:"institution_name"`
	Status PartnerEndCustomerOAuthStatusUpdatedValues `json:"status"`
}

// NewPartnerEndCustomerOAuthStatusUpdatedWebhook instantiates a new PartnerEndCustomerOAuthStatusUpdatedWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerEndCustomerOAuthStatusUpdatedWebhook(webhookType string, webhookCode string, endCustomerClientId string, environment WebhookEnvironmentValues, institutionId string, institutionName string, status PartnerEndCustomerOAuthStatusUpdatedValues) *PartnerEndCustomerOAuthStatusUpdatedWebhook {
	this := PartnerEndCustomerOAuthStatusUpdatedWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.EndCustomerClientId = endCustomerClientId
	this.Environment = environment
	this.InstitutionId = institutionId
	this.InstitutionName = institutionName
	this.Status = status
	return &this
}

// NewPartnerEndCustomerOAuthStatusUpdatedWebhookWithDefaults instantiates a new PartnerEndCustomerOAuthStatusUpdatedWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerEndCustomerOAuthStatusUpdatedWebhookWithDefaults() *PartnerEndCustomerOAuthStatusUpdatedWebhook {
	this := PartnerEndCustomerOAuthStatusUpdatedWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *PartnerEndCustomerOAuthStatusUpdatedWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerOAuthStatusUpdatedWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *PartnerEndCustomerOAuthStatusUpdatedWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *PartnerEndCustomerOAuthStatusUpdatedWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerOAuthStatusUpdatedWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *PartnerEndCustomerOAuthStatusUpdatedWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetEndCustomerClientId returns the EndCustomerClientId field value
func (o *PartnerEndCustomerOAuthStatusUpdatedWebhook) GetEndCustomerClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndCustomerClientId
}

// GetEndCustomerClientIdOk returns a tuple with the EndCustomerClientId field value
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerOAuthStatusUpdatedWebhook) GetEndCustomerClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndCustomerClientId, true
}

// SetEndCustomerClientId sets field value
func (o *PartnerEndCustomerOAuthStatusUpdatedWebhook) SetEndCustomerClientId(v string) {
	o.EndCustomerClientId = v
}

// GetEnvironment returns the Environment field value
func (o *PartnerEndCustomerOAuthStatusUpdatedWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerOAuthStatusUpdatedWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *PartnerEndCustomerOAuthStatusUpdatedWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

// GetInstitutionId returns the InstitutionId field value
func (o *PartnerEndCustomerOAuthStatusUpdatedWebhook) GetInstitutionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerOAuthStatusUpdatedWebhook) GetInstitutionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstitutionId, true
}

// SetInstitutionId sets field value
func (o *PartnerEndCustomerOAuthStatusUpdatedWebhook) SetInstitutionId(v string) {
	o.InstitutionId = v
}

// GetInstitutionName returns the InstitutionName field value
func (o *PartnerEndCustomerOAuthStatusUpdatedWebhook) GetInstitutionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstitutionName
}

// GetInstitutionNameOk returns a tuple with the InstitutionName field value
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerOAuthStatusUpdatedWebhook) GetInstitutionNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstitutionName, true
}

// SetInstitutionName sets field value
func (o *PartnerEndCustomerOAuthStatusUpdatedWebhook) SetInstitutionName(v string) {
	o.InstitutionName = v
}

// GetStatus returns the Status field value
func (o *PartnerEndCustomerOAuthStatusUpdatedWebhook) GetStatus() PartnerEndCustomerOAuthStatusUpdatedValues {
	if o == nil {
		var ret PartnerEndCustomerOAuthStatusUpdatedValues
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerOAuthStatusUpdatedWebhook) GetStatusOk() (*PartnerEndCustomerOAuthStatusUpdatedValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PartnerEndCustomerOAuthStatusUpdatedWebhook) SetStatus(v PartnerEndCustomerOAuthStatusUpdatedValues) {
	o.Status = v
}

func (o PartnerEndCustomerOAuthStatusUpdatedWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["end_customer_client_id"] = o.EndCustomerClientId
	}
	if true {
		toSerialize["environment"] = o.Environment
	}
	if true {
		toSerialize["institution_id"] = o.InstitutionId
	}
	if true {
		toSerialize["institution_name"] = o.InstitutionName
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullablePartnerEndCustomerOAuthStatusUpdatedWebhook struct {
	value *PartnerEndCustomerOAuthStatusUpdatedWebhook
	isSet bool
}

func (v NullablePartnerEndCustomerOAuthStatusUpdatedWebhook) Get() *PartnerEndCustomerOAuthStatusUpdatedWebhook {
	return v.value
}

func (v *NullablePartnerEndCustomerOAuthStatusUpdatedWebhook) Set(val *PartnerEndCustomerOAuthStatusUpdatedWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerEndCustomerOAuthStatusUpdatedWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerEndCustomerOAuthStatusUpdatedWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerEndCustomerOAuthStatusUpdatedWebhook(val *PartnerEndCustomerOAuthStatusUpdatedWebhook) *NullablePartnerEndCustomerOAuthStatusUpdatedWebhook {
	return &NullablePartnerEndCustomerOAuthStatusUpdatedWebhook{value: val, isSet: true}
}

func (v NullablePartnerEndCustomerOAuthStatusUpdatedWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerEndCustomerOAuthStatusUpdatedWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


