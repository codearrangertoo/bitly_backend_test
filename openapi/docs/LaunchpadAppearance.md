# LaunchpadAppearance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LaunchpadId** | Pointer to **string** |  | [optional] 
**AvatarPreference** | Pointer to **string** |  | [optional] 
**HideBitlyLogo** | **bool** |  | 
**DisplayName** | Pointer to **string** |  | [optional] 
**AvatarBackgroundColor** | Pointer to **string** |  | [optional] 
**LaunchpadTextColor** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ButtonAppearance** | Pointer to [**LaunchpadButtonAppearance**](LaunchpadButtonAppearance.md) |  | [optional] 
**AvatarImageUrl** | Pointer to **string** |  | [optional] 
**StylePreference** | Pointer to **string** |  | [optional] 
**Background** | Pointer to **string** |  | [optional] 
**ThemeId** | Pointer to **int32** |  | [optional] 
**AvatarTextColor** | Pointer to **string** |  | [optional] 
**Layout** | Pointer to **string** |  | [optional] 

## Methods

### NewLaunchpadAppearance

`func NewLaunchpadAppearance(hideBitlyLogo bool, ) *LaunchpadAppearance`

NewLaunchpadAppearance instantiates a new LaunchpadAppearance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLaunchpadAppearanceWithDefaults

`func NewLaunchpadAppearanceWithDefaults() *LaunchpadAppearance`

NewLaunchpadAppearanceWithDefaults instantiates a new LaunchpadAppearance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLaunchpadId

`func (o *LaunchpadAppearance) GetLaunchpadId() string`

GetLaunchpadId returns the LaunchpadId field if non-nil, zero value otherwise.

### GetLaunchpadIdOk

`func (o *LaunchpadAppearance) GetLaunchpadIdOk() (*string, bool)`

GetLaunchpadIdOk returns a tuple with the LaunchpadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchpadId

`func (o *LaunchpadAppearance) SetLaunchpadId(v string)`

SetLaunchpadId sets LaunchpadId field to given value.

### HasLaunchpadId

`func (o *LaunchpadAppearance) HasLaunchpadId() bool`

HasLaunchpadId returns a boolean if a field has been set.

### GetAvatarPreference

`func (o *LaunchpadAppearance) GetAvatarPreference() string`

GetAvatarPreference returns the AvatarPreference field if non-nil, zero value otherwise.

### GetAvatarPreferenceOk

`func (o *LaunchpadAppearance) GetAvatarPreferenceOk() (*string, bool)`

GetAvatarPreferenceOk returns a tuple with the AvatarPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarPreference

`func (o *LaunchpadAppearance) SetAvatarPreference(v string)`

SetAvatarPreference sets AvatarPreference field to given value.

### HasAvatarPreference

`func (o *LaunchpadAppearance) HasAvatarPreference() bool`

HasAvatarPreference returns a boolean if a field has been set.

### GetHideBitlyLogo

`func (o *LaunchpadAppearance) GetHideBitlyLogo() bool`

GetHideBitlyLogo returns the HideBitlyLogo field if non-nil, zero value otherwise.

### GetHideBitlyLogoOk

`func (o *LaunchpadAppearance) GetHideBitlyLogoOk() (*bool, bool)`

GetHideBitlyLogoOk returns a tuple with the HideBitlyLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideBitlyLogo

`func (o *LaunchpadAppearance) SetHideBitlyLogo(v bool)`

SetHideBitlyLogo sets HideBitlyLogo field to given value.


### GetDisplayName

