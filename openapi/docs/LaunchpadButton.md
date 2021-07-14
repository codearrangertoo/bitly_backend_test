# LaunchpadButton

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LaunchpadId** | Pointer to **string** |  | [optional] 
**Bitlink** | Pointer to **string** |  | [optional] 
**ButtonId** | Pointer to **string** |  | [optional] 
**Keyword** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**IsActive** | **bool** |  | 
**Domain** | Pointer to **string** |  | [optional] 
**SortOrder** | Pointer to **int32** |  | [optional] 
**LongUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewLaunchpadButton

`func NewLaunchpadButton(isActive bool, ) *LaunchpadButton`

NewLaunchpadButton instantiates a new LaunchpadButton object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLaunchpadButtonWithDefaults

`func NewLaunchpadButtonWithDefaults() *LaunchpadButton`

NewLaunchpadButtonWithDefaults instantiates a new LaunchpadButton object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLaunchpadId

`func (o *LaunchpadButton) GetLaunchpadId() string`

GetLaunchpadId returns the LaunchpadId field if non-nil, zero value otherwise.

### GetLaunchpadIdOk

`func (o *LaunchpadButton) GetLaunchpadIdOk() (*string, bool)`

GetLaunchpadIdOk returns a tuple with the LaunchpadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchpadId

`func (o *LaunchpadButton) SetLaunchpadId(v string)`

SetLaunchpadId sets LaunchpadId field to given value.

### HasLaunchpadId

`func (o *LaunchpadButton) HasLaunchpadId() bool`

HasLaunchpadId returns a boolean if a field has been set.

### GetBitlink

`func (o *LaunchpadButton) GetBitlink() string`

GetBitlink returns the Bitlink field if non-nil, zero value otherwise.

### GetBitlinkOk

`func (o *LaunchpadButton) GetBitlinkOk() (*string, bool)`

GetBitlinkOk returns a tuple with the Bitlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitlink

`func (o *LaunchpadButton) SetBitlink(v string)`

SetBitlink sets Bitlink field to given value.

### HasBitlink

`func (o *LaunchpadButton) HasBitlink() bool`

HasBitlink returns a boolean if a field has been set.

### GetButtonId

`func (o *LaunchpadButton) GetButtonId() string`

GetButtonId returns the ButtonId field if non-nil, zero value otherwise.

### GetButtonIdOk

`func (o *LaunchpadButton) GetButtonIdOk() (*string, bool)`

GetButtonIdOk returns a tuple with the ButtonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonId

`func (o *LaunchpadButton) SetButtonId(v string)`

SetButtonId sets ButtonId field to given value.

### HasButtonId

`func (o *LaunchpadButton) HasButtonId() bool`

HasButtonId returns a boolean if a field has been set.

### GetKeyword

`func (o *LaunchpadButton) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *LaunchpadButton) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *LaunchpadButton) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *LaunchpadButton) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### GetTitle

`func (o *LaunchpadButton) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LaunchpadButton) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LaunchpadButton) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LaunchpadButton) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetIsActive

`func (o *LaunchpadButton) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *LaunchpadButton) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *LaunchpadButton) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetDomain

`func (o *LaunchpadButton) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *LaunchpadButton) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *LaunchpadButton) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *LaunchpadButton) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetSortOrder

`func (o *LaunchpadButton) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *LaunchpadButton) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *LaunchpadButton) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *LaunchpadButton) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetLongUrl

`func (o *LaunchpadButton) GetLongUrl() string`

GetLongUrl returns the LongUrl field if non-nil, zero value otherwise.

### GetLongUrlOk

`func (o *LaunchpadButton) GetLongUrlOk() (*string, bool)`

GetLongUrlOk returns a tuple with the LongUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongUrl

`func (o *LaunchpadButton) SetLongUrl(v string)`

SetLongUrl sets LongUrl field to given value.

### HasLongUrl

`func (o *LaunchpadButton) HasLongUrl() bool`

HasLongUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


