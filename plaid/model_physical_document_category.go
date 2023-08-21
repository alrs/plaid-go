/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.419.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// PhysicalDocumentCategory The type of identity document detected in the images provided. Will always be one of the following values:    `drivers_license` - A driver's license issued by the associated country, establishing identity without any guarantee as to citizenship, and granting driving privileges    `id_card` - A general national identification card, distinct from driver's licenses as it only establishes identity    `passport` - A travel passport issued by the associated country for one of its citizens    `residence_permit_card` - An identity document issued by the associated country permitting a foreign citizen to <em>temporarily</em> reside there    `resident_card` - An identity document issued by the associated country permitting a foreign citizen to <em>permanently</em> reside there    `visa` - An identity document issued by the associated country permitting a foreign citizen entry for a short duration and for a specific purpose, typically no longer than 6 months  Note: This value may be different from the ID type that the user selects within Link. For example, if they select \"Driver's License\" but then submit a picture of a passport, this field will say `passport`
type PhysicalDocumentCategory string

var _ = fmt.Printf

// List of PhysicalDocumentCategory
const (
	PHYSICALDOCUMENTCATEGORY_DRIVERS_LICENSE PhysicalDocumentCategory = "drivers_license"
	PHYSICALDOCUMENTCATEGORY_ID_CARD PhysicalDocumentCategory = "id_card"
	PHYSICALDOCUMENTCATEGORY_PASSPORT PhysicalDocumentCategory = "passport"
	PHYSICALDOCUMENTCATEGORY_RESIDENCE_PERMIT_CARD PhysicalDocumentCategory = "residence_permit_card"
	PHYSICALDOCUMENTCATEGORY_RESIDENT_CARD PhysicalDocumentCategory = "resident_card"
	PHYSICALDOCUMENTCATEGORY_VISA PhysicalDocumentCategory = "visa"
)

var allowedPhysicalDocumentCategoryEnumValues = []PhysicalDocumentCategory{
	"drivers_license",
	"id_card",
	"passport",
	"residence_permit_card",
	"resident_card",
	"visa",
}

func (v *PhysicalDocumentCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := PhysicalDocumentCategory(value)


	*v = enumTypeValue
	return nil
}

// NewPhysicalDocumentCategoryFromValue returns a pointer to a valid PhysicalDocumentCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPhysicalDocumentCategoryFromValue(v string) (*PhysicalDocumentCategory, error) {
	ev := PhysicalDocumentCategory(v)


	return &ev, nil
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

