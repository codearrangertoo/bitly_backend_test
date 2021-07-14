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

// OAuthAppWithOwnerLogin struct for OAuthAppWithOwnerLogin
type OAuthAppWithOwnerLogin struct {
	OwnerLogin *string `json:"owner_login,omitempty"`
	Link *string `json:"link,omitempty"`
	Name *string `json:"name,omitempty"`
	ClientId *string `json:"client_id,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewOAuthAppWithOwnerLogin instantiates a new OAuthAppWithOwnerLogin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthAppWithOwnerLogin() *OAuthAppWithOwnerLogin {
	this := OAuthAppWithOwnerLogin{}
	return &this
}

// NewOAuthAppWithOwnerLoginWithDefaults instantiates a new OAuthAppWithOwnerLogin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthAppWithOwnerLoginWithDefaults() *OAuthAppWithOwnerLogin {
	this := OAuthAppWithOwnerLogin{}
	return &this
}

// GetOwnerLogin returns the OwnerLogin field value if set, zero value otherwise.
func (o *OAuthAppWithOwnerLogin) GetOwnerLogin() string {
	if o == nil || o.OwnerLogin == nil {
		var ret string
		return ret
	}
	return *o.OwnerLogin
}

// GetOwnerLoginOk returns a tuple with the OwnerLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthAppWithOwnerLogin) GetOwnerLoginOk() (*string, bool) {
	if o == nil || o.OwnerLogin == nil {
		return nil, false
	}
	return o.OwnerLogin, true
}

// HasOwnerLogin returns a boolean if a field has been set.
func (o *OAuthAppWithOwnerLogin) HasOwnerLogin() bool {
	if o != nil && o.OwnerLogin != nil {
		return true
	}

	return false
}

// SetOwnerLogin gets a reference to the given string and assigns it to the OwnerLogin field.
func (o *OAuthAppWithOwnerLogin) SetOwnerLogin(v string) {
	o.OwnerLogin = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *OAuthAppWithOwnerLogin) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthAppWithOwnerLogin) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *OAuthAppWithOwnerLogin) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *OAuthAppWithOwnerLogin) SetLink(v string) {
	o.Link = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OAuthAppWithOwnerLogin) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthAppWithOwnerLogin) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OAuthAppWithOwnerLogin) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OAuthAppWithOwnerLogin) SetName(v string) {
	o.Name = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *OAuthAppWithOwnerLogin) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthAppWithOwnerLogin) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OAuthAppWithOwnerLogin) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *OAuthAppWithOwnerLogin) SetClientId(v string) {
	o.ClientId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OAuthAppWithOwnerLogin) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthAppWithOwnerLogin) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OAuthAppWithOwnerLogin) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OAuthAppWithOwnerLogin) SetDescription(v string) {
	o.Description = &v
}

func (o OAuthAppWithOwnerLogin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OwnerLogin != nil {
		toSerialize["owner_login"] = o.OwnerLogin
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableOAuthAppWithOwnerLogin struct {
	value *OAuthAppWithOwnerLogin
	isSet bool
}

func (v NullableOAuthAppWithOwnerLogin) Get() *OAuthAppWithOwnerLogin {
	return v.value
}

func (v *NullableOAuthAppWithOwnerLogin) Set(val *OAuthAppWithOwnerLogin) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthAppWithOwnerLogin) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthAppWithOwnerLogin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthAppWithOwnerLogin(val *OAuthAppWithOwnerLogin) *NullableOAuthAppWithOwnerLogin {
	return &NullableOAuthAppWithOwnerLogin{value: val, isSet: true}
}

func (v NullableOAuthAppWithOwnerLogin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthAppWithOwnerLogin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


