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

// CreditFreddieMacReportingInformationVOA24 Information about an report identifier and a report name.
type CreditFreddieMacReportingInformationVOA24 struct {
	// Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
	ReportDateTime *string `json:"ReportDateTime,omitempty"`
	// Documentation not found in the MISMO model viewer and not provided by Freddie Mac. The value can only be \"ReportID\"
	ReportIdentifierType *string `json:"ReportIdentifierType,omitempty"`
	// Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
	ReportingInformationIdentifier string `json:"ReportingInformationIdentifier"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacReportingInformationVOA24 CreditFreddieMacReportingInformationVOA24

// NewCreditFreddieMacReportingInformationVOA24 instantiates a new CreditFreddieMacReportingInformationVOA24 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacReportingInformationVOA24(reportingInformationIdentifier string) *CreditFreddieMacReportingInformationVOA24 {
	this := CreditFreddieMacReportingInformationVOA24{}
	this.ReportingInformationIdentifier = reportingInformationIdentifier
	return &this
}

// NewCreditFreddieMacReportingInformationVOA24WithDefaults instantiates a new CreditFreddieMacReportingInformationVOA24 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacReportingInformationVOA24WithDefaults() *CreditFreddieMacReportingInformationVOA24 {
	this := CreditFreddieMacReportingInformationVOA24{}
	return &this
}

// GetReportDateTime returns the ReportDateTime field value if set, zero value otherwise.
func (o *CreditFreddieMacReportingInformationVOA24) GetReportDateTime() string {
	if o == nil || o.ReportDateTime == nil {
		var ret string
		return ret
	}
	return *o.ReportDateTime
}

// GetReportDateTimeOk returns a tuple with the ReportDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacReportingInformationVOA24) GetReportDateTimeOk() (*string, bool) {
	if o == nil || o.ReportDateTime == nil {
		return nil, false
	}
	return o.ReportDateTime, true
}

// HasReportDateTime returns a boolean if a field has been set.
func (o *CreditFreddieMacReportingInformationVOA24) HasReportDateTime() bool {
	if o != nil && o.ReportDateTime != nil {
		return true
	}

	return false
}

// SetReportDateTime gets a reference to the given string and assigns it to the ReportDateTime field.
func (o *CreditFreddieMacReportingInformationVOA24) SetReportDateTime(v string) {
	o.ReportDateTime = &v
}

// GetReportIdentifierType returns the ReportIdentifierType field value if set, zero value otherwise.
func (o *CreditFreddieMacReportingInformationVOA24) GetReportIdentifierType() string {
	if o == nil || o.ReportIdentifierType == nil {
		var ret string
		return ret
	}
	return *o.ReportIdentifierType
}

// GetReportIdentifierTypeOk returns a tuple with the ReportIdentifierType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacReportingInformationVOA24) GetReportIdentifierTypeOk() (*string, bool) {
	if o == nil || o.ReportIdentifierType == nil {
		return nil, false
	}
	return o.ReportIdentifierType, true
}

// HasReportIdentifierType returns a boolean if a field has been set.
func (o *CreditFreddieMacReportingInformationVOA24) HasReportIdentifierType() bool {
	if o != nil && o.ReportIdentifierType != nil {
		return true
	}

	return false
}

// SetReportIdentifierType gets a reference to the given string and assigns it to the ReportIdentifierType field.
func (o *CreditFreddieMacReportingInformationVOA24) SetReportIdentifierType(v string) {
	o.ReportIdentifierType = &v
}

// GetReportingInformationIdentifier returns the ReportingInformationIdentifier field value
func (o *CreditFreddieMacReportingInformationVOA24) GetReportingInformationIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportingInformationIdentifier
}

// GetReportingInformationIdentifierOk returns a tuple with the ReportingInformationIdentifier field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacReportingInformationVOA24) GetReportingInformationIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReportingInformationIdentifier, true
}

// SetReportingInformationIdentifier sets field value
func (o *CreditFreddieMacReportingInformationVOA24) SetReportingInformationIdentifier(v string) {
	o.ReportingInformationIdentifier = v
}

func (o CreditFreddieMacReportingInformationVOA24) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReportDateTime != nil {
		toSerialize["ReportDateTime"] = o.ReportDateTime
	}
	if o.ReportIdentifierType != nil {
		toSerialize["ReportIdentifierType"] = o.ReportIdentifierType
	}
	if true {
		toSerialize["ReportingInformationIdentifier"] = o.ReportingInformationIdentifier
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacReportingInformationVOA24) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacReportingInformationVOA24 := _CreditFreddieMacReportingInformationVOA24{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacReportingInformationVOA24); err == nil {
		*o = CreditFreddieMacReportingInformationVOA24(varCreditFreddieMacReportingInformationVOA24)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ReportDateTime")
		delete(additionalProperties, "ReportIdentifierType")
		delete(additionalProperties, "ReportingInformationIdentifier")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieMacReportingInformationVOA24 struct {
	value *CreditFreddieMacReportingInformationVOA24
	isSet bool
}

func (v NullableCreditFreddieMacReportingInformationVOA24) Get() *CreditFreddieMacReportingInformationVOA24 {
	return v.value
}

func (v *NullableCreditFreddieMacReportingInformationVOA24) Set(val *CreditFreddieMacReportingInformationVOA24) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacReportingInformationVOA24) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacReportingInformationVOA24) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacReportingInformationVOA24(val *CreditFreddieMacReportingInformationVOA24) *NullableCreditFreddieMacReportingInformationVOA24 {
	return &NullableCreditFreddieMacReportingInformationVOA24{value: val, isSet: true}
}

func (v NullableCreditFreddieMacReportingInformationVOA24) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacReportingInformationVOA24) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


