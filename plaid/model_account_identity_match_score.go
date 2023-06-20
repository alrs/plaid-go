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
)

// AccountIdentityMatchScore Identity match scores for an account
type AccountIdentityMatchScore struct {
	// Plaid’s unique identifier for the account. This value will not change unless Plaid can't reconcile the account with the data returned by the financial institution. This may occur, for example, when the name of the account changes. If this happens a new `account_id` will be assigned to the account.  The `account_id` can also change if the `access_token` is deleted and the same credentials that were used to generate that `access_token` are used to generate a new `access_token` on a later date. In that case, the new `account_id` will be different from the old `account_id`.  If an account with a specific `account_id` disappears instead of changing, the account is likely closed. Closed accounts are not returned by the Plaid API.  Like all Plaid identifiers, the `account_id` is case sensitive.
	AccountId string `json:"account_id"`
	Balances AccountBalance `json:"balances"`
	// The last 2-4 alphanumeric characters of an account's official account number. Note that the mask may be non-unique between an Item's accounts, and it may also not match the mask that the bank displays to the user.
	Mask NullableString `json:"mask"`
	// The name of the account, either assigned by the user or by the financial institution itself
	Name string `json:"name"`
	// The official name of the account as given by the financial institution
	OfficialName NullableString `json:"official_name"`
	Type AccountType `json:"type"`
	Subtype NullableAccountSubtype `json:"subtype"`
	// The current verification status of an Auth Item initiated through Automated or Manual micro-deposits.  Returned for Auth Items only.  `pending_automatic_verification`: The Item is pending automatic verification  `pending_manual_verification`: The Item is pending manual micro-deposit verification. Items remain in this state until the user successfully verifies the two amounts.  `automatically_verified`: The Item has successfully been automatically verified   `manually_verified`: The Item has successfully been manually verified  `verification_expired`: Plaid was unable to automatically verify the deposit within 7 calendar days and will no longer attempt to validate the Item. Users may retry by submitting their information again through Link.  `verification_failed`: The Item failed manual micro-deposit verification because the user exhausted all 3 verification attempts. Users may retry by submitting their information again through Link.   
	VerificationStatus *string `json:"verification_status,omitempty"`
	// A unique and persistent identifier for accounts that can be used to trace multiple instances of the same account across different Items for depository accounts. This is currently an opt-in field and only supported for Chase Items.
	PersistentAccountId *string `json:"persistent_account_id,omitempty"`
	LegalName NullableNameMatchScore `json:"legal_name,omitempty"`
	PhoneNumber NullablePhoneNumberMatchScore `json:"phone_number,omitempty"`
	EmailAddress NullableEmailAddressMatchScore `json:"email_address,omitempty"`
	Address NullableAddressMatchScore `json:"address,omitempty"`
}

// NewAccountIdentityMatchScore instantiates a new AccountIdentityMatchScore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountIdentityMatchScore(accountId string, balances AccountBalance, mask NullableString, name string, officialName NullableString, type_ AccountType, subtype NullableAccountSubtype) *AccountIdentityMatchScore {
	this := AccountIdentityMatchScore{}
	this.AccountId = accountId
	this.Balances = balances
	this.Mask = mask
	this.Name = name
	this.OfficialName = officialName
	this.Type = type_
	this.Subtype = subtype
	return &this
}

// NewAccountIdentityMatchScoreWithDefaults instantiates a new AccountIdentityMatchScore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountIdentityMatchScoreWithDefaults() *AccountIdentityMatchScore {
	this := AccountIdentityMatchScore{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *AccountIdentityMatchScore) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AccountIdentityMatchScore) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AccountIdentityMatchScore) SetAccountId(v string) {
	o.AccountId = v
}

// GetBalances returns the Balances field value
func (o *AccountIdentityMatchScore) GetBalances() AccountBalance {
	if o == nil {
		var ret AccountBalance
		return ret
	}

	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value
// and a boolean to check if the value has been set.
func (o *AccountIdentityMatchScore) GetBalancesOk() (*AccountBalance, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Balances, true
}

// SetBalances sets field value
func (o *AccountIdentityMatchScore) SetBalances(v AccountBalance) {
	o.Balances = v
}

// GetMask returns the Mask field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AccountIdentityMatchScore) GetMask() string {
	if o == nil || o.Mask.Get() == nil {
		var ret string
		return ret
	}

	return *o.Mask.Get()
}

