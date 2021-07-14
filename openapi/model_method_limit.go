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

// MethodLimit struct for MethodLimit
type MethodLimit struct {
	Count *int32 `json:"count,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewMethodLimit instantiates a new MethodLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMethodLimit() *MethodLimit {
	this := MethodLimit{}
	return &this
}

// NewMethodLimitWithDefaults instantiates a new MethodLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMethodLimitWithDefaults() *MethodLimit {
	this := MethodLimit{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *MethodLimit) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MethodLimit) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *MethodLimit) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *MethodLimit) SetCount(v int32) {
	o.Count = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *MethodLimit) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MethodLimit) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *MethodLimit) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *MethodLimit) SetLimit(v int32) {
	o.Limit = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MethodLimit) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MethodLimit) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MethodLimit) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MethodLimit) SetName(v string) {
	o.Name = &v
}

func (o MethodLimit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableMethodLimit struct {
	value *MethodLimit
	isSet bool
}

func (v NullableMethodLimit) Get() *MethodLimit {
	return v.value
}

func (v *NullableMethodLimit) Set(val *MethodLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableMethodLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableMethodLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMethodLimit(val *MethodLimit) *NullableMethodLimit {
	return &NullableMethodLimit{value: val, isSet: true}
}

func (v NullableMethodLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMethodLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


