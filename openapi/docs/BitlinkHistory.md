# BitlinkHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**References** | Pointer to **map[string]string** |  | [optional] 
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

### NewBitlinkHistory

`func NewBitlinkHistory() *BitlinkHistory`

NewBitlinkHistory instantiates a new BitlinkHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBitlinkHistoryWithDefaults

`func NewBitlinkHistoryWithDefaults() *BitlinkHistory`

NewBitlinkHistoryWithDefaults instantiates a new BitlinkHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferences

`func (o *BitlinkHistory) GetReferences() map[string]string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *BitlinkHistory) GetReferencesOk() (*map[string]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *BitlinkHistory) SetReferences(v map[string]string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *BitlinkHistory) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetHash

`func (o *BitlinkHistory) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *BitlinkHistory) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *BitlinkHistory) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *BitlinkHistory) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetGroupGuid

`func (o *BitlinkHistory) GetGroupGuid() string`

GetGroupGuid returns the GroupGuid field if non-nil, zero value otherwise.

### GetGroupGuidOk

`func (o *BitlinkHistory) GetGroupGuidOk() (*string, bool)`

GetGroupGuidOk returns a tuple with the GroupGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupGuid

`func (o *BitlinkHistory) SetGroupGuid(v string)`

SetGroupGuid sets GroupGuid field to given value.

### HasGroupGuid

`func (o *BitlinkHistory) HasGroupGuid() bool`

HasGroupGuid returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BitlinkHistory) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BitlinkHistory) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BitlinkHistory) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BitlinkHistory) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetIsActive

`func (o *BitlinkHistory) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *BitlinkHistory) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *BitlinkHistory) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *BitlinkHistory) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLongUrl

`func (o *BitlinkHistory) GetLongUrl() string`

GetLongUrl returns the LongUrl field if non-nil, zero value otherwise.

### GetLongUrlOk

`func (o *BitlinkHistory) GetLongUrlOk() (*string, bool)`

GetLongUrlOk returns a tuple with the LongUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongUrl

`func (o *BitlinkHistory) SetLongUrl(v string)`

SetLongUrl sets LongUrl field to given value.

### HasLongUrl

`func (o *BitlinkHistory) HasLongUrl() bool`

HasLongUrl returns a boolean if a field has been set.

### GetDeactivatedAt

`func (o *BitlinkHistory) GetDeactivatedAt() string`

GetDeactivatedAt returns the DeactivatedAt field if non-nil, zero value otherwise.

### GetDeactivatedAtOk

`func (o *BitlinkHistory) GetDeactivatedAtOk() (*string, bool)`

GetDeactivatedAtOk returns a tuple with the DeactivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivatedAt

`func (o *BitlinkHistory) SetDeactivatedAt(v string)`

SetDeactivatedAt sets DeactivatedAt field to given value.

### HasDeactivatedAt

`func (o *BitlinkHistory) HasDeactivatedAt() bool`

HasDeactivatedAt returns a boolean if a field has been set.

### GetLogin

`func (o *BitlinkHistory) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *BitlinkHistory) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *BitlinkHistory) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *BitlinkHistory) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetTargetBitlinkId

`func (o *BitlinkHistory) GetTargetBitlinkId() string`

GetTargetBitlinkId returns the TargetBitlinkId field if non-nil, zero value otherwise.

### GetTargetBitlinkIdOk

`func (o *BitlinkHistory) GetTargetBitlinkIdOk() (*string, bool)`

GetTargetBitlinkIdOk returns a tuple with the TargetBitlinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBitlinkId

`func (o *BitlinkHistory) SetTargetBitlinkId(v string)`

SetTargetBitlinkId sets TargetBitlinkId field to given value.

### HasTargetBitlinkId

`func (o *BitlinkHistory) HasTargetBitlinkId() bool`

HasTargetBitlinkId returns a boolean if a field has been set.

### GetId

`func (o *BitlinkHistory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BitlinkHistory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BitlinkHistory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BitlinkHistory) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


