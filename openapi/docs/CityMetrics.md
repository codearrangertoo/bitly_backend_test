# CityMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Units** | Pointer to **int32** |  | [optional] 
**Facet** | Pointer to **string** |  | [optional] 
**UnitReference** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**Metrics** | Pointer to [**[]CityMetric**](CityMetric.md) |  | [optional] 
**OtherMetrics** | Pointer to [**OtherMetrics**](OtherMetrics.md) |  | [optional] 

## Methods

### NewCityMetrics

`func NewCityMetrics() *CityMetrics`

NewCityMetrics instantiates a new CityMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCityMetricsWithDefaults

`func NewCityMetricsWithDefaults() *CityMetrics`

NewCityMetricsWithDefaults instantiates a new CityMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnits

`func (o *CityMetrics) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *CityMetrics) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *CityMetrics) SetUnits(v int32)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *CityMetrics) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetFacet

`func (o *CityMetrics) GetFacet() string`

GetFacet returns the Facet field if non-nil, zero value otherwise.

### GetFacetOk

`func (o *CityMetrics) GetFacetOk() (*string, bool)`

GetFacetOk returns a tuple with the Facet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacet

`func (o *CityMetrics) SetFacet(v string)`

SetFacet sets Facet field to given value.

### HasFacet

`func (o *CityMetrics) HasFacet() bool`

HasFacet returns a boolean if a field has been set.

### GetUnitReference

`func (o *CityMetrics) GetUnitReference() string`

GetUnitReference returns the UnitReference field if non-nil, zero value otherwise.

### GetUnitReferenceOk

`func (o *CityMetrics) GetUnitReferenceOk() (*string, bool)`

GetUnitReferenceOk returns a tuple with the UnitReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitReference

`func (o *CityMetrics) SetUnitReference(v string)`

SetUnitReference sets UnitReference field to given value.

### HasUnitReference

`func (o *CityMetrics) HasUnitReference() bool`

HasUnitReference returns a boolean if a field has been set.

### GetUnit

`func (o *CityMetrics) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CityMetrics) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CityMetrics) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *CityMetrics) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetMetrics

`func (o *CityMetrics) GetMetrics() []CityMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *CityMetrics) GetMetricsOk() (*[]CityMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *CityMetrics) SetMetrics(v []CityMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *CityMetrics) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetOtherMetrics

`func (o *CityMetrics) GetOtherMetrics() OtherMetrics`

GetOtherMetrics returns the OtherMetrics field if non-nil, zero value otherwise.

### GetOtherMetricsOk

`func (o *CityMetrics) GetOtherMetricsOk() (*OtherMetrics, bool)`

GetOtherMetricsOk returns a tuple with the OtherMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherMetrics

`func (o *CityMetrics) SetOtherMetrics(v OtherMetrics)`

SetOtherMetrics sets OtherMetrics field to given value.

### HasOtherMetrics

`func (o *CityMetrics) HasOtherMetrics() bool`

HasOtherMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


