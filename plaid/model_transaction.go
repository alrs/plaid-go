/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.121.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"time"
)

// Transaction A representation of a transaction
type Transaction struct {
	// Please use the `payment_channel` field, `transaction_type` will be deprecated in the future.  `digital:` transactions that took place online.  `place:` transactions that were made at a physical location.  `special:` transactions that relate to banks, e.g. fees or deposits.  `unresolved:` transactions that do not fit into the other three types. 
	TransactionType *string `json:"transaction_type,omitempty"`
	// The ID of a posted transaction's associated pending transaction, where applicable.
	PendingTransactionId NullableString `json:"pending_transaction_id"`
	// The ID of the category to which this transaction belongs. For a full list of categories, see [`/categories/get`](https://plaid.com/docs/api/products/transactions/#categoriesget).  If the `transactions` object was returned by an Assets endpoint such as `/asset_report/get/` or `/asset_report/pdf/get`, this field will only appear in an Asset Report with Insights.
	CategoryId NullableString `json:"category_id"`
	// A hierarchical array of the categories to which this transaction belongs. For a full list of categories, see [`/categories/get`](https://plaid.com/docs/api/products/transactions/#categoriesget).  If the `transactions` object was returned by an Assets endpoint such as `/asset_report/get/` or `/asset_report/pdf/get`, this field will only appear in an Asset Report with Insights.
	Category []string `json:"category"`
	Location Location `json:"location"`
	PaymentMeta PaymentMeta `json:"payment_meta"`
	// The name of the account owner. This field is not typically populated and only relevant when dealing with sub-accounts.
	AccountOwner NullableString `json:"account_owner"`
	// The merchant name or transaction description.  If the `transactions` object was returned by a Transactions endpoint such as `/transactions/get`, this field will always appear. If the `transactions` object was returned by an Assets endpoint such as `/asset_report/get/` or `/asset_report/pdf/get`, this field will only appear in an Asset Report with Insights.
	Name string `json:"name"`
	// The string returned by the financial institution to describe the transaction. For transactions returned by `/transactions/get`, this field is in beta and will be omitted unless the client is both enrolled in the closed beta program and has set `options.include_original_description` to `true`.
	OriginalDescription NullableString `json:"original_description,omitempty"`
	// The ID of the account in which this transaction occurred.
	AccountId string `json:"account_id"`
	// The settled value of the transaction, denominated in the account's currency, as stated in `iso_currency_code` or `unofficial_currency_code`. Positive values when money moves out of the account; negative values when money moves in. For example, debit card purchases are positive; credit card payments, direct deposits, and refunds are negative.
	Amount float64 `json:"amount"`
	// The ISO-4217 currency code of the transaction. Always `null` if `unofficial_currency_code` is non-null.
	IsoCurrencyCode NullableString `json:"iso_currency_code"`
	// The unofficial currency code associated with the transaction. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.
	UnofficialCurrencyCode NullableString `json:"unofficial_currency_code"`
	// For pending transactions, the date that the transaction occurred; for posted transactions, the date that the transaction posted. Both dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ( `YYYY-MM-DD` ).
	Date string `json:"date"`
	// When `true`, identifies the transaction as pending or unsettled. Pending transaction details (name, type, amount, category ID) may change before they are settled.
	Pending bool `json:"pending"`
	// The unique ID of the transaction. Like all Plaid identifiers, the `transaction_id` is case sensitive.
	TransactionId string `json:"transaction_id"`
	// The merchant name, as extracted by Plaid from the `name` field.
	MerchantName NullableString `json:"merchant_name,omitempty"`
	// The check number of the transaction. This field is only populated for check transactions.
	CheckNumber NullableString `json:"check_number,omitempty"`
	// The channel used to make a payment. `online:` transactions that took place online.  `in store:` transactions that were made at a physical location.  `other:` transactions that relate to banks, e.g. fees or deposits.  This field replaces the `transaction_type` field. 
	PaymentChannel string `json:"payment_channel"`
	// The date that the transaction was authorized. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ( `YYYY-MM-DD` ). The `authorized_date` field uses machine learning to determine a transaction date for transactions where the `date_transacted` is not available. If the `date_transacted` field is present and not `null`, the `authorized_date` field will have the same value as the `date_transacted` field.
	AuthorizedDate NullableString `json:"authorized_date"`
	// Date and time when a transaction was authorized in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ( `YYYY-MM-DDTHH:mm:ssZ` ).  This field is returned for select financial institutions and comes as provided by the institution. It may contain default time values (such as 00:00:00).
	AuthorizedDatetime NullableTime `json:"authorized_datetime"`
	// Date and time when a transaction was posted in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ( `YYYY-MM-DDTHH:mm:ssZ` ).  This field is returned for select financial institutions and comes as provided by the institution. It may contain default time values (such as 00:00:00).
	Datetime NullableTime `json:"datetime"`
	TransactionCode NullableTransactionCode `json:"transaction_code"`
	PersonalFinanceCategory NullablePersonalFinanceCategory `json:"personal_finance_category,omitempty"`
}

