/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.181.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// AssetTransactionCategoryType Asset Transaction Category Type Enumerated derived by Vendor.
type AssetTransactionCategoryType string

// List of AssetTransactionCategoryType
const (
	ASSETTRANSACTIONCATEGORYTYPE_ATM_FEE AssetTransactionCategoryType = "ATMFee"
	ASSETTRANSACTIONCATEGORYTYPE_ADVERTISING AssetTransactionCategoryType = "Advertising"
	ASSETTRANSACTIONCATEGORYTYPE_AIR_TRAVEL AssetTransactionCategoryType = "AirTravel"
	ASSETTRANSACTIONCATEGORYTYPE_ALCOHOL_BARS AssetTransactionCategoryType = "AlcoholBars"
	ASSETTRANSACTIONCATEGORYTYPE_ALLOWANCE AssetTransactionCategoryType = "Allowance"
	ASSETTRANSACTIONCATEGORYTYPE_AMUSEMENT AssetTransactionCategoryType = "Amusement"
	ASSETTRANSACTIONCATEGORYTYPE_ARTS AssetTransactionCategoryType = "Arts"
	ASSETTRANSACTIONCATEGORYTYPE_AUTO_TRANSPORT AssetTransactionCategoryType = "AutoTransport"
	ASSETTRANSACTIONCATEGORYTYPE_AUTO_INSURANCE AssetTransactionCategoryType = "AutoInsurance"
	ASSETTRANSACTIONCATEGORYTYPE_AUTO_PAYMENT AssetTransactionCategoryType = "AutoPayment"
	ASSETTRANSACTIONCATEGORYTYPE_BABY_SUPPLIES AssetTransactionCategoryType = "BabySupplies"
	ASSETTRANSACTIONCATEGORYTYPE_BABYSITTER_DAYCARE AssetTransactionCategoryType = "BabysitterDaycare"
	ASSETTRANSACTIONCATEGORYTYPE_BANK_FEE AssetTransactionCategoryType = "BankFee"
	ASSETTRANSACTIONCATEGORYTYPE_BILLS_UTILITIES AssetTransactionCategoryType = "BillsUtilities"
	ASSETTRANSACTIONCATEGORYTYPE_BONUS AssetTransactionCategoryType = "Bonus"
	ASSETTRANSACTIONCATEGORYTYPE_BOOKS_SUPPLIES AssetTransactionCategoryType = "BooksSupplies"
	ASSETTRANSACTIONCATEGORYTYPE_BUSINESS_SERVICES AssetTransactionCategoryType = "Business Services"
	ASSETTRANSACTIONCATEGORYTYPE_BUY AssetTransactionCategoryType = "Buy"
	ASSETTRANSACTIONCATEGORYTYPE_CASH_ATM AssetTransactionCategoryType = "CashATM"
	ASSETTRANSACTIONCATEGORYTYPE_CHARITY AssetTransactionCategoryType = "Charity"
	ASSETTRANSACTIONCATEGORYTYPE_CHECK AssetTransactionCategoryType = "Check"
	ASSETTRANSACTIONCATEGORYTYPE_CHILD_SUPPORT AssetTransactionCategoryType = "ChildSupport"
	ASSETTRANSACTIONCATEGORYTYPE_CLOTHING AssetTransactionCategoryType = "Clothing"
	ASSETTRANSACTIONCATEGORYTYPE_COFFEE_SHOPS AssetTransactionCategoryType = "CoffeeShops"
	ASSETTRANSACTIONCATEGORYTYPE_CREDIT_CARD_PAYMENT AssetTransactionCategoryType = "CreditCardPayment"
	ASSETTRANSACTIONCATEGORYTYPE_DENTIST AssetTransactionCategoryType = "Dentist"
	ASSETTRANSACTIONCATEGORYTYPE_DOCTOR AssetTransactionCategoryType = "Doctor"
	ASSETTRANSACTIONCATEGORYTYPE_EDUCATION AssetTransactionCategoryType = "Education"
	ASSETTRANSACTIONCATEGORYTYPE_ELECTRONICS_SOFTWARE AssetTransactionCategoryType = "ElectronicsSoftware"
	ASSETTRANSACTIONCATEGORYTYPE_ENTERTAINMENT AssetTransactionCategoryType = "Entertainment"
	ASSETTRANSACTIONCATEGORYTYPE_EYECARE AssetTransactionCategoryType = "Eyecare"
	ASSETTRANSACTIONCATEGORYTYPE_FAST_FOOD AssetTransactionCategoryType = "FastFood"
	ASSETTRANSACTIONCATEGORYTYPE_FEDERAL_TAX AssetTransactionCategoryType = "FederalTax"
	ASSETTRANSACTIONCATEGORYTYPE_FEES_CHARGES AssetTransactionCategoryType = "FeesCharges"
	ASSETTRANSACTIONCATEGORYTYPE_FINANCE_CHARGE AssetTransactionCategoryType = "FinanceCharge"
	ASSETTRANSACTIONCATEGORYTYPE_FINANCIAL AssetTransactionCategoryType = "Financial"
	ASSETTRANSACTIONCATEGORYTYPE_FINANCIAL_ADVISOR AssetTransactionCategoryType = "FinancialAdvisor"
	ASSETTRANSACTIONCATEGORYTYPE_FOOD_DINING AssetTransactionCategoryType = "FoodDining"
	ASSETTRANSACTIONCATEGORYTYPE_FURNISHINGS AssetTransactionCategoryType = "Furnishings"
	ASSETTRANSACTIONCATEGORYTYPE_GAS_FUEL AssetTransactionCategoryType = "GasFuel"
	ASSETTRANSACTIONCATEGORYTYPE_GIFTS_DONATIONS AssetTransactionCategoryType = "GiftsDonations"
	ASSETTRANSACTIONCATEGORYTYPE_GROCERIES AssetTransactionCategoryType = "Groceries"
	ASSETTRANSACTIONCATEGORYTYPE_GYM AssetTransactionCategoryType = "Gym"
	ASSETTRANSACTIONCATEGORYTYPE_HAIR AssetTransactionCategoryType = "Hair"
	ASSETTRANSACTIONCATEGORYTYPE_HEALTH_FITNESS AssetTransactionCategoryType = "HealthFitness"
	ASSETTRANSACTIONCATEGORYTYPE_HEALTH_INSURANCE AssetTransactionCategoryType = "HealthInsurance"
	ASSETTRANSACTIONCATEGORYTYPE_HOBBIES AssetTransactionCategoryType = "Hobbies"
	ASSETTRANSACTIONCATEGORYTYPE_HOME AssetTransactionCategoryType = "Home"
	ASSETTRANSACTIONCATEGORYTYPE_HOME_IMPROVEMENT AssetTransactionCategoryType = "HomeImprovement"
	ASSETTRANSACTIONCATEGORYTYPE_HOME_INSURANCE AssetTransactionCategoryType = "HomeInsurance"
	ASSETTRANSACTIONCATEGORYTYPE_HOME_PHONE AssetTransactionCategoryType = "HomePhone"
	ASSETTRANSACTIONCATEGORYTYPE_HOME_SERVICES AssetTransactionCategoryType = "HomeServices"
	ASSETTRANSACTIONCATEGORYTYPE_HOME_SUPPLIES AssetTransactionCategoryType = "HomeSupplies"
	ASSETTRANSACTIONCATEGORYTYPE_HOTEL AssetTransactionCategoryType = "Hotel"
	ASSETTRANSACTIONCATEGORYTYPE_INCOME AssetTransactionCategoryType = "Income"
	ASSETTRANSACTIONCATEGORYTYPE_INTEREST_INCOME AssetTransactionCategoryType = "InterestIncome"
	ASSETTRANSACTIONCATEGORYTYPE_INTERNET AssetTransactionCategoryType = "Internet"
	ASSETTRANSACTIONCATEGORYTYPE_INVESTMENTS AssetTransactionCategoryType = "Investments"
	ASSETTRANSACTIONCATEGORYTYPE_KIDS AssetTransactionCategoryType = "Kids"
	ASSETTRANSACTIONCATEGORYTYPE_KIDS_ACTIVITIES AssetTransactionCategoryType = "KidsActivities"
	ASSETTRANSACTIONCATEGORYTYPE_LATE_FEE AssetTransactionCategoryType = "LateFee"
	ASSETTRANSACTIONCATEGORYTYPE_LAUNDRY AssetTransactionCategoryType = "Laundry"
	ASSETTRANSACTIONCATEGORYTYPE_LAWN_GARDEN AssetTransactionCategoryType = "LawnGarden"
	ASSETTRANSACTIONCATEGORYTYPE_LEGAL AssetTransactionCategoryType = "Legal"
	ASSETTRANSACTIONCATEGORYTYPE_LIFE_INSURANCE AssetTransactionCategoryType = "LifeInsurance"
	ASSETTRANSACTIONCATEGORYTYPE_LOAN_INSURANCE AssetTransactionCategoryType = "LoanInsurance"
	ASSETTRANSACTIONCATEGORYTYPE_LOAN_PAYMENT AssetTransactionCategoryType = "LoanPayment"
	ASSETTRANSACTIONCATEGORYTYPE_LOANS AssetTransactionCategoryType = "Loans"
	ASSETTRANSACTIONCATEGORYTYPE_MOBILE_PHONE AssetTransactionCategoryType = "MobilePhone"
	ASSETTRANSACTIONCATEGORYTYPE_MORTGAGE_RENT AssetTransactionCategoryType = "MortgageRent"
	ASSETTRANSACTIONCATEGORYTYPE_MOVIES_DVDS AssetTransactionCategoryType = "MoviesDVDs"
	ASSETTRANSACTIONCATEGORYTYPE_MUSIC AssetTransactionCategoryType = "Music"
	ASSETTRANSACTIONCATEGORYTYPE_NEWSPAPERS_MAGAZINES AssetTransactionCategoryType = "NewspapersMagazines"
	ASSETTRANSACTIONCATEGORYTYPE_OFFICE_SUPPLIES AssetTransactionCategoryType = "OfficeSupplies"
	ASSETTRANSACTIONCATEGORYTYPE_PARKING AssetTransactionCategoryType = "Parking"
	ASSETTRANSACTIONCATEGORYTYPE_PAYCHECK AssetTransactionCategoryType = "Paycheck"
	ASSETTRANSACTIONCATEGORYTYPE_PERSONAL_CARE AssetTransactionCategoryType = "PersonalCare"
	ASSETTRANSACTIONCATEGORYTYPE_PET_FOOD_SUPPLIES AssetTransactionCategoryType = "PetFoodSupplies"
	ASSETTRANSACTIONCATEGORYTYPE_PET_GROOMING AssetTransactionCategoryType = "PetGrooming"
	ASSETTRANSACTIONCATEGORYTYPE_PETS AssetTransactionCategoryType = "Pets"
	ASSETTRANSACTIONCATEGORYTYPE_PHARMACY AssetTransactionCategoryType = "Pharmacy"
	ASSETTRANSACTIONCATEGORYTYPE_PRINTING AssetTransactionCategoryType = "Printing"
	ASSETTRANSACTIONCATEGORYTYPE_PROPERTY_TAX AssetTransactionCategoryType = "Property Tax"
	ASSETTRANSACTIONCATEGORYTYPE_PUBLIC_TRANSPORTATION AssetTransactionCategoryType = "Public Transportation"
	ASSETTRANSACTIONCATEGORYTYPE_REIMBURSEMENT AssetTransactionCategoryType = "Reimbursement"
	ASSETTRANSACTIONCATEGORYTYPE_RENTAL_CAR_TAXI AssetTransactionCategoryType = "RentalCarTaxi"
	ASSETTRANSACTIONCATEGORYTYPE_RESTAURANTS AssetTransactionCategoryType = "Restaurants"
	ASSETTRANSACTIONCATEGORYTYPE_SALES_TAX AssetTransactionCategoryType = "SalesTax"
	ASSETTRANSACTIONCATEGORYTYPE_SERVICE_PARTS AssetTransactionCategoryType = "ServiceParts"
	ASSETTRANSACTIONCATEGORYTYPE_SERVICE_FEE AssetTransactionCategoryType = "ServiceFee"
	ASSETTRANSACTIONCATEGORYTYPE_SHIPPING AssetTransactionCategoryType = "Shipping"
	ASSETTRANSACTIONCATEGORYTYPE_SHOPPING AssetTransactionCategoryType = "Shopping"
	ASSETTRANSACTIONCATEGORYTYPE_SPA_MASSAGE AssetTransactionCategoryType = "SpaMassage"
	ASSETTRANSACTIONCATEGORYTYPE_SPORTING_GOODS AssetTransactionCategoryType = "SportingGoods"
	ASSETTRANSACTIONCATEGORYTYPE_SPORTS AssetTransactionCategoryType = "Sports"
	ASSETTRANSACTIONCATEGORYTYPE_STATE_TAX AssetTransactionCategoryType = "StateTax"
	ASSETTRANSACTIONCATEGORYTYPE_STUDENT_LOAN AssetTransactionCategoryType = "Student Loan"
	ASSETTRANSACTIONCATEGORYTYPE_TAXES AssetTransactionCategoryType = "Taxes"
	ASSETTRANSACTIONCATEGORYTYPE_TELEVISION AssetTransactionCategoryType = "Television"
	ASSETTRANSACTIONCATEGORYTYPE_TOYS AssetTransactionCategoryType = "Toys"
	ASSETTRANSACTIONCATEGORYTYPE_TRANSFER AssetTransactionCategoryType = "Transfer"
	ASSETTRANSACTIONCATEGORYTYPE_TRAVEL AssetTransactionCategoryType = "Travel"
	ASSETTRANSACTIONCATEGORYTYPE_TUITION AssetTransactionCategoryType = "Tuition"
	ASSETTRANSACTIONCATEGORYTYPE_UNCATEGORIZED AssetTransactionCategoryType = "Uncategorized"
	ASSETTRANSACTIONCATEGORYTYPE_UTILITIES AssetTransactionCategoryType = "Utilities"
	ASSETTRANSACTIONCATEGORYTYPE_VACATION AssetTransactionCategoryType = "Vacation"
	ASSETTRANSACTIONCATEGORYTYPE_VETERINARY AssetTransactionCategoryType = "Veterinary"
)