`func (o *LaunchpadAppearance) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *LaunchpadAppearance) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *LaunchpadAppearance) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *LaunchpadAppearance) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAvatarBackgroundColor

`func (o *LaunchpadAppearance) GetAvatarBackgroundColor() string`

GetAvatarBackgroundColor returns the AvatarBackgroundColor field if non-nil, zero value otherwise.

### GetAvatarBackgroundColorOk

`func (o *LaunchpadAppearance) GetAvatarBackgroundColorOk() (*string, bool)`

GetAvatarBackgroundColorOk returns a tuple with the AvatarBackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarBackgroundColor

`func (o *LaunchpadAppearance) SetAvatarBackgroundColor(v string)`

SetAvatarBackgroundColor sets AvatarBackgroundColor field to given value.

### HasAvatarBackgroundColor

`func (o *LaunchpadAppearance) HasAvatarBackgroundColor() bool`

HasAvatarBackgroundColor returns a boolean if a field has been set.

### GetLaunchpadTextColor

`func (o *LaunchpadAppearance) GetLaunchpadTextColor() string`

GetLaunchpadTextColor returns the LaunchpadTextColor field if non-nil, zero value otherwise.

### GetLaunchpadTextColorOk

`func (o *LaunchpadAppearance) GetLaunchpadTextColorOk() (*string, bool)`

GetLaunchpadTextColorOk returns a tuple with the LaunchpadTextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchpadTextColor

`func (o *LaunchpadAppearance) SetLaunchpadTextColor(v string)`

SetLaunchpadTextColor sets LaunchpadTextColor field to given value.

### HasLaunchpadTextColor

`func (o *LaunchpadAppearance) HasLaunchpadTextColor() bool`

HasLaunchpadTextColor returns a boolean if a field has been set.

### GetDescription

`func (o *LaunchpadAppearance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LaunchpadAppearance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LaunchpadAppearance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LaunchpadAppearance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetButtonAppearance

`func (o *LaunchpadAppearance) GetButtonAppearance() LaunchpadButtonAppearance`

GetButtonAppearance returns the ButtonAppearance field if non-nil, zero value otherwise.

### GetButtonAppearanceOk

`func (o *LaunchpadAppearance) GetButtonAppearanceOk() (*LaunchpadButtonAppearance, bool)`

GetButtonAppearanceOk returns a tuple with the ButtonAppearance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonAppearance

`func (o *LaunchpadAppearance) SetButtonAppearance(v LaunchpadButtonAppearance)`

SetButtonAppearance sets ButtonAppearance field to given value.

### HasButtonAppearance

`func (o *LaunchpadAppearance) HasButtonAppearance() bool`

HasButtonAppearance returns a boolean if a field has been set.

### GetAvatarImageUrl

`func (o *LaunchpadAppearance) GetAvatarImageUrl() string`

GetAvatarImageUrl returns the AvatarImageUrl field if non-nil, zero value otherwise.

### GetAvatarImageUrlOk

`func (o *LaunchpadAppearance) GetAvatarImageUrlOk() (*string, bool)`

GetAvatarImageUrlOk returns a tuple with the AvatarImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarImageUrl

`func (o *LaunchpadAppearance) SetAvatarImageUrl(v string)`

SetAvatarImageUrl sets AvatarImageUrl field to given value.

### HasAvatarImageUrl

`func (o *LaunchpadAppearance) HasAvatarImageUrl() bool`

HasAvatarImageUrl returns a boolean if a field has been set.

### GetStylePreference

`func (o *LaunchpadAppearance) GetStylePreference() string`

GetStylePreference returns the StylePreference field if non-nil, zero value otherwise.

### GetStylePreferenceOk

`func (o *LaunchpadAppearance) GetStylePreferenceOk() (*string, bool)`

GetStylePreferenceOk returns a tuple with the StylePreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStylePreference

`func (o *LaunchpadAppearance) SetStylePreference(v string)`

SetStylePreference sets StylePreference field to given value.

### HasStylePreference

`func (o *LaunchpadAppearance) HasStylePreference() bool`

HasStylePreference returns a boolean if a field has been set.

### GetBackground

`func (o *LaunchpadAppearance) GetBackground() string`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *LaunchpadAppearance) GetBackgroundOk() (*string, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *LaunchpadAppearance) SetBackground(v string)`

SetBackground sets Background field to given value.

### HasBackground

`func (o *LaunchpadAppearance) HasBackground() bool`

HasBackground returns a boolean if a field has been set.

### GetThemeId

`func (o *LaunchpadAppearance) GetThemeId() int32`

GetThemeId returns the ThemeId field if non-nil, zero value otherwise.

### GetThemeIdOk

`func (o *LaunchpadAppearance) GetThemeIdOk() (*int32, bool)`

GetThemeIdOk returns a tuple with the ThemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeId

`func (o *LaunchpadAppearance) SetThemeId(v int32)`

SetThemeId sets ThemeId field to given value.

### HasThemeId

`func (o *LaunchpadAppearance) HasThemeId() bool`

HasThemeId returns a boolean if a field has been set.

### GetAvatarTextColor

`func (o *LaunchpadAppearance) GetAvatarTextColor() string`

GetAvatarTextColor returns the AvatarTextColor field if non-nil, zero value otherwise.

### GetAvatarTextColorOk

`func (o *LaunchpadAppearance) GetAvatarTextColorOk() (*string, bool)`

GetAvatarTextColorOk returns a tuple with the AvatarTextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarTextColor

`func (o *LaunchpadAppearance) SetAvatarTextColor(v string)`

SetAvatarTextColor sets AvatarTextColor field to given value.

### HasAvatarTextColor

`func (o *LaunchpadAppearance) HasAvatarTextColor() bool`

HasAvatarTextColor returns a boolean if a field has been set.

### GetLayout

`func (o *LaunchpadAppearance) GetLayout() string`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *LaunchpadAppearance) GetLayoutOk() (*string, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *LaunchpadAppearance) SetLayout(v string)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *LaunchpadAppearance) HasLayout() bool`

HasLayout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


