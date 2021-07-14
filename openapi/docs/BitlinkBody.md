# BitlinkBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**References** | Pointer to **map[string]string** |  | [optional] 
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

### NewBitlinkBody

`func NewBitlinkBody() *BitlinkBody`

NewBitlinkBody instantiates a new BitlinkBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBitlinkBodyWithDefaults

`func NewBitlinkBodyWithDefaults() *BitlinkBody`

NewBitlinkBodyWithDefaults instantiates a new BitlinkBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferences

`func (o *BitlinkBody) GetReferences() map[string]string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *BitlinkBody) GetReferencesOk() (*map[string]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *BitlinkBody) SetReferences(v map[string]string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *BitlinkBody) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetArchived

`func (o *BitlinkBody) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *BitlinkBody) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *BitlinkBody) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *BitlinkBody) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetTags

`func (o *BitlinkBody) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BitlinkBody) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BitlinkBody) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BitlinkBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BitlinkBody) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BitlinkBody) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BitlinkBody) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BitlinkBody) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetTitle

`func (o *BitlinkBody) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BitlinkBody) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BitlinkBody) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BitlinkBody) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDeeplinks

`func (o *BitlinkBody) GetDeeplinks() []DeeplinkRule`

GetDeeplinks returns the Deeplinks field if non-nil, zero value otherwise.

### GetDeeplinksOk

`func (o *BitlinkBody) GetDeeplinksOk() (*[]DeeplinkRule, bool)`

GetDeeplinksOk returns a tuple with the Deeplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeeplinks

`func (o *BitlinkBody) SetDeeplinks(v []DeeplinkRule)`

SetDeeplinks sets Deeplinks field to given value.

### HasDeeplinks

`func (o *BitlinkBody) HasDeeplinks() bool`

HasDeeplinks returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BitlinkBody) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BitlinkBody) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BitlinkBody) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BitlinkBody) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLongUrl

`func (o *BitlinkBody) GetLongUrl() string`

GetLongUrl returns the LongUrl field if non-nil, zero value otherwise.

### GetLongUrlOk

`func (o *BitlinkBody) GetLongUrlOk() (*string, bool)`

GetLongUrlOk returns a tuple with the LongUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongUrl

`func (o *BitlinkBody) SetLongUrl(v string)`

SetLongUrl sets LongUrl field to given value.

### HasLongUrl

`func (o *BitlinkBody) HasLongUrl() bool`

HasLongUrl returns a boolean if a field has been set.

### GetClientId

`func (o *BitlinkBody) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *BitlinkBody) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *BitlinkBody) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *BitlinkBody) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetCustomBitlinks

`func (o *BitlinkBody) GetCustomBitlinks() []string`

GetCustomBitlinks returns the CustomBitlinks field if non-nil, zero value otherwise.

### GetCustomBitlinksOk

`func (o *BitlinkBody) GetCustomBitlinksOk() (*[]string, bool)`

GetCustomBitlinksOk returns a tuple with the CustomBitlinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBitlinks

`func (o *BitlinkBody) SetCustomBitlinks(v []string)`

SetCustomBitlinks sets CustomBitlinks field to given value.

### HasCustomBitlinks

`func (o *BitlinkBody) HasCustomBitlinks() bool`

HasCustomBitlinks returns a boolean if a field has been set.

### GetLink

`func (o *BitlinkBody) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *BitlinkBody) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *BitlinkBody) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *BitlinkBody) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetLaunchpadIds

`func (o *BitlinkBody) GetLaunchpadIds() []string`

GetLaunchpadIds returns the LaunchpadIds field if non-nil, zero value otherwise.

### GetLaunchpadIdsOk

`func (o *BitlinkBody) GetLaunchpadIdsOk() (*[]string, bool)`

GetLaunchpadIdsOk returns a tuple with the LaunchpadIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchpadIds

`func (o *BitlinkBody) SetLaunchpadIds(v []string)`

SetLaunchpadIds sets LaunchpadIds field to given value.

### HasLaunchpadIds

`func (o *BitlinkBody) HasLaunchpadIds() bool`

HasLaunchpadIds returns a boolean if a field has been set.

### GetId

`func (o *BitlinkBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BitlinkBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BitlinkBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BitlinkBody) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


