# Shorten

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupGuid** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] [default to "bit.ly"]
**LongUrl** | **string** |  | 

## Methods

### NewShorten

`func NewShorten(longUrl string, ) *Shorten`

NewShorten instantiates a new Shorten object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShortenWithDefaults

`func NewShortenWithDefaults() *Shorten`

NewShortenWithDefaults instantiates a new Shorten object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupGuid

`func (o *Shorten) GetGroupGuid() string`

GetGroupGuid returns the GroupGuid field if non-nil, zero value otherwise.

### GetGroupGuidOk

`func (o *Shorten) GetGroupGuidOk() (*string, bool)`

GetGroupGuidOk returns a tuple with the GroupGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupGuid

`func (o *Shorten) SetGroupGuid(v string)`

SetGroupGuid sets GroupGuid field to given value.

### HasGroupGuid

`func (o *Shorten) HasGroupGuid() bool`

HasGroupGuid returns a boolean if a field has been set.

### GetDomain

`func (o *Shorten) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Shorten) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Shorten) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Shorten) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetLongUrl

`func (o *Shorten) GetLongUrl() string`

GetLongUrl returns the LongUrl field if non-nil, zero value otherwise.

### GetLongUrlOk

`func (o *Shorten) GetLongUrlOk() (*string, bool)`

GetLongUrlOk returns a tuple with the LongUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongUrl

`func (o *Shorten) SetLongUrl(v string)`

SetLongUrl sets LongUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


