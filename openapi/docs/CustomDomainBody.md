# CustomDomainBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeeplinkApps** | Pointer to [**[]MinimalDeeplinkApp**](MinimalDeeplinkApp.md) |  | [optional] 
**HstsEnabled** | Pointer to **bool** |  | [optional] 
**UpgradeInsecureRequests** | Pointer to **bool** |  | [optional] 
**CustomDomain** | Pointer to **string** |  | [optional] 
**ValidationError** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **int32** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**RootRedirect** | Pointer to **string** |  | [optional] 
**SslConfigurationError** | Pointer to **string** |  | [optional] 
**SslAutoconfigureError** | Pointer to **bool** |  | [optional] 
**GroupGuids** | Pointer to **[]string** |  | [optional] 
**SslCert** | Pointer to [**CustomDomainBodySslCert**](CustomDomainBodySslCert.md) |  | [optional] 
**WildcardRedirect** | Pointer to **string** |  | [optional] 
**ConfigurationLastCheckTs** | Pointer to **string** |  | [optional] 
**HttpsEnabled** | Pointer to **bool** |  | [optional] 
**ValidationStatus** | Pointer to **string** |  | [optional] 
**HttpsBitlinks** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustomDomainBody

`func NewCustomDomainBody() *CustomDomainBody`

NewCustomDomainBody instantiates a new CustomDomainBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDomainBodyWithDefaults

`func NewCustomDomainBodyWithDefaults() *CustomDomainBody`

NewCustomDomainBodyWithDefaults instantiates a new CustomDomainBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeeplinkApps

`func (o *CustomDomainBody) GetDeeplinkApps() []MinimalDeeplinkApp`

GetDeeplinkApps returns the DeeplinkApps field if non-nil, zero value otherwise.

### GetDeeplinkAppsOk

`func (o *CustomDomainBody) GetDeeplinkAppsOk() (*[]MinimalDeeplinkApp, bool)`

GetDeeplinkAppsOk returns a tuple with the DeeplinkApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeeplinkApps

`func (o *CustomDomainBody) SetDeeplinkApps(v []MinimalDeeplinkApp)`

SetDeeplinkApps sets DeeplinkApps field to given value.

### HasDeeplinkApps

`func (o *CustomDomainBody) HasDeeplinkApps() bool`

HasDeeplinkApps returns a boolean if a field has been set.

### GetHstsEnabled

`func (o *CustomDomainBody) GetHstsEnabled() bool`

GetHstsEnabled returns the HstsEnabled field if non-nil, zero value otherwise.

### GetHstsEnabledOk

`func (o *CustomDomainBody) GetHstsEnabledOk() (*bool, bool)`

GetHstsEnabledOk returns a tuple with the HstsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsEnabled

`func (o *CustomDomainBody) SetHstsEnabled(v bool)`

SetHstsEnabled sets HstsEnabled field to given value.

### HasHstsEnabled

`func (o *CustomDomainBody) HasHstsEnabled() bool`

HasHstsEnabled returns a boolean if a field has been set.

### GetUpgradeInsecureRequests

`func (o *CustomDomainBody) GetUpgradeInsecureRequests() bool`

GetUpgradeInsecureRequests returns the UpgradeInsecureRequests field if non-nil, zero value otherwise.

### GetUpgradeInsecureRequestsOk

`func (o *CustomDomainBody) GetUpgradeInsecureRequestsOk() (*bool, bool)`

GetUpgradeInsecureRequestsOk returns a tuple with the UpgradeInsecureRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeInsecureRequests

`func (o *CustomDomainBody) SetUpgradeInsecureRequests(v bool)`

SetUpgradeInsecureRequests sets UpgradeInsecureRequests field to given value.

### HasUpgradeInsecureRequests

`func (o *CustomDomainBody) HasUpgradeInsecureRequests() bool`

HasUpgradeInsecureRequests returns a boolean if a field has been set.

### GetCustomDomain

`func (o *CustomDomainBody) GetCustomDomain() string`

GetCustomDomain returns the CustomDomain field if non-nil, zero value otherwise.

### GetCustomDomainOk

`func (o *CustomDomainBody) GetCustomDomainOk() (*string, bool)`

