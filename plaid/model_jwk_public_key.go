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

// JWKPublicKey A JSON Web Key (JWK) that can be used in conjunction with [JWT libraries](https://jwt.io/#libraries-io) to verify Plaid webhooks
type JWKPublicKey struct {
	// The alg member identifies the cryptographic algorithm family used with the key.
	Alg string `json:"alg"`
	// The crv member identifies the cryptographic curve used with the key.
	Crv string `json:"crv"`
	// The kid (Key ID) member can be used to match a specific key. This can be used, for instance, to choose among a set of keys within the JWK during key rollover.
	Kid string `json:"kid"`
	// The kty (key type) parameter identifies the cryptographic algorithm family used with the key, such as RSA or EC.
	Kty string `json:"kty"`
	// The use (public key use) parameter identifies the intended use of the public key.
	Use string `json:"use"`
	// The x member contains the x coordinate for the elliptic curve point.
	X string `json:"x"`
	// The y member contains the y coordinate for the elliptic curve point.
	Y string `json:"y"`
	// The timestamp when the key was created, in Unix time.
	CreatedAt int32 `json:"created_at"`
	// The timestamp when the key expired, in Unix time.
	ExpiredAt NullableInt32 `json:"expired_at"`
	AdditionalProperties map[string]interface{}
}

type _JWKPublicKey JWKPublicKey

// NewJWKPublicKey instantiates a new JWKPublicKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJWKPublicKey(alg string, crv string, kid string, kty string, use string, x string, y string, createdAt int32, expiredAt NullableInt32) *JWKPublicKey {
	this := JWKPublicKey{}
	this.Alg = alg
	this.Crv = crv
	this.Kid = kid
	this.Kty = kty
	this.Use = use
	this.X = x
	this.Y = y
	this.CreatedAt = createdAt
	this.ExpiredAt = expiredAt
	return &this
}

// NewJWKPublicKeyWithDefaults instantiates a new JWKPublicKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJWKPublicKeyWithDefaults() *JWKPublicKey {
	this := JWKPublicKey{}
	return &this
}

// GetAlg returns the Alg field value
func (o *JWKPublicKey) GetAlg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alg
}

// GetAlgOk returns a tuple with the Alg field value
// and a boolean to check if the value has been set.
func (o *JWKPublicKey) GetAlgOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Alg, true
}

// SetAlg sets field value
func (o *JWKPublicKey) SetAlg(v string) {
	o.Alg = v
}

// GetCrv returns the Crv field value
func (o *JWKPublicKey) GetCrv() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Crv
}

// GetCrvOk returns a tuple with the Crv field value
// and a boolean to check if the value has been set.
func (o *JWKPublicKey) GetCrvOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Crv, true
}

// SetCrv sets field value
func (o *JWKPublicKey) SetCrv(v string) {
	o.Crv = v
}

// GetKid returns the Kid field value
func (o *JWKPublicKey) GetKid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kid
}

// GetKidOk returns a tuple with the Kid field value
// and a boolean to check if the value has been set.
func (o *JWKPublicKey) GetKidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kid, true
}

// SetKid sets field value
func (o *JWKPublicKey) SetKid(v string) {
	o.Kid = v
}

// GetKty returns the Kty field value
func (o *JWKPublicKey) GetKty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kty
}

// GetKtyOk returns a tuple with the Kty field value
// and a boolean to check if the value has been set.
func (o *JWKPublicKey) GetKtyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kty, true
}

// SetKty sets field value
func (o *JWKPublicKey) SetKty(v string) {
	o.Kty = v
}

// GetUse returns the Use field value
func (o *JWKPublicKey) GetUse() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Use
}

// GetUseOk returns a tuple with the Use field value
// and a boolean to check if the value has been set.
func (o *JWKPublicKey) GetUseOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Use, true
}

// SetUse sets field value
func (o *JWKPublicKey) SetUse(v string) {
	o.Use = v
}

// GetX returns the X field value
func (o *JWKPublicKey) GetX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *JWKPublicKey) GetXOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *JWKPublicKey) SetX(v string) {
	o.X = v
}

// GetY returns the Y field value
func (o *JWKPublicKey) GetY() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *JWKPublicKey) GetYOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *JWKPublicKey) SetY(v string) {
	o.Y = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *JWKPublicKey) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *JWKPublicKey) GetCreatedAtOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *JWKPublicKey) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetExpiredAt returns the ExpiredAt field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *JWKPublicKey) GetExpiredAt() int32 {
	if o == nil || o.ExpiredAt.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ExpiredAt.Get()
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JWKPublicKey) GetExpiredAtOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpiredAt.Get(), o.ExpiredAt.IsSet()
}

// SetExpiredAt sets field value
func (o *JWKPublicKey) SetExpiredAt(v int32) {
	o.ExpiredAt.Set(&v)
}

func (o JWKPublicKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["alg"] = o.Alg
	}
	if true {
		toSerialize["crv"] = o.Crv
	}
	if true {
		toSerialize["kid"] = o.Kid
	}
	if true {
		toSerialize["kty"] = o.Kty
	}
	if true {
		toSerialize["use"] = o.Use
	}
	if true {
		toSerialize["x"] = o.X
	}
	if true {
		toSerialize["y"] = o.Y
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["expired_at"] = o.ExpiredAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *JWKPublicKey) UnmarshalJSON(bytes []byte) (err error) {
	varJWKPublicKey := _JWKPublicKey{}

	if err = json.Unmarshal(bytes, &varJWKPublicKey); err == nil {
		*o = JWKPublicKey(varJWKPublicKey)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "alg")
		delete(additionalProperties, "crv")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "use")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "expired_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJWKPublicKey struct {
	value *JWKPublicKey
	isSet bool
}

func (v NullableJWKPublicKey) Get() *JWKPublicKey {
	return v.value
}

func (v *NullableJWKPublicKey) Set(val *JWKPublicKey) {
	v.value = val
	v.isSet = true
}

func (v NullableJWKPublicKey) IsSet() bool {
	return v.isSet
}

func (v *NullableJWKPublicKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJWKPublicKey(val *JWKPublicKey) *NullableJWKPublicKey {
	return &NullableJWKPublicKey{value: val, isSet: true}
}

func (v NullableJWKPublicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJWKPublicKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


