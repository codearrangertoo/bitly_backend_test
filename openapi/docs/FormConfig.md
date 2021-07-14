# FormConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to [**[]FormSelectOptions**](FormSelectOptions.md) |  | [optional] 

## Methods

### NewFormConfig

`func NewFormConfig() *FormConfig`

NewFormConfig instantiates a new FormConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormConfigWithDefaults

`func NewFormConfigWithDefaults() *FormConfig`

NewFormConfigWithDefaults instantiates a new FormConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *FormConfig) GetOptions() []FormSelectOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FormConfig) GetOptionsOk() (*[]FormSelectOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FormConfig) SetOptions(v []FormSelectOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *FormConfig) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


