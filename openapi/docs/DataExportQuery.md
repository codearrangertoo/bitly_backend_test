# DataExportQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | Pointer to **string** |  | [optional] 
**UnitReference** | Pointer to **int32** |  | [optional] 
**LinkDeeplinks** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**CampaignIds** | Pointer to **[]string** |  | [optional] 
**ModifiedAfter** | Pointer to **int32** |  | [optional] 
**Keyword** | Pointer to **string** |  | [optional] 
**CreatedAfter** | Pointer to **int32** |  | [optional] 
**CustomBitlink** | Pointer to **string** |  | [optional] 
**Units** | Pointer to **int32** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**ReportType** | Pointer to **string** |  | [optional] 
**CreatedBefore** | Pointer to **int32** |  | [optional] 
**Emails** | Pointer to **[]string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 

## Methods

### NewDataExportQuery

`func NewDataExportQuery() *DataExportQuery`

NewDataExportQuery instantiates a new DataExportQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataExportQueryWithDefaults

`func NewDataExportQueryWithDefaults() *DataExportQuery`

NewDataExportQueryWithDefaults instantiates a new DataExportQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *DataExportQuery) GetArchived() string`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *DataExportQuery) GetArchivedOk() (*string, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *DataExportQuery) SetArchived(v string)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *DataExportQuery) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetUnitReference

`func (o *DataExportQuery) GetUnitReference() int32`

GetUnitReference returns the UnitReference field if non-nil, zero value otherwise.

### GetUnitReferenceOk

`func (o *DataExportQuery) GetUnitReferenceOk() (*int32, bool)`

GetUnitReferenceOk returns a tuple with the UnitReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitReference

`func (o *DataExportQuery) SetUnitReference(v int32)`

SetUnitReference sets UnitReference field to given value.

### HasUnitReference

`func (o *DataExportQuery) HasUnitReference() bool`

HasUnitReference returns a boolean if a field has been set.

### GetLinkDeeplinks

`func (o *DataExportQuery) GetLinkDeeplinks() string`

GetLinkDeeplinks returns the LinkDeeplinks field if non-nil, zero value otherwise.

### GetLinkDeeplinksOk

`func (o *DataExportQuery) GetLinkDeeplinksOk() (*string, bool)`

GetLinkDeeplinksOk returns a tuple with the LinkDeeplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDeeplinks

`func (o *DataExportQuery) SetLinkDeeplinks(v string)`

SetLinkDeeplinks sets LinkDeeplinks field to given value.

### HasLinkDeeplinks

`func (o *DataExportQuery) HasLinkDeeplinks() bool`

HasLinkDeeplinks returns a boolean if a field has been set.

### GetTags

`func (o *DataExportQuery) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DataExportQuery) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DataExportQuery) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DataExportQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCampaignIds

`func (o *DataExportQuery) GetCampaignIds() []string`

GetCampaignIds returns the CampaignIds field if non-nil, zero value otherwise.

### GetCampaignIdsOk

`func (o *DataExportQuery) GetCampaignIdsOk() (*[]string, bool)`

GetCampaignIdsOk returns a tuple with the CampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignIds

`func (o *DataExportQuery) SetCampaignIds(v []string)`

SetCampaignIds sets CampaignIds field to given value.

### HasCampaignIds

`func (o *DataExportQuery) HasCampaignIds() bool`

HasCampaignIds returns a boolean if a field has been set.

### GetModifiedAfter

`func (o *DataExportQuery) GetModifiedAfter() int32`

GetModifiedAfter returns the ModifiedAfter field if non-nil, zero value otherwise.

### GetModifiedAfterOk

`func (o *DataExportQuery) GetModifiedAfterOk() (*int32, bool)`

GetModifiedAfterOk returns a tuple with the ModifiedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAfter

`func (o *DataExportQuery) SetModifiedAfter(v int32)`

SetModifiedAfter sets ModifiedAfter field to given value.

### HasModifiedAfter

`func (o *DataExportQuery) HasModifiedAfter() bool`

HasModifiedAfter returns a boolean if a field has been set.

### GetKeyword

`func (o *DataExportQuery) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *DataExportQuery) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *DataExportQuery) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *DataExportQuery) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### GetCreatedAfter

`func (o *DataExportQuery) GetCreatedAfter() int32`

GetCreatedAfter returns the CreatedAfter field if non-nil, zero value otherwise.

### GetCreatedAfterOk

`func (o *DataExportQuery) GetCreatedAfterOk() (*int32, bool)`

GetCreatedAfterOk returns a tuple with the CreatedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAfter

`func (o *DataExportQuery) SetCreatedAfter(v int32)`

SetCreatedAfter sets CreatedAfter field to given value.

### HasCreatedAfter

`func (o *DataExportQuery) HasCreatedAfter() bool`

HasCreatedAfter returns a boolean if a field has been set.

### GetCustomBitlink

`func (o *DataExportQuery) GetCustomBitlink() string`

GetCustomBitlink returns the CustomBitlink field if non-nil, zero value otherwise.

### GetCustomBitlinkOk

`func (o *DataExportQuery) GetCustomBitlinkOk() (*string, bool)`

GetCustomBitlinkOk returns a tuple with the CustomBitlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBitlink

`func (o *DataExportQuery) SetCustomBitlink(v string)`

SetCustomBitlink sets CustomBitlink field to given value.

### HasCustomBitlink

`func (o *DataExportQuery) HasCustomBitlink() bool`

HasCustomBitlink returns a boolean if a field has been set.

### GetUnits

`func (o *DataExportQuery) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *DataExportQuery) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *DataExportQuery) SetUnits(v int32)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *DataExportQuery) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetQuery

`func (o *DataExportQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *DataExportQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *DataExportQuery) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *DataExportQuery) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetReportType

`func (o *DataExportQuery) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *DataExportQuery) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *DataExportQuery) SetReportType(v string)`

SetReportType sets ReportType field to given value.

### HasReportType

`func (o *DataExportQuery) HasReportType() bool`

HasReportType returns a boolean if a field has been set.

### GetCreatedBefore

`func (o *DataExportQuery) GetCreatedBefore() int32`

GetCreatedBefore returns the CreatedBefore field if non-nil, zero value otherwise.

### GetCreatedBeforeOk

`func (o *DataExportQuery) GetCreatedBeforeOk() (*int32, bool)`

GetCreatedBeforeOk returns a tuple with the CreatedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBefore

`func (o *DataExportQuery) SetCreatedBefore(v int32)`

SetCreatedBefore sets CreatedBefore field to given value.

### HasCreatedBefore

`func (o *DataExportQuery) HasCreatedBefore() bool`

HasCreatedBefore returns a boolean if a field has been set.

### GetEmails

`func (o *DataExportQuery) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *DataExportQuery) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *DataExportQuery) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *DataExportQuery) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetUnit

`func (o *DataExportQuery) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *DataExportQuery) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *DataExportQuery) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *DataExportQuery) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