// NewTransaction instantiates a new Transaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransaction(pendingTransactionId NullableString, categoryId NullableString, category []string, location Location, paymentMeta PaymentMeta, accountOwner NullableString, name string, accountId string, amount float64, isoCurrencyCode NullableString, unofficialCurrencyCode NullableString, date string, pending bool, transactionId string, paymentChannel string, authorizedDate NullableString, authorizedDatetime NullableTime, datetime NullableTime, transactionCode NullableTransactionCode) *Transaction {
	this := Transaction{}
	this.PendingTransactionId = pendingTransactionId
	this.CategoryId = categoryId
	this.Category = category
	this.Location = location
	this.PaymentMeta = paymentMeta
	this.AccountOwner = accountOwner
	this.Name = name
	this.AccountId = accountId
	this.Amount = amount
	this.IsoCurrencyCode = isoCurrencyCode
	this.UnofficialCurrencyCode = unofficialCurrencyCode
	this.Date = date
	this.Pending = pending
	this.TransactionId = transactionId
	this.PaymentChannel = paymentChannel
	this.AuthorizedDate = authorizedDate
	this.AuthorizedDatetime = authorizedDatetime
	this.Datetime = datetime
	this.TransactionCode = transactionCode
	return &this
}

// NewTransactionWithDefaults instantiates a new Transaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionWithDefaults() *Transaction {
	this := Transaction{}
	return &this
}

// GetTransactionType returns the TransactionType field value if set, zero value otherwise.
func (o *Transaction) GetTransactionType() string {
	if o == nil || o.TransactionType == nil {
		var ret string
		return ret
	}
	return *o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetTransactionTypeOk() (*string, bool) {
	if o == nil || o.TransactionType == nil {
		return nil, false
	}
	return o.TransactionType, true
}

// HasTransactionType returns a boolean if a field has been set.
func (o *Transaction) HasTransactionType() bool {
	if o != nil && o.TransactionType != nil {
		return true
	}

	return false
}

// SetTransactionType gets a reference to the given string and assigns it to the TransactionType field.
func (o *Transaction) SetTransactionType(v string) {
	o.TransactionType = &v
}

// GetPendingTransactionId returns the PendingTransactionId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetPendingTransactionId() string {
	if o == nil || o.PendingTransactionId.Get() == nil {
		var ret string
		return ret
	}

	return *o.PendingTransactionId.Get()
}

// GetPendingTransactionIdOk returns a tuple with the PendingTransactionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetPendingTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PendingTransactionId.Get(), o.PendingTransactionId.IsSet()
}

// SetPendingTransactionId sets field value
func (o *Transaction) SetPendingTransactionId(v string) {
	o.PendingTransactionId.Set(&v)
}

// GetCategoryId returns the CategoryId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetCategoryId() string {
	if o == nil || o.CategoryId.Get() == nil {
		var ret string
		return ret
	}

	return *o.CategoryId.Get()
}

// GetCategoryIdOk returns a tuple with the CategoryId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetCategoryIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CategoryId.Get(), o.CategoryId.IsSet()
}

// SetCategoryId sets field value
func (o *Transaction) SetCategoryId(v string) {
	o.CategoryId.Set(&v)
}

