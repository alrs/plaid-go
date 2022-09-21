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
)

// ItemImportRequest ItemImportRequest defines the request schema for `/item/import`
type ItemImportRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Array of product strings
	Products []Products `json:"products"`
	UserAuth ItemImportRequestUserAuth `json:"user_auth"`
	Options *ItemImportRequestOptions `json:"options,omitempty"`
}

// NewItemImportRequest instantiates a new ItemImportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemImportRequest(products []Products, userAuth ItemImportRequestUserAuth) *ItemImportRequest {
	this := ItemImportRequest{}
	this.Products = products
	this.UserAuth = userAuth
	return &this
}

// NewItemImportRequestWithDefaults instantiates a new ItemImportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemImportRequestWithDefaults() *ItemImportRequest {
	this := ItemImportRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ItemImportRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemImportRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ItemImportRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ItemImportRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ItemImportRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemImportRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ItemImportRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ItemImportRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetProducts returns the Products field value
func (o *ItemImportRequest) GetProducts() []Products {
	if o == nil {
		var ret []Products
		return ret
	}

	return o.Products
}

// GetProductsOk returns a tuple with the Products field value
// and a boolean to check if the value has been set.
func (o *ItemImportRequest) GetProductsOk() (*[]Products, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Products, true
}

// SetProducts sets field value
func (o *ItemImportRequest) SetProducts(v []Products) {
	o.Products = v
}

// GetUserAuth returns the UserAuth field value
func (o *ItemImportRequest) GetUserAuth() ItemImportRequestUserAuth {
	if o == nil {
		var ret ItemImportRequestUserAuth
		return ret
	}

	return o.UserAuth
}

// GetUserAuthOk returns a tuple with the UserAuth field value
// and a boolean to check if the value has been set.
func (o *ItemImportRequest) GetUserAuthOk() (*ItemImportRequestUserAuth, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserAuth, true
}

// SetUserAuth sets field value
func (o *ItemImportRequest) SetUserAuth(v ItemImportRequestUserAuth) {
	o.UserAuth = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ItemImportRequest) GetOptions() ItemImportRequestOptions {
	if o == nil || o.Options == nil {
		var ret ItemImportRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemImportRequest) GetOptionsOk() (*ItemImportRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ItemImportRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given ItemImportRequestOptions and assigns it to the Options field.
func (o *ItemImportRequest) SetOptions(v ItemImportRequestOptions) {
	o.Options = &v
}

func (o ItemImportRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["products"] = o.Products
	}
	if true {
		toSerialize["user_auth"] = o.UserAuth
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableItemImportRequest struct {
	value *ItemImportRequest
	isSet bool
}

func (v NullableItemImportRequest) Get() *ItemImportRequest {
	return v.value
}

func (v *NullableItemImportRequest) Set(val *ItemImportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableItemImportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableItemImportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemImportRequest(val *ItemImportRequest) *NullableItemImportRequest {
	return &NullableItemImportRequest{value: val, isSet: true}
}

func (v NullableItemImportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemImportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


