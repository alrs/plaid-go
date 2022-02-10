/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.64.14
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// LiabilityOverride Used to configure Sandbox test data for the Liabilities product
type LiabilityOverride struct {
	// The type of the liability object, either `credit` or `student`. Mortgages are not currently supported in the custom Sandbox.
	Type string `json:"type"`
	// The purchase APR percentage value. For simplicity, this is the only interest rate used to calculate interest charges. Can only be set if `type` is `credit`.
	PurchaseApr float32 `json:"purchase_apr"`
	// The cash APR percentage value. Can only be set if `type` is `credit`.
	CashApr float32 `json:"cash_apr"`
	// The balance transfer APR percentage value. Can only be set if `type` is `credit`. Can only be set if `type` is `credit`.
	BalanceTransferApr float32 `json:"balance_transfer_apr"`
	// The special APR percentage value. Can only be set if `type` is `credit`.
	SpecialApr float32 `json:"special_apr"`
	// Override the `last_payment_amount` field. Can only be set if `type` is `credit`.
	LastPaymentAmount float32 `json:"last_payment_amount"`
	// Override the `minimum_payment_amount` field. Can only be set if `type` is `credit` or `student`.
	MinimumPaymentAmount float32 `json:"minimum_payment_amount"`
	// Override the `is_overdue` field
	IsOverdue bool `json:"is_overdue"`
	// The date on which the loan was initially lent, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format. Can only be set if `type` is `student`.
	OriginationDate string `json:"origination_date"`
	// The original loan principal. Can only be set if `type` is `student`.
	Principal float32 `json:"principal"`
	// The interest rate on the loan as a percentage. Can only be set if `type` is `student`.
	NominalApr float32 `json:"nominal_apr"`
	// If set, interest capitalization begins at the given number of months after loan origination. By default interest is never capitalized. Can only be set if `type` is `student`.
	InterestCapitalizationGracePeriodMonths float32 `json:"interest_capitalization_grace_period_months"`
	RepaymentModel StudentLoanRepaymentModel `json:"repayment_model"`
	// Override the `expected_payoff_date` field. Can only be set if `type` is `student`.
	ExpectedPayoffDate string `json:"expected_payoff_date"`
	// Override the `guarantor` field. Can only be set if `type` is `student`.
	Guarantor string `json:"guarantor"`
	// Override the `is_federal` field. Can only be set if `type` is `student`.
	IsFederal bool `json:"is_federal"`
	// Override the `loan_name` field. Can only be set if `type` is `student`.
	LoanName string `json:"loan_name"`
	LoanStatus StudentLoanStatus `json:"loan_status"`
	// Override the `payment_reference_number` field. Can only be set if `type` is `student`.
	PaymentReferenceNumber string `json:"payment_reference_number"`
	PslfStatus PSLFStatus `json:"pslf_status"`
	// Override the `repayment_plan.description` field. Can only be set if `type` is `student`.
	RepaymentPlanDescription string `json:"repayment_plan_description"`
	// Override the `repayment_plan.type` field. Can only be set if `type` is `student`. Possible values are: `\"extended graduated\"`, `\"extended standard\"`, `\"graduated\"`, `\"income-contingent repayment\"`, `\"income-based repayment\"`, `\"interest only\"`, `\"other\"`, `\"pay as you earn\"`, `\"revised pay as you earn\"`, or `\"standard\"`.
	RepaymentPlanType string `json:"repayment_plan_type"`
	// Override the `sequence_number` field. Can only be set if `type` is `student`.
	SequenceNumber string `json:"sequence_number"`
	ServicerAddress Address `json:"servicer_address"`
	AdditionalProperties map[string]interface{}
}

type _LiabilityOverride LiabilityOverride