GetCustomDomainOk returns a tuple with the CustomDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomain

`func (o *CustomDomainBody) SetCustomDomain(v string)`

SetCustomDomain sets CustomDomain field to given value.

### HasCustomDomain

`func (o *CustomDomainBody) HasCustomDomain() bool`

HasCustomDomain returns a boolean if a field has been set.

### GetValidationError

`func (o *CustomDomainBody) GetValidationError() string`

GetValidationError returns the ValidationError field if non-nil, zero value otherwise.

### GetValidationErrorOk

`func (o *CustomDomainBody) GetValidationErrorOk() (*string, bool)`

GetValidationErrorOk returns a tuple with the ValidationError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationError

`func (o *CustomDomainBody) SetValidationError(v string)`

SetValidationError sets ValidationError field to given value.

### HasValidationError

`func (o *CustomDomainBody) HasValidationError() bool`

HasValidationError returns a boolean if a field has been set.

### GetCreated

`func (o *CustomDomainBody) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomDomainBody) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CustomDomainBody) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CustomDomainBody) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetIsActive

`func (o *CustomDomainBody) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CustomDomainBody) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CustomDomainBody) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CustomDomainBody) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetRootRedirect

`func (o *CustomDomainBody) GetRootRedirect() string`

GetRootRedirect returns the RootRedirect field if non-nil, zero value otherwise.

### GetRootRedirectOk

`func (o *CustomDomainBody) GetRootRedirectOk() (*string, bool)`

GetRootRedirectOk returns a tuple with the RootRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootRedirect

`func (o *CustomDomainBody) SetRootRedirect(v string)`

SetRootRedirect sets RootRedirect field to given value.

### HasRootRedirect

`func (o *CustomDomainBody) HasRootRedirect() bool`

HasRootRedirect returns a boolean if a field has been set.

### GetSslConfigurationError

`func (o *CustomDomainBody) GetSslConfigurationError() string`

GetSslConfigurationError returns the SslConfigurationError field if non-nil, zero value otherwise.

### GetSslConfigurationErrorOk

`func (o *CustomDomainBody) GetSslConfigurationErrorOk() (*string, bool)`

GetSslConfigurationErrorOk returns a tuple with the SslConfigurationError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslConfigurationError

`func (o *CustomDomainBody) SetSslConfigurationError(v string)`

SetSslConfigurationError sets SslConfigurationError field to given value.

### HasSslConfigurationError

`func (o *CustomDomainBody) HasSslConfigurationError() bool`

HasSslConfigurationError returns a boolean if a field has been set.

### GetSslAutoconfigureError

`func (o *CustomDomainBody) GetSslAutoconfigureError() bool`

GetSslAutoconfigureError returns the SslAutoconfigureError field if non-nil, zero value otherwise.

### GetSslAutoconfigureErrorOk

`func (o *CustomDomainBody) GetSslAutoconfigureErrorOk() (*bool, bool)`

GetSslAutoconfigureErrorOk returns a tuple with the SslAutoconfigureError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslAutoconfigureError

`func (o *CustomDomainBody) SetSslAutoconfigureError(v bool)`

SetSslAutoconfigureError sets SslAutoconfigureError field to given value.

### HasSslAutoconfigureError

`func (o *CustomDomainBody) HasSslAutoconfigureError() bool`

HasSslAutoconfigureError returns a boolean if a field has been set.

### GetGroupGuids

`func (o *CustomDomainBody) GetGroupGuids() []string`

GetGroupGuids returns the GroupGuids field if non-nil, zero value otherwise.

### GetGroupGuidsOk

`func (o *CustomDomainBody) GetGroupGuidsOk() (*[]string, bool)`

GetGroupGuidsOk returns a tuple with the GroupGuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupGuids

`func (o *CustomDomainBody) SetGroupGuids(v []string)`

SetGroupGuids sets GroupGuids field to given value.

### HasGroupGuids

`func (o *CustomDomainBody) HasGroupGuids() bool`

HasGroupGuids returns a boolean if a field has been set.

### GetSslCert

`func (o *CustomDomainBody) GetSslCert() CustomDomainBodySslCert`

