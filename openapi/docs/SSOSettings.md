# SSOSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SloEndpoint** | Pointer to **string** |  | [optional] 
**DefaultGroupGuid** | Pointer to **string** |  | [optional] 
**IdpUrl** | Pointer to **string** |  | [optional] 
**Certificate** | Pointer to **string** |  | [optional] 
**OrganizationGuid** | Pointer to **string** |  | [optional] 
**SamlEndpoint** | Pointer to **string** |  | [optional] 
**IdentityProvider** | Pointer to **string** |  | [optional] 
**Domains** | Pointer to **[]string** |  | [optional] 
**IssuerUrl** | Pointer to **string** |  | [optional] 
**UrlSlug** | Pointer to **string** |  | [optional] 

## Methods

### NewSSOSettings

`func NewSSOSettings() *SSOSettings`

NewSSOSettings instantiates a new SSOSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSOSettingsWithDefaults

`func NewSSOSettingsWithDefaults() *SSOSettings`

NewSSOSettingsWithDefaults instantiates a new SSOSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSloEndpoint

`func (o *SSOSettings) GetSloEndpoint() string`

GetSloEndpoint returns the SloEndpoint field if non-nil, zero value otherwise.

### GetSloEndpointOk

`func (o *SSOSettings) GetSloEndpointOk() (*string, bool)`

GetSloEndpointOk returns a tuple with the SloEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloEndpoint

`func (o *SSOSettings) SetSloEndpoint(v string)`

SetSloEndpoint sets SloEndpoint field to given value.

### HasSloEndpoint

`func (o *SSOSettings) HasSloEndpoint() bool`

HasSloEndpoint returns a boolean if a field has been set.

### GetDefaultGroupGuid

`func (o *SSOSettings) GetDefaultGroupGuid() string`

GetDefaultGroupGuid returns the DefaultGroupGuid field if non-nil, zero value otherwise.

### GetDefaultGroupGuidOk

`func (o *SSOSettings) GetDefaultGroupGuidOk() (*string, bool)`

GetDefaultGroupGuidOk returns a tuple with the DefaultGroupGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroupGuid

`func (o *SSOSettings) SetDefaultGroupGuid(v string)`

SetDefaultGroupGuid sets DefaultGroupGuid field to given value.

### HasDefaultGroupGuid

`func (o *SSOSettings) HasDefaultGroupGuid() bool`

HasDefaultGroupGuid returns a boolean if a field has been set.

### GetIdpUrl

`func (o *SSOSettings) GetIdpUrl() string`

GetIdpUrl returns the IdpUrl field if non-nil, zero value otherwise.

### GetIdpUrlOk

`func (o *SSOSettings) GetIdpUrlOk() (*string, bool)`

GetIdpUrlOk returns a tuple with the IdpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpUrl

`func (o *SSOSettings) SetIdpUrl(v string)`

SetIdpUrl sets IdpUrl field to given value.

### HasIdpUrl

`func (o *SSOSettings) HasIdpUrl() bool`

HasIdpUrl returns a boolean if a field has been set.

### GetCertificate

`func (o *SSOSettings) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SSOSettings) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SSOSettings) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *SSOSettings) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetOrganizationGuid

`func (o *SSOSettings) GetOrganizationGuid() string`

GetOrganizationGuid returns the OrganizationGuid field if non-nil, zero value otherwise.

### GetOrganizationGuidOk

`func (o *SSOSettings) GetOrganizationGuidOk() (*string, bool)`

GetOrganizationGuidOk returns a tuple with the OrganizationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationGuid

`func (o *SSOSettings) SetOrganizationGuid(v string)`

SetOrganizationGuid sets OrganizationGuid field to given value.

### HasOrganizationGuid

`func (o *SSOSettings) HasOrganizationGuid() bool`

HasOrganizationGuid returns a boolean if a field has been set.

### GetSamlEndpoint

`func (o *SSOSettings) GetSamlEndpoint() string`

GetSamlEndpoint returns the SamlEndpoint field if non-nil, zero value otherwise.

### GetSamlEndpointOk

`func (o *SSOSettings) GetSamlEndpointOk() (*string, bool)`

GetSamlEndpointOk returns a tuple with the SamlEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlEndpoint

`func (o *SSOSettings) SetSamlEndpoint(v string)`

SetSamlEndpoint sets SamlEndpoint field to given value.

### HasSamlEndpoint

`func (o *SSOSettings) HasSamlEndpoint() bool`

HasSamlEndpoint returns a boolean if a field has been set.

### GetIdentityProvider

`func (o *SSOSettings) GetIdentityProvider() string`

GetIdentityProvider returns the IdentityProvider field if non-nil, zero value otherwise.

### GetIdentityProviderOk

`func (o *SSOSettings) GetIdentityProviderOk() (*string, bool)`

GetIdentityProviderOk returns a tuple with the IdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProvider

`func (o *SSOSettings) SetIdentityProvider(v string)`

SetIdentityProvider sets IdentityProvider field to given value.

### HasIdentityProvider

`func (o *SSOSettings) HasIdentityProvider() bool`

HasIdentityProvider returns a boolean if a field has been set.

### GetDomains

`func (o *SSOSettings) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *SSOSettings) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *SSOSettings) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *SSOSettings) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetIssuerUrl

`func (o *SSOSettings) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *SSOSettings) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *SSOSettings) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.

### HasIssuerUrl

`func (o *SSOSettings) HasIssuerUrl() bool`

HasIssuerUrl returns a boolean if a field has been set.

### GetUrlSlug

`func (o *SSOSettings) GetUrlSlug() string`

GetUrlSlug returns the UrlSlug field if non-nil, zero value otherwise.

### GetUrlSlugOk

`func (o *SSOSettings) GetUrlSlugOk() (*string, bool)`

GetUrlSlugOk returns a tuple with the UrlSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlSlug

`func (o *SSOSettings) SetUrlSlug(v string)`

SetUrlSlug sets UrlSlug field to given value.

### HasUrlSlug

`func (o *SSOSettings) HasUrlSlug() bool`

HasUrlSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


