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
)

// CreditFreddieMacVerificationOfAssetVOA24 Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
type CreditFreddieMacVerificationOfAssetVOA24 struct {
	REPORTING_INFORMATION CreditFreddieMacReportingInformationVOA24 `json:"REPORTING_INFORMATION"`
	SERVICE_PRODUCT_FULFILLMENT ServiceProductFulfillment `json:"SERVICE_PRODUCT_FULFILLMENT"`
	VERIFICATION_OF_ASSET_RESPONSE CreditFreddieMacVerificationOfAssetResponseVOA24 `json:"VERIFICATION_OF_ASSET_RESPONSE"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacVerificationOfAssetVOA24 CreditFreddieMacVerificationOfAssetVOA24

// NewCreditFreddieMacVerificationOfAssetVOA24 instantiates a new CreditFreddieMacVerificationOfAssetVOA24 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacVerificationOfAssetVOA24(rEPORTINGINFORMATION CreditFreddieMacReportingInformationVOA24, sERVICEPRODUCTFULFILLMENT ServiceProductFulfillment, vERIFICATIONOFASSETRESPONSE CreditFreddieMacVerificationOfAssetResponseVOA24) *CreditFreddieMacVerificationOfAssetVOA24 {
	this := CreditFreddieMacVerificationOfAssetVOA24{}
	this.REPORTING_INFORMATION = rEPORTINGINFORMATION
	this.SERVICE_PRODUCT_FULFILLMENT = sERVICEPRODUCTFULFILLMENT
	this.VERIFICATION_OF_ASSET_RESPONSE = vERIFICATIONOFASSETRESPONSE
	return &this
}

// NewCreditFreddieMacVerificationOfAssetVOA24WithDefaults instantiates a new CreditFreddieMacVerificationOfAssetVOA24 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacVerificationOfAssetVOA24WithDefaults() *CreditFreddieMacVerificationOfAssetVOA24 {
	this := CreditFreddieMacVerificationOfAssetVOA24{}
	return &this
}

// GetREPORTING_INFORMATION returns the REPORTING_INFORMATION field value
func (o *CreditFreddieMacVerificationOfAssetVOA24) GetREPORTING_INFORMATION() CreditFreddieMacReportingInformationVOA24 {
	if o == nil {
		var ret CreditFreddieMacReportingInformationVOA24
		return ret
	}

	return o.REPORTING_INFORMATION
}

// GetREPORTING_INFORMATIONOk returns a tuple with the REPORTING_INFORMATION field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacVerificationOfAssetVOA24) GetREPORTING_INFORMATIONOk() (*CreditFreddieMacReportingInformationVOA24, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.REPORTING_INFORMATION, true
}

// SetREPORTING_INFORMATION sets field value
func (o *CreditFreddieMacVerificationOfAssetVOA24) SetREPORTING_INFORMATION(v CreditFreddieMacReportingInformationVOA24) {
	o.REPORTING_INFORMATION = v
}

// GetSERVICE_PRODUCT_FULFILLMENT returns the SERVICE_PRODUCT_FULFILLMENT field value
func (o *CreditFreddieMacVerificationOfAssetVOA24) GetSERVICE_PRODUCT_FULFILLMENT() ServiceProductFulfillment {
	if o == nil {
		var ret ServiceProductFulfillment
		return ret
	}

	return o.SERVICE_PRODUCT_FULFILLMENT
}

// GetSERVICE_PRODUCT_FULFILLMENTOk returns a tuple with the SERVICE_PRODUCT_FULFILLMENT field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacVerificationOfAssetVOA24) GetSERVICE_PRODUCT_FULFILLMENTOk() (*ServiceProductFulfillment, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SERVICE_PRODUCT_FULFILLMENT, true
}

// SetSERVICE_PRODUCT_FULFILLMENT sets field value
func (o *CreditFreddieMacVerificationOfAssetVOA24) SetSERVICE_PRODUCT_FULFILLMENT(v ServiceProductFulfillment) {
	o.SERVICE_PRODUCT_FULFILLMENT = v
}

// GetVERIFICATION_OF_ASSET_RESPONSE returns the VERIFICATION_OF_ASSET_RESPONSE field value
func (o *CreditFreddieMacVerificationOfAssetVOA24) GetVERIFICATION_OF_ASSET_RESPONSE() CreditFreddieMacVerificationOfAssetResponseVOA24 {
	if o == nil {
		var ret CreditFreddieMacVerificationOfAssetResponseVOA24
		return ret
	}

	return o.VERIFICATION_OF_ASSET_RESPONSE
}

// GetVERIFICATION_OF_ASSET_RESPONSEOk returns a tuple with the VERIFICATION_OF_ASSET_RESPONSE field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacVerificationOfAssetVOA24) GetVERIFICATION_OF_ASSET_RESPONSEOk() (*CreditFreddieMacVerificationOfAssetResponseVOA24, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VERIFICATION_OF_ASSET_RESPONSE, true
}

// SetVERIFICATION_OF_ASSET_RESPONSE sets field value
func (o *CreditFreddieMacVerificationOfAssetVOA24) SetVERIFICATION_OF_ASSET_RESPONSE(v CreditFreddieMacVerificationOfAssetResponseVOA24) {
	o.VERIFICATION_OF_ASSET_RESPONSE = v
}

func (o CreditFreddieMacVerificationOfAssetVOA24) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["REPORTING_INFORMATION"] = o.REPORTING_INFORMATION
	}
	if true {
		toSerialize["SERVICE_PRODUCT_FULFILLMENT"] = o.SERVICE_PRODUCT_FULFILLMENT
	}
	if true {
		toSerialize["VERIFICATION_OF_ASSET_RESPONSE"] = o.VERIFICATION_OF_ASSET_RESPONSE
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacVerificationOfAssetVOA24) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacVerificationOfAssetVOA24 := _CreditFreddieMacVerificationOfAssetVOA24{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacVerificationOfAssetVOA24); err == nil {
		*o = CreditFreddieMacVerificationOfAssetVOA24(varCreditFreddieMacVerificationOfAssetVOA24)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "REPORTING_INFORMATION")
		delete(additionalProperties, "SERVICE_PRODUCT_FULFILLMENT")
		delete(additionalProperties, "VERIFICATION_OF_ASSET_RESPONSE")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieMacVerificationOfAssetVOA24 struct {
	value *CreditFreddieMacVerificationOfAssetVOA24
	isSet bool
}

func (v NullableCreditFreddieMacVerificationOfAssetVOA24) Get() *CreditFreddieMacVerificationOfAssetVOA24 {
	return v.value
}

func (v *NullableCreditFreddieMacVerificationOfAssetVOA24) Set(val *CreditFreddieMacVerificationOfAssetVOA24) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacVerificationOfAssetVOA24) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacVerificationOfAssetVOA24) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacVerificationOfAssetVOA24(val *CreditFreddieMacVerificationOfAssetVOA24) *NullableCreditFreddieMacVerificationOfAssetVOA24 {
	return &NullableCreditFreddieMacVerificationOfAssetVOA24{value: val, isSet: true}
}

func (v NullableCreditFreddieMacVerificationOfAssetVOA24) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacVerificationOfAssetVOA24) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


