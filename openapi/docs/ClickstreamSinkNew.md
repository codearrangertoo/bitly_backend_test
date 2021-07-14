# ClickstreamSinkNew

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **string** | Location of the sink | [optional] 
**Name** | Pointer to **string** | Descriptive name for the sink | [optional] 
**AuthKey** | Pointer to **string** | Auth key for syncing to destination | [optional] 
**Provider** | Pointer to **string** | Provider of sink | [optional] 

## Methods

### NewClickstreamSinkNew

`func NewClickstreamSinkNew() *ClickstreamSinkNew`

NewClickstreamSinkNew instantiates a new ClickstreamSinkNew object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClickstreamSinkNewWithDefaults

`func NewClickstreamSinkNewWithDefaults() *ClickstreamSinkNew`

NewClickstreamSinkNewWithDefaults instantiates a new ClickstreamSinkNew object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *ClickstreamSinkNew) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ClickstreamSinkNew) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ClickstreamSinkNew) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *ClickstreamSinkNew) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetName

`func (o *ClickstreamSinkNew) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClickstreamSinkNew) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClickstreamSinkNew) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClickstreamSinkNew) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAuthKey

`func (o *ClickstreamSinkNew) GetAuthKey() string`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *ClickstreamSinkNew) GetAuthKeyOk() (*string, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *ClickstreamSinkNew) SetAuthKey(v string)`

SetAuthKey sets AuthKey field to given value.

### HasAuthKey

`func (o *ClickstreamSinkNew) HasAuthKey() bool`

HasAuthKey returns a boolean if a field has been set.

### GetProvider

`func (o *ClickstreamSinkNew) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ClickstreamSinkNew) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ClickstreamSinkNew) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ClickstreamSinkNew) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


