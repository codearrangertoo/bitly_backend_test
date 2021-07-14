# NSQProducer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpPort** | Pointer to **int32** |  | [optional] 
**BroadcastAddress** | Pointer to **string** |  | [optional] 
**TcpPort** | Pointer to **int32** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**RemoteAddress** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewNSQProducer

`func NewNSQProducer() *NSQProducer`

NewNSQProducer instantiates a new NSQProducer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNSQProducerWithDefaults

`func NewNSQProducerWithDefaults() *NSQProducer`

NewNSQProducerWithDefaults instantiates a new NSQProducer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpPort

`func (o *NSQProducer) GetHttpPort() int32`

GetHttpPort returns the HttpPort field if non-nil, zero value otherwise.

### GetHttpPortOk

`func (o *NSQProducer) GetHttpPortOk() (*int32, bool)`

GetHttpPortOk returns a tuple with the HttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpPort

`func (o *NSQProducer) SetHttpPort(v int32)`

SetHttpPort sets HttpPort field to given value.

### HasHttpPort

`func (o *NSQProducer) HasHttpPort() bool`

HasHttpPort returns a boolean if a field has been set.

### GetBroadcastAddress

`func (o *NSQProducer) GetBroadcastAddress() string`

GetBroadcastAddress returns the BroadcastAddress field if non-nil, zero value otherwise.

### GetBroadcastAddressOk

`func (o *NSQProducer) GetBroadcastAddressOk() (*string, bool)`

GetBroadcastAddressOk returns a tuple with the BroadcastAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastAddress

`func (o *NSQProducer) SetBroadcastAddress(v string)`

SetBroadcastAddress sets BroadcastAddress field to given value.

### HasBroadcastAddress

`func (o *NSQProducer) HasBroadcastAddress() bool`

HasBroadcastAddress returns a boolean if a field has been set.

### GetTcpPort

`func (o *NSQProducer) GetTcpPort() int32`

GetTcpPort returns the TcpPort field if non-nil, zero value otherwise.

### GetTcpPortOk

`func (o *NSQProducer) GetTcpPortOk() (*int32, bool)`

GetTcpPortOk returns a tuple with the TcpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpPort

`func (o *NSQProducer) SetTcpPort(v int32)`

SetTcpPort sets TcpPort field to given value.

### HasTcpPort

`func (o *NSQProducer) HasTcpPort() bool`

HasTcpPort returns a boolean if a field has been set.

### GetHostname

`func (o *NSQProducer) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NSQProducer) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NSQProducer) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *NSQProducer) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *NSQProducer) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *NSQProducer) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *NSQProducer) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *NSQProducer) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetVersion

`func (o *NSQProducer) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NSQProducer) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NSQProducer) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NSQProducer) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


