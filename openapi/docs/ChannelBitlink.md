# ChannelBitlink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeywordLink** | Pointer to **string** |  | [optional] 
**CreatedTs** | **int32** |  | 
**BitlinkId** | **string** |  | 
**ChannelGuid** | **string** |  | 
**Title** | Pointer to **string** |  | [optional] 
**CampaignIds** | Pointer to **[]string** |  | [optional] 
**TotalClicks** | Pointer to **int32** |  | [optional] 
**AggregateLink** | Pointer to **string** |  | [optional] 
**CampaignGuid** | **string** |  | 
**Link** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**LongUrl** | Pointer to **string** |  | [optional] 
**Clicks** | Pointer to [**[]CampaignClickData**](CampaignClickData.md) |  | [optional] 

## Methods

### NewChannelBitlink

`func NewChannelBitlink(createdTs int32, bitlinkId string, channelGuid string, campaignGuid string, ) *ChannelBitlink`

NewChannelBitlink instantiates a new ChannelBitlink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelBitlinkWithDefaults

`func NewChannelBitlinkWithDefaults() *ChannelBitlink`

NewChannelBitlinkWithDefaults instantiates a new ChannelBitlink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeywordLink

`func (o *ChannelBitlink) GetKeywordLink() string`

GetKeywordLink returns the KeywordLink field if non-nil, zero value otherwise.

### GetKeywordLinkOk

`func (o *ChannelBitlink) GetKeywordLinkOk() (*string, bool)`

GetKeywordLinkOk returns a tuple with the KeywordLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywordLink

`func (o *ChannelBitlink) SetKeywordLink(v string)`

SetKeywordLink sets KeywordLink field to given value.

### HasKeywordLink

`func (o *ChannelBitlink) HasKeywordLink() bool`

HasKeywordLink returns a boolean if a field has been set.

### GetCreatedTs

`func (o *ChannelBitlink) GetCreatedTs() int32`

GetCreatedTs returns the CreatedTs field if non-nil, zero value otherwise.

### GetCreatedTsOk

`func (o *ChannelBitlink) GetCreatedTsOk() (*int32, bool)`

GetCreatedTsOk returns a tuple with the CreatedTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTs

`func (o *ChannelBitlink) SetCreatedTs(v int32)`

SetCreatedTs sets CreatedTs field to given value.


### GetBitlinkId

`func (o *ChannelBitlink) GetBitlinkId() string`

GetBitlinkId returns the BitlinkId field if non-nil, zero value otherwise.

### GetBitlinkIdOk

`func (o *ChannelBitlink) GetBitlinkIdOk() (*string, bool)`

GetBitlinkIdOk returns a tuple with the BitlinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitlinkId

`func (o *ChannelBitlink) SetBitlinkId(v string)`

SetBitlinkId sets BitlinkId field to given value.


### GetChannelGuid

`func (o *ChannelBitlink) GetChannelGuid() string`

GetChannelGuid returns the ChannelGuid field if non-nil, zero value otherwise.

### GetChannelGuidOk

`func (o *ChannelBitlink) GetChannelGuidOk() (*string, bool)`

GetChannelGuidOk returns a tuple with the ChannelGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelGuid

`func (o *ChannelBitlink) SetChannelGuid(v string)`

SetChannelGuid sets ChannelGuid field to given value.


### GetTitle

`func (o *ChannelBitlink) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChannelBitlink) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChannelBitlink) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ChannelBitlink) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCampaignIds

`func (o *ChannelBitlink) GetCampaignIds() []string`

GetCampaignIds returns the CampaignIds field if non-nil, zero value otherwise.

### GetCampaignIdsOk

`func (o *ChannelBitlink) GetCampaignIdsOk() (*[]string, bool)`

GetCampaignIdsOk returns a tuple with the CampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignIds

`func (o *ChannelBitlink) SetCampaignIds(v []string)`

SetCampaignIds sets CampaignIds field to given value.

### HasCampaignIds

`func (o *ChannelBitlink) HasCampaignIds() bool`

HasCampaignIds returns a boolean if a field has been set.

### GetTotalClicks

`func (o *ChannelBitlink) GetTotalClicks() int32`

GetTotalClicks returns the TotalClicks field if non-nil, zero value otherwise.

### GetTotalClicksOk

`func (o *ChannelBitlink) GetTotalClicksOk() (*int32, bool)`

GetTotalClicksOk returns a tuple with the TotalClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalClicks

`func (o *ChannelBitlink) SetTotalClicks(v int32)`

SetTotalClicks sets TotalClicks field to given value.

### HasTotalClicks

`func (o *ChannelBitlink) HasTotalClicks() bool`

HasTotalClicks returns a boolean if a field has been set.

### GetAggregateLink

`func (o *ChannelBitlink) GetAggregateLink() string`

GetAggregateLink returns the AggregateLink field if non-nil, zero value otherwise.

### GetAggregateLinkOk

`func (o *ChannelBitlink) GetAggregateLinkOk() (*string, bool)`

GetAggregateLinkOk returns a tuple with the AggregateLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateLink

`func (o *ChannelBitlink) SetAggregateLink(v string)`

SetAggregateLink sets AggregateLink field to given value.

### HasAggregateLink

`func (o *ChannelBitlink) HasAggregateLink() bool`

HasAggregateLink returns a boolean if a field has been set.

### GetCampaignGuid

`func (o *ChannelBitlink) GetCampaignGuid() string`

GetCampaignGuid returns the CampaignGuid field if non-nil, zero value otherwise.

### GetCampaignGuidOk

`func (o *ChannelBitlink) GetCampaignGuidOk() (*string, bool)`

GetCampaignGuidOk returns a tuple with the CampaignGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignGuid

`func (o *ChannelBitlink) SetCampaignGuid(v string)`

SetCampaignGuid sets CampaignGuid field to given value.


### GetLink

`func (o *ChannelBitlink) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *ChannelBitlink) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *ChannelBitlink) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *ChannelBitlink) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetClientId

`func (o *ChannelBitlink) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ChannelBitlink) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ChannelBitlink) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ChannelBitlink) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetLongUrl

`func (o *ChannelBitlink) GetLongUrl() string`

GetLongUrl returns the LongUrl field if non-nil, zero value otherwise.

### GetLongUrlOk

`func (o *ChannelBitlink) GetLongUrlOk() (*string, bool)`

GetLongUrlOk returns a tuple with the LongUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongUrl

`func (o *ChannelBitlink) SetLongUrl(v string)`

SetLongUrl sets LongUrl field to given value.

### HasLongUrl

`func (o *ChannelBitlink) HasLongUrl() bool`

HasLongUrl returns a boolean if a field has been set.

### GetClicks

`func (o *ChannelBitlink) GetClicks() []CampaignClickData`

GetClicks returns the Clicks field if non-nil, zero value otherwise.

### GetClicksOk

`func (o *ChannelBitlink) GetClicksOk() (*[]CampaignClickData, bool)`

GetClicksOk returns a tuple with the Clicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicks

`func (o *ChannelBitlink) SetClicks(v []CampaignClickData)`

SetClicks sets Clicks field to given value.

### HasClicks

`func (o *ChannelBitlink) HasClicks() bool`

HasClicks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


