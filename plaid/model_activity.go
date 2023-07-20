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

// Activity Describes a consent activity.
type Activity struct {
	Activity ActivityType `json:"activity"`
	// The date this activity was initiated [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format in UTC.
	InitiatedDate string `json:"initiated_date"`
	// A unique identifier for the activity
	Id string `json:"id"`
	// Application ID of the client who initiated the activity.
	Initiator string `json:"initiator"`
	State ActionState `json:"state"`
	// This field will map to the application ID that is returned from /item/applications/list, or provided to the institution in an oauth redirect.
	TargetApplicationId *string `json:"target_application_id,omitempty"`
	Scopes NullableScopesNullable `json:"scopes,omitempty"`
}

// NewActivity instantiates a new Activity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivity(activity ActivityType, initiatedDate string, id string, initiator string, state ActionState) *Activity {
	this := Activity{}
	this.Activity = activity
	this.InitiatedDate = initiatedDate
	this.Id = id
	this.Initiator = initiator
	this.State = state
	return &this
}

// NewActivityWithDefaults instantiates a new Activity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityWithDefaults() *Activity {
	this := Activity{}
	return &this
}

// GetActivity returns the Activity field value
func (o *Activity) GetActivity() ActivityType {
	if o == nil {
		var ret ActivityType
		return ret
	}

	return o.Activity
}

// GetActivityOk returns a tuple with the Activity field value
// and a boolean to check if the value has been set.
func (o *Activity) GetActivityOk() (*ActivityType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Activity, true
}

// SetActivity sets field value
func (o *Activity) SetActivity(v ActivityType) {
	o.Activity = v
}

// GetInitiatedDate returns the InitiatedDate field value
func (o *Activity) GetInitiatedDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InitiatedDate
}

// GetInitiatedDateOk returns a tuple with the InitiatedDate field value
// and a boolean to check if the value has been set.
func (o *Activity) GetInitiatedDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InitiatedDate, true
}

// SetInitiatedDate sets field value
func (o *Activity) SetInitiatedDate(v string) {
	o.InitiatedDate = v
}

// GetId returns the Id field value
func (o *Activity) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Activity) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Activity) SetId(v string) {
	o.Id = v
}

// GetInitiator returns the Initiator field value
func (o *Activity) GetInitiator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value
// and a boolean to check if the value has been set.
func (o *Activity) GetInitiatorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Initiator, true
}

// SetInitiator sets field value
func (o *Activity) SetInitiator(v string) {
	o.Initiator = v
}

// GetState returns the State field value
func (o *Activity) GetState() ActionState {
	if o == nil {
		var ret ActionState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Activity) GetStateOk() (*ActionState, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *Activity) SetState(v ActionState) {
	o.State = v
}

// GetTargetApplicationId returns the TargetApplicationId field value if set, zero value otherwise.
func (o *Activity) GetTargetApplicationId() string {
	if o == nil || o.TargetApplicationId == nil {
		var ret string
		return ret
	}
	return *o.TargetApplicationId
}

// GetTargetApplicationIdOk returns a tuple with the TargetApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetTargetApplicationIdOk() (*string, bool) {
	if o == nil || o.TargetApplicationId == nil {
		return nil, false
	}
	return o.TargetApplicationId, true
}

// HasTargetApplicationId returns a boolean if a field has been set.
func (o *Activity) HasTargetApplicationId() bool {
	if o != nil && o.TargetApplicationId != nil {
		return true
	}

	return false
}

// SetTargetApplicationId gets a reference to the given string and assigns it to the TargetApplicationId field.
func (o *Activity) SetTargetApplicationId(v string) {
	o.TargetApplicationId = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Activity) GetScopes() ScopesNullable {
	if o == nil || o.Scopes.Get() == nil {
		var ret ScopesNullable
		return ret
	}
	return *o.Scopes.Get()
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Activity) GetScopesOk() (*ScopesNullable, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Scopes.Get(), o.Scopes.IsSet()
}

// HasScopes returns a boolean if a field has been set.
func (o *Activity) HasScopes() bool {
	if o != nil && o.Scopes.IsSet() {
		return true
	}

	return false
}

// SetScopes gets a reference to the given NullableScopesNullable and assigns it to the Scopes field.
func (o *Activity) SetScopes(v ScopesNullable) {
	o.Scopes.Set(&v)
}
// SetScopesNil sets the value for Scopes to be an explicit nil
func (o *Activity) SetScopesNil() {
	o.Scopes.Set(nil)
}

// UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
func (o *Activity) UnsetScopes() {
	o.Scopes.Unset()
}

func (o Activity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["activity"] = o.Activity
	}
	if true {
		toSerialize["initiated_date"] = o.InitiatedDate
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["initiator"] = o.Initiator
	}
	if true {
		toSerialize["state"] = o.State
	}
	if o.TargetApplicationId != nil {
		toSerialize["target_application_id"] = o.TargetApplicationId
	}
	if o.Scopes.IsSet() {
		toSerialize["scopes"] = o.Scopes.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableActivity struct {
	value *Activity
	isSet bool
}

func (v NullableActivity) Get() *Activity {
	return v.value
}

func (v *NullableActivity) Set(val *Activity) {
	v.value = val
	v.isSet = true
}

func (v NullableActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivity(val *Activity) *NullableActivity {
	return &NullableActivity{value: val, isSet: true}
}

func (v NullableActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


