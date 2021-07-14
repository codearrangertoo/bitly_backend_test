# Clicks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Units** | Pointer to **int32** |  | [optional] 
**UnitReference** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**LinkClicks** | Pointer to [**[]LinkClicks**](LinkClicks.md) |  | [optional] 

## Methods

### NewClicks

`func NewClicks() *Clicks`

NewClicks instantiates a new Clicks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClicksWithDefaults

`func NewClicksWithDefaults() *Clicks`

NewClicksWithDefaults instantiates a new Clicks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnits

`func (o *Clicks) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *Clicks) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *Clicks) SetUnits(v int32)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *Clicks) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetUnitReference

`func (o *Clicks) GetUnitReference() string`

GetUnitReference returns the UnitReference field if non-nil, zero value otherwise.

### GetUnitReferenceOk

`func (o *Clicks) GetUnitReferenceOk() (*string, bool)`

GetUnitReferenceOk returns a tuple with the UnitReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitReference

`func (o *Clicks) SetUnitReference(v string)`

SetUnitReference sets UnitReference field to given value.

### HasUnitReference

`func (o *Clicks) HasUnitReference() bool`

HasUnitReference returns a boolean if a field has been set.

### GetUnit

`func (o *Clicks) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Clicks) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Clicks) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Clicks) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetLinkClicks

`func (o *Clicks) GetLinkClicks() []LinkClicks`

GetLinkClicks returns the LinkClicks field if non-nil, zero value otherwise.

### GetLinkClicksOk

`func (o *Clicks) GetLinkClicksOk() (*[]LinkClicks, bool)`

GetLinkClicksOk returns a tuple with the LinkClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkClicks

`func (o *Clicks) SetLinkClicks(v []LinkClicks)`

SetLinkClicks sets LinkClicks field to given value.

### HasLinkClicks

`func (o *Clicks) HasLinkClicks() bool`

HasLinkClicks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