// NewLiabilityOverride instantiates a new LiabilityOverride object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiabilityOverride(type_ string, purchaseApr float32, cashApr float32, balanceTransferApr float32, specialApr float32, lastPaymentAmount float32, minimumPaymentAmount float32, isOverdue bool, originationDate string, principal float32, nominalApr float32, interestCapitalizationGracePeriodMonths float32, repaymentModel StudentLoanRepaymentModel, expectedPayoffDate string, guarantor string, isFederal bool, loanName string, loanStatus StudentLoanStatus, paymentReferenceNumber string, pslfStatus PSLFStatus, repaymentPlanDescription string, repaymentPlanType string, sequenceNumber string, servicerAddress Address) *LiabilityOverride {
	this := LiabilityOverride{}
	this.Type = type_
	this.PurchaseApr = purchaseApr
	this.CashApr = cashApr
	this.BalanceTransferApr = balanceTransferApr
	this.SpecialApr = specialApr
	this.LastPaymentAmount = lastPaymentAmount
	this.MinimumPaymentAmount = minimumPaymentAmount
	this.IsOverdue = isOverdue
	this.OriginationDate = originationDate
	this.Principal = principal
	this.NominalApr = nominalApr
	this.InterestCapitalizationGracePeriodMonths = interestCapitalizationGracePeriodMonths
	this.RepaymentModel = repaymentModel
	this.ExpectedPayoffDate = expectedPayoffDate
	this.Guarantor = guarantor
	this.IsFederal = isFederal
	this.LoanName = loanName
	this.LoanStatus = loanStatus
	this.PaymentReferenceNumber = paymentReferenceNumber
	this.PslfStatus = pslfStatus
	this.RepaymentPlanDescription = repaymentPlanDescription
	this.RepaymentPlanType = repaymentPlanType
	this.SequenceNumber = sequenceNumber
	this.ServicerAddress = servicerAddress
	return &this
}

// NewLiabilityOverrideWithDefaults instantiates a new LiabilityOverride object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiabilityOverrideWithDefaults() *LiabilityOverride {
	this := LiabilityOverride{}
	return &this
}

// GetType returns the Type field value
func (o *LiabilityOverride) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LiabilityOverride) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LiabilityOverride) SetType(v string) {
	o.Type = v
}

// GetPurchaseApr returns the PurchaseApr field value
func (o *LiabilityOverride) GetPurchaseApr() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PurchaseApr
}

// GetPurchaseAprOk returns a tuple with the PurchaseApr field value
// and a boolean to check if the value has been set.
func (o *LiabilityOverride) GetPurchaseAprOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PurchaseApr, true
}

// SetPurchaseApr sets field value
func (o *LiabilityOverride) SetPurchaseApr(v float32) {
	o.PurchaseApr = v
}

// GetCashApr returns the CashApr field value
func (o *LiabilityOverride) GetCashApr() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CashApr
}

// GetCashAprOk returns a tuple with the CashApr field value
// and a boolean to check if the value has been set.
func (o *LiabilityOverride) GetCashAprOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CashApr, true
}

// SetCashApr sets field value
func (o *LiabilityOverride) SetCashApr(v float32) {
	o.CashApr = v
}

// GetBalanceTransferApr returns the BalanceTransferApr field value
func (o *LiabilityOverride) GetBalanceTransferApr() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BalanceTransferApr
}

// GetBalanceTransferAprOk returns a tuple with the BalanceTransferApr field value
// and a boolean to check if the value has been set.
func (o *LiabilityOverride) GetBalanceTransferAprOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BalanceTransferApr, true
}

// SetBalanceTransferApr sets field value
func (o *LiabilityOverride) SetBalanceTransferApr(v float32) {
	o.BalanceTransferApr = v
}

// GetSpecialApr returns the SpecialApr field value
func (o *LiabilityOverride) GetSpecialApr() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SpecialApr
}

// GetSpecialAprOk returns a tuple with the SpecialApr field value
// and a boolean to check if the value has been set.
func (o *LiabilityOverride) GetSpecialAprOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SpecialApr, true
}

// SetSpecialApr sets field value
func (o *LiabilityOverride) SetSpecialApr(v float32) {
	o.SpecialApr = v
}

// GetLastPaymentAmount returns the LastPaymentAmount field value
func (o *LiabilityOverride) GetLastPaymentAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LastPaymentAmount
}

// GetLastPaymentAmountOk returns a tuple with the LastPaymentAmount field value
// and a boolean to check if the value has been set.
func (o *LiabilityOverride) GetLastPaymentAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastPaymentAmount, true
}

// SetLastPaymentAmount sets field value
func (o *LiabilityOverride) SetLastPaymentAmount(v float32) {
	o.LastPaymentAmount = v
}

// GetMinimumPaymentAmount returns the MinimumPaymentAmount field value
func (o *LiabilityOverride) GetMinimumPaymentAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MinimumPaymentAmount
}

// GetMinimumPaymentAmountOk returns a tuple with the MinimumPaymentAmount field value
// and a boolean to check if the value has been set.
func (o *LiabilityOverride) GetMinimumPaymentAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MinimumPaymentAmount, true
}

