# Conflict

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to [**[]FieldError**](FieldError.md) |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewConflict

`func NewConflict() *Conflict`

NewConflict instantiates a new Conflict object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConflictWithDefaults

`func NewConflictWithDefaults() *Conflict`

NewConflictWithDefaults instantiates a new Conflict object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *Conflict) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Conflict) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Conflict) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Conflict) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrors

`func (o *Conflict) GetErrors() []FieldError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Conflict) GetErrorsOk() (*[]FieldError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Conflict) SetErrors(v []FieldError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Conflict) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetResource

`func (o *Conflict) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *Conflict) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *Conflict) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *Conflict) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetDescription

`func (o *Conflict) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Conflict) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Conflict) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Conflict) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


