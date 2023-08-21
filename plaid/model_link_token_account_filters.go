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
)

// LinkTokenAccountFilters By default, Link will provide limited account filtering: it will only display Institutions that are compatible with all products supplied in the `products` parameter of `/link/token/create`, and, if `auth` is specified in the `products` array, will also filter out accounts other than `checking` and `savings` accounts on the Account Select pane. You can further limit the accounts shown in Link by using `account_filters` to specify the account subtypes to be shown in Link. Only the specified subtypes will be shown. This filtering applies to both the Account Select view (if enabled) and the Institution Select view. Institutions that do not support the selected subtypes will be omitted from Link. To indicate that all subtypes should be shown, use the value `\"all\"`. If the `account_filters` filter is used, any account type for which a filter is not specified will be entirely omitted from Link. For a full list of valid types and subtypes, see the [Account schema](https://plaid.com/docs/api/accounts#account-type-schema).  For institutions using OAuth, the filter will not affect the list of accounts shown by the bank in the OAuth window. 
type LinkTokenAccountFilters struct {
	Depository *DepositoryFilter `json:"depository,omitempty"`
	Credit *CreditFilter `json:"credit,omitempty"`
	Loan *LoanFilter `json:"loan,omitempty"`
	Investment *InvestmentFilter `json:"investment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinkTokenAccountFilters LinkTokenAccountFilters

// NewLinkTokenAccountFilters instantiates a new LinkTokenAccountFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenAccountFilters() *LinkTokenAccountFilters {
	this := LinkTokenAccountFilters{}
	return &this
}

// NewLinkTokenAccountFiltersWithDefaults instantiates a new LinkTokenAccountFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenAccountFiltersWithDefaults() *LinkTokenAccountFilters {
	this := LinkTokenAccountFilters{}
	return &this
}

// GetDepository returns the Depository field value if set, zero value otherwise.
func (o *LinkTokenAccountFilters) GetDepository() DepositoryFilter {
	if o == nil || o.Depository == nil {
		var ret DepositoryFilter
		return ret
	}
	return *o.Depository
}

// GetDepositoryOk returns a tuple with the Depository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenAccountFilters) GetDepositoryOk() (*DepositoryFilter, bool) {
	if o == nil || o.Depository == nil {
		return nil, false
	}
	return o.Depository, true
}

// HasDepository returns a boolean if a field has been set.
func (o *LinkTokenAccountFilters) HasDepository() bool {
	if o != nil && o.Depository != nil {
		return true
	}

	return false
}

// SetDepository gets a reference to the given DepositoryFilter and assigns it to the Depository field.
func (o *LinkTokenAccountFilters) SetDepository(v DepositoryFilter) {
	o.Depository = &v
}

// GetCredit returns the Credit field value if set, zero value otherwise.
func (o *LinkTokenAccountFilters) GetCredit() CreditFilter {
	if o == nil || o.Credit == nil {
		var ret CreditFilter
		return ret
	}
	return *o.Credit
}

// GetCreditOk returns a tuple with the Credit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenAccountFilters) GetCreditOk() (*CreditFilter, bool) {
	if o == nil || o.Credit == nil {
		return nil, false
	}
	return o.Credit, true
}

// HasCredit returns a boolean if a field has been set.
func (o *LinkTokenAccountFilters) HasCredit() bool {
	if o != nil && o.Credit != nil {
		return true
	}

	return false
}

// SetCredit gets a reference to the given CreditFilter and assigns it to the Credit field.
func (o *LinkTokenAccountFilters) SetCredit(v CreditFilter) {
	o.Credit = &v
}

// GetLoan returns the Loan field value if set, zero value otherwise.
func (o *LinkTokenAccountFilters) GetLoan() LoanFilter {
	if o == nil || o.Loan == nil {
		var ret LoanFilter
		return ret
	}
	return *o.Loan
}

// GetLoanOk returns a tuple with the Loan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenAccountFilters) GetLoanOk() (*LoanFilter, bool) {
	if o == nil || o.Loan == nil {
		return nil, false
	}
	return o.Loan, true
}

// HasLoan returns a boolean if a field has been set.
func (o *LinkTokenAccountFilters) HasLoan() bool {
	if o != nil && o.Loan != nil {
		return true
	}

	return false
}

// SetLoan gets a reference to the given LoanFilter and assigns it to the Loan field.
func (o *LinkTokenAccountFilters) SetLoan(v LoanFilter) {
	o.Loan = &v
}

// GetInvestment returns the Investment field value if set, zero value otherwise.
func (o *LinkTokenAccountFilters) GetInvestment() InvestmentFilter {
	if o == nil || o.Investment == nil {
		var ret InvestmentFilter
		return ret
	}
	return *o.Investment
}

// GetInvestmentOk returns a tuple with the Investment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenAccountFilters) GetInvestmentOk() (*InvestmentFilter, bool) {
	if o == nil || o.Investment == nil {
		return nil, false
	}
	return o.Investment, true
}

// HasInvestment returns a boolean if a field has been set.
func (o *LinkTokenAccountFilters) HasInvestment() bool {
	if o != nil && o.Investment != nil {
		return true
	}

	return false
}

// SetInvestment gets a reference to the given InvestmentFilter and assigns it to the Investment field.
func (o *LinkTokenAccountFilters) SetInvestment(v InvestmentFilter) {
	o.Investment = &v
}

func (o LinkTokenAccountFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Depository != nil {
		toSerialize["depository"] = o.Depository
	}
	if o.Credit != nil {
		toSerialize["credit"] = o.Credit
	}
	if o.Loan != nil {
		toSerialize["loan"] = o.Loan
	}
	if o.Investment != nil {
		toSerialize["investment"] = o.Investment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkTokenAccountFilters) UnmarshalJSON(bytes []byte) (err error) {
	varLinkTokenAccountFilters := _LinkTokenAccountFilters{}

	if err = json.Unmarshal(bytes, &varLinkTokenAccountFilters); err == nil {
		*o = LinkTokenAccountFilters(varLinkTokenAccountFilters)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "depository")
		delete(additionalProperties, "credit")
		delete(additionalProperties, "loan")
		delete(additionalProperties, "investment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkTokenAccountFilters struct {
	value *LinkTokenAccountFilters
	isSet bool
}

func (v NullableLinkTokenAccountFilters) Get() *LinkTokenAccountFilters {
	return v.value
}

func (v *NullableLinkTokenAccountFilters) Set(val *LinkTokenAccountFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenAccountFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenAccountFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenAccountFilters(val *LinkTokenAccountFilters) *NullableLinkTokenAccountFilters {
	return &NullableLinkTokenAccountFilters{value: val, isSet: true}
}

func (v NullableLinkTokenAccountFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenAccountFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


