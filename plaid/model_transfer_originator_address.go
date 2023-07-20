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

// TransferOriginatorAddress The originator's address.
type TransferOriginatorAddress struct {
	City string `json:"city"`
	Street string `json:"street"`
	Region string `json:"region"`
	PostalCode string `json:"postal_code"`
	// ISO-3166-1 alpha-2 country code standard.
	CountryCode string `json:"country_code"`
}

// NewTransferOriginatorAddress instantiates a new TransferOriginatorAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferOriginatorAddress(city string, street string, region string, postalCode string, countryCode string) *TransferOriginatorAddress {
	this := TransferOriginatorAddress{}
	this.City = city
	this.Street = street
	this.Region = region
	this.PostalCode = postalCode
	this.CountryCode = countryCode
	return &this
}

// NewTransferOriginatorAddressWithDefaults instantiates a new TransferOriginatorAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferOriginatorAddressWithDefaults() *TransferOriginatorAddress {
	this := TransferOriginatorAddress{}
	return &this
}

// GetCity returns the City field value
func (o *TransferOriginatorAddress) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *TransferOriginatorAddress) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *TransferOriginatorAddress) SetCity(v string) {
	o.City = v
}

// GetStreet returns the Street field value
func (o *TransferOriginatorAddress) GetStreet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Street
}

// GetStreetOk returns a tuple with the Street field value
// and a boolean to check if the value has been set.
func (o *TransferOriginatorAddress) GetStreetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Street, true
}

// SetStreet sets field value
func (o *TransferOriginatorAddress) SetStreet(v string) {
	o.Street = v
}

// GetRegion returns the Region field value
func (o *TransferOriginatorAddress) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *TransferOriginatorAddress) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *TransferOriginatorAddress) SetRegion(v string) {
	o.Region = v
}

// GetPostalCode returns the PostalCode field value
func (o *TransferOriginatorAddress) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *TransferOriginatorAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *TransferOriginatorAddress) SetPostalCode(v string) {
	o.PostalCode = v
}

// GetCountryCode returns the CountryCode field value
func (o *TransferOriginatorAddress) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *TransferOriginatorAddress) GetCountryCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *TransferOriginatorAddress) SetCountryCode(v string) {
	o.CountryCode = v
}

func (o TransferOriginatorAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["city"] = o.City
	}
	if true {
		toSerialize["street"] = o.Street
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if true {
		toSerialize["postal_code"] = o.PostalCode
	}
	if true {
		toSerialize["country_code"] = o.CountryCode
	}
	return json.Marshal(toSerialize)
}

type NullableTransferOriginatorAddress struct {
	value *TransferOriginatorAddress
	isSet bool
}

func (v NullableTransferOriginatorAddress) Get() *TransferOriginatorAddress {
	return v.value
}

func (v *NullableTransferOriginatorAddress) Set(val *TransferOriginatorAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferOriginatorAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferOriginatorAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferOriginatorAddress(val *TransferOriginatorAddress) *NullableTransferOriginatorAddress {
	return &NullableTransferOriginatorAddress{value: val, isSet: true}
}

func (v NullableTransferOriginatorAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferOriginatorAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


