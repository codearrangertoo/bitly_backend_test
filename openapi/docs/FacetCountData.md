# FacetCountData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** |  | [optional] 
**Ts** | Pointer to **int32** |  | [optional] 
**Items** | Pointer to [**[]FacetCountItem**](FacetCountItem.md) |  | [optional] 

## Methods

### NewFacetCountData

`func NewFacetCountData() *FacetCountData`

NewFacetCountData instantiates a new FacetCountData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacetCountDataWithDefaults

`func NewFacetCountDataWithDefaults() *FacetCountData`

NewFacetCountDataWithDefaults instantiates a new FacetCountData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *FacetCountData) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *FacetCountData) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *FacetCountData) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *FacetCountData) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetTs

`func (o *FacetCountData) GetTs() int32`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *FacetCountData) GetTsOk() (*int32, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *FacetCountData) SetTs(v int32)`

SetTs sets Ts field to given value.

### HasTs

`func (o *FacetCountData) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetItems

`func (o *FacetCountData) GetItems() []FacetCountItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FacetCountData) GetItemsOk() (*[]FacetCountItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FacetCountData) SetItems(v []FacetCountItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *FacetCountData) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


