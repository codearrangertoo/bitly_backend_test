# CountItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Ts** | Pointer to **string** |  | [optional] 

## Methods

### NewCountItem

`func NewCountItem() *CountItem`

NewCountItem instantiates a new CountItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountItemWithDefaults

`func NewCountItemWithDefaults() *CountItem`

NewCountItemWithDefaults instantiates a new CountItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CountItem) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CountItem) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CountItem) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CountItem) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTs

`func (o *CountItem) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *CountItem) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *CountItem) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *CountItem) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


