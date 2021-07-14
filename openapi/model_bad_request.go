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

// BadRequest BAD_REQUEST
type BadRequest struct {
	Message *string `json:"message,omitempty"`
	Errors *[]FieldError `json:"errors,omitempty"`
	Resource *string `json:"resource,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewBadRequest instantiates a new BadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBadRequest() *BadRequest {
	this := BadRequest{}
	return &this
}

// NewBadRequestWithDefaults instantiates a new BadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBadRequestWithDefaults() *BadRequest {
	this := BadRequest{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BadRequest) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequest) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BadRequest) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BadRequest) SetMessage(v string) {
	o.Message = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *BadRequest) GetErrors() []FieldError {
	if o == nil || o.Errors == nil {
		var ret []FieldError
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequest) GetErrorsOk() (*[]FieldError, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *BadRequest) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []FieldError and assigns it to the Errors field.
func (o *BadRequest) SetErrors(v []FieldError) {
	o.Errors = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *BadRequest) GetResource() string {
	if o == nil || o.Resource == nil {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequest) GetResourceOk() (*string, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *BadRequest) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *BadRequest) SetResource(v string) {
	o.Resource = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BadRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BadRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BadRequest) SetDescription(v string) {
	o.Description = &v
}

func (o BadRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableBadRequest struct {
	value *BadRequest
	isSet bool
}

func (v NullableBadRequest) Get() *BadRequest {
	return v.value
}

func (v *NullableBadRequest) Set(val *BadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBadRequest(val *BadRequest) *NullableBadRequest {
	return &NullableBadRequest{value: val, isSet: true}
}

func (v NullableBadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


