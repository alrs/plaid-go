/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.181.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"time"
)

// Security Contains details about a security
type Security struct {
	// A unique, Plaid-specific identifier for the security, used to associate securities with holdings. Like all Plaid identifiers, the `security_id` is case sensitive.
	SecurityId string `json:"security_id"`
	// 12-character ISIN, a globally unique securities identifier.
	Isin NullableString `json:"isin"`
	// 9-character CUSIP, an identifier assigned to North American securities.
	Cusip NullableString `json:"cusip"`
	// 7-character SEDOL, an identifier assigned to securities in the UK.
	Sedol NullableString `json:"sedol"`
	// An identifier given to the security by the institution
	InstitutionSecurityId NullableString `json:"institution_security_id"`
	// If `institution_security_id` is present, this field indicates the Plaid `institution_id` of the institution to whom the identifier belongs.
	InstitutionId NullableString `json:"institution_id"`
	// In certain cases, Plaid will provide the ID of another security whose performance resembles this security, typically when the original security has low volume, or when a private security can be modeled with a publicly traded security.
	ProxySecurityId NullableString `json:"proxy_security_id"`
	// A descriptive name for the security, suitable for display.
	Name NullableString `json:"name"`
	// The security’s trading symbol for publicly traded securities, and otherwise a short identifier if available.
	TickerSymbol NullableString `json:"ticker_symbol"`
	// Indicates that a security is a highly liquid asset and can be treated like cash.
	IsCashEquivalent NullableBool `json:"is_cash_equivalent"`
	// The security type of the holding. Valid security types are:  `cash`: Cash, currency, and money market funds  `cryptocurrency`: Digital or virtual currencies  `derivative`: Options, warrants, and other derivative instruments  `equity`: Domestic and foreign equities  `etf`: Multi-asset exchange-traded investment funds  `fixed income`: Bonds and certificates of deposit (CDs)  `loan`: Loans and loan receivables  `mutual fund`: Open- and closed-end vehicles pooling funds of multiple investors  `other`: Unknown or other investment types
	Type NullableString `json:"type"`
	// Price of the security at the close of the previous trading session. Null for non-public securities.  If the security is a foreign currency this field will be updated daily and will be priced in USD.  If the security is a cryptocurrency, this field will be updated multiple times a day. As crypto prices can fluctuate quickly and data may become stale sooner than other asset classes, please refer to update_datetime with the time when the price was last updated. 
	ClosePrice NullableFloat64 `json:"close_price"`
	// Date for which `close_price` is accurate. Always `null` if `close_price` is `null`.
	ClosePriceAsOf NullableString `json:"close_price_as_of"`
	// Date and time at which close_price is accurate, in ISO 8601 format (YYYY-MM-DDTHH:mm:ssZ). Always null if close_price is null.
	UpdateDatetime NullableTime `json:"update_datetime,omitempty"`
	// The ISO-4217 currency code of the price given. Always `null` if `unofficial_currency_code` is non-`null`.
	IsoCurrencyCode NullableString `json:"iso_currency_code"`
	// The unofficial currency code associated with the security. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.
	UnofficialCurrencyCode NullableString `json:"unofficial_currency_code"`
	AdditionalProperties map[string]interface{}
}

type _Security Security

// NewSecurity instantiates a new Security object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurity(securityId string, isin NullableString, cusip NullableString, sedol NullableString, institutionSecurityId NullableString, institutionId NullableString, proxySecurityId NullableString, name NullableString, tickerSymbol NullableString, isCashEquivalent NullableBool, type_ NullableString, closePrice NullableFloat64, closePriceAsOf NullableString, isoCurrencyCode NullableString, unofficialCurrencyCode NullableString) *Security {
	this := Security{}
	this.SecurityId = securityId
	this.Isin = isin
	this.Cusip = cusip
	this.Sedol = sedol
	this.InstitutionSecurityId = institutionSecurityId
	this.InstitutionId = institutionId
	this.ProxySecurityId = proxySecurityId
	this.Name = name
	this.TickerSymbol = tickerSymbol
	this.IsCashEquivalent = isCashEquivalent
	this.Type = type_
	this.ClosePrice = closePrice
	this.ClosePriceAsOf = closePriceAsOf
	this.IsoCurrencyCode = isoCurrencyCode
	this.UnofficialCurrencyCode = unofficialCurrencyCode
	return &this
}

// NewSecurityWithDefaults instantiates a new Security object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityWithDefaults() *Security {
	this := Security{}
	return &this
}

// GetSecurityId returns the SecurityId field value
func (o *Security) GetSecurityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecurityId
}

// GetSecurityIdOk returns a tuple with the SecurityId field value
// and a boolean to check if the value has been set.
func (o *Security) GetSecurityIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SecurityId, true
}

// SetSecurityId sets field value
func (o *Security) SetSecurityId(v string) {
	o.SecurityId = v
}

