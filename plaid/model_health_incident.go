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
	"time"
)

// HealthIncident struct for HealthIncident
type HealthIncident struct {
	// The start date of the incident, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format, e.g. `\"2020-10-30T15:26:48Z\"`.
	StartDate time.Time `json:"start_date"`
	// The end date of the incident, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format, e.g. `\"2020-10-30T15:26:48Z\"`.
	EndDate *time.Time `json:"end_date,omitempty"`
	// The title of the incident
	Title string `json:"title"`
	// Updates on the health incident.
	IncidentUpdates []IncidentUpdate `json:"incident_updates"`
	AdditionalProperties map[string]interface{}
}

type _HealthIncident HealthIncident

// NewHealthIncident instantiates a new HealthIncident object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthIncident(startDate time.Time, title string, incidentUpdates []IncidentUpdate) *HealthIncident {
	this := HealthIncident{}
	this.StartDate = startDate
	this.Title = title
	this.IncidentUpdates = incidentUpdates
	return &this
}

// NewHealthIncidentWithDefaults instantiates a new HealthIncident object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthIncidentWithDefaults() *HealthIncident {
	this := HealthIncident{}
	return &this
}

// GetStartDate returns the StartDate field value
func (o *HealthIncident) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *HealthIncident) GetStartDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *HealthIncident) SetStartDate(v time.Time) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *HealthIncident) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthIncident) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *HealthIncident) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *HealthIncident) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetTitle returns the Title field value
func (o *HealthIncident) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *HealthIncident) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *HealthIncident) SetTitle(v string) {
	o.Title = v
}

// GetIncidentUpdates returns the IncidentUpdates field value
func (o *HealthIncident) GetIncidentUpdates() []IncidentUpdate {
	if o == nil {
		var ret []IncidentUpdate
		return ret
	}

	return o.IncidentUpdates
}

// GetIncidentUpdatesOk returns a tuple with the IncidentUpdates field value
// and a boolean to check if the value has been set.
func (o *HealthIncident) GetIncidentUpdatesOk() (*[]IncidentUpdate, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IncidentUpdates, true
}

// SetIncidentUpdates sets field value
func (o *HealthIncident) SetIncidentUpdates(v []IncidentUpdate) {
	o.IncidentUpdates = v
}

func (o HealthIncident) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["start_date"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["end_date"] = o.EndDate
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["incident_updates"] = o.IncidentUpdates
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HealthIncident) UnmarshalJSON(bytes []byte) (err error) {
	varHealthIncident := _HealthIncident{}

	if err = json.Unmarshal(bytes, &varHealthIncident); err == nil {
		*o = HealthIncident(varHealthIncident)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "start_date")
		delete(additionalProperties, "end_date")
		delete(additionalProperties, "title")
		delete(additionalProperties, "incident_updates")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHealthIncident struct {
	value *HealthIncident
	isSet bool
}

func (v NullableHealthIncident) Get() *HealthIncident {
	return v.value
}

func (v *NullableHealthIncident) Set(val *HealthIncident) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthIncident) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthIncident) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthIncident(val *HealthIncident) *NullableHealthIncident {
	return &NullableHealthIncident{value: val, isSet: true}
}

func (v NullableHealthIncident) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthIncident) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


