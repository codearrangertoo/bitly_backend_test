# WebhookUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**GroupGuid** | Pointer to **string** |  | [optional] 
**FetchTags** | Pointer to **bool** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**OrganizationGuid** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**OauthUrl** | Pointer to **string** |  | [optional] 
**Guid** | **string** |  | 
**Event** | Pointer to **string** |  | [optional] 

## Methods

### NewWebhookUpdate

`func NewWebhookUpdate(guid string, ) *WebhookUpdate`

NewWebhookUpdate instantiates a new WebhookUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookUpdateWithDefaults

`func NewWebhookUpdateWithDefaults() *WebhookUpdate`

NewWebhookUpdateWithDefaults instantiates a new WebhookUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhookUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *WebhookUpdate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookUpdate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookUpdate) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WebhookUpdate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetGroupGuid

`func (o *WebhookUpdate) GetGroupGuid() string`

GetGroupGuid returns the GroupGuid field if non-nil, zero value otherwise.

### GetGroupGuidOk

`func (o *WebhookUpdate) GetGroupGuidOk() (*string, bool)`

GetGroupGuidOk returns a tuple with the GroupGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupGuid

`func (o *WebhookUpdate) SetGroupGuid(v string)`

SetGroupGuid sets GroupGuid field to given value.

### HasGroupGuid

`func (o *WebhookUpdate) HasGroupGuid() bool`

HasGroupGuid returns a boolean if a field has been set.

### GetFetchTags

`func (o *WebhookUpdate) GetFetchTags() bool`

GetFetchTags returns the FetchTags field if non-nil, zero value otherwise.

### GetFetchTagsOk

`func (o *WebhookUpdate) GetFetchTagsOk() (*bool, bool)`

GetFetchTagsOk returns a tuple with the FetchTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchTags

`func (o *WebhookUpdate) SetFetchTags(v bool)`

SetFetchTags sets FetchTags field to given value.

### HasFetchTags

`func (o *WebhookUpdate) HasFetchTags() bool`

HasFetchTags returns a boolean if a field has been set.

### GetIsActive

`func (o *WebhookUpdate) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *WebhookUpdate) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *WebhookUpdate) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *WebhookUpdate) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetOrganizationGuid

`func (o *WebhookUpdate) GetOrganizationGuid() string`

GetOrganizationGuid returns the OrganizationGuid field if non-nil, zero value otherwise.

### GetOrganizationGuidOk

`func (o *WebhookUpdate) GetOrganizationGuidOk() (*string, bool)`

GetOrganizationGuidOk returns a tuple with the OrganizationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationGuid

`func (o *WebhookUpdate) SetOrganizationGuid(v string)`

SetOrganizationGuid sets OrganizationGuid field to given value.

### HasOrganizationGuid

`func (o *WebhookUpdate) HasOrganizationGuid() bool`

HasOrganizationGuid returns a boolean if a field has been set.

### GetClientId

`func (o *WebhookUpdate) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *WebhookUpdate) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *WebhookUpdate) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *WebhookUpdate) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *WebhookUpdate) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *WebhookUpdate) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *WebhookUpdate) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *WebhookUpdate) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetOauthUrl

`func (o *WebhookUpdate) GetOauthUrl() string`

GetOauthUrl returns the OauthUrl field if non-nil, zero value otherwise.

### GetOauthUrlOk

`func (o *WebhookUpdate) GetOauthUrlOk() (*string, bool)`

GetOauthUrlOk returns a tuple with the OauthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthUrl

`func (o *WebhookUpdate) SetOauthUrl(v string)`

SetOauthUrl sets OauthUrl field to given value.

### HasOauthUrl

`func (o *WebhookUpdate) HasOauthUrl() bool`

HasOauthUrl returns a boolean if a field has been set.

### GetGuid

`func (o *WebhookUpdate) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *WebhookUpdate) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *WebhookUpdate) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetEvent

`func (o *WebhookUpdate) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookUpdate) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookUpdate) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *WebhookUpdate) HasEvent() bool`

HasEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


