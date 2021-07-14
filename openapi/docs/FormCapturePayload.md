# FormCapturePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubmitText** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ResponseStatus** | Pointer to **string** |  | [optional] 
**Pages** | Pointer to [**[]FormPage**](FormPage.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**GroupGuid** | Pointer to **string** |  | [optional] 
**Login** | Pointer to **string** |  | [optional] 
**SkipText** | Pointer to **string** |  | [optional] 
**OrgGuid** | Pointer to **string** |  | [optional] 

## Methods

### NewFormCapturePayload

`func NewFormCapturePayload() *FormCapturePayload`

NewFormCapturePayload instantiates a new FormCapturePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormCapturePayloadWithDefaults

`func NewFormCapturePayloadWithDefaults() *FormCapturePayload`

NewFormCapturePayloadWithDefaults instantiates a new FormCapturePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubmitText

`func (o *FormCapturePayload) GetSubmitText() string`

GetSubmitText returns the SubmitText field if non-nil, zero value otherwise.

### GetSubmitTextOk

`func (o *FormCapturePayload) GetSubmitTextOk() (*string, bool)`

GetSubmitTextOk returns a tuple with the SubmitText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitText

`func (o *FormCapturePayload) SetSubmitText(v string)`

SetSubmitText sets SubmitText field to given value.

### HasSubmitText

`func (o *FormCapturePayload) HasSubmitText() bool`

HasSubmitText returns a boolean if a field has been set.

### GetName

`func (o *FormCapturePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormCapturePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormCapturePayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FormCapturePayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResponseStatus

`func (o *FormCapturePayload) GetResponseStatus() string`

GetResponseStatus returns the ResponseStatus field if non-nil, zero value otherwise.

### GetResponseStatusOk

`func (o *FormCapturePayload) GetResponseStatusOk() (*string, bool)`

GetResponseStatusOk returns a tuple with the ResponseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatus

`func (o *FormCapturePayload) SetResponseStatus(v string)`

SetResponseStatus sets ResponseStatus field to given value.

### HasResponseStatus

`func (o *FormCapturePayload) HasResponseStatus() bool`

HasResponseStatus returns a boolean if a field has been set.

### GetPages

`func (o *FormCapturePayload) GetPages() []FormPage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *FormCapturePayload) GetPagesOk() (*[]FormPage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *FormCapturePayload) SetPages(v []FormPage)`

SetPages sets Pages field to given value.

### HasPages

`func (o *FormCapturePayload) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetVersion

`func (o *FormCapturePayload) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FormCapturePayload) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FormCapturePayload) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FormCapturePayload) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetGroupGuid

`func (o *FormCapturePayload) GetGroupGuid() string`

GetGroupGuid returns the GroupGuid field if non-nil, zero value otherwise.

### GetGroupGuidOk

`func (o *FormCapturePayload) GetGroupGuidOk() (*string, bool)`

GetGroupGuidOk returns a tuple with the GroupGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupGuid

`func (o *FormCapturePayload) SetGroupGuid(v string)`

SetGroupGuid sets GroupGuid field to given value.

### HasGroupGuid

`func (o *FormCapturePayload) HasGroupGuid() bool`

HasGroupGuid returns a boolean if a field has been set.

### GetLogin

`func (o *FormCapturePayload) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *FormCapturePayload) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *FormCapturePayload) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *FormCapturePayload) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetSkipText

`func (o *FormCapturePayload) GetSkipText() string`

GetSkipText returns the SkipText field if non-nil, zero value otherwise.

### GetSkipTextOk

`func (o *FormCapturePayload) GetSkipTextOk() (*string, bool)`

GetSkipTextOk returns a tuple with the SkipText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipText

`func (o *FormCapturePayload) SetSkipText(v string)`

SetSkipText sets SkipText field to given value.

### HasSkipText

`func (o *FormCapturePayload) HasSkipText() bool`

HasSkipText returns a boolean if a field has been set.

### GetOrgGuid

`func (o *FormCapturePayload) GetOrgGuid() string`

GetOrgGuid returns the OrgGuid field if non-nil, zero value otherwise.

### GetOrgGuidOk

`func (o *FormCapturePayload) GetOrgGuidOk() (*string, bool)`

GetOrgGuidOk returns a tuple with the OrgGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgGuid

`func (o *FormCapturePayload) SetOrgGuid(v string)`

SetOrgGuid sets OrgGuid field to given value.

### HasOrgGuid

`func (o *FormCapturePayload) HasOrgGuid() bool`

HasOrgGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


