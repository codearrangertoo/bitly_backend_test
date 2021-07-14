# Tier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Family** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **int32** |  | [optional] 
**BillingPeriodType** | Pointer to **string** |  | [optional] 
**ConsumableFeatures** | Pointer to [**[]DefaultConsumableFeature**](DefaultConsumableFeature.md) |  | [optional] 
**SortOrder** | Pointer to **int32** |  | [optional] 
**IsSelfService** | Pointer to **bool** |  | [optional] 
**AccessFeatures** | Pointer to [**[]DefaultAccessFeature**](DefaultAccessFeature.md) |  | [optional] 
**DisplayIcon** | Pointer to **string** |  | [optional] 
**IsPaid** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewTier

`func NewTier() *Tier`

NewTier instantiates a new Tier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierWithDefaults

`func NewTierWithDefaults() *Tier`

NewTierWithDefaults instantiates a new Tier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *Tier) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Tier) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Tier) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Tier) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *Tier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tier) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Tier) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFamily

`func (o *Tier) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *Tier) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *Tier) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *Tier) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetPrice

`func (o *Tier) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Tier) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Tier) SetPrice(v int32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Tier) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetBillingPeriodType

`func (o *Tier) GetBillingPeriodType() string`

GetBillingPeriodType returns the BillingPeriodType field if non-nil, zero value otherwise.

### GetBillingPeriodTypeOk

`func (o *Tier) GetBillingPeriodTypeOk() (*string, bool)`

GetBillingPeriodTypeOk returns a tuple with the BillingPeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodType

`func (o *Tier) SetBillingPeriodType(v string)`

SetBillingPeriodType sets BillingPeriodType field to given value.

### HasBillingPeriodType

`func (o *Tier) HasBillingPeriodType() bool`

HasBillingPeriodType returns a boolean if a field has been set.

### GetConsumableFeatures

`func (o *Tier) GetConsumableFeatures() []DefaultConsumableFeature`

GetConsumableFeatures returns the ConsumableFeatures field if non-nil, zero value otherwise.

### GetConsumableFeaturesOk

`func (o *Tier) GetConsumableFeaturesOk() (*[]DefaultConsumableFeature, bool)`

GetConsumableFeaturesOk returns a tuple with the ConsumableFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumableFeatures

`func (o *Tier) SetConsumableFeatures(v []DefaultConsumableFeature)`

SetConsumableFeatures sets ConsumableFeatures field to given value.

### HasConsumableFeatures

`func (o *Tier) HasConsumableFeatures() bool`

HasConsumableFeatures returns a boolean if a field has been set.

### GetSortOrder

`func (o *Tier) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *Tier) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *Tier) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *Tier) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetIsSelfService

`func (o *Tier) GetIsSelfService() bool`

GetIsSelfService returns the IsSelfService field if non-nil, zero value otherwise.

### GetIsSelfServiceOk

`func (o *Tier) GetIsSelfServiceOk() (*bool, bool)`

GetIsSelfServiceOk returns a tuple with the IsSelfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelfService

`func (o *Tier) SetIsSelfService(v bool)`

SetIsSelfService sets IsSelfService field to given value.

### HasIsSelfService

`func (o *Tier) HasIsSelfService() bool`

HasIsSelfService returns a boolean if a field has been set.

### GetAccessFeatures

`func (o *Tier) GetAccessFeatures() []DefaultAccessFeature`

GetAccessFeatures returns the AccessFeatures field if non-nil, zero value otherwise.

### GetAccessFeaturesOk

`func (o *Tier) GetAccessFeaturesOk() (*[]DefaultAccessFeature, bool)`

GetAccessFeaturesOk returns a tuple with the AccessFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessFeatures

`func (o *Tier) SetAccessFeatures(v []DefaultAccessFeature)`

SetAccessFeatures sets AccessFeatures field to given value.

### HasAccessFeatures

`func (o *Tier) HasAccessFeatures() bool`

HasAccessFeatures returns a boolean if a field has been set.

### GetDisplayIcon

`func (o *Tier) GetDisplayIcon() string`

GetDisplayIcon returns the DisplayIcon field if non-nil, zero value otherwise.

### GetDisplayIconOk

`func (o *Tier) GetDisplayIconOk() (*string, bool)`

GetDisplayIconOk returns a tuple with the DisplayIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayIcon

`func (o *Tier) SetDisplayIcon(v string)`

SetDisplayIcon sets DisplayIcon field to given value.

### HasDisplayIcon

`func (o *Tier) HasDisplayIcon() bool`

HasDisplayIcon returns a boolean if a field has been set.

### GetIsPaid

`func (o *Tier) GetIsPaid() bool`

GetIsPaid returns the IsPaid field if non-nil, zero value otherwise.

### GetIsPaidOk

`func (o *Tier) GetIsPaidOk() (*bool, bool)`

GetIsPaidOk returns a tuple with the IsPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaid

`func (o *Tier) SetIsPaid(v bool)`

SetIsPaid sets IsPaid field to given value.

### HasIsPaid

`func (o *Tier) HasIsPaid() bool`

HasIsPaid returns a boolean if a field has been set.

### GetDescription

`func (o *Tier) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Tier) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Tier) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Tier) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


