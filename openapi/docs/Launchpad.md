# Launchpad

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LaunchpadId** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Keyword** | Pointer to **string** |  | [optional] 
**IsActive** | **bool** |  | 
**Buttons** | Pointer to [**[]LaunchpadButton**](LaunchpadButton.md) |  | [optional] 
**LaunchpadAppearance** | Pointer to [**LaunchpadAppearance**](LaunchpadAppearance.md) |  | [optional] 

## Methods

### NewLaunchpad

`func NewLaunchpad(isActive bool, ) *Launchpad`

NewLaunchpad instantiates a new Launchpad object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLaunchpadWithDefaults

`func NewLaunchpadWithDefaults() *Launchpad`

NewLaunchpadWithDefaults instantiates a new Launchpad object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLaunchpadId

`func (o *Launchpad) GetLaunchpadId() string`

GetLaunchpadId returns the LaunchpadId field if non-nil, zero value otherwise.

### GetLaunchpadIdOk

`func (o *Launchpad) GetLaunchpadIdOk() (*string, bool)`

GetLaunchpadIdOk returns a tuple with the LaunchpadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchpadId

`func (o *Launchpad) SetLaunchpadId(v string)`

SetLaunchpadId sets LaunchpadId field to given value.

### HasLaunchpadId

`func (o *Launchpad) HasLaunchpadId() bool`

HasLaunchpadId returns a boolean if a field has been set.

### GetDomain

`func (o *Launchpad) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Launchpad) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Launchpad) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Launchpad) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetKeyword

`func (o *Launchpad) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *Launchpad) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *Launchpad) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *Launchpad) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### GetIsActive

`func (o *Launchpad) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Launchpad) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Launchpad) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetButtons

`func (o *Launchpad) GetButtons() []LaunchpadButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *Launchpad) GetButtonsOk() (*[]LaunchpadButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *Launchpad) SetButtons(v []LaunchpadButton)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *Launchpad) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetLaunchpadAppearance

`func (o *Launchpad) GetLaunchpadAppearance() LaunchpadAppearance`

GetLaunchpadAppearance returns the LaunchpadAppearance field if non-nil, zero value otherwise.

### GetLaunchpadAppearanceOk

`func (o *Launchpad) GetLaunchpadAppearanceOk() (*LaunchpadAppearance, bool)`

GetLaunchpadAppearanceOk returns a tuple with the LaunchpadAppearance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchpadAppearance

`func (o *Launchpad) SetLaunchpadAppearance(v LaunchpadAppearance)`

SetLaunchpadAppearance sets LaunchpadAppearance field to given value.

### HasLaunchpadAppearance

`func (o *Launchpad) HasLaunchpadAppearance() bool`

HasLaunchpadAppearance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


