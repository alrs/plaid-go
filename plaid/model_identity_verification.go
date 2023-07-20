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
	"time"
)

// IdentityVerification A identity verification attempt represents a customer's attempt to verify their identity, reflecting the required steps for completing the session, the results for each step, and information collected in the process.
type IdentityVerification struct {
	// ID of the associated Identity Verification attempt.
	Id string `json:"id"`
	// A unique ID that identifies the end user in your system. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the Link Token Create `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.
	ClientUserId string `json:"client_user_id"`
	// An ISO8601 formatted timestamp.
	CreatedAt time.Time `json:"created_at"`
	// An ISO8601 formatted timestamp.
	CompletedAt NullableTime `json:"completed_at"`
	// The ID for the Identity Verification preceding this session. This field will only be filled if the current Identity Verification is a retry of a previous attempt.
	PreviousAttemptId NullableString `json:"previous_attempt_id"`
	// A shareable URL that can be sent directly to the user to complete verification
	ShareableUrl NullableString `json:"shareable_url"`
	Template IdentityVerificationTemplateReference `json:"template"`
	User IdentityVerificationUserData `json:"user"`
	Status IdentityVerificationStatus `json:"status"`
	Steps IdentityVerificationStepSummary `json:"steps"`
	DocumentaryVerification NullableDocumentaryVerification `json:"documentary_verification"`
	SelfieCheck NullableSelfieCheck `json:"selfie_check"`
	KycCheck NullableKYCCheckDetails `json:"kyc_check"`
	RiskCheck NullableRiskCheckDetails `json:"risk_check"`
	// ID of the associated screening.
	WatchlistScreeningId NullableString `json:"watchlist_screening_id"`
	// An ISO8601 formatted timestamp.
	RedactedAt NullableTime `json:"redacted_at"`
	AdditionalProperties map[string]interface{}
}

type _IdentityVerification IdentityVerification

// NewIdentityVerification instantiates a new IdentityVerification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityVerification(id string, clientUserId string, createdAt time.Time, completedAt NullableTime, previousAttemptId NullableString, shareableUrl NullableString, template IdentityVerificationTemplateReference, user IdentityVerificationUserData, status IdentityVerificationStatus, steps IdentityVerificationStepSummary, documentaryVerification NullableDocumentaryVerification, selfieCheck NullableSelfieCheck, kycCheck NullableKYCCheckDetails, riskCheck NullableRiskCheckDetails, watchlistScreeningId NullableString, redactedAt NullableTime) *IdentityVerification {
	this := IdentityVerification{}
	this.Id = id
	this.ClientUserId = clientUserId
	this.CreatedAt = createdAt
	this.CompletedAt = completedAt
	this.PreviousAttemptId = previousAttemptId
	this.ShareableUrl = shareableUrl
	this.Template = template
	this.User = user
	this.Status = status
	this.Steps = steps
	this.DocumentaryVerification = documentaryVerification
	this.SelfieCheck = selfieCheck
	this.KycCheck = kycCheck
	this.RiskCheck = riskCheck
	this.WatchlistScreeningId = watchlistScreeningId
	this.RedactedAt = redactedAt
	return &this
}

// NewIdentityVerificationWithDefaults instantiates a new IdentityVerification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityVerificationWithDefaults() *IdentityVerification {
	this := IdentityVerification{}
	return &this
}

// GetId returns the Id field value
func (o *IdentityVerification) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IdentityVerification) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IdentityVerification) SetId(v string) {
	o.Id = v
}

// GetClientUserId returns the ClientUserId field value
func (o *IdentityVerification) GetClientUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientUserId
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value
// and a boolean to check if the value has been set.
func (o *IdentityVerification) GetClientUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientUserId, true
}

// SetClientUserId sets field value
func (o *IdentityVerification) SetClientUserId(v string) {
	o.ClientUserId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *IdentityVerification) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *IdentityVerification) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *IdentityVerification) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCompletedAt returns the CompletedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *IdentityVerification) GetCompletedAt() time.Time {
	if o == nil || o.CompletedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerification) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// SetCompletedAt sets field value
func (o *IdentityVerification) SetCompletedAt(v time.Time) {
	o.CompletedAt.Set(&v)
}

// GetPreviousAttemptId returns the PreviousAttemptId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IdentityVerification) GetPreviousAttemptId() string {
	if o == nil || o.PreviousAttemptId.Get() == nil {
		var ret string
		return ret
	}

	return *o.PreviousAttemptId.Get()
}

