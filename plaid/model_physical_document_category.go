/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.181.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// PhysicalDocumentCategory The type of identity document detected in the images provided. Will always be one of the following values:    `drivers_license` - A driver's license for the associated country    `id_card` - A general national identification card, distinct from driver's licenses    `passport` - A passport for the associated country    `residence_permit_card` - An identity document permitting a foreign citizen to <em>temporarily</em> reside in the associated country    `resident_card` - An identity document permitting a foreign citizen to <em>permanently</em> reside in the associated country  Note: This value may be different from the ID type that the user selects within Link. For example, if they select \"Driver's License\" but then submit a picture of a passport, this field will say `passport`
type PhysicalDocumentCategory string

// List of PhysicalDocumentCategory
const (
	PHYSICALDOCUMENTCATEGORY_DRIVERS_LICENSE PhysicalDocumentCategory = "drivers_license"
	PHYSICALDOCUMENTCATEGORY_ID_CARD PhysicalDocumentCategory = "id_card"
	PHYSICALDOCUMENTCATEGORY_PASSPORT PhysicalDocumentCategory = "passport"
	PHYSICALDOCUMENTCATEGORY_RESIDENCE_PERMIT_CARD PhysicalDocumentCategory = "residence_permit_card"
	PHYSICALDOCUMENTCATEGORY_RESIDENT_CARD PhysicalDocumentCategory = "resident_card"
)

var allowedPhysicalDocumentCategoryEnumValues = []PhysicalDocumentCategory{
	"drivers_license",
	"id_card",
	"passport",
	"residence_permit_card",
	"resident_card",
}

func (v *PhysicalDocumentCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PhysicalDocumentCategory(value)
	for _, existing := range allowedPhysicalDocumentCategoryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PhysicalDocumentCategory", value)
}

// NewPhysicalDocumentCategoryFromValue returns a pointer to a valid PhysicalDocumentCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPhysicalDocumentCategoryFromValue(v string) (*PhysicalDocumentCategory, error) {
	ev := PhysicalDocumentCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PhysicalDocumentCategory: valid values are %v", v, allowedPhysicalDocumentCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PhysicalDocumentCategory) IsValid() bool {
	for _, existing := range allowedPhysicalDocumentCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PhysicalDocumentCategory value
func (v PhysicalDocumentCategory) Ptr() *PhysicalDocumentCategory {
	return &v
}

type NullablePhysicalDocumentCategory struct {
	value *PhysicalDocumentCategory
	isSet bool
}

func (v NullablePhysicalDocumentCategory) Get() *PhysicalDocumentCategory {
	return v.value
}

func (v *NullablePhysicalDocumentCategory) Set(val *PhysicalDocumentCategory) {
	v.value = val
	v.isSet = true
}

func (v NullablePhysicalDocumentCategory) IsSet() bool {
	return v.isSet
}

func (v *NullablePhysicalDocumentCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhysicalDocumentCategory(val *PhysicalDocumentCategory) *NullablePhysicalDocumentCategory {
	return &NullablePhysicalDocumentCategory{value: val, isSet: true}
}

func (v NullablePhysicalDocumentCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhysicalDocumentCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