// GetCategory returns the Category field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *Transaction) GetCategory() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetCategoryOk() (*[]string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *Transaction) SetCategory(v []string) {
	o.Category = v
}

// GetLocation returns the Location field value
func (o *Transaction) GetLocation() Location {
	if o == nil {
		var ret Location
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetLocationOk() (*Location, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *Transaction) SetLocation(v Location) {
	o.Location = v
}

// GetPaymentMeta returns the PaymentMeta field value
func (o *Transaction) GetPaymentMeta() PaymentMeta {
	if o == nil {
		var ret PaymentMeta
		return ret
	}

	return o.PaymentMeta
}

// GetPaymentMetaOk returns a tuple with the PaymentMeta field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetPaymentMetaOk() (*PaymentMeta, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentMeta, true
}

// SetPaymentMeta sets field value
func (o *Transaction) SetPaymentMeta(v PaymentMeta) {
	o.PaymentMeta = v
}

// GetAccountOwner returns the AccountOwner field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetAccountOwner() string {
	if o == nil || o.AccountOwner.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccountOwner.Get()
}

// GetAccountOwnerOk returns a tuple with the AccountOwner field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetAccountOwnerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountOwner.Get(), o.AccountOwner.IsSet()
}

// SetAccountOwner sets field value
func (o *Transaction) SetAccountOwner(v string) {
	o.AccountOwner.Set(&v)
}

// GetName returns the Name field value
func (o *Transaction) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Transaction) SetName(v string) {
	o.Name = v
}

// GetOriginalDescription returns the OriginalDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetOriginalDescription() string {
	if o == nil || o.OriginalDescription.Get() == nil {
		var ret string
		return ret
	}
	return *o.OriginalDescription.Get()
}

// GetOriginalDescriptionOk returns a tuple with the OriginalDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetOriginalDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginalDescription.Get(), o.OriginalDescription.IsSet()
}

// HasOriginalDescription returns a boolean if a field has been set.
func (o *Transaction) HasOriginalDescription() bool {
	if o != nil && o.OriginalDescription.IsSet() {
		return true
	}

	return false
}

// SetOriginalDescription gets a reference to the given NullableString and assigns it to the OriginalDescription field.
func (o *Transaction) SetOriginalDescription(v string) {
	o.OriginalDescription.Set(&v)
}
// SetOriginalDescriptionNil sets the value for OriginalDescription to be an explicit nil
func (o *Transaction) SetOriginalDescriptionNil() {
	o.OriginalDescription.Set(nil)
}

// UnsetOriginalDescription ensures that no value is present for OriginalDescription, not even an explicit nil
func (o *Transaction) UnsetOriginalDescription() {
	o.OriginalDescription.Unset()
}

// GetAccountId returns the AccountId field value
func (o *Transaction) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *Transaction) SetAccountId(v string) {
	o.AccountId = v
}

// GetAmount returns the Amount field value
func (o *Transaction) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Transaction) SetAmount(v float64) {
	o.Amount = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.IsoCurrencyCode.Get()
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsoCurrencyCode.Get(), o.IsoCurrencyCode.IsSet()
}

// SetIsoCurrencyCode sets field value
func (o *Transaction) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode.Set(&v)
}

// GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetUnofficialCurrencyCode() string {
	if o == nil || o.UnofficialCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.UnofficialCurrencyCode.Get()
}

// GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetUnofficialCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnofficialCurrencyCode.Get(), o.UnofficialCurrencyCode.IsSet()
}

// SetUnofficialCurrencyCode sets field value
func (o *Transaction) SetUnofficialCurrencyCode(v string) {
	o.UnofficialCurrencyCode.Set(&v)
}

// GetDate returns the Date field value
func (o *Transaction) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *Transaction) SetDate(v string) {
	o.Date = v
}

// GetPending returns the Pending field value
func (o *Transaction) GetPending() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Pending
}

// GetPendingOk returns a tuple with the Pending field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetPendingOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pending, true
}

// SetPending sets field value
func (o *Transaction) SetPending(v bool) {
	o.Pending = v
}

// GetTransactionId returns the TransactionId field value
func (o *Transaction) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *Transaction) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetMerchantName returns the MerchantName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetMerchantName() string {
	if o == nil || o.MerchantName.Get() == nil {
		var ret string
		return ret
	}
	return *o.MerchantName.Get()
}