// GetPreviousAttemptIdOk returns a tuple with the PreviousAttemptId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerification) GetPreviousAttemptIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PreviousAttemptId.Get(), o.PreviousAttemptId.IsSet()
}

// SetPreviousAttemptId sets field value
func (o *IdentityVerification) SetPreviousAttemptId(v string) {
	o.PreviousAttemptId.Set(&v)
}

// GetShareableUrl returns the ShareableUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IdentityVerification) GetShareableUrl() string {
	if o == nil || o.ShareableUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.ShareableUrl.Get()
}

// GetShareableUrlOk returns a tuple with the ShareableUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerification) GetShareableUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ShareableUrl.Get(), o.ShareableUrl.IsSet()
}

// SetShareableUrl sets field value
func (o *IdentityVerification) SetShareableUrl(v string) {
	o.ShareableUrl.Set(&v)
}

// GetTemplate returns the Template field value
func (o *IdentityVerification) GetTemplate() IdentityVerificationTemplateReference {
	if o == nil {
		var ret IdentityVerificationTemplateReference
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *IdentityVerification) GetTemplateOk() (*IdentityVerificationTemplateReference, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *IdentityVerification) SetTemplate(v IdentityVerificationTemplateReference) {
	o.Template = v
}

// GetUser returns the User field value
func (o *IdentityVerification) GetUser() IdentityVerificationUserData {
	if o == nil {
		var ret IdentityVerificationUserData
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *IdentityVerification) GetUserOk() (*IdentityVerificationUserData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *IdentityVerification) SetUser(v IdentityVerificationUserData) {
	o.User = v
}

// GetStatus returns the Status field value
func (o *IdentityVerification) GetStatus() IdentityVerificationStatus {
	if o == nil {
		var ret IdentityVerificationStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *IdentityVerification) GetStatusOk() (*IdentityVerificationStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *IdentityVerification) SetStatus(v IdentityVerificationStatus) {
	o.Status = v
}

// GetSteps returns the Steps field value
func (o *IdentityVerification) GetSteps() IdentityVerificationStepSummary {
	if o == nil {
		var ret IdentityVerificationStepSummary
		return ret
	}

	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value
// and a boolean to check if the value has been set.
func (o *IdentityVerification) GetStepsOk() (*IdentityVerificationStepSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Steps, true
}

// SetSteps sets field value
func (o *IdentityVerification) SetSteps(v IdentityVerificationStepSummary) {
	o.Steps = v
}

// GetDocumentaryVerification returns the DocumentaryVerification field value
// If the value is explicit nil, the zero value for DocumentaryVerification will be returned
func (o *IdentityVerification) GetDocumentaryVerification() DocumentaryVerification {
	if o == nil || o.DocumentaryVerification.Get() == nil {
		var ret DocumentaryVerification
		return ret
	}

	return *o.DocumentaryVerification.Get()
}

// GetDocumentaryVerificationOk returns a tuple with the DocumentaryVerification field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerification) GetDocumentaryVerificationOk() (*DocumentaryVerification, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DocumentaryVerification.Get(), o.DocumentaryVerification.IsSet()
}

// SetDocumentaryVerification sets field value
func (o *IdentityVerification) SetDocumentaryVerification(v DocumentaryVerification) {
	o.DocumentaryVerification.Set(&v)
}

// GetSelfieCheck returns the SelfieCheck field value
// If the value is explicit nil, the zero value for SelfieCheck will be returned
func (o *IdentityVerification) GetSelfieCheck() SelfieCheck {
	if o == nil || o.SelfieCheck.Get() == nil {
		var ret SelfieCheck
		return ret
	}

	return *o.SelfieCheck.Get()
}

// GetSelfieCheckOk returns a tuple with the SelfieCheck field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerification) GetSelfieCheckOk() (*SelfieCheck, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SelfieCheck.Get(), o.SelfieCheck.IsSet()
}

// SetSelfieCheck sets field value
func (o *IdentityVerification) SetSelfieCheck(v SelfieCheck) {
	o.SelfieCheck.Set(&v)
}

// GetKycCheck returns the KycCheck field value
// If the value is explicit nil, the zero value for KYCCheckDetails will be returned
func (o *IdentityVerification) GetKycCheck() KYCCheckDetails {
	if o == nil || o.KycCheck.Get() == nil {
		var ret KYCCheckDetails
		return ret
	}

	return *o.KycCheck.Get()
}

// GetKycCheckOk returns a tuple with the KycCheck field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerification) GetKycCheckOk() (*KYCCheckDetails, bool) {
	if o == nil  {
		return nil, false
	}
	return o.KycCheck.Get(), o.KycCheck.IsSet()
}

// SetKycCheck sets field value
func (o *IdentityVerification) SetKycCheck(v KYCCheckDetails) {
	o.KycCheck.Set(&v)
}

// GetRiskCheck returns the RiskCheck field value
// If the value is explicit nil, the zero value for RiskCheckDetails will be returned
func (o *IdentityVerification) GetRiskCheck() RiskCheckDetails {
	if o == nil || o.RiskCheck.Get() == nil {
		var ret RiskCheckDetails
		return ret
	}

	return *o.RiskCheck.Get()
}

// GetRiskCheckOk returns a tuple with the RiskCheck field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerification) GetRiskCheckOk() (*RiskCheckDetails, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RiskCheck.Get(), o.RiskCheck.IsSet()
}

