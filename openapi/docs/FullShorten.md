# FullShorten

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** |  | [optional] [default to "bit.ly"]
**Title** | Pointer to **string** |  | [optional] 
**GroupGuid** | Pointer to **string** | Always include a specific group and custom domain in your shorten calls. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Deeplinks** | Pointer to [**[]Deeplink**](Deeplink.md) |  | [optional] 
**LongUrl** | **string** |  | 

## Methods

### NewFullShorten

`func NewFullShorten(longUrl string, ) *FullShorten`

NewFullShorten instantiates a new FullShorten object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullShortenWithDefaults

`func NewFullShortenWithDefaults() *FullShorten`

NewFullShortenWithDefaults instantiates a new FullShorten object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *FullShorten) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *FullShorten) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *FullShorten) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *FullShorten) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetTitle

`func (o *FullShorten) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FullShorten) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FullShorten) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FullShorten) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetGroupGuid

`func (o *FullShorten) GetGroupGuid() string`

GetGroupGuid returns the GroupGuid field if non-nil, zero value otherwise.

### GetGroupGuidOk

`func (o *FullShorten) GetGroupGuidOk() (*string, bool)`

GetGroupGuidOk returns a tuple with the GroupGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupGuid

`func (o *FullShorten) SetGroupGuid(v string)`

SetGroupGuid sets GroupGuid field to given value.

### HasGroupGuid

`func (o *FullShorten) HasGroupGuid() bool`

HasGroupGuid returns a boolean if a field has been set.

### GetTags

`func (o *FullShorten) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FullShorten) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FullShorten) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FullShorten) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDeeplinks

`func (o *FullShorten) GetDeeplinks() []Deeplink`

GetDeeplinks returns the Deeplinks field if non-nil, zero value otherwise.

### GetDeeplinksOk

`func (o *FullShorten) GetDeeplinksOk() (*[]Deeplink, bool)`

GetDeeplinksOk returns a tuple with the Deeplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeeplinks

`func (o *FullShorten) SetDeeplinks(v []Deeplink)`

SetDeeplinks sets Deeplinks field to given value.

### HasDeeplinks

`func (o *FullShorten) HasDeeplinks() bool`

HasDeeplinks returns a boolean if a field has been set.

### GetLongUrl

`func (o *FullShorten) GetLongUrl() string`

GetLongUrl returns the LongUrl field if non-nil, zero value otherwise.

### GetLongUrlOk

`func (o *FullShorten) GetLongUrlOk() (*string, bool)`

GetLongUrlOk returns a tuple with the LongUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongUrl

`func (o *FullShorten) SetLongUrl(v string)`

SetLongUrl sets LongUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


