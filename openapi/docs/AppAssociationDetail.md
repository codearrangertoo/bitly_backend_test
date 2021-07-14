# AppAssociationDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | Pointer to **int32** |  | [optional] 
**Fallback** | Pointer to **bool** |  | [optional] 
**AppGuid** | Pointer to **string** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 

## Methods

### NewAppAssociationDetail

`func NewAppAssociationDetail() *AppAssociationDetail`

NewAppAssociationDetail instantiates a new AppAssociationDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAssociationDetailWithDefaults

`func NewAppAssociationDetailWithDefaults() *AppAssociationDetail`

NewAppAssociationDetailWithDefaults instantiates a new AppAssociationDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *AppAssociationDetail) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *AppAssociationDetail) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *AppAssociationDetail) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *AppAssociationDetail) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetFallback

`func (o *AppAssociationDetail) GetFallback() bool`

GetFallback returns the Fallback field if non-nil, zero value otherwise.

### GetFallbackOk

`func (o *AppAssociationDetail) GetFallbackOk() (*bool, bool)`

GetFallbackOk returns a tuple with the Fallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback

`func (o *AppAssociationDetail) SetFallback(v bool)`

SetFallback sets Fallback field to given value.

### HasFallback

`func (o *AppAssociationDetail) HasFallback() bool`

HasFallback returns a boolean if a field has been set.

### GetAppGuid

`func (o *AppAssociationDetail) GetAppGuid() string`

GetAppGuid returns the AppGuid field if non-nil, zero value otherwise.

### GetAppGuidOk

`func (o *AppAssociationDetail) GetAppGuidOk() (*string, bool)`

GetAppGuidOk returns a tuple with the AppGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppGuid

`func (o *AppAssociationDetail) SetAppGuid(v string)`

SetAppGuid sets AppGuid field to given value.

### HasAppGuid

`func (o *AppAssociationDetail) HasAppGuid() bool`

HasAppGuid returns a boolean if a field has been set.

### GetOs

`func (o *AppAssociationDetail) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *AppAssociationDetail) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *AppAssociationDetail) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *AppAssociationDetail) HasOs() bool`

HasOs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