// SetRiskCheck sets field value
func (o *IdentityVerification) SetRiskCheck(v RiskCheckDetails) {
	o.RiskCheck.Set(&v)
}

// GetWatchlistScreeningId returns the WatchlistScreeningId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IdentityVerification) GetWatchlistScreeningId() string {
	if o == nil || o.WatchlistScreeningId.Get() == nil {
		var ret string
		return ret
	}

	return *o.WatchlistScreeningId.Get()
}

// GetWatchlistScreeningIdOk returns a tuple with the WatchlistScreeningId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerification) GetWatchlistScreeningIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WatchlistScreeningId.Get(), o.WatchlistScreeningId.IsSet()
}

// SetWatchlistScreeningId sets field value
func (o *IdentityVerification) SetWatchlistScreeningId(v string) {
	o.WatchlistScreeningId.Set(&v)
}

// GetRedactedAt returns the RedactedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *IdentityVerification) GetRedactedAt() time.Time {
	if o == nil || o.RedactedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.RedactedAt.Get()
}

// GetRedactedAtOk returns a tuple with the RedactedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerification) GetRedactedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RedactedAt.Get(), o.RedactedAt.IsSet()
}

// SetRedactedAt sets field value
func (o *IdentityVerification) SetRedactedAt(v time.Time) {
	o.RedactedAt.Set(&v)
}

func (o IdentityVerification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["client_user_id"] = o.ClientUserId
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["completed_at"] = o.CompletedAt.Get()
	}
	if true {
		toSerialize["previous_attempt_id"] = o.PreviousAttemptId.Get()
	}
	if true {
		toSerialize["shareable_url"] = o.ShareableUrl.Get()
	}
	if true {
		toSerialize["template"] = o.Template
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["steps"] = o.Steps
	}
	if true {
		toSerialize["documentary_verification"] = o.DocumentaryVerification.Get()
	}
	if true {
		toSerialize["selfie_check"] = o.SelfieCheck.Get()
	}
	if true {
		toSerialize["kyc_check"] = o.KycCheck.Get()
	}
	if true {
		toSerialize["risk_check"] = o.RiskCheck.Get()
	}
	if true {
		toSerialize["watchlist_screening_id"] = o.WatchlistScreeningId.Get()
	}
	if true {
		toSerialize["redacted_at"] = o.RedactedAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityVerification) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityVerification := _IdentityVerification{}

	if err = json.Unmarshal(bytes, &varIdentityVerification); err == nil {
		*o = IdentityVerification(varIdentityVerification)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "client_user_id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "completed_at")
		delete(additionalProperties, "previous_attempt_id")
		delete(additionalProperties, "shareable_url")
		delete(additionalProperties, "template")
		delete(additionalProperties, "user")
		delete(additionalProperties, "status")
		delete(additionalProperties, "steps")
		delete(additionalProperties, "documentary_verification")
		delete(additionalProperties, "selfie_check")
		delete(additionalProperties, "kyc_check")
		delete(additionalProperties, "risk_check")
		delete(additionalProperties, "watchlist_screening_id")
		delete(additionalProperties, "redacted_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityVerification struct {
	value *IdentityVerification
	isSet bool
}

func (v NullableIdentityVerification) Get() *IdentityVerification {
	return v.value
}

func (v *NullableIdentityVerification) Set(val *IdentityVerification) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityVerification(val *IdentityVerification) *NullableIdentityVerification {
	return &NullableIdentityVerification{value: val, isSet: true}
}

func (v NullableIdentityVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


