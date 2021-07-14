# BitlinkHistoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** | The backhalf of the underlying hash link | [optional] 
**GroupGuid** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**LongUrl** | Pointer to **string** |  | [optional] 
**DeactivatedAt** | Pointer to **string** |  | [optional] 
**Login** | Pointer to **string** |  | [optional] 
**TargetBitlinkId** | Pointer to **string** | The domain and backhalf of the underlying hash link | [optional] 
**Id** | Pointer to **string** | The domain and backhalf of the parent override | [optional] 

## Methods

### NewBitlinkHistoryAllOf

`func NewBitlinkHistoryAllOf() *BitlinkHistoryAllOf`

NewBitlinkHistoryAllOf instantiates a new BitlinkHistoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBitlinkHistoryAllOfWithDefaults

`func NewBitlinkHistoryAllOfWithDefaults() *BitlinkHistoryAllOf`

NewBitlinkHistoryAllOfWithDefaults instantiates a new BitlinkHistoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *BitlinkHistoryAllOf) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *BitlinkHistoryAllOf) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *BitlinkHistoryAllOf) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *BitlinkHistoryAllOf) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetGroupGuid

`func (o *BitlinkHistoryAllOf) GetGroupGuid() string`

GetGroupGuid returns the GroupGuid field if non-nil, zero value otherwise.

### GetGroupGuidOk

`func (o *BitlinkHistoryAllOf) GetGroupGuidOk() (*string, bool)`

GetGroupGuidOk returns a tuple with the GroupGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupGuid

`func (o *BitlinkHistoryAllOf) SetGroupGuid(v string)`

SetGroupGuid sets GroupGuid field to given value.

### HasGroupGuid

`func (o *BitlinkHistoryAllOf) HasGroupGuid() bool`

HasGroupGuid returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BitlinkHistoryAllOf) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BitlinkHistoryAllOf) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BitlinkHistoryAllOf) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BitlinkHistoryAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetIsActive

`func (o *BitlinkHistoryAllOf) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *BitlinkHistoryAllOf) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *BitlinkHistoryAllOf) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *BitlinkHistoryAllOf) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLongUrl

`func (o *BitlinkHistoryAllOf) GetLongUrl() string`

GetLongUrl returns the LongUrl field if non-nil, zero value otherwise.

### GetLongUrlOk

`func (o *BitlinkHistoryAllOf) GetLongUrlOk() (*string, bool)`

GetLongUrlOk returns a tuple with the LongUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongUrl

`func (o *BitlinkHistoryAllOf) SetLongUrl(v string)`

SetLongUrl sets LongUrl field to given value.

### HasLongUrl

`func (o *BitlinkHistoryAllOf) HasLongUrl() bool`

HasLongUrl returns a boolean if a field has been set.

### GetDeactivatedAt

`func (o *BitlinkHistoryAllOf) GetDeactivatedAt() string`

GetDeactivatedAt returns the DeactivatedAt field if non-nil, zero value otherwise.

### GetDeactivatedAtOk

`func (o *BitlinkHistoryAllOf) GetDeactivatedAtOk() (*string, bool)`

GetDeactivatedAtOk returns a tuple with the DeactivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivatedAt

`func (o *BitlinkHistoryAllOf) SetDeactivatedAt(v string)`

SetDeactivatedAt sets DeactivatedAt field to given value.

### HasDeactivatedAt

`func (o *BitlinkHistoryAllOf) HasDeactivatedAt() bool`

HasDeactivatedAt returns a boolean if a field has been set.

### GetLogin

`func (o *BitlinkHistoryAllOf) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *BitlinkHistoryAllOf) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *BitlinkHistoryAllOf) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *BitlinkHistoryAllOf) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetTargetBitlinkId

`func (o *BitlinkHistoryAllOf) GetTargetBitlinkId() string`

GetTargetBitlinkId returns the TargetBitlinkId field if non-nil, zero value otherwise.

### GetTargetBitlinkIdOk

`func (o *BitlinkHistoryAllOf) GetTargetBitlinkIdOk() (*string, bool)`

GetTargetBitlinkIdOk returns a tuple with the TargetBitlinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBitlinkId

`func (o *BitlinkHistoryAllOf) SetTargetBitlinkId(v string)`

SetTargetBitlinkId sets TargetBitlinkId field to given value.

### HasTargetBitlinkId

`func (o *BitlinkHistoryAllOf) HasTargetBitlinkId() bool`

HasTargetBitlinkId returns a boolean if a field has been set.

### GetId

`func (o *BitlinkHistoryAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BitlinkHistoryAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BitlinkHistoryAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BitlinkHistoryAllOf) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


