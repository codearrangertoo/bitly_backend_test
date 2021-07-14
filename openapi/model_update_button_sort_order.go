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

// UpdateButtonSortOrder struct for UpdateButtonSortOrder
type UpdateButtonSortOrder struct {
	Buttons *[]LaunchpadButtonSortOrder `json:"buttons,omitempty"`
}

// NewUpdateButtonSortOrder instantiates a new UpdateButtonSortOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateButtonSortOrder() *UpdateButtonSortOrder {
	this := UpdateButtonSortOrder{}
	return &this
}

// NewUpdateButtonSortOrderWithDefaults instantiates a new UpdateButtonSortOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateButtonSortOrderWithDefaults() *UpdateButtonSortOrder {
	this := UpdateButtonSortOrder{}
	return &this
}

// GetButtons returns the Buttons field value if set, zero value otherwise.
func (o *UpdateButtonSortOrder) GetButtons() []LaunchpadButtonSortOrder {
	if o == nil || o.Buttons == nil {
		var ret []LaunchpadButtonSortOrder
		return ret
	}
	return *o.Buttons
}

// GetButtonsOk returns a tuple with the Buttons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateButtonSortOrder) GetButtonsOk() (*[]LaunchpadButtonSortOrder, bool) {
	if o == nil || o.Buttons == nil {
		return nil, false
	}
	return o.Buttons, true
}

// HasButtons returns a boolean if a field has been set.
func (o *UpdateButtonSortOrder) HasButtons() bool {
	if o != nil && o.Buttons != nil {
		return true
	}

	return false
}

// SetButtons gets a reference to the given []LaunchpadButtonSortOrder and assigns it to the Buttons field.
func (o *UpdateButtonSortOrder) SetButtons(v []LaunchpadButtonSortOrder) {
	o.Buttons = &v
}

func (o UpdateButtonSortOrder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Buttons != nil {
		toSerialize["buttons"] = o.Buttons
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateButtonSortOrder struct {
	value *UpdateButtonSortOrder
	isSet bool
}

func (v NullableUpdateButtonSortOrder) Get() *UpdateButtonSortOrder {
	return v.value
}

func (v *NullableUpdateButtonSortOrder) Set(val *UpdateButtonSortOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateButtonSortOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateButtonSortOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateButtonSortOrder(val *UpdateButtonSortOrder) *NullableUpdateButtonSortOrder {
	return &NullableUpdateButtonSortOrder{value: val, isSet: true}
}

func (v NullableUpdateButtonSortOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateButtonSortOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