// GetIsin returns the Isin field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Security) GetIsin() string {
	if o == nil || o.Isin.Get() == nil {
		var ret string
		return ret
	}

	return *o.Isin.Get()
}

// GetIsinOk returns a tuple with the Isin field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetIsinOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Isin.Get(), o.Isin.IsSet()
}

// SetIsin sets field value
func (o *Security) SetIsin(v string) {
	o.Isin.Set(&v)
}

// GetCusip returns the Cusip field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Security) GetCusip() string {
	if o == nil || o.Cusip.Get() == nil {
		var ret string
		return ret
	}

	return *o.Cusip.Get()
}

// GetCusipOk returns a tuple with the Cusip field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetCusipOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Cusip.Get(), o.Cusip.IsSet()
}

// SetCusip sets field value
func (o *Security) SetCusip(v string) {
	o.Cusip.Set(&v)
}

// GetSedol returns the Sedol field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Security) GetSedol() string {
	if o == nil || o.Sedol.Get() == nil {
		var ret string
		return ret
	}

	return *o.Sedol.Get()
}

// GetSedolOk returns a tuple with the Sedol field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetSedolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Sedol.Get(), o.Sedol.IsSet()
}

// SetSedol sets field value
func (o *Security) SetSedol(v string) {
	o.Sedol.Set(&v)
}

// GetInstitutionSecurityId returns the InstitutionSecurityId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Security) GetInstitutionSecurityId() string {
	if o == nil || o.InstitutionSecurityId.Get() == nil {
		var ret string
		return ret
	}

	return *o.InstitutionSecurityId.Get()
}

// GetInstitutionSecurityIdOk returns a tuple with the InstitutionSecurityId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetInstitutionSecurityIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InstitutionSecurityId.Get(), o.InstitutionSecurityId.IsSet()
}

// SetInstitutionSecurityId sets field value
func (o *Security) SetInstitutionSecurityId(v string) {
	o.InstitutionSecurityId.Set(&v)
}

// GetInstitutionId returns the InstitutionId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Security) GetInstitutionId() string {
	if o == nil || o.InstitutionId.Get() == nil {
		var ret string
		return ret
	}

	return *o.InstitutionId.Get()
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetInstitutionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InstitutionId.Get(), o.InstitutionId.IsSet()
}

// SetInstitutionId sets field value
func (o *Security) SetInstitutionId(v string) {
	o.InstitutionId.Set(&v)
}

// GetProxySecurityId returns the ProxySecurityId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Security) GetProxySecurityId() string {
	if o == nil || o.ProxySecurityId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ProxySecurityId.Get()
}

// GetProxySecurityIdOk returns a tuple with the ProxySecurityId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetProxySecurityIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProxySecurityId.Get(), o.ProxySecurityId.IsSet()
}

// SetProxySecurityId sets field value
func (o *Security) SetProxySecurityId(v string) {
	o.ProxySecurityId.Set(&v)
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Security) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *Security) SetName(v string) {
	o.Name.Set(&v)
}

// GetTickerSymbol returns the TickerSymbol field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Security) GetTickerSymbol() string {
	if o == nil || o.TickerSymbol.Get() == nil {
		var ret string
		return ret
	}

	return *o.TickerSymbol.Get()
}

// GetTickerSymbolOk returns a tuple with the TickerSymbol field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetTickerSymbolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TickerSymbol.Get(), o.TickerSymbol.IsSet()
}

// SetTickerSymbol sets field value
func (o *Security) SetTickerSymbol(v string) {
	o.TickerSymbol.Set(&v)
}

// GetIsCashEquivalent returns the IsCashEquivalent field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *Security) GetIsCashEquivalent() bool {
	if o == nil || o.IsCashEquivalent.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsCashEquivalent.Get()
}

// GetIsCashEquivalentOk returns a tuple with the IsCashEquivalent field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetIsCashEquivalentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsCashEquivalent.Get(), o.IsCashEquivalent.IsSet()
}

// SetIsCashEquivalent sets field value
func (o *Security) SetIsCashEquivalent(v bool) {
	o.IsCashEquivalent.Set(&v)
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Security) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *Security) SetType(v string) {
	o.Type.Set(&v)
}

// GetClosePrice returns the ClosePrice field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *Security) GetClosePrice() float64 {
	if o == nil || o.ClosePrice.Get() == nil {
		var ret float64
		return ret
	}

	return *o.ClosePrice.Get()
}

// GetClosePriceOk returns a tuple with the ClosePrice field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetClosePriceOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClosePrice.Get(), o.ClosePrice.IsSet()
}

// SetClosePrice sets field value
func (o *Security) SetClosePrice(v float64) {
	o.ClosePrice.Set(&v)
}

