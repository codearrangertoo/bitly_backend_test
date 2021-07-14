# InvitationsError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Emails** | Pointer to **[]string** |  | [optional] 

## Methods

### NewInvitationsError

`func NewInvitationsError() *InvitationsError`

NewInvitationsError instantiates a new InvitationsError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationsErrorWithDefaults

`func NewInvitationsErrorWithDefaults() *InvitationsError`

NewInvitationsErrorWithDefaults instantiates a new InvitationsError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *InvitationsError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InvitationsError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InvitationsError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InvitationsError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetEmails

`func (o *InvitationsError) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *InvitationsError) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *InvitationsError) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *InvitationsError) HasEmails() bool`

HasEmails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


