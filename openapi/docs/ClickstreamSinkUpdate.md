# ClickstreamSinkUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **string** | Location of the sink | [optional] 
**Name** | Pointer to **string** | Descriptive name for the sink | [optional] 
**AuthKey** | Pointer to **string** | Auth key for syncing to destination | [optional] 
**Provider** | Pointer to **string** | Provider of sink | [optional] 
**Deactivated** | Pointer to **bool** | False if sink should be activated | [optional] 
**Schedule** | Pointer to **string** | Delivery schedule | [optional] 

## Methods

### NewClickstreamSinkUpdate

`func NewClickstreamSinkUpdate() *ClickstreamSinkUpdate`

NewClickstreamSinkUpdate instantiates a new ClickstreamSinkUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClickstreamSinkUpdateWithDefaults

`func NewClickstreamSinkUpdateWithDefaults() *ClickstreamSinkUpdate`

NewClickstreamSinkUpdateWithDefaults instantiates a new ClickstreamSinkUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *ClickstreamSinkUpdate) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ClickstreamSinkUpdate) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ClickstreamSinkUpdate) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *ClickstreamSinkUpdate) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetName

`func (o *ClickstreamSinkUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClickstreamSinkUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClickstreamSinkUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClickstreamSinkUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAuthKey

`func (o *ClickstreamSinkUpdate) GetAuthKey() string`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *ClickstreamSinkUpdate) GetAuthKeyOk() (*string, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *ClickstreamSinkUpdate) SetAuthKey(v string)`

SetAuthKey sets AuthKey field to given value.

### HasAuthKey

`func (o *ClickstreamSinkUpdate) HasAuthKey() bool`

HasAuthKey returns a boolean if a field has been set.

### GetProvider

`func (o *ClickstreamSinkUpdate) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ClickstreamSinkUpdate) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ClickstreamSinkUpdate) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ClickstreamSinkUpdate) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetDeactivated

`func (o *ClickstreamSinkUpdate) GetDeactivated() bool`

GetDeactivated returns the Deactivated field if non-nil, zero value otherwise.

### GetDeactivatedOk

`func (o *ClickstreamSinkUpdate) GetDeactivatedOk() (*bool, bool)`

GetDeactivatedOk returns a tuple with the Deactivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivated

`func (o *ClickstreamSinkUpdate) SetDeactivated(v bool)`

SetDeactivated sets Deactivated field to given value.

### HasDeactivated

`func (o *ClickstreamSinkUpdate) HasDeactivated() bool`

HasDeactivated returns a boolean if a field has been set.

### GetSchedule

`func (o *ClickstreamSinkUpdate) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ClickstreamSinkUpdate) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ClickstreamSinkUpdate) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ClickstreamSinkUpdate) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


