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
	"fmt"
)

// RiskCheckEmailDomainIsFreeProvider Indicates whether the email address domain is a free provider such as Gmail or Hotmail if known.
type RiskCheckEmailDomainIsFreeProvider string

var _ = fmt.Printf

// List of RiskCheckEmailDomainIsFreeProvider
const (
	RISKCHECKEMAILDOMAINISFREEPROVIDER_YES RiskCheckEmailDomainIsFreeProvider = "yes"
	RISKCHECKEMAILDOMAINISFREEPROVIDER_NO RiskCheckEmailDomainIsFreeProvider = "no"
	RISKCHECKEMAILDOMAINISFREEPROVIDER_NO_DATA RiskCheckEmailDomainIsFreeProvider = "no_data"
)

var allowedRiskCheckEmailDomainIsFreeProviderEnumValues = []RiskCheckEmailDomainIsFreeProvider{
	"yes",
	"no",
	"no_data",
}

func (v *RiskCheckEmailDomainIsFreeProvider) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := RiskCheckEmailDomainIsFreeProvider(value)


	*v = enumTypeValue
	return nil
}

// NewRiskCheckEmailDomainIsFreeProviderFromValue returns a pointer to a valid RiskCheckEmailDomainIsFreeProvider
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRiskCheckEmailDomainIsFreeProviderFromValue(v string) (*RiskCheckEmailDomainIsFreeProvider, error) {
	ev := RiskCheckEmailDomainIsFreeProvider(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RiskCheckEmailDomainIsFreeProvider) IsValid() bool {
	for _, existing := range allowedRiskCheckEmailDomainIsFreeProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RiskCheckEmailDomainIsFreeProvider value
func (v RiskCheckEmailDomainIsFreeProvider) Ptr() *RiskCheckEmailDomainIsFreeProvider {
	return &v
}

type NullableRiskCheckEmailDomainIsFreeProvider struct {
	value *RiskCheckEmailDomainIsFreeProvider
	isSet bool
}

func (v NullableRiskCheckEmailDomainIsFreeProvider) Get() *RiskCheckEmailDomainIsFreeProvider {
	return v.value
}

func (v *NullableRiskCheckEmailDomainIsFreeProvider) Set(val *RiskCheckEmailDomainIsFreeProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskCheckEmailDomainIsFreeProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskCheckEmailDomainIsFreeProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskCheckEmailDomainIsFreeProvider(val *RiskCheckEmailDomainIsFreeProvider) *NullableRiskCheckEmailDomainIsFreeProvider {
	return &NullableRiskCheckEmailDomainIsFreeProvider{value: val, isSet: true}
}

func (v NullableRiskCheckEmailDomainIsFreeProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskCheckEmailDomainIsFreeProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

