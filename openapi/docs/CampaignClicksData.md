# CampaignClicksData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTs** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**BrandGuid** | Pointer to **string** |  | [optional] 
**CampaignChannels** | Pointer to [**[]CampaignChannelClicks**](CampaignChannelClicks.md) |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**ModifiedTs** | Pointer to **int32** |  | [optional] 
**TotalClicks** | Pointer to **int32** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewCampaignClicksData

`func NewCampaignClicksData() *CampaignClicksData`

NewCampaignClicksData instantiates a new CampaignClicksData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignClicksDataWithDefaults

`func NewCampaignClicksDataWithDefaults() *CampaignClicksData`

NewCampaignClicksDataWithDefaults instantiates a new CampaignClicksData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTs

`func (o *CampaignClicksData) GetCreatedTs() int32`

GetCreatedTs returns the CreatedTs field if non-nil, zero value otherwise.

### GetCreatedTsOk

`func (o *CampaignClicksData) GetCreatedTsOk() (*int32, bool)`

GetCreatedTsOk returns a tuple with the CreatedTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTs

`func (o *CampaignClicksData) SetCreatedTs(v int32)`

SetCreatedTs sets CreatedTs field to given value.

### HasCreatedTs

`func (o *CampaignClicksData) HasCreatedTs() bool`

HasCreatedTs returns a boolean if a field has been set.

### GetDescription

`func (o *CampaignClicksData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignClicksData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignClicksData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CampaignClicksData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBrandGuid

`func (o *CampaignClicksData) GetBrandGuid() string`

GetBrandGuid returns the BrandGuid field if non-nil, zero value otherwise.

### GetBrandGuidOk

`func (o *CampaignClicksData) GetBrandGuidOk() (*string, bool)`

GetBrandGuidOk returns a tuple with the BrandGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandGuid

`func (o *CampaignClicksData) SetBrandGuid(v string)`

SetBrandGuid sets BrandGuid field to given value.

### HasBrandGuid

`func (o *CampaignClicksData) HasBrandGuid() bool`

HasBrandGuid returns a boolean if a field has been set.

### GetCampaignChannels

`func (o *CampaignClicksData) GetCampaignChannels() []CampaignChannelClicks`

GetCampaignChannels returns the CampaignChannels field if non-nil, zero value otherwise.

### GetCampaignChannelsOk

`func (o *CampaignClicksData) GetCampaignChannelsOk() (*[]CampaignChannelClicks, bool)`

GetCampaignChannelsOk returns a tuple with the CampaignChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignChannels

`func (o *CampaignClicksData) SetCampaignChannels(v []CampaignChannelClicks)`

SetCampaignChannels sets CampaignChannels field to given value.

### HasCampaignChannels

`func (o *CampaignClicksData) HasCampaignChannels() bool`

HasCampaignChannels returns a boolean if a field has been set.

### GetCreatedBy

`func (o *CampaignClicksData) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CampaignClicksData) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CampaignClicksData) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *CampaignClicksData) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetModifiedTs

`func (o *CampaignClicksData) GetModifiedTs() int32`

GetModifiedTs returns the ModifiedTs field if non-nil, zero value otherwise.

### GetModifiedTsOk

`func (o *CampaignClicksData) GetModifiedTsOk() (*int32, bool)`

GetModifiedTsOk returns a tuple with the ModifiedTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTs

`func (o *CampaignClicksData) SetModifiedTs(v int32)`

SetModifiedTs sets ModifiedTs field to given value.

### HasModifiedTs

`func (o *CampaignClicksData) HasModifiedTs() bool`

HasModifiedTs returns a boolean if a field has been set.

### GetTotalClicks

`func (o *CampaignClicksData) GetTotalClicks() int32`

GetTotalClicks returns the TotalClicks field if non-nil, zero value otherwise.

### GetTotalClicksOk

`func (o *CampaignClicksData) GetTotalClicksOk() (*int32, bool)`

GetTotalClicksOk returns a tuple with the TotalClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalClicks

`func (o *CampaignClicksData) SetTotalClicks(v int32)`

SetTotalClicks sets TotalClicks field to given value.

### HasTotalClicks

`func (o *CampaignClicksData) HasTotalClicks() bool`

HasTotalClicks returns a boolean if a field has been set.

### GetGuid

`func (o *CampaignClicksData) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *CampaignClicksData) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *CampaignClicksData) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *CampaignClicksData) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetName

`func (o *CampaignClicksData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignClicksData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignClicksData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CampaignClicksData) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


