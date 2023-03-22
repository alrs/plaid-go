/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.343.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// SignalAddressData Data about the components comprising an address.
type SignalAddressData struct {
	// The full city name
	City *string `json:"city,omitempty"`
	// The region or state Example: `\"NC\"`
	Region NullableString `json:"region,omitempty"`
	// The full street address Example: `\"564 Main Street, APT 15\"`
	Street *string `json:"street,omitempty"`
	// The postal code
	PostalCode NullableString `json:"postal_code,omitempty"`
	// The ISO 3166-1 alpha-2 country code
	Country NullableString `json:"country,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SignalAddressData SignalAddressData

// NewSignalAddressData instantiates a new SignalAddressData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignalAddressData() *SignalAddressData {
	this := SignalAddressData{}
	return &this
}

// NewSignalAddressDataWithDefaults instantiates a new SignalAddressData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignalAddressDataWithDefaults() *SignalAddressData {
	this := SignalAddressData{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *SignalAddressData) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalAddressData) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *SignalAddressData) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *SignalAddressData) SetCity(v string) {
	o.City = &v
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignalAddressData) GetRegion() string {
	if o == nil || o.Region.Get() == nil {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignalAddressData) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *SignalAddressData) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *SignalAddressData) SetRegion(v string) {
	o.Region.Set(&v)
}
// SetRegionNil sets the value for Region to be an explicit nil
func (o *SignalAddressData) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *SignalAddressData) UnsetRegion() {
	o.Region.Unset()
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *SignalAddressData) GetStreet() string {
	if o == nil || o.Street == nil {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalAddressData) GetStreetOk() (*string, bool) {
	if o == nil || o.Street == nil {
		return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *SignalAddressData) HasStreet() bool {
	if o != nil && o.Street != nil {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *SignalAddressData) SetStreet(v string) {
	o.Street = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignalAddressData) GetPostalCode() string {
	if o == nil || o.PostalCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignalAddressData) GetPostalCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// HasPostalCode returns a boolean if a field has been set.
func (o *SignalAddressData) HasPostalCode() bool {
	if o != nil && o.PostalCode.IsSet() {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given NullableString and assigns it to the PostalCode field.
func (o *SignalAddressData) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}
// SetPostalCodeNil sets the value for PostalCode to be an explicit nil
func (o *SignalAddressData) SetPostalCodeNil() {
	o.PostalCode.Set(nil)
}

// UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
func (o *SignalAddressData) UnsetPostalCode() {
	o.PostalCode.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignalAddressData) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignalAddressData) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *SignalAddressData) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *SignalAddressData) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *SignalAddressData) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *SignalAddressData) UnsetCountry() {
	o.Country.Unset()
}

func (o SignalAddressData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	if o.Street != nil {
		toSerialize["street"] = o.Street
	}
	if o.PostalCode.IsSet() {
		toSerialize["postal_code"] = o.PostalCode.Get()
	}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SignalAddressData) UnmarshalJSON(bytes []byte) (err error) {
	varSignalAddressData := _SignalAddressData{}

	if err = json.Unmarshal(bytes, &varSignalAddressData); err == nil {
		*o = SignalAddressData(varSignalAddressData)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "city")
		delete(additionalProperties, "region")
		delete(additionalProperties, "street")
		delete(additionalProperties, "postal_code")
		delete(additionalProperties, "country")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSignalAddressData struct {
	value *SignalAddressData
	isSet bool
}

func (v NullableSignalAddressData) Get() *SignalAddressData {
	return v.value
}

func (v *NullableSignalAddressData) Set(val *SignalAddressData) {
	v.value = val
	v.isSet = true
}

func (v NullableSignalAddressData) IsSet() bool {
	return v.isSet
}

func (v *NullableSignalAddressData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignalAddressData(val *SignalAddressData) *NullableSignalAddressData {
	return &NullableSignalAddressData{value: val, isSet: true}
}

func (v NullableSignalAddressData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignalAddressData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


