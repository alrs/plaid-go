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

// WatchlistScreeningIndividualUpdateResponse The screening object allows you to represent a customer in your system, update their profile, and search for them on various watchlists. Note: Rejected customers will not receive new hits, regardless of program configuration.
type WatchlistScreeningIndividualUpdateResponse struct {
	// ID of the associated screening.
	Id string `json:"id"`
	SearchTerms WatchlistScreeningSearchTerms `json:"search_terms"`
	// ID of the associated user.
	Assignee NullableString `json:"assignee"`
	Status WatchlistScreeningStatus `json:"status"`
	// An identifier to help you connect this object to your internal systems. For example, your database ID corresponding to this object.
	ClientUserId NullableString `json:"client_user_id"`
	AuditTrail WatchlistScreeningAuditTrail `json:"audit_trail"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _WatchlistScreeningIndividualUpdateResponse WatchlistScreeningIndividualUpdateResponse

// NewWatchlistScreeningIndividualUpdateResponse instantiates a new WatchlistScreeningIndividualUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningIndividualUpdateResponse(id string, searchTerms WatchlistScreeningSearchTerms, assignee NullableString, status WatchlistScreeningStatus, clientUserId NullableString, auditTrail WatchlistScreeningAuditTrail, requestId string) *WatchlistScreeningIndividualUpdateResponse {
	this := WatchlistScreeningIndividualUpdateResponse{}
	this.Id = id
	this.SearchTerms = searchTerms
	this.Assignee = assignee
	this.Status = status
	this.ClientUserId = clientUserId
	this.AuditTrail = auditTrail
	this.RequestId = requestId
	return &this
}

// NewWatchlistScreeningIndividualUpdateResponseWithDefaults instantiates a new WatchlistScreeningIndividualUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningIndividualUpdateResponseWithDefaults() *WatchlistScreeningIndividualUpdateResponse {
	this := WatchlistScreeningIndividualUpdateResponse{}
	return &this
}

// GetId returns the Id field value
func (o *WatchlistScreeningIndividualUpdateResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualUpdateResponse) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WatchlistScreeningIndividualUpdateResponse) SetId(v string) {
	o.Id = v
}

// GetSearchTerms returns the SearchTerms field value
func (o *WatchlistScreeningIndividualUpdateResponse) GetSearchTerms() WatchlistScreeningSearchTerms {
	if o == nil {
		var ret WatchlistScreeningSearchTerms
		return ret
	}

	return o.SearchTerms
}

// GetSearchTermsOk returns a tuple with the SearchTerms field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualUpdateResponse) GetSearchTermsOk() (*WatchlistScreeningSearchTerms, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SearchTerms, true
}

// SetSearchTerms sets field value
func (o *WatchlistScreeningIndividualUpdateResponse) SetSearchTerms(v WatchlistScreeningSearchTerms) {
	o.SearchTerms = v
}

// GetAssignee returns the Assignee field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchlistScreeningIndividualUpdateResponse) GetAssignee() string {
	if o == nil || o.Assignee.Get() == nil {
		var ret string
		return ret
	}

	return *o.Assignee.Get()
}

// GetAssigneeOk returns a tuple with the Assignee field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningIndividualUpdateResponse) GetAssigneeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Assignee.Get(), o.Assignee.IsSet()
}

// SetAssignee sets field value
func (o *WatchlistScreeningIndividualUpdateResponse) SetAssignee(v string) {
	o.Assignee.Set(&v)
}

// GetStatus returns the Status field value
func (o *WatchlistScreeningIndividualUpdateResponse) GetStatus() WatchlistScreeningStatus {
	if o == nil {
		var ret WatchlistScreeningStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualUpdateResponse) GetStatusOk() (*WatchlistScreeningStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WatchlistScreeningIndividualUpdateResponse) SetStatus(v WatchlistScreeningStatus) {
	o.Status = v
}

// GetClientUserId returns the ClientUserId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchlistScreeningIndividualUpdateResponse) GetClientUserId() string {
	if o == nil || o.ClientUserId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ClientUserId.Get()
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningIndividualUpdateResponse) GetClientUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientUserId.Get(), o.ClientUserId.IsSet()
}

// SetClientUserId sets field value
func (o *WatchlistScreeningIndividualUpdateResponse) SetClientUserId(v string) {
	o.ClientUserId.Set(&v)
}

// GetAuditTrail returns the AuditTrail field value
func (o *WatchlistScreeningIndividualUpdateResponse) GetAuditTrail() WatchlistScreeningAuditTrail {
	if o == nil {
		var ret WatchlistScreeningAuditTrail
		return ret
	}

	return o.AuditTrail
}

// GetAuditTrailOk returns a tuple with the AuditTrail field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualUpdateResponse) GetAuditTrailOk() (*WatchlistScreeningAuditTrail, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuditTrail, true
}

// SetAuditTrail sets field value
func (o *WatchlistScreeningIndividualUpdateResponse) SetAuditTrail(v WatchlistScreeningAuditTrail) {
	o.AuditTrail = v
}

// GetRequestId returns the RequestId field value
func (o *WatchlistScreeningIndividualUpdateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualUpdateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *WatchlistScreeningIndividualUpdateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o WatchlistScreeningIndividualUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["search_terms"] = o.SearchTerms
	}
	if true {
		toSerialize["assignee"] = o.Assignee.Get()
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["client_user_id"] = o.ClientUserId.Get()
	}
	if true {
		toSerialize["audit_trail"] = o.AuditTrail
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WatchlistScreeningIndividualUpdateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWatchlistScreeningIndividualUpdateResponse := _WatchlistScreeningIndividualUpdateResponse{}

	if err = json.Unmarshal(bytes, &varWatchlistScreeningIndividualUpdateResponse); err == nil {
		*o = WatchlistScreeningIndividualUpdateResponse(varWatchlistScreeningIndividualUpdateResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "search_terms")
		delete(additionalProperties, "assignee")
		delete(additionalProperties, "status")
		delete(additionalProperties, "client_user_id")
		delete(additionalProperties, "audit_trail")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWatchlistScreeningIndividualUpdateResponse struct {
	value *WatchlistScreeningIndividualUpdateResponse
	isSet bool
}

func (v NullableWatchlistScreeningIndividualUpdateResponse) Get() *WatchlistScreeningIndividualUpdateResponse {
	return v.value
}

func (v *NullableWatchlistScreeningIndividualUpdateResponse) Set(val *WatchlistScreeningIndividualUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningIndividualUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningIndividualUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningIndividualUpdateResponse(val *WatchlistScreeningIndividualUpdateResponse) *NullableWatchlistScreeningIndividualUpdateResponse {
	return &NullableWatchlistScreeningIndividualUpdateResponse{value: val, isSet: true}
}

func (v NullableWatchlistScreeningIndividualUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningIndividualUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


