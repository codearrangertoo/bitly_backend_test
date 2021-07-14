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

// BillingAccountID struct for BillingAccountID
type BillingAccountID struct {
	BillingAccountId *string `json:"billing_account_id,omitempty"`
}

// NewBillingAccountID instantiates a new BillingAccountID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingAccountID() *BillingAccountID {
	this := BillingAccountID{}
	return &this
}

// NewBillingAccountIDWithDefaults instantiates a new BillingAccountID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingAccountIDWithDefaults() *BillingAccountID {
	this := BillingAccountID{}
	return &this
}

// GetBillingAccountId returns the BillingAccountId field value if set, zero value otherwise.
func (o *BillingAccountID) GetBillingAccountId() string {
	if o == nil || o.BillingAccountId == nil {
		var ret string
		return ret
	}
	return *o.BillingAccountId
}

// GetBillingAccountIdOk returns a tuple with the BillingAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAccountID) GetBillingAccountIdOk() (*string, bool) {
	if o == nil || o.BillingAccountId == nil {
		return nil, false
	}
	return o.BillingAccountId, true
}

// HasBillingAccountId returns a boolean if a field has been set.
func (o *BillingAccountID) HasBillingAccountId() bool {
	if o != nil && o.BillingAccountId != nil {
		return true
	}

	return false
}

// SetBillingAccountId gets a reference to the given string and assigns it to the BillingAccountId field.
func (o *BillingAccountID) SetBillingAccountId(v string) {
	o.BillingAccountId = &v
}

func (o BillingAccountID) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BillingAccountId != nil {
		toSerialize["billing_account_id"] = o.BillingAccountId
	}
	return json.Marshal(toSerialize)
}

type NullableBillingAccountID struct {
	value *BillingAccountID
	isSet bool
}

func (v NullableBillingAccountID) Get() *BillingAccountID {
	return v.value
}

func (v *NullableBillingAccountID) Set(val *BillingAccountID) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingAccountID) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingAccountID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingAccountID(val *BillingAccountID) *NullableBillingAccountID {
	return &NullableBillingAccountID{value: val, isSet: true}
}

func (v NullableBillingAccountID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingAccountID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


