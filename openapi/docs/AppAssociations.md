# AppAssociations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AndroidInstallPreference** | Pointer to [**InstallPreference**](InstallPreference.md) |  | [optional] 
**IosInstallPreference** | Pointer to [**InstallPreference**](InstallPreference.md) |  | [optional] 
**AndroidApps** | Pointer to [**[]AppAssociationDetail**](AppAssociationDetail.md) |  | [optional] 
**CustomDomain** | Pointer to **string** |  | [optional] 
**IosApps** | Pointer to [**[]AppAssociationDetail**](AppAssociationDetail.md) |  | [optional] 

## Methods

### NewAppAssociations

`func NewAppAssociations() *AppAssociations`

NewAppAssociations instantiates a new AppAssociations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAssociationsWithDefaults

`func NewAppAssociationsWithDefaults() *AppAssociations`

NewAppAssociationsWithDefaults instantiates a new AppAssociations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAndroidInstallPreference

`func (o *AppAssociations) GetAndroidInstallPreference() InstallPreference`

GetAndroidInstallPreference returns the AndroidInstallPreference field if non-nil, zero value otherwise.

### GetAndroidInstallPreferenceOk

`func (o *AppAssociations) GetAndroidInstallPreferenceOk() (*InstallPreference, bool)`

GetAndroidInstallPreferenceOk returns a tuple with the AndroidInstallPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidInstallPreference

`func (o *AppAssociations) SetAndroidInstallPreference(v InstallPreference)`

SetAndroidInstallPreference sets AndroidInstallPreference field to given value.

### HasAndroidInstallPreference

`func (o *AppAssociations) HasAndroidInstallPreference() bool`

HasAndroidInstallPreference returns a boolean if a field has been set.

### GetIosInstallPreference

`func (o *AppAssociations) GetIosInstallPreference() InstallPreference`

GetIosInstallPreference returns the IosInstallPreference field if non-nil, zero value otherwise.

### GetIosInstallPreferenceOk

`func (o *AppAssociations) GetIosInstallPreferenceOk() (*InstallPreference, bool)`

GetIosInstallPreferenceOk returns a tuple with the IosInstallPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosInstallPreference

`func (o *AppAssociations) SetIosInstallPreference(v InstallPreference)`

SetIosInstallPreference sets IosInstallPreference field to given value.

### HasIosInstallPreference

`func (o *AppAssociations) HasIosInstallPreference() bool`

HasIosInstallPreference returns a boolean if a field has been set.

### GetAndroidApps

`func (o *AppAssociations) GetAndroidApps() []AppAssociationDetail`

GetAndroidApps returns the AndroidApps field if non-nil, zero value otherwise.

### GetAndroidAppsOk

`func (o *AppAssociations) GetAndroidAppsOk() (*[]AppAssociationDetail, bool)`

GetAndroidAppsOk returns a tuple with the AndroidApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidApps

`func (o *AppAssociations) SetAndroidApps(v []AppAssociationDetail)`

SetAndroidApps sets AndroidApps field to given value.

### HasAndroidApps

`func (o *AppAssociations) HasAndroidApps() bool`

HasAndroidApps returns a boolean if a field has been set.

### GetCustomDomain

`func (o *AppAssociations) GetCustomDomain() string`

GetCustomDomain returns the CustomDomain field if non-nil, zero value otherwise.

### GetCustomDomainOk

`func (o *AppAssociations) GetCustomDomainOk() (*string, bool)`

GetCustomDomainOk returns a tuple with the CustomDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomain

`func (o *AppAssociations) SetCustomDomain(v string)`

SetCustomDomain sets CustomDomain field to given value.

### HasCustomDomain

`func (o *AppAssociations) HasCustomDomain() bool`

HasCustomDomain returns a boolean if a field has been set.

### GetIosApps

`func (o *AppAssociations) GetIosApps() []AppAssociationDetail`

GetIosApps returns the IosApps field if non-nil, zero value otherwise.

### GetIosAppsOk

`func (o *AppAssociations) GetIosAppsOk() (*[]AppAssociationDetail, bool)`

GetIosAppsOk returns a tuple with the IosApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosApps

`func (o *AppAssociations) SetIosApps(v []AppAssociationDetail)`

SetIosApps sets IosApps field to given value.

### HasIosApps

`func (o *AppAssociations) HasIosApps() bool`

HasIosApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


