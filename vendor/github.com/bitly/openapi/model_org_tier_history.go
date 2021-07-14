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

// OrgTierHistory struct for OrgTierHistory
type OrgTierHistory struct {
	Status *string `json:"status,omitempty"`
	// ISO timestamp
	Created *string `json:"created,omitempty"`
	// ISO timestamp
	Deactivated *string `json:"deactivated,omitempty"`
	// ISO timestamp
	TierChangeDate *string `json:"tier_change_date,omitempty"`
	// ISO timestamp
	Activated *string `json:"activated,omitempty"`
	// ISO timestamp
	Modified *string `json:"modified,omitempty"`
	TierName *string `json:"tier_name,omitempty"`
	OrgGuid *string `json:"org_guid,omitempty"`
}

// NewOrgTierHistory instantiates a new OrgTierHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgTierHistory() *OrgTierHistory {
	this := OrgTierHistory{}
	return &this
}

// NewOrgTierHistoryWithDefaults instantiates a new OrgTierHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgTierHistoryWithDefaults() *OrgTierHistory {
	this := OrgTierHistory{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrgTierHistory) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgTierHistory) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrgTierHistory) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrgTierHistory) SetStatus(v string) {
	o.Status = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *OrgTierHistory) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgTierHistory) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *OrgTierHistory) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *OrgTierHistory) SetCreated(v string) {
	o.Created = &v
}

// GetDeactivated returns the Deactivated field value if set, zero value otherwise.
func (o *OrgTierHistory) GetDeactivated() string {
	if o == nil || o.Deactivated == nil {
		var ret string
		return ret
	}
	return *o.Deactivated
}

// GetDeactivatedOk returns a tuple with the Deactivated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgTierHistory) GetDeactivatedOk() (*string, bool) {
	if o == nil || o.Deactivated == nil {
		return nil, false
	}
	return o.Deactivated, true
}

// HasDeactivated returns a boolean if a field has been set.
func (o *OrgTierHistory) HasDeactivated() bool {
	if o != nil && o.Deactivated != nil {
		return true
	}

	return false
}

// SetDeactivated gets a reference to the given string and assigns it to the Deactivated field.
func (o *OrgTierHistory) SetDeactivated(v string) {
	o.Deactivated = &v
}

// GetTierChangeDate returns the TierChangeDate field value if set, zero value otherwise.
func (o *OrgTierHistory) GetTierChangeDate() string {
	if o == nil || o.TierChangeDate == nil {
		var ret string
		return ret
	}
	return *o.TierChangeDate
}

// GetTierChangeDateOk returns a tuple with the TierChangeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgTierHistory) GetTierChangeDateOk() (*string, bool) {
	if o == nil || o.TierChangeDate == nil {
		return nil, false
	}
	return o.TierChangeDate, true
}

// HasTierChangeDate returns a boolean if a field has been set.
func (o *OrgTierHistory) HasTierChangeDate() bool {
	if o != nil && o.TierChangeDate != nil {
		return true
	}

	return false
}

// SetTierChangeDate gets a reference to the given string and assigns it to the TierChangeDate field.
func (o *OrgTierHistory) SetTierChangeDate(v string) {
	o.TierChangeDate = &v
}

// GetActivated returns the Activated field value if set, zero value otherwise.
func (o *OrgTierHistory) GetActivated() string {
	if o == nil || o.Activated == nil {
		var ret string
		return ret
	}
	return *o.Activated
}

// GetActivatedOk returns a tuple with the Activated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgTierHistory) GetActivatedOk() (*string, bool) {
	if o == nil || o.Activated == nil {
		return nil, false
	}
	return o.Activated, true
}

// HasActivated returns a boolean if a field has been set.
func (o *OrgTierHistory) HasActivated() bool {
	if o != nil && o.Activated != nil {
		return true
	}

	return false
}

// SetActivated gets a reference to the given string and assigns it to the Activated field.
func (o *OrgTierHistory) SetActivated(v string) {
	o.Activated = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *OrgTierHistory) GetModified() string {
	if o == nil || o.Modified == nil {
		var ret string
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgTierHistory) GetModifiedOk() (*string, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *OrgTierHistory) HasModified() bool {
	if o != nil && o.Modified != nil {
		return true
	}

	return false
}

// SetModified gets a reference to the given string and assigns it to the Modified field.
func (o *OrgTierHistory) SetModified(v string) {
	o.Modified = &v
}

// GetTierName returns the TierName field value if set, zero value otherwise.
func (o *OrgTierHistory) GetTierName() string {
	if o == nil || o.TierName == nil {
		var ret string
		return ret
	}
	return *o.TierName
}

// GetTierNameOk returns a tuple with the TierName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgTierHistory) GetTierNameOk() (*string, bool) {
	if o == nil || o.TierName == nil {
		return nil, false
	}
	return o.TierName, true
}

// HasTierName returns a boolean if a field has been set.
func (o *OrgTierHistory) HasTierName() bool {
	if o != nil && o.TierName != nil {
		return true
	}

	return false
}

// SetTierName gets a reference to the given string and assigns it to the TierName field.
func (o *OrgTierHistory) SetTierName(v string) {
	o.TierName = &v
}

// GetOrgGuid returns the OrgGuid field value if set, zero value otherwise.
func (o *OrgTierHistory) GetOrgGuid() string {
	if o == nil || o.OrgGuid == nil {
		var ret string
		return ret
	}
	return *o.OrgGuid
}

// GetOrgGuidOk returns a tuple with the OrgGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgTierHistory) GetOrgGuidOk() (*string, bool) {
	if o == nil || o.OrgGuid == nil {
		return nil, false
	}
	return o.OrgGuid, true
}

// HasOrgGuid returns a boolean if a field has been set.
func (o *OrgTierHistory) HasOrgGuid() bool {
	if o != nil && o.OrgGuid != nil {
		return true
	}

	return false
}

// SetOrgGuid gets a reference to the given string and assigns it to the OrgGuid field.
func (o *OrgTierHistory) SetOrgGuid(v string) {
	o.OrgGuid = &v
}

func (o OrgTierHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Deactivated != nil {
		toSerialize["deactivated"] = o.Deactivated
	}
	if o.TierChangeDate != nil {
		toSerialize["tier_change_date"] = o.TierChangeDate
	}
	if o.Activated != nil {
		toSerialize["activated"] = o.Activated
	}
	if o.Modified != nil {
		toSerialize["modified"] = o.Modified
	}
	if o.TierName != nil {
		toSerialize["tier_name"] = o.TierName
	}
	if o.OrgGuid != nil {
		toSerialize["org_guid"] = o.OrgGuid
	}
	return json.Marshal(toSerialize)
}

type NullableOrgTierHistory struct {
	value *OrgTierHistory
	isSet bool
}

func (v NullableOrgTierHistory) Get() *OrgTierHistory {
	return v.value
}

func (v *NullableOrgTierHistory) Set(val *OrgTierHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgTierHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgTierHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgTierHistory(val *OrgTierHistory) *NullableOrgTierHistory {
	return &NullableOrgTierHistory{value: val, isSet: true}
}

func (v NullableOrgTierHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgTierHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


