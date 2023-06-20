/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.379.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"time"
)

// WatchlistScreeningHit Data from a government watchlist or PEP list that has been attached to the screening.
type WatchlistScreeningHit struct {
	// ID of the associated screening hit.
	Id string `json:"id"`
	ReviewStatus WatchlistScreeningHitStatus `json:"review_status"`
	// An ISO8601 formatted timestamp.
	FirstActive time.Time `json:"first_active"`
	// An ISO8601 formatted timestamp.
	InactiveSince NullableTime `json:"inactive_since"`
	// An ISO8601 formatted timestamp.
	HistoricalSince NullableTime `json:"historical_since"`
	ListCode IndividualWatchlistCode `json:"list_code"`
	// A universal identifier for a watchlist individual that is stable across searches and updates.
	PlaidUid string `json:"plaid_uid"`
	// The identifier provided by the source sanction or watchlist. When one is not provided by the source, this is `null`.
	SourceUid NullableString `json:"source_uid"`
	Analysis *ScreeningHitAnalysis `json:"analysis,omitempty"`
	Data *ScreeningHitData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WatchlistScreeningHit WatchlistScreeningHit

// NewWatchlistScreeningHit instantiates a new WatchlistScreeningHit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningHit(id string, reviewStatus WatchlistScreeningHitStatus, firstActive time.Time, inactiveSince NullableTime, historicalSince NullableTime, listCode IndividualWatchlistCode, plaidUid string, sourceUid NullableString) *WatchlistScreeningHit {
	this := WatchlistScreeningHit{}
	this.Id = id
	this.ReviewStatus = reviewStatus
	this.FirstActive = firstActive
	this.InactiveSince = inactiveSince
	this.HistoricalSince = historicalSince
	this.ListCode = listCode
	this.PlaidUid = plaidUid
	this.SourceUid = sourceUid
	return &this
}

// NewWatchlistScreeningHitWithDefaults instantiates a new WatchlistScreeningHit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningHitWithDefaults() *WatchlistScreeningHit {
	this := WatchlistScreeningHit{}
	return &this
}

// GetId returns the Id field value
func (o *WatchlistScreeningHit) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningHit) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WatchlistScreeningHit) SetId(v string) {
	o.Id = v
}

// GetReviewStatus returns the ReviewStatus field value
func (o *WatchlistScreeningHit) GetReviewStatus() WatchlistScreeningHitStatus {
	if o == nil {
		var ret WatchlistScreeningHitStatus
		return ret
	}

	return o.ReviewStatus
}

// GetReviewStatusOk returns a tuple with the ReviewStatus field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningHit) GetReviewStatusOk() (*WatchlistScreeningHitStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReviewStatus, true
}

// SetReviewStatus sets field value
func (o *WatchlistScreeningHit) SetReviewStatus(v WatchlistScreeningHitStatus) {
	o.ReviewStatus = v
}

// GetFirstActive returns the FirstActive field value
func (o *WatchlistScreeningHit) GetFirstActive() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.FirstActive
}

// GetFirstActiveOk returns a tuple with the FirstActive field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningHit) GetFirstActiveOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FirstActive, true
}

// SetFirstActive sets field value
func (o *WatchlistScreeningHit) SetFirstActive(v time.Time) {
	o.FirstActive = v
}

// GetInactiveSince returns the InactiveSince field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *WatchlistScreeningHit) GetInactiveSince() time.Time {
	if o == nil || o.InactiveSince.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.InactiveSince.Get()
}

// GetInactiveSinceOk returns a tuple with the InactiveSince field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningHit) GetInactiveSinceOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InactiveSince.Get(), o.InactiveSince.IsSet()
}

// SetInactiveSince sets field value
func (o *WatchlistScreeningHit) SetInactiveSince(v time.Time) {
	o.InactiveSince.Set(&v)
}

// GetHistoricalSince returns the HistoricalSince field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *WatchlistScreeningHit) GetHistoricalSince() time.Time {
	if o == nil || o.HistoricalSince.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.HistoricalSince.Get()
}

// GetHistoricalSinceOk returns a tuple with the HistoricalSince field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningHit) GetHistoricalSinceOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HistoricalSince.Get(), o.HistoricalSince.IsSet()
}

// SetHistoricalSince sets field value
func (o *WatchlistScreeningHit) SetHistoricalSince(v time.Time) {
	o.HistoricalSince.Set(&v)
}

// GetListCode returns the ListCode field value
func (o *WatchlistScreeningHit) GetListCode() IndividualWatchlistCode {
	if o == nil {
		var ret IndividualWatchlistCode
		return ret
	}

	return o.ListCode
}