// SetMinimumPaymentAmount sets field value
func (o *LiabilityOverride) SetMinimumPaymentAmount(v float32) {
	o.MinimumPaymentAmount = v
}

// GetIsOverdue returns the IsOverdue field value
func (o *LiabilityOverride) GetIsOverdue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsOverdue
}

// GetIsOverdueOk returns a tuple with the IsOverdue field value
// and a boolean to check if the value has been set.
func (o *LiabilityOverride) GetIsOverdueOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsOverdue, true
}

// SetIsOverdue sets field value
func (o *LiabilityOverride) SetIsOverdue(v bool) {
	o.IsOverdue = v
}

// GetOriginationDate returns the OriginationDate field value
func (o *LiabilityOverride) GetOriginationDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginationDate
}

// GetOriginationDateOk returns a tuple with the OriginationDate field value
// and a boolean to check if the value has been set.
func (o *LiabilityOverride) GetOriginationDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OriginationDate, true
}

// SetOriginationDate sets field value
func (o *LiabilityOverride) SetOriginationDate(v string) {
	o.OriginationDate = v
}

// GetPrincipal returns the Principal field value
func (o *LiabilityOverride) GetPrincipal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value
// and a boolean to check if the value has been set.
func (o *LiabilityOverride) GetPrincipalOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Principal, true
}

// SetPrincipal sets field value
func (o *LiabilityOverride) SetPrincipal(v float32) {
	o.Principal = v
}

// GetNominalApr returns the NominalApr field value
func (o *LiabilityOverride) GetNominalApr() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NominalApr
}

// GetNominalAprOk returns a tuple with the NominalApr field value
// and a boolean to check if the value has been set.
func (o *LiabilityOverride) GetNominalAprOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NominalApr, true
}

// SetNominalApr sets field value
func (o *LiabilityOverride) SetNominalApr(v float32) {
	o.NominalApr = v
}

// GetInterestCapitalizationGracePeriodMonths returns the InterestCapitalizationGracePeriodMonths field value
func (o *LiabilityOverride) GetInterestCapitalizationGracePeriodMonths() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.InterestCapitalizationGracePeriodMonths
}

// GetInterestCapitalizationGracePeriodMonthsOk returns a tuple with the InterestCapitalizationGracePeriodMonths field value
// and a boolean to check if the value has been set.
func (o *LiabilityOverride) GetInterestCapitalizationGracePeriodMonthsOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InterestCapitalizationGracePeriodMonths, true
}

// SetInterestCapitalizationGracePeriodMonths sets field value
func (o *LiabilityOverride) SetInterestCapitalizationGracePeriodMonths(v float32) {
	o.InterestCapitalizationGracePeriodMonths = v
}

// GetRepaymentModel returns the RepaymentModel field value
func (o *LiabilityOverride) GetRepaymentModel() StudentLoanRepaymentModel {
	if o == nil {
		var ret StudentLoanRepaymentModel
		return ret
	}

	return o.RepaymentModel
}

// GetRepaymentModelOk returns a tuple with the RepaymentModel field value
// and a boolean to check if the value has been set.
func (o *LiabilityOverride) GetRepaymentModelOk() (*StudentLoanRepaymentModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RepaymentModel, true
}

// SetRepaymentModel sets field value
func (o *LiabilityOverride) SetRepaymentModel(v StudentLoanRepaymentModel) {
	o.RepaymentModel = v
}

// GetExpectedPayoffDate returns the ExpectedPayoffDate field value
func (o *LiabilityOverride) GetExpectedPayoffDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpectedPayoffDate
}

// GetExpectedPayoffDateOk returns a tuple with the ExpectedPayoffDate field value
// and a boolean to check if the value has been set.
func (o *LiabilityOverride) GetExpectedPayoffDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExpectedPayoffDate, true
}

// SetExpectedPayoffDate sets field value
func (o *LiabilityOverride) SetExpectedPayoffDate(v string) {
	o.ExpectedPayoffDate = v
}

// GetGuarantor returns the Guarantor field value
func (o *LiabilityOverride) GetGuarantor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Guarantor
}

// GetGuarantorOk returns a tuple with the Guarantor field value
// and a boolean to check if the value has been set.
func (o *LiabilityOverride) GetGuarantorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Guarantor, true
}

// SetGuarantor sets field value
func (o *LiabilityOverride) SetGuarantor(v string) {
	o.Guarantor = v
}

