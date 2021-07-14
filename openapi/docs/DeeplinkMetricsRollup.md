# DeeplinkMetricsRollup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitReference** | Pointer to **string** | data returned up to this reference timestamp | [optional] 
**Ios** | Pointer to [**DeeplinkMetric**](DeeplinkMetric.md) |  | [optional] 
**Other** | Pointer to [**DeeplinkMetric**](DeeplinkMetric.md) |  | [optional] 
**Units** | Pointer to **int32** | the number of units queried for | [optional] [default to -1]
**Android** | Pointer to [**DeeplinkMetric**](DeeplinkMetric.md) |  | [optional] 
**Unit** | Pointer to [**TimeUnit**](TimeUnit.md) |  | [optional] [default to DAY]

## Methods

### NewDeeplinkMetricsRollup

`func NewDeeplinkMetricsRollup() *DeeplinkMetricsRollup`

NewDeeplinkMetricsRollup instantiates a new DeeplinkMetricsRollup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeeplinkMetricsRollupWithDefaults

`func NewDeeplinkMetricsRollupWithDefaults() *DeeplinkMetricsRollup`

NewDeeplinkMetricsRollupWithDefaults instantiates a new DeeplinkMetricsRollup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitReference

`func (o *DeeplinkMetricsRollup) GetUnitReference() string`

GetUnitReference returns the UnitReference field if non-nil, zero value otherwise.

### GetUnitReferenceOk

`func (o *DeeplinkMetricsRollup) GetUnitReferenceOk() (*string, bool)`

GetUnitReferenceOk returns a tuple with the UnitReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitReference

`func (o *DeeplinkMetricsRollup) SetUnitReference(v string)`

SetUnitReference sets UnitReference field to given value.

### HasUnitReference

`func (o *DeeplinkMetricsRollup) HasUnitReference() bool`

HasUnitReference returns a boolean if a field has been set.

### GetIos

`func (o *DeeplinkMetricsRollup) GetIos() DeeplinkMetric`

GetIos returns the Ios field if non-nil, zero value otherwise.

### GetIosOk

`func (o *DeeplinkMetricsRollup) GetIosOk() (*DeeplinkMetric, bool)`

GetIosOk returns a tuple with the Ios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIos

`func (o *DeeplinkMetricsRollup) SetIos(v DeeplinkMetric)`

SetIos sets Ios field to given value.

### HasIos

`func (o *DeeplinkMetricsRollup) HasIos() bool`

HasIos returns a boolean if a field has been set.

### GetOther

`func (o *DeeplinkMetricsRollup) GetOther() DeeplinkMetric`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *DeeplinkMetricsRollup) GetOtherOk() (*DeeplinkMetric, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *DeeplinkMetricsRollup) SetOther(v DeeplinkMetric)`

SetOther sets Other field to given value.

### HasOther

`func (o *DeeplinkMetricsRollup) HasOther() bool`

HasOther returns a boolean if a field has been set.

### GetUnits

`func (o *DeeplinkMetricsRollup) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *DeeplinkMetricsRollup) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *DeeplinkMetricsRollup) SetUnits(v int32)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *DeeplinkMetricsRollup) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetAndroid

`func (o *DeeplinkMetricsRollup) GetAndroid() DeeplinkMetric`

GetAndroid returns the Android field if non-nil, zero value otherwise.

### GetAndroidOk

`func (o *DeeplinkMetricsRollup) GetAndroidOk() (*DeeplinkMetric, bool)`

GetAndroidOk returns a tuple with the Android field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroid

`func (o *DeeplinkMetricsRollup) SetAndroid(v DeeplinkMetric)`

SetAndroid sets Android field to given value.

### HasAndroid

`func (o *DeeplinkMetricsRollup) HasAndroid() bool`

HasAndroid returns a boolean if a field has been set.

### GetUnit

`func (o *DeeplinkMetricsRollup) GetUnit() TimeUnit`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *DeeplinkMetricsRollup) GetUnitOk() (*TimeUnit, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *DeeplinkMetricsRollup) SetUnit(v TimeUnit)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *DeeplinkMetricsRollup) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


