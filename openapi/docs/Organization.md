# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**References** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** |  | 
**Bsds** | **[]string** |  | 
**Created** | **string** |  | 
**IsActive** | **bool** |  | 
**Modified** | **string** |  | 
**TierDisplayName** | **string** |  | 
**TierFamily** | **string** |  | 
**Tier** | **string** |  | 
**Role** | **string** |  | 
**Guid** | **string** |  | 

## Methods

### NewOrganization

`func NewOrganization(name string, bsds []string, created string, isActive bool, modified string, tierDisplayName string, tierFamily string, tier string, role string, guid string, ) *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferences

`func (o *Organization) GetReferences() map[string]string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *Organization) GetReferencesOk() (*map[string]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *Organization) SetReferences(v map[string]string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *Organization) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetName

`func (o *Organization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Organization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Organization) SetName(v string)`

SetName sets Name field to given value.


### GetBsds

`func (o *Organization) GetBsds() []string`

GetBsds returns the Bsds field if non-nil, zero value otherwise.

### GetBsdsOk

`func (o *Organization) GetBsdsOk() (*[]string, bool)`

GetBsdsOk returns a tuple with the Bsds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsds

`func (o *Organization) SetBsds(v []string)`

SetBsds sets Bsds field to given value.


### GetCreated

`func (o *Organization) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Organization) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Organization) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetIsActive

`func (o *Organization) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Organization) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Organization) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetModified

`func (o *Organization) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Organization) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Organization) SetModified(v string)`

SetModified sets Modified field to given value.


### GetTierDisplayName

`func (o *Organization) GetTierDisplayName() string`

GetTierDisplayName returns the TierDisplayName field if non-nil, zero value otherwise.

### GetTierDisplayNameOk

`func (o *Organization) GetTierDisplayNameOk() (*string, bool)`

GetTierDisplayNameOk returns a tuple with the TierDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierDisplayName

`func (o *Organization) SetTierDisplayName(v string)`

SetTierDisplayName sets TierDisplayName field to given value.


### GetTierFamily

`func (o *Organization) GetTierFamily() string`

GetTierFamily returns the TierFamily field if non-nil, zero value otherwise.

### GetTierFamilyOk

`func (o *Organization) GetTierFamilyOk() (*string, bool)`

GetTierFamilyOk returns a tuple with the TierFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierFamily

`func (o *Organization) SetTierFamily(v string)`

SetTierFamily sets TierFamily field to given value.


### GetTier

`func (o *Organization) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *Organization) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *Organization) SetTier(v string)`

SetTier sets Tier field to given value.


### GetRole

`func (o *Organization) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Organization) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Organization) SetRole(v string)`

SetRole sets Role field to given value.


### GetGuid

`func (o *Organization) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Organization) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Organization) SetGuid(v string)`

SetGuid sets Guid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


