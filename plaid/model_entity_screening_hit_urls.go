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

// EntityScreeningHitUrls URLs associated with the entity screening hit
type EntityScreeningHitUrls struct {
	// An 'http' or 'https' URL (must begin with either of those).
	Url string `json:"url"`
	AdditionalProperties map[string]interface{}
}

type _EntityScreeningHitUrls EntityScreeningHitUrls

// NewEntityScreeningHitUrls instantiates a new EntityScreeningHitUrls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityScreeningHitUrls(url string) *EntityScreeningHitUrls {
	this := EntityScreeningHitUrls{}
	this.Url = url
	return &this
}

// NewEntityScreeningHitUrlsWithDefaults instantiates a new EntityScreeningHitUrls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityScreeningHitUrlsWithDefaults() *EntityScreeningHitUrls {
	this := EntityScreeningHitUrls{}
	return &this
}

// GetUrl returns the Url field value
func (o *EntityScreeningHitUrls) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitUrls) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *EntityScreeningHitUrls) SetUrl(v string) {
	o.Url = v
}

func (o EntityScreeningHitUrls) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntityScreeningHitUrls) UnmarshalJSON(bytes []byte) (err error) {
	varEntityScreeningHitUrls := _EntityScreeningHitUrls{}

	if err = json.Unmarshal(bytes, &varEntityScreeningHitUrls); err == nil {
		*o = EntityScreeningHitUrls(varEntityScreeningHitUrls)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntityScreeningHitUrls struct {
	value *EntityScreeningHitUrls
	isSet bool
}

func (v NullableEntityScreeningHitUrls) Get() *EntityScreeningHitUrls {
	return v.value
}

func (v *NullableEntityScreeningHitUrls) Set(val *EntityScreeningHitUrls) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityScreeningHitUrls) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityScreeningHitUrls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityScreeningHitUrls(val *EntityScreeningHitUrls) *NullableEntityScreeningHitUrls {
	return &NullableEntityScreeningHitUrls{value: val, isSet: true}
}

func (v NullableEntityScreeningHitUrls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityScreeningHitUrls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


