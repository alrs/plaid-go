/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.210.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TransactionsEnrichGetResponse TransactionsEnrichGetResponse defines the response schema for `/transactions/enrich`.
type TransactionsEnrichGetResponse struct {
	// A list of enriched transactions.
	EnrichedTransactions []ClientProvidedEnrichedTransaction `json:"enriched_transactions"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId *string `json:"request_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransactionsEnrichGetResponse TransactionsEnrichGetResponse

// NewTransactionsEnrichGetResponse instantiates a new TransactionsEnrichGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsEnrichGetResponse(enrichedTransactions []ClientProvidedEnrichedTransaction) *TransactionsEnrichGetResponse {
	this := TransactionsEnrichGetResponse{}
	this.EnrichedTransactions = enrichedTransactions
	return &this
}

// NewTransactionsEnrichGetResponseWithDefaults instantiates a new TransactionsEnrichGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsEnrichGetResponseWithDefaults() *TransactionsEnrichGetResponse {
	this := TransactionsEnrichGetResponse{}
	return &this
}

// GetEnrichedTransactions returns the EnrichedTransactions field value
func (o *TransactionsEnrichGetResponse) GetEnrichedTransactions() []ClientProvidedEnrichedTransaction {
	if o == nil {
		var ret []ClientProvidedEnrichedTransaction
		return ret
	}

	return o.EnrichedTransactions
}

// GetEnrichedTransactionsOk returns a tuple with the EnrichedTransactions field value
// and a boolean to check if the value has been set.
func (o *TransactionsEnrichGetResponse) GetEnrichedTransactionsOk() (*[]ClientProvidedEnrichedTransaction, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EnrichedTransactions, true
}

// SetEnrichedTransactions sets field value
func (o *TransactionsEnrichGetResponse) SetEnrichedTransactions(v []ClientProvidedEnrichedTransaction) {
	o.EnrichedTransactions = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *TransactionsEnrichGetResponse) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsEnrichGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *TransactionsEnrichGetResponse) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *TransactionsEnrichGetResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o TransactionsEnrichGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enriched_transactions"] = o.EnrichedTransactions
	}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransactionsEnrichGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransactionsEnrichGetResponse := _TransactionsEnrichGetResponse{}

	if err = json.Unmarshal(bytes, &varTransactionsEnrichGetResponse); err == nil {
		*o = TransactionsEnrichGetResponse(varTransactionsEnrichGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "enriched_transactions")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransactionsEnrichGetResponse struct {
	value *TransactionsEnrichGetResponse
	isSet bool
}

func (v NullableTransactionsEnrichGetResponse) Get() *TransactionsEnrichGetResponse {
	return v.value
}

func (v *NullableTransactionsEnrichGetResponse) Set(val *TransactionsEnrichGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsEnrichGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsEnrichGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsEnrichGetResponse(val *TransactionsEnrichGetResponse) *NullableTransactionsEnrichGetResponse {
	return &NullableTransactionsEnrichGetResponse{value: val, isSet: true}
}

func (v NullableTransactionsEnrichGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsEnrichGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

