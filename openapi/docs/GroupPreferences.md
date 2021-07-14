# GroupPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupGuid** | Pointer to **string** |  | [optional] 
**DomainPreference** | Pointer to **string** |  | [optional] 

## Methods

### NewGroupPreferences

`func NewGroupPreferences() *GroupPreferences`

NewGroupPreferences instantiates a new GroupPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupPreferencesWithDefaults

`func NewGroupPreferencesWithDefaults() *GroupPreferences`

NewGroupPreferencesWithDefaults instantiates a new GroupPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupGuid

`func (o *GroupPreferences) GetGroupGuid() string`

GetGroupGuid returns the GroupGuid field if non-nil, zero value otherwise.

### GetGroupGuidOk

`func (o *GroupPreferences) GetGroupGuidOk() (*string, bool)`

GetGroupGuidOk returns a tuple with the GroupGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupGuid

`func (o *GroupPreferences) SetGroupGuid(v string)`

SetGroupGuid sets GroupGuid field to given value.

### HasGroupGuid

`func (o *GroupPreferences) HasGroupGuid() bool`

HasGroupGuid returns a boolean if a field has been set.

### GetDomainPreference

`func (o *GroupPreferences) GetDomainPreference() string`

GetDomainPreference returns the DomainPreference field if non-nil, zero value otherwise.

### GetDomainPreferenceOk

`func (o *GroupPreferences) GetDomainPreferenceOk() (*string, bool)`

GetDomainPreferenceOk returns a tuple with the DomainPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainPreference

`func (o *GroupPreferences) SetDomainPreference(v string)`

SetDomainPreference sets DomainPreference field to given value.

### HasDomainPreference

`func (o *GroupPreferences) HasDomainPreference() bool`

HasDomainPreference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


