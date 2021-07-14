# ReferrersByDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Referrers** | Pointer to [**[]Metric**](Metric.md) |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 

## Methods

### NewReferrersByDomain

`func NewReferrersByDomain() *ReferrersByDomain`

NewReferrersByDomain instantiates a new ReferrersByDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferrersByDomainWithDefaults

`func NewReferrersByDomainWithDefaults() *ReferrersByDomain`

NewReferrersByDomainWithDefaults instantiates a new ReferrersByDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferrers

`func (o *ReferrersByDomain) GetReferrers() []Metric`

GetReferrers returns the Referrers field if non-nil, zero value otherwise.

### GetReferrersOk

`func (o *ReferrersByDomain) GetReferrersOk() (*[]Metric, bool)`

GetReferrersOk returns a tuple with the Referrers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrers

`func (o *ReferrersByDomain) SetReferrers(v []Metric)`

SetReferrers sets Referrers field to given value.

### HasReferrers

`func (o *ReferrersByDomain) HasReferrers() bool`

HasReferrers returns a boolean if a field has been set.

### GetNetwork

`func (o *ReferrersByDomain) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ReferrersByDomain) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ReferrersByDomain) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ReferrersByDomain) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


