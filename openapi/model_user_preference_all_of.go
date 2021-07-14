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

// UserPreferenceAllOf struct for UserPreferenceAllOf
type UserPreferenceAllOf struct {
	Created *string `json:"created,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	Modified *string `json:"modified,omitempty"`
	Value *string `json:"value,omitempty"`
	Preference *string `json:"preference,omitempty"`
	Login *string `json:"login,omitempty"`
	Expired *bool `json:"expired,omitempty"`
	LastSeen *string `json:"last_seen,omitempty"`
}

// NewUserPreferenceAllOf instantiates a new UserPreferenceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPreferenceAllOf() *UserPreferenceAllOf {
	this := UserPreferenceAllOf{}
	return &this
}

// NewUserPreferenceAllOfWithDefaults instantiates a new UserPreferenceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPreferenceAllOfWithDefaults() *UserPreferenceAllOf {
	this := UserPreferenceAllOf{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *UserPreferenceAllOf) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferenceAllOf) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *UserPreferenceAllOf) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *UserPreferenceAllOf) SetCreated(v string) {
	o.Created = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *UserPreferenceAllOf) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferenceAllOf) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *UserPreferenceAllOf) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *UserPreferenceAllOf) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *UserPreferenceAllOf) GetModified() string {
	if o == nil || o.Modified == nil {
		var ret string
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferenceAllOf) GetModifiedOk() (*string, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *UserPreferenceAllOf) HasModified() bool {
	if o != nil && o.Modified != nil {
		return true
	}

	return false
}

// SetModified gets a reference to the given string and assigns it to the Modified field.
func (o *UserPreferenceAllOf) SetModified(v string) {
	o.Modified = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *UserPreferenceAllOf) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferenceAllOf) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UserPreferenceAllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *UserPreferenceAllOf) SetValue(v string) {
	o.Value = &v
}

// GetPreference returns the Preference field value if set, zero value otherwise.
func (o *UserPreferenceAllOf) GetPreference() string {
	if o == nil || o.Preference == nil {
		var ret string
		return ret
	}
	return *o.Preference
}

// GetPreferenceOk returns a tuple with the Preference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferenceAllOf) GetPreferenceOk() (*string, bool) {
	if o == nil || o.Preference == nil {
		return nil, false
	}
	return o.Preference, true
}

// HasPreference returns a boolean if a field has been set.
func (o *UserPreferenceAllOf) HasPreference() bool {
	if o != nil && o.Preference != nil {
		return true
	}

	return false
}

// SetPreference gets a reference to the given string and assigns it to the Preference field.
func (o *UserPreferenceAllOf) SetPreference(v string) {
	o.Preference = &v
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *UserPreferenceAllOf) GetLogin() string {
	if o == nil || o.Login == nil {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferenceAllOf) GetLoginOk() (*string, bool) {
	if o == nil || o.Login == nil {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *UserPreferenceAllOf) HasLogin() bool {
	if o != nil && o.Login != nil {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *UserPreferenceAllOf) SetLogin(v string) {
	o.Login = &v
}

// GetExpired returns the Expired field value if set, zero value otherwise.
func (o *UserPreferenceAllOf) GetExpired() bool {
	if o == nil || o.Expired == nil {
		var ret bool
		return ret
	}
	return *o.Expired
}

// GetExpiredOk returns a tuple with the Expired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferenceAllOf) GetExpiredOk() (*bool, bool) {
	if o == nil || o.Expired == nil {
		return nil, false
	}
	return o.Expired, true
}

// HasExpired returns a boolean if a field has been set.
func (o *UserPreferenceAllOf) HasExpired() bool {
	if o != nil && o.Expired != nil {
		return true
	}

	return false
}

// SetExpired gets a reference to the given bool and assigns it to the Expired field.
func (o *UserPreferenceAllOf) SetExpired(v bool) {
	o.Expired = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *UserPreferenceAllOf) GetLastSeen() string {
	if o == nil || o.LastSeen == nil {
		var ret string
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferenceAllOf) GetLastSeenOk() (*string, bool) {
	if o == nil || o.LastSeen == nil {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *UserPreferenceAllOf) HasLastSeen() bool {
	if o != nil && o.LastSeen != nil {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given string and assigns it to the LastSeen field.
func (o *UserPreferenceAllOf) SetLastSeen(v string) {
	o.LastSeen = &v
}

func (o UserPreferenceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.IsActive != nil {
		toSerialize["is_active"] = o.IsActive
	}
	if o.Modified != nil {
		toSerialize["modified"] = o.Modified
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Preference != nil {
		toSerialize["preference"] = o.Preference
	}
	if o.Login != nil {
		toSerialize["login"] = o.Login
	}
	if o.Expired != nil {
		toSerialize["expired"] = o.Expired
	}
	if o.LastSeen != nil {
		toSerialize["last_seen"] = o.LastSeen
	}
	return json.Marshal(toSerialize)
}

type NullableUserPreferenceAllOf struct {
	value *UserPreferenceAllOf
	isSet bool
}

func (v NullableUserPreferenceAllOf) Get() *UserPreferenceAllOf {
	return v.value
}

func (v *NullableUserPreferenceAllOf) Set(val *UserPreferenceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPreferenceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPreferenceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPreferenceAllOf(val *UserPreferenceAllOf) *NullableUserPreferenceAllOf {
	return &NullableUserPreferenceAllOf{value: val, isSet: true}
}

func (v NullableUserPreferenceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPreferenceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


