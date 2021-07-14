# OAuthApps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**[]OAuthApp**](OAuthApp.md) |  | [optional] 

## Methods

### NewOAuthApps

`func NewOAuthApps() *OAuthApps`

NewOAuthApps instantiates a new OAuthApps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthAppsWithDefaults

`func NewOAuthAppsWithDefaults() *OAuthApps`

NewOAuthAppsWithDefaults instantiates a new OAuthApps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *OAuthApps) GetApplications() []OAuthApp`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *OAuthApps) GetApplicationsOk() (*[]OAuthApp, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *OAuthApps) SetApplications(v []OAuthApp)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *OAuthApps) HasApplications() bool`

HasApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


