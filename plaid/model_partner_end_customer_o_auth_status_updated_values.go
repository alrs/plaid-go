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

// PartnerEndCustomerOAuthStatusUpdatedValues The OAuth status of the update
type PartnerEndCustomerOAuthStatusUpdatedValues string

var _ = fmt.Printf

// List of PartnerEndCustomerOAuthStatusUpdatedValues
const (
	PARTNERENDCUSTOMEROAUTHSTATUSUPDATEDVALUES_NOT_STARTED PartnerEndCustomerOAuthStatusUpdatedValues = "not-started"
	PARTNERENDCUSTOMEROAUTHSTATUSUPDATEDVALUES_PROCESSING PartnerEndCustomerOAuthStatusUpdatedValues = "processing"
	PARTNERENDCUSTOMEROAUTHSTATUSUPDATEDVALUES_APPROVED PartnerEndCustomerOAuthStatusUpdatedValues = "approved"
	PARTNERENDCUSTOMEROAUTHSTATUSUPDATEDVALUES_ENABLED PartnerEndCustomerOAuthStatusUpdatedValues = "enabled"
	PARTNERENDCUSTOMEROAUTHSTATUSUPDATEDVALUES_ATTENTION_REQUIRED PartnerEndCustomerOAuthStatusUpdatedValues = "attention-required"
)

var allowedPartnerEndCustomerOAuthStatusUpdatedValuesEnumValues = []PartnerEndCustomerOAuthStatusUpdatedValues{
	"not-started",
	"processing",
	"approved",
	"enabled",
	"attention-required",
}

func (v *PartnerEndCustomerOAuthStatusUpdatedValues) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := PartnerEndCustomerOAuthStatusUpdatedValues(value)


	*v = enumTypeValue
	return nil
}

// NewPartnerEndCustomerOAuthStatusUpdatedValuesFromValue returns a pointer to a valid PartnerEndCustomerOAuthStatusUpdatedValues
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPartnerEndCustomerOAuthStatusUpdatedValuesFromValue(v string) (*PartnerEndCustomerOAuthStatusUpdatedValues, error) {
	ev := PartnerEndCustomerOAuthStatusUpdatedValues(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PartnerEndCustomerOAuthStatusUpdatedValues) IsValid() bool {
	for _, existing := range allowedPartnerEndCustomerOAuthStatusUpdatedValuesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PartnerEndCustomerOAuthStatusUpdatedValues value
func (v PartnerEndCustomerOAuthStatusUpdatedValues) Ptr() *PartnerEndCustomerOAuthStatusUpdatedValues {
	return &v
}

type NullablePartnerEndCustomerOAuthStatusUpdatedValues struct {
	value *PartnerEndCustomerOAuthStatusUpdatedValues
	isSet bool
}

func (v NullablePartnerEndCustomerOAuthStatusUpdatedValues) Get() *PartnerEndCustomerOAuthStatusUpdatedValues {
	return v.value
}

func (v *NullablePartnerEndCustomerOAuthStatusUpdatedValues) Set(val *PartnerEndCustomerOAuthStatusUpdatedValues) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerEndCustomerOAuthStatusUpdatedValues) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerEndCustomerOAuthStatusUpdatedValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerEndCustomerOAuthStatusUpdatedValues(val *PartnerEndCustomerOAuthStatusUpdatedValues) *NullablePartnerEndCustomerOAuthStatusUpdatedValues {
	return &NullablePartnerEndCustomerOAuthStatusUpdatedValues{value: val, isSet: true}
}

func (v NullablePartnerEndCustomerOAuthStatusUpdatedValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerEndCustomerOAuthStatusUpdatedValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

