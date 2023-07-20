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

// TransferAuthorizationUserInRequest The legal name and other information for the account holder.
type TransferAuthorizationUserInRequest struct {
	// The user's legal name.
	LegalName string `json:"legal_name"`
	// The user's phone number. In order to qualify for a guaranteed transfer, at least one of `phone_number` or `email_address` must be provided.
	PhoneNumber *string `json:"phone_number,omitempty"`
	// The user's email address. In order to qualify for a guaranteed transfer, at least one of `phone_number` or `email_address` must be provided.
	EmailAddress *string `json:"email_address,omitempty"`
	Address *TransferUserAddressInRequest `json:"address,omitempty"`
}

// NewTransferAuthorizationUserInRequest instantiates a new TransferAuthorizationUserInRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferAuthorizationUserInRequest(legalName string) *TransferAuthorizationUserInRequest {
	this := TransferAuthorizationUserInRequest{}
	this.LegalName = legalName
	return &this
}

// NewTransferAuthorizationUserInRequestWithDefaults instantiates a new TransferAuthorizationUserInRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferAuthorizationUserInRequestWithDefaults() *TransferAuthorizationUserInRequest {
	this := TransferAuthorizationUserInRequest{}
	return &this
}

// GetLegalName returns the LegalName field value
func (o *TransferAuthorizationUserInRequest) GetLegalName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LegalName
}

// GetLegalNameOk returns a tuple with the LegalName field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationUserInRequest) GetLegalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LegalName, true
}

// SetLegalName sets field value
func (o *TransferAuthorizationUserInRequest) SetLegalName(v string) {
	o.LegalName = v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *TransferAuthorizationUserInRequest) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationUserInRequest) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *TransferAuthorizationUserInRequest) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *TransferAuthorizationUserInRequest) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *TransferAuthorizationUserInRequest) GetEmailAddress() string {
	if o == nil || o.EmailAddress == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationUserInRequest) GetEmailAddressOk() (*string, bool) {
	if o == nil || o.EmailAddress == nil {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *TransferAuthorizationUserInRequest) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *TransferAuthorizationUserInRequest) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *TransferAuthorizationUserInRequest) GetAddress() TransferUserAddressInRequest {
	if o == nil || o.Address == nil {
		var ret TransferUserAddressInRequest
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationUserInRequest) GetAddressOk() (*TransferUserAddressInRequest, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *TransferAuthorizationUserInRequest) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given TransferUserAddressInRequest and assigns it to the Address field.
func (o *TransferAuthorizationUserInRequest) SetAddress(v TransferUserAddressInRequest) {
	o.Address = &v
}

func (o TransferAuthorizationUserInRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
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

type NullableTransferAuthorizationUserInRequest struct {
	value *TransferAuthorizationUserInRequest
	isSet bool
}

func (v NullableTransferAuthorizationUserInRequest) Get() *TransferAuthorizationUserInRequest {
	return v.value
}

func (v *NullableTransferAuthorizationUserInRequest) Set(val *TransferAuthorizationUserInRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferAuthorizationUserInRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferAuthorizationUserInRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferAuthorizationUserInRequest(val *TransferAuthorizationUserInRequest) *NullableTransferAuthorizationUserInRequest {
	return &NullableTransferAuthorizationUserInRequest{value: val, isSet: true}
}

func (v NullableTransferAuthorizationUserInRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferAuthorizationUserInRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


