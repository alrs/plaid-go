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

// PaymentInitiationConsentCreateRequest PaymentInitiationConsentCreateRequest defines the request schema for `/payment_initiation/consent/create`
type PaymentInitiationConsentCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The ID of the recipient the payment consent is for. The created consent can be used to transfer funds to this recipient only.
	RecipientId string `json:"recipient_id"`
	// A reference for the payment consent. This must be an alphanumeric string with at most 18 characters and must not contain any special characters.
	Reference string `json:"reference"`
	// An array of payment consent scopes.
	Scopes []PaymentInitiationConsentScope `json:"scopes"`
	Constraints PaymentInitiationConsentConstraints `json:"constraints"`
	Options NullableExternalPaymentInitiationConsentOptions `json:"options,omitempty"`
}

// NewPaymentInitiationConsentCreateRequest instantiates a new PaymentInitiationConsentCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationConsentCreateRequest(recipientId string, reference string, scopes []PaymentInitiationConsentScope, constraints PaymentInitiationConsentConstraints) *PaymentInitiationConsentCreateRequest {
	this := PaymentInitiationConsentCreateRequest{}
	this.RecipientId = recipientId
	this.Reference = reference
	this.Scopes = scopes
	this.Constraints = constraints
	return &this
}

// NewPaymentInitiationConsentCreateRequestWithDefaults instantiates a new PaymentInitiationConsentCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationConsentCreateRequestWithDefaults() *PaymentInitiationConsentCreateRequest {
	this := PaymentInitiationConsentCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PaymentInitiationConsentCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsentCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PaymentInitiationConsentCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PaymentInitiationConsentCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *PaymentInitiationConsentCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsentCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *PaymentInitiationConsentCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *PaymentInitiationConsentCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetRecipientId returns the RecipientId field value
func (o *PaymentInitiationConsentCreateRequest) GetRecipientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientId
}

// GetRecipientIdOk returns a tuple with the RecipientId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsentCreateRequest) GetRecipientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecipientId, true
}

// SetRecipientId sets field value
func (o *PaymentInitiationConsentCreateRequest) SetRecipientId(v string) {
	o.RecipientId = v
}

// GetReference returns the Reference field value
func (o *PaymentInitiationConsentCreateRequest) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsentCreateRequest) GetReferenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *PaymentInitiationConsentCreateRequest) SetReference(v string) {
	o.Reference = v
}

// GetScopes returns the Scopes field value
func (o *PaymentInitiationConsentCreateRequest) GetScopes() []PaymentInitiationConsentScope {
	if o == nil {
		var ret []PaymentInitiationConsentScope
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsentCreateRequest) GetScopesOk() (*[]PaymentInitiationConsentScope, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Scopes, true
}

// SetScopes sets field value
func (o *PaymentInitiationConsentCreateRequest) SetScopes(v []PaymentInitiationConsentScope) {
	o.Scopes = v
}

// GetConstraints returns the Constraints field value
func (o *PaymentInitiationConsentCreateRequest) GetConstraints() PaymentInitiationConsentConstraints {
	if o == nil {
		var ret PaymentInitiationConsentConstraints
		return ret
	}

	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsentCreateRequest) GetConstraintsOk() (*PaymentInitiationConsentConstraints, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Constraints, true
}

// SetConstraints sets field value
func (o *PaymentInitiationConsentCreateRequest) SetConstraints(v PaymentInitiationConsentConstraints) {
	o.Constraints = v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationConsentCreateRequest) GetOptions() ExternalPaymentInitiationConsentOptions {
	if o == nil || o.Options.Get() == nil {
		var ret ExternalPaymentInitiationConsentOptions
		return ret
	}
	return *o.Options.Get()
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationConsentCreateRequest) GetOptionsOk() (*ExternalPaymentInitiationConsentOptions, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Options.Get(), o.Options.IsSet()
}

// HasOptions returns a boolean if a field has been set.
func (o *PaymentInitiationConsentCreateRequest) HasOptions() bool {
	if o != nil && o.Options.IsSet() {
		return true
	}

	return false
}

// SetOptions gets a reference to the given NullableExternalPaymentInitiationConsentOptions and assigns it to the Options field.
func (o *PaymentInitiationConsentCreateRequest) SetOptions(v ExternalPaymentInitiationConsentOptions) {
	o.Options.Set(&v)
}
// SetOptionsNil sets the value for Options to be an explicit nil
func (o *PaymentInitiationConsentCreateRequest) SetOptionsNil() {
	o.Options.Set(nil)
}

// UnsetOptions ensures that no value is present for Options, not even an explicit nil
func (o *PaymentInitiationConsentCreateRequest) UnsetOptions() {
	o.Options.Unset()
}

func (o PaymentInitiationConsentCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["recipient_id"] = o.RecipientId
	}
	if true {
		toSerialize["reference"] = o.Reference
	}
	if true {
		toSerialize["scopes"] = o.Scopes
	}
	if true {
		toSerialize["constraints"] = o.Constraints
	}
	if o.Options.IsSet() {
		toSerialize["options"] = o.Options.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentInitiationConsentCreateRequest struct {
	value *PaymentInitiationConsentCreateRequest
	isSet bool
}

func (v NullablePaymentInitiationConsentCreateRequest) Get() *PaymentInitiationConsentCreateRequest {
	return v.value
}

func (v *NullablePaymentInitiationConsentCreateRequest) Set(val *PaymentInitiationConsentCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationConsentCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationConsentCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationConsentCreateRequest(val *PaymentInitiationConsentCreateRequest) *NullablePaymentInitiationConsentCreateRequest {
	return &NullablePaymentInitiationConsentCreateRequest{value: val, isSet: true}
}

func (v NullablePaymentInitiationConsentCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationConsentCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


