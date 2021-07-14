# PlanLimitsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationGuid** | Pointer to **string** |  | [optional] 
**PlanLimits** | Pointer to [**[]PlanLimit**](PlanLimit.md) |  | [optional] 

## Methods

### NewPlanLimitsAllOf

`func NewPlanLimitsAllOf() *PlanLimitsAllOf`

NewPlanLimitsAllOf instantiates a new PlanLimitsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanLimitsAllOfWithDefaults

`func NewPlanLimitsAllOfWithDefaults() *PlanLimitsAllOf`

NewPlanLimitsAllOfWithDefaults instantiates a new PlanLimitsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationGuid

`func (o *PlanLimitsAllOf) GetOrganizationGuid() string`

GetOrganizationGuid returns the OrganizationGuid field if non-nil, zero value otherwise.

### GetOrganizationGuidOk

`func (o *PlanLimitsAllOf) GetOrganizationGuidOk() (*string, bool)`

GetOrganizationGuidOk returns a tuple with the OrganizationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationGuid

`func (o *PlanLimitsAllOf) SetOrganizationGuid(v string)`

SetOrganizationGuid sets OrganizationGuid field to given value.

### HasOrganizationGuid

`func (o *PlanLimitsAllOf) HasOrganizationGuid() bool`

HasOrganizationGuid returns a boolean if a field has been set.

### GetPlanLimits

`func (o *PlanLimitsAllOf) GetPlanLimits() []PlanLimit`

GetPlanLimits returns the PlanLimits field if non-nil, zero value otherwise.

### GetPlanLimitsOk

`func (o *PlanLimitsAllOf) GetPlanLimitsOk() (*[]PlanLimit, bool)`

GetPlanLimitsOk returns a tuple with the PlanLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanLimits

`func (o *PlanLimitsAllOf) SetPlanLimits(v []PlanLimit)`

SetPlanLimits sets PlanLimits field to given value.

### HasPlanLimits

`func (o *PlanLimitsAllOf) HasPlanLimits() bool`

HasPlanLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


