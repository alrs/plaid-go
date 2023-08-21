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
)

// LinkSessionExit An object representing an [onExit](https://plaid.com/docs/link/web/#onexit) callback from Link.
type LinkSessionExit struct {
	Error NullablePlaidError `json:"error"`
	Metadata NullableLinkSessionExitMetadata `json:"metadata"`
	AdditionalProperties map[string]interface{}
}

type _LinkSessionExit LinkSessionExit

// NewLinkSessionExit instantiates a new LinkSessionExit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkSessionExit(error_ NullablePlaidError, metadata NullableLinkSessionExitMetadata) *LinkSessionExit {
	this := LinkSessionExit{}
	this.Error = error_
	this.Metadata = metadata
	return &this
}

// NewLinkSessionExitWithDefaults instantiates a new LinkSessionExit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkSessionExitWithDefaults() *LinkSessionExit {
	this := LinkSessionExit{}
	return &this
}

// GetError returns the Error field value
// If the value is explicit nil, the zero value for PlaidError will be returned
func (o *LinkSessionExit) GetError() PlaidError {
	if o == nil || o.Error.Get() == nil {
		var ret PlaidError
		return ret
	}

	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkSessionExit) GetErrorOk() (*PlaidError, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// SetError sets field value
func (o *LinkSessionExit) SetError(v PlaidError) {
	o.Error.Set(&v)
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for LinkSessionExitMetadata will be returned
func (o *LinkSessionExit) GetMetadata() LinkSessionExitMetadata {
	if o == nil || o.Metadata.Get() == nil {
		var ret LinkSessionExitMetadata
		return ret
	}

	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkSessionExit) GetMetadataOk() (*LinkSessionExitMetadata, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// SetMetadata sets field value
func (o *LinkSessionExit) SetMetadata(v LinkSessionExitMetadata) {
	o.Metadata.Set(&v)
}

func (o LinkSessionExit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["error"] = o.Error.Get()
	}
	if true {
		toSerialize["metadata"] = o.Metadata.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkSessionExit) UnmarshalJSON(bytes []byte) (err error) {
	varLinkSessionExit := _LinkSessionExit{}

	if err = json.Unmarshal(bytes, &varLinkSessionExit); err == nil {
		*o = LinkSessionExit(varLinkSessionExit)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkSessionExit struct {
	value *LinkSessionExit
	isSet bool
}

func (v NullableLinkSessionExit) Get() *LinkSessionExit {
	return v.value
}

func (v *NullableLinkSessionExit) Set(val *LinkSessionExit) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkSessionExit) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkSessionExit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkSessionExit(val *LinkSessionExit) *NullableLinkSessionExit {
	return &NullableLinkSessionExit{value: val, isSet: true}
}

func (v NullableLinkSessionExit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkSessionExit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


