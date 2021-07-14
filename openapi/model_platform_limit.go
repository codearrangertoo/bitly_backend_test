/*
 * Bitly API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.0.0
 * Contact: api@bitly.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PlatformLimit struct for PlatformLimit
type PlatformLimit struct {
	Endpoint *string `json:"endpoint,omitempty"`
	Methods *[]MethodLimit `json:"methods,omitempty"`
}

// NewPlatformLimit instantiates a new PlatformLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformLimit() *PlatformLimit {
	this := PlatformLimit{}
	return &this
}

// NewPlatformLimitWithDefaults instantiates a new PlatformLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformLimitWithDefaults() *PlatformLimit {
	this := PlatformLimit{}
	return &this
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *PlatformLimit) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformLimit) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *PlatformLimit) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *PlatformLimit) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetMethods returns the Methods field value if set, zero value otherwise.
func (o *PlatformLimit) GetMethods() []MethodLimit {
	if o == nil || o.Methods == nil {
		var ret []MethodLimit
		return ret
	}
	return *o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformLimit) GetMethodsOk() (*[]MethodLimit, bool) {
	if o == nil || o.Methods == nil {
		return nil, false
	}
	return o.Methods, true
}

// HasMethods returns a boolean if a field has been set.
func (o *PlatformLimit) HasMethods() bool {
	if o != nil && o.Methods != nil {
		return true
	}

	return false
}

// SetMethods gets a reference to the given []MethodLimit and assigns it to the Methods field.
func (o *PlatformLimit) SetMethods(v []MethodLimit) {
	o.Methods = &v
}

func (o PlatformLimit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.Methods != nil {
		toSerialize["methods"] = o.Methods
	}
	return json.Marshal(toSerialize)
}

type NullablePlatformLimit struct {
	value *PlatformLimit
	isSet bool
}

func (v NullablePlatformLimit) Get() *PlatformLimit {
	return v.value
}

func (v *NullablePlatformLimit) Set(val *PlatformLimit) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformLimit) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformLimit(val *PlatformLimit) *NullablePlatformLimit {
	return &NullablePlatformLimit{value: val, isSet: true}
}

func (v NullablePlatformLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


