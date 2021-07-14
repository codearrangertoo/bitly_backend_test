# UserPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Preferences** | Pointer to [**[]UserPreference**](UserPreference.md) |  | [optional] 

## Methods

### NewUserPreferences

`func NewUserPreferences() *UserPreferences`

NewUserPreferences instantiates a new UserPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPreferencesWithDefaults

`func NewUserPreferencesWithDefaults() *UserPreferences`

NewUserPreferencesWithDefaults instantiates a new UserPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreferences

`func (o *UserPreferences) GetPreferences() []UserPreference`

GetPreferences returns the Preferences field if non-nil, zero value otherwise.

### GetPreferencesOk

`func (o *UserPreferences) GetPreferencesOk() (*[]UserPreference, bool)`

GetPreferencesOk returns a tuple with the Preferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferences

`func (o *UserPreferences) SetPreferences(v []UserPreference)`

SetPreferences sets Preferences field to given value.

### HasPreferences

`func (o *UserPreferences) HasPreferences() bool`

HasPreferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


