# DeviceMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Units** | Pointer to **int32** |  | [optional] 
**Facet** | Pointer to **string** |  | [optional] 
**UnitReference** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**Metrics** | Pointer to [**[]DeviceMetric**](DeviceMetric.md) |  | [optional] 

## Methods

### NewDeviceMetrics

`func NewDeviceMetrics() *DeviceMetrics`

NewDeviceMetrics instantiates a new DeviceMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceMetricsWithDefaults

`func NewDeviceMetricsWithDefaults() *DeviceMetrics`

NewDeviceMetricsWithDefaults instantiates a new DeviceMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnits

`func (o *DeviceMetrics) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *DeviceMetrics) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *DeviceMetrics) SetUnits(v int32)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *DeviceMetrics) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetFacet

`func (o *DeviceMetrics) GetFacet() string`

GetFacet returns the Facet field if non-nil, zero value otherwise.

### GetFacetOk

`func (o *DeviceMetrics) GetFacetOk() (*string, bool)`

GetFacetOk returns a tuple with the Facet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacet

`func (o *DeviceMetrics) SetFacet(v string)`

SetFacet sets Facet field to given value.

### HasFacet

`func (o *DeviceMetrics) HasFacet() bool`

HasFacet returns a boolean if a field has been set.

### GetUnitReference

`func (o *DeviceMetrics) GetUnitReference() string`

GetUnitReference returns the UnitReference field if non-nil, zero value otherwise.

### GetUnitReferenceOk

`func (o *DeviceMetrics) GetUnitReferenceOk() (*string, bool)`

GetUnitReferenceOk returns a tuple with the UnitReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitReference

`func (o *DeviceMetrics) SetUnitReference(v string)`

SetUnitReference sets UnitReference field to given value.

### HasUnitReference

`func (o *DeviceMetrics) HasUnitReference() bool`

HasUnitReference returns a boolean if a field has been set.

### GetUnit

`func (o *DeviceMetrics) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *DeviceMetrics) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *DeviceMetrics) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *DeviceMetrics) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetMetrics

`func (o *DeviceMetrics) GetMetrics() []DeviceMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *DeviceMetrics) GetMetricsOk() (*[]DeviceMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *DeviceMetrics) SetMetrics(v []DeviceMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *DeviceMetrics) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