var allowedAssetTransactionCategoryTypeEnumValues = []AssetTransactionCategoryType{
	"ATMFee",
	"Advertising",
	"AirTravel",
	"AlcoholBars",
	"Allowance",
	"Amusement",
	"Arts",
	"AutoTransport",
	"AutoInsurance",
	"AutoPayment",
	"BabySupplies",
	"BabysitterDaycare",
	"BankFee",
	"BillsUtilities",
	"Bonus",
	"BooksSupplies",
	"Business Services",
	"Buy",
	"CashATM",
	"Charity",
	"Check",
	"ChildSupport",
	"Clothing",
	"CoffeeShops",
	"CreditCardPayment",
	"Dentist",
	"Doctor",
	"Education",
	"ElectronicsSoftware",
	"Entertainment",
	"Eyecare",
	"FastFood",
	"FederalTax",
	"FeesCharges",
	"FinanceCharge",
	"Financial",
	"FinancialAdvisor",
	"FoodDining",
	"Furnishings",
	"GasFuel",
	"GiftsDonations",
	"Groceries",
	"Gym",
	"Hair",
	"HealthFitness",
	"HealthInsurance",
	"Hobbies",
	"Home",
	"HomeImprovement",
	"HomeInsurance",
	"HomePhone",
	"HomeServices",
	"HomeSupplies",
	"Hotel",
	"Income",
	"InterestIncome",
	"Internet",
	"Investments",
	"Kids",
	"KidsActivities",
	"LateFee",
	"Laundry",
	"LawnGarden",
	"Legal",
	"LifeInsurance",
	"LoanInsurance",
	"LoanPayment",
	"Loans",
	"MobilePhone",
	"MortgageRent",
	"MoviesDVDs",
	"Music",
	"NewspapersMagazines",
	"OfficeSupplies",
	"Parking",
	"Paycheck",
	"PersonalCare",
	"PetFoodSupplies",
	"PetGrooming",
	"Pets",
	"Pharmacy",
	"Printing",
	"Property Tax",
	"Public Transportation",
	"Reimbursement",
	"RentalCarTaxi",
	"Restaurants",
	"SalesTax",
	"ServiceParts",
	"ServiceFee",
	"Shipping",
	"Shopping",
	"SpaMassage",
	"SportingGoods",
	"Sports",
	"StateTax",
	"Student Loan",
	"Taxes",
	"Television",
	"Toys",
	"Transfer",
	"Travel",
	"Tuition",
	"Uncategorized",
	"Utilities",
	"Vacation",
	"Veterinary",
}

