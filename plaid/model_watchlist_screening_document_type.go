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
	"fmt"
)

// WatchlistScreeningDocumentType The kind of official document represented by this object.  `birth_certificate` - A certificate of birth  `drivers_license` - A license to operate a motor vehicle  `immigration_number` - Immigration or residence documents  `military_id` - Identification issued by a military group  `other` - Any document not covered by other categories  `passport` - An official passport issue by a government  `personal_identification` - Any generic personal identification that is not covered by other categories  `ration_card` - Identification that entitles the holder to rations  `ssn` - United States Social Security Number  `student_id` - Identification issued by an educational institution  `tax_id` - Identification issued for the purpose of collecting taxes  `travel_document` - Visas, entry permits, refugee documents, etc.  `voter_id` - Identification issued for the purpose of voting
type WatchlistScreeningDocumentType string

var _ = fmt.Printf

// List of WatchlistScreeningDocumentType
const (
	WATCHLISTSCREENINGDOCUMENTTYPE_BIRTH_CERTIFICATE WatchlistScreeningDocumentType = "birth_certificate"
	WATCHLISTSCREENINGDOCUMENTTYPE_DRIVERS_LICENSE WatchlistScreeningDocumentType = "drivers_license"
	WATCHLISTSCREENINGDOCUMENTTYPE_IMMIGRATION_NUMBER WatchlistScreeningDocumentType = "immigration_number"
	WATCHLISTSCREENINGDOCUMENTTYPE_MILITARY_ID WatchlistScreeningDocumentType = "military_id"
	WATCHLISTSCREENINGDOCUMENTTYPE_OTHER WatchlistScreeningDocumentType = "other"
	WATCHLISTSCREENINGDOCUMENTTYPE_PASSPORT WatchlistScreeningDocumentType = "passport"
	WATCHLISTSCREENINGDOCUMENTTYPE_PERSONAL_IDENTIFICATION WatchlistScreeningDocumentType = "personal_identification"
	WATCHLISTSCREENINGDOCUMENTTYPE_RATION_CARD WatchlistScreeningDocumentType = "ration_card"
	WATCHLISTSCREENINGDOCUMENTTYPE_SSN WatchlistScreeningDocumentType = "ssn"
	WATCHLISTSCREENINGDOCUMENTTYPE_STUDENT_ID WatchlistScreeningDocumentType = "student_id"
	WATCHLISTSCREENINGDOCUMENTTYPE_TAX_ID WatchlistScreeningDocumentType = "tax_id"
	WATCHLISTSCREENINGDOCUMENTTYPE_TRAVEL_DOCUMENT WatchlistScreeningDocumentType = "travel_document"
	WATCHLISTSCREENINGDOCUMENTTYPE_VOTER_ID WatchlistScreeningDocumentType = "voter_id"
)

var allowedWatchlistScreeningDocumentTypeEnumValues = []WatchlistScreeningDocumentType{
	"birth_certificate",
	"drivers_license",
	"immigration_number",
	"military_id",
	"other",
	"passport",
	"personal_identification",
	"ration_card",
	"ssn",
	"student_id",
	"tax_id",
	"travel_document",
	"voter_id",
}

func (v *WatchlistScreeningDocumentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := WatchlistScreeningDocumentType(value)


	*v = enumTypeValue
	return nil
}

// NewWatchlistScreeningDocumentTypeFromValue returns a pointer to a valid WatchlistScreeningDocumentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWatchlistScreeningDocumentTypeFromValue(v string) (*WatchlistScreeningDocumentType, error) {
	ev := WatchlistScreeningDocumentType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WatchlistScreeningDocumentType) IsValid() bool {
	for _, existing := range allowedWatchlistScreeningDocumentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WatchlistScreeningDocumentType value
func (v WatchlistScreeningDocumentType) Ptr() *WatchlistScreeningDocumentType {
	return &v
}

type NullableWatchlistScreeningDocumentType struct {
	value *WatchlistScreeningDocumentType
	isSet bool
}

func (v NullableWatchlistScreeningDocumentType) Get() *WatchlistScreeningDocumentType {
	return v.value
}

func (v *NullableWatchlistScreeningDocumentType) Set(val *WatchlistScreeningDocumentType) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningDocumentType) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningDocumentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningDocumentType(val *WatchlistScreeningDocumentType) *NullableWatchlistScreeningDocumentType {
	return &NullableWatchlistScreeningDocumentType{value: val, isSet: true}
}

func (v NullableWatchlistScreeningDocumentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningDocumentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

