/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.419.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// PartnerEndCustomerQuestionnaireStatus The status of the end customer's security questionnaire.
type PartnerEndCustomerQuestionnaireStatus string

var _ = fmt.Printf

// List of PartnerEndCustomerQuestionnaireStatus
const (
	PARTNERENDCUSTOMERQUESTIONNAIRESTATUS_NOT_STARTED PartnerEndCustomerQuestionnaireStatus = "NOT_STARTED"
	PARTNERENDCUSTOMERQUESTIONNAIRESTATUS_RECEIVED PartnerEndCustomerQuestionnaireStatus = "RECEIVED"
	PARTNERENDCUSTOMERQUESTIONNAIRESTATUS_COMPLETE PartnerEndCustomerQuestionnaireStatus = "COMPLETE"
)

var allowedPartnerEndCustomerQuestionnaireStatusEnumValues = []PartnerEndCustomerQuestionnaireStatus{
	"NOT_STARTED",
	"RECEIVED",
	"COMPLETE",
}

func (v *PartnerEndCustomerQuestionnaireStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := PartnerEndCustomerQuestionnaireStatus(value)


	*v = enumTypeValue
	return nil
}

// NewPartnerEndCustomerQuestionnaireStatusFromValue returns a pointer to a valid PartnerEndCustomerQuestionnaireStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPartnerEndCustomerQuestionnaireStatusFromValue(v string) (*PartnerEndCustomerQuestionnaireStatus, error) {
	ev := PartnerEndCustomerQuestionnaireStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PartnerEndCustomerQuestionnaireStatus) IsValid() bool {
	for _, existing := range allowedPartnerEndCustomerQuestionnaireStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PartnerEndCustomerQuestionnaireStatus value
func (v PartnerEndCustomerQuestionnaireStatus) Ptr() *PartnerEndCustomerQuestionnaireStatus {
	return &v
}

type NullablePartnerEndCustomerQuestionnaireStatus struct {
	value *PartnerEndCustomerQuestionnaireStatus
	isSet bool
}

func (v NullablePartnerEndCustomerQuestionnaireStatus) Get() *PartnerEndCustomerQuestionnaireStatus {
	return v.value
}

func (v *NullablePartnerEndCustomerQuestionnaireStatus) Set(val *PartnerEndCustomerQuestionnaireStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerEndCustomerQuestionnaireStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerEndCustomerQuestionnaireStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerEndCustomerQuestionnaireStatus(val *PartnerEndCustomerQuestionnaireStatus) *NullablePartnerEndCustomerQuestionnaireStatus {
	return &NullablePartnerEndCustomerQuestionnaireStatus{value: val, isSet: true}
}

func (v NullablePartnerEndCustomerQuestionnaireStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerEndCustomerQuestionnaireStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

