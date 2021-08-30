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

// AccountAccess Allow or disallow product access by account. Unlisted (e.g. missing) accounts will be considered `new_accounts`.
type AccountAccess struct {
	// The unique account identifier for this account. This value must match that returned by the data access API for this account.
	UniqueId string `json:"unique_id"`
	// Allow the application to see this account (and associated details, including balance) in the list of accounts. If unset, defaults to `true`.
	Authorized NullableBool `json:"authorized,omitempty"`
}

// NewAccountAccess instantiates a new AccountAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountAccess(uniqueId string) *AccountAccess {
	this := AccountAccess{}
	this.UniqueId = uniqueId
	var authorized bool = true
	this.Authorized = *NewNullableBool(&authorized)
	return &this
}

// NewAccountAccessWithDefaults instantiates a new AccountAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountAccessWithDefaults() *AccountAccess {
	this := AccountAccess{}
	var authorized bool = true
	this.Authorized = *NewNullableBool(&authorized)
	return &this
}

// GetUniqueId returns the UniqueId field value
func (o *AccountAccess) GetUniqueId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UniqueId
}

// GetUniqueIdOk returns a tuple with the UniqueId field value
// and a boolean to check if the value has been set.
func (o *AccountAccess) GetUniqueIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UniqueId, true
}

// SetUniqueId sets field value
func (o *AccountAccess) SetUniqueId(v string) {
	o.UniqueId = v
}

// GetAuthorized returns the Authorized field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountAccess) GetAuthorized() bool {
	if o == nil || o.Authorized.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Authorized.Get()
}

// GetAuthorizedOk returns a tuple with the Authorized field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountAccess) GetAuthorizedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Authorized.Get(), o.Authorized.IsSet()
}

// HasAuthorized returns a boolean if a field has been set.
func (o *AccountAccess) HasAuthorized() bool {
	if o != nil && o.Authorized.IsSet() {
		return true
	}

	return false
}

// SetAuthorized gets a reference to the given NullableBool and assigns it to the Authorized field.
func (o *AccountAccess) SetAuthorized(v bool) {
	o.Authorized.Set(&v)
}
// SetAuthorizedNil sets the value for Authorized to be an explicit nil
func (o *AccountAccess) SetAuthorizedNil() {
	o.Authorized.Set(nil)
}

// UnsetAuthorized ensures that no value is present for Authorized, not even an explicit nil
func (o *AccountAccess) UnsetAuthorized() {
	o.Authorized.Unset()
}

func (o AccountAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["unique_id"] = o.UniqueId
	}
	if o.Authorized.IsSet() {
		toSerialize["authorized"] = o.Authorized.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAccountAccess struct {
	value *AccountAccess
	isSet bool
}

func (v NullableAccountAccess) Get() *AccountAccess {
	return v.value
}

func (v *NullableAccountAccess) Set(val *AccountAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAccess(val *AccountAccess) *NullableAccountAccess {
	return &NullableAccountAccess{value: val, isSet: true}
}

func (v NullableAccountAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


