/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.79.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"time"
)

// ItemStatusInvestments Information about the last successful and failed investments update for the Item.
type ItemStatusInvestments struct {
	// [ISO 8601](https://wikipedia.org/wiki/ISO_8601) timestamp of the last successful investments update for the Item. The status will update each time Plaid successfully connects with the institution, regardless of whether any new data is available in the update.
	LastSuccessfulUpdate NullableTime `json:"last_successful_update,omitempty"`
	// [ISO 8601](https://wikipedia.org/wiki/ISO_8601) timestamp of the last failed investments update for the Item. The status will update each time Plaid fails an attempt to connect with the institution, regardless of whether any new data is available in the update.
	LastFailedUpdate NullableTime `json:"last_failed_update,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ItemStatusInvestments ItemStatusInvestments

// NewItemStatusInvestments instantiates a new ItemStatusInvestments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemStatusInvestments() *ItemStatusInvestments {
	this := ItemStatusInvestments{}
	return &this
}

// NewItemStatusInvestmentsWithDefaults instantiates a new ItemStatusInvestments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemStatusInvestmentsWithDefaults() *ItemStatusInvestments {
	this := ItemStatusInvestments{}
	return &this
}

// GetLastSuccessfulUpdate returns the LastSuccessfulUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemStatusInvestments) GetLastSuccessfulUpdate() time.Time {
	if o == nil || o.LastSuccessfulUpdate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSuccessfulUpdate.Get()
}

// GetLastSuccessfulUpdateOk returns a tuple with the LastSuccessfulUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemStatusInvestments) GetLastSuccessfulUpdateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastSuccessfulUpdate.Get(), o.LastSuccessfulUpdate.IsSet()
}

// HasLastSuccessfulUpdate returns a boolean if a field has been set.
func (o *ItemStatusInvestments) HasLastSuccessfulUpdate() bool {
	if o != nil && o.LastSuccessfulUpdate.IsSet() {
		return true
	}

	return false
}

// SetLastSuccessfulUpdate gets a reference to the given NullableTime and assigns it to the LastSuccessfulUpdate field.
func (o *ItemStatusInvestments) SetLastSuccessfulUpdate(v time.Time) {
	o.LastSuccessfulUpdate.Set(&v)
}
// SetLastSuccessfulUpdateNil sets the value for LastSuccessfulUpdate to be an explicit nil
func (o *ItemStatusInvestments) SetLastSuccessfulUpdateNil() {
	o.LastSuccessfulUpdate.Set(nil)
}

// UnsetLastSuccessfulUpdate ensures that no value is present for LastSuccessfulUpdate, not even an explicit nil
func (o *ItemStatusInvestments) UnsetLastSuccessfulUpdate() {
	o.LastSuccessfulUpdate.Unset()
}

// GetLastFailedUpdate returns the LastFailedUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemStatusInvestments) GetLastFailedUpdate() time.Time {
	if o == nil || o.LastFailedUpdate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastFailedUpdate.Get()
}

// GetLastFailedUpdateOk returns a tuple with the LastFailedUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemStatusInvestments) GetLastFailedUpdateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastFailedUpdate.Get(), o.LastFailedUpdate.IsSet()
}

// HasLastFailedUpdate returns a boolean if a field has been set.
func (o *ItemStatusInvestments) HasLastFailedUpdate() bool {
	if o != nil && o.LastFailedUpdate.IsSet() {
		return true
	}

	return false
}

// SetLastFailedUpdate gets a reference to the given NullableTime and assigns it to the LastFailedUpdate field.
func (o *ItemStatusInvestments) SetLastFailedUpdate(v time.Time) {
	o.LastFailedUpdate.Set(&v)
}
// SetLastFailedUpdateNil sets the value for LastFailedUpdate to be an explicit nil
func (o *ItemStatusInvestments) SetLastFailedUpdateNil() {
	o.LastFailedUpdate.Set(nil)
}

// UnsetLastFailedUpdate ensures that no value is present for LastFailedUpdate, not even an explicit nil
func (o *ItemStatusInvestments) UnsetLastFailedUpdate() {
	o.LastFailedUpdate.Unset()
}

func (o ItemStatusInvestments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastSuccessfulUpdate.IsSet() {
		toSerialize["last_successful_update"] = o.LastSuccessfulUpdate.Get()
	}
	if o.LastFailedUpdate.IsSet() {
		toSerialize["last_failed_update"] = o.LastFailedUpdate.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ItemStatusInvestments) UnmarshalJSON(bytes []byte) (err error) {
	varItemStatusInvestments := _ItemStatusInvestments{}

	if err = json.Unmarshal(bytes, &varItemStatusInvestments); err == nil {
		*o = ItemStatusInvestments(varItemStatusInvestments)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "last_successful_update")
		delete(additionalProperties, "last_failed_update")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableItemStatusInvestments struct {
	value *ItemStatusInvestments
	isSet bool
}

func (v NullableItemStatusInvestments) Get() *ItemStatusInvestments {
	return v.value
}

func (v *NullableItemStatusInvestments) Set(val *ItemStatusInvestments) {
	v.value = val
	v.isSet = true
}

func (v NullableItemStatusInvestments) IsSet() bool {
	return v.isSet
}

func (v *NullableItemStatusInvestments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemStatusInvestments(val *ItemStatusInvestments) *NullableItemStatusInvestments {
	return &NullableItemStatusInvestments{value: val, isSet: true}
}

func (v NullableItemStatusInvestments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemStatusInvestments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


