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

// StudentLoan Contains details about a student loan account
type StudentLoan struct {
	// The ID of the account that this liability belongs to.
	AccountId NullableString `json:"account_id"`
	// The account number of the loan. For some institutions, this may be a masked version of the number (e.g., the last 4 digits instead of the entire number).
	AccountNumber NullableString `json:"account_number"`
	// The dates on which loaned funds were disbursed or will be disbursed. These are often in the past. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).
	DisbursementDates []string `json:"disbursement_dates"`
	// The date when the student loan is expected to be paid off. Availability for this field is limited. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).
	ExpectedPayoffDate NullableString `json:"expected_payoff_date"`
	// The guarantor of the student loan.
	Guarantor NullableString `json:"guarantor"`
	// The interest rate on the loan as a percentage.
	InterestRatePercentage float32 `json:"interest_rate_percentage"`
	// `true` if a payment is currently overdue. Availability for this field is limited.
	IsOverdue NullableBool `json:"is_overdue"`
	// The amount of the last payment.
	LastPaymentAmount NullableFloat32 `json:"last_payment_amount"`
	// The date of the last payment. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).
	LastPaymentDate NullableString `json:"last_payment_date"`
	// The date of the last statement. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).
	LastStatementIssueDate NullableString `json:"last_statement_issue_date"`
	// The type of loan, e.g., \"Consolidation Loans\".
	LoanName NullableString `json:"loan_name"`
	LoanStatus StudentLoanStatus `json:"loan_status"`
	// The minimum payment due for the next billing cycle. There are some exceptions: Some institutions require a minimum payment across all loans associated with an account number. Our API presents that same minimum payment amount on each loan. The institutions that do this are: Great Lakes ( `ins_116861`), Firstmark (`ins_116295`), Commonbond Firstmark Services (`ins_116950`), Nelnet (`ins_116528`), EdFinancial Services (`ins_116304`), Granite State (`ins_116308`), and Oklahoma Student Loan Authority (`ins_116945`). Firstmark (`ins_116295` ) and Navient (`ins_116248`) will display as $0 if there is an autopay program in effect.
	MinimumPaymentAmount NullableFloat32 `json:"minimum_payment_amount"`
	// The due date for the next payment. The due date is `null` if a payment is not expected. A payment is not expected if `loan_status.type` is `deferment`, `in_school`, `consolidated`, `paid in full`, or `transferred`. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).
	NextPaymentDueDate NullableString `json:"next_payment_due_date"`
	// The date on which the loan was initially lent. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). 
	OriginationDate NullableString `json:"origination_date"`
	// The original principal balance of the loan.
	OriginationPrincipalAmount NullableFloat32 `json:"origination_principal_amount"`
	// The total dollar amount of the accrued interest balance. For Sallie Mae ( `ins_116944`), this amount is included in the current balance of the loan, so this field will return as `null`.
	OutstandingInterestAmount NullableFloat32 `json:"outstanding_interest_amount"`
	// The relevant account number that should be used to reference this loan for payments. In the majority of cases, `payment_reference_number` will match a`ccount_number,` but in some institutions, such as Great Lakes (`ins_116861`), it will be different.
	PaymentReferenceNumber NullableString `json:"payment_reference_number"`
	PslfStatus PSLFStatus `json:"pslf_status"`
	RepaymentPlan StudentRepaymentPlan `json:"repayment_plan"`
	// The sequence number of the student loan. Heartland ECSI (`ins_116948`) does not make this field available.
	SequenceNumber NullableString `json:"sequence_number"`
	ServicerAddress ServicerAddressData `json:"servicer_address"`
	// The year to date (YTD) interest paid. Availability for this field is limited.
	YtdInterestPaid NullableFloat32 `json:"ytd_interest_paid"`
	// The year to date (YTD) principal paid. Availability for this field is limited.
	YtdPrincipalPaid NullableFloat32 `json:"ytd_principal_paid"`
	AdditionalProperties map[string]interface{}
}

type _StudentLoan StudentLoan

