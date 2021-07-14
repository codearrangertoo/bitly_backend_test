# DeeplinkApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThirdPartyAppId** | **string** |  | 
**CreatedTs** | Pointer to **int32** |  | [optional] 
**Scheme** | Pointer to **string** |  | [optional] 
**InstallUrl** | **string** |  | 
**Name** | **string** |  | 
**LegacyAppId** | Pointer to **string** |  | [optional] 
**IconUrl** | Pointer to **string** |  | [optional] 
**OrganizationGuid** | Pointer to **string** |  | [optional] 
**ModifiedTs** | Pointer to **int32** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**AppleAppEntitlementId** | Pointer to **string** |  | [optional] 
**AndroidSha256** | Pointer to **[]string** |  | [optional] 
**Os** | **string** |  | 

## Methods

### NewDeeplinkApp

`func NewDeeplinkApp(thirdPartyAppId string, installUrl string, name string, os string, ) *DeeplinkApp`

NewDeeplinkApp instantiates a new DeeplinkApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeeplinkAppWithDefaults

`func NewDeeplinkAppWithDefaults() *DeeplinkApp`

NewDeeplinkAppWithDefaults instantiates a new DeeplinkApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThirdPartyAppId

`func (o *DeeplinkApp) GetThirdPartyAppId() string`

GetThirdPartyAppId returns the ThirdPartyAppId field if non-nil, zero value otherwise.

### GetThirdPartyAppIdOk

`func (o *DeeplinkApp) GetThirdPartyAppIdOk() (*string, bool)`

GetThirdPartyAppIdOk returns a tuple with the ThirdPartyAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyAppId

`func (o *DeeplinkApp) SetThirdPartyAppId(v string)`

SetThirdPartyAppId sets ThirdPartyAppId field to given value.


### GetCreatedTs

`func (o *DeeplinkApp) GetCreatedTs() int32`

GetCreatedTs returns the CreatedTs field if non-nil, zero value otherwise.

### GetCreatedTsOk

`func (o *DeeplinkApp) GetCreatedTsOk() (*int32, bool)`

GetCreatedTsOk returns a tuple with the CreatedTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTs

`func (o *DeeplinkApp) SetCreatedTs(v int32)`

SetCreatedTs sets CreatedTs field to given value.

### HasCreatedTs

`func (o *DeeplinkApp) HasCreatedTs() bool`

HasCreatedTs returns a boolean if a field has been set.

### GetScheme

`func (o *DeeplinkApp) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *DeeplinkApp) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *DeeplinkApp) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *DeeplinkApp) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetInstallUrl

`func (o *DeeplinkApp) GetInstallUrl() string`

GetInstallUrl returns the InstallUrl field if non-nil, zero value otherwise.

### GetInstallUrlOk

`func (o *DeeplinkApp) GetInstallUrlOk() (*string, bool)`

GetInstallUrlOk returns a tuple with the InstallUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallUrl

`func (o *DeeplinkApp) SetInstallUrl(v string)`

SetInstallUrl sets InstallUrl field to given value.


### GetName

`func (o *DeeplinkApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeeplinkApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeeplinkApp) SetName(v string)`

SetName sets Name field to given value.


### GetLegacyAppId

`func (o *DeeplinkApp) GetLegacyAppId() string`

GetLegacyAppId returns the LegacyAppId field if non-nil, zero value otherwise.

### GetLegacyAppIdOk

`func (o *DeeplinkApp) GetLegacyAppIdOk() (*string, bool)`

GetLegacyAppIdOk returns a tuple with the LegacyAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyAppId

`func (o *DeeplinkApp) SetLegacyAppId(v string)`

SetLegacyAppId sets LegacyAppId field to given value.

### HasLegacyAppId

`func (o *DeeplinkApp) HasLegacyAppId() bool`

HasLegacyAppId returns a boolean if a field has been set.

### GetIconUrl

`func (o *DeeplinkApp) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *DeeplinkApp) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *DeeplinkApp) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *DeeplinkApp) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetOrganizationGuid

`func (o *DeeplinkApp) GetOrganizationGuid() string`

GetOrganizationGuid returns the OrganizationGuid field if non-nil, zero value otherwise.

### GetOrganizationGuidOk

`func (o *DeeplinkApp) GetOrganizationGuidOk() (*string, bool)`

GetOrganizationGuidOk returns a tuple with the OrganizationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationGuid

`func (o *DeeplinkApp) SetOrganizationGuid(v string)`

SetOrganizationGuid sets OrganizationGuid field to given value.

### HasOrganizationGuid

`func (o *DeeplinkApp) HasOrganizationGuid() bool`

HasOrganizationGuid returns a boolean if a field has been set.

### GetModifiedTs

`func (o *DeeplinkApp) GetModifiedTs() int32`

GetModifiedTs returns the ModifiedTs field if non-nil, zero value otherwise.

### GetModifiedTsOk

`func (o *DeeplinkApp) GetModifiedTsOk() (*int32, bool)`

GetModifiedTsOk returns a tuple with the ModifiedTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTs

`func (o *DeeplinkApp) SetModifiedTs(v int32)`

SetModifiedTs sets ModifiedTs field to given value.

### HasModifiedTs

`func (o *DeeplinkApp) HasModifiedTs() bool`

HasModifiedTs returns a boolean if a field has been set.

### GetGuid

`func (o *DeeplinkApp) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *DeeplinkApp) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *DeeplinkApp) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *DeeplinkApp) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetAppleAppEntitlementId

`func (o *DeeplinkApp) GetAppleAppEntitlementId() string`

GetAppleAppEntitlementId returns the AppleAppEntitlementId field if non-nil, zero value otherwise.

### GetAppleAppEntitlementIdOk

`func (o *DeeplinkApp) GetAppleAppEntitlementIdOk() (*string, bool)`

GetAppleAppEntitlementIdOk returns a tuple with the AppleAppEntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleAppEntitlementId

`func (o *DeeplinkApp) SetAppleAppEntitlementId(v string)`

SetAppleAppEntitlementId sets AppleAppEntitlementId field to given value.

### HasAppleAppEntitlementId

`func (o *DeeplinkApp) HasAppleAppEntitlementId() bool`

HasAppleAppEntitlementId returns a boolean if a field has been set.

### GetAndroidSha256

`func (o *DeeplinkApp) GetAndroidSha256() []string`

GetAndroidSha256 returns the AndroidSha256 field if non-nil, zero value otherwise.

### GetAndroidSha256Ok

`func (o *DeeplinkApp) GetAndroidSha256Ok() (*[]string, bool)`

GetAndroidSha256Ok returns a tuple with the AndroidSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidSha256

`func (o *DeeplinkApp) SetAndroidSha256(v []string)`

SetAndroidSha256 sets AndroidSha256 field to given value.

### HasAndroidSha256

`func (o *DeeplinkApp) HasAndroidSha256() bool`

HasAndroidSha256 returns a boolean if a field has been set.

### GetOs

`func (o *DeeplinkApp) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *DeeplinkApp) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *DeeplinkApp) SetOs(v string)`

SetOs sets Os field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


