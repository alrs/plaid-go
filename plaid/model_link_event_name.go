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
	"fmt"
)

// LinkEventName A string representing the event that has just occurred in the Link flow.
type LinkEventName string

var _ = fmt.Printf

// List of LinkEventName
const (
	LINKEVENTNAME_BANK_INCOME_INSIGHTS_COMPLETED LinkEventName = "BANK_INCOME_INSIGHTS_COMPLETED"
	LINKEVENTNAME_CLOSE_OAUTH LinkEventName = "CLOSE_OAUTH"
	LINKEVENTNAME_ERROR LinkEventName = "ERROR"
	LINKEVENTNAME_EXIT LinkEventName = "EXIT"
	LINKEVENTNAME_FAIL_OAUTH LinkEventName = "FAIL_OAUTH"
	LINKEVENTNAME_HANDOFF LinkEventName = "HANDOFF"
	LINKEVENTNAME_OPEN LinkEventName = "OPEN"
	LINKEVENTNAME_OPEN_MY_PLAID LinkEventName = "OPEN_MY_PLAID"
	LINKEVENTNAME_OPEN_OAUTH LinkEventName = "OPEN_OAUTH"
	LINKEVENTNAME_SEARCH_INSTITUTION LinkEventName = "SEARCH_INSTITUTION"
	LINKEVENTNAME_SELECT_AUTH_TYPE LinkEventName = "SELECT_AUTH_TYPE"
	LINKEVENTNAME_SELECT_BRAND LinkEventName = "SELECT_BRAND"
	LINKEVENTNAME_SELECT_DEGRADED_INSTITUTION LinkEventName = "SELECT_DEGRADED_INSTITUTION"
	LINKEVENTNAME_SELECT_DOWN_INSTITUTION LinkEventName = "SELECT_DOWN_INSTITUTION"
	LINKEVENTNAME_SELECT_INSTITUTION LinkEventName = "SELECT_INSTITUTION"
	LINKEVENTNAME_SUBMIT_ACCOUNT_NUMBER LinkEventName = "SUBMIT_ACCOUNT_NUMBER"
	LINKEVENTNAME_SUBMIT_CREDENTIALS LinkEventName = "SUBMIT_CREDENTIALS"
	LINKEVENTNAME_SUBMIT_DOCUMENTS LinkEventName = "SUBMIT_DOCUMENTS"
	LINKEVENTNAME_SUBMIT_DOCUMENTS_ERROR LinkEventName = "SUBMIT_DOCUMENTS_ERROR"
	LINKEVENTNAME_SUBMIT_DOCUMENTS_SUCCESS LinkEventName = "SUBMIT_DOCUMENTS_SUCCESS"
	LINKEVENTNAME_SUBMIT_MFA LinkEventName = "SUBMIT_MFA"
	LINKEVENTNAME_SUBMIT_ROUTING_NUMBER LinkEventName = "SUBMIT_ROUTING_NUMBER"
	LINKEVENTNAME_TRANSITION_VIEW LinkEventName = "TRANSITION_VIEW"
	LINKEVENTNAME_VIEW_DATA_TYPES LinkEventName = "VIEW_DATA_TYPES"
)

var allowedLinkEventNameEnumValues = []LinkEventName{
	"BANK_INCOME_INSIGHTS_COMPLETED",
	"CLOSE_OAUTH",
	"ERROR",
	"EXIT",
	"FAIL_OAUTH",
	"HANDOFF",
	"OPEN",
	"OPEN_MY_PLAID",
	"OPEN_OAUTH",
	"SEARCH_INSTITUTION",
	"SELECT_AUTH_TYPE",
	"SELECT_BRAND",
	"SELECT_DEGRADED_INSTITUTION",
	"SELECT_DOWN_INSTITUTION",
	"SELECT_INSTITUTION",
	"SUBMIT_ACCOUNT_NUMBER",
	"SUBMIT_CREDENTIALS",
	"SUBMIT_DOCUMENTS",
	"SUBMIT_DOCUMENTS_ERROR",
	"SUBMIT_DOCUMENTS_SUCCESS",
	"SUBMIT_MFA",
	"SUBMIT_ROUTING_NUMBER",
	"TRANSITION_VIEW",
	"VIEW_DATA_TYPES",
}

func (v *LinkEventName) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := LinkEventName(value)


	*v = enumTypeValue
	return nil
}

// NewLinkEventNameFromValue returns a pointer to a valid LinkEventName
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLinkEventNameFromValue(v string) (*LinkEventName, error) {
	ev := LinkEventName(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LinkEventName) IsValid() bool {
	for _, existing := range allowedLinkEventNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LinkEventName value
func (v LinkEventName) Ptr() *LinkEventName {
	return &v
}

type NullableLinkEventName struct {
	value *LinkEventName
	isSet bool
}

func (v NullableLinkEventName) Get() *LinkEventName {
	return v.value
}

func (v *NullableLinkEventName) Set(val *LinkEventName) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkEventName) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkEventName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkEventName(val *LinkEventName) *NullableLinkEventName {
	return &NullableLinkEventName{value: val, isSet: true}
}

func (v NullableLinkEventName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkEventName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
