# BaseMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Units** | Pointer to **int32** |  | [optional] 
**Facet** | Pointer to **string** |  | [optional] 
**UnitReference** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 

## Methods

### NewBaseMetrics

`func NewBaseMetrics() *BaseMetrics`

NewBaseMetrics instantiates a new BaseMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseMetricsWithDefaults

`func NewBaseMetricsWithDefaults() *BaseMetrics`

NewBaseMetricsWithDefaults instantiates a new BaseMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnits

`func (o *BaseMetrics) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *BaseMetrics) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *BaseMetrics) SetUnits(v int32)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *BaseMetrics) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetFacet

`func (o *BaseMetrics) GetFacet() string`

GetFacet returns the Facet field if non-nil, zero value otherwise.

### GetFacetOk

`func (o *BaseMetrics) GetFacetOk() (*string, bool)`

GetFacetOk returns a tuple with the Facet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacet

`func (o *BaseMetrics) SetFacet(v string)`

SetFacet sets Facet field to given value.

### HasFacet

`func (o *BaseMetrics) HasFacet() bool`

HasFacet returns a boolean if a field has been set.

### GetUnitReference

`func (o *BaseMetrics) GetUnitReference() string`

GetUnitReference returns the UnitReference field if non-nil, zero value otherwise.

### GetUnitReferenceOk

`func (o *BaseMetrics) GetUnitReferenceOk() (*string, bool)`

GetUnitReferenceOk returns a tuple with the UnitReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitReference

`func (o *BaseMetrics) SetUnitReference(v string)`

SetUnitReference sets UnitReference field to given value.

### HasUnitReference

`func (o *BaseMetrics) HasUnitReference() bool`

HasUnitReference returns a boolean if a field has been set.

### GetUnit

`func (o *BaseMetrics) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *BaseMetrics) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *BaseMetrics) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *BaseMetrics) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


