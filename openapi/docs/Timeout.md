# Timeout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to [**[]FieldError**](FieldError.md) |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewTimeout

`func NewTimeout() *Timeout`

NewTimeout instantiates a new Timeout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeoutWithDefaults

`func NewTimeoutWithDefaults() *Timeout`

NewTimeoutWithDefaults instantiates a new Timeout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *Timeout) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Timeout) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Timeout) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Timeout) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrors

`func (o *Timeout) GetErrors() []FieldError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Timeout) GetErrorsOk() (*[]FieldError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Timeout) SetErrors(v []FieldError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Timeout) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetResource

`func (o *Timeout) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *Timeout) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *Timeout) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *Timeout) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetDescription

`func (o *Timeout) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Timeout) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Timeout) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Timeout) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


