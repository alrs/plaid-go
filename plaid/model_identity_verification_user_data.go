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

// IdentityVerificationUserData The identity data that was either collected from the user or provided via API in order to perform an identity verification.
type IdentityVerificationUserData struct {
	// A phone number in E.164 format.
	PhoneNumber NullableString `json:"phone_number,omitempty"`
	// A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).
	DateOfBirth NullableString `json:"date_of_birth"`
	// An IPv4 or IPV6 address.
	IpAddress NullableString `json:"ip_address"`
	// A valid email address.
	EmailAddress NullableString `json:"email_address"`
	Name NullableUserName `json:"name"`
	Address NullableIdentityVerificationUserAddress `json:"address"`
	IdNumber NullableUserIDNumber `json:"id_number"`
}

// NewIdentityVerificationUserData instantiates a new IdentityVerificationUserData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityVerificationUserData(dateOfBirth NullableString, ipAddress NullableString, emailAddress NullableString, name NullableUserName, address NullableIdentityVerificationUserAddress, idNumber NullableUserIDNumber) *IdentityVerificationUserData {
	this := IdentityVerificationUserData{}
	this.DateOfBirth = dateOfBirth
	this.IpAddress = ipAddress
	this.EmailAddress = emailAddress
	this.Name = name
	this.Address = address
	this.IdNumber = idNumber
	return &this
}

// NewIdentityVerificationUserDataWithDefaults instantiates a new IdentityVerificationUserData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityVerificationUserDataWithDefaults() *IdentityVerificationUserData {
	this := IdentityVerificationUserData{}
	return &this
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityVerificationUserData) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationUserData) GetPhoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *IdentityVerificationUserData) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullableString and assigns it to the PhoneNumber field.
func (o *IdentityVerificationUserData) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}
// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *IdentityVerificationUserData) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *IdentityVerificationUserData) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

// GetDateOfBirth returns the DateOfBirth field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IdentityVerificationUserData) GetDateOfBirth() string {
	if o == nil || o.DateOfBirth.Get() == nil {
		var ret string
		return ret
	}

	return *o.DateOfBirth.Get()
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationUserData) GetDateOfBirthOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DateOfBirth.Get(), o.DateOfBirth.IsSet()
}

// SetDateOfBirth sets field value
func (o *IdentityVerificationUserData) SetDateOfBirth(v string) {
	o.DateOfBirth.Set(&v)
}

// GetIpAddress returns the IpAddress field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IdentityVerificationUserData) GetIpAddress() string {
	if o == nil || o.IpAddress.Get() == nil {
		var ret string
		return ret
	}

	return *o.IpAddress.Get()
}

// GetIpAddressOk returns a tuple with the IpAddress field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationUserData) GetIpAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IpAddress.Get(), o.IpAddress.IsSet()
}

// SetIpAddress sets field value
func (o *IdentityVerificationUserData) SetIpAddress(v string) {
	o.IpAddress.Set(&v)
}

// GetEmailAddress returns the EmailAddress field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IdentityVerificationUserData) GetEmailAddress() string {
	if o == nil || o.EmailAddress.Get() == nil {
		var ret string
		return ret
	}

	return *o.EmailAddress.Get()
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationUserData) GetEmailAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmailAddress.Get(), o.EmailAddress.IsSet()
}

// SetEmailAddress sets field value
func (o *IdentityVerificationUserData) SetEmailAddress(v string) {
	o.EmailAddress.Set(&v)
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for UserName will be returned
func (o *IdentityVerificationUserData) GetName() UserName {
	if o == nil || o.Name.Get() == nil {
		var ret UserName
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationUserData) GetNameOk() (*UserName, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *IdentityVerificationUserData) SetName(v UserName) {
	o.Name.Set(&v)
}

// GetAddress returns the Address field value
// If the value is explicit nil, the zero value for IdentityVerificationUserAddress will be returned
func (o *IdentityVerificationUserData) GetAddress() IdentityVerificationUserAddress {
	if o == nil || o.Address.Get() == nil {
		var ret IdentityVerificationUserAddress
		return ret
	}

	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationUserData) GetAddressOk() (*IdentityVerificationUserAddress, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// SetAddress sets field value
func (o *IdentityVerificationUserData) SetAddress(v IdentityVerificationUserAddress) {
	o.Address.Set(&v)
}

// GetIdNumber returns the IdNumber field value
// If the value is explicit nil, the zero value for UserIDNumber will be returned
func (o *IdentityVerificationUserData) GetIdNumber() UserIDNumber {
	if o == nil || o.IdNumber.Get() == nil {
		var ret UserIDNumber
		return ret
	}

	return *o.IdNumber.Get()
}

// GetIdNumberOk returns a tuple with the IdNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationUserData) GetIdNumberOk() (*UserIDNumber, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IdNumber.Get(), o.IdNumber.IsSet()
}

// SetIdNumber sets field value
func (o *IdentityVerificationUserData) SetIdNumber(v UserIDNumber) {
	o.IdNumber.Set(&v)
}

func (o IdentityVerificationUserData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PhoneNumber.IsSet() {
		toSerialize["phone_number"] = o.PhoneNumber.Get()
	}
	if true {
		toSerialize["date_of_birth"] = o.DateOfBirth.Get()
	}
	if true {
		toSerialize["ip_address"] = o.IpAddress.Get()
	}
	if true {
		toSerialize["email_address"] = o.EmailAddress.Get()
	}
	if true {
		toSerialize["name"] = o.Name.Get()
	}
	if true {
		toSerialize["address"] = o.Address.Get()
	}
	if true {
		toSerialize["id_number"] = o.IdNumber.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityVerificationUserData struct {
	value *IdentityVerificationUserData
	isSet bool
}

func (v NullableIdentityVerificationUserData) Get() *IdentityVerificationUserData {
	return v.value
}

func (v *NullableIdentityVerificationUserData) Set(val *IdentityVerificationUserData) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityVerificationUserData) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityVerificationUserData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityVerificationUserData(val *IdentityVerificationUserData) *NullableIdentityVerificationUserData {
	return &NullableIdentityVerificationUserData{value: val, isSet: true}
}

func (v NullableIdentityVerificationUserData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityVerificationUserData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


