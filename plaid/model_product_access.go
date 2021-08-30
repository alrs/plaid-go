/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.26.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// ProductAccess The product access being requested. Used to or disallow product access across all accounts. If unset, defaults to all products allowed.
type ProductAccess struct {
	// Allow access to statements. If unset, defaults to `true`.
	Statements NullableBool `json:"statements,omitempty"`
	// Allow access to the Identity product (name, email, phone, address). If unset, defaults to `true`.
	Identity NullableBool `json:"identity,omitempty"`
	// Allow access to account number details. If unset, defaults to `true`.
	Auth NullableBool `json:"auth,omitempty"`
	// Allow access to transaction details. If unset, defaults to `true`.
	Transactions NullableBool `json:"transactions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductAccess ProductAccess

// NewProductAccess instantiates a new ProductAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAccess() *ProductAccess {
	this := ProductAccess{}
	var statements bool = true
	this.Statements = *NewNullableBool(&statements)
	var identity bool = true
	this.Identity = *NewNullableBool(&identity)
	var auth bool = true
	this.Auth = *NewNullableBool(&auth)
	var transactions bool = true
	this.Transactions = *NewNullableBool(&transactions)
	return &this
}

// NewProductAccessWithDefaults instantiates a new ProductAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAccessWithDefaults() *ProductAccess {
	this := ProductAccess{}
	var statements bool = true
	this.Statements = *NewNullableBool(&statements)
	var identity bool = true
	this.Identity = *NewNullableBool(&identity)
	var auth bool = true
	this.Auth = *NewNullableBool(&auth)
	var transactions bool = true
	this.Transactions = *NewNullableBool(&transactions)
	return &this
}

// GetStatements returns the Statements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAccess) GetStatements() bool {
	if o == nil || o.Statements.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Statements.Get()
}

// GetStatementsOk returns a tuple with the Statements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductAccess) GetStatementsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Statements.Get(), o.Statements.IsSet()
}

// HasStatements returns a boolean if a field has been set.
func (o *ProductAccess) HasStatements() bool {
	if o != nil && o.Statements.IsSet() {
		return true
	}

	return false
}

// SetStatements gets a reference to the given NullableBool and assigns it to the Statements field.
func (o *ProductAccess) SetStatements(v bool) {
	o.Statements.Set(&v)
}
// SetStatementsNil sets the value for Statements to be an explicit nil
func (o *ProductAccess) SetStatementsNil() {
	o.Statements.Set(nil)
}

// UnsetStatements ensures that no value is present for Statements, not even an explicit nil
func (o *ProductAccess) UnsetStatements() {
	o.Statements.Unset()
}

// GetIdentity returns the Identity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAccess) GetIdentity() bool {
	if o == nil || o.Identity.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Identity.Get()
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductAccess) GetIdentityOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Identity.Get(), o.Identity.IsSet()
}

// HasIdentity returns a boolean if a field has been set.
func (o *ProductAccess) HasIdentity() bool {
	if o != nil && o.Identity.IsSet() {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given NullableBool and assigns it to the Identity field.
func (o *ProductAccess) SetIdentity(v bool) {
	o.Identity.Set(&v)
}
// SetIdentityNil sets the value for Identity to be an explicit nil
func (o *ProductAccess) SetIdentityNil() {
	o.Identity.Set(nil)
}

// UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
func (o *ProductAccess) UnsetIdentity() {
	o.Identity.Unset()
}

// GetAuth returns the Auth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAccess) GetAuth() bool {
	if o == nil || o.Auth.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Auth.Get()
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductAccess) GetAuthOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Auth.Get(), o.Auth.IsSet()
}

// HasAuth returns a boolean if a field has been set.
func (o *ProductAccess) HasAuth() bool {
	if o != nil && o.Auth.IsSet() {
		return true
	}

	return false
}

// SetAuth gets a reference to the given NullableBool and assigns it to the Auth field.
func (o *ProductAccess) SetAuth(v bool) {
	o.Auth.Set(&v)
}
// SetAuthNil sets the value for Auth to be an explicit nil
func (o *ProductAccess) SetAuthNil() {
	o.Auth.Set(nil)
}

// UnsetAuth ensures that no value is present for Auth, not even an explicit nil
func (o *ProductAccess) UnsetAuth() {
	o.Auth.Unset()
}

// GetTransactions returns the Transactions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAccess) GetTransactions() bool {
	if o == nil || o.Transactions.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Transactions.Get()
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductAccess) GetTransactionsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Transactions.Get(), o.Transactions.IsSet()
}

// HasTransactions returns a boolean if a field has been set.
func (o *ProductAccess) HasTransactions() bool {
	if o != nil && o.Transactions.IsSet() {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given NullableBool and assigns it to the Transactions field.
func (o *ProductAccess) SetTransactions(v bool) {
	o.Transactions.Set(&v)
}
// SetTransactionsNil sets the value for Transactions to be an explicit nil
func (o *ProductAccess) SetTransactionsNil() {
	o.Transactions.Set(nil)
}

// UnsetTransactions ensures that no value is present for Transactions, not even an explicit nil
func (o *ProductAccess) UnsetTransactions() {
	o.Transactions.Unset()
}

func (o ProductAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Statements.IsSet() {
		toSerialize["statements"] = o.Statements.Get()
	}
	if o.Identity.IsSet() {
		toSerialize["identity"] = o.Identity.Get()
	}
	if o.Auth.IsSet() {
		toSerialize["auth"] = o.Auth.Get()
	}
	if o.Transactions.IsSet() {
		toSerialize["transactions"] = o.Transactions.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProductAccess) UnmarshalJSON(bytes []byte) (err error) {
	varProductAccess := _ProductAccess{}

	if err = json.Unmarshal(bytes, &varProductAccess); err == nil {
		*o = ProductAccess(varProductAccess)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "statements")
		delete(additionalProperties, "identity")
		delete(additionalProperties, "auth")
		delete(additionalProperties, "transactions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProductAccess struct {
	value *ProductAccess
	isSet bool
}

func (v NullableProductAccess) Get() *ProductAccess {
	return v.value
}

func (v *NullableProductAccess) Set(val *ProductAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAccess(val *ProductAccess) *NullableProductAccess {
	return &NullableProductAccess{value: val, isSet: true}
}

func (v NullableProductAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


