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

// CampaignAllOf struct for CampaignAllOf
type CampaignAllOf struct {
	Name *string `json:"name,omitempty"`
	// ISO timestamp
	Created *string `json:"created,omitempty"`
	GroupGuid *string `json:"group_guid,omitempty"`
	// ISO timestamp
	Modified *string `json:"modified,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	Guid *string `json:"guid,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewCampaignAllOf instantiates a new CampaignAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignAllOf() *CampaignAllOf {
	this := CampaignAllOf{}
	return &this
}

// NewCampaignAllOfWithDefaults instantiates a new CampaignAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignAllOfWithDefaults() *CampaignAllOf {
	this := CampaignAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CampaignAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CampaignAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CampaignAllOf) SetName(v string) {
	o.Name = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *CampaignAllOf) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOf) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CampaignAllOf) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *CampaignAllOf) SetCreated(v string) {
	o.Created = &v
}

// GetGroupGuid returns the GroupGuid field value if set, zero value otherwise.
func (o *CampaignAllOf) GetGroupGuid() string {
	if o == nil || o.GroupGuid == nil {
		var ret string
		return ret
	}
	return *o.GroupGuid
}

// GetGroupGuidOk returns a tuple with the GroupGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOf) GetGroupGuidOk() (*string, bool) {
	if o == nil || o.GroupGuid == nil {
		return nil, false
	}
	return o.GroupGuid, true
}

// HasGroupGuid returns a boolean if a field has been set.
func (o *CampaignAllOf) HasGroupGuid() bool {
	if o != nil && o.GroupGuid != nil {
		return true
	}

	return false
}

// SetGroupGuid gets a reference to the given string and assigns it to the GroupGuid field.
func (o *CampaignAllOf) SetGroupGuid(v string) {
	o.GroupGuid = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *CampaignAllOf) GetModified() string {
	if o == nil || o.Modified == nil {
		var ret string
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOf) GetModifiedOk() (*string, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *CampaignAllOf) HasModified() bool {
	if o != nil && o.Modified != nil {
		return true
	}

	return false
}

// SetModified gets a reference to the given string and assigns it to the Modified field.
func (o *CampaignAllOf) SetModified(v string) {
	o.Modified = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *CampaignAllOf) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOf) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *CampaignAllOf) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *CampaignAllOf) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *CampaignAllOf) GetGuid() string {
	if o == nil || o.Guid == nil {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOf) GetGuidOk() (*string, bool) {
	if o == nil || o.Guid == nil {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *CampaignAllOf) HasGuid() bool {
	if o != nil && o.Guid != nil {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *CampaignAllOf) SetGuid(v string) {
	o.Guid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CampaignAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CampaignAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CampaignAllOf) SetDescription(v string) {
	o.Description = &v
}

func (o CampaignAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.GroupGuid != nil {
		toSerialize["group_guid"] = o.GroupGuid
	}
	if o.Modified != nil {
		toSerialize["modified"] = o.Modified
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.Guid != nil {
		toSerialize["guid"] = o.Guid
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableCampaignAllOf struct {
	value *CampaignAllOf
	isSet bool
}

func (v NullableCampaignAllOf) Get() *CampaignAllOf {
	return v.value
}

func (v *NullableCampaignAllOf) Set(val *CampaignAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignAllOf(val *CampaignAllOf) *NullableCampaignAllOf {
	return &NullableCampaignAllOf{value: val, isSet: true}
}

func (v NullableCampaignAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