// GetClosePriceAsOf returns the ClosePriceAsOf field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Security) GetClosePriceAsOf() string {
	if o == nil || o.ClosePriceAsOf.Get() == nil {
		var ret string
		return ret
	}

	return *o.ClosePriceAsOf.Get()
}

// GetClosePriceAsOfOk returns a tuple with the ClosePriceAsOf field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetClosePriceAsOfOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClosePriceAsOf.Get(), o.ClosePriceAsOf.IsSet()
}

// SetClosePriceAsOf sets field value
func (o *Security) SetClosePriceAsOf(v string) {
	o.ClosePriceAsOf.Set(&v)
}

// GetUpdateDatetime returns the UpdateDatetime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Security) GetUpdateDatetime() time.Time {
	if o == nil || o.UpdateDatetime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdateDatetime.Get()
}

// GetUpdateDatetimeOk returns a tuple with the UpdateDatetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetUpdateDatetimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdateDatetime.Get(), o.UpdateDatetime.IsSet()
}

// HasUpdateDatetime returns a boolean if a field has been set.
func (o *Security) HasUpdateDatetime() bool {
	if o != nil && o.UpdateDatetime.IsSet() {
		return true
	}

	return false
}

// SetUpdateDatetime gets a reference to the given NullableTime and assigns it to the UpdateDatetime field.
func (o *Security) SetUpdateDatetime(v time.Time) {
	o.UpdateDatetime.Set(&v)
}
// SetUpdateDatetimeNil sets the value for UpdateDatetime to be an explicit nil
func (o *Security) SetUpdateDatetimeNil() {
	o.UpdateDatetime.Set(nil)
}

// UnsetUpdateDatetime ensures that no value is present for UpdateDatetime, not even an explicit nil
func (o *Security) UnsetUpdateDatetime() {
	o.UpdateDatetime.Unset()
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Security) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.IsoCurrencyCode.Get()
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsoCurrencyCode.Get(), o.IsoCurrencyCode.IsSet()
}

// SetIsoCurrencyCode sets field value
func (o *Security) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode.Set(&v)
}

// GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Security) GetUnofficialCurrencyCode() string {
	if o == nil || o.UnofficialCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.UnofficialCurrencyCode.Get()
}

// GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetUnofficialCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnofficialCurrencyCode.Get(), o.UnofficialCurrencyCode.IsSet()
}

// SetUnofficialCurrencyCode sets field value
func (o *Security) SetUnofficialCurrencyCode(v string) {
	o.UnofficialCurrencyCode.Set(&v)
}

func (o Security) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["security_id"] = o.SecurityId
	}
	if true {
		toSerialize["isin"] = o.Isin.Get()
	}
	if true {
		toSerialize["cusip"] = o.Cusip.Get()
	}
	if true {
		toSerialize["sedol"] = o.Sedol.Get()
	}
	if true {
		toSerialize["institution_security_id"] = o.InstitutionSecurityId.Get()
	}
	if true {
		toSerialize["institution_id"] = o.InstitutionId.Get()
	}
	if true {
		toSerialize["proxy_security_id"] = o.ProxySecurityId.Get()
	}
	if true {
		toSerialize["name"] = o.Name.Get()
	}
	if true {
		toSerialize["ticker_symbol"] = o.TickerSymbol.Get()
	}
	if true {
		toSerialize["is_cash_equivalent"] = o.IsCashEquivalent.Get()
	}
	if true {
		toSerialize["type"] = o.Type.Get()
	}
	if true {
		toSerialize["close_price"] = o.ClosePrice.Get()
	}
	if true {
		toSerialize["close_price_as_of"] = o.ClosePriceAsOf.Get()
	}
	if o.UpdateDatetime.IsSet() {
		toSerialize["update_datetime"] = o.UpdateDatetime.Get()
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

func (o *Security) UnmarshalJSON(bytes []byte) (err error) {
	varSecurity := _Security{}

	if err = json.Unmarshal(bytes, &varSecurity); err == nil {
		*o = Security(varSecurity)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "security_id")
		delete(additionalProperties, "isin")
		delete(additionalProperties, "cusip")
		delete(additionalProperties, "sedol")
		delete(additionalProperties, "institution_security_id")
		delete(additionalProperties, "institution_id")
		delete(additionalProperties, "proxy_security_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "ticker_symbol")
		delete(additionalProperties, "is_cash_equivalent")
		delete(additionalProperties, "type")
		delete(additionalProperties, "close_price")
		delete(additionalProperties, "close_price_as_of")
		delete(additionalProperties, "update_datetime")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "unofficial_currency_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurity struct {
	value *Security
	isSet bool
}

func (v NullableSecurity) Get() *Security {
	return v.value
}

func (v *NullableSecurity) Set(val *Security) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurity(val *Security) *NullableSecurity {
	return &NullableSecurity{value: val, isSet: true}
}

func (v NullableSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


