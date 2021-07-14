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

// FieldError struct for FieldError
type FieldError struct {
	Field *string `json:"field,omitempty"`
	Message *string `json:"message,omitempty"`
	ErrorCode *string `json:"error_code,omitempty"`
}

// NewFieldError instantiates a new FieldError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldError() *FieldError {
	this := FieldError{}
	return &this
}

// NewFieldErrorWithDefaults instantiates a new FieldError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldErrorWithDefaults() *FieldError {
	this := FieldError{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *FieldError) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldError) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *FieldError) HasField() bool {
	if o != nil && o.Field != nil {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *FieldError) SetField(v string) {
	o.Field = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *FieldError) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldError) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *FieldError) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *FieldError) SetMessage(v string) {
	o.Message = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *FieldError) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldError) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *FieldError) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *FieldError) SetErrorCode(v string) {
	o.ErrorCode = &v
}

func (o FieldError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.ErrorCode != nil {
		toSerialize["error_code"] = o.ErrorCode
	}
	return json.Marshal(toSerialize)
}

type NullableFieldError struct {
	value *FieldError
	isSet bool
}

func (v NullableFieldError) Get() *FieldError {
	return v.value
}

func (v *NullableFieldError) Set(val *FieldError) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldError) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldError(val *FieldError) *NullableFieldError {
	return &NullableFieldError{value: val, isSet: true}
}

func (v NullableFieldError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