// NewStudentLoan instantiates a new StudentLoan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStudentLoan(accountId NullableString, accountNumber NullableString, disbursementDates []string, expectedPayoffDate NullableString, guarantor NullableString, interestRatePercentage float32, isOverdue NullableBool, lastPaymentAmount NullableFloat32, lastPaymentDate NullableString, lastStatementIssueDate NullableString, loanName NullableString, loanStatus StudentLoanStatus, minimumPaymentAmount NullableFloat32, nextPaymentDueDate NullableString, originationDate NullableString, originationPrincipalAmount NullableFloat32, outstandingInterestAmount NullableFloat32, paymentReferenceNumber NullableString, pslfStatus PSLFStatus, repaymentPlan StudentRepaymentPlan, sequenceNumber NullableString, servicerAddress ServicerAddressData, ytdInterestPaid NullableFloat32, ytdPrincipalPaid NullableFloat32) *StudentLoan {
	this := StudentLoan{}
	this.AccountId = accountId
	this.AccountNumber = accountNumber
	this.DisbursementDates = disbursementDates
	this.ExpectedPayoffDate = expectedPayoffDate
	this.Guarantor = guarantor
	this.InterestRatePercentage = interestRatePercentage
	this.IsOverdue = isOverdue
	this.LastPaymentAmount = lastPaymentAmount
	this.LastPaymentDate = lastPaymentDate
	this.LastStatementIssueDate = lastStatementIssueDate
	this.LoanName = loanName
	this.LoanStatus = loanStatus
	this.MinimumPaymentAmount = minimumPaymentAmount
	this.NextPaymentDueDate = nextPaymentDueDate
	this.OriginationDate = originationDate
	this.OriginationPrincipalAmount = originationPrincipalAmount
	this.OutstandingInterestAmount = outstandingInterestAmount
	this.PaymentReferenceNumber = paymentReferenceNumber
	this.PslfStatus = pslfStatus
	this.RepaymentPlan = repaymentPlan
	this.SequenceNumber = sequenceNumber
	this.ServicerAddress = servicerAddress
	this.YtdInterestPaid = ytdInterestPaid
	this.YtdPrincipalPaid = ytdPrincipalPaid
	return &this
}

// NewStudentLoanWithDefaults instantiates a new StudentLoan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStudentLoanWithDefaults() *StudentLoan {
	this := StudentLoan{}
	return &this
}

// GetAccountId returns the AccountId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StudentLoan) GetAccountId() string {
	if o == nil || o.AccountId.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccountId.Get()
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StudentLoan) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountId.Get(), o.AccountId.IsSet()
}

// SetAccountId sets field value
func (o *StudentLoan) SetAccountId(v string) {
	o.AccountId.Set(&v)
}

// GetAccountNumber returns the AccountNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StudentLoan) GetAccountNumber() string {
	if o == nil || o.AccountNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccountNumber.Get()
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StudentLoan) GetAccountNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountNumber.Get(), o.AccountNumber.IsSet()
}

// SetAccountNumber sets field value
func (o *StudentLoan) SetAccountNumber(v string) {
	o.AccountNumber.Set(&v)
}

// GetDisbursementDates returns the DisbursementDates field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *StudentLoan) GetDisbursementDates() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DisbursementDates
}

// GetDisbursementDatesOk returns a tuple with the DisbursementDates field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StudentLoan) GetDisbursementDatesOk() (*[]string, bool) {
	if o == nil || o.DisbursementDates == nil {
		return nil, false
	}
	return &o.DisbursementDates, true
}

// SetDisbursementDates sets field value
func (o *StudentLoan) SetDisbursementDates(v []string) {
	o.DisbursementDates = v
}

// GetExpectedPayoffDate returns the ExpectedPayoffDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StudentLoan) GetExpectedPayoffDate() string {
	if o == nil || o.ExpectedPayoffDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExpectedPayoffDate.Get()
}

// GetExpectedPayoffDateOk returns a tuple with the ExpectedPayoffDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StudentLoan) GetExpectedPayoffDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpectedPayoffDate.Get(), o.ExpectedPayoffDate.IsSet()
}

// SetExpectedPayoffDate sets field value
func (o *StudentLoan) SetExpectedPayoffDate(v string) {
	o.ExpectedPayoffDate.Set(&v)
}

