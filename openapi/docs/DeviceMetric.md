# DeviceMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clicks** | Pointer to **int32** |  | [optional] 
**DeviceType** | Pointer to **string** |  | [optional] 

## Methods

### NewDeviceMetric

`func NewDeviceMetric() *DeviceMetric`

NewDeviceMetric instantiates a new DeviceMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceMetricWithDefaults

`func NewDeviceMetricWithDefaults() *DeviceMetric`

NewDeviceMetricWithDefaults instantiates a new DeviceMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClicks

`func (o *DeviceMetric) GetClicks() int32`

GetClicks returns the Clicks field if non-nil, zero value otherwise.

### GetClicksOk

`func (o *DeviceMetric) GetClicksOk() (*int32, bool)`

GetClicksOk returns a tuple with the Clicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicks

`func (o *DeviceMetric) SetClicks(v int32)`

SetClicks sets Clicks field to given value.

### HasClicks

`func (o *DeviceMetric) HasClicks() bool`

HasClicks returns a boolean if a field has been set.

### GetDeviceType

`func (o *DeviceMetric) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DeviceMetric) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DeviceMetric) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *DeviceMetric) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


