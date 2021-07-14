# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UiAction** | **string** |  | 
**Referrer** | Pointer to **string** |  | [optional] 
**Action** | **string** |  | 
**InitiatedBy** | **string** |  | 
**GroupGuid** | Pointer to **string** |  | [optional] 
**UiActionDate** | **string** |  | 
**Login** | **string** |  | 
**OrgGuid** | **string** |  | 

## Methods

### NewEvent

`func NewEvent(uiAction string, action string, initiatedBy string, uiActionDate string, login string, orgGuid string, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUiAction

`func (o *Event) GetUiAction() string`

GetUiAction returns the UiAction field if non-nil, zero value otherwise.

### GetUiActionOk

`func (o *Event) GetUiActionOk() (*string, bool)`

GetUiActionOk returns a tuple with the UiAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiAction

`func (o *Event) SetUiAction(v string)`

SetUiAction sets UiAction field to given value.


### GetReferrer

`func (o *Event) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *Event) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *Event) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.

### HasReferrer

`func (o *Event) HasReferrer() bool`

HasReferrer returns a boolean if a field has been set.

### GetAction

`func (o *Event) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Event) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Event) SetAction(v string)`

SetAction sets Action field to given value.


### GetInitiatedBy

`func (o *Event) GetInitiatedBy() string`

GetInitiatedBy returns the InitiatedBy field if non-nil, zero value otherwise.

### GetInitiatedByOk

`func (o *Event) GetInitiatedByOk() (*string, bool)`

GetInitiatedByOk returns a tuple with the InitiatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatedBy

`func (o *Event) SetInitiatedBy(v string)`

SetInitiatedBy sets InitiatedBy field to given value.


### GetGroupGuid

`func (o *Event) GetGroupGuid() string`

GetGroupGuid returns the GroupGuid field if non-nil, zero value otherwise.

### GetGroupGuidOk

`func (o *Event) GetGroupGuidOk() (*string, bool)`

GetGroupGuidOk returns a tuple with the GroupGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupGuid

`func (o *Event) SetGroupGuid(v string)`

SetGroupGuid sets GroupGuid field to given value.

### HasGroupGuid

`func (o *Event) HasGroupGuid() bool`

HasGroupGuid returns a boolean if a field has been set.

### GetUiActionDate

`func (o *Event) GetUiActionDate() string`

GetUiActionDate returns the UiActionDate field if non-nil, zero value otherwise.

### GetUiActionDateOk

`func (o *Event) GetUiActionDateOk() (*string, bool)`

GetUiActionDateOk returns a tuple with the UiActionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiActionDate

`func (o *Event) SetUiActionDate(v string)`

SetUiActionDate sets UiActionDate field to given value.


### GetLogin

`func (o *Event) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *Event) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *Event) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetOrgGuid

`func (o *Event) GetOrgGuid() string`

GetOrgGuid returns the OrgGuid field if non-nil, zero value otherwise.

### GetOrgGuidOk

`func (o *Event) GetOrgGuidOk() (*string, bool)`

GetOrgGuidOk returns a tuple with the OrgGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgGuid

`func (o *Event) SetOrgGuid(v string)`

SetOrgGuid sets OrgGuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


