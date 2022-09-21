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
)

// W2 W2 is an object that represents income data taken from a W2 tax document.
type W2 struct {
	Employer *PaystubEmployer `json:"employer,omitempty"`
	Employee *Employee `json:"employee,omitempty"`
	// The tax year of the W2 document.
	TaxYear NullableString `json:"tax_year,omitempty"`
	// An employee identification number or EIN.
	EmployerIdNumber NullableString `json:"employer_id_number,omitempty"`
	// Wages from tips and other compensation.
	WagesTipsOtherComp NullableString `json:"wages_tips_other_comp,omitempty"`
	// Federal income tax withheld for the tax year.
	FederalIncomeTaxWithheld NullableString `json:"federal_income_tax_withheld,omitempty"`
	// Wages from social security.
	SocialSecurityWages NullableString `json:"social_security_wages,omitempty"`
	// Social security tax withheld for the tax year.
	SocialSecurityTaxWithheld NullableString `json:"social_security_tax_withheld,omitempty"`
	// Wages and tips from medicare.
	MedicareWagesAndTips NullableString `json:"medicare_wages_and_tips,omitempty"`
	// Medicare tax withheld for the tax year.
	MedicareTaxWithheld NullableString `json:"medicare_tax_withheld,omitempty"`
	// Tips from social security.
	SocialSecurityTips NullableString `json:"social_security_tips,omitempty"`
	// Allocated tips.
	AllocatedTips NullableString `json:"allocated_tips,omitempty"`
	// Contents from box 9 on the W2.
	Box9 NullableString `json:"box_9,omitempty"`
	// Dependent care benefits.
	DependentCareBenefits NullableString `json:"dependent_care_benefits,omitempty"`
	// Nonqualified plans.
	NonqualifiedPlans NullableString `json:"nonqualified_plans,omitempty"`
	Box12 *[]W2Box12 `json:"box_12,omitempty"`
	// Statutory employee.
	StatutoryEmployee NullableString `json:"statutory_employee,omitempty"`
	// Retirement plan.
	RetirementPlan NullableString `json:"retirement_plan,omitempty"`
	// Third party sick pay.
	ThirdPartySickPay NullableString `json:"third_party_sick_pay,omitempty"`
	// Other.
	Other NullableString `json:"other,omitempty"`
	StateAndLocalWages *[]W2StateAndLocalWages `json:"state_and_local_wages,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _W2 W2

// NewW2 instantiates a new W2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewW2() *W2 {
	this := W2{}
	return &this
}

// NewW2WithDefaults instantiates a new W2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewW2WithDefaults() *W2 {
	this := W2{}
	return &this
}

// GetEmployer returns the Employer field value if set, zero value otherwise.
func (o *W2) GetEmployer() PaystubEmployer {
	if o == nil || o.Employer == nil {
		var ret PaystubEmployer
		return ret
	}
	return *o.Employer
}

// GetEmployerOk returns a tuple with the Employer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *W2) GetEmployerOk() (*PaystubEmployer, bool) {
	if o == nil || o.Employer == nil {
		return nil, false
	}
	return o.Employer, true
}

// HasEmployer returns a boolean if a field has been set.
func (o *W2) HasEmployer() bool {
	if o != nil && o.Employer != nil {
		return true
	}

	return false
}

// SetEmployer gets a reference to the given PaystubEmployer and assigns it to the Employer field.
func (o *W2) SetEmployer(v PaystubEmployer) {
	o.Employer = &v
}

// GetEmployee returns the Employee field value if set, zero value otherwise.
func (o *W2) GetEmployee() Employee {
	if o == nil || o.Employee == nil {
		var ret Employee
		return ret
	}
	return *o.Employee
}

// GetEmployeeOk returns a tuple with the Employee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *W2) GetEmployeeOk() (*Employee, bool) {
	if o == nil || o.Employee == nil {
		return nil, false
	}
	return o.Employee, true
}

// HasEmployee returns a boolean if a field has been set.
func (o *W2) HasEmployee() bool {
	if o != nil && o.Employee != nil {
		return true
	}

	return false
}

// SetEmployee gets a reference to the given Employee and assigns it to the Employee field.
func (o *W2) SetEmployee(v Employee) {
	o.Employee = &v
}

// GetTaxYear returns the TaxYear field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2) GetTaxYear() string {
	if o == nil || o.TaxYear.Get() == nil {
		var ret string
		return ret
	}
	return *o.TaxYear.Get()
}

// GetTaxYearOk returns a tuple with the TaxYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2) GetTaxYearOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TaxYear.Get(), o.TaxYear.IsSet()
}

// HasTaxYear returns a boolean if a field has been set.
func (o *W2) HasTaxYear() bool {
	if o != nil && o.TaxYear.IsSet() {
		return true
	}

	return false
}

// SetTaxYear gets a reference to the given NullableString and assigns it to the TaxYear field.
func (o *W2) SetTaxYear(v string) {
	o.TaxYear.Set(&v)
}
// SetTaxYearNil sets the value for TaxYear to be an explicit nil
func (o *W2) SetTaxYearNil() {
	o.TaxYear.Set(nil)
}

// UnsetTaxYear ensures that no value is present for TaxYear, not even an explicit nil
func (o *W2) UnsetTaxYear() {
	o.TaxYear.Unset()
}

// GetEmployerIdNumber returns the EmployerIdNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2) GetEmployerIdNumber() string {
	if o == nil || o.EmployerIdNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.EmployerIdNumber.Get()
}

// GetEmployerIdNumberOk returns a tuple with the EmployerIdNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2) GetEmployerIdNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmployerIdNumber.Get(), o.EmployerIdNumber.IsSet()
}

// HasEmployerIdNumber returns a boolean if a field has been set.
func (o *W2) HasEmployerIdNumber() bool {
	if o != nil && o.EmployerIdNumber.IsSet() {
		return true
	}

	return false
}

// SetEmployerIdNumber gets a reference to the given NullableString and assigns it to the EmployerIdNumber field.
func (o *W2) SetEmployerIdNumber(v string) {
	o.EmployerIdNumber.Set(&v)
}
// SetEmployerIdNumberNil sets the value for EmployerIdNumber to be an explicit nil
func (o *W2) SetEmployerIdNumberNil() {
	o.EmployerIdNumber.Set(nil)
}

// UnsetEmployerIdNumber ensures that no value is present for EmployerIdNumber, not even an explicit nil
func (o *W2) UnsetEmployerIdNumber() {
	o.EmployerIdNumber.Unset()
}

// GetWagesTipsOtherComp returns the WagesTipsOtherComp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2) GetWagesTipsOtherComp() string {
	if o == nil || o.WagesTipsOtherComp.Get() == nil {
		var ret string
		return ret
	}
	return *o.WagesTipsOtherComp.Get()
}

// GetWagesTipsOtherCompOk returns a tuple with the WagesTipsOtherComp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2) GetWagesTipsOtherCompOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WagesTipsOtherComp.Get(), o.WagesTipsOtherComp.IsSet()
}

// HasWagesTipsOtherComp returns a boolean if a field has been set.
func (o *W2) HasWagesTipsOtherComp() bool {
	if o != nil && o.WagesTipsOtherComp.IsSet() {
		return true
	}

	return false
}

// SetWagesTipsOtherComp gets a reference to the given NullableString and assigns it to the WagesTipsOtherComp field.
func (o *W2) SetWagesTipsOtherComp(v string) {
	o.WagesTipsOtherComp.Set(&v)
}
// SetWagesTipsOtherCompNil sets the value for WagesTipsOtherComp to be an explicit nil
func (o *W2) SetWagesTipsOtherCompNil() {
	o.WagesTipsOtherComp.Set(nil)
}

// UnsetWagesTipsOtherComp ensures that no value is present for WagesTipsOtherComp, not even an explicit nil
func (o *W2) UnsetWagesTipsOtherComp() {
	o.WagesTipsOtherComp.Unset()
}

// GetFederalIncomeTaxWithheld returns the FederalIncomeTaxWithheld field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2) GetFederalIncomeTaxWithheld() string {
	if o == nil || o.FederalIncomeTaxWithheld.Get() == nil {
		var ret string
		return ret
	}
	return *o.FederalIncomeTaxWithheld.Get()
}

// GetFederalIncomeTaxWithheldOk returns a tuple with the FederalIncomeTaxWithheld field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2) GetFederalIncomeTaxWithheldOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FederalIncomeTaxWithheld.Get(), o.FederalIncomeTaxWithheld.IsSet()
}

// HasFederalIncomeTaxWithheld returns a boolean if a field has been set.
func (o *W2) HasFederalIncomeTaxWithheld() bool {
	if o != nil && o.FederalIncomeTaxWithheld.IsSet() {
		return true
	}

	return false
}

// SetFederalIncomeTaxWithheld gets a reference to the given NullableString and assigns it to the FederalIncomeTaxWithheld field.
func (o *W2) SetFederalIncomeTaxWithheld(v string) {
	o.FederalIncomeTaxWithheld.Set(&v)
}
// SetFederalIncomeTaxWithheldNil sets the value for FederalIncomeTaxWithheld to be an explicit nil
func (o *W2) SetFederalIncomeTaxWithheldNil() {
	o.FederalIncomeTaxWithheld.Set(nil)
}

// UnsetFederalIncomeTaxWithheld ensures that no value is present for FederalIncomeTaxWithheld, not even an explicit nil
func (o *W2) UnsetFederalIncomeTaxWithheld() {
	o.FederalIncomeTaxWithheld.Unset()
}

// GetSocialSecurityWages returns the SocialSecurityWages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2) GetSocialSecurityWages() string {
	if o == nil || o.SocialSecurityWages.Get() == nil {
		var ret string
		return ret
	}
	return *o.SocialSecurityWages.Get()
}

// GetSocialSecurityWagesOk returns a tuple with the SocialSecurityWages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2) GetSocialSecurityWagesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SocialSecurityWages.Get(), o.SocialSecurityWages.IsSet()
}

// HasSocialSecurityWages returns a boolean if a field has been set.
func (o *W2) HasSocialSecurityWages() bool {
	if o != nil && o.SocialSecurityWages.IsSet() {
		return true
	}

	return false
}

// SetSocialSecurityWages gets a reference to the given NullableString and assigns it to the SocialSecurityWages field.
func (o *W2) SetSocialSecurityWages(v string) {
	o.SocialSecurityWages.Set(&v)
}
// SetSocialSecurityWagesNil sets the value for SocialSecurityWages to be an explicit nil
func (o *W2) SetSocialSecurityWagesNil() {
	o.SocialSecurityWages.Set(nil)
}

// UnsetSocialSecurityWages ensures that no value is present for SocialSecurityWages, not even an explicit nil
func (o *W2) UnsetSocialSecurityWages() {
	o.SocialSecurityWages.Unset()
}

// GetSocialSecurityTaxWithheld returns the SocialSecurityTaxWithheld field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2) GetSocialSecurityTaxWithheld() string {
	if o == nil || o.SocialSecurityTaxWithheld.Get() == nil {
		var ret string
		return ret
	}
	return *o.SocialSecurityTaxWithheld.Get()
}

// GetSocialSecurityTaxWithheldOk returns a tuple with the SocialSecurityTaxWithheld field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2) GetSocialSecurityTaxWithheldOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SocialSecurityTaxWithheld.Get(), o.SocialSecurityTaxWithheld.IsSet()
}

// HasSocialSecurityTaxWithheld returns a boolean if a field has been set.
func (o *W2) HasSocialSecurityTaxWithheld() bool {
	if o != nil && o.SocialSecurityTaxWithheld.IsSet() {
		return true
	}

	return false
}

// SetSocialSecurityTaxWithheld gets a reference to the given NullableString and assigns it to the SocialSecurityTaxWithheld field.
func (o *W2) SetSocialSecurityTaxWithheld(v string) {
	o.SocialSecurityTaxWithheld.Set(&v)
}
// SetSocialSecurityTaxWithheldNil sets the value for SocialSecurityTaxWithheld to be an explicit nil
func (o *W2) SetSocialSecurityTaxWithheldNil() {
	o.SocialSecurityTaxWithheld.Set(nil)
}

// UnsetSocialSecurityTaxWithheld ensures that no value is present for SocialSecurityTaxWithheld, not even an explicit nil
func (o *W2) UnsetSocialSecurityTaxWithheld() {
	o.SocialSecurityTaxWithheld.Unset()
}

// GetMedicareWagesAndTips returns the MedicareWagesAndTips field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2) GetMedicareWagesAndTips() string {
	if o == nil || o.MedicareWagesAndTips.Get() == nil {
		var ret string
		return ret
	}
	return *o.MedicareWagesAndTips.Get()
}

// GetMedicareWagesAndTipsOk returns a tuple with the MedicareWagesAndTips field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2) GetMedicareWagesAndTipsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MedicareWagesAndTips.Get(), o.MedicareWagesAndTips.IsSet()
}

// HasMedicareWagesAndTips returns a boolean if a field has been set.
func (o *W2) HasMedicareWagesAndTips() bool {
	if o != nil && o.MedicareWagesAndTips.IsSet() {
		return true
	}

	return false
}

// SetMedicareWagesAndTips gets a reference to the given NullableString and assigns it to the MedicareWagesAndTips field.
func (o *W2) SetMedicareWagesAndTips(v string) {
	o.MedicareWagesAndTips.Set(&v)
}
// SetMedicareWagesAndTipsNil sets the value for MedicareWagesAndTips to be an explicit nil
func (o *W2) SetMedicareWagesAndTipsNil() {
	o.MedicareWagesAndTips.Set(nil)
}

// UnsetMedicareWagesAndTips ensures that no value is present for MedicareWagesAndTips, not even an explicit nil
func (o *W2) UnsetMedicareWagesAndTips() {
	o.MedicareWagesAndTips.Unset()
}

// GetMedicareTaxWithheld returns the MedicareTaxWithheld field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2) GetMedicareTaxWithheld() string {
	if o == nil || o.MedicareTaxWithheld.Get() == nil {
		var ret string
		return ret
	}
	return *o.MedicareTaxWithheld.Get()
}

// GetMedicareTaxWithheldOk returns a tuple with the MedicareTaxWithheld field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2) GetMedicareTaxWithheldOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MedicareTaxWithheld.Get(), o.MedicareTaxWithheld.IsSet()
}

// HasMedicareTaxWithheld returns a boolean if a field has been set.
func (o *W2) HasMedicareTaxWithheld() bool {
	if o != nil && o.MedicareTaxWithheld.IsSet() {
		return true
	}

	return false
}

// SetMedicareTaxWithheld gets a reference to the given NullableString and assigns it to the MedicareTaxWithheld field.
func (o *W2) SetMedicareTaxWithheld(v string) {
	o.MedicareTaxWithheld.Set(&v)
}
// SetMedicareTaxWithheldNil sets the value for MedicareTaxWithheld to be an explicit nil
func (o *W2) SetMedicareTaxWithheldNil() {
	o.MedicareTaxWithheld.Set(nil)
}

// UnsetMedicareTaxWithheld ensures that no value is present for MedicareTaxWithheld, not even an explicit nil
func (o *W2) UnsetMedicareTaxWithheld() {
	o.MedicareTaxWithheld.Unset()
}

// GetSocialSecurityTips returns the SocialSecurityTips field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2) GetSocialSecurityTips() string {
	if o == nil || o.SocialSecurityTips.Get() == nil {
		var ret string
		return ret
	}
	return *o.SocialSecurityTips.Get()
}

// GetSocialSecurityTipsOk returns a tuple with the SocialSecurityTips field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2) GetSocialSecurityTipsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SocialSecurityTips.Get(), o.SocialSecurityTips.IsSet()
}

// HasSocialSecurityTips returns a boolean if a field has been set.
func (o *W2) HasSocialSecurityTips() bool {
	if o != nil && o.SocialSecurityTips.IsSet() {
		return true
	}

	return false
}

// SetSocialSecurityTips gets a reference to the given NullableString and assigns it to the SocialSecurityTips field.
func (o *W2) SetSocialSecurityTips(v string) {
	o.SocialSecurityTips.Set(&v)
}
// SetSocialSecurityTipsNil sets the value for SocialSecurityTips to be an explicit nil
func (o *W2) SetSocialSecurityTipsNil() {
	o.SocialSecurityTips.Set(nil)
}

// UnsetSocialSecurityTips ensures that no value is present for SocialSecurityTips, not even an explicit nil
func (o *W2) UnsetSocialSecurityTips() {
	o.SocialSecurityTips.Unset()
}

// GetAllocatedTips returns the AllocatedTips field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2) GetAllocatedTips() string {
	if o == nil || o.AllocatedTips.Get() == nil {
		var ret string
		return ret
	}
	return *o.AllocatedTips.Get()
}

// GetAllocatedTipsOk returns a tuple with the AllocatedTips field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2) GetAllocatedTipsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AllocatedTips.Get(), o.AllocatedTips.IsSet()
}

// HasAllocatedTips returns a boolean if a field has been set.
func (o *W2) HasAllocatedTips() bool {
	if o != nil && o.AllocatedTips.IsSet() {
		return true
	}

	return false
}

// SetAllocatedTips gets a reference to the given NullableString and assigns it to the AllocatedTips field.
func (o *W2) SetAllocatedTips(v string) {
	o.AllocatedTips.Set(&v)
}
// SetAllocatedTipsNil sets the value for AllocatedTips to be an explicit nil
func (o *W2) SetAllocatedTipsNil() {
	o.AllocatedTips.Set(nil)
}

// UnsetAllocatedTips ensures that no value is present for AllocatedTips, not even an explicit nil
func (o *W2) UnsetAllocatedTips() {
	o.AllocatedTips.Unset()
}

// GetBox9 returns the Box9 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2) GetBox9() string {
	if o == nil || o.Box9.Get() == nil {
		var ret string
		return ret
	}
	return *o.Box9.Get()
}

// GetBox9Ok returns a tuple with the Box9 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2) GetBox9Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Box9.Get(), o.Box9.IsSet()
}

// HasBox9 returns a boolean if a field has been set.
func (o *W2) HasBox9() bool {
	if o != nil && o.Box9.IsSet() {
		return true
	}

	return false
}

// SetBox9 gets a reference to the given NullableString and assigns it to the Box9 field.
func (o *W2) SetBox9(v string) {
	o.Box9.Set(&v)
}
// SetBox9Nil sets the value for Box9 to be an explicit nil
func (o *W2) SetBox9Nil() {
	o.Box9.Set(nil)
}

// UnsetBox9 ensures that no value is present for Box9, not even an explicit nil
func (o *W2) UnsetBox9() {
	o.Box9.Unset()
}

// GetDependentCareBenefits returns the DependentCareBenefits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2) GetDependentCareBenefits() string {
	if o == nil || o.DependentCareBenefits.Get() == nil {
		var ret string
		return ret
	}
	return *o.DependentCareBenefits.Get()
}

// GetDependentCareBenefitsOk returns a tuple with the DependentCareBenefits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2) GetDependentCareBenefitsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DependentCareBenefits.Get(), o.DependentCareBenefits.IsSet()
}

// HasDependentCareBenefits returns a boolean if a field has been set.
func (o *W2) HasDependentCareBenefits() bool {
	if o != nil && o.DependentCareBenefits.IsSet() {
		return true
	}

	return false
}

// SetDependentCareBenefits gets a reference to the given NullableString and assigns it to the DependentCareBenefits field.
func (o *W2) SetDependentCareBenefits(v string) {
	o.DependentCareBenefits.Set(&v)
}
// SetDependentCareBenefitsNil sets the value for DependentCareBenefits to be an explicit nil
func (o *W2) SetDependentCareBenefitsNil() {
	o.DependentCareBenefits.Set(nil)
}

// UnsetDependentCareBenefits ensures that no value is present for DependentCareBenefits, not even an explicit nil
func (o *W2) UnsetDependentCareBenefits() {
	o.DependentCareBenefits.Unset()
}

// GetNonqualifiedPlans returns the NonqualifiedPlans field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2) GetNonqualifiedPlans() string {
	if o == nil || o.NonqualifiedPlans.Get() == nil {
		var ret string
		return ret
	}
	return *o.NonqualifiedPlans.Get()
}

// GetNonqualifiedPlansOk returns a tuple with the NonqualifiedPlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2) GetNonqualifiedPlansOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NonqualifiedPlans.Get(), o.NonqualifiedPlans.IsSet()
}

// HasNonqualifiedPlans returns a boolean if a field has been set.
func (o *W2) HasNonqualifiedPlans() bool {
	if o != nil && o.NonqualifiedPlans.IsSet() {
		return true
	}

	return false
}

// SetNonqualifiedPlans gets a reference to the given NullableString and assigns it to the NonqualifiedPlans field.
func (o *W2) SetNonqualifiedPlans(v string) {
	o.NonqualifiedPlans.Set(&v)
}
// SetNonqualifiedPlansNil sets the value for NonqualifiedPlans to be an explicit nil
func (o *W2) SetNonqualifiedPlansNil() {
	o.NonqualifiedPlans.Set(nil)
}

// UnsetNonqualifiedPlans ensures that no value is present for NonqualifiedPlans, not even an explicit nil
func (o *W2) UnsetNonqualifiedPlans() {
	o.NonqualifiedPlans.Unset()
}

// GetBox12 returns the Box12 field value if set, zero value otherwise.
func (o *W2) GetBox12() []W2Box12 {
	if o == nil || o.Box12 == nil {
		var ret []W2Box12
		return ret
	}
	return *o.Box12
}

// GetBox12Ok returns a tuple with the Box12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *W2) GetBox12Ok() (*[]W2Box12, bool) {
	if o == nil || o.Box12 == nil {
		return nil, false
	}
	return o.Box12, true
}

// HasBox12 returns a boolean if a field has been set.
func (o *W2) HasBox12() bool {
	if o != nil && o.Box12 != nil {
		return true
	}

	return false
}

// SetBox12 gets a reference to the given []W2Box12 and assigns it to the Box12 field.
func (o *W2) SetBox12(v []W2Box12) {
	o.Box12 = &v
}

// GetStatutoryEmployee returns the StatutoryEmployee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2) GetStatutoryEmployee() string {
	if o == nil || o.StatutoryEmployee.Get() == nil {
		var ret string
		return ret
	}
	return *o.StatutoryEmployee.Get()
}

// GetStatutoryEmployeeOk returns a tuple with the StatutoryEmployee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2) GetStatutoryEmployeeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StatutoryEmployee.Get(), o.StatutoryEmployee.IsSet()
}

// HasStatutoryEmployee returns a boolean if a field has been set.
func (o *W2) HasStatutoryEmployee() bool {
	if o != nil && o.StatutoryEmployee.IsSet() {
		return true
	}

	return false
}

// SetStatutoryEmployee gets a reference to the given NullableString and assigns it to the StatutoryEmployee field.
func (o *W2) SetStatutoryEmployee(v string) {
	o.StatutoryEmployee.Set(&v)
}
// SetStatutoryEmployeeNil sets the value for StatutoryEmployee to be an explicit nil
func (o *W2) SetStatutoryEmployeeNil() {
	o.StatutoryEmployee.Set(nil)
}

// UnsetStatutoryEmployee ensures that no value is present for StatutoryEmployee, not even an explicit nil
func (o *W2) UnsetStatutoryEmployee() {
	o.StatutoryEmployee.Unset()
}

// GetRetirementPlan returns the RetirementPlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2) GetRetirementPlan() string {
	if o == nil || o.RetirementPlan.Get() == nil {
		var ret string
		return ret
	}
	return *o.RetirementPlan.Get()
}

// GetRetirementPlanOk returns a tuple with the RetirementPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2) GetRetirementPlanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RetirementPlan.Get(), o.RetirementPlan.IsSet()
}

// HasRetirementPlan returns a boolean if a field has been set.
func (o *W2) HasRetirementPlan() bool {
	if o != nil && o.RetirementPlan.IsSet() {
		return true
	}

	return false
}

// SetRetirementPlan gets a reference to the given NullableString and assigns it to the RetirementPlan field.
func (o *W2) SetRetirementPlan(v string) {
	o.RetirementPlan.Set(&v)
}
// SetRetirementPlanNil sets the value for RetirementPlan to be an explicit nil
func (o *W2) SetRetirementPlanNil() {
	o.RetirementPlan.Set(nil)
}

// UnsetRetirementPlan ensures that no value is present for RetirementPlan, not even an explicit nil
func (o *W2) UnsetRetirementPlan() {
	o.RetirementPlan.Unset()
}

// GetThirdPartySickPay returns the ThirdPartySickPay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2) GetThirdPartySickPay() string {
	if o == nil || o.ThirdPartySickPay.Get() == nil {
		var ret string
		return ret
	}
	return *o.ThirdPartySickPay.Get()
}

// GetThirdPartySickPayOk returns a tuple with the ThirdPartySickPay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2) GetThirdPartySickPayOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ThirdPartySickPay.Get(), o.ThirdPartySickPay.IsSet()
}

// HasThirdPartySickPay returns a boolean if a field has been set.
func (o *W2) HasThirdPartySickPay() bool {
	if o != nil && o.ThirdPartySickPay.IsSet() {
		return true
	}

	return false
}

// SetThirdPartySickPay gets a reference to the given NullableString and assigns it to the ThirdPartySickPay field.
func (o *W2) SetThirdPartySickPay(v string) {
	o.ThirdPartySickPay.Set(&v)
}
// SetThirdPartySickPayNil sets the value for ThirdPartySickPay to be an explicit nil
func (o *W2) SetThirdPartySickPayNil() {
	o.ThirdPartySickPay.Set(nil)
}

// UnsetThirdPartySickPay ensures that no value is present for ThirdPartySickPay, not even an explicit nil
func (o *W2) UnsetThirdPartySickPay() {
	o.ThirdPartySickPay.Unset()
}

// GetOther returns the Other field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *W2) GetOther() string {
	if o == nil || o.Other.Get() == nil {
		var ret string
		return ret
	}
	return *o.Other.Get()
}

// GetOtherOk returns a tuple with the Other field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *W2) GetOtherOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Other.Get(), o.Other.IsSet()
}

// HasOther returns a boolean if a field has been set.
func (o *W2) HasOther() bool {
	if o != nil && o.Other.IsSet() {
		return true
	}

	return false
}

// SetOther gets a reference to the given NullableString and assigns it to the Other field.
func (o *W2) SetOther(v string) {
	o.Other.Set(&v)
}
// SetOtherNil sets the value for Other to be an explicit nil
func (o *W2) SetOtherNil() {
	o.Other.Set(nil)
}

// UnsetOther ensures that no value is present for Other, not even an explicit nil
func (o *W2) UnsetOther() {
	o.Other.Unset()
}

// GetStateAndLocalWages returns the StateAndLocalWages field value if set, zero value otherwise.
func (o *W2) GetStateAndLocalWages() []W2StateAndLocalWages {
	if o == nil || o.StateAndLocalWages == nil {
		var ret []W2StateAndLocalWages
		return ret
	}
	return *o.StateAndLocalWages
}

// GetStateAndLocalWagesOk returns a tuple with the StateAndLocalWages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *W2) GetStateAndLocalWagesOk() (*[]W2StateAndLocalWages, bool) {
	if o == nil || o.StateAndLocalWages == nil {
		return nil, false
	}
	return o.StateAndLocalWages, true
}

// HasStateAndLocalWages returns a boolean if a field has been set.
func (o *W2) HasStateAndLocalWages() bool {
	if o != nil && o.StateAndLocalWages != nil {
		return true
	}

	return false
}

// SetStateAndLocalWages gets a reference to the given []W2StateAndLocalWages and assigns it to the StateAndLocalWages field.
func (o *W2) SetStateAndLocalWages(v []W2StateAndLocalWages) {
	o.StateAndLocalWages = &v
}

func (o W2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Employer != nil {
		toSerialize["employer"] = o.Employer
	}
	if o.Employee != nil {
		toSerialize["employee"] = o.Employee
	}
	if o.TaxYear.IsSet() {
		toSerialize["tax_year"] = o.TaxYear.Get()
	}
	if o.EmployerIdNumber.IsSet() {
		toSerialize["employer_id_number"] = o.EmployerIdNumber.Get()
	}
	if o.WagesTipsOtherComp.IsSet() {
		toSerialize["wages_tips_other_comp"] = o.WagesTipsOtherComp.Get()
	}
	if o.FederalIncomeTaxWithheld.IsSet() {
		toSerialize["federal_income_tax_withheld"] = o.FederalIncomeTaxWithheld.Get()
	}
	if o.SocialSecurityWages.IsSet() {
		toSerialize["social_security_wages"] = o.SocialSecurityWages.Get()
	}
	if o.SocialSecurityTaxWithheld.IsSet() {
		toSerialize["social_security_tax_withheld"] = o.SocialSecurityTaxWithheld.Get()
	}
	if o.MedicareWagesAndTips.IsSet() {
		toSerialize["medicare_wages_and_tips"] = o.MedicareWagesAndTips.Get()
	}
	if o.MedicareTaxWithheld.IsSet() {
		toSerialize["medicare_tax_withheld"] = o.MedicareTaxWithheld.Get()
	}
	if o.SocialSecurityTips.IsSet() {
		toSerialize["social_security_tips"] = o.SocialSecurityTips.Get()
	}
	if o.AllocatedTips.IsSet() {
		toSerialize["allocated_tips"] = o.AllocatedTips.Get()
	}
	if o.Box9.IsSet() {
		toSerialize["box_9"] = o.Box9.Get()
	}
	if o.DependentCareBenefits.IsSet() {
		toSerialize["dependent_care_benefits"] = o.DependentCareBenefits.Get()
	}
	if o.NonqualifiedPlans.IsSet() {
		toSerialize["nonqualified_plans"] = o.NonqualifiedPlans.Get()
	}
	if o.Box12 != nil {
		toSerialize["box_12"] = o.Box12
	}
	if o.StatutoryEmployee.IsSet() {
		toSerialize["statutory_employee"] = o.StatutoryEmployee.Get()
	}
	if o.RetirementPlan.IsSet() {
		toSerialize["retirement_plan"] = o.RetirementPlan.Get()
	}
	if o.ThirdPartySickPay.IsSet() {
		toSerialize["third_party_sick_pay"] = o.ThirdPartySickPay.Get()
	}
	if o.Other.IsSet() {
		toSerialize["other"] = o.Other.Get()
	}
	if o.StateAndLocalWages != nil {
		toSerialize["state_and_local_wages"] = o.StateAndLocalWages
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *W2) UnmarshalJSON(bytes []byte) (err error) {
	varW2 := _W2{}

	if err = json.Unmarshal(bytes, &varW2); err == nil {
		*o = W2(varW2)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "employer")
		delete(additionalProperties, "employee")
		delete(additionalProperties, "tax_year")
		delete(additionalProperties, "employer_id_number")
		delete(additionalProperties, "wages_tips_other_comp")
		delete(additionalProperties, "federal_income_tax_withheld")
		delete(additionalProperties, "social_security_wages")
		delete(additionalProperties, "social_security_tax_withheld")
		delete(additionalProperties, "medicare_wages_and_tips")
		delete(additionalProperties, "medicare_tax_withheld")
		delete(additionalProperties, "social_security_tips")
		delete(additionalProperties, "allocated_tips")
		delete(additionalProperties, "box_9")
		delete(additionalProperties, "dependent_care_benefits")
		delete(additionalProperties, "nonqualified_plans")
		delete(additionalProperties, "box_12")
		delete(additionalProperties, "statutory_employee")
		delete(additionalProperties, "retirement_plan")
		delete(additionalProperties, "third_party_sick_pay")
		delete(additionalProperties, "other")
		delete(additionalProperties, "state_and_local_wages")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableW2 struct {
	value *W2
	isSet bool
}

func (v NullableW2) Get() *W2 {
	return v.value
}

func (v *NullableW2) Set(val *W2) {
	v.value = val
	v.isSet = true
}

func (v NullableW2) IsSet() bool {
	return v.isSet
}

func (v *NullableW2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableW2(val *W2) *NullableW2 {
	return &NullableW2{value: val, isSet: true}
}

func (v NullableW2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableW2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


