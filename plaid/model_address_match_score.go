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

// AddressMatchScore Score found by matching address provided by the API with the address on the account at the financial institution. The score can range from 0 to 100 where 100 is a perfect match and 0 is a no match. If the account contains multiple owners, the maximum match score is filled.
type AddressMatchScore struct {
	// Match score for address. The score can range from 0 to 100 where 100 is a perfect match and 0 is a no match. If the address is missing from either the API or financial institution, this is null.
	Score NullableInt32 `json:"score,omitempty"`
	// postal code was provided for both and was a match
	IsPostalCodeMatch NullableBool `json:"is_postal_code_match,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AddressMatchScore AddressMatchScore

// NewAddressMatchScore instantiates a new AddressMatchScore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressMatchScore() *AddressMatchScore {
	this := AddressMatchScore{}
	return &this
}

// NewAddressMatchScoreWithDefaults instantiates a new AddressMatchScore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressMatchScoreWithDefaults() *AddressMatchScore {
	this := AddressMatchScore{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressMatchScore) GetScore() int32 {
	if o == nil || o.Score.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Score.Get()
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressMatchScore) GetScoreOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Score.Get(), o.Score.IsSet()
}

// HasScore returns a boolean if a field has been set.
func (o *AddressMatchScore) HasScore() bool {
	if o != nil && o.Score.IsSet() {
		return true
	}

	return false
}

// SetScore gets a reference to the given NullableInt32 and assigns it to the Score field.
func (o *AddressMatchScore) SetScore(v int32) {
	o.Score.Set(&v)
}
// SetScoreNil sets the value for Score to be an explicit nil
func (o *AddressMatchScore) SetScoreNil() {
	o.Score.Set(nil)
}

// UnsetScore ensures that no value is present for Score, not even an explicit nil
func (o *AddressMatchScore) UnsetScore() {
	o.Score.Unset()
}

// GetIsPostalCodeMatch returns the IsPostalCodeMatch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressMatchScore) GetIsPostalCodeMatch() bool {
	if o == nil || o.IsPostalCodeMatch.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsPostalCodeMatch.Get()
}

// GetIsPostalCodeMatchOk returns a tuple with the IsPostalCodeMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressMatchScore) GetIsPostalCodeMatchOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsPostalCodeMatch.Get(), o.IsPostalCodeMatch.IsSet()
}

// HasIsPostalCodeMatch returns a boolean if a field has been set.
func (o *AddressMatchScore) HasIsPostalCodeMatch() bool {
	if o != nil && o.IsPostalCodeMatch.IsSet() {
		return true
	}

	return false
}

// SetIsPostalCodeMatch gets a reference to the given NullableBool and assigns it to the IsPostalCodeMatch field.
func (o *AddressMatchScore) SetIsPostalCodeMatch(v bool) {
	o.IsPostalCodeMatch.Set(&v)
}
// SetIsPostalCodeMatchNil sets the value for IsPostalCodeMatch to be an explicit nil
func (o *AddressMatchScore) SetIsPostalCodeMatchNil() {
	o.IsPostalCodeMatch.Set(nil)
}

// UnsetIsPostalCodeMatch ensures that no value is present for IsPostalCodeMatch, not even an explicit nil
func (o *AddressMatchScore) UnsetIsPostalCodeMatch() {
	o.IsPostalCodeMatch.Unset()
}

func (o AddressMatchScore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Score.IsSet() {
		toSerialize["score"] = o.Score.Get()
	}
	if o.IsPostalCodeMatch.IsSet() {
		toSerialize["is_postal_code_match"] = o.IsPostalCodeMatch.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AddressMatchScore) UnmarshalJSON(bytes []byte) (err error) {
	varAddressMatchScore := _AddressMatchScore{}

	if err = json.Unmarshal(bytes, &varAddressMatchScore); err == nil {
		*o = AddressMatchScore(varAddressMatchScore)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "score")
		delete(additionalProperties, "is_postal_code_match")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAddressMatchScore struct {
	value *AddressMatchScore
	isSet bool
}

func (v NullableAddressMatchScore) Get() *AddressMatchScore {
	return v.value
}

func (v *NullableAddressMatchScore) Set(val *AddressMatchScore) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressMatchScore) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressMatchScore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressMatchScore(val *AddressMatchScore) *NullableAddressMatchScore {
	return &NullableAddressMatchScore{value: val, isSet: true}
}

func (v NullableAddressMatchScore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressMatchScore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


