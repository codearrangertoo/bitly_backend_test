# EmailBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**IsPrimary** | Pointer to **bool** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 

## Methods

### NewEmailBody

`func NewEmailBody() *EmailBody`

NewEmailBody instantiates a new EmailBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailBodyWithDefaults

`func NewEmailBodyWithDefaults() *EmailBody`

NewEmailBodyWithDefaults instantiates a new EmailBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *EmailBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailBody) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmailBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsPrimary

`func (o *EmailBody) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *EmailBody) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *EmailBody) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *EmailBody) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetEmail

`func (o *EmailBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailBody) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EmailBody) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


