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

// Views struct for Views
type Views struct {
	LaunchpadViews *[]ViewMetric `json:"launchpad_views,omitempty"`
	Units *int32 `json:"units,omitempty"`
	Facet *string `json:"facet,omitempty"`
	Unit *string `json:"unit,omitempty"`
	UnitReference *string `json:"unit_reference,omitempty"`
}

// NewViews instantiates a new Views object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViews() *Views {
	this := Views{}
	return &this
}

// NewViewsWithDefaults instantiates a new Views object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewsWithDefaults() *Views {
	this := Views{}
	return &this
}

// GetLaunchpadViews returns the LaunchpadViews field value if set, zero value otherwise.
func (o *Views) GetLaunchpadViews() []ViewMetric {
	if o == nil || o.LaunchpadViews == nil {
		var ret []ViewMetric
		return ret
	}
	return *o.LaunchpadViews
}

// GetLaunchpadViewsOk returns a tuple with the LaunchpadViews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Views) GetLaunchpadViewsOk() (*[]ViewMetric, bool) {
	if o == nil || o.LaunchpadViews == nil {
		return nil, false
	}
	return o.LaunchpadViews, true
}

// HasLaunchpadViews returns a boolean if a field has been set.
func (o *Views) HasLaunchpadViews() bool {
	if o != nil && o.LaunchpadViews != nil {
		return true
	}

	return false
}

// SetLaunchpadViews gets a reference to the given []ViewMetric and assigns it to the LaunchpadViews field.
func (o *Views) SetLaunchpadViews(v []ViewMetric) {
	o.LaunchpadViews = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *Views) GetUnits() int32 {
	if o == nil || o.Units == nil {
		var ret int32
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Views) GetUnitsOk() (*int32, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *Views) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given int32 and assigns it to the Units field.
func (o *Views) SetUnits(v int32) {
	o.Units = &v
}

// GetFacet returns the Facet field value if set, zero value otherwise.
func (o *Views) GetFacet() string {
	if o == nil || o.Facet == nil {
		var ret string
		return ret
	}
	return *o.Facet
}

// GetFacetOk returns a tuple with the Facet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Views) GetFacetOk() (*string, bool) {
	if o == nil || o.Facet == nil {
		return nil, false
	}
	return o.Facet, true
}

// HasFacet returns a boolean if a field has been set.
func (o *Views) HasFacet() bool {
	if o != nil && o.Facet != nil {
		return true
	}

	return false
}

// SetFacet gets a reference to the given string and assigns it to the Facet field.
func (o *Views) SetFacet(v string) {
	o.Facet = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *Views) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Views) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *Views) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *Views) SetUnit(v string) {
	o.Unit = &v
}

// GetUnitReference returns the UnitReference field value if set, zero value otherwise.
func (o *Views) GetUnitReference() string {
	if o == nil || o.UnitReference == nil {
		var ret string
		return ret
	}
	return *o.UnitReference
}

// GetUnitReferenceOk returns a tuple with the UnitReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Views) GetUnitReferenceOk() (*string, bool) {
	if o == nil || o.UnitReference == nil {
		return nil, false
	}
	return o.UnitReference, true
}

// HasUnitReference returns a boolean if a field has been set.
func (o *Views) HasUnitReference() bool {
	if o != nil && o.UnitReference != nil {
		return true
	}

	return false
}

// SetUnitReference gets a reference to the given string and assigns it to the UnitReference field.
func (o *Views) SetUnitReference(v string) {
	o.UnitReference = &v
}

func (o Views) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LaunchpadViews != nil {
		toSerialize["launchpad_views"] = o.LaunchpadViews
	}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	if o.Facet != nil {
		toSerialize["facet"] = o.Facet
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.UnitReference != nil {
		toSerialize["unit_reference"] = o.UnitReference
	}
	return json.Marshal(toSerialize)
}

type NullableViews struct {
	value *Views
	isSet bool
}

func (v NullableViews) Get() *Views {
	return v.value
}

func (v *NullableViews) Set(val *Views) {
	v.value = val
	v.isSet = true
}

func (v NullableViews) IsSet() bool {
	return v.isSet
}

func (v *NullableViews) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViews(val *Views) *NullableViews {
	return &NullableViews{value: val, isSet: true}
}

func (v NullableViews) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViews) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


