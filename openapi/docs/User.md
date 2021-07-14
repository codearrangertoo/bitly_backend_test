# User

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

## Methods

### NewUser

`func NewUser(name string, created string, isActive bool, modified string, isSsoUser bool, is2faEnabled bool, login string, emails []Email, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultGroupGuid

`func (o *User) GetDefaultGroupGuid() string`

GetDefaultGroupGuid returns the DefaultGroupGuid field if non-nil, zero value otherwise.

### GetDefaultGroupGuidOk

`func (o *User) GetDefaultGroupGuidOk() (*string, bool)`

GetDefaultGroupGuidOk returns a tuple with the DefaultGroupGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroupGuid

`func (o *User) SetDefaultGroupGuid(v string)`

SetDefaultGroupGuid sets DefaultGroupGuid field to given value.

### HasDefaultGroupGuid

`func (o *User) HasDefaultGroupGuid() bool`

HasDefaultGroupGuid returns a boolean if a field has been set.

### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v string)`

SetName sets Name field to given value.


### GetCreated

`func (o *User) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *User) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *User) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetIsActive

`func (o *User) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *User) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *User) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetModified

`func (o *User) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *User) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *User) SetModified(v string)`

SetModified sets Modified field to given value.


### GetIsSsoUser

`func (o *User) GetIsSsoUser() bool`

GetIsSsoUser returns the IsSsoUser field if non-nil, zero value otherwise.

### GetIsSsoUserOk

`func (o *User) GetIsSsoUserOk() (*bool, bool)`

GetIsSsoUserOk returns a tuple with the IsSsoUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSsoUser

`func (o *User) SetIsSsoUser(v bool)`

SetIsSsoUser sets IsSsoUser field to given value.


### GetIs2faEnabled

`func (o *User) GetIs2faEnabled() bool`

GetIs2faEnabled returns the Is2faEnabled field if non-nil, zero value otherwise.

### GetIs2faEnabledOk

`func (o *User) GetIs2faEnabledOk() (*bool, bool)`

GetIs2faEnabledOk returns a tuple with the Is2faEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIs2faEnabled

`func (o *User) SetIs2faEnabled(v bool)`

SetIs2faEnabled sets Is2faEnabled field to given value.


### GetLogin

`func (o *User) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *User) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *User) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetEmails

`func (o *User) GetEmails() []Email`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *User) GetEmailsOk() (*[]Email, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *User) SetEmails(v []Email)`

SetEmails sets Emails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


