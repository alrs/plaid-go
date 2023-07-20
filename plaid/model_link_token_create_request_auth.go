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

// LinkTokenCreateRequestAuth Specifies options for initializing Link for use with the Auth product. This field can be used to enable or disable extended Auth flows for the resulting Link session. Omitting any field will result in a default that can be configured by your account manager.
type LinkTokenCreateRequestAuth struct {
	// Specifies whether Auth Type Select is enabled for the Link session, allowing the end user to choose between linking instantly or manually prior to selecting their financial institution. Note that this can only be true if `same_day_microdeposits_enabled` is set to true.
	AuthTypeSelectEnabled *bool `json:"auth_type_select_enabled,omitempty"`
	// Specifies whether the Link session is enabled for the Automated Micro-deposits flow.
	AutomatedMicrodepositsEnabled *bool `json:"automated_microdeposits_enabled,omitempty"`
	// Specifies whether the Link session is enabled for the Instant Match flow. As of November 2022, Instant Match will be enabled by default. Instant Match can be disabled by setting this field to `false`.
	InstantMatchEnabled *bool `json:"instant_match_enabled,omitempty"`
	// Specifies whether the Link session is enabled for the Same Day Micro-deposits flow.
	SameDayMicrodepositsEnabled *bool `json:"same_day_microdeposits_enabled,omitempty"`
	// Specifies what type of Reroute to Credentials pane should be used in the Link session for the Same Day Micro-deposits flow.
	RerouteToCredentials *string `json:"reroute_to_credentials,omitempty"`
	// This field has been deprecated in favor of `auth_type_select_enabled`.
	FlowType *string `json:"flow_type,omitempty"`
}

// NewLinkTokenCreateRequestAuth instantiates a new LinkTokenCreateRequestAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenCreateRequestAuth() *LinkTokenCreateRequestAuth {
	this := LinkTokenCreateRequestAuth{}
	var authTypeSelectEnabled bool = false
	this.AuthTypeSelectEnabled = &authTypeSelectEnabled
	return &this
}

// NewLinkTokenCreateRequestAuthWithDefaults instantiates a new LinkTokenCreateRequestAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenCreateRequestAuthWithDefaults() *LinkTokenCreateRequestAuth {
	this := LinkTokenCreateRequestAuth{}
	var authTypeSelectEnabled bool = false
	this.AuthTypeSelectEnabled = &authTypeSelectEnabled
	return &this
}

// GetAuthTypeSelectEnabled returns the AuthTypeSelectEnabled field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestAuth) GetAuthTypeSelectEnabled() bool {
	if o == nil || o.AuthTypeSelectEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AuthTypeSelectEnabled
}

// GetAuthTypeSelectEnabledOk returns a tuple with the AuthTypeSelectEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestAuth) GetAuthTypeSelectEnabledOk() (*bool, bool) {
	if o == nil || o.AuthTypeSelectEnabled == nil {
		return nil, false
	}
	return o.AuthTypeSelectEnabled, true
}

// HasAuthTypeSelectEnabled returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestAuth) HasAuthTypeSelectEnabled() bool {
	if o != nil && o.AuthTypeSelectEnabled != nil {
		return true
	}

	return false
}

// SetAuthTypeSelectEnabled gets a reference to the given bool and assigns it to the AuthTypeSelectEnabled field.
func (o *LinkTokenCreateRequestAuth) SetAuthTypeSelectEnabled(v bool) {
	o.AuthTypeSelectEnabled = &v
}

// GetAutomatedMicrodepositsEnabled returns the AutomatedMicrodepositsEnabled field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestAuth) GetAutomatedMicrodepositsEnabled() bool {
	if o == nil || o.AutomatedMicrodepositsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AutomatedMicrodepositsEnabled
}

// GetAutomatedMicrodepositsEnabledOk returns a tuple with the AutomatedMicrodepositsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestAuth) GetAutomatedMicrodepositsEnabledOk() (*bool, bool) {
	if o == nil || o.AutomatedMicrodepositsEnabled == nil {
		return nil, false
	}
	return o.AutomatedMicrodepositsEnabled, true
}

// HasAutomatedMicrodepositsEnabled returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestAuth) HasAutomatedMicrodepositsEnabled() bool {
	if o != nil && o.AutomatedMicrodepositsEnabled != nil {
		return true
	}

	return false
}

// SetAutomatedMicrodepositsEnabled gets a reference to the given bool and assigns it to the AutomatedMicrodepositsEnabled field.
func (o *LinkTokenCreateRequestAuth) SetAutomatedMicrodepositsEnabled(v bool) {
	o.AutomatedMicrodepositsEnabled = &v
}

// GetInstantMatchEnabled returns the InstantMatchEnabled field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestAuth) GetInstantMatchEnabled() bool {
	if o == nil || o.InstantMatchEnabled == nil {
		var ret bool
		return ret
	}
	return *o.InstantMatchEnabled
}

