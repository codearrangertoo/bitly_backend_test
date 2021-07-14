# FeatureUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumableFeatures** | [**[]ConsumableFeatureFull**](ConsumableFeatureFull.md) | an array of all of the organizations consumable features | 
**AccessFeatures** | [**[]AccessFeature**](AccessFeature.md) | an array of all of the organizations access features | 

## Methods

### NewFeatureUsage

`func NewFeatureUsage(consumableFeatures []ConsumableFeatureFull, accessFeatures []AccessFeature, ) *FeatureUsage`

NewFeatureUsage instantiates a new FeatureUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureUsageWithDefaults

`func NewFeatureUsageWithDefaults() *FeatureUsage`

NewFeatureUsageWithDefaults instantiates a new FeatureUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumableFeatures

`func (o *FeatureUsage) GetConsumableFeatures() []ConsumableFeatureFull`

GetConsumableFeatures returns the ConsumableFeatures field if non-nil, zero value otherwise.

### GetConsumableFeaturesOk

`func (o *FeatureUsage) GetConsumableFeaturesOk() (*[]ConsumableFeatureFull, bool)`

GetConsumableFeaturesOk returns a tuple with the ConsumableFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumableFeatures

`func (o *FeatureUsage) SetConsumableFeatures(v []ConsumableFeatureFull)`

SetConsumableFeatures sets ConsumableFeatures field to given value.


### GetAccessFeatures

`func (o *FeatureUsage) GetAccessFeatures() []AccessFeature`

GetAccessFeatures returns the AccessFeatures field if non-nil, zero value otherwise.

### GetAccessFeaturesOk

`func (o *FeatureUsage) GetAccessFeaturesOk() (*[]AccessFeature, bool)`

GetAccessFeaturesOk returns a tuple with the AccessFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessFeatures

`func (o *FeatureUsage) SetAccessFeatures(v []AccessFeature)`

SetAccessFeatures sets AccessFeatures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


