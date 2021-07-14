# InvitationsCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleName** | **string** |  | 
**Emails** | **[]string** |  | 
**Groups** | Pointer to [**[]InvitationGroup**](InvitationGroup.md) |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 

## Methods

### NewInvitationsCreate

`func NewInvitationsCreate(roleName string, emails []string, ) *InvitationsCreate`

NewInvitationsCreate instantiates a new InvitationsCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationsCreateWithDefaults

`func NewInvitationsCreateWithDefaults() *InvitationsCreate`

NewInvitationsCreateWithDefaults instantiates a new InvitationsCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleName

`func (o *InvitationsCreate) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *InvitationsCreate) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *InvitationsCreate) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetEmails

`func (o *InvitationsCreate) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *InvitationsCreate) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *InvitationsCreate) SetEmails(v []string)`

SetEmails sets Emails field to given value.


### GetGroups

`func (o *InvitationsCreate) GetGroups() []InvitationGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *InvitationsCreate) GetGroupsOk() (*[]InvitationGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *InvitationsCreate) SetGroups(v []InvitationGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *InvitationsCreate) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetCreated

`func (o *InvitationsCreate) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InvitationsCreate) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *InvitationsCreate) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *InvitationsCreate) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


