# PlatformLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | Pointer to **string** |  | [optional] 
**Methods** | Pointer to [**[]MethodLimit**](MethodLimit.md) |  | [optional] 

## Methods

### NewPlatformLimit

`func NewPlatformLimit() *PlatformLimit`

NewPlatformLimit instantiates a new PlatformLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformLimitWithDefaults

`func NewPlatformLimitWithDefaults() *PlatformLimit`

NewPlatformLimitWithDefaults instantiates a new PlatformLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *PlatformLimit) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *PlatformLimit) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *PlatformLimit) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *PlatformLimit) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetMethods

`func (o *PlatformLimit) GetMethods() []MethodLimit`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *PlatformLimit) GetMethodsOk() (*[]MethodLimit, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *PlatformLimit) SetMethods(v []MethodLimit)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *PlatformLimit) HasMethods() bool`

HasMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