// GetIsFederal returns the IsFederal field value
func (o *LiabilityOverride) GetIsFederal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFederal
}

// GetIsFederalOk returns a tuple with the IsFederal field value
// and a boolean to check if the value has been set.
func (o *LiabilityOverride) GetIsFederalOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsFederal, true
}

// SetIsFederal sets field value
func (o *LiabilityOverride) SetIsFederal(v bool) {
	o.IsFederal = v
}

// GetLoanName returns the LoanName field value
func (o *LiabilityOverride) GetLoanName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoanName
}

// GetLoanNameOk returns a tuple with the LoanName field value
// and a boolean to check if the value has been set.
func (o *LiabilityOverride) GetLoanNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LoanName, true
}

// SetLoanName sets field value
func (o *LiabilityOverride) SetLoanName(v string) {
	o.LoanName = v
}

// GetLoanStatus returns the LoanStatus field value
func (o *LiabilityOverride) GetLoanStatus() StudentLoanStatus {
	if o == nil {
		var ret StudentLoanStatus
		return ret
	}

	return o.LoanStatus
}

// GetLoanStatusOk returns a tuple with the LoanStatus field value
// and a boolean to check if the value has been set.
func (o *LiabilityOverride) GetLoanStatusOk() (*StudentLoanStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LoanStatus, true
}

// SetLoanStatus sets field value
func (o *LiabilityOverride) SetLoanStatus(v StudentLoanStatus) {
	o.LoanStatus = v
}

// GetPaymentReferenceNumber returns the PaymentReferenceNumber field value
func (o *LiabilityOverride) GetPaymentReferenceNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentReferenceNumber
}

// GetPaymentReferenceNumberOk returns a tuple with the PaymentReferenceNumber field value
// and a boolean to check if the value has been set.
func (o *LiabilityOverride) GetPaymentReferenceNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentReferenceNumber, true
}

// SetPaymentReferenceNumber sets field value
func (o *LiabilityOverride) SetPaymentReferenceNumber(v string) {
	o.PaymentReferenceNumber = v
}

// GetPslfStatus returns the PslfStatus field value
func (o *LiabilityOverride) GetPslfStatus() PSLFStatus {
	if o == nil {
		var ret PSLFStatus
		return ret
	}

	return o.PslfStatus
}

// GetPslfStatusOk returns a tuple with the PslfStatus field value
// and a boolean to check if the value has been set.
func (o *LiabilityOverride) GetPslfStatusOk() (*PSLFStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PslfStatus, true
}

// SetPslfStatus sets field value
func (o *LiabilityOverride) SetPslfStatus(v PSLFStatus) {
	o.PslfStatus = v
}

// GetRepaymentPlanDescription returns the RepaymentPlanDescription field value
func (o *LiabilityOverride) GetRepaymentPlanDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RepaymentPlanDescription
}

// GetRepaymentPlanDescriptionOk returns a tuple with the RepaymentPlanDescription field value
// and a boolean to check if the value has been set.
func (o *LiabilityOverride) GetRepaymentPlanDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RepaymentPlanDescription, true
}

// SetRepaymentPlanDescription sets field value
func (o *LiabilityOverride) SetRepaymentPlanDescription(v string) {
	o.RepaymentPlanDescription = v
}

// GetRepaymentPlanType returns the RepaymentPlanType field value
func (o *LiabilityOverride) GetRepaymentPlanType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RepaymentPlanType
}

// GetRepaymentPlanTypeOk returns a tuple with the RepaymentPlanType field value
// and a boolean to check if the value has been set.
func (o *LiabilityOverride) GetRepaymentPlanTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RepaymentPlanType, true
}

// SetRepaymentPlanType sets field value
func (o *LiabilityOverride) SetRepaymentPlanType(v string) {
	o.RepaymentPlanType = v
}

// GetSequenceNumber returns the SequenceNumber field value
func (o *LiabilityOverride) GetSequenceNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value
// and a boolean to check if the value has been set.
func (o *LiabilityOverride) GetSequenceNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SequenceNumber, true
}

// SetSequenceNumber sets field value
func (o *LiabilityOverride) SetSequenceNumber(v string) {
	o.SequenceNumber = v
}

// GetServicerAddress returns the ServicerAddress field value
func (o *LiabilityOverride) GetServicerAddress() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.ServicerAddress
}

