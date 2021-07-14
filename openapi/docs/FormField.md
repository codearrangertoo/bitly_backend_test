# FormField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**FormConfig**](FormConfig.md) |  | [optional] 
**Response** | Pointer to **string** |  | [optional] 

## Methods

### NewFormField

`func NewFormField() *FormField`

NewFormField instantiates a new FormField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldWithDefaults

`func NewFormFieldWithDefaults() *FormField`

NewFormFieldWithDefaults instantiates a new FormField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldType

`func (o *FormField) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *FormField) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *FormField) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *FormField) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### GetName

`func (o *FormField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormField) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FormField) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequired

`func (o *FormField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FormField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FormField) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FormField) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetLabel

`func (o *FormField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormField) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FormField) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetConfig

`func (o *FormField) GetConfig() FormConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *FormField) GetConfigOk() (*FormConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *FormField) SetConfig(v FormConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *FormField) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetResponse

`func (o *FormField) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *FormField) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *FormField) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *FormField) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


