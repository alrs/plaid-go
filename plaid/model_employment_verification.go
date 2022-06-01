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

// EmploymentVerification An object containing proof of employment data for an individual
type EmploymentVerification struct {
	Status NullableEmploymentVerificationStatus `json:"status,omitempty"`
	// Start of employment in ISO 8601 format (YYYY-MM-DD).
	StartDate NullableString `json:"start_date,omitempty"`
	// End of employment, if applicable. Provided in ISO 8601 format (YYY-MM-DD).
	EndDate NullableString `json:"end_date,omitempty"`
	Employer *EmployerVerification `json:"employer,omitempty"`
	// Current title of employee.
	Title NullableString `json:"title,omitempty"`
	PlatformIds *PlatformIds `json:"platform_ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmploymentVerification EmploymentVerification

// NewEmploymentVerification instantiates a new EmploymentVerification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmploymentVerification() *EmploymentVerification {
	this := EmploymentVerification{}
	return &this
}

// NewEmploymentVerificationWithDefaults instantiates a new EmploymentVerification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmploymentVerificationWithDefaults() *EmploymentVerification {
	this := EmploymentVerification{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmploymentVerification) GetStatus() EmploymentVerificationStatus {
	if o == nil || o.Status.Get() == nil {
		var ret EmploymentVerificationStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmploymentVerification) GetStatusOk() (*EmploymentVerificationStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *EmploymentVerification) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableEmploymentVerificationStatus and assigns it to the Status field.
func (o *EmploymentVerification) SetStatus(v EmploymentVerificationStatus) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *EmploymentVerification) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *EmploymentVerification) UnsetStatus() {
	o.Status.Unset()
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmploymentVerification) GetStartDate() string {
	if o == nil || o.StartDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmploymentVerification) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *EmploymentVerification) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableString and assigns it to the StartDate field.
func (o *EmploymentVerification) SetStartDate(v string) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *EmploymentVerification) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *EmploymentVerification) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmploymentVerification) GetEndDate() string {
	if o == nil || o.EndDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmploymentVerification) GetEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *EmploymentVerification) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableString and assigns it to the EndDate field.
func (o *EmploymentVerification) SetEndDate(v string) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *EmploymentVerification) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *EmploymentVerification) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetEmployer returns the Employer field value if set, zero value otherwise.
func (o *EmploymentVerification) GetEmployer() EmployerVerification {
	if o == nil || o.Employer == nil {
		var ret EmployerVerification
		return ret
	}
	return *o.Employer
}

// GetEmployerOk returns a tuple with the Employer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmploymentVerification) GetEmployerOk() (*EmployerVerification, bool) {
	if o == nil || o.Employer == nil {
		return nil, false
	}
	return o.Employer, true
}

// HasEmployer returns a boolean if a field has been set.
func (o *EmploymentVerification) HasEmployer() bool {
	if o != nil && o.Employer != nil {
		return true
	}

	return false
}

// SetEmployer gets a reference to the given EmployerVerification and assigns it to the Employer field.
func (o *EmploymentVerification) SetEmployer(v EmployerVerification) {
	o.Employer = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmploymentVerification) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmploymentVerification) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *EmploymentVerification) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *EmploymentVerification) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *EmploymentVerification) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *EmploymentVerification) UnsetTitle() {
	o.Title.Unset()
}

// GetPlatformIds returns the PlatformIds field value if set, zero value otherwise.
func (o *EmploymentVerification) GetPlatformIds() PlatformIds {
	if o == nil || o.PlatformIds == nil {
		var ret PlatformIds
		return ret
	}
	return *o.PlatformIds
}

// GetPlatformIdsOk returns a tuple with the PlatformIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmploymentVerification) GetPlatformIdsOk() (*PlatformIds, bool) {
	if o == nil || o.PlatformIds == nil {
		return nil, false
	}
	return o.PlatformIds, true
}

// HasPlatformIds returns a boolean if a field has been set.
func (o *EmploymentVerification) HasPlatformIds() bool {
	if o != nil && o.PlatformIds != nil {
		return true
	}

	return false
}

// SetPlatformIds gets a reference to the given PlatformIds and assigns it to the PlatformIds field.
func (o *EmploymentVerification) SetPlatformIds(v PlatformIds) {
	o.PlatformIds = &v
}

func (o EmploymentVerification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.StartDate.IsSet() {
		toSerialize["start_date"] = o.StartDate.Get()
	}
	if o.EndDate.IsSet() {
		toSerialize["end_date"] = o.EndDate.Get()
	}
	if o.Employer != nil {
		toSerialize["employer"] = o.Employer
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.PlatformIds != nil {
		toSerialize["platform_ids"] = o.PlatformIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EmploymentVerification) UnmarshalJSON(bytes []byte) (err error) {
	varEmploymentVerification := _EmploymentVerification{}

	if err = json.Unmarshal(bytes, &varEmploymentVerification); err == nil {
		*o = EmploymentVerification(varEmploymentVerification)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "start_date")
		delete(additionalProperties, "end_date")
		delete(additionalProperties, "employer")
		delete(additionalProperties, "title")
		delete(additionalProperties, "platform_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmploymentVerification struct {
	value *EmploymentVerification
	isSet bool
}

func (v NullableEmploymentVerification) Get() *EmploymentVerification {
	return v.value
}

func (v *NullableEmploymentVerification) Set(val *EmploymentVerification) {
	v.value = val
	v.isSet = true
}

func (v NullableEmploymentVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableEmploymentVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmploymentVerification(val *EmploymentVerification) *NullableEmploymentVerification {
	return &NullableEmploymentVerification{value: val, isSet: true}
}

func (v NullableEmploymentVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmploymentVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


