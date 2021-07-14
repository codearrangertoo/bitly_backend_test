# BadGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to [**[]FieldError**](FieldError.md) |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBadGateway

`func NewBadGateway() *BadGateway`

NewBadGateway instantiates a new BadGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadGatewayWithDefaults

`func NewBadGatewayWithDefaults() *BadGateway`

NewBadGatewayWithDefaults instantiates a new BadGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *BadGateway) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BadGateway) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BadGateway) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BadGateway) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrors

`func (o *BadGateway) GetErrors() []FieldError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BadGateway) GetErrorsOk() (*[]FieldError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BadGateway) SetErrors(v []FieldError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BadGateway) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetResource

`func (o *BadGateway) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *BadGateway) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *BadGateway) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *BadGateway) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetDescription

`func (o *BadGateway) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BadGateway) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BadGateway) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BadGateway) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