// GetGuarantor returns the Guarantor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StudentLoan) GetGuarantor() string {
	if o == nil || o.Guarantor.Get() == nil {
		var ret string
		return ret
	}

	return *o.Guarantor.Get()
}

// GetGuarantorOk returns a tuple with the Guarantor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StudentLoan) GetGuarantorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Guarantor.Get(), o.Guarantor.IsSet()
}

// SetGuarantor sets field value
func (o *StudentLoan) SetGuarantor(v string) {
	o.Guarantor.Set(&v)
}

// GetInterestRatePercentage returns the InterestRatePercentage field value
func (o *StudentLoan) GetInterestRatePercentage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.InterestRatePercentage
}

// GetInterestRatePercentageOk returns a tuple with the InterestRatePercentage field value
// and a boolean to check if the value has been set.
func (o *StudentLoan) GetInterestRatePercentageOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InterestRatePercentage, true
}

// SetInterestRatePercentage sets field value
func (o *StudentLoan) SetInterestRatePercentage(v float32) {
	o.InterestRatePercentage = v
}

// GetIsOverdue returns the IsOverdue field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *StudentLoan) GetIsOverdue() bool {
	if o == nil || o.IsOverdue.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsOverdue.Get()
}

// GetIsOverdueOk returns a tuple with the IsOverdue field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StudentLoan) GetIsOverdueOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsOverdue.Get(), o.IsOverdue.IsSet()
}

// SetIsOverdue sets field value
func (o *StudentLoan) SetIsOverdue(v bool) {
	o.IsOverdue.Set(&v)
}

// GetLastPaymentAmount returns the LastPaymentAmount field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *StudentLoan) GetLastPaymentAmount() float32 {
	if o == nil || o.LastPaymentAmount.Get() == nil {
		var ret float32
		return ret
	}

	return *o.LastPaymentAmount.Get()
}

// GetLastPaymentAmountOk returns a tuple with the LastPaymentAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StudentLoan) GetLastPaymentAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastPaymentAmount.Get(), o.LastPaymentAmount.IsSet()
}

// SetLastPaymentAmount sets field value
func (o *StudentLoan) SetLastPaymentAmount(v float32) {
	o.LastPaymentAmount.Set(&v)
}

// GetLastPaymentDate returns the LastPaymentDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StudentLoan) GetLastPaymentDate() string {
	if o == nil || o.LastPaymentDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastPaymentDate.Get()
}

// GetLastPaymentDateOk returns a tuple with the LastPaymentDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StudentLoan) GetLastPaymentDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastPaymentDate.Get(), o.LastPaymentDate.IsSet()
}

// SetLastPaymentDate sets field value
func (o *StudentLoan) SetLastPaymentDate(v string) {
	o.LastPaymentDate.Set(&v)
}

// GetLastStatementIssueDate returns the LastStatementIssueDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StudentLoan) GetLastStatementIssueDate() string {
	if o == nil || o.LastStatementIssueDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastStatementIssueDate.Get()
}

// GetLastStatementIssueDateOk returns a tuple with the LastStatementIssueDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StudentLoan) GetLastStatementIssueDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastStatementIssueDate.Get(), o.LastStatementIssueDate.IsSet()
}

// SetLastStatementIssueDate sets field value
func (o *StudentLoan) SetLastStatementIssueDate(v string) {
	o.LastStatementIssueDate.Set(&v)
}

// GetLoanName returns the LoanName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StudentLoan) GetLoanName() string {
	if o == nil || o.LoanName.Get() == nil {
		var ret string
		return ret
	}

	return *o.LoanName.Get()
}

// GetLoanNameOk returns a tuple with the LoanName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StudentLoan) GetLoanNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LoanName.Get(), o.LoanName.IsSet()
}

// SetLoanName sets field value
func (o *StudentLoan) SetLoanName(v string) {
	o.LoanName.Set(&v)
}

// GetLoanStatus returns the LoanStatus field value
func (o *StudentLoan) GetLoanStatus() StudentLoanStatus {
	if o == nil {
		var ret StudentLoanStatus
		return ret
	}

	return o.LoanStatus
}

