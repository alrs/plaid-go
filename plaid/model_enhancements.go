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
)

// Enhancements A grouping of the Plaid produced transaction enhancement fields.
type Enhancements struct {
	// The name of the primary counterparty, such as the merchant or the financial institution, as extracted by Plaid from the raw description.
	MerchantName NullableString `json:"merchant_name,omitempty"`
	// The website associated with this transaction, if available.
	Website NullableString `json:"website,omitempty"`
	// The URL of a logo associated with this transaction, if available. The logo is formatted as a 100x100 pixel PNG file.
	LogoUrl NullableString `json:"logo_url,omitempty"`
	// The check number of the transaction. This field is only populated for check transactions.
	CheckNumber NullableString `json:"check_number,omitempty"`
	PaymentChannel PaymentChannel `json:"payment_channel"`
	// The ID of the category to which this transaction belongs. For a full list of categories, see [`/categories/get`](https://plaid.com/docs/api/products/transactions/#categoriesget).
	CategoryId NullableString `json:"category_id"`
	// A hierarchical array of the categories to which this transaction belongs. For a full list of categories, see [`/categories/get`](https://plaid.com/docs/api/products/transactions/#categoriesget).
	Category []string `json:"category"`
	Location Location `json:"location"`
	PersonalFinanceCategory NullablePersonalFinanceCategory `json:"personal_finance_category,omitempty"`
	// A link to the icon associated with the primary personal finance category. The logo will always be 100x100 pixels.
	PersonalFinanceCategoryIconUrl *string `json:"personal_finance_category_icon_url,omitempty"`
	// The counterparties present in the transaction. Counterparties, such as the merchant or the financial institution, are extracted by Plaid from the raw description.
	Counterparties *[]Counterparty `json:"counterparties,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Enhancements Enhancements

// NewEnhancements instantiates a new Enhancements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnhancements(paymentChannel PaymentChannel, categoryId NullableString, category []string, location Location) *Enhancements {
	this := Enhancements{}
	this.PaymentChannel = paymentChannel
	this.CategoryId = categoryId
	this.Category = category
	this.Location = location
	return &this
}

// NewEnhancementsWithDefaults instantiates a new Enhancements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnhancementsWithDefaults() *Enhancements {
	this := Enhancements{}
	return &this
}

// GetMerchantName returns the MerchantName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Enhancements) GetMerchantName() string {
	if o == nil || o.MerchantName.Get() == nil {
		var ret string
		return ret
	}
	return *o.MerchantName.Get()
}

// GetMerchantNameOk returns a tuple with the MerchantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Enhancements) GetMerchantNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MerchantName.Get(), o.MerchantName.IsSet()
}

// HasMerchantName returns a boolean if a field has been set.
func (o *Enhancements) HasMerchantName() bool {
	if o != nil && o.MerchantName.IsSet() {
		return true
	}

	return false
}

// SetMerchantName gets a reference to the given NullableString and assigns it to the MerchantName field.
func (o *Enhancements) SetMerchantName(v string) {
	o.MerchantName.Set(&v)
}
// SetMerchantNameNil sets the value for MerchantName to be an explicit nil
func (o *Enhancements) SetMerchantNameNil() {
	o.MerchantName.Set(nil)
}

// UnsetMerchantName ensures that no value is present for MerchantName, not even an explicit nil
func (o *Enhancements) UnsetMerchantName() {
	o.MerchantName.Unset()
}

// GetWebsite returns the Website field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Enhancements) GetWebsite() string {
	if o == nil || o.Website.Get() == nil {
		var ret string
		return ret
	}
	return *o.Website.Get()
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Enhancements) GetWebsiteOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Website.Get(), o.Website.IsSet()
}

// HasWebsite returns a boolean if a field has been set.
func (o *Enhancements) HasWebsite() bool {
	if o != nil && o.Website.IsSet() {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given NullableString and assigns it to the Website field.
func (o *Enhancements) SetWebsite(v string) {
	o.Website.Set(&v)
}
// SetWebsiteNil sets the value for Website to be an explicit nil
func (o *Enhancements) SetWebsiteNil() {
	o.Website.Set(nil)
}

// UnsetWebsite ensures that no value is present for Website, not even an explicit nil
func (o *Enhancements) UnsetWebsite() {
	o.Website.Unset()
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Enhancements) GetLogoUrl() string {
	if o == nil || o.LogoUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.LogoUrl.Get()
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Enhancements) GetLogoUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LogoUrl.Get(), o.LogoUrl.IsSet()
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *Enhancements) HasLogoUrl() bool {
	if o != nil && o.LogoUrl.IsSet() {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given NullableString and assigns it to the LogoUrl field.
func (o *Enhancements) SetLogoUrl(v string) {
	o.LogoUrl.Set(&v)
}
// SetLogoUrlNil sets the value for LogoUrl to be an explicit nil
func (o *Enhancements) SetLogoUrlNil() {
	o.LogoUrl.Set(nil)
}

// UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
func (o *Enhancements) UnsetLogoUrl() {
	o.LogoUrl.Unset()
}

// GetCheckNumber returns the CheckNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Enhancements) GetCheckNumber() string {
	if o == nil || o.CheckNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.CheckNumber.Get()
}

// GetCheckNumberOk returns a tuple with the CheckNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Enhancements) GetCheckNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CheckNumber.Get(), o.CheckNumber.IsSet()
}

// HasCheckNumber returns a boolean if a field has been set.
func (o *Enhancements) HasCheckNumber() bool {
	if o != nil && o.CheckNumber.IsSet() {
		return true
	}

	return false
}

// SetCheckNumber gets a reference to the given NullableString and assigns it to the CheckNumber field.
func (o *Enhancements) SetCheckNumber(v string) {
	o.CheckNumber.Set(&v)
}
// SetCheckNumberNil sets the value for CheckNumber to be an explicit nil
func (o *Enhancements) SetCheckNumberNil() {
	o.CheckNumber.Set(nil)
}

// UnsetCheckNumber ensures that no value is present for CheckNumber, not even an explicit nil
func (o *Enhancements) UnsetCheckNumber() {
	o.CheckNumber.Unset()
}

// GetPaymentChannel returns the PaymentChannel field value
func (o *Enhancements) GetPaymentChannel() PaymentChannel {
	if o == nil {
		var ret PaymentChannel
		return ret
	}

	return o.PaymentChannel
}

// GetPaymentChannelOk returns a tuple with the PaymentChannel field value
// and a boolean to check if the value has been set.
func (o *Enhancements) GetPaymentChannelOk() (*PaymentChannel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentChannel, true
}

// SetPaymentChannel sets field value
func (o *Enhancements) SetPaymentChannel(v PaymentChannel) {
	o.PaymentChannel = v
}

// GetCategoryId returns the CategoryId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Enhancements) GetCategoryId() string {
	if o == nil || o.CategoryId.Get() == nil {
		var ret string
		return ret
	}

	return *o.CategoryId.Get()
}

// GetCategoryIdOk returns a tuple with the CategoryId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Enhancements) GetCategoryIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CategoryId.Get(), o.CategoryId.IsSet()
}

// SetCategoryId sets field value
func (o *Enhancements) SetCategoryId(v string) {
	o.CategoryId.Set(&v)
}

// GetCategory returns the Category field value
func (o *Enhancements) GetCategory() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *Enhancements) GetCategoryOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *Enhancements) SetCategory(v []string) {
	o.Category = v
}

// GetLocation returns the Location field value
func (o *Enhancements) GetLocation() Location {
	if o == nil {
		var ret Location
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *Enhancements) GetLocationOk() (*Location, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *Enhancements) SetLocation(v Location) {
	o.Location = v
}

// GetPersonalFinanceCategory returns the PersonalFinanceCategory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Enhancements) GetPersonalFinanceCategory() PersonalFinanceCategory {
	if o == nil || o.PersonalFinanceCategory.Get() == nil {
		var ret PersonalFinanceCategory
		return ret
	}
	return *o.PersonalFinanceCategory.Get()
}

// GetPersonalFinanceCategoryOk returns a tuple with the PersonalFinanceCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Enhancements) GetPersonalFinanceCategoryOk() (*PersonalFinanceCategory, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PersonalFinanceCategory.Get(), o.PersonalFinanceCategory.IsSet()
}

// HasPersonalFinanceCategory returns a boolean if a field has been set.
func (o *Enhancements) HasPersonalFinanceCategory() bool {
	if o != nil && o.PersonalFinanceCategory.IsSet() {
		return true
	}

	return false
}

// SetPersonalFinanceCategory gets a reference to the given NullablePersonalFinanceCategory and assigns it to the PersonalFinanceCategory field.
func (o *Enhancements) SetPersonalFinanceCategory(v PersonalFinanceCategory) {
	o.PersonalFinanceCategory.Set(&v)
}
// SetPersonalFinanceCategoryNil sets the value for PersonalFinanceCategory to be an explicit nil
func (o *Enhancements) SetPersonalFinanceCategoryNil() {
	o.PersonalFinanceCategory.Set(nil)
}

// UnsetPersonalFinanceCategory ensures that no value is present for PersonalFinanceCategory, not even an explicit nil
func (o *Enhancements) UnsetPersonalFinanceCategory() {
	o.PersonalFinanceCategory.Unset()
}

// GetPersonalFinanceCategoryIconUrl returns the PersonalFinanceCategoryIconUrl field value if set, zero value otherwise.
func (o *Enhancements) GetPersonalFinanceCategoryIconUrl() string {
	if o == nil || o.PersonalFinanceCategoryIconUrl == nil {
		var ret string
		return ret
	}
	return *o.PersonalFinanceCategoryIconUrl
}

// GetPersonalFinanceCategoryIconUrlOk returns a tuple with the PersonalFinanceCategoryIconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Enhancements) GetPersonalFinanceCategoryIconUrlOk() (*string, bool) {
	if o == nil || o.PersonalFinanceCategoryIconUrl == nil {
		return nil, false
	}
	return o.PersonalFinanceCategoryIconUrl, true
}

// HasPersonalFinanceCategoryIconUrl returns a boolean if a field has been set.
func (o *Enhancements) HasPersonalFinanceCategoryIconUrl() bool {
	if o != nil && o.PersonalFinanceCategoryIconUrl != nil {
		return true
	}

	return false
}

// SetPersonalFinanceCategoryIconUrl gets a reference to the given string and assigns it to the PersonalFinanceCategoryIconUrl field.
func (o *Enhancements) SetPersonalFinanceCategoryIconUrl(v string) {
	o.PersonalFinanceCategoryIconUrl = &v
}

// GetCounterparties returns the Counterparties field value if set, zero value otherwise.
func (o *Enhancements) GetCounterparties() []Counterparty {
	if o == nil || o.Counterparties == nil {
		var ret []Counterparty
		return ret
	}
	return *o.Counterparties
}

// GetCounterpartiesOk returns a tuple with the Counterparties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Enhancements) GetCounterpartiesOk() (*[]Counterparty, bool) {
	if o == nil || o.Counterparties == nil {
		return nil, false
	}
	return o.Counterparties, true
}

// HasCounterparties returns a boolean if a field has been set.
func (o *Enhancements) HasCounterparties() bool {
	if o != nil && o.Counterparties != nil {
		return true
	}

	return false
}

// SetCounterparties gets a reference to the given []Counterparty and assigns it to the Counterparties field.
func (o *Enhancements) SetCounterparties(v []Counterparty) {
	o.Counterparties = &v
}

func (o Enhancements) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MerchantName.IsSet() {
		toSerialize["merchant_name"] = o.MerchantName.Get()
	}
	if o.Website.IsSet() {
		toSerialize["website"] = o.Website.Get()
	}
	if o.LogoUrl.IsSet() {
		toSerialize["logo_url"] = o.LogoUrl.Get()
	}
	if o.CheckNumber.IsSet() {
		toSerialize["check_number"] = o.CheckNumber.Get()
	}
	if true {
		toSerialize["payment_channel"] = o.PaymentChannel
	}
	if true {
		toSerialize["category_id"] = o.CategoryId.Get()
	}
	if true {
		toSerialize["category"] = o.Category
	}
	if true {
		toSerialize["location"] = o.Location
	}
	if o.PersonalFinanceCategory.IsSet() {
		toSerialize["personal_finance_category"] = o.PersonalFinanceCategory.Get()
	}
	if o.PersonalFinanceCategoryIconUrl != nil {
		toSerialize["personal_finance_category_icon_url"] = o.PersonalFinanceCategoryIconUrl
	}
	if o.Counterparties != nil {
		toSerialize["counterparties"] = o.Counterparties
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Enhancements) UnmarshalJSON(bytes []byte) (err error) {
	varEnhancements := _Enhancements{}

	if err = json.Unmarshal(bytes, &varEnhancements); err == nil {
		*o = Enhancements(varEnhancements)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "merchant_name")
		delete(additionalProperties, "website")
		delete(additionalProperties, "logo_url")
		delete(additionalProperties, "check_number")
		delete(additionalProperties, "payment_channel")
		delete(additionalProperties, "category_id")
		delete(additionalProperties, "category")
		delete(additionalProperties, "location")
		delete(additionalProperties, "personal_finance_category")
		delete(additionalProperties, "personal_finance_category_icon_url")
		delete(additionalProperties, "counterparties")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnhancements struct {
	value *Enhancements
	isSet bool
}

func (v NullableEnhancements) Get() *Enhancements {
	return v.value
}

func (v *NullableEnhancements) Set(val *Enhancements) {
	v.value = val
	v.isSet = true
}

func (v NullableEnhancements) IsSet() bool {
	return v.isSet
}

func (v *NullableEnhancements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnhancements(val *Enhancements) *NullableEnhancements {
	return &NullableEnhancements{value: val, isSet: true}
}

func (v NullableEnhancements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnhancements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