// GetInstantMatchEnabledOk returns a tuple with the InstantMatchEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestAuth) GetInstantMatchEnabledOk() (*bool, bool) {
	if o == nil || o.InstantMatchEnabled == nil {
		return nil, false
	}
	return o.InstantMatchEnabled, true
}

// HasInstantMatchEnabled returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestAuth) HasInstantMatchEnabled() bool {
	if o != nil && o.InstantMatchEnabled != nil {
		return true
	}

	return false
}

// SetInstantMatchEnabled gets a reference to the given bool and assigns it to the InstantMatchEnabled field.
func (o *LinkTokenCreateRequestAuth) SetInstantMatchEnabled(v bool) {
	o.InstantMatchEnabled = &v
}

// GetSameDayMicrodepositsEnabled returns the SameDayMicrodepositsEnabled field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestAuth) GetSameDayMicrodepositsEnabled() bool {
	if o == nil || o.SameDayMicrodepositsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SameDayMicrodepositsEnabled
}

// GetSameDayMicrodepositsEnabledOk returns a tuple with the SameDayMicrodepositsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestAuth) GetSameDayMicrodepositsEnabledOk() (*bool, bool) {
	if o == nil || o.SameDayMicrodepositsEnabled == nil {
		return nil, false
	}
	return o.SameDayMicrodepositsEnabled, true
}

// HasSameDayMicrodepositsEnabled returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestAuth) HasSameDayMicrodepositsEnabled() bool {
	if o != nil && o.SameDayMicrodepositsEnabled != nil {
		return true
	}

	return false
}

// SetSameDayMicrodepositsEnabled gets a reference to the given bool and assigns it to the SameDayMicrodepositsEnabled field.
func (o *LinkTokenCreateRequestAuth) SetSameDayMicrodepositsEnabled(v bool) {
	o.SameDayMicrodepositsEnabled = &v
}

// GetRerouteToCredentials returns the RerouteToCredentials field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestAuth) GetRerouteToCredentials() string {
	if o == nil || o.RerouteToCredentials == nil {
		var ret string
		return ret
	}
	return *o.RerouteToCredentials
}

// GetRerouteToCredentialsOk returns a tuple with the RerouteToCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestAuth) GetRerouteToCredentialsOk() (*string, bool) {
	if o == nil || o.RerouteToCredentials == nil {
		return nil, false
	}
	return o.RerouteToCredentials, true
}

// HasRerouteToCredentials returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestAuth) HasRerouteToCredentials() bool {
	if o != nil && o.RerouteToCredentials != nil {
		return true
	}

	return false
}

// SetRerouteToCredentials gets a reference to the given string and assigns it to the RerouteToCredentials field.
func (o *LinkTokenCreateRequestAuth) SetRerouteToCredentials(v string) {
	o.RerouteToCredentials = &v
}

// GetFlowType returns the FlowType field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestAuth) GetFlowType() string {
	if o == nil || o.FlowType == nil {
		var ret string
		return ret
	}
	return *o.FlowType
}

// GetFlowTypeOk returns a tuple with the FlowType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestAuth) GetFlowTypeOk() (*string, bool) {
	if o == nil || o.FlowType == nil {
		return nil, false
	}
	return o.FlowType, true
}

// HasFlowType returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestAuth) HasFlowType() bool {
	if o != nil && o.FlowType != nil {
		return true
	}

	return false
}

// SetFlowType gets a reference to the given string and assigns it to the FlowType field.
func (o *LinkTokenCreateRequestAuth) SetFlowType(v string) {
	o.FlowType = &v
}

func (o LinkTokenCreateRequestAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthTypeSelectEnabled != nil {
		toSerialize["auth_type_select_enabled"] = o.AuthTypeSelectEnabled
	}
	if o.AutomatedMicrodepositsEnabled != nil {
		toSerialize["automated_microdeposits_enabled"] = o.AutomatedMicrodepositsEnabled
	}
	if o.InstantMatchEnabled != nil {
		toSerialize["instant_match_enabled"] = o.InstantMatchEnabled
	}
	if o.SameDayMicrodepositsEnabled != nil {
		toSerialize["same_day_microdeposits_enabled"] = o.SameDayMicrodepositsEnabled
	}
	if o.RerouteToCredentials != nil {
		toSerialize["reroute_to_credentials"] = o.RerouteToCredentials
	}
	if o.FlowType != nil {
		toSerialize["flow_type"] = o.FlowType
	}
	return json.Marshal(toSerialize)
}

type NullableLinkTokenCreateRequestAuth struct {
	value *LinkTokenCreateRequestAuth
	isSet bool
}

func (v NullableLinkTokenCreateRequestAuth) Get() *LinkTokenCreateRequestAuth {
	return v.value
}

func (v *NullableLinkTokenCreateRequestAuth) Set(val *LinkTokenCreateRequestAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenCreateRequestAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenCreateRequestAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenCreateRequestAuth(val *LinkTokenCreateRequestAuth) *NullableLinkTokenCreateRequestAuth {
	return &NullableLinkTokenCreateRequestAuth{value: val, isSet: true}
}

func (v NullableLinkTokenCreateRequestAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenCreateRequestAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


