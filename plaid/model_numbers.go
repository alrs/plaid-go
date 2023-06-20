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

// Numbers Account and bank identifier number data used to configure the test account. All values are optional.
type Numbers struct {
	// Will be used for the account number.
	Account *string `json:"account,omitempty"`
	// Must be a valid ACH routing number.
	AchRouting *string `json:"ach_routing,omitempty"`
	// Must be a valid wire transfer routing number.
	AchWireRouting *string `json:"ach_wire_routing,omitempty"`
	// EFT institution number. Must be specified alongside `eft_branch`.
	EftInstitution *string `json:"eft_institution,omitempty"`
	// EFT branch number. Must be specified alongside `eft_institution`.
	EftBranch *string `json:"eft_branch,omitempty"`
	// Bank identifier code (BIC). Must be specified alongside `international_iban`.
	InternationalBic *string `json:"international_bic,omitempty"`
	// International bank account number (IBAN). If no account number is specified via `account`, will also be used as the account number by default. Must be specified alongside `international_bic`.
	InternationalIban *string `json:"international_iban,omitempty"`
	// BACS sort code
	BacsSortCode *string `json:"bacs_sort_code,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Numbers Numbers

// NewNumbers instantiates a new Numbers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumbers() *Numbers {
	this := Numbers{}
	return &this
}

// NewNumbersWithDefaults instantiates a new Numbers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumbersWithDefaults() *Numbers {
	this := Numbers{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *Numbers) GetAccount() string {
	if o == nil || o.Account == nil {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Numbers) GetAccountOk() (*string, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *Numbers) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *Numbers) SetAccount(v string) {
	o.Account = &v
}

// GetAchRouting returns the AchRouting field value if set, zero value otherwise.
func (o *Numbers) GetAchRouting() string {
	if o == nil || o.AchRouting == nil {
		var ret string
		return ret
	}
	return *o.AchRouting
}

// GetAchRoutingOk returns a tuple with the AchRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Numbers) GetAchRoutingOk() (*string, bool) {
	if o == nil || o.AchRouting == nil {
		return nil, false
	}
	return o.AchRouting, true
}

// HasAchRouting returns a boolean if a field has been set.
func (o *Numbers) HasAchRouting() bool {
	if o != nil && o.AchRouting != nil {
		return true
	}

	return false
}

// SetAchRouting gets a reference to the given string and assigns it to the AchRouting field.
func (o *Numbers) SetAchRouting(v string) {
	o.AchRouting = &v
}

// GetAchWireRouting returns the AchWireRouting field value if set, zero value otherwise.
func (o *Numbers) GetAchWireRouting() string {
	if o == nil || o.AchWireRouting == nil {
		var ret string
		return ret
	}
	return *o.AchWireRouting
}

// GetAchWireRoutingOk returns a tuple with the AchWireRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Numbers) GetAchWireRoutingOk() (*string, bool) {
	if o == nil || o.AchWireRouting == nil {
		return nil, false
	}
	return o.AchWireRouting, true
}

// HasAchWireRouting returns a boolean if a field has been set.
func (o *Numbers) HasAchWireRouting() bool {
	if o != nil && o.AchWireRouting != nil {
		return true
	}

	return false
}

// SetAchWireRouting gets a reference to the given string and assigns it to the AchWireRouting field.
func (o *Numbers) SetAchWireRouting(v string) {
	o.AchWireRouting = &v
}

// GetEftInstitution returns the EftInstitution field value if set, zero value otherwise.
func (o *Numbers) GetEftInstitution() string {
	if o == nil || o.EftInstitution == nil {
		var ret string
		return ret
	}
	return *o.EftInstitution
}

// GetEftInstitutionOk returns a tuple with the EftInstitution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Numbers) GetEftInstitutionOk() (*string, bool) {
	if o == nil || o.EftInstitution == nil {
		return nil, false
	}
	return o.EftInstitution, true
}

// HasEftInstitution returns a boolean if a field has been set.
func (o *Numbers) HasEftInstitution() bool {
	if o != nil && o.EftInstitution != nil {
		return true
	}

	return false
}

// SetEftInstitution gets a reference to the given string and assigns it to the EftInstitution field.
func (o *Numbers) SetEftInstitution(v string) {
	o.EftInstitution = &v
}

// GetEftBranch returns the EftBranch field value if set, zero value otherwise.
func (o *Numbers) GetEftBranch() string {
	if o == nil || o.EftBranch == nil {
		var ret string
		return ret
	}
	return *o.EftBranch
}

// GetEftBranchOk returns a tuple with the EftBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Numbers) GetEftBranchOk() (*string, bool) {
	if o == nil || o.EftBranch == nil {
		return nil, false
	}
	return o.EftBranch, true
}

// HasEftBranch returns a boolean if a field has been set.
func (o *Numbers) HasEftBranch() bool {
	if o != nil && o.EftBranch != nil {
		return true
	}

	return false
}

// SetEftBranch gets a reference to the given string and assigns it to the EftBranch field.
func (o *Numbers) SetEftBranch(v string) {
	o.EftBranch = &v
}

// GetInternationalBic returns the InternationalBic field value if set, zero value otherwise.
func (o *Numbers) GetInternationalBic() string {
	if o == nil || o.InternationalBic == nil {
		var ret string
		return ret
	}
	return *o.InternationalBic
}

// GetInternationalBicOk returns a tuple with the InternationalBic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Numbers) GetInternationalBicOk() (*string, bool) {
	if o == nil || o.InternationalBic == nil {
		return nil, false
	}
	return o.InternationalBic, true
}

// HasInternationalBic returns a boolean if a field has been set.
func (o *Numbers) HasInternationalBic() bool {
	if o != nil && o.InternationalBic != nil {
		return true
	}

	return false
}

// SetInternationalBic gets a reference to the given string and assigns it to the InternationalBic field.
func (o *Numbers) SetInternationalBic(v string) {
	o.InternationalBic = &v
}

// GetInternationalIban returns the InternationalIban field value if set, zero value otherwise.
func (o *Numbers) GetInternationalIban() string {
	if o == nil || o.InternationalIban == nil {
		var ret string
		return ret
	}
	return *o.InternationalIban
}

// GetInternationalIbanOk returns a tuple with the InternationalIban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Numbers) GetInternationalIbanOk() (*string, bool) {
	if o == nil || o.InternationalIban == nil {
		return nil, false
	}
	return o.InternationalIban, true
}

// HasInternationalIban returns a boolean if a field has been set.
func (o *Numbers) HasInternationalIban() bool {
	if o != nil && o.InternationalIban != nil {
		return true
	}

	return false
}

// SetInternationalIban gets a reference to the given string and assigns it to the InternationalIban field.
func (o *Numbers) SetInternationalIban(v string) {
	o.InternationalIban = &v
}

// GetBacsSortCode returns the BacsSortCode field value if set, zero value otherwise.
func (o *Numbers) GetBacsSortCode() string {
	if o == nil || o.BacsSortCode == nil {
		var ret string
		return ret
	}
	return *o.BacsSortCode
}

// GetBacsSortCodeOk returns a tuple with the BacsSortCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Numbers) GetBacsSortCodeOk() (*string, bool) {
	if o == nil || o.BacsSortCode == nil {
		return nil, false
	}
	return o.BacsSortCode, true
}

// HasBacsSortCode returns a boolean if a field has been set.
func (o *Numbers) HasBacsSortCode() bool {
	if o != nil && o.BacsSortCode != nil {
		return true
	}

	return false
}

// SetBacsSortCode gets a reference to the given string and assigns it to the BacsSortCode field.
func (o *Numbers) SetBacsSortCode(v string) {
	o.BacsSortCode = &v
}

func (o Numbers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.AchRouting != nil {
		toSerialize["ach_routing"] = o.AchRouting
	}
	if o.AchWireRouting != nil {
		toSerialize["ach_wire_routing"] = o.AchWireRouting
	}
	if o.EftInstitution != nil {
		toSerialize["eft_institution"] = o.EftInstitution
	}
	if o.EftBranch != nil {
		toSerialize["eft_branch"] = o.EftBranch
	}
	if o.InternationalBic != nil {
		toSerialize["international_bic"] = o.InternationalBic
	}
	if o.InternationalIban != nil {
		toSerialize["international_iban"] = o.InternationalIban
	}
	if o.BacsSortCode != nil {
		toSerialize["bacs_sort_code"] = o.BacsSortCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Numbers) UnmarshalJSON(bytes []byte) (err error) {
	varNumbers := _Numbers{}

	if err = json.Unmarshal(bytes, &varNumbers); err == nil {
		*o = Numbers(varNumbers)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account")
		delete(additionalProperties, "ach_routing")
		delete(additionalProperties, "ach_wire_routing")
		delete(additionalProperties, "eft_institution")
		delete(additionalProperties, "eft_branch")
		delete(additionalProperties, "international_bic")
		delete(additionalProperties, "international_iban")
		delete(additionalProperties, "bacs_sort_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNumbers struct {
	value *Numbers
	isSet bool
}

func (v NullableNumbers) Get() *Numbers {
	return v.value
}

func (v *NullableNumbers) Set(val *Numbers) {
	v.value = val
	v.isSet = true
}

func (v NullableNumbers) IsSet() bool {
	return v.isSet
}

func (v *NullableNumbers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumbers(val *Numbers) *NullableNumbers {
	return &NullableNumbers{value: val, isSet: true}
}

func (v NullableNumbers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumbers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


