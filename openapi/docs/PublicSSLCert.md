# PublicSSLCert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomCert** | Pointer to **bool** |  | [optional] 
**ValidEnd** | Pointer to **int32** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 

## Methods

### NewPublicSSLCert

`func NewPublicSSLCert() *PublicSSLCert`

NewPublicSSLCert instantiates a new PublicSSLCert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSSLCertWithDefaults

`func NewPublicSSLCertWithDefaults() *PublicSSLCert`

NewPublicSSLCertWithDefaults instantiates a new PublicSSLCert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomCert

`func (o *PublicSSLCert) GetCustomCert() bool`

GetCustomCert returns the CustomCert field if non-nil, zero value otherwise.

### GetCustomCertOk

`func (o *PublicSSLCert) GetCustomCertOk() (*bool, bool)`

GetCustomCertOk returns a tuple with the CustomCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCert

`func (o *PublicSSLCert) SetCustomCert(v bool)`

SetCustomCert sets CustomCert field to given value.

### HasCustomCert

`func (o *PublicSSLCert) HasCustomCert() bool`

HasCustomCert returns a boolean if a field has been set.

### GetValidEnd

`func (o *PublicSSLCert) GetValidEnd() int32`

GetValidEnd returns the ValidEnd field if non-nil, zero value otherwise.

### GetValidEndOk

`func (o *PublicSSLCert) GetValidEndOk() (*int32, bool)`

GetValidEndOk returns a tuple with the ValidEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidEnd

`func (o *PublicSSLCert) SetValidEnd(v int32)`

SetValidEnd sets ValidEnd field to given value.

### HasValidEnd

`func (o *PublicSSLCert) HasValidEnd() bool`

HasValidEnd returns a boolean if a field has been set.

### GetIssuer

`func (o *PublicSSLCert) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *PublicSSLCert) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *PublicSSLCert) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *PublicSSLCert) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


