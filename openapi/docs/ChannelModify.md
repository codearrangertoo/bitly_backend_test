# ChannelModify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupGuid** | Pointer to **string** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Modified** | Pointer to **string** | ISO_TIMESTAMP | [optional] 
**Created** | Pointer to **string** | ISO TIMESTAMP | [optional] 
**Bitlinks** | Pointer to [**[]BaseChannelBitlink**](BaseChannelBitlink.md) |  | [optional] 

## Methods

### NewChannelModify

`func NewChannelModify() *ChannelModify`

NewChannelModify instantiates a new ChannelModify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelModifyWithDefaults

`func NewChannelModifyWithDefaults() *ChannelModify`

NewChannelModifyWithDefaults instantiates a new ChannelModify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupGuid

`func (o *ChannelModify) GetGroupGuid() string`

GetGroupGuid returns the GroupGuid field if non-nil, zero value otherwise.

### GetGroupGuidOk

`func (o *ChannelModify) GetGroupGuidOk() (*string, bool)`

GetGroupGuidOk returns a tuple with the GroupGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupGuid

`func (o *ChannelModify) SetGroupGuid(v string)`

SetGroupGuid sets GroupGuid field to given value.

### HasGroupGuid

`func (o *ChannelModify) HasGroupGuid() bool`

HasGroupGuid returns a boolean if a field has been set.

### GetGuid

`func (o *ChannelModify) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ChannelModify) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ChannelModify) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ChannelModify) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetName

`func (o *ChannelModify) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChannelModify) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChannelModify) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChannelModify) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModified

`func (o *ChannelModify) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ChannelModify) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ChannelModify) SetModified(v string)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ChannelModify) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetCreated

`func (o *ChannelModify) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ChannelModify) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ChannelModify) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ChannelModify) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetBitlinks

`func (o *ChannelModify) GetBitlinks() []BaseChannelBitlink`

GetBitlinks returns the Bitlinks field if non-nil, zero value otherwise.

### GetBitlinksOk

`func (o *ChannelModify) GetBitlinksOk() (*[]BaseChannelBitlink, bool)`

GetBitlinksOk returns a tuple with the Bitlinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitlinks

`func (o *ChannelModify) SetBitlinks(v []BaseChannelBitlink)`

SetBitlinks sets Bitlinks field to given value.

### HasBitlinks

`func (o *ChannelModify) HasBitlinks() bool`

HasBitlinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