// GetLoanStatusOk returns a tuple with the LoanStatus field value
// and a boolean to check if the value has been set.
func (o *StudentLoan) GetLoanStatusOk() (*StudentLoanStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LoanStatus, true
}

// SetLoanStatus sets field value
func (o *StudentLoan) SetLoanStatus(v StudentLoanStatus) {
	o.LoanStatus = v
}

// GetMinimumPaymentAmount returns the MinimumPaymentAmount field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *StudentLoan) GetMinimumPaymentAmount() float32 {
	if o == nil || o.MinimumPaymentAmount.Get() == nil {
		var ret float32
		return ret
	}

	return *o.MinimumPaymentAmount.Get()
}

// GetMinimumPaymentAmountOk returns a tuple with the MinimumPaymentAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StudentLoan) GetMinimumPaymentAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MinimumPaymentAmount.Get(), o.MinimumPaymentAmount.IsSet()
}

// SetMinimumPaymentAmount sets field value
func (o *StudentLoan) SetMinimumPaymentAmount(v float32) {
	o.MinimumPaymentAmount.Set(&v)
}

// GetNextPaymentDueDate returns the NextPaymentDueDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StudentLoan) GetNextPaymentDueDate() string {
	if o == nil || o.NextPaymentDueDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextPaymentDueDate.Get()
}

// GetNextPaymentDueDateOk returns a tuple with the NextPaymentDueDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StudentLoan) GetNextPaymentDueDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextPaymentDueDate.Get(), o.NextPaymentDueDate.IsSet()
}

// SetNextPaymentDueDate sets field value
func (o *StudentLoan) SetNextPaymentDueDate(v string) {
	o.NextPaymentDueDate.Set(&v)
}

// GetOriginationDate returns the OriginationDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StudentLoan) GetOriginationDate() string {
	if o == nil || o.OriginationDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.OriginationDate.Get()
}

// GetOriginationDateOk returns a tuple with the OriginationDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StudentLoan) GetOriginationDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginationDate.Get(), o.OriginationDate.IsSet()
}

// SetOriginationDate sets field value
func (o *StudentLoan) SetOriginationDate(v string) {
	o.OriginationDate.Set(&v)
}

// GetOriginationPrincipalAmount returns the OriginationPrincipalAmount field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *StudentLoan) GetOriginationPrincipalAmount() float32 {
	if o == nil || o.OriginationPrincipalAmount.Get() == nil {
		var ret float32
		return ret
	}

	return *o.OriginationPrincipalAmount.Get()
}

// GetOriginationPrincipalAmountOk returns a tuple with the OriginationPrincipalAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StudentLoan) GetOriginationPrincipalAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginationPrincipalAmount.Get(), o.OriginationPrincipalAmount.IsSet()
}

// SetOriginationPrincipalAmount sets field value
func (o *StudentLoan) SetOriginationPrincipalAmount(v float32) {
	o.OriginationPrincipalAmount.Set(&v)
}

// GetOutstandingInterestAmount returns the OutstandingInterestAmount field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *StudentLoan) GetOutstandingInterestAmount() float32 {
	if o == nil || o.OutstandingInterestAmount.Get() == nil {
		var ret float32
		return ret
	}

	return *o.OutstandingInterestAmount.Get()
}

// GetOutstandingInterestAmountOk returns a tuple with the OutstandingInterestAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StudentLoan) GetOutstandingInterestAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OutstandingInterestAmount.Get(), o.OutstandingInterestAmount.IsSet()
}

// SetOutstandingInterestAmount sets field value
func (o *StudentLoan) SetOutstandingInterestAmount(v float32) {
	o.OutstandingInterestAmount.Set(&v)
}

// GetPaymentReferenceNumber returns the PaymentReferenceNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StudentLoan) GetPaymentReferenceNumber() string {
	if o == nil || o.PaymentReferenceNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.PaymentReferenceNumber.Get()
}

// GetPaymentReferenceNumberOk returns a tuple with the PaymentReferenceNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StudentLoan) GetPaymentReferenceNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaymentReferenceNumber.Get(), o.PaymentReferenceNumber.IsSet()
}