func (v *AssetTransactionCategoryType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AssetTransactionCategoryType(value)
	for _, existing := range allowedAssetTransactionCategoryTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AssetTransactionCategoryType", value)
}

// NewAssetTransactionCategoryTypeFromValue returns a pointer to a valid AssetTransactionCategoryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAssetTransactionCategoryTypeFromValue(v string) (*AssetTransactionCategoryType, error) {
	ev := AssetTransactionCategoryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AssetTransactionCategoryType: valid values are %v", v, allowedAssetTransactionCategoryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AssetTransactionCategoryType) IsValid() bool {
	for _, existing := range allowedAssetTransactionCategoryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AssetTransactionCategoryType value
func (v AssetTransactionCategoryType) Ptr() *AssetTransactionCategoryType {
	return &v
}

type NullableAssetTransactionCategoryType struct {
	value *AssetTransactionCategoryType
	isSet bool
}

func (v NullableAssetTransactionCategoryType) Get() *AssetTransactionCategoryType {
	return v.value
}

func (v *NullableAssetTransactionCategoryType) Set(val *AssetTransactionCategoryType) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetTransactionCategoryType) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetTransactionCategoryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetTransactionCategoryType(val *AssetTransactionCategoryType) *NullableAssetTransactionCategoryType {
	return &NullableAssetTransactionCategoryType{value: val, isSet: true}
}

func (v NullableAssetTransactionCategoryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetTransactionCategoryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

