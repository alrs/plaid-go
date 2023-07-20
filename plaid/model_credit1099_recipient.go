/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.391.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// Credit1099Recipient An object representing a recipient used in both 1099-K and 1099-MISC tax documents.
type Credit1099Recipient struct {
	Address *CreditPayStubAddress `json:"address,omitempty"`
	// Name of recipient.
	Name NullableString `json:"name,omitempty"`
	// Tax identification number of recipient.
	Tin NullableString `json:"tin,omitempty"`
	// Account number number of recipient.
	AccountNumber NullableString `json:"account_number,omitempty"`
	// Checked if FACTA is a filing requirement.
	FactaFilingRequirement NullableString `json:"facta_filing_requirement,omitempty"`
	// Checked if 2nd TIN exists.
	SecondTinExists NullableString `json:"second_tin_exists,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Credit1099Recipient Credit1099Recipient

// NewCredit1099Recipient instantiates a new Credit1099Recipient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredit1099Recipient() *Credit1099Recipient {
	this := Credit1099Recipient{}
	return &this
}

// NewCredit1099RecipientWithDefaults instantiates a new Credit1099Recipient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredit1099RecipientWithDefaults() *Credit1099Recipient {
	this := Credit1099Recipient{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Credit1099Recipient) GetAddress() CreditPayStubAddress {
	if o == nil || o.Address == nil {
		var ret CreditPayStubAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit1099Recipient) GetAddressOk() (*CreditPayStubAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Credit1099Recipient) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given CreditPayStubAddress and assigns it to the Address field.
func (o *Credit1099Recipient) SetAddress(v CreditPayStubAddress) {
	o.Address = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credit1099Recipient) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credit1099Recipient) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Credit1099Recipient) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Credit1099Recipient) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Credit1099Recipient) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Credit1099Recipient) UnsetName() {
	o.Name.Unset()
}

// GetTin returns the Tin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credit1099Recipient) GetTin() string {
	if o == nil || o.Tin.Get() == nil {
		var ret string
		return ret
	}
	return *o.Tin.Get()
}

// GetTinOk returns a tuple with the Tin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credit1099Recipient) GetTinOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Tin.Get(), o.Tin.IsSet()
}

// HasTin returns a boolean if a field has been set.
func (o *Credit1099Recipient) HasTin() bool {
	if o != nil && o.Tin.IsSet() {
		return true
	}

	return false
}

// SetTin gets a reference to the given NullableString and assigns it to the Tin field.
func (o *Credit1099Recipient) SetTin(v string) {
	o.Tin.Set(&v)
}
// SetTinNil sets the value for Tin to be an explicit nil
func (o *Credit1099Recipient) SetTinNil() {
	o.Tin.Set(nil)
}

// UnsetTin ensures that no value is present for Tin, not even an explicit nil
func (o *Credit1099Recipient) UnsetTin() {
	o.Tin.Unset()
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credit1099Recipient) GetAccountNumber() string {
	if o == nil || o.AccountNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountNumber.Get()
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credit1099Recipient) GetAccountNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountNumber.Get(), o.AccountNumber.IsSet()
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *Credit1099Recipient) HasAccountNumber() bool {
	if o != nil && o.AccountNumber.IsSet() {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given NullableString and assigns it to the AccountNumber field.
func (o *Credit1099Recipient) SetAccountNumber(v string) {
	o.AccountNumber.Set(&v)
}
// SetAccountNumberNil sets the value for AccountNumber to be an explicit nil
func (o *Credit1099Recipient) SetAccountNumberNil() {
	o.AccountNumber.Set(nil)
}

// UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil
func (o *Credit1099Recipient) UnsetAccountNumber() {
	o.AccountNumber.Unset()
}

// GetFactaFilingRequirement returns the FactaFilingRequirement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credit1099Recipient) GetFactaFilingRequirement() string {
	if o == nil || o.FactaFilingRequirement.Get() == nil {
		var ret string
		return ret
	}
	return *o.FactaFilingRequirement.Get()
}

// GetFactaFilingRequirementOk returns a tuple with the FactaFilingRequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credit1099Recipient) GetFactaFilingRequirementOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FactaFilingRequirement.Get(), o.FactaFilingRequirement.IsSet()
}

// HasFactaFilingRequirement returns a boolean if a field has been set.
func (o *Credit1099Recipient) HasFactaFilingRequirement() bool {
	if o != nil && o.FactaFilingRequirement.IsSet() {
		return true
	}

	return false
}

// SetFactaFilingRequirement gets a reference to the given NullableString and assigns it to the FactaFilingRequirement field.
func (o *Credit1099Recipient) SetFactaFilingRequirement(v string) {
	o.FactaFilingRequirement.Set(&v)
}
// SetFactaFilingRequirementNil sets the value for FactaFilingRequirement to be an explicit nil
func (o *Credit1099Recipient) SetFactaFilingRequirementNil() {
	o.FactaFilingRequirement.Set(nil)
}

// UnsetFactaFilingRequirement ensures that no value is present for FactaFilingRequirement, not even an explicit nil
func (o *Credit1099Recipient) UnsetFactaFilingRequirement() {
	o.FactaFilingRequirement.Unset()
}

// GetSecondTinExists returns the SecondTinExists field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credit1099Recipient) GetSecondTinExists() string {
	if o == nil || o.SecondTinExists.Get() == nil {
		var ret string
		return ret
	}
	return *o.SecondTinExists.Get()
}

// GetSecondTinExistsOk returns a tuple with the SecondTinExists field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credit1099Recipient) GetSecondTinExistsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SecondTinExists.Get(), o.SecondTinExists.IsSet()
}

// HasSecondTinExists returns a boolean if a field has been set.
func (o *Credit1099Recipient) HasSecondTinExists() bool {
	if o != nil && o.SecondTinExists.IsSet() {
		return true
	}

	return false
}

// SetSecondTinExists gets a reference to the given NullableString and assigns it to the SecondTinExists field.
func (o *Credit1099Recipient) SetSecondTinExists(v string) {
	o.SecondTinExists.Set(&v)
}
// SetSecondTinExistsNil sets the value for SecondTinExists to be an explicit nil
func (o *Credit1099Recipient) SetSecondTinExistsNil() {
	o.SecondTinExists.Set(nil)
}

// UnsetSecondTinExists ensures that no value is present for SecondTinExists, not even an explicit nil
func (o *Credit1099Recipient) UnsetSecondTinExists() {
	o.SecondTinExists.Unset()
}

func (o Credit1099Recipient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Tin.IsSet() {
		toSerialize["tin"] = o.Tin.Get()
	}
	if o.AccountNumber.IsSet() {
		toSerialize["account_number"] = o.AccountNumber.Get()
	}
	if o.FactaFilingRequirement.IsSet() {
		toSerialize["facta_filing_requirement"] = o.FactaFilingRequirement.Get()
	}
	if o.SecondTinExists.IsSet() {
		toSerialize["second_tin_exists"] = o.SecondTinExists.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Credit1099Recipient) UnmarshalJSON(bytes []byte) (err error) {
	varCredit1099Recipient := _Credit1099Recipient{}

	if err = json.Unmarshal(bytes, &varCredit1099Recipient); err == nil {
		*o = Credit1099Recipient(varCredit1099Recipient)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "name")
		delete(additionalProperties, "tin")
		delete(additionalProperties, "account_number")
		delete(additionalProperties, "facta_filing_requirement")
		delete(additionalProperties, "second_tin_exists")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCredit1099Recipient struct {
	value *Credit1099Recipient
	isSet bool
}

func (v NullableCredit1099Recipient) Get() *Credit1099Recipient {
	return v.value
}

func (v *NullableCredit1099Recipient) Set(val *Credit1099Recipient) {
	v.value = val
	v.isSet = true
}

func (v NullableCredit1099Recipient) IsSet() bool {
	return v.isSet
}

func (v *NullableCredit1099Recipient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredit1099Recipient(val *Credit1099Recipient) *NullableCredit1099Recipient {
	return &NullableCredit1099Recipient{value: val, isSet: true}
}

func (v NullableCredit1099Recipient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredit1099Recipient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


