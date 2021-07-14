# CustomDomains

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpgradeRequired** | Pointer to **bool** |  | [optional] 
**CustomDomains** | Pointer to [**[]CustomDomainBody**](CustomDomainBody.md) |  | [optional] 

## Methods

### NewCustomDomains

`func NewCustomDomains() *CustomDomains`

NewCustomDomains instantiates a new CustomDomains object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDomainsWithDefaults

`func NewCustomDomainsWithDefaults() *CustomDomains`

NewCustomDomainsWithDefaults instantiates a new CustomDomains object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpgradeRequired

`func (o *CustomDomains) GetUpgradeRequired() bool`

GetUpgradeRequired returns the UpgradeRequired field if non-nil, zero value otherwise.

### GetUpgradeRequiredOk

`func (o *CustomDomains) GetUpgradeRequiredOk() (*bool, bool)`

GetUpgradeRequiredOk returns a tuple with the UpgradeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeRequired

`func (o *CustomDomains) SetUpgradeRequired(v bool)`

SetUpgradeRequired sets UpgradeRequired field to given value.

### HasUpgradeRequired

`func (o *CustomDomains) HasUpgradeRequired() bool`

HasUpgradeRequired returns a boolean if a field has been set.

### GetCustomDomains

`func (o *CustomDomains) GetCustomDomains() []CustomDomainBody`

GetCustomDomains returns the CustomDomains field if non-nil, zero value otherwise.

### GetCustomDomainsOk

`func (o *CustomDomains) GetCustomDomainsOk() (*[]CustomDomainBody, bool)`

GetCustomDomainsOk returns a tuple with the CustomDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomains

`func (o *CustomDomains) SetCustomDomains(v []CustomDomainBody)`

SetCustomDomains sets CustomDomains field to given value.

### HasCustomDomains

`func (o *CustomDomains) HasCustomDomains() bool`

HasCustomDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


