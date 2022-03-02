/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.79.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// Holding A securities holding at an institution.
type Holding struct {
	// The Plaid `account_id` associated with the holding.
	AccountId string `json:"account_id"`
	// The Plaid `security_id` associated with the holding.
	SecurityId string `json:"security_id"`
	// The last price given by the institution for this security.
	InstitutionPrice float64 `json:"institution_price"`
	// The date at which `institution_price` was current.
	InstitutionPriceAsOf NullableString `json:"institution_price_as_of"`
	// The value of the holding, as reported by the institution.
	InstitutionValue float64 `json:"institution_value"`
	// The cost basis of the holding.
	CostBasis NullableFloat64 `json:"cost_basis"`
	// The total quantity of the asset held, as reported by the financial institution. If the security is an option, `quantity` will reflect the total number of options (typically the number of contracts multiplied by 100), not the number of contracts.
	Quantity float64 `json:"quantity"`
	// The ISO-4217 currency code of the holding. Always `null` if `unofficial_currency_code` is non-`null`.
	IsoCurrencyCode NullableString `json:"iso_currency_code"`
	// The unofficial currency code associated with the holding. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s. 
	UnofficialCurrencyCode NullableString `json:"unofficial_currency_code"`
	AdditionalProperties map[string]interface{}
}

type _Holding Holding

// NewHolding instantiates a new Holding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHolding(accountId string, securityId string, institutionPrice float64, institutionPriceAsOf NullableString, institutionValue float64, costBasis NullableFloat64, quantity float64, isoCurrencyCode NullableString, unofficialCurrencyCode NullableString) *Holding {
	this := Holding{}
	this.AccountId = accountId
	this.SecurityId = securityId
	this.InstitutionPrice = institutionPrice
	this.InstitutionPriceAsOf = institutionPriceAsOf
	this.InstitutionValue = institutionValue
	this.CostBasis = costBasis
	this.Quantity = quantity
	this.IsoCurrencyCode = isoCurrencyCode
	this.UnofficialCurrencyCode = unofficialCurrencyCode
	return &this
}

// NewHoldingWithDefaults instantiates a new Holding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHoldingWithDefaults() *Holding {
	this := Holding{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *Holding) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *Holding) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *Holding) SetAccountId(v string) {
	o.AccountId = v
}

// GetSecurityId returns the SecurityId field value
func (o *Holding) GetSecurityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecurityId
}

// GetSecurityIdOk returns a tuple with the SecurityId field value
// and a boolean to check if the value has been set.
func (o *Holding) GetSecurityIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SecurityId, true
}

// SetSecurityId sets field value
func (o *Holding) SetSecurityId(v string) {
	o.SecurityId = v
}

// GetInstitutionPrice returns the InstitutionPrice field value
func (o *Holding) GetInstitutionPrice() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.InstitutionPrice
}

// GetInstitutionPriceOk returns a tuple with the InstitutionPrice field value
// and a boolean to check if the value has been set.
func (o *Holding) GetInstitutionPriceOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstitutionPrice, true
}

// SetInstitutionPrice sets field value
func (o *Holding) SetInstitutionPrice(v float64) {
	o.InstitutionPrice = v
}

// GetInstitutionPriceAsOf returns the InstitutionPriceAsOf field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Holding) GetInstitutionPriceAsOf() string {
	if o == nil || o.InstitutionPriceAsOf.Get() == nil {
		var ret string
		return ret
	}

	return *o.InstitutionPriceAsOf.Get()
}

// GetInstitutionPriceAsOfOk returns a tuple with the InstitutionPriceAsOf field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Holding) GetInstitutionPriceAsOfOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InstitutionPriceAsOf.Get(), o.InstitutionPriceAsOf.IsSet()
}

// SetInstitutionPriceAsOf sets field value
func (o *Holding) SetInstitutionPriceAsOf(v string) {
	o.InstitutionPriceAsOf.Set(&v)
}

// GetInstitutionValue returns the InstitutionValue field value
func (o *Holding) GetInstitutionValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.InstitutionValue
}

// GetInstitutionValueOk returns a tuple with the InstitutionValue field value
// and a boolean to check if the value has been set.
func (o *Holding) GetInstitutionValueOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstitutionValue, true
}

// SetInstitutionValue sets field value
func (o *Holding) SetInstitutionValue(v float64) {
	o.InstitutionValue = v
}

// GetCostBasis returns the CostBasis field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *Holding) GetCostBasis() float64 {
	if o == nil || o.CostBasis.Get() == nil {
		var ret float64
		return ret
	}

	return *o.CostBasis.Get()
}

// GetCostBasisOk returns a tuple with the CostBasis field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Holding) GetCostBasisOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CostBasis.Get(), o.CostBasis.IsSet()
}

// SetCostBasis sets field value
func (o *Holding) SetCostBasis(v float64) {
	o.CostBasis.Set(&v)
}

// GetQuantity returns the Quantity field value
func (o *Holding) GetQuantity() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *Holding) GetQuantityOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *Holding) SetQuantity(v float64) {
	o.Quantity = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Holding) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.IsoCurrencyCode.Get()
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Holding) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsoCurrencyCode.Get(), o.IsoCurrencyCode.IsSet()
}

// SetIsoCurrencyCode sets field value
func (o *Holding) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode.Set(&v)
}

// GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Holding) GetUnofficialCurrencyCode() string {
	if o == nil || o.UnofficialCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.UnofficialCurrencyCode.Get()
}

// GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Holding) GetUnofficialCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnofficialCurrencyCode.Get(), o.UnofficialCurrencyCode.IsSet()
}

// SetUnofficialCurrencyCode sets field value
func (o *Holding) SetUnofficialCurrencyCode(v string) {
	o.UnofficialCurrencyCode.Set(&v)
}

func (o Holding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["security_id"] = o.SecurityId
	}
	if true {
		toSerialize["institution_price"] = o.InstitutionPrice
	}
	if true {
		toSerialize["institution_price_as_of"] = o.InstitutionPriceAsOf.Get()
	}
	if true {
		toSerialize["institution_value"] = o.InstitutionValue
	}
	if true {
		toSerialize["cost_basis"] = o.CostBasis.Get()
	}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode.Get()
	}
	if true {
		toSerialize["unofficial_currency_code"] = o.UnofficialCurrencyCode.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Holding) UnmarshalJSON(bytes []byte) (err error) {
	varHolding := _Holding{}

	if err = json.Unmarshal(bytes, &varHolding); err == nil {
		*o = Holding(varHolding)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "security_id")
		delete(additionalProperties, "institution_price")
		delete(additionalProperties, "institution_price_as_of")
		delete(additionalProperties, "institution_value")
		delete(additionalProperties, "cost_basis")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "unofficial_currency_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHolding struct {
	value *Holding
	isSet bool
}

func (v NullableHolding) Get() *Holding {
	return v.value
}

func (v *NullableHolding) Set(val *Holding) {
	v.value = val
	v.isSet = true
}

func (v NullableHolding) IsSet() bool {
	return v.isSet
}

func (v *NullableHolding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHolding(val *Holding) *NullableHolding {
	return &NullableHolding{value: val, isSet: true}
}

func (v NullableHolding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHolding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


