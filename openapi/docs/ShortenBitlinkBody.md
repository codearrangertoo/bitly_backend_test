# ShortenBitlinkBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**References** | Pointer to **map[string]string** |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Deeplinks** | Pointer to [**[]DeeplinkRule**](DeeplinkRule.md) |  | [optional] 
**LongUrl** | Pointer to **string** |  | [optional] 
**CustomBitlinks** | Pointer to **[]string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 

## Methods

### NewShortenBitlinkBody

`func NewShortenBitlinkBody() *ShortenBitlinkBody`

NewShortenBitlinkBody instantiates a new ShortenBitlinkBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShortenBitlinkBodyWithDefaults

`func NewShortenBitlinkBodyWithDefaults() *ShortenBitlinkBody`

NewShortenBitlinkBodyWithDefaults instantiates a new ShortenBitlinkBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferences

`func (o *ShortenBitlinkBody) GetReferences() map[string]string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ShortenBitlinkBody) GetReferencesOk() (*map[string]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ShortenBitlinkBody) SetReferences(v map[string]string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ShortenBitlinkBody) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetArchived

`func (o *ShortenBitlinkBody) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ShortenBitlinkBody) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ShortenBitlinkBody) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *ShortenBitlinkBody) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetTags

`func (o *ShortenBitlinkBody) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ShortenBitlinkBody) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ShortenBitlinkBody) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ShortenBitlinkBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ShortenBitlinkBody) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ShortenBitlinkBody) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ShortenBitlinkBody) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ShortenBitlinkBody) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeeplinks

`func (o *ShortenBitlinkBody) GetDeeplinks() []DeeplinkRule`

GetDeeplinks returns the Deeplinks field if non-nil, zero value otherwise.

### GetDeeplinksOk

`func (o *ShortenBitlinkBody) GetDeeplinksOk() (*[]DeeplinkRule, bool)`

GetDeeplinksOk returns a tuple with the Deeplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeeplinks

`func (o *ShortenBitlinkBody) SetDeeplinks(v []DeeplinkRule)`

SetDeeplinks sets Deeplinks field to given value.

### HasDeeplinks

`func (o *ShortenBitlinkBody) HasDeeplinks() bool`

HasDeeplinks returns a boolean if a field has been set.

### GetLongUrl

`func (o *ShortenBitlinkBody) GetLongUrl() string`

GetLongUrl returns the LongUrl field if non-nil, zero value otherwise.

### GetLongUrlOk

`func (o *ShortenBitlinkBody) GetLongUrlOk() (*string, bool)`

GetLongUrlOk returns a tuple with the LongUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongUrl

`func (o *ShortenBitlinkBody) SetLongUrl(v string)`

SetLongUrl sets LongUrl field to given value.

### HasLongUrl

`func (o *ShortenBitlinkBody) HasLongUrl() bool`

HasLongUrl returns a boolean if a field has been set.

### GetCustomBitlinks

`func (o *ShortenBitlinkBody) GetCustomBitlinks() []string`

GetCustomBitlinks returns the CustomBitlinks field if non-nil, zero value otherwise.

### GetCustomBitlinksOk

`func (o *ShortenBitlinkBody) GetCustomBitlinksOk() (*[]string, bool)`

GetCustomBitlinksOk returns a tuple with the CustomBitlinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBitlinks

`func (o *ShortenBitlinkBody) SetCustomBitlinks(v []string)`

SetCustomBitlinks sets CustomBitlinks field to given value.

### HasCustomBitlinks

`func (o *ShortenBitlinkBody) HasCustomBitlinks() bool`

HasCustomBitlinks returns a boolean if a field has been set.

### GetLink

`func (o *ShortenBitlinkBody) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *ShortenBitlinkBody) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *ShortenBitlinkBody) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *ShortenBitlinkBody) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetId

`func (o *ShortenBitlinkBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShortenBitlinkBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShortenBitlinkBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShortenBitlinkBody) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