// GetServicerAddressOk returns a tuple with the ServicerAddress field value
// and a boolean to check if the value has been set.
func (o *LiabilityOverride) GetServicerAddressOk() (*Address, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServicerAddress, true
}

// SetServicerAddress sets field value
func (o *LiabilityOverride) SetServicerAddress(v Address) {
	o.ServicerAddress = v
}

func (o LiabilityOverride) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["purchase_apr"] = o.PurchaseApr
	}
	if true {
		toSerialize["cash_apr"] = o.CashApr
	}
	if true {
		toSerialize["balance_transfer_apr"] = o.BalanceTransferApr
	}
	if true {
		toSerialize["special_apr"] = o.SpecialApr
	}
	if true {
		toSerialize["last_payment_amount"] = o.LastPaymentAmount
	}
	if true {
		toSerialize["minimum_payment_amount"] = o.MinimumPaymentAmount
	}
	if true {
		toSerialize["is_overdue"] = o.IsOverdue
	}
	if true {
		toSerialize["origination_date"] = o.OriginationDate
	}
	if true {
		toSerialize["principal"] = o.Principal
	}
	if true {
		toSerialize["nominal_apr"] = o.NominalApr
	}
	if true {
		toSerialize["interest_capitalization_grace_period_months"] = o.InterestCapitalizationGracePeriodMonths
	}
	if true {
		toSerialize["repayment_model"] = o.RepaymentModel
	}
	if true {
		toSerialize["expected_payoff_date"] = o.ExpectedPayoffDate
	}
	if true {
		toSerialize["guarantor"] = o.Guarantor
	}
	if true {
		toSerialize["is_federal"] = o.IsFederal
	}
	if true {
		toSerialize["loan_name"] = o.LoanName
	}
	if true {
		toSerialize["loan_status"] = o.LoanStatus
	}
	if true {
		toSerialize["payment_reference_number"] = o.PaymentReferenceNumber
	}
	if true {
		toSerialize["pslf_status"] = o.PslfStatus
	}
	if true {
		toSerialize["repayment_plan_description"] = o.RepaymentPlanDescription
	}
	if true {
		toSerialize["repayment_plan_type"] = o.RepaymentPlanType
	}
	if true {
		toSerialize["sequence_number"] = o.SequenceNumber
	}
	if true {
		toSerialize["servicer_address"] = o.ServicerAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LiabilityOverride) UnmarshalJSON(bytes []byte) (err error) {
	varLiabilityOverride := _LiabilityOverride{}

	if err = json.Unmarshal(bytes, &varLiabilityOverride); err == nil {
		*o = LiabilityOverride(varLiabilityOverride)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "purchase_apr")
		delete(additionalProperties, "cash_apr")
		delete(additionalProperties, "balance_transfer_apr")
		delete(additionalProperties, "special_apr")
		delete(additionalProperties, "last_payment_amount")
		delete(additionalProperties, "minimum_payment_amount")
		delete(additionalProperties, "is_overdue")
		delete(additionalProperties, "origination_date")
		delete(additionalProperties, "principal")
		delete(additionalProperties, "nominal_apr")
		delete(additionalProperties, "interest_capitalization_grace_period_months")
		delete(additionalProperties, "repayment_model")
		delete(additionalProperties, "expected_payoff_date")
		delete(additionalProperties, "guarantor")
		delete(additionalProperties, "is_federal")
		delete(additionalProperties, "loan_name")
		delete(additionalProperties, "loan_status")
		delete(additionalProperties, "payment_reference_number")
		delete(additionalProperties, "pslf_status")
		delete(additionalProperties, "repayment_plan_description")
		delete(additionalProperties, "repayment_plan_type")
		delete(additionalProperties, "sequence_number")
		delete(additionalProperties, "servicer_address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLiabilityOverride struct {
	value *LiabilityOverride
	isSet bool
}

func (v NullableLiabilityOverride) Get() *LiabilityOverride {
	return v.value
}

func (v *NullableLiabilityOverride) Set(val *LiabilityOverride) {
	v.value = val
	v.isSet = true
}

func (v NullableLiabilityOverride) IsSet() bool {
	return v.isSet
}

func (v *NullableLiabilityOverride) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiabilityOverride(val *LiabilityOverride) *NullableLiabilityOverride {
	return &NullableLiabilityOverride{value: val, isSet: true}
}

func (v NullableLiabilityOverride) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiabilityOverride) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


