# UpdateLaunchpad

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LaunchpadId** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**IsActive** | **bool** |  | 
**LaunchpadAppearance** | Pointer to [**LaunchpadAppearance**](LaunchpadAppearance.md) |  | [optional] 
**Keyword** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateLaunchpad

`func NewUpdateLaunchpad(isActive bool, ) *UpdateLaunchpad`

NewUpdateLaunchpad instantiates a new UpdateLaunchpad object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLaunchpadWithDefaults

`func NewUpdateLaunchpadWithDefaults() *UpdateLaunchpad`

NewUpdateLaunchpadWithDefaults instantiates a new UpdateLaunchpad object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLaunchpadId

`func (o *UpdateLaunchpad) GetLaunchpadId() string`

GetLaunchpadId returns the LaunchpadId field if non-nil, zero value otherwise.

### GetLaunchpadIdOk

`func (o *UpdateLaunchpad) GetLaunchpadIdOk() (*string, bool)`

GetLaunchpadIdOk returns a tuple with the LaunchpadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchpadId

`func (o *UpdateLaunchpad) SetLaunchpadId(v string)`

SetLaunchpadId sets LaunchpadId field to given value.

### HasLaunchpadId

`func (o *UpdateLaunchpad) HasLaunchpadId() bool`

HasLaunchpadId returns a boolean if a field has been set.

### GetDomain

`func (o *UpdateLaunchpad) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UpdateLaunchpad) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UpdateLaunchpad) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *UpdateLaunchpad) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetIsActive

`func (o *UpdateLaunchpad) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateLaunchpad) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateLaunchpad) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetLaunchpadAppearance

`func (o *UpdateLaunchpad) GetLaunchpadAppearance() LaunchpadAppearance`

GetLaunchpadAppearance returns the LaunchpadAppearance field if non-nil, zero value otherwise.

### GetLaunchpadAppearanceOk

`func (o *UpdateLaunchpad) GetLaunchpadAppearanceOk() (*LaunchpadAppearance, bool)`

GetLaunchpadAppearanceOk returns a tuple with the LaunchpadAppearance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchpadAppearance

`func (o *UpdateLaunchpad) SetLaunchpadAppearance(v LaunchpadAppearance)`

SetLaunchpadAppearance sets LaunchpadAppearance field to given value.

### HasLaunchpadAppearance

`func (o *UpdateLaunchpad) HasLaunchpadAppearance() bool`

HasLaunchpadAppearance returns a boolean if a field has been set.

### GetKeyword

`func (o *UpdateLaunchpad) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *UpdateLaunchpad) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *UpdateLaunchpad) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *UpdateLaunchpad) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


