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
)

// ClientProvidedTransaction A client-provided transaction for Plaid to enrich.
type ClientProvidedTransaction struct {
	// A unique ID for the transaction used to help you tie data back to your systems.
	Id string `json:"id"`
	// The raw description of the transaction.
	Description string `json:"description"`
	// The absolute value of the transaction (>= 0)
	Amount float64 `json:"amount"`
	Direction *EnrichTransactionDirection `json:"direction,omitempty"`
	// The ISO-4217 currency code of the transaction, e.g., USD.
	IsoCurrencyCode string `json:"iso_currency_code"`
	AdditionalProperties map[string]interface{}
}

type _ClientProvidedTransaction ClientProvidedTransaction

// NewClientProvidedTransaction instantiates a new ClientProvidedTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientProvidedTransaction(id string, description string, amount float64, isoCurrencyCode string) *ClientProvidedTransaction {
	this := ClientProvidedTransaction{}
	this.Id = id
	this.Description = description
	this.Amount = amount
	this.IsoCurrencyCode = isoCurrencyCode
	return &this
}

// NewClientProvidedTransactionWithDefaults instantiates a new ClientProvidedTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientProvidedTransactionWithDefaults() *ClientProvidedTransaction {
	this := ClientProvidedTransaction{}
	return &this
}

// GetId returns the Id field value
func (o *ClientProvidedTransaction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ClientProvidedTransaction) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ClientProvidedTransaction) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value
func (o *ClientProvidedTransaction) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ClientProvidedTransaction) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ClientProvidedTransaction) SetDescription(v string) {
	o.Description = v
}

// GetAmount returns the Amount field value
func (o *ClientProvidedTransaction) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ClientProvidedTransaction) GetAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ClientProvidedTransaction) SetAmount(v float64) {
	o.Amount = v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *ClientProvidedTransaction) GetDirection() EnrichTransactionDirection {
	if o == nil || o.Direction == nil {
		var ret EnrichTransactionDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProvidedTransaction) GetDirectionOk() (*EnrichTransactionDirection, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *ClientProvidedTransaction) HasDirection() bool {
	if o != nil && o.Direction != nil {
		return true
	}

	return false
}

// SetDirection gets a reference to the given EnrichTransactionDirection and assigns it to the Direction field.
func (o *ClientProvidedTransaction) SetDirection(v EnrichTransactionDirection) {
	o.Direction = &v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
func (o *ClientProvidedTransaction) GetIsoCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *ClientProvidedTransaction) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsoCurrencyCode, true
}

// SetIsoCurrencyCode sets field value
func (o *ClientProvidedTransaction) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = v
}

func (o ClientProvidedTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.Direction != nil {
		toSerialize["direction"] = o.Direction
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ClientProvidedTransaction) UnmarshalJSON(bytes []byte) (err error) {
	varClientProvidedTransaction := _ClientProvidedTransaction{}

	if err = json.Unmarshal(bytes, &varClientProvidedTransaction); err == nil {
		*o = ClientProvidedTransaction(varClientProvidedTransaction)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "direction")
		delete(additionalProperties, "iso_currency_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientProvidedTransaction struct {
	value *ClientProvidedTransaction
	isSet bool
}

func (v NullableClientProvidedTransaction) Get() *ClientProvidedTransaction {
	return v.value
}

func (v *NullableClientProvidedTransaction) Set(val *ClientProvidedTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableClientProvidedTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableClientProvidedTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientProvidedTransaction(val *ClientProvidedTransaction) *NullableClientProvidedTransaction {
	return &NullableClientProvidedTransaction{value: val, isSet: true}
}

func (v NullableClientProvidedTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientProvidedTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


