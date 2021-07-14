# Clickstream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Descriptive name for the clickstream | [optional] 
**Created** | Pointer to **time.Time** | Clickstream creation time | [optional] 
**Deactivated** | Pointer to **time.Time** | Time of deactivation, if any | [optional] 
**Modified** | Pointer to **time.Time** | Last modified time | [optional] 
**Query** | Pointer to **string** | Clickstream query | [optional] 
**Id** | Pointer to **string** | Unique short id | [optional] 

## Methods

### NewClickstream

`func NewClickstream() *Clickstream`

NewClickstream instantiates a new Clickstream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClickstreamWithDefaults

`func NewClickstreamWithDefaults() *Clickstream`

NewClickstreamWithDefaults instantiates a new Clickstream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Clickstream) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Clickstream) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Clickstream) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Clickstream) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreated

`func (o *Clickstream) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Clickstream) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Clickstream) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Clickstream) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDeactivated

`func (o *Clickstream) GetDeactivated() time.Time`

GetDeactivated returns the Deactivated field if non-nil, zero value otherwise.

### GetDeactivatedOk

`func (o *Clickstream) GetDeactivatedOk() (*time.Time, bool)`

GetDeactivatedOk returns a tuple with the Deactivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivated

`func (o *Clickstream) SetDeactivated(v time.Time)`

SetDeactivated sets Deactivated field to given value.

### HasDeactivated

`func (o *Clickstream) HasDeactivated() bool`

HasDeactivated returns a boolean if a field has been set.

### GetModified

`func (o *Clickstream) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Clickstream) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Clickstream) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Clickstream) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetQuery

`func (o *Clickstream) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Clickstream) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Clickstream) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *Clickstream) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetId

`func (o *Clickstream) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Clickstream) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Clickstream) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Clickstream) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


