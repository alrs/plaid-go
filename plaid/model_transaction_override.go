/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.419.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TransactionOverride Data to populate as test transaction data. If not specified, random transactions will be generated instead.
type TransactionOverride struct {
	// The date of the transaction, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format. Transactions in Sandbox will move from pending to posted once their transaction date has been reached. If a `date_transacted` is not provided by the institution, a transaction date may be available in the [`authorized_date`](https://plaid.com/docs/api/products/transactions/#transactions-get-response-transactions-authorized-date) field.
	DateTransacted string `json:"date_transacted"`
	// The date the transaction posted, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format. Posted dates in the past or present will result in posted transactions; posted dates in the future will result in pending transactions.
	DatePosted string `json:"date_posted"`
	// The transaction amount. Can be negative.
	Amount float64 `json:"amount"`
	// The transaction description.
	Description string `json:"description"`
	// The ISO-4217 format currency code for the transaction.
	Currency *string `json:"currency,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransactionOverride TransactionOverride

// NewTransactionOverride instantiates a new TransactionOverride object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionOverride(dateTransacted string, datePosted string, amount float64, description string) *TransactionOverride {
	this := TransactionOverride{}
	this.DateTransacted = dateTransacted
	this.DatePosted = datePosted
	this.Amount = amount
	this.Description = description
	return &this
}

// NewTransactionOverrideWithDefaults instantiates a new TransactionOverride object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionOverrideWithDefaults() *TransactionOverride {
	this := TransactionOverride{}
	return &this
}

// GetDateTransacted returns the DateTransacted field value
func (o *TransactionOverride) GetDateTransacted() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateTransacted
}

// GetDateTransactedOk returns a tuple with the DateTransacted field value
// and a boolean to check if the value has been set.
func (o *TransactionOverride) GetDateTransactedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DateTransacted, true
}

// SetDateTransacted sets field value
func (o *TransactionOverride) SetDateTransacted(v string) {
	o.DateTransacted = v
}

// GetDatePosted returns the DatePosted field value
func (o *TransactionOverride) GetDatePosted() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatePosted
}

// GetDatePostedOk returns a tuple with the DatePosted field value
// and a boolean to check if the value has been set.
func (o *TransactionOverride) GetDatePostedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DatePosted, true
}

// SetDatePosted sets field value
func (o *TransactionOverride) SetDatePosted(v string) {
	o.DatePosted = v
}

// GetAmount returns the Amount field value
func (o *TransactionOverride) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransactionOverride) GetAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransactionOverride) SetAmount(v float64) {
	o.Amount = v
}

// GetDescription returns the Description field value
func (o *TransactionOverride) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TransactionOverride) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TransactionOverride) SetDescription(v string) {
	o.Description = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *TransactionOverride) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOverride) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *TransactionOverride) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *TransactionOverride) SetCurrency(v string) {
	o.Currency = &v
}

func (o TransactionOverride) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["date_transacted"] = o.DateTransacted
	}
	if true {
		toSerialize["date_posted"] = o.DatePosted
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransactionOverride) UnmarshalJSON(bytes []byte) (err error) {
	varTransactionOverride := _TransactionOverride{}

	if err = json.Unmarshal(bytes, &varTransactionOverride); err == nil {
		*o = TransactionOverride(varTransactionOverride)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "date_transacted")
		delete(additionalProperties, "date_posted")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "description")
		delete(additionalProperties, "currency")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransactionOverride struct {
	value *TransactionOverride
	isSet bool
}

func (v NullableTransactionOverride) Get() *TransactionOverride {
	return v.value
}

func (v *NullableTransactionOverride) Set(val *TransactionOverride) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionOverride) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionOverride) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionOverride(val *TransactionOverride) *NullableTransactionOverride {
	return &NullableTransactionOverride{value: val, isSet: true}
}

func (v NullableTransactionOverride) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionOverride) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


