# TooManyRequests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to [**[]FieldError**](FieldError.md) |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewTooManyRequests

`func NewTooManyRequests() *TooManyRequests`

NewTooManyRequests instantiates a new TooManyRequests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTooManyRequestsWithDefaults

`func NewTooManyRequestsWithDefaults() *TooManyRequests`

NewTooManyRequestsWithDefaults instantiates a new TooManyRequests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *TooManyRequests) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TooManyRequests) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TooManyRequests) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TooManyRequests) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrors

`func (o *TooManyRequests) GetErrors() []FieldError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *TooManyRequests) GetErrorsOk() (*[]FieldError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *TooManyRequests) SetErrors(v []FieldError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *TooManyRequests) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetResource

`func (o *TooManyRequests) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *TooManyRequests) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *TooManyRequests) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *TooManyRequests) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetDescription

`func (o *TooManyRequests) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TooManyRequests) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TooManyRequests) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TooManyRequests) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


