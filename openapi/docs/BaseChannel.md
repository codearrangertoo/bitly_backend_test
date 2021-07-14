# BaseChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupGuid** | Pointer to **string** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Modified** | Pointer to **string** | ISO_TIMESTAMP | [optional] 
**Created** | Pointer to **string** | ISO TIMESTAMP | [optional] 

## Methods

### NewBaseChannel

`func NewBaseChannel() *BaseChannel`

NewBaseChannel instantiates a new BaseChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseChannelWithDefaults

`func NewBaseChannelWithDefaults() *BaseChannel`

NewBaseChannelWithDefaults instantiates a new BaseChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupGuid

`func (o *BaseChannel) GetGroupGuid() string`

GetGroupGuid returns the GroupGuid field if non-nil, zero value otherwise.

### GetGroupGuidOk

`func (o *BaseChannel) GetGroupGuidOk() (*string, bool)`

GetGroupGuidOk returns a tuple with the GroupGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupGuid

`func (o *BaseChannel) SetGroupGuid(v string)`

SetGroupGuid sets GroupGuid field to given value.

### HasGroupGuid

`func (o *BaseChannel) HasGroupGuid() bool`

HasGroupGuid returns a boolean if a field has been set.

### GetGuid

`func (o *BaseChannel) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *BaseChannel) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *BaseChannel) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *BaseChannel) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetName

`func (o *BaseChannel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseChannel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseChannel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseChannel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModified

`func (o *BaseChannel) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *BaseChannel) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *BaseChannel) SetModified(v string)`

SetModified sets Modified field to given value.

### HasModified

`func (o *BaseChannel) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetCreated

`func (o *BaseChannel) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BaseChannel) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BaseChannel) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *BaseChannel) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


