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

// IdentityVerificationCreateRequestUser User information collected outside of Link, most likely via your own onboarding process.  Each of the following identity fields are optional:  `email_address`  `phone_number`  `date_of_birth`  `name`  `address`  `id_number`  Specifically, these fields are optional in that they can either be fully provided (satisfying every required field in their subschema) or omitted from the request entirely by not providing the key or value. Providing these fields via the API will result in Link skipping the data collection process for the associated user. All verification steps enabled in the associated Identity Verification Template will still be run. Verification steps will either be run immediately, or once the user completes the `accept_tos` step, depending on the value provided to the `gave_consent` field. If you are not using the shareable URL feature, you can optionally provide these fields via `/link/token/create` instead; both `/identity_verification/create` and `/link/token/create` are valid ways to provide this information. Note that if you provide a non-`null` user data object via `/identity_verification/create`, any user data fields entered via `/link/token/create` for the same `client_user_id` will be ignored when prefilling Link.
type IdentityVerificationCreateRequestUser struct {
	// A valid email address.
	EmailAddress *string `json:"email_address,omitempty"`
	// A phone number in E.164 format.
	PhoneNumber NullableString `json:"phone_number,omitempty"`
	// A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).
	DateOfBirth *string `json:"date_of_birth,omitempty"`
	Name NullableIdentityVerificationRequestUserName `json:"name,omitempty"`
	Address NullableUserAddress `json:"address,omitempty"`
	IdNumber NullableUserIDNumber `json:"id_number,omitempty"`
	// Specifying `user.client_user_id` is deprecated. Please provide `client_user_id` at the root level instead.
	ClientUserId NullableString `json:"client_user_id,omitempty"`
}

// NewIdentityVerificationCreateRequestUser instantiates a new IdentityVerificationCreateRequestUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityVerificationCreateRequestUser() *IdentityVerificationCreateRequestUser {
	this := IdentityVerificationCreateRequestUser{}
	return &this
}

// NewIdentityVerificationCreateRequestUserWithDefaults instantiates a new IdentityVerificationCreateRequestUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityVerificationCreateRequestUserWithDefaults() *IdentityVerificationCreateRequestUser {
	this := IdentityVerificationCreateRequestUser{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *IdentityVerificationCreateRequestUser) GetEmailAddress() string {
	if o == nil || o.EmailAddress == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityVerificationCreateRequestUser) GetEmailAddressOk() (*string, bool) {
	if o == nil || o.EmailAddress == nil {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *IdentityVerificationCreateRequestUser) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *IdentityVerificationCreateRequestUser) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityVerificationCreateRequestUser) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationCreateRequestUser) GetPhoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *IdentityVerificationCreateRequestUser) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullableString and assigns it to the PhoneNumber field.
func (o *IdentityVerificationCreateRequestUser) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}
// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *IdentityVerificationCreateRequestUser) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *IdentityVerificationCreateRequestUser) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise.
func (o *IdentityVerificationCreateRequestUser) GetDateOfBirth() string {
	if o == nil || o.DateOfBirth == nil {
		var ret string
		return ret
	}
	return *o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityVerificationCreateRequestUser) GetDateOfBirthOk() (*string, bool) {
	if o == nil || o.DateOfBirth == nil {
		return nil, false
	}
	return o.DateOfBirth, true
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *IdentityVerificationCreateRequestUser) HasDateOfBirth() bool {
	if o != nil && o.DateOfBirth != nil {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given string and assigns it to the DateOfBirth field.
func (o *IdentityVerificationCreateRequestUser) SetDateOfBirth(v string) {
	o.DateOfBirth = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityVerificationCreateRequestUser) GetName() IdentityVerificationRequestUserName {
	if o == nil || o.Name.Get() == nil {
		var ret IdentityVerificationRequestUserName
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationCreateRequestUser) GetNameOk() (*IdentityVerificationRequestUserName, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *IdentityVerificationCreateRequestUser) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableIdentityVerificationRequestUserName and assigns it to the Name field.
func (o *IdentityVerificationCreateRequestUser) SetName(v IdentityVerificationRequestUserName) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *IdentityVerificationCreateRequestUser) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *IdentityVerificationCreateRequestUser) UnsetName() {
	o.Name.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityVerificationCreateRequestUser) GetAddress() UserAddress {
	if o == nil || o.Address.Get() == nil {
		var ret UserAddress
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationCreateRequestUser) GetAddressOk() (*UserAddress, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *IdentityVerificationCreateRequestUser) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableUserAddress and assigns it to the Address field.
func (o *IdentityVerificationCreateRequestUser) SetAddress(v UserAddress) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *IdentityVerificationCreateRequestUser) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *IdentityVerificationCreateRequestUser) UnsetAddress() {
	o.Address.Unset()
}

