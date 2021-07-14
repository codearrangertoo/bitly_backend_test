# Views

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LaunchpadViews** | Pointer to [**[]ViewMetric**](ViewMetric.md) |  | [optional] 
**Units** | Pointer to **int32** |  | [optional] 
**Facet** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**UnitReference** | Pointer to **string** |  | [optional] 

## Methods

### NewViews

`func NewViews() *Views`

NewViews instantiates a new Views object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewsWithDefaults

`func NewViewsWithDefaults() *Views`

NewViewsWithDefaults instantiates a new Views object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLaunchpadViews

`func (o *Views) GetLaunchpadViews() []ViewMetric`

GetLaunchpadViews returns the LaunchpadViews field if non-nil, zero value otherwise.

### GetLaunchpadViewsOk

`func (o *Views) GetLaunchpadViewsOk() (*[]ViewMetric, bool)`

GetLaunchpadViewsOk returns a tuple with the LaunchpadViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchpadViews

`func (o *Views) SetLaunchpadViews(v []ViewMetric)`

SetLaunchpadViews sets LaunchpadViews field to given value.

### HasLaunchpadViews

`func (o *Views) HasLaunchpadViews() bool`

HasLaunchpadViews returns a boolean if a field has been set.

### GetUnits

`func (o *Views) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *Views) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *Views) SetUnits(v int32)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *Views) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetFacet

`func (o *Views) GetFacet() string`

GetFacet returns the Facet field if non-nil, zero value otherwise.

### GetFacetOk

`func (o *Views) GetFacetOk() (*string, bool)`

GetFacetOk returns a tuple with the Facet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacet

`func (o *Views) SetFacet(v string)`

SetFacet sets Facet field to given value.

### HasFacet

`func (o *Views) HasFacet() bool`

HasFacet returns a boolean if a field has been set.

### GetUnit

`func (o *Views) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Views) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Views) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Views) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUnitReference

`func (o *Views) GetUnitReference() string`

GetUnitReference returns the UnitReference field if non-nil, zero value otherwise.

### GetUnitReferenceOk

`func (o *Views) GetUnitReferenceOk() (*string, bool)`

GetUnitReferenceOk returns a tuple with the UnitReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitReference

`func (o *Views) SetUnitReference(v string)`

SetUnitReference sets UnitReference field to given value.

### HasUnitReference

`func (o *Views) HasUnitReference() bool`

HasUnitReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


