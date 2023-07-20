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

// Counterparty The counterparty, such as the merchant or financial institution, is extracted by Plaid from the raw description.
type Counterparty struct {
	// The name of the counterparty, such as the merchant or the financial institution, as extracted by Plaid from the raw description.
	Name string `json:"name"`
	// A unique, stable, Plaid-generated id that maps to the counterparty.
	EntityId NullableString `json:"entity_id,omitempty"`
	Type CounterpartyType `json:"type"`
	// The website associated with the counterparty.
	Website NullableString `json:"website"`
	// The URL of a logo associated with the counterparty, if available. The logo is formatted as a 100x100 pixel PNG file.
	LogoUrl NullableString `json:"logo_url"`
	AdditionalProperties map[string]interface{}
}

type _Counterparty Counterparty

// NewCounterparty instantiates a new Counterparty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCounterparty(name string, type_ CounterpartyType, website NullableString, logoUrl NullableString) *Counterparty {
	this := Counterparty{}
	this.Name = name
	this.Type = type_
	this.Website = website
	this.LogoUrl = logoUrl
	return &this
}

// NewCounterpartyWithDefaults instantiates a new Counterparty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCounterpartyWithDefaults() *Counterparty {
	this := Counterparty{}
	return &this
}

// GetName returns the Name field value
func (o *Counterparty) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Counterparty) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Counterparty) SetName(v string) {
	o.Name = v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Counterparty) GetEntityId() string {
	if o == nil || o.EntityId.Get() == nil {
		var ret string
		return ret
	}
	return *o.EntityId.Get()
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Counterparty) GetEntityIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EntityId.Get(), o.EntityId.IsSet()
}

// HasEntityId returns a boolean if a field has been set.
func (o *Counterparty) HasEntityId() bool {
	if o != nil && o.EntityId.IsSet() {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given NullableString and assigns it to the EntityId field.
func (o *Counterparty) SetEntityId(v string) {
	o.EntityId.Set(&v)
}
// SetEntityIdNil sets the value for EntityId to be an explicit nil
func (o *Counterparty) SetEntityIdNil() {
	o.EntityId.Set(nil)
}

// UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
func (o *Counterparty) UnsetEntityId() {
	o.EntityId.Unset()
}

// GetType returns the Type field value
func (o *Counterparty) GetType() CounterpartyType {
	if o == nil {
		var ret CounterpartyType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Counterparty) GetTypeOk() (*CounterpartyType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Counterparty) SetType(v CounterpartyType) {
	o.Type = v
}

// GetWebsite returns the Website field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Counterparty) GetWebsite() string {
	if o == nil || o.Website.Get() == nil {
		var ret string
		return ret
	}

	return *o.Website.Get()
}

// GetWebsiteOk returns a tuple with the Website field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Counterparty) GetWebsiteOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Website.Get(), o.Website.IsSet()
}

// SetWebsite sets field value
func (o *Counterparty) SetWebsite(v string) {
	o.Website.Set(&v)
}

// GetLogoUrl returns the LogoUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Counterparty) GetLogoUrl() string {
	if o == nil || o.LogoUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.LogoUrl.Get()
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Counterparty) GetLogoUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LogoUrl.Get(), o.LogoUrl.IsSet()
}

// SetLogoUrl sets field value
func (o *Counterparty) SetLogoUrl(v string) {
	o.LogoUrl.Set(&v)
}

func (o Counterparty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.EntityId.IsSet() {
		toSerialize["entity_id"] = o.EntityId.Get()
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["website"] = o.Website.Get()
	}
	if true {
		toSerialize["logo_url"] = o.LogoUrl.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Counterparty) UnmarshalJSON(bytes []byte) (err error) {
	varCounterparty := _Counterparty{}

	if err = json.Unmarshal(bytes, &varCounterparty); err == nil {
		*o = Counterparty(varCounterparty)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "entity_id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "website")
		delete(additionalProperties, "logo_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCounterparty struct {
	value *Counterparty
	isSet bool
}

func (v NullableCounterparty) Get() *Counterparty {
	return v.value
}

func (v *NullableCounterparty) Set(val *Counterparty) {
	v.value = val
	v.isSet = true
}

func (v NullableCounterparty) IsSet() bool {
	return v.isSet
}

func (v *NullableCounterparty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCounterparty(val *Counterparty) *NullableCounterparty {
	return &NullableCounterparty{value: val, isSet: true}
}

func (v NullableCounterparty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCounterparty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


