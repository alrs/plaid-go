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

// TransferUserInRequestDeprecated The legal name and other information for the account holder.
type TransferUserInRequestDeprecated struct {
	// The user's legal name.
	LegalName *string `json:"legal_name,omitempty"`
	// The user's phone number.
	PhoneNumber *string `json:"phone_number,omitempty"`
	// The user's email address.
	EmailAddress *string `json:"email_address,omitempty"`
	Address *TransferUserAddressInRequest `json:"address,omitempty"`
}

// NewTransferUserInRequestDeprecated instantiates a new TransferUserInRequestDeprecated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferUserInRequestDeprecated() *TransferUserInRequestDeprecated {
	this := TransferUserInRequestDeprecated{}
	return &this
}

// NewTransferUserInRequestDeprecatedWithDefaults instantiates a new TransferUserInRequestDeprecated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferUserInRequestDeprecatedWithDefaults() *TransferUserInRequestDeprecated {
	this := TransferUserInRequestDeprecated{}
	return &this
}

// GetLegalName returns the LegalName field value if set, zero value otherwise.
func (o *TransferUserInRequestDeprecated) GetLegalName() string {
	if o == nil || o.LegalName == nil {
		var ret string
		return ret
	}
	return *o.LegalName
}

// GetLegalNameOk returns a tuple with the LegalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferUserInRequestDeprecated) GetLegalNameOk() (*string, bool) {
	if o == nil || o.LegalName == nil {
		return nil, false
	}
	return o.LegalName, true
}

// HasLegalName returns a boolean if a field has been set.
func (o *TransferUserInRequestDeprecated) HasLegalName() bool {
	if o != nil && o.LegalName != nil {
		return true
	}

	return false
}

// SetLegalName gets a reference to the given string and assigns it to the LegalName field.
func (o *TransferUserInRequestDeprecated) SetLegalName(v string) {
	o.LegalName = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *TransferUserInRequestDeprecated) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferUserInRequestDeprecated) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *TransferUserInRequestDeprecated) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *TransferUserInRequestDeprecated) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *TransferUserInRequestDeprecated) GetEmailAddress() string {
	if o == nil || o.EmailAddress == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferUserInRequestDeprecated) GetEmailAddressOk() (*string, bool) {
	if o == nil || o.EmailAddress == nil {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *TransferUserInRequestDeprecated) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *TransferUserInRequestDeprecated) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *TransferUserInRequestDeprecated) GetAddress() TransferUserAddressInRequest {
	if o == nil || o.Address == nil {
		var ret TransferUserAddressInRequest
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferUserInRequestDeprecated) GetAddressOk() (*TransferUserAddressInRequest, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *TransferUserInRequestDeprecated) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given TransferUserAddressInRequest and assigns it to the Address field.
func (o *TransferUserInRequestDeprecated) SetAddress(v TransferUserAddressInRequest) {
	o.Address = &v
}

func (o TransferUserInRequestDeprecated) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LegalName != nil {
		toSerialize["legal_name"] = o.LegalName
	}
	if o.PhoneNumber != nil {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if o.EmailAddress != nil {
		toSerialize["email_address"] = o.EmailAddress
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableTransferUserInRequestDeprecated struct {
	value *TransferUserInRequestDeprecated
	isSet bool
}

func (v NullableTransferUserInRequestDeprecated) Get() *TransferUserInRequestDeprecated {
	return v.value
}

func (v *NullableTransferUserInRequestDeprecated) Set(val *TransferUserInRequestDeprecated) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferUserInRequestDeprecated) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferUserInRequestDeprecated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferUserInRequestDeprecated(val *TransferUserInRequestDeprecated) *NullableTransferUserInRequestDeprecated {
	return &NullableTransferUserInRequestDeprecated{value: val, isSet: true}
}

func (v NullableTransferUserInRequestDeprecated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferUserInRequestDeprecated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


