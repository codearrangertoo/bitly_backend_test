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

// Shorten struct for Shorten
type Shorten struct {
	GroupGuid *string `json:"group_guid,omitempty"`
	Domain *string `json:"domain,omitempty"`
	LongUrl string `json:"long_url"`
}

// NewShorten instantiates a new Shorten object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShorten(longUrl string) *Shorten {
	this := Shorten{}
	var domain string = "bit.ly"
	this.Domain = &domain
	this.LongUrl = longUrl
	return &this
}

// NewShortenWithDefaults instantiates a new Shorten object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShortenWithDefaults() *Shorten {
	this := Shorten{}
	var domain string = "bit.ly"
	this.Domain = &domain
	return &this
}

// GetGroupGuid returns the GroupGuid field value if set, zero value otherwise.
func (o *Shorten) GetGroupGuid() string {
	if o == nil || o.GroupGuid == nil {
		var ret string
		return ret
	}
	return *o.GroupGuid
}

// GetGroupGuidOk returns a tuple with the GroupGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shorten) GetGroupGuidOk() (*string, bool) {
	if o == nil || o.GroupGuid == nil {
		return nil, false
	}
	return o.GroupGuid, true
}

// HasGroupGuid returns a boolean if a field has been set.
func (o *Shorten) HasGroupGuid() bool {
	if o != nil && o.GroupGuid != nil {
		return true
	}

	return false
}

// SetGroupGuid gets a reference to the given string and assigns it to the GroupGuid field.
func (o *Shorten) SetGroupGuid(v string) {
	o.GroupGuid = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *Shorten) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shorten) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *Shorten) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *Shorten) SetDomain(v string) {
	o.Domain = &v
}

// GetLongUrl returns the LongUrl field value
func (o *Shorten) GetLongUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LongUrl
}

// GetLongUrlOk returns a tuple with the LongUrl field value
// and a boolean to check if the value has been set.
func (o *Shorten) GetLongUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LongUrl, true
}

// SetLongUrl sets field value
func (o *Shorten) SetLongUrl(v string) {
	o.LongUrl = v
}

func (o Shorten) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupGuid != nil {
		toSerialize["group_guid"] = o.GroupGuid
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if true {
		toSerialize["long_url"] = o.LongUrl
	}
	return json.Marshal(toSerialize)
}

type NullableShorten struct {
	value *Shorten
	isSet bool
}

func (v NullableShorten) Get() *Shorten {
	return v.value
}

func (v *NullableShorten) Set(val *Shorten) {
	v.value = val
	v.isSet = true
}

func (v NullableShorten) IsSet() bool {
	return v.isSet
}

func (v *NullableShorten) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShorten(val *Shorten) *NullableShorten {
	return &NullableShorten{value: val, isSet: true}
}

func (v NullableShorten) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShorten) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