// GetMerchantNameOk returns a tuple with the MerchantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetMerchantNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MerchantName.Get(), o.MerchantName.IsSet()
}

// HasMerchantName returns a boolean if a field has been set.
func (o *Transaction) HasMerchantName() bool {
	if o != nil && o.MerchantName.IsSet() {
		return true
	}

	return false
}

// SetMerchantName gets a reference to the given NullableString and assigns it to the MerchantName field.
func (o *Transaction) SetMerchantName(v string) {
	o.MerchantName.Set(&v)
}
// SetMerchantNameNil sets the value for MerchantName to be an explicit nil
func (o *Transaction) SetMerchantNameNil() {
	o.MerchantName.Set(nil)
}

// UnsetMerchantName ensures that no value is present for MerchantName, not even an explicit nil
func (o *Transaction) UnsetMerchantName() {
	o.MerchantName.Unset()
}

// GetCheckNumber returns the CheckNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetCheckNumber() string {
	if o == nil || o.CheckNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.CheckNumber.Get()
}

// GetCheckNumberOk returns a tuple with the CheckNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetCheckNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CheckNumber.Get(), o.CheckNumber.IsSet()
}

// HasCheckNumber returns a boolean if a field has been set.
func (o *Transaction) HasCheckNumber() bool {
	if o != nil && o.CheckNumber.IsSet() {
		return true
	}

	return false
}

// SetCheckNumber gets a reference to the given NullableString and assigns it to the CheckNumber field.
func (o *Transaction) SetCheckNumber(v string) {
	o.CheckNumber.Set(&v)
}
// SetCheckNumberNil sets the value for CheckNumber to be an explicit nil
func (o *Transaction) SetCheckNumberNil() {
	o.CheckNumber.Set(nil)
}

// UnsetCheckNumber ensures that no value is present for CheckNumber, not even an explicit nil
func (o *Transaction) UnsetCheckNumber() {
	o.CheckNumber.Unset()
}

// GetPaymentChannel returns the PaymentChannel field value
func (o *Transaction) GetPaymentChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentChannel
}

// GetPaymentChannelOk returns a tuple with the PaymentChannel field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetPaymentChannelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentChannel, true
}

// SetPaymentChannel sets field value
func (o *Transaction) SetPaymentChannel(v string) {
	o.PaymentChannel = v
}

// GetAuthorizedDate returns the AuthorizedDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetAuthorizedDate() string {
	if o == nil || o.AuthorizedDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.AuthorizedDate.Get()
}

// GetAuthorizedDateOk returns a tuple with the AuthorizedDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetAuthorizedDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AuthorizedDate.Get(), o.AuthorizedDate.IsSet()
}

// SetAuthorizedDate sets field value
func (o *Transaction) SetAuthorizedDate(v string) {
	o.AuthorizedDate.Set(&v)
}

// GetAuthorizedDatetime returns the AuthorizedDatetime field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Transaction) GetAuthorizedDatetime() time.Time {
	if o == nil || o.AuthorizedDatetime.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.AuthorizedDatetime.Get()
}

// GetAuthorizedDatetimeOk returns a tuple with the AuthorizedDatetime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetAuthorizedDatetimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AuthorizedDatetime.Get(), o.AuthorizedDatetime.IsSet()
}

// SetAuthorizedDatetime sets field value
func (o *Transaction) SetAuthorizedDatetime(v time.Time) {
	o.AuthorizedDatetime.Set(&v)
}

// GetDatetime returns the Datetime field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Transaction) GetDatetime() time.Time {
	if o == nil || o.Datetime.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Datetime.Get()
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetDatetimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Datetime.Get(), o.Datetime.IsSet()
}

// SetDatetime sets field value
func (o *Transaction) SetDatetime(v time.Time) {
	o.Datetime.Set(&v)
}

// GetTransactionCode returns the TransactionCode field value
// If the value is explicit nil, the zero value for TransactionCode will be returned
func (o *Transaction) GetTransactionCode() TransactionCode {
	if o == nil || o.TransactionCode.Get() == nil {
		var ret TransactionCode
		return ret
	}

	return *o.TransactionCode.Get()
}

