# DeeplinkRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bitlink** | Pointer to **string** |  | [optional] 
**InstallUrl** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** | ISO timestamp | [optional] 
**AppUriPath** | Pointer to **string** |  | [optional] 
**Modified** | Pointer to **string** | ISO timestamp | [optional] 
**InstallType** | Pointer to **string** |  | [optional] 
**AppGuid** | Pointer to **string** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**BrandGuid** | Pointer to **string** |  | [optional] 

## Methods

### NewDeeplinkRule

`func NewDeeplinkRule() *DeeplinkRule`

NewDeeplinkRule instantiates a new DeeplinkRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeeplinkRuleWithDefaults

`func NewDeeplinkRuleWithDefaults() *DeeplinkRule`

NewDeeplinkRuleWithDefaults instantiates a new DeeplinkRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBitlink

`func (o *DeeplinkRule) GetBitlink() string`

GetBitlink returns the Bitlink field if non-nil, zero value otherwise.

### GetBitlinkOk

`func (o *DeeplinkRule) GetBitlinkOk() (*string, bool)`

GetBitlinkOk returns a tuple with the Bitlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitlink

`func (o *DeeplinkRule) SetBitlink(v string)`

SetBitlink sets Bitlink field to given value.

### HasBitlink

`func (o *DeeplinkRule) HasBitlink() bool`

HasBitlink returns a boolean if a field has been set.

### GetInstallUrl

`func (o *DeeplinkRule) GetInstallUrl() string`

GetInstallUrl returns the InstallUrl field if non-nil, zero value otherwise.

### GetInstallUrlOk

`func (o *DeeplinkRule) GetInstallUrlOk() (*string, bool)`

GetInstallUrlOk returns a tuple with the InstallUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallUrl

`func (o *DeeplinkRule) SetInstallUrl(v string)`

SetInstallUrl sets InstallUrl field to given value.

### HasInstallUrl

`func (o *DeeplinkRule) HasInstallUrl() bool`

HasInstallUrl returns a boolean if a field has been set.

### GetCreated

`func (o *DeeplinkRule) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeeplinkRule) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeeplinkRule) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DeeplinkRule) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetAppUriPath

`func (o *DeeplinkRule) GetAppUriPath() string`

GetAppUriPath returns the AppUriPath field if non-nil, zero value otherwise.

### GetAppUriPathOk

`func (o *DeeplinkRule) GetAppUriPathOk() (*string, bool)`

GetAppUriPathOk returns a tuple with the AppUriPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUriPath

`func (o *DeeplinkRule) SetAppUriPath(v string)`

SetAppUriPath sets AppUriPath field to given value.

### HasAppUriPath

`func (o *DeeplinkRule) HasAppUriPath() bool`

HasAppUriPath returns a boolean if a field has been set.

### GetModified

`func (o *DeeplinkRule) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *DeeplinkRule) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *DeeplinkRule) SetModified(v string)`

SetModified sets Modified field to given value.

### HasModified

`func (o *DeeplinkRule) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetInstallType

`func (o *DeeplinkRule) GetInstallType() string`

GetInstallType returns the InstallType field if non-nil, zero value otherwise.

### GetInstallTypeOk

`func (o *DeeplinkRule) GetInstallTypeOk() (*string, bool)`

GetInstallTypeOk returns a tuple with the InstallType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallType

`func (o *DeeplinkRule) SetInstallType(v string)`

SetInstallType sets InstallType field to given value.

### HasInstallType

`func (o *DeeplinkRule) HasInstallType() bool`

HasInstallType returns a boolean if a field has been set.

### GetAppGuid

`func (o *DeeplinkRule) GetAppGuid() string`

GetAppGuid returns the AppGuid field if non-nil, zero value otherwise.

### GetAppGuidOk

`func (o *DeeplinkRule) GetAppGuidOk() (*string, bool)`

GetAppGuidOk returns a tuple with the AppGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppGuid

`func (o *DeeplinkRule) SetAppGuid(v string)`

SetAppGuid sets AppGuid field to given value.

### HasAppGuid

`func (o *DeeplinkRule) HasAppGuid() bool`

HasAppGuid returns a boolean if a field has been set.

### GetGuid

`func (o *DeeplinkRule) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *DeeplinkRule) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *DeeplinkRule) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *DeeplinkRule) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetOs

`func (o *DeeplinkRule) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *DeeplinkRule) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *DeeplinkRule) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *DeeplinkRule) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetBrandGuid

`func (o *DeeplinkRule) GetBrandGuid() string`

GetBrandGuid returns the BrandGuid field if non-nil, zero value otherwise.

### GetBrandGuidOk

`func (o *DeeplinkRule) GetBrandGuidOk() (*string, bool)`

GetBrandGuidOk returns a tuple with the BrandGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandGuid

`func (o *DeeplinkRule) SetBrandGuid(v string)`

SetBrandGuid sets BrandGuid field to given value.

### HasBrandGuid

`func (o *DeeplinkRule) HasBrandGuid() bool`

HasBrandGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


