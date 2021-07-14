# ClickstreamSink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Descriptive name for the sink | [optional] 
**Schedule** | Pointer to **string** | Delivery schedule | [optional] 
**Deactivated** | Pointer to **time.Time** | Time of deactivation, if any | [optional] 
**Destination** | Pointer to **string** | Location of the sink | [optional] 
**Modified** | Pointer to **time.Time** | Last modified time | [optional] 
**Created** | Pointer to **time.Time** | Sink creation time | [optional] 
**SinkOriginId** | Pointer to **int32** | Sink origin ID classification | [optional] 
**ClientId** | Pointer to **string** | Unique short id of client | [optional] 
**Provider** | Pointer to **string** | Provider of sink | [optional] 
**LastSync** | Pointer to **time.Time** | Last synced time | [optional] 
**Id** | Pointer to **string** | Unique short id | [optional] 
**ClickstreamId** | Pointer to **string** | Unique short id of clickstream | [optional] 

## Methods

### NewClickstreamSink

`func NewClickstreamSink() *ClickstreamSink`

NewClickstreamSink instantiates a new ClickstreamSink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClickstreamSinkWithDefaults

`func NewClickstreamSinkWithDefaults() *ClickstreamSink`

NewClickstreamSinkWithDefaults instantiates a new ClickstreamSink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClickstreamSink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClickstreamSink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClickstreamSink) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClickstreamSink) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchedule

`func (o *ClickstreamSink) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ClickstreamSink) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ClickstreamSink) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ClickstreamSink) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetDeactivated

`func (o *ClickstreamSink) GetDeactivated() time.Time`

GetDeactivated returns the Deactivated field if non-nil, zero value otherwise.

### GetDeactivatedOk

`func (o *ClickstreamSink) GetDeactivatedOk() (*time.Time, bool)`

GetDeactivatedOk returns a tuple with the Deactivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivated

`func (o *ClickstreamSink) SetDeactivated(v time.Time)`

SetDeactivated sets Deactivated field to given value.

### HasDeactivated

`func (o *ClickstreamSink) HasDeactivated() bool`

HasDeactivated returns a boolean if a field has been set.

### GetDestination

`func (o *ClickstreamSink) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ClickstreamSink) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ClickstreamSink) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *ClickstreamSink) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetModified

`func (o *ClickstreamSink) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ClickstreamSink) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ClickstreamSink) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ClickstreamSink) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetCreated

`func (o *ClickstreamSink) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ClickstreamSink) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ClickstreamSink) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ClickstreamSink) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetSinkOriginId

`func (o *ClickstreamSink) GetSinkOriginId() int32`

GetSinkOriginId returns the SinkOriginId field if non-nil, zero value otherwise.

### GetSinkOriginIdOk

`func (o *ClickstreamSink) GetSinkOriginIdOk() (*int32, bool)`

GetSinkOriginIdOk returns a tuple with the SinkOriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSinkOriginId

`func (o *ClickstreamSink) SetSinkOriginId(v int32)`

SetSinkOriginId sets SinkOriginId field to given value.

### HasSinkOriginId

`func (o *ClickstreamSink) HasSinkOriginId() bool`

HasSinkOriginId returns a boolean if a field has been set.

### GetClientId

`func (o *ClickstreamSink) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ClickstreamSink) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ClickstreamSink) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ClickstreamSink) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetProvider

`func (o *ClickstreamSink) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ClickstreamSink) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ClickstreamSink) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ClickstreamSink) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetLastSync

`func (o *ClickstreamSink) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *ClickstreamSink) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *ClickstreamSink) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *ClickstreamSink) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetId

`func (o *ClickstreamSink) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClickstreamSink) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClickstreamSink) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClickstreamSink) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClickstreamId

`func (o *ClickstreamSink) GetClickstreamId() string`

GetClickstreamId returns the ClickstreamId field if non-nil, zero value otherwise.

### GetClickstreamIdOk

`func (o *ClickstreamSink) GetClickstreamIdOk() (*string, bool)`

GetClickstreamIdOk returns a tuple with the ClickstreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickstreamId

`func (o *ClickstreamSink) SetClickstreamId(v string)`

SetClickstreamId sets ClickstreamId field to given value.

### HasClickstreamId

`func (o *ClickstreamSink) HasClickstreamId() bool`

HasClickstreamId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


