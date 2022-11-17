/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.205.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// EmailAddressMatchScore Score found by matching email provided by the API with the email on the account at the financial institution. If the account contains multiple owners, the maximum match score is filled.
type EmailAddressMatchScore struct {
	// Match score for normalized email. 100 is a perfect match and 0 is a no match. If the email is missing from either the API or financial institution, this is empty.
	Score NullableInt32 `json:"score,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmailAddressMatchScore EmailAddressMatchScore

// NewEmailAddressMatchScore instantiates a new EmailAddressMatchScore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailAddressMatchScore() *EmailAddressMatchScore {
	this := EmailAddressMatchScore{}
	return &this
}

// NewEmailAddressMatchScoreWithDefaults instantiates a new EmailAddressMatchScore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailAddressMatchScoreWithDefaults() *EmailAddressMatchScore {
	this := EmailAddressMatchScore{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailAddressMatchScore) GetScore() int32 {
	if o == nil || o.Score.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Score.Get()
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailAddressMatchScore) GetScoreOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Score.Get(), o.Score.IsSet()
}

// HasScore returns a boolean if a field has been set.
func (o *EmailAddressMatchScore) HasScore() bool {
	if o != nil && o.Score.IsSet() {
		return true
	}

	return false
}

// SetScore gets a reference to the given NullableInt32 and assigns it to the Score field.
func (o *EmailAddressMatchScore) SetScore(v int32) {
	o.Score.Set(&v)
}
// SetScoreNil sets the value for Score to be an explicit nil
func (o *EmailAddressMatchScore) SetScoreNil() {
	o.Score.Set(nil)
}

// UnsetScore ensures that no value is present for Score, not even an explicit nil
func (o *EmailAddressMatchScore) UnsetScore() {
	o.Score.Unset()
}

func (o EmailAddressMatchScore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Score.IsSet() {
		toSerialize["score"] = o.Score.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EmailAddressMatchScore) UnmarshalJSON(bytes []byte) (err error) {
	varEmailAddressMatchScore := _EmailAddressMatchScore{}

	if err = json.Unmarshal(bytes, &varEmailAddressMatchScore); err == nil {
		*o = EmailAddressMatchScore(varEmailAddressMatchScore)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "score")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmailAddressMatchScore struct {
	value *EmailAddressMatchScore
	isSet bool
}

func (v NullableEmailAddressMatchScore) Get() *EmailAddressMatchScore {
	return v.value
}

func (v *NullableEmailAddressMatchScore) Set(val *EmailAddressMatchScore) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailAddressMatchScore) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailAddressMatchScore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailAddressMatchScore(val *EmailAddressMatchScore) *NullableEmailAddressMatchScore {
	return &NullableEmailAddressMatchScore{value: val, isSet: true}
}

func (v NullableEmailAddressMatchScore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailAddressMatchScore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


