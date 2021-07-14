# UserPasswordChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalPassword** | **string** |  | 
**NewPassword** | **string** |  | 

## Methods

### NewUserPasswordChange

`func NewUserPasswordChange(originalPassword string, newPassword string, ) *UserPasswordChange`

NewUserPasswordChange instantiates a new UserPasswordChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPasswordChangeWithDefaults

`func NewUserPasswordChangeWithDefaults() *UserPasswordChange`

NewUserPasswordChangeWithDefaults instantiates a new UserPasswordChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalPassword

`func (o *UserPasswordChange) GetOriginalPassword() string`

GetOriginalPassword returns the OriginalPassword field if non-nil, zero value otherwise.

### GetOriginalPasswordOk

`func (o *UserPasswordChange) GetOriginalPasswordOk() (*string, bool)`

GetOriginalPasswordOk returns a tuple with the OriginalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPassword

`func (o *UserPasswordChange) SetOriginalPassword(v string)`

SetOriginalPassword sets OriginalPassword field to given value.


### GetNewPassword

`func (o *UserPasswordChange) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *UserPasswordChange) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *UserPasswordChange) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