// GetListCodeOk returns a tuple with the ListCode field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningHit) GetListCodeOk() (*IndividualWatchlistCode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ListCode, true
}

// SetListCode sets field value
func (o *WatchlistScreeningHit) SetListCode(v IndividualWatchlistCode) {
	o.ListCode = v
}

// GetPlaidUid returns the PlaidUid field value
func (o *WatchlistScreeningHit) GetPlaidUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlaidUid
}

// GetPlaidUidOk returns a tuple with the PlaidUid field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningHit) GetPlaidUidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PlaidUid, true
}

// SetPlaidUid sets field value
func (o *WatchlistScreeningHit) SetPlaidUid(v string) {
	o.PlaidUid = v
}

// GetSourceUid returns the SourceUid field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchlistScreeningHit) GetSourceUid() string {
	if o == nil || o.SourceUid.Get() == nil {
		var ret string
		return ret
	}

	return *o.SourceUid.Get()
}

// GetSourceUidOk returns a tuple with the SourceUid field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningHit) GetSourceUidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceUid.Get(), o.SourceUid.IsSet()
}

// SetSourceUid sets field value
func (o *WatchlistScreeningHit) SetSourceUid(v string) {
	o.SourceUid.Set(&v)
}

// GetAnalysis returns the Analysis field value if set, zero value otherwise.
func (o *WatchlistScreeningHit) GetAnalysis() ScreeningHitAnalysis {
	if o == nil || o.Analysis == nil {
		var ret ScreeningHitAnalysis
		return ret
	}
	return *o.Analysis
}

// GetAnalysisOk returns a tuple with the Analysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningHit) GetAnalysisOk() (*ScreeningHitAnalysis, bool) {
	if o == nil || o.Analysis == nil {
		return nil, false
	}
	return o.Analysis, true
}

// HasAnalysis returns a boolean if a field has been set.
func (o *WatchlistScreeningHit) HasAnalysis() bool {
	if o != nil && o.Analysis != nil {
		return true
	}

	return false
}

// SetAnalysis gets a reference to the given ScreeningHitAnalysis and assigns it to the Analysis field.
func (o *WatchlistScreeningHit) SetAnalysis(v ScreeningHitAnalysis) {
	o.Analysis = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *WatchlistScreeningHit) GetData() ScreeningHitData {
	if o == nil || o.Data == nil {
		var ret ScreeningHitData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningHit) GetDataOk() (*ScreeningHitData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *WatchlistScreeningHit) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ScreeningHitData and assigns it to the Data field.
func (o *WatchlistScreeningHit) SetData(v ScreeningHitData) {
	o.Data = &v
}

func (o WatchlistScreeningHit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["review_status"] = o.ReviewStatus
	}
	if true {
		toSerialize["first_active"] = o.FirstActive
	}
	if true {
		toSerialize["inactive_since"] = o.InactiveSince.Get()
	}
	if true {
		toSerialize["historical_since"] = o.HistoricalSince.Get()
	}
	if true {
		toSerialize["list_code"] = o.ListCode
	}
	if true {
		toSerialize["plaid_uid"] = o.PlaidUid
	}
	if true {
		toSerialize["source_uid"] = o.SourceUid.Get()
	}
	if o.Analysis != nil {
		toSerialize["analysis"] = o.Analysis
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WatchlistScreeningHit) UnmarshalJSON(bytes []byte) (err error) {
	varWatchlistScreeningHit := _WatchlistScreeningHit{}

	if err = json.Unmarshal(bytes, &varWatchlistScreeningHit); err == nil {
		*o = WatchlistScreeningHit(varWatchlistScreeningHit)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "review_status")
		delete(additionalProperties, "first_active")
		delete(additionalProperties, "inactive_since")
		delete(additionalProperties, "historical_since")
		delete(additionalProperties, "list_code")
		delete(additionalProperties, "plaid_uid")
		delete(additionalProperties, "source_uid")
		delete(additionalProperties, "analysis")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWatchlistScreeningHit struct {
	value *WatchlistScreeningHit
	isSet bool
}

func (v NullableWatchlistScreeningHit) Get() *WatchlistScreeningHit {
	return v.value
}

func (v *NullableWatchlistScreeningHit) Set(val *WatchlistScreeningHit) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningHit) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningHit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningHit(val *WatchlistScreeningHit) *NullableWatchlistScreeningHit {
	return &NullableWatchlistScreeningHit{value: val, isSet: true}
}

func (v NullableWatchlistScreeningHit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningHit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


