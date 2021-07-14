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

// DeactivateUser struct for DeactivateUser
type DeactivateUser struct {
	Feedback *Feedback `json:"feedback,omitempty"`
	ConfirmText *string `json:"confirm_text,omitempty"`
}

// NewDeactivateUser instantiates a new DeactivateUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeactivateUser() *DeactivateUser {
	this := DeactivateUser{}
	return &this
}

// NewDeactivateUserWithDefaults instantiates a new DeactivateUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeactivateUserWithDefaults() *DeactivateUser {
	this := DeactivateUser{}
	return &this
}

// GetFeedback returns the Feedback field value if set, zero value otherwise.
func (o *DeactivateUser) GetFeedback() Feedback {
	if o == nil || o.Feedback == nil {
		var ret Feedback
		return ret
	}
	return *o.Feedback
}

// GetFeedbackOk returns a tuple with the Feedback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeactivateUser) GetFeedbackOk() (*Feedback, bool) {
	if o == nil || o.Feedback == nil {
		return nil, false
	}
	return o.Feedback, true
}

// HasFeedback returns a boolean if a field has been set.
func (o *DeactivateUser) HasFeedback() bool {
	if o != nil && o.Feedback != nil {
		return true
	}

	return false
}

// SetFeedback gets a reference to the given Feedback and assigns it to the Feedback field.
func (o *DeactivateUser) SetFeedback(v Feedback) {
	o.Feedback = &v
}

// GetConfirmText returns the ConfirmText field value if set, zero value otherwise.
func (o *DeactivateUser) GetConfirmText() string {
	if o == nil || o.ConfirmText == nil {
		var ret string
		return ret
	}
	return *o.ConfirmText
}

// GetConfirmTextOk returns a tuple with the ConfirmText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeactivateUser) GetConfirmTextOk() (*string, bool) {
	if o == nil || o.ConfirmText == nil {
		return nil, false
	}
	return o.ConfirmText, true
}

// HasConfirmText returns a boolean if a field has been set.
func (o *DeactivateUser) HasConfirmText() bool {
	if o != nil && o.ConfirmText != nil {
		return true
	}

	return false
}

// SetConfirmText gets a reference to the given string and assigns it to the ConfirmText field.
func (o *DeactivateUser) SetConfirmText(v string) {
	o.ConfirmText = &v
}

func (o DeactivateUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Feedback != nil {
		toSerialize["feedback"] = o.Feedback
	}
	if o.ConfirmText != nil {
		toSerialize["confirm_text"] = o.ConfirmText
	}
	return json.Marshal(toSerialize)
}

type NullableDeactivateUser struct {
	value *DeactivateUser
	isSet bool
}

func (v NullableDeactivateUser) Get() *DeactivateUser {
	return v.value
}

func (v *NullableDeactivateUser) Set(val *DeactivateUser) {
	v.value = val
	v.isSet = true
}

func (v NullableDeactivateUser) IsSet() bool {
	return v.isSet
}

func (v *NullableDeactivateUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeactivateUser(val *DeactivateUser) *NullableDeactivateUser {
	return &NullableDeactivateUser{value: val, isSet: true}
}

func (v NullableDeactivateUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeactivateUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


