# CampaignChannelClicks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTs** | Pointer to **int32** |  | [optional] 
**ChannelGuid** | Pointer to **string** |  | [optional] 
**TotalClicks** | Pointer to **int32** |  | [optional] 
**ChannelName** | Pointer to **string** |  | [optional] 
**CampaignGuid** | Pointer to **string** |  | [optional] 
**ChannelBitlinks** | Pointer to [**ChannelBitlinks**](ChannelBitlinks.md) |  | [optional] 
**Clicks** | Pointer to [**[]CampaignClickData**](CampaignClickData.md) |  | [optional] 

## Methods

### NewCampaignChannelClicks

`func NewCampaignChannelClicks() *CampaignChannelClicks`

NewCampaignChannelClicks instantiates a new CampaignChannelClicks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignChannelClicksWithDefaults

`func NewCampaignChannelClicksWithDefaults() *CampaignChannelClicks`

NewCampaignChannelClicksWithDefaults instantiates a new CampaignChannelClicks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTs

`func (o *CampaignChannelClicks) GetCreatedTs() int32`

GetCreatedTs returns the CreatedTs field if non-nil, zero value otherwise.

### GetCreatedTsOk

`func (o *CampaignChannelClicks) GetCreatedTsOk() (*int32, bool)`

GetCreatedTsOk returns a tuple with the CreatedTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTs

`func (o *CampaignChannelClicks) SetCreatedTs(v int32)`

SetCreatedTs sets CreatedTs field to given value.

### HasCreatedTs

`func (o *CampaignChannelClicks) HasCreatedTs() bool`

HasCreatedTs returns a boolean if a field has been set.

### GetChannelGuid

`func (o *CampaignChannelClicks) GetChannelGuid() string`

GetChannelGuid returns the ChannelGuid field if non-nil, zero value otherwise.

### GetChannelGuidOk

`func (o *CampaignChannelClicks) GetChannelGuidOk() (*string, bool)`

GetChannelGuidOk returns a tuple with the ChannelGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelGuid

`func (o *CampaignChannelClicks) SetChannelGuid(v string)`

SetChannelGuid sets ChannelGuid field to given value.

### HasChannelGuid

`func (o *CampaignChannelClicks) HasChannelGuid() bool`

HasChannelGuid returns a boolean if a field has been set.

### GetTotalClicks

`func (o *CampaignChannelClicks) GetTotalClicks() int32`

GetTotalClicks returns the TotalClicks field if non-nil, zero value otherwise.

### GetTotalClicksOk

`func (o *CampaignChannelClicks) GetTotalClicksOk() (*int32, bool)`

GetTotalClicksOk returns a tuple with the TotalClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalClicks

`func (o *CampaignChannelClicks) SetTotalClicks(v int32)`

SetTotalClicks sets TotalClicks field to given value.

### HasTotalClicks

`func (o *CampaignChannelClicks) HasTotalClicks() bool`

HasTotalClicks returns a boolean if a field has been set.

### GetChannelName

`func (o *CampaignChannelClicks) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *CampaignChannelClicks) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *CampaignChannelClicks) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *CampaignChannelClicks) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### GetCampaignGuid

`func (o *CampaignChannelClicks) GetCampaignGuid() string`

GetCampaignGuid returns the CampaignGuid field if non-nil, zero value otherwise.

### GetCampaignGuidOk

`func (o *CampaignChannelClicks) GetCampaignGuidOk() (*string, bool)`

GetCampaignGuidOk returns a tuple with the CampaignGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignGuid

`func (o *CampaignChannelClicks) SetCampaignGuid(v string)`

SetCampaignGuid sets CampaignGuid field to given value.

### HasCampaignGuid

`func (o *CampaignChannelClicks) HasCampaignGuid() bool`

HasCampaignGuid returns a boolean if a field has been set.

### GetChannelBitlinks

`func (o *CampaignChannelClicks) GetChannelBitlinks() ChannelBitlinks`

GetChannelBitlinks returns the ChannelBitlinks field if non-nil, zero value otherwise.

### GetChannelBitlinksOk

`func (o *CampaignChannelClicks) GetChannelBitlinksOk() (*ChannelBitlinks, bool)`

GetChannelBitlinksOk returns a tuple with the ChannelBitlinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelBitlinks

`func (o *CampaignChannelClicks) SetChannelBitlinks(v ChannelBitlinks)`

SetChannelBitlinks sets ChannelBitlinks field to given value.

### HasChannelBitlinks

`func (o *CampaignChannelClicks) HasChannelBitlinks() bool`

HasChannelBitlinks returns a boolean if a field has been set.

### GetClicks

`func (o *CampaignChannelClicks) GetClicks() []CampaignClickData`

GetClicks returns the Clicks field if non-nil, zero value otherwise.

### GetClicksOk

`func (o *CampaignChannelClicks) GetClicksOk() (*[]CampaignClickData, bool)`

GetClicksOk returns a tuple with the Clicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicks

`func (o *CampaignChannelClicks) SetClicks(v []CampaignClickData)`

SetClicks sets Clicks field to given value.

### HasClicks

`func (o *CampaignChannelClicks) HasClicks() bool`

HasClicks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


