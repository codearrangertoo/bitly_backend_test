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

// BSDSearchResult struct for BSDSearchResult
type BSDSearchResult struct {
	Type *string `json:"type,omitempty"`
	Domain *string `json:"domain,omitempty"`
	Subdomain *string `json:"subdomain,omitempty"`
	Link *string `json:"link,omitempty"`
	Zone *string `json:"zone,omitempty"`
}

// NewBSDSearchResult instantiates a new BSDSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBSDSearchResult() *BSDSearchResult {
	this := BSDSearchResult{}
	return &this
}

// NewBSDSearchResultWithDefaults instantiates a new BSDSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBSDSearchResultWithDefaults() *BSDSearchResult {
	this := BSDSearchResult{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BSDSearchResult) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BSDSearchResult) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BSDSearchResult) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BSDSearchResult) SetType(v string) {
	o.Type = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *BSDSearchResult) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BSDSearchResult) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *BSDSearchResult) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *BSDSearchResult) SetDomain(v string) {
	o.Domain = &v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *BSDSearchResult) GetSubdomain() string {
	if o == nil || o.Subdomain == nil {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BSDSearchResult) GetSubdomainOk() (*string, bool) {
	if o == nil || o.Subdomain == nil {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *BSDSearchResult) HasSubdomain() bool {
	if o != nil && o.Subdomain != nil {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *BSDSearchResult) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *BSDSearchResult) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BSDSearchResult) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *BSDSearchResult) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *BSDSearchResult) SetLink(v string) {
	o.Link = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *BSDSearchResult) GetZone() string {
	if o == nil || o.Zone == nil {
		var ret string
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BSDSearchResult) GetZoneOk() (*string, bool) {
	if o == nil || o.Zone == nil {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *BSDSearchResult) HasZone() bool {
	if o != nil && o.Zone != nil {
		return true
	}

	return false
}

// SetZone gets a reference to the given string and assigns it to the Zone field.
func (o *BSDSearchResult) SetZone(v string) {
	o.Zone = &v
}

func (o BSDSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Subdomain != nil {
		toSerialize["subdomain"] = o.Subdomain
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}
	if o.Zone != nil {
		toSerialize["zone"] = o.Zone
	}
	return json.Marshal(toSerialize)
}

type NullableBSDSearchResult struct {
	value *BSDSearchResult
	isSet bool
}

func (v NullableBSDSearchResult) Get() *BSDSearchResult {
	return v.value
}

func (v *NullableBSDSearchResult) Set(val *BSDSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableBSDSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableBSDSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBSDSearchResult(val *BSDSearchResult) *NullableBSDSearchResult {
	return &NullableBSDSearchResult{value: val, isSet: true}
}

func (v NullableBSDSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBSDSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


