# WebhookCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Url** | **string** |  | 
**GroupGuid** | Pointer to **string** |  | [optional] 
**FetchTags** | Pointer to **bool** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**OrganizationGuid** | **string** |  | 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**OauthUrl** | Pointer to **string** |  | [optional] 
**Event** | **string** |  | 

## Methods

### NewWebhookCreate

`func NewWebhookCreate(name string, url string, organizationGuid string, event string, ) *WebhookCreate`

NewWebhookCreate instantiates a new WebhookCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookCreateWithDefaults

`func NewWebhookCreateWithDefaults() *WebhookCreate`

NewWebhookCreateWithDefaults instantiates a new WebhookCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhookCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookCreate) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *WebhookCreate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookCreate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookCreate) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetGroupGuid

`func (o *WebhookCreate) GetGroupGuid() string`

GetGroupGuid returns the GroupGuid field if non-nil, zero value otherwise.

### GetGroupGuidOk

`func (o *WebhookCreate) GetGroupGuidOk() (*string, bool)`

GetGroupGuidOk returns a tuple with the GroupGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupGuid

`func (o *WebhookCreate) SetGroupGuid(v string)`

SetGroupGuid sets GroupGuid field to given value.

### HasGroupGuid

`func (o *WebhookCreate) HasGroupGuid() bool`

HasGroupGuid returns a boolean if a field has been set.

### GetFetchTags

`func (o *WebhookCreate) GetFetchTags() bool`

GetFetchTags returns the FetchTags field if non-nil, zero value otherwise.

### GetFetchTagsOk

`func (o *WebhookCreate) GetFetchTagsOk() (*bool, bool)`

GetFetchTagsOk returns a tuple with the FetchTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchTags

`func (o *WebhookCreate) SetFetchTags(v bool)`

SetFetchTags sets FetchTags field to given value.

### HasFetchTags

`func (o *WebhookCreate) HasFetchTags() bool`

HasFetchTags returns a boolean if a field has been set.

### GetIsActive

`func (o *WebhookCreate) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *WebhookCreate) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *WebhookCreate) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *WebhookCreate) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetOrganizationGuid

`func (o *WebhookCreate) GetOrganizationGuid() string`

GetOrganizationGuid returns the OrganizationGuid field if non-nil, zero value otherwise.

### GetOrganizationGuidOk

`func (o *WebhookCreate) GetOrganizationGuidOk() (*string, bool)`

GetOrganizationGuidOk returns a tuple with the OrganizationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationGuid

`func (o *WebhookCreate) SetOrganizationGuid(v string)`

SetOrganizationGuid sets OrganizationGuid field to given value.


### GetClientId

`func (o *WebhookCreate) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *WebhookCreate) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *WebhookCreate) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *WebhookCreate) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *WebhookCreate) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *WebhookCreate) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *WebhookCreate) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *WebhookCreate) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetOauthUrl

`func (o *WebhookCreate) GetOauthUrl() string`

GetOauthUrl returns the OauthUrl field if non-nil, zero value otherwise.

### GetOauthUrlOk

`func (o *WebhookCreate) GetOauthUrlOk() (*string, bool)`

GetOauthUrlOk returns a tuple with the OauthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthUrl

`func (o *WebhookCreate) SetOauthUrl(v string)`

SetOauthUrl sets OauthUrl field to given value.

### HasOauthUrl

`func (o *WebhookCreate) HasOauthUrl() bool`

HasOauthUrl returns a boolean if a field has been set.

### GetEvent

`func (o *WebhookCreate) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookCreate) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookCreate) SetEvent(v string)`

SetEvent sets Event field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