// GetIdNumber returns the IdNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityVerificationCreateRequestUser) GetIdNumber() UserIDNumber {
	if o == nil || o.IdNumber.Get() == nil {
		var ret UserIDNumber
		return ret
	}
	return *o.IdNumber.Get()
}

// GetIdNumberOk returns a tuple with the IdNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationCreateRequestUser) GetIdNumberOk() (*UserIDNumber, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IdNumber.Get(), o.IdNumber.IsSet()
}

// HasIdNumber returns a boolean if a field has been set.
func (o *IdentityVerificationCreateRequestUser) HasIdNumber() bool {
	if o != nil && o.IdNumber.IsSet() {
		return true
	}

	return false
}

// SetIdNumber gets a reference to the given NullableUserIDNumber and assigns it to the IdNumber field.
func (o *IdentityVerificationCreateRequestUser) SetIdNumber(v UserIDNumber) {
	o.IdNumber.Set(&v)
}
// SetIdNumberNil sets the value for IdNumber to be an explicit nil
func (o *IdentityVerificationCreateRequestUser) SetIdNumberNil() {
	o.IdNumber.Set(nil)
}

// UnsetIdNumber ensures that no value is present for IdNumber, not even an explicit nil
func (o *IdentityVerificationCreateRequestUser) UnsetIdNumber() {
	o.IdNumber.Unset()
}

// GetClientUserId returns the ClientUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityVerificationCreateRequestUser) GetClientUserId() string {
	if o == nil || o.ClientUserId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientUserId.Get()
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationCreateRequestUser) GetClientUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientUserId.Get(), o.ClientUserId.IsSet()
}

// HasClientUserId returns a boolean if a field has been set.
func (o *IdentityVerificationCreateRequestUser) HasClientUserId() bool {
	if o != nil && o.ClientUserId.IsSet() {
		return true
	}

	return false
}

// SetClientUserId gets a reference to the given NullableString and assigns it to the ClientUserId field.
func (o *IdentityVerificationCreateRequestUser) SetClientUserId(v string) {
	o.ClientUserId.Set(&v)
}
// SetClientUserIdNil sets the value for ClientUserId to be an explicit nil
func (o *IdentityVerificationCreateRequestUser) SetClientUserIdNil() {
	o.ClientUserId.Set(nil)
}

// UnsetClientUserId ensures that no value is present for ClientUserId, not even an explicit nil
func (o *IdentityVerificationCreateRequestUser) UnsetClientUserId() {
	o.ClientUserId.Unset()
}

func (o IdentityVerificationCreateRequestUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EmailAddress != nil {
		toSerialize["email_address"] = o.EmailAddress
	}
	if o.PhoneNumber.IsSet() {
		toSerialize["phone_number"] = o.PhoneNumber.Get()
	}
	if o.DateOfBirth != nil {
		toSerialize["date_of_birth"] = o.DateOfBirth
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.IdNumber.IsSet() {
		toSerialize["id_number"] = o.IdNumber.Get()
	}
	if o.ClientUserId.IsSet() {
		toSerialize["client_user_id"] = o.ClientUserId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityVerificationCreateRequestUser struct {
	value *IdentityVerificationCreateRequestUser
	isSet bool
}

func (v NullableIdentityVerificationCreateRequestUser) Get() *IdentityVerificationCreateRequestUser {
	return v.value
}

func (v *NullableIdentityVerificationCreateRequestUser) Set(val *IdentityVerificationCreateRequestUser) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityVerificationCreateRequestUser) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityVerificationCreateRequestUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityVerificationCreateRequestUser(val *IdentityVerificationCreateRequestUser) *NullableIdentityVerificationCreateRequestUser {
	return &NullableIdentityVerificationCreateRequestUser{value: val, isSet: true}
}

func (v NullableIdentityVerificationCreateRequestUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityVerificationCreateRequestUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


