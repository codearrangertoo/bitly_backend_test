# Authorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to [**OAuthAppWithOwnerLogin**](OAuthAppWithOwnerLogin.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** | ISO timestamp | [optional] 

## Methods

### NewAuthorization

`func NewAuthorization() *Authorization`

NewAuthorization instantiates a new Authorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationWithDefaults

`func NewAuthorizationWithDefaults() *Authorization`

NewAuthorizationWithDefaults instantiates a new Authorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *Authorization) GetApp() OAuthAppWithOwnerLogin`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *Authorization) GetAppOk() (*OAuthAppWithOwnerLogin, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *Authorization) SetApp(v OAuthAppWithOwnerLogin)`

SetApp sets App field to given value.

### HasApp

`func (o *Authorization) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetId

`func (o *Authorization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Authorization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Authorization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Authorization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientId

`func (o *Authorization) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Authorization) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Authorization) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Authorization) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetCreated

`func (o *Authorization) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Authorization) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Authorization) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Authorization) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


