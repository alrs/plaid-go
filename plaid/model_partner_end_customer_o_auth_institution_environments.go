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

// PartnerEndCustomerOAuthInstitutionEnvironments Registration statuses by environment.
type PartnerEndCustomerOAuthInstitutionEnvironments struct {
	Development *PartnerEndCustomerOAuthInstitutionApplicationStatus `json:"development,omitempty"`
	Production *PartnerEndCustomerOAuthInstitutionApplicationStatus `json:"production,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PartnerEndCustomerOAuthInstitutionEnvironments PartnerEndCustomerOAuthInstitutionEnvironments

// NewPartnerEndCustomerOAuthInstitutionEnvironments instantiates a new PartnerEndCustomerOAuthInstitutionEnvironments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerEndCustomerOAuthInstitutionEnvironments() *PartnerEndCustomerOAuthInstitutionEnvironments {
	this := PartnerEndCustomerOAuthInstitutionEnvironments{}
	return &this
}

// NewPartnerEndCustomerOAuthInstitutionEnvironmentsWithDefaults instantiates a new PartnerEndCustomerOAuthInstitutionEnvironments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerEndCustomerOAuthInstitutionEnvironmentsWithDefaults() *PartnerEndCustomerOAuthInstitutionEnvironments {
	this := PartnerEndCustomerOAuthInstitutionEnvironments{}
	return &this
}

// GetDevelopment returns the Development field value if set, zero value otherwise.
func (o *PartnerEndCustomerOAuthInstitutionEnvironments) GetDevelopment() PartnerEndCustomerOAuthInstitutionApplicationStatus {
	if o == nil || o.Development == nil {
		var ret PartnerEndCustomerOAuthInstitutionApplicationStatus
		return ret
	}
	return *o.Development
}

// GetDevelopmentOk returns a tuple with the Development field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerOAuthInstitutionEnvironments) GetDevelopmentOk() (*PartnerEndCustomerOAuthInstitutionApplicationStatus, bool) {
	if o == nil || o.Development == nil {
		return nil, false
	}
	return o.Development, true
}

// HasDevelopment returns a boolean if a field has been set.
func (o *PartnerEndCustomerOAuthInstitutionEnvironments) HasDevelopment() bool {
	if o != nil && o.Development != nil {
		return true
	}

	return false
}

// SetDevelopment gets a reference to the given PartnerEndCustomerOAuthInstitutionApplicationStatus and assigns it to the Development field.
func (o *PartnerEndCustomerOAuthInstitutionEnvironments) SetDevelopment(v PartnerEndCustomerOAuthInstitutionApplicationStatus) {
	o.Development = &v
}

// GetProduction returns the Production field value if set, zero value otherwise.
func (o *PartnerEndCustomerOAuthInstitutionEnvironments) GetProduction() PartnerEndCustomerOAuthInstitutionApplicationStatus {
	if o == nil || o.Production == nil {
		var ret PartnerEndCustomerOAuthInstitutionApplicationStatus
		return ret
	}
	return *o.Production
}

// GetProductionOk returns a tuple with the Production field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerOAuthInstitutionEnvironments) GetProductionOk() (*PartnerEndCustomerOAuthInstitutionApplicationStatus, bool) {
	if o == nil || o.Production == nil {
		return nil, false
	}
	return o.Production, true
}

// HasProduction returns a boolean if a field has been set.
func (o *PartnerEndCustomerOAuthInstitutionEnvironments) HasProduction() bool {
	if o != nil && o.Production != nil {
		return true
	}

	return false
}

// SetProduction gets a reference to the given PartnerEndCustomerOAuthInstitutionApplicationStatus and assigns it to the Production field.
func (o *PartnerEndCustomerOAuthInstitutionEnvironments) SetProduction(v PartnerEndCustomerOAuthInstitutionApplicationStatus) {
	o.Production = &v
}

func (o PartnerEndCustomerOAuthInstitutionEnvironments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Development != nil {
		toSerialize["development"] = o.Development
	}
	if o.Production != nil {
		toSerialize["production"] = o.Production
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PartnerEndCustomerOAuthInstitutionEnvironments) UnmarshalJSON(bytes []byte) (err error) {
	varPartnerEndCustomerOAuthInstitutionEnvironments := _PartnerEndCustomerOAuthInstitutionEnvironments{}

	if err = json.Unmarshal(bytes, &varPartnerEndCustomerOAuthInstitutionEnvironments); err == nil {
		*o = PartnerEndCustomerOAuthInstitutionEnvironments(varPartnerEndCustomerOAuthInstitutionEnvironments)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "development")
		delete(additionalProperties, "production")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePartnerEndCustomerOAuthInstitutionEnvironments struct {
	value *PartnerEndCustomerOAuthInstitutionEnvironments
	isSet bool
}

func (v NullablePartnerEndCustomerOAuthInstitutionEnvironments) Get() *PartnerEndCustomerOAuthInstitutionEnvironments {
	return v.value
}

func (v *NullablePartnerEndCustomerOAuthInstitutionEnvironments) Set(val *PartnerEndCustomerOAuthInstitutionEnvironments) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerEndCustomerOAuthInstitutionEnvironments) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerEndCustomerOAuthInstitutionEnvironments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerEndCustomerOAuthInstitutionEnvironments(val *PartnerEndCustomerOAuthInstitutionEnvironments) *NullablePartnerEndCustomerOAuthInstitutionEnvironments {
	return &NullablePartnerEndCustomerOAuthInstitutionEnvironments{value: val, isSet: true}
}

func (v NullablePartnerEndCustomerOAuthInstitutionEnvironments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerEndCustomerOAuthInstitutionEnvironments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