// GetMaskOk returns a tuple with the Mask field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountIdentityMatchScore) GetMaskOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Mask.Get(), o.Mask.IsSet()
}

// SetMask sets field value
func (o *AccountIdentityMatchScore) SetMask(v string) {
	o.Mask.Set(&v)
}

// GetName returns the Name field value
func (o *AccountIdentityMatchScore) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountIdentityMatchScore) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountIdentityMatchScore) SetName(v string) {
	o.Name = v
}

// GetOfficialName returns the OfficialName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AccountIdentityMatchScore) GetOfficialName() string {
	if o == nil || o.OfficialName.Get() == nil {
		var ret string
		return ret
	}

	return *o.OfficialName.Get()
}

// GetOfficialNameOk returns a tuple with the OfficialName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountIdentityMatchScore) GetOfficialNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OfficialName.Get(), o.OfficialName.IsSet()
}

// SetOfficialName sets field value
func (o *AccountIdentityMatchScore) SetOfficialName(v string) {
	o.OfficialName.Set(&v)
}

// GetType returns the Type field value
func (o *AccountIdentityMatchScore) GetType() AccountType {
	if o == nil {
		var ret AccountType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccountIdentityMatchScore) GetTypeOk() (*AccountType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccountIdentityMatchScore) SetType(v AccountType) {
	o.Type = v
}

// GetSubtype returns the Subtype field value
// If the value is explicit nil, the zero value for AccountSubtype will be returned
func (o *AccountIdentityMatchScore) GetSubtype() AccountSubtype {
	if o == nil || o.Subtype.Get() == nil {
		var ret AccountSubtype
		return ret
	}

	return *o.Subtype.Get()
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountIdentityMatchScore) GetSubtypeOk() (*AccountSubtype, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Subtype.Get(), o.Subtype.IsSet()
}

// SetSubtype sets field value
func (o *AccountIdentityMatchScore) SetSubtype(v AccountSubtype) {
	o.Subtype.Set(&v)
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *AccountIdentityMatchScore) GetVerificationStatus() string {
	if o == nil || o.VerificationStatus == nil {
		var ret string
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountIdentityMatchScore) GetVerificationStatusOk() (*string, bool) {
	if o == nil || o.VerificationStatus == nil {
		return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *AccountIdentityMatchScore) HasVerificationStatus() bool {
	if o != nil && o.VerificationStatus != nil {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given string and assigns it to the VerificationStatus field.
func (o *AccountIdentityMatchScore) SetVerificationStatus(v string) {
	o.VerificationStatus = &v
}

// GetPersistentAccountId returns the PersistentAccountId field value if set, zero value otherwise.
func (o *AccountIdentityMatchScore) GetPersistentAccountId() string {
	if o == nil || o.PersistentAccountId == nil {
		var ret string
		return ret
	}
	return *o.PersistentAccountId
}

// GetPersistentAccountIdOk returns a tuple with the PersistentAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountIdentityMatchScore) GetPersistentAccountIdOk() (*string, bool) {
	if o == nil || o.PersistentAccountId == nil {
		return nil, false
	}
	return o.PersistentAccountId, true
}

// HasPersistentAccountId returns a boolean if a field has been set.
func (o *AccountIdentityMatchScore) HasPersistentAccountId() bool {
	if o != nil && o.PersistentAccountId != nil {
		return true
	}

	return false
}

// SetPersistentAccountId gets a reference to the given string and assigns it to the PersistentAccountId field.
func (o *AccountIdentityMatchScore) SetPersistentAccountId(v string) {
	o.PersistentAccountId = &v
}

// GetLegalName returns the LegalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountIdentityMatchScore) GetLegalName() NameMatchScore {
	if o == nil || o.LegalName.Get() == nil {
		var ret NameMatchScore
		return ret
	}
	return *o.LegalName.Get()
}

// GetLegalNameOk returns a tuple with the LegalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountIdentityMatchScore) GetLegalNameOk() (*NameMatchScore, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LegalName.Get(), o.LegalName.IsSet()
}

// HasLegalName returns a boolean if a field has been set.
func (o *AccountIdentityMatchScore) HasLegalName() bool {
	if o != nil && o.LegalName.IsSet() {
		return true
	}

	return false
}

// SetLegalName gets a reference to the given NullableNameMatchScore and assigns it to the LegalName field.
func (o *AccountIdentityMatchScore) SetLegalName(v NameMatchScore) {
	o.LegalName.Set(&v)
}
// SetLegalNameNil sets the value for LegalName to be an explicit nil
func (o *AccountIdentityMatchScore) SetLegalNameNil() {
	o.LegalName.Set(nil)
}

// UnsetLegalName ensures that no value is present for LegalName, not even an explicit nil
func (o *AccountIdentityMatchScore) UnsetLegalName() {
	o.LegalName.Unset()
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountIdentityMatchScore) GetPhoneNumber() PhoneNumberMatchScore {
	if o == nil || o.PhoneNumber.Get() == nil {
		var ret PhoneNumberMatchScore
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountIdentityMatchScore) GetPhoneNumberOk() (*PhoneNumberMatchScore, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *AccountIdentityMatchScore) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullablePhoneNumberMatchScore and assigns it to the PhoneNumber field.
func (o *AccountIdentityMatchScore) SetPhoneNumber(v PhoneNumberMatchScore) {
	o.PhoneNumber.Set(&v)
}
// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *AccountIdentityMatchScore) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *AccountIdentityMatchScore) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountIdentityMatchScore) GetEmailAddress() EmailAddressMatchScore {
	if o == nil || o.EmailAddress.Get() == nil {
		var ret EmailAddressMatchScore
		return ret
	}
	return *o.EmailAddress.Get()
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountIdentityMatchScore) GetEmailAddressOk() (*EmailAddressMatchScore, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmailAddress.Get(), o.EmailAddress.IsSet()
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *AccountIdentityMatchScore) HasEmailAddress() bool {
	if o != nil && o.EmailAddress.IsSet() {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given NullableEmailAddressMatchScore and assigns it to the EmailAddress field.
func (o *AccountIdentityMatchScore) SetEmailAddress(v EmailAddressMatchScore) {
	o.EmailAddress.Set(&v)
}
// SetEmailAddressNil sets the value for EmailAddress to be an explicit nil
func (o *AccountIdentityMatchScore) SetEmailAddressNil() {
	o.EmailAddress.Set(nil)
}

// UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
func (o *AccountIdentityMatchScore) UnsetEmailAddress() {
	o.EmailAddress.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountIdentityMatchScore) GetAddress() AddressMatchScore {
	if o == nil || o.Address.Get() == nil {
		var ret AddressMatchScore
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountIdentityMatchScore) GetAddressOk() (*AddressMatchScore, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *AccountIdentityMatchScore) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableAddressMatchScore and assigns it to the Address field.
func (o *AccountIdentityMatchScore) SetAddress(v AddressMatchScore) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *AccountIdentityMatchScore) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *AccountIdentityMatchScore) UnsetAddress() {
	o.Address.Unset()
}

func (o AccountIdentityMatchScore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["balances"] = o.Balances
	}
	if true {
		toSerialize["mask"] = o.Mask.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["official_name"] = o.OfficialName.Get()
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["subtype"] = o.Subtype.Get()
	}
	if o.VerificationStatus != nil {
		toSerialize["verification_status"] = o.VerificationStatus
	}
	if o.PersistentAccountId != nil {
		toSerialize["persistent_account_id"] = o.PersistentAccountId
	}
	if o.LegalName.IsSet() {
		toSerialize["legal_name"] = o.LegalName.Get()
	}
	if o.PhoneNumber.IsSet() {
		toSerialize["phone_number"] = o.PhoneNumber.Get()
	}
	if o.EmailAddress.IsSet() {
		toSerialize["email_address"] = o.EmailAddress.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAccountIdentityMatchScore struct {
	value *AccountIdentityMatchScore
	isSet bool
}

func (v NullableAccountIdentityMatchScore) Get() *AccountIdentityMatchScore {
	return v.value
}

func (v *NullableAccountIdentityMatchScore) Set(val *AccountIdentityMatchScore) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountIdentityMatchScore) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountIdentityMatchScore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountIdentityMatchScore(val *AccountIdentityMatchScore) *NullableAccountIdentityMatchScore {
	return &NullableAccountIdentityMatchScore{value: val, isSet: true}
}

func (v NullableAccountIdentityMatchScore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountIdentityMatchScore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


