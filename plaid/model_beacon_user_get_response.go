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
	"time"
)

// BeaconUserGetResponse A Beacon User represents an end user that has been scanned against the Beacon Network.
type BeaconUserGetResponse struct {
	// ID of the associated Beacon User.
	Id string `json:"id"`
	// An ISO8601 formatted timestamp.
	CreatedAt time.Time `json:"created_at"`
	// An ISO8601 formatted timestamp. This field indicates the last time the resource was modified.
	UpdatedAt time.Time `json:"updated_at"`
	Status BeaconUserStatus `json:"status"`
	// ID of the associated Beacon Program.
	ProgramId string `json:"program_id"`
	// A unique ID that identifies the end user in your system. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.
	ClientUserId string `json:"client_user_id"`
	User BeaconUserData `json:"user"`
	AuditTrail BeaconAuditTrail `json:"audit_trail"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _BeaconUserGetResponse BeaconUserGetResponse

// NewBeaconUserGetResponse instantiates a new BeaconUserGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBeaconUserGetResponse(id string, createdAt time.Time, updatedAt time.Time, status BeaconUserStatus, programId string, clientUserId string, user BeaconUserData, auditTrail BeaconAuditTrail, requestId string) *BeaconUserGetResponse {
	this := BeaconUserGetResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Status = status
	this.ProgramId = programId
	this.ClientUserId = clientUserId
	this.User = user
	this.AuditTrail = auditTrail
	this.RequestId = requestId
	return &this
}

// NewBeaconUserGetResponseWithDefaults instantiates a new BeaconUserGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBeaconUserGetResponseWithDefaults() *BeaconUserGetResponse {
	this := BeaconUserGetResponse{}
	return &this
}

// GetId returns the Id field value
func (o *BeaconUserGetResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BeaconUserGetResponse) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BeaconUserGetResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *BeaconUserGetResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *BeaconUserGetResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *BeaconUserGetResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *BeaconUserGetResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *BeaconUserGetResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *BeaconUserGetResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetStatus returns the Status field value
func (o *BeaconUserGetResponse) GetStatus() BeaconUserStatus {
	if o == nil {
		var ret BeaconUserStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BeaconUserGetResponse) GetStatusOk() (*BeaconUserStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BeaconUserGetResponse) SetStatus(v BeaconUserStatus) {
	o.Status = v
}

// GetProgramId returns the ProgramId field value
func (o *BeaconUserGetResponse) GetProgramId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProgramId
}

// GetProgramIdOk returns a tuple with the ProgramId field value
// and a boolean to check if the value has been set.
func (o *BeaconUserGetResponse) GetProgramIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProgramId, true
}

// SetProgramId sets field value
func (o *BeaconUserGetResponse) SetProgramId(v string) {
	o.ProgramId = v
}

// GetClientUserId returns the ClientUserId field value
func (o *BeaconUserGetResponse) GetClientUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientUserId
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value
// and a boolean to check if the value has been set.
func (o *BeaconUserGetResponse) GetClientUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientUserId, true
}

// SetClientUserId sets field value
func (o *BeaconUserGetResponse) SetClientUserId(v string) {
	o.ClientUserId = v
}

// GetUser returns the User field value
func (o *BeaconUserGetResponse) GetUser() BeaconUserData {
	if o == nil {
		var ret BeaconUserData
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *BeaconUserGetResponse) GetUserOk() (*BeaconUserData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *BeaconUserGetResponse) SetUser(v BeaconUserData) {
	o.User = v
}

// GetAuditTrail returns the AuditTrail field value
func (o *BeaconUserGetResponse) GetAuditTrail() BeaconAuditTrail {
	if o == nil {
		var ret BeaconAuditTrail
		return ret
	}

	return o.AuditTrail
}

// GetAuditTrailOk returns a tuple with the AuditTrail field value
// and a boolean to check if the value has been set.
func (o *BeaconUserGetResponse) GetAuditTrailOk() (*BeaconAuditTrail, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuditTrail, true
}

// SetAuditTrail sets field value
func (o *BeaconUserGetResponse) SetAuditTrail(v BeaconAuditTrail) {
	o.AuditTrail = v
}

// GetRequestId returns the RequestId field value
func (o *BeaconUserGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *BeaconUserGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *BeaconUserGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o BeaconUserGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["program_id"] = o.ProgramId
	}
	if true {
		toSerialize["client_user_id"] = o.ClientUserId
	}
	if true {
		toSerialize["user"] = o.User
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

func (o *BeaconUserGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBeaconUserGetResponse := _BeaconUserGetResponse{}

	if err = json.Unmarshal(bytes, &varBeaconUserGetResponse); err == nil {
		*o = BeaconUserGetResponse(varBeaconUserGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "status")
		delete(additionalProperties, "program_id")
		delete(additionalProperties, "client_user_id")
		delete(additionalProperties, "user")
		delete(additionalProperties, "audit_trail")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBeaconUserGetResponse struct {
	value *BeaconUserGetResponse
	isSet bool
}

func (v NullableBeaconUserGetResponse) Get() *BeaconUserGetResponse {
	return v.value
}

func (v *NullableBeaconUserGetResponse) Set(val *BeaconUserGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBeaconUserGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBeaconUserGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeaconUserGetResponse(val *BeaconUserGetResponse) *NullableBeaconUserGetResponse {
	return &NullableBeaconUserGetResponse{value: val, isSet: true}
}

func (v NullableBeaconUserGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeaconUserGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

