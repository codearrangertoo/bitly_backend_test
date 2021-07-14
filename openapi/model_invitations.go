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

// Invitations struct for Invitations
type Invitations struct {
	Invitations *[]Invitation `json:"invitations,omitempty"`
}

// NewInvitations instantiates a new Invitations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitations() *Invitations {
	this := Invitations{}
	return &this
}

// NewInvitationsWithDefaults instantiates a new Invitations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationsWithDefaults() *Invitations {
	this := Invitations{}
	return &this
}

// GetInvitations returns the Invitations field value if set, zero value otherwise.
func (o *Invitations) GetInvitations() []Invitation {
	if o == nil || o.Invitations == nil {
		var ret []Invitation
		return ret
	}
	return *o.Invitations
}

// GetInvitationsOk returns a tuple with the Invitations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitations) GetInvitationsOk() (*[]Invitation, bool) {
	if o == nil || o.Invitations == nil {
		return nil, false
	}
	return o.Invitations, true
}

// HasInvitations returns a boolean if a field has been set.
func (o *Invitations) HasInvitations() bool {
	if o != nil && o.Invitations != nil {
		return true
	}

	return false
}

// SetInvitations gets a reference to the given []Invitation and assigns it to the Invitations field.
func (o *Invitations) SetInvitations(v []Invitation) {
	o.Invitations = &v
}

func (o Invitations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Invitations != nil {
		toSerialize["invitations"] = o.Invitations
	}
	return json.Marshal(toSerialize)
}

type NullableInvitations struct {
	value *Invitations
	isSet bool
}

func (v NullableInvitations) Get() *Invitations {
	return v.value
}

func (v *NullableInvitations) Set(val *Invitations) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitations) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitations(val *Invitations) *NullableInvitations {
	return &NullableInvitations{value: val, isSet: true}
}

func (v NullableInvitations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


