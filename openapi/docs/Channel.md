# Channel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupGuid** | Pointer to **string** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Modified** | Pointer to **string** | ISO_TIMESTAMP | [optional] 
**Created** | Pointer to **string** | ISO TIMESTAMP | [optional] 
**References** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewChannel

`func NewChannel() *Channel`

NewChannel instantiates a new Channel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelWithDefaults

`func NewChannelWithDefaults() *Channel`

NewChannelWithDefaults instantiates a new Channel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupGuid

`func (o *Channel) GetGroupGuid() string`

GetGroupGuid returns the GroupGuid field if non-nil, zero value otherwise.

### GetGroupGuidOk

`func (o *Channel) GetGroupGuidOk() (*string, bool)`

GetGroupGuidOk returns a tuple with the GroupGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupGuid

`func (o *Channel) SetGroupGuid(v string)`

SetGroupGuid sets GroupGuid field to given value.

### HasGroupGuid

`func (o *Channel) HasGroupGuid() bool`

HasGroupGuid returns a boolean if a field has been set.

### GetGuid

`func (o *Channel) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Channel) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Channel) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *Channel) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetName

`func (o *Channel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Channel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Channel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Channel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModified

`func (o *Channel) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Channel) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Channel) SetModified(v string)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Channel) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetCreated

`func (o *Channel) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Channel) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Channel) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Channel) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetReferences

`func (o *Channel) GetReferences() map[string]string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *Channel) GetReferencesOk() (*map[string]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *Channel) SetReferences(v map[string]string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *Channel) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


