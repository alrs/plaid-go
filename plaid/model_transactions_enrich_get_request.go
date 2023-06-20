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

// TransactionsEnrichGetRequest TransactionsEnrichGetRequest defines the request schema for `/transactions/enrich`.
type TransactionsEnrichGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The account type for the requested transactions (either `depository` or `credit`).
	AccountType string `json:"account_type"`
	// An array of transaction objects to be enriched by Plaid. Maximum of 100 transactions per request.
	Transactions []ClientProvidedTransaction `json:"transactions"`
	Options *TransactionsEnrichRequestOptions `json:"options,omitempty"`
}

// NewTransactionsEnrichGetRequest instantiates a new TransactionsEnrichGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsEnrichGetRequest(accountType string, transactions []ClientProvidedTransaction) *TransactionsEnrichGetRequest {
	this := TransactionsEnrichGetRequest{}
	this.AccountType = accountType
	this.Transactions = transactions
	return &this
}

// NewTransactionsEnrichGetRequestWithDefaults instantiates a new TransactionsEnrichGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsEnrichGetRequestWithDefaults() *TransactionsEnrichGetRequest {
	this := TransactionsEnrichGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransactionsEnrichGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsEnrichGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransactionsEnrichGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransactionsEnrichGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransactionsEnrichGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsEnrichGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransactionsEnrichGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransactionsEnrichGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccountType returns the AccountType field value
func (o *TransactionsEnrichGetRequest) GetAccountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *TransactionsEnrichGetRequest) GetAccountTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *TransactionsEnrichGetRequest) SetAccountType(v string) {
	o.AccountType = v
}

// GetTransactions returns the Transactions field value
func (o *TransactionsEnrichGetRequest) GetTransactions() []ClientProvidedTransaction {
	if o == nil {
		var ret []ClientProvidedTransaction
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *TransactionsEnrichGetRequest) GetTransactionsOk() (*[]ClientProvidedTransaction, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Transactions, true
}

// SetTransactions sets field value
func (o *TransactionsEnrichGetRequest) SetTransactions(v []ClientProvidedTransaction) {
	o.Transactions = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *TransactionsEnrichGetRequest) GetOptions() TransactionsEnrichRequestOptions {
	if o == nil || o.Options == nil {
		var ret TransactionsEnrichRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsEnrichGetRequest) GetOptionsOk() (*TransactionsEnrichRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *TransactionsEnrichGetRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given TransactionsEnrichRequestOptions and assigns it to the Options field.
func (o *TransactionsEnrichGetRequest) SetOptions(v TransactionsEnrichRequestOptions) {
	o.Options = &v
}

func (o TransactionsEnrichGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["account_type"] = o.AccountType
	}
	if true {
		toSerialize["transactions"] = o.Transactions
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionsEnrichGetRequest struct {
	value *TransactionsEnrichGetRequest
	isSet bool
}

func (v NullableTransactionsEnrichGetRequest) Get() *TransactionsEnrichGetRequest {
	return v.value
}

func (v *NullableTransactionsEnrichGetRequest) Set(val *TransactionsEnrichGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsEnrichGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsEnrichGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsEnrichGetRequest(val *TransactionsEnrichGetRequest) *NullableTransactionsEnrichGetRequest {
	return &NullableTransactionsEnrichGetRequest{value: val, isSet: true}
}

func (v NullableTransactionsEnrichGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsEnrichGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


