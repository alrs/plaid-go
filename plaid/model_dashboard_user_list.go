/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.121.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// DashboardUserList List of dashboard users
type DashboardUserList struct {
	// List of dashboard users
	DashboardUsers []DashboardUser `json:"dashboard_users"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
}

// NewDashboardUserList instantiates a new DashboardUserList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardUserList(dashboardUsers []DashboardUser, requestId string) *DashboardUserList {
	this := DashboardUserList{}
	this.DashboardUsers = dashboardUsers
	this.RequestId = requestId
	return &this
}

// NewDashboardUserListWithDefaults instantiates a new DashboardUserList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardUserListWithDefaults() *DashboardUserList {
	this := DashboardUserList{}
	return &this
}

// GetDashboardUsers returns the DashboardUsers field value
func (o *DashboardUserList) GetDashboardUsers() []DashboardUser {
	if o == nil {
		var ret []DashboardUser
		return ret
	}

	return o.DashboardUsers
}

// GetDashboardUsersOk returns a tuple with the DashboardUsers field value
// and a boolean to check if the value has been set.
func (o *DashboardUserList) GetDashboardUsersOk() (*[]DashboardUser, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DashboardUsers, true
}

// SetDashboardUsers sets field value
func (o *DashboardUserList) SetDashboardUsers(v []DashboardUser) {
	o.DashboardUsers = v
}

// GetRequestId returns the RequestId field value
func (o *DashboardUserList) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *DashboardUserList) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *DashboardUserList) SetRequestId(v string) {
	o.RequestId = v
}

func (o DashboardUserList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dashboard_users"] = o.DashboardUsers
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}
	return json.Marshal(toSerialize)
}

type NullableDashboardUserList struct {
	value *DashboardUserList
	isSet bool
}

func (v NullableDashboardUserList) Get() *DashboardUserList {
	return v.value
}

func (v *NullableDashboardUserList) Set(val *DashboardUserList) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardUserList) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardUserList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardUserList(val *DashboardUserList) *NullableDashboardUserList {
	return &NullableDashboardUserList{value: val, isSet: true}
}

func (v NullableDashboardUserList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardUserList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


