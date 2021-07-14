# Bitlinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Links** | Pointer to [**[]BitlinkBody**](BitlinkBody.md) |  | [optional] 

## Methods

### NewBitlinks

`func NewBitlinks() *Bitlinks`

NewBitlinks instantiates a new Bitlinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBitlinksWithDefaults

`func NewBitlinksWithDefaults() *Bitlinks`

NewBitlinksWithDefaults instantiates a new Bitlinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *Bitlinks) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *Bitlinks) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *Bitlinks) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *Bitlinks) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetLinks

`func (o *Bitlinks) GetLinks() []BitlinkBody`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Bitlinks) GetLinksOk() (*[]BitlinkBody, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Bitlinks) SetLinks(v []BitlinkBody)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Bitlinks) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


