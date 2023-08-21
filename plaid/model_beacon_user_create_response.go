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

// BeaconUserCreateResponse A Beacon User represents an end user that has been scanned against the Beacon Network.
type BeaconUserCreateResponse struct {
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

type _BeaconUserCreateResponse BeaconUserCreateResponse

// NewBeaconUserCreateResponse instantiates a new BeaconUserCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBeaconUserCreateResponse(id string, createdAt time.Time, updatedAt time.Time, status BeaconUserStatus, programId string, clientUserId string, user BeaconUserData, auditTrail BeaconAuditTrail, requestId string) *BeaconUserCreateResponse {
	this := BeaconUserCreateResponse{}
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

// NewBeaconUserCreateResponseWithDefaults instantiates a new BeaconUserCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBeaconUserCreateResponseWithDefaults() *BeaconUserCreateResponse {
	this := BeaconUserCreateResponse{}
	return &this
}

// GetId returns the Id field value
func (o *BeaconUserCreateResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BeaconUserCreateResponse) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BeaconUserCreateResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *BeaconUserCreateResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *BeaconUserCreateResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *BeaconUserCreateResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *BeaconUserCreateResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *BeaconUserCreateResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *BeaconUserCreateResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetStatus returns the Status field value
func (o *BeaconUserCreateResponse) GetStatus() BeaconUserStatus {
	if o == nil {
		var ret BeaconUserStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BeaconUserCreateResponse) GetStatusOk() (*BeaconUserStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BeaconUserCreateResponse) SetStatus(v BeaconUserStatus) {
	o.Status = v
}

// GetProgramId returns the ProgramId field value
func (o *BeaconUserCreateResponse) GetProgramId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProgramId
}

// GetProgramIdOk returns a tuple with the ProgramId field value
// and a boolean to check if the value has been set.
func (o *BeaconUserCreateResponse) GetProgramIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProgramId, true
}

// SetProgramId sets field value
func (o *BeaconUserCreateResponse) SetProgramId(v string) {
	o.ProgramId = v
}

// GetClientUserId returns the ClientUserId field value
func (o *BeaconUserCreateResponse) GetClientUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientUserId
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value
// and a boolean to check if the value has been set.
func (o *BeaconUserCreateResponse) GetClientUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientUserId, true
}

// SetClientUserId sets field value
func (o *BeaconUserCreateResponse) SetClientUserId(v string) {
	o.ClientUserId = v
}

// GetUser returns the User field value
func (o *BeaconUserCreateResponse) GetUser() BeaconUserData {
	if o == nil {
		var ret BeaconUserData
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *BeaconUserCreateResponse) GetUserOk() (*BeaconUserData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *BeaconUserCreateResponse) SetUser(v BeaconUserData) {
	o.User = v
}

// GetAuditTrail returns the AuditTrail field value
func (o *BeaconUserCreateResponse) GetAuditTrail() BeaconAuditTrail {
	if o == nil {
		var ret BeaconAuditTrail
		return ret
	}

	return o.AuditTrail
}

// GetAuditTrailOk returns a tuple with the AuditTrail field value
// and a boolean to check if the value has been set.
func (o *BeaconUserCreateResponse) GetAuditTrailOk() (*BeaconAuditTrail, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuditTrail, true
}

// SetAuditTrail sets field value
func (o *BeaconUserCreateResponse) SetAuditTrail(v BeaconAuditTrail) {
	o.AuditTrail = v
}

// GetRequestId returns the RequestId field value
func (o *BeaconUserCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *BeaconUserCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *BeaconUserCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o BeaconUserCreateResponse) MarshalJSON() ([]byte, error) {
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

func (o *BeaconUserCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBeaconUserCreateResponse := _BeaconUserCreateResponse{}

	if err = json.Unmarshal(bytes, &varBeaconUserCreateResponse); err == nil {
		*o = BeaconUserCreateResponse(varBeaconUserCreateResponse)
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

type NullableBeaconUserCreateResponse struct {
	value *BeaconUserCreateResponse
	isSet bool
}

func (v NullableBeaconUserCreateResponse) Get() *BeaconUserCreateResponse {
	return v.value
}

func (v *NullableBeaconUserCreateResponse) Set(val *BeaconUserCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBeaconUserCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBeaconUserCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeaconUserCreateResponse(val *BeaconUserCreateResponse) *NullableBeaconUserCreateResponse {
	return &NullableBeaconUserCreateResponse{value: val, isSet: true}
}

func (v NullableBeaconUserCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeaconUserCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


