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

// ImageQuality A high level description of the quality of the image the user submitted.  For example, an image that is blurry, distorted by glare from a nearby light source, or improperly framed might be marked as low or medium quality. Poor quality images are more likely to fail OCR and/or template conformity checks.  Note: By default, Plaid will let a user recapture document images twice before failing the entire session if we attribute the failure to low image quality.
type ImageQuality string

var _ = fmt.Printf

// List of ImageQuality
const (
	IMAGEQUALITY_HIGH ImageQuality = "high"
	IMAGEQUALITY_MEDIUM ImageQuality = "medium"
	IMAGEQUALITY_LOW ImageQuality = "low"
)

var allowedImageQualityEnumValues = []ImageQuality{
	"high",
	"medium",
	"low",
}

func (v *ImageQuality) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := ImageQuality(value)


	*v = enumTypeValue
	return nil
}

// NewImageQualityFromValue returns a pointer to a valid ImageQuality
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewImageQualityFromValue(v string) (*ImageQuality, error) {
	ev := ImageQuality(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImageQuality) IsValid() bool {
	for _, existing := range allowedImageQualityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImageQuality value
func (v ImageQuality) Ptr() *ImageQuality {
	return &v
}

type NullableImageQuality struct {
	value *ImageQuality
	isSet bool
}

func (v NullableImageQuality) Get() *ImageQuality {
	return v.value
}

func (v *NullableImageQuality) Set(val *ImageQuality) {
	v.value = val
	v.isSet = true
}

func (v NullableImageQuality) IsSet() bool {
	return v.isSet
}

func (v *NullableImageQuality) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageQuality(val *ImageQuality) *NullableImageQuality {
	return &NullableImageQuality{value: val, isSet: true}
}

func (v NullableImageQuality) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageQuality) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

