# BitlinkUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Deeplinks** | Pointer to [**[]DeeplinkRule**](DeeplinkRule.md) |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**LongUrl** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**CustomBitlinks** | Pointer to **[]string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**LaunchpadIds** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 

## Methods

### NewBitlinkUpdate

`func NewBitlinkUpdate() *BitlinkUpdate`

NewBitlinkUpdate instantiates a new BitlinkUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBitlinkUpdateWithDefaults

`func NewBitlinkUpdateWithDefaults() *BitlinkUpdate`

NewBitlinkUpdateWithDefaults instantiates a new BitlinkUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *BitlinkUpdate) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *BitlinkUpdate) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *BitlinkUpdate) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *BitlinkUpdate) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetTags

`func (o *BitlinkUpdate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BitlinkUpdate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BitlinkUpdate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BitlinkUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BitlinkUpdate) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BitlinkUpdate) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BitlinkUpdate) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BitlinkUpdate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetTitle

`func (o *BitlinkUpdate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BitlinkUpdate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BitlinkUpdate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BitlinkUpdate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDeeplinks

`func (o *BitlinkUpdate) GetDeeplinks() []DeeplinkRule`

GetDeeplinks returns the Deeplinks field if non-nil, zero value otherwise.

### GetDeeplinksOk

`func (o *BitlinkUpdate) GetDeeplinksOk() (*[]DeeplinkRule, bool)`

GetDeeplinksOk returns a tuple with the Deeplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeeplinks

`func (o *BitlinkUpdate) SetDeeplinks(v []DeeplinkRule)`

SetDeeplinks sets Deeplinks field to given value.

### HasDeeplinks

`func (o *BitlinkUpdate) HasDeeplinks() bool`

HasDeeplinks returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BitlinkUpdate) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BitlinkUpdate) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BitlinkUpdate) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BitlinkUpdate) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLongUrl

`func (o *BitlinkUpdate) GetLongUrl() string`

GetLongUrl returns the LongUrl field if non-nil, zero value otherwise.

### GetLongUrlOk

`func (o *BitlinkUpdate) GetLongUrlOk() (*string, bool)`

GetLongUrlOk returns a tuple with the LongUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongUrl

`func (o *BitlinkUpdate) SetLongUrl(v string)`

SetLongUrl sets LongUrl field to given value.

### HasLongUrl

`func (o *BitlinkUpdate) HasLongUrl() bool`

HasLongUrl returns a boolean if a field has been set.

### GetClientId

`func (o *BitlinkUpdate) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *BitlinkUpdate) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *BitlinkUpdate) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *BitlinkUpdate) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetCustomBitlinks

`func (o *BitlinkUpdate) GetCustomBitlinks() []string`

GetCustomBitlinks returns the CustomBitlinks field if non-nil, zero value otherwise.

### GetCustomBitlinksOk

`func (o *BitlinkUpdate) GetCustomBitlinksOk() (*[]string, bool)`

GetCustomBitlinksOk returns a tuple with the CustomBitlinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBitlinks

`func (o *BitlinkUpdate) SetCustomBitlinks(v []string)`

SetCustomBitlinks sets CustomBitlinks field to given value.

### HasCustomBitlinks

`func (o *BitlinkUpdate) HasCustomBitlinks() bool`

HasCustomBitlinks returns a boolean if a field has been set.

### GetLink

`func (o *BitlinkUpdate) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *BitlinkUpdate) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *BitlinkUpdate) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *BitlinkUpdate) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetLaunchpadIds

`func (o *BitlinkUpdate) GetLaunchpadIds() []string`

GetLaunchpadIds returns the LaunchpadIds field if non-nil, zero value otherwise.

### GetLaunchpadIdsOk

`func (o *BitlinkUpdate) GetLaunchpadIdsOk() (*[]string, bool)`

GetLaunchpadIdsOk returns a tuple with the LaunchpadIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchpadIds

`func (o *BitlinkUpdate) SetLaunchpadIds(v []string)`

SetLaunchpadIds sets LaunchpadIds field to given value.

### HasLaunchpadIds

`func (o *BitlinkUpdate) HasLaunchpadIds() bool`

HasLaunchpadIds returns a boolean if a field has been set.

### GetId

`func (o *BitlinkUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BitlinkUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BitlinkUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BitlinkUpdate) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