// SetPaymentReferenceNumber sets field value
func (o *StudentLoan) SetPaymentReferenceNumber(v string) {
	o.PaymentReferenceNumber.Set(&v)
}

// GetPslfStatus returns the PslfStatus field value
func (o *StudentLoan) GetPslfStatus() PSLFStatus {
	if o == nil {
		var ret PSLFStatus
		return ret
	}

	return o.PslfStatus
}

// GetPslfStatusOk returns a tuple with the PslfStatus field value
// and a boolean to check if the value has been set.
func (o *StudentLoan) GetPslfStatusOk() (*PSLFStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PslfStatus, true
}

// SetPslfStatus sets field value
func (o *StudentLoan) SetPslfStatus(v PSLFStatus) {
	o.PslfStatus = v
}

// GetRepaymentPlan returns the RepaymentPlan field value
func (o *StudentLoan) GetRepaymentPlan() StudentRepaymentPlan {
	if o == nil {
		var ret StudentRepaymentPlan
		return ret
	}

	return o.RepaymentPlan
}

// GetRepaymentPlanOk returns a tuple with the RepaymentPlan field value
// and a boolean to check if the value has been set.
func (o *StudentLoan) GetRepaymentPlanOk() (*StudentRepaymentPlan, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RepaymentPlan, true
}

// SetRepaymentPlan sets field value
func (o *StudentLoan) SetRepaymentPlan(v StudentRepaymentPlan) {
	o.RepaymentPlan = v
}

// GetSequenceNumber returns the SequenceNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StudentLoan) GetSequenceNumber() string {
	if o == nil || o.SequenceNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.SequenceNumber.Get()
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StudentLoan) GetSequenceNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SequenceNumber.Get(), o.SequenceNumber.IsSet()
}

// SetSequenceNumber sets field value
func (o *StudentLoan) SetSequenceNumber(v string) {
	o.SequenceNumber.Set(&v)
}

// GetServicerAddress returns the ServicerAddress field value
func (o *StudentLoan) GetServicerAddress() ServicerAddressData {
	if o == nil {
		var ret ServicerAddressData
		return ret
	}

	return o.ServicerAddress
}

// GetServicerAddressOk returns a tuple with the ServicerAddress field value
// and a boolean to check if the value has been set.
func (o *StudentLoan) GetServicerAddressOk() (*ServicerAddressData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServicerAddress, true
}

// SetServicerAddress sets field value
func (o *StudentLoan) SetServicerAddress(v ServicerAddressData) {
	o.ServicerAddress = v
}

// GetYtdInterestPaid returns the YtdInterestPaid field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *StudentLoan) GetYtdInterestPaid() float32 {
	if o == nil || o.YtdInterestPaid.Get() == nil {
		var ret float32
		return ret
	}

	return *o.YtdInterestPaid.Get()
}

// GetYtdInterestPaidOk returns a tuple with the YtdInterestPaid field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StudentLoan) GetYtdInterestPaidOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.YtdInterestPaid.Get(), o.YtdInterestPaid.IsSet()
}

// SetYtdInterestPaid sets field value
func (o *StudentLoan) SetYtdInterestPaid(v float32) {
	o.YtdInterestPaid.Set(&v)
}

// GetYtdPrincipalPaid returns the YtdPrincipalPaid field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *StudentLoan) GetYtdPrincipalPaid() float32 {
	if o == nil || o.YtdPrincipalPaid.Get() == nil {
		var ret float32
		return ret
	}

	return *o.YtdPrincipalPaid.Get()
}

// GetYtdPrincipalPaidOk returns a tuple with the YtdPrincipalPaid field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StudentLoan) GetYtdPrincipalPaidOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.YtdPrincipalPaid.Get(), o.YtdPrincipalPaid.IsSet()
}

// SetYtdPrincipalPaid sets field value
func (o *StudentLoan) SetYtdPrincipalPaid(v float32) {
	o.YtdPrincipalPaid.Set(&v)
}

