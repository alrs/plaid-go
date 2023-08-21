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

// CreditBankIncomeCategory The income category. Note that the `CASH` value has been deprecated and is used only for existing legacy implementations. It has been replaced by the new categories `CASH_DEPOSIT` (representing cash or check deposits) and `TRANSFER_FROM_APPLICATION` (representing cash transfers originating from apps, such as Zelle or Venmo).
type CreditBankIncomeCategory string

var _ = fmt.Printf

// List of CreditBankIncomeCategory
const (
	CREDITBANKINCOMECATEGORY_SALARY CreditBankIncomeCategory = "SALARY"
	CREDITBANKINCOMECATEGORY_UNEMPLOYMENT CreditBankIncomeCategory = "UNEMPLOYMENT"
	CREDITBANKINCOMECATEGORY_CASH CreditBankIncomeCategory = "CASH"
	CREDITBANKINCOMECATEGORY_GIG_ECONOMY CreditBankIncomeCategory = "GIG_ECONOMY"
	CREDITBANKINCOMECATEGORY_RENTAL CreditBankIncomeCategory = "RENTAL"
	CREDITBANKINCOMECATEGORY_CHILD_SUPPORT CreditBankIncomeCategory = "CHILD_SUPPORT"
	CREDITBANKINCOMECATEGORY_MILITARY CreditBankIncomeCategory = "MILITARY"
	CREDITBANKINCOMECATEGORY_RETIREMENT CreditBankIncomeCategory = "RETIREMENT"
	CREDITBANKINCOMECATEGORY_LONG_TERM_DISABILITY CreditBankIncomeCategory = "LONG_TERM_DISABILITY"
	CREDITBANKINCOMECATEGORY_BANK_INTEREST CreditBankIncomeCategory = "BANK_INTEREST"
	CREDITBANKINCOMECATEGORY_CASH_DEPOSIT CreditBankIncomeCategory = "CASH_DEPOSIT"
	CREDITBANKINCOMECATEGORY_TRANSFER_FROM_APPLICATION CreditBankIncomeCategory = "TRANSFER_FROM_APPLICATION"
	CREDITBANKINCOMECATEGORY_TAX_REFUND CreditBankIncomeCategory = "TAX_REFUND"
	CREDITBANKINCOMECATEGORY_BENEFIT_OTHER CreditBankIncomeCategory = "BENEFIT_OTHER"
	CREDITBANKINCOMECATEGORY_OTHER CreditBankIncomeCategory = "OTHER"
)

var allowedCreditBankIncomeCategoryEnumValues = []CreditBankIncomeCategory{
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
	"CASH_DEPOSIT",
	"TRANSFER_FROM_APPLICATION",
	"TAX_REFUND",
	"BENEFIT_OTHER",
	"OTHER",
}

func (v *CreditBankIncomeCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := CreditBankIncomeCategory(value)


	*v = enumTypeValue
	return nil
}

// NewCreditBankIncomeCategoryFromValue returns a pointer to a valid CreditBankIncomeCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreditBankIncomeCategoryFromValue(v string) (*CreditBankIncomeCategory, error) {
	ev := CreditBankIncomeCategory(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreditBankIncomeCategory) IsValid() bool {
	for _, existing := range allowedCreditBankIncomeCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreditBankIncomeCategory value
func (v CreditBankIncomeCategory) Ptr() *CreditBankIncomeCategory {
	return &v
}

type NullableCreditBankIncomeCategory struct {
	value *CreditBankIncomeCategory
	isSet bool
}

func (v NullableCreditBankIncomeCategory) Get() *CreditBankIncomeCategory {
	return v.value
}

func (v *NullableCreditBankIncomeCategory) Set(val *CreditBankIncomeCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankIncomeCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankIncomeCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankIncomeCategory(val *CreditBankIncomeCategory) *NullableCreditBankIncomeCategory {
	return &NullableCreditBankIncomeCategory{value: val, isSet: true}
}

func (v NullableCreditBankIncomeCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankIncomeCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

