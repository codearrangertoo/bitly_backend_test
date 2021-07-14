# CityViewMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Units** | Pointer to **int32** |  | [optional] 
**Facet** | Pointer to **string** |  | [optional] 
**UnitReference** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**LaunchpadViews** | Pointer to [**[]ViewMetric**](ViewMetric.md) |  | [optional] 
**OtherMetrics** | Pointer to [**OtherViewMetrics**](OtherViewMetrics.md) |  | [optional] 

## Methods

### NewCityViewMetrics

`func NewCityViewMetrics() *CityViewMetrics`

NewCityViewMetrics instantiates a new CityViewMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCityViewMetricsWithDefaults

`func NewCityViewMetricsWithDefaults() *CityViewMetrics`

NewCityViewMetricsWithDefaults instantiates a new CityViewMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnits

`func (o *CityViewMetrics) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *CityViewMetrics) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *CityViewMetrics) SetUnits(v int32)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *CityViewMetrics) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetFacet

`func (o *CityViewMetrics) GetFacet() string`

GetFacet returns the Facet field if non-nil, zero value otherwise.

### GetFacetOk

`func (o *CityViewMetrics) GetFacetOk() (*string, bool)`

GetFacetOk returns a tuple with the Facet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacet

`func (o *CityViewMetrics) SetFacet(v string)`

SetFacet sets Facet field to given value.

### HasFacet

`func (o *CityViewMetrics) HasFacet() bool`

HasFacet returns a boolean if a field has been set.

### GetUnitReference

`func (o *CityViewMetrics) GetUnitReference() string`

GetUnitReference returns the UnitReference field if non-nil, zero value otherwise.

### GetUnitReferenceOk

`func (o *CityViewMetrics) GetUnitReferenceOk() (*string, bool)`

GetUnitReferenceOk returns a tuple with the UnitReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitReference

`func (o *CityViewMetrics) SetUnitReference(v string)`

SetUnitReference sets UnitReference field to given value.

### HasUnitReference

`func (o *CityViewMetrics) HasUnitReference() bool`

HasUnitReference returns a boolean if a field has been set.

### GetUnit

`func (o *CityViewMetrics) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CityViewMetrics) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CityViewMetrics) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *CityViewMetrics) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetLaunchpadViews

`func (o *CityViewMetrics) GetLaunchpadViews() []ViewMetric`

GetLaunchpadViews returns the LaunchpadViews field if non-nil, zero value otherwise.

### GetLaunchpadViewsOk

`func (o *CityViewMetrics) GetLaunchpadViewsOk() (*[]ViewMetric, bool)`

GetLaunchpadViewsOk returns a tuple with the LaunchpadViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchpadViews

`func (o *CityViewMetrics) SetLaunchpadViews(v []ViewMetric)`

SetLaunchpadViews sets LaunchpadViews field to given value.

### HasLaunchpadViews

`func (o *CityViewMetrics) HasLaunchpadViews() bool`

HasLaunchpadViews returns a boolean if a field has been set.

### GetOtherMetrics

`func (o *CityViewMetrics) GetOtherMetrics() OtherViewMetrics`

GetOtherMetrics returns the OtherMetrics field if non-nil, zero value otherwise.

### GetOtherMetricsOk

`func (o *CityViewMetrics) GetOtherMetricsOk() (*OtherViewMetrics, bool)`

GetOtherMetricsOk returns a tuple with the OtherMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherMetrics

`func (o *CityViewMetrics) SetOtherMetrics(v OtherViewMetrics)`

SetOtherMetrics sets OtherMetrics field to given value.

### HasOtherMetrics

`func (o *CityViewMetrics) HasOtherMetrics() bool`

HasOtherMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