func (o StudentLoan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId.Get()
	}
	if true {
		toSerialize["account_number"] = o.AccountNumber.Get()
	}
	if o.DisbursementDates != nil {
		toSerialize["disbursement_dates"] = o.DisbursementDates
	}
	if true {
		toSerialize["expected_payoff_date"] = o.ExpectedPayoffDate.Get()
	}
	if true {
		toSerialize["guarantor"] = o.Guarantor.Get()
	}
	if true {
		toSerialize["interest_rate_percentage"] = o.InterestRatePercentage
	}
	if true {
		toSerialize["is_overdue"] = o.IsOverdue.Get()
	}
	if true {
		toSerialize["last_payment_amount"] = o.LastPaymentAmount.Get()
	}
	if true {
		toSerialize["last_payment_date"] = o.LastPaymentDate.Get()
	}
	if true {
		toSerialize["last_statement_issue_date"] = o.LastStatementIssueDate.Get()
	}
	if true {
		toSerialize["loan_name"] = o.LoanName.Get()
	}
	if true {
		toSerialize["loan_status"] = o.LoanStatus
	}
	if true {
		toSerialize["minimum_payment_amount"] = o.MinimumPaymentAmount.Get()
	}
	if true {
		toSerialize["next_payment_due_date"] = o.NextPaymentDueDate.Get()
	}
	if true {
		toSerialize["origination_date"] = o.OriginationDate.Get()
	}
	if true {
		toSerialize["origination_principal_amount"] = o.OriginationPrincipalAmount.Get()
	}
	if true {
		toSerialize["outstanding_interest_amount"] = o.OutstandingInterestAmount.Get()
	}
	if true {
		toSerialize["payment_reference_number"] = o.PaymentReferenceNumber.Get()
	}
	if true {
		toSerialize["pslf_status"] = o.PslfStatus
	}
	if true {
		toSerialize["repayment_plan"] = o.RepaymentPlan
	}
	if true {
		toSerialize["sequence_number"] = o.SequenceNumber.Get()
	}
	if true {
		toSerialize["servicer_address"] = o.ServicerAddress
	}
	if true {
		toSerialize["ytd_interest_paid"] = o.YtdInterestPaid.Get()
	}
	if true {
		toSerialize["ytd_principal_paid"] = o.YtdPrincipalPaid.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StudentLoan) UnmarshalJSON(bytes []byte) (err error) {
	varStudentLoan := _StudentLoan{}

	if err = json.Unmarshal(bytes, &varStudentLoan); err == nil {
		*o = StudentLoan(varStudentLoan)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "account_number")
		delete(additionalProperties, "disbursement_dates")
		delete(additionalProperties, "expected_payoff_date")
		delete(additionalProperties, "guarantor")
		delete(additionalProperties, "interest_rate_percentage")
		delete(additionalProperties, "is_overdue")
		delete(additionalProperties, "last_payment_amount")
		delete(additionalProperties, "last_payment_date")
		delete(additionalProperties, "last_statement_issue_date")
		delete(additionalProperties, "loan_name")
		delete(additionalProperties, "loan_status")
		delete(additionalProperties, "minimum_payment_amount")
		delete(additionalProperties, "next_payment_due_date")
		delete(additionalProperties, "origination_date")
		delete(additionalProperties, "origination_principal_amount")
		delete(additionalProperties, "outstanding_interest_amount")
		delete(additionalProperties, "payment_reference_number")
		delete(additionalProperties, "pslf_status")
		delete(additionalProperties, "repayment_plan")
		delete(additionalProperties, "sequence_number")
		delete(additionalProperties, "servicer_address")
		delete(additionalProperties, "ytd_interest_paid")
		delete(additionalProperties, "ytd_principal_paid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStudentLoan struct {
	value *StudentLoan
	isSet bool
}

func (v NullableStudentLoan) Get() *StudentLoan {
	return v.value
}

func (v *NullableStudentLoan) Set(val *StudentLoan) {
	v.value = val
	v.isSet = true
}

func (v NullableStudentLoan) IsSet() bool {
	return v.isSet
}

func (v *NullableStudentLoan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStudentLoan(val *StudentLoan) *NullableStudentLoan {
	return &NullableStudentLoan{value: val, isSet: true}
}

func (v NullableStudentLoan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStudentLoan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


