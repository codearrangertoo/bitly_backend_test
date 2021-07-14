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

// UserInternal struct for UserInternal
type UserInternal struct {
	DefaultGroupGuid *string `json:"default_group_guid,omitempty"`
	Name string `json:"name"`
	Created string `json:"created"`
	IsActive bool `json:"is_active"`
	Modified string `json:"modified"`
	IsSsoUser bool `json:"is_sso_user"`
	Is2faEnabled bool `json:"is_2fa_enabled"`
	Login string `json:"login"`
	Emails []Email `json:"emails"`
	RoleName *string `json:"role_name,omitempty"`
}

// NewUserInternal instantiates a new UserInternal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInternal(name string, created string, isActive bool, modified string, isSsoUser bool, is2faEnabled bool, login string, emails []Email) *UserInternal {
	this := UserInternal{}
	this.Name = name
	this.Created = created
	this.IsActive = isActive
	this.Modified = modified
	this.IsSsoUser = isSsoUser
	this.Is2faEnabled = is2faEnabled
	this.Login = login
	this.Emails = emails
	return &this
}

// NewUserInternalWithDefaults instantiates a new UserInternal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInternalWithDefaults() *UserInternal {
	this := UserInternal{}
	return &this
}

// GetDefaultGroupGuid returns the DefaultGroupGuid field value if set, zero value otherwise.
func (o *UserInternal) GetDefaultGroupGuid() string {
	if o == nil || o.DefaultGroupGuid == nil {
		var ret string
		return ret
	}
	return *o.DefaultGroupGuid
}

// GetDefaultGroupGuidOk returns a tuple with the DefaultGroupGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInternal) GetDefaultGroupGuidOk() (*string, bool) {
	if o == nil || o.DefaultGroupGuid == nil {
		return nil, false
	}
	return o.DefaultGroupGuid, true
}

// HasDefaultGroupGuid returns a boolean if a field has been set.
func (o *UserInternal) HasDefaultGroupGuid() bool {
	if o != nil && o.DefaultGroupGuid != nil {
		return true
	}

	return false
}

// SetDefaultGroupGuid gets a reference to the given string and assigns it to the DefaultGroupGuid field.
func (o *UserInternal) SetDefaultGroupGuid(v string) {
	o.DefaultGroupGuid = &v
}

// GetName returns the Name field value
func (o *UserInternal) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserInternal) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserInternal) SetName(v string) {
	o.Name = v
}

// GetCreated returns the Created field value
func (o *UserInternal) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *UserInternal) GetCreatedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *UserInternal) SetCreated(v string) {
	o.Created = v
}

// GetIsActive returns the IsActive field value
func (o *UserInternal) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *UserInternal) GetIsActiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *UserInternal) SetIsActive(v bool) {
	o.IsActive = v
}

// GetModified returns the Modified field value
func (o *UserInternal) GetModified() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *UserInternal) GetModifiedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *UserInternal) SetModified(v string) {
	o.Modified = v
}

// GetIsSsoUser returns the IsSsoUser field value
func (o *UserInternal) GetIsSsoUser() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSsoUser
}

// GetIsSsoUserOk returns a tuple with the IsSsoUser field value
// and a boolean to check if the value has been set.
func (o *UserInternal) GetIsSsoUserOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsSsoUser, true
}

// SetIsSsoUser sets field value
func (o *UserInternal) SetIsSsoUser(v bool) {
	o.IsSsoUser = v
}

// GetIs2faEnabled returns the Is2faEnabled field value
func (o *UserInternal) GetIs2faEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Is2faEnabled
}

// GetIs2faEnabledOk returns a tuple with the Is2faEnabled field value
// and a boolean to check if the value has been set.
func (o *UserInternal) GetIs2faEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Is2faEnabled, true
}

// SetIs2faEnabled sets field value
func (o *UserInternal) SetIs2faEnabled(v bool) {
	o.Is2faEnabled = v
}

// GetLogin returns the Login field value
func (o *UserInternal) GetLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Login
}

// GetLoginOk returns a tuple with the Login field value
// and a boolean to check if the value has been set.
func (o *UserInternal) GetLoginOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Login, true
}

// SetLogin sets field value
func (o *UserInternal) SetLogin(v string) {
	o.Login = v
}

// GetEmails returns the Emails field value
func (o *UserInternal) GetEmails() []Email {
	if o == nil {
		var ret []Email
		return ret
	}

	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value
// and a boolean to check if the value has been set.
func (o *UserInternal) GetEmailsOk() (*[]Email, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Emails, true
}

// SetEmails sets field value
func (o *UserInternal) SetEmails(v []Email) {
	o.Emails = v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *UserInternal) GetRoleName() string {
	if o == nil || o.RoleName == nil {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInternal) GetRoleNameOk() (*string, bool) {
	if o == nil || o.RoleName == nil {
		return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *UserInternal) HasRoleName() bool {
	if o != nil && o.RoleName != nil {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *UserInternal) SetRoleName(v string) {
	o.RoleName = &v
}

func (o UserInternal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultGroupGuid != nil {
		toSerialize["default_group_guid"] = o.DefaultGroupGuid
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["is_active"] = o.IsActive
	}
	if true {
		toSerialize["modified"] = o.Modified
	}
	if true {
		toSerialize["is_sso_user"] = o.IsSsoUser
	}
	if true {
		toSerialize["is_2fa_enabled"] = o.Is2faEnabled
	}
	if true {
		toSerialize["login"] = o.Login
	}
	if true {
		toSerialize["emails"] = o.Emails
	}
	if o.RoleName != nil {
		toSerialize["role_name"] = o.RoleName
	}
	return json.Marshal(toSerialize)
}

type NullableUserInternal struct {
	value *UserInternal
	isSet bool
}

func (v NullableUserInternal) Get() *UserInternal {
	return v.value
}

func (v *NullableUserInternal) Set(val *UserInternal) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInternal) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInternal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInternal(val *UserInternal) *NullableUserInternal {
	return &NullableUserInternal{value: val, isSet: true}
}

func (v NullableUserInternal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInternal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