GetSslCert returns the SslCert field if non-nil, zero value otherwise.

### GetSslCertOk

`func (o *CustomDomainBody) GetSslCertOk() (*CustomDomainBodySslCert, bool)`

GetSslCertOk returns a tuple with the SslCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCert

`func (o *CustomDomainBody) SetSslCert(v CustomDomainBodySslCert)`

SetSslCert sets SslCert field to given value.

### HasSslCert

`func (o *CustomDomainBody) HasSslCert() bool`

HasSslCert returns a boolean if a field has been set.

### GetWildcardRedirect

`func (o *CustomDomainBody) GetWildcardRedirect() string`

GetWildcardRedirect returns the WildcardRedirect field if non-nil, zero value otherwise.

### GetWildcardRedirectOk

`func (o *CustomDomainBody) GetWildcardRedirectOk() (*string, bool)`

GetWildcardRedirectOk returns a tuple with the WildcardRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcardRedirect

`func (o *CustomDomainBody) SetWildcardRedirect(v string)`

SetWildcardRedirect sets WildcardRedirect field to given value.

### HasWildcardRedirect

`func (o *CustomDomainBody) HasWildcardRedirect() bool`

HasWildcardRedirect returns a boolean if a field has been set.

### GetConfigurationLastCheckTs

`func (o *CustomDomainBody) GetConfigurationLastCheckTs() string`

GetConfigurationLastCheckTs returns the ConfigurationLastCheckTs field if non-nil, zero value otherwise.

### GetConfigurationLastCheckTsOk

`func (o *CustomDomainBody) GetConfigurationLastCheckTsOk() (*string, bool)`

GetConfigurationLastCheckTsOk returns a tuple with the ConfigurationLastCheckTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationLastCheckTs

`func (o *CustomDomainBody) SetConfigurationLastCheckTs(v string)`

SetConfigurationLastCheckTs sets ConfigurationLastCheckTs field to given value.

### HasConfigurationLastCheckTs

`func (o *CustomDomainBody) HasConfigurationLastCheckTs() bool`

HasConfigurationLastCheckTs returns a boolean if a field has been set.

### GetHttpsEnabled

`func (o *CustomDomainBody) GetHttpsEnabled() bool`

GetHttpsEnabled returns the HttpsEnabled field if non-nil, zero value otherwise.

### GetHttpsEnabledOk

`func (o *CustomDomainBody) GetHttpsEnabledOk() (*bool, bool)`

GetHttpsEnabledOk returns a tuple with the HttpsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsEnabled

`func (o *CustomDomainBody) SetHttpsEnabled(v bool)`

SetHttpsEnabled sets HttpsEnabled field to given value.

### HasHttpsEnabled

`func (o *CustomDomainBody) HasHttpsEnabled() bool`

HasHttpsEnabled returns a boolean if a field has been set.

### GetValidationStatus

`func (o *CustomDomainBody) GetValidationStatus() string`

GetValidationStatus returns the ValidationStatus field if non-nil, zero value otherwise.

### GetValidationStatusOk

`func (o *CustomDomainBody) GetValidationStatusOk() (*string, bool)`

GetValidationStatusOk returns a tuple with the ValidationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationStatus

`func (o *CustomDomainBody) SetValidationStatus(v string)`

SetValidationStatus sets ValidationStatus field to given value.

### HasValidationStatus

`func (o *CustomDomainBody) HasValidationStatus() bool`

HasValidationStatus returns a boolean if a field has been set.

### GetHttpsBitlinks

`func (o *CustomDomainBody) GetHttpsBitlinks() bool`

GetHttpsBitlinks returns the HttpsBitlinks field if non-nil, zero value otherwise.

### GetHttpsBitlinksOk

`func (o *CustomDomainBody) GetHttpsBitlinksOk() (*bool, bool)`

GetHttpsBitlinksOk returns a tuple with the HttpsBitlinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsBitlinks

`func (o *CustomDomainBody) SetHttpsBitlinks(v bool)`

SetHttpsBitlinks sets HttpsBitlinks field to given value.

### HasHttpsBitlinks

`func (o *CustomDomainBody) HasHttpsBitlinks() bool`

HasHttpsBitlinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


