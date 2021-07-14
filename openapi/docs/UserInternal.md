# UserInternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultGroupGuid** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Created** | **string** |  | 
**IsActive** | **bool** |  | 
**Modified** | **string** |  | 
**IsSsoUser** | **bool** |  | 
**Is2faEnabled** | **bool** |  | 
**Login** | **string** |  | 
**Emails** | [**[]Email**](Email.md) |  | 
**RoleName** | Pointer to **string** |  | [optional] 

## Methods

### NewUserInternal

`func NewUserInternal(name string, created string, isActive bool, modified string, isSsoUser bool, is2faEnabled bool, login string, emails []Email, ) *UserInternal`

NewUserInternal instantiates a new UserInternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInternalWithDefaults

`func NewUserInternalWithDefaults() *UserInternal`

NewUserInternalWithDefaults instantiates a new UserInternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultGroupGuid

`func (o *UserInternal) GetDefaultGroupGuid() string`

GetDefaultGroupGuid returns the DefaultGroupGuid field if non-nil, zero value otherwise.

### GetDefaultGroupGuidOk

`func (o *UserInternal) GetDefaultGroupGuidOk() (*string, bool)`

GetDefaultGroupGuidOk returns a tuple with the DefaultGroupGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroupGuid

`func (o *UserInternal) SetDefaultGroupGuid(v string)`

SetDefaultGroupGuid sets DefaultGroupGuid field to given value.

### HasDefaultGroupGuid

`func (o *UserInternal) HasDefaultGroupGuid() bool`

HasDefaultGroupGuid returns a boolean if a field has been set.

### GetName

`func (o *UserInternal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserInternal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserInternal) SetName(v string)`

SetName sets Name field to given value.


### GetCreated

`func (o *UserInternal) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UserInternal) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UserInternal) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetIsActive

`func (o *UserInternal) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UserInternal) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UserInternal) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetModified

`func (o *UserInternal) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *UserInternal) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *UserInternal) SetModified(v string)`

SetModified sets Modified field to given value.


### GetIsSsoUser

`func (o *UserInternal) GetIsSsoUser() bool`

GetIsSsoUser returns the IsSsoUser field if non-nil, zero value otherwise.

### GetIsSsoUserOk

`func (o *UserInternal) GetIsSsoUserOk() (*bool, bool)`

GetIsSsoUserOk returns a tuple with the IsSsoUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSsoUser

`func (o *UserInternal) SetIsSsoUser(v bool)`

SetIsSsoUser sets IsSsoUser field to given value.


### GetIs2faEnabled

`func (o *UserInternal) GetIs2faEnabled() bool`

GetIs2faEnabled returns the Is2faEnabled field if non-nil, zero value otherwise.

### GetIs2faEnabledOk

`func (o *UserInternal) GetIs2faEnabledOk() (*bool, bool)`

GetIs2faEnabledOk returns a tuple with the Is2faEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIs2faEnabled

`func (o *UserInternal) SetIs2faEnabled(v bool)`

SetIs2faEnabled sets Is2faEnabled field to given value.


### GetLogin

`func (o *UserInternal) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *UserInternal) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *UserInternal) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetEmails

`func (o *UserInternal) GetEmails() []Email`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *UserInternal) GetEmailsOk() (*[]Email, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *UserInternal) SetEmails(v []Email)`

SetEmails sets Emails field to given value.


### GetRoleName

`func (o *UserInternal) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *UserInternal) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *UserInternal) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *UserInternal) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


