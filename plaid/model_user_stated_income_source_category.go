/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.343.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// UserStatedIncomeSourceCategory The income category for a specified income source
type UserStatedIncomeSourceCategory string

var _ = fmt.Printf

// List of UserStatedIncomeSourceCategory
const (
	USERSTATEDINCOMESOURCECATEGORY_OTHER UserStatedIncomeSourceCategory = "OTHER"
	USERSTATEDINCOMESOURCECATEGORY_SALARY UserStatedIncomeSourceCategory = "SALARY"
	USERSTATEDINCOMESOURCECATEGORY_UNEMPLOYMENT UserStatedIncomeSourceCategory = "UNEMPLOYMENT"
	USERSTATEDINCOMESOURCECATEGORY_CASH UserStatedIncomeSourceCategory = "CASH"
	USERSTATEDINCOMESOURCECATEGORY_GIG_ECONOMY UserStatedIncomeSourceCategory = "GIG_ECONOMY"
	USERSTATEDINCOMESOURCECATEGORY_RENTAL UserStatedIncomeSourceCategory = "RENTAL"
	USERSTATEDINCOMESOURCECATEGORY_CHILD_SUPPORT UserStatedIncomeSourceCategory = "CHILD_SUPPORT"
	USERSTATEDINCOMESOURCECATEGORY_MILITARY UserStatedIncomeSourceCategory = "MILITARY"
	USERSTATEDINCOMESOURCECATEGORY_RETIREMENT UserStatedIncomeSourceCategory = "RETIREMENT"
	USERSTATEDINCOMESOURCECATEGORY_LONG_TERM_DISABILITY UserStatedIncomeSourceCategory = "LONG_TERM_DISABILITY"
	USERSTATEDINCOMESOURCECATEGORY_BANK_INTEREST UserStatedIncomeSourceCategory = "BANK_INTEREST"
)

var allowedUserStatedIncomeSourceCategoryEnumValues = []UserStatedIncomeSourceCategory{
	"OTHER",
	"SALARY",
	"UNEMPLOYMENT",
	"CASH",
	"GIG_ECONOMY",
	"RENTAL",
	"CHILD_SUPPORT",
	"MILITARY",
	"RETIREMENT",
	"LONG_TERM_DISABILITY",
	"BANK_INTEREST",
}

func (v *UserStatedIncomeSourceCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := UserStatedIncomeSourceCategory(value)


	*v = enumTypeValue
	return nil
}

// NewUserStatedIncomeSourceCategoryFromValue returns a pointer to a valid UserStatedIncomeSourceCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserStatedIncomeSourceCategoryFromValue(v string) (*UserStatedIncomeSourceCategory, error) {
	ev := UserStatedIncomeSourceCategory(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserStatedIncomeSourceCategory) IsValid() bool {
	for _, existing := range allowedUserStatedIncomeSourceCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserStatedIncomeSourceCategory value
func (v UserStatedIncomeSourceCategory) Ptr() *UserStatedIncomeSourceCategory {
	return &v
}

type NullableUserStatedIncomeSourceCategory struct {
	value *UserStatedIncomeSourceCategory
	isSet bool
}

func (v NullableUserStatedIncomeSourceCategory) Get() *UserStatedIncomeSourceCategory {
	return v.value
}

func (v *NullableUserStatedIncomeSourceCategory) Set(val *UserStatedIncomeSourceCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableUserStatedIncomeSourceCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableUserStatedIncomeSourceCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserStatedIncomeSourceCategory(val *UserStatedIncomeSourceCategory) *NullableUserStatedIncomeSourceCategory {
	return &NullableUserStatedIncomeSourceCategory{value: val, isSet: true}
}

func (v NullableUserStatedIncomeSourceCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserStatedIncomeSourceCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