// GetTransactionCodeOk returns a tuple with the TransactionCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetTransactionCodeOk() (*TransactionCode, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TransactionCode.Get(), o.TransactionCode.IsSet()
}

// SetTransactionCode sets field value
func (o *Transaction) SetTransactionCode(v TransactionCode) {
	o.TransactionCode.Set(&v)
}

// GetPersonalFinanceCategory returns the PersonalFinanceCategory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetPersonalFinanceCategory() PersonalFinanceCategory {
	if o == nil || o.PersonalFinanceCategory.Get() == nil {
		var ret PersonalFinanceCategory
		return ret
	}
	return *o.PersonalFinanceCategory.Get()
}

// GetPersonalFinanceCategoryOk returns a tuple with the PersonalFinanceCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetPersonalFinanceCategoryOk() (*PersonalFinanceCategory, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PersonalFinanceCategory.Get(), o.PersonalFinanceCategory.IsSet()
}

// HasPersonalFinanceCategory returns a boolean if a field has been set.
func (o *Transaction) HasPersonalFinanceCategory() bool {
	if o != nil && o.PersonalFinanceCategory.IsSet() {
		return true
	}

	return false
}

// SetPersonalFinanceCategory gets a reference to the given NullablePersonalFinanceCategory and assigns it to the PersonalFinanceCategory field.
func (o *Transaction) SetPersonalFinanceCategory(v PersonalFinanceCategory) {
	o.PersonalFinanceCategory.Set(&v)
}
// SetPersonalFinanceCategoryNil sets the value for PersonalFinanceCategory to be an explicit nil
func (o *Transaction) SetPersonalFinanceCategoryNil() {
	o.PersonalFinanceCategory.Set(nil)
}

// UnsetPersonalFinanceCategory ensures that no value is present for PersonalFinanceCategory, not even an explicit nil
func (o *Transaction) UnsetPersonalFinanceCategory() {
	o.PersonalFinanceCategory.Unset()
}

func (o Transaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TransactionType != nil {
		toSerialize["transaction_type"] = o.TransactionType
	}
	if true {
		toSerialize["pending_transaction_id"] = o.PendingTransactionId.Get()
	}
	if true {
		toSerialize["category_id"] = o.CategoryId.Get()
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if true {
		toSerialize["location"] = o.Location
	}
	if true {
		toSerialize["payment_meta"] = o.PaymentMeta
	}
	if true {
		toSerialize["account_owner"] = o.AccountOwner.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.OriginalDescription.IsSet() {
		toSerialize["original_description"] = o.OriginalDescription.Get()
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode.Get()
	}
	if true {
		toSerialize["unofficial_currency_code"] = o.UnofficialCurrencyCode.Get()
	}
	if true {
		toSerialize["date"] = o.Date
	}
	if true {
		toSerialize["pending"] = o.Pending
	}
	if true {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if o.MerchantName.IsSet() {
		toSerialize["merchant_name"] = o.MerchantName.Get()
	}
	if o.CheckNumber.IsSet() {
		toSerialize["check_number"] = o.CheckNumber.Get()
	}
	if true {
		toSerialize["payment_channel"] = o.PaymentChannel
	}
	if true {
		toSerialize["authorized_date"] = o.AuthorizedDate.Get()
	}
	if true {
		toSerialize["authorized_datetime"] = o.AuthorizedDatetime.Get()
	}
	if true {
		toSerialize["datetime"] = o.Datetime.Get()
	}
	if true {
		toSerialize["transaction_code"] = o.TransactionCode.Get()
	}
	if o.PersonalFinanceCategory.IsSet() {
		toSerialize["personal_finance_category"] = o.PersonalFinanceCategory.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTransaction struct {
	value *Transaction
	isSet bool
}

func (v NullableTransaction) Get() *Transaction {
	return v.value
}

func (v *NullableTransaction) Set(val *Transaction) {
	v.value = val
	v.isSet = true
}

func (v NullableTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransaction(val *Transaction) *NullableTransaction {
	return &NullableTransaction{value: val, isSet: true}
}

func (v NullableTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


