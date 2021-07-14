# PlanLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**References** | Pointer to **map[string]string** |  | [optional] 
**OrganizationGuid** | Pointer to **string** |  | [optional] 
**PlanLimits** | Pointer to [**[]PlanLimit**](PlanLimit.md) |  | [optional] 

## Methods

### NewPlanLimits

`func NewPlanLimits() *PlanLimits`

NewPlanLimits instantiates a new PlanLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanLimitsWithDefaults

`func NewPlanLimitsWithDefaults() *PlanLimits`

NewPlanLimitsWithDefaults instantiates a new PlanLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferences

`func (o *PlanLimits) GetReferences() map[string]string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *PlanLimits) GetReferencesOk() (*map[string]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *PlanLimits) SetReferences(v map[string]string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *PlanLimits) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetOrganizationGuid

`func (o *PlanLimits) GetOrganizationGuid() string`

GetOrganizationGuid returns the OrganizationGuid field if non-nil, zero value otherwise.

### GetOrganizationGuidOk

`func (o *PlanLimits) GetOrganizationGuidOk() (*string, bool)`

GetOrganizationGuidOk returns a tuple with the OrganizationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationGuid

`func (o *PlanLimits) SetOrganizationGuid(v string)`

SetOrganizationGuid sets OrganizationGuid field to given value.

### HasOrganizationGuid

`func (o *PlanLimits) HasOrganizationGuid() bool`

HasOrganizationGuid returns a boolean if a field has been set.

### GetPlanLimits

`func (o *PlanLimits) GetPlanLimits() []PlanLimit`

GetPlanLimits returns the PlanLimits field if non-nil, zero value otherwise.

### GetPlanLimitsOk

`func (o *PlanLimits) GetPlanLimitsOk() (*[]PlanLimit, bool)`

GetPlanLimitsOk returns a tuple with the PlanLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanLimits

`func (o *PlanLimits) SetPlanLimits(v []PlanLimit)`

SetPlanLimits sets PlanLimits field to given value.

### HasPlanLimits

`func (o *PlanLimits) HasPlanLimits() bool`

HasPlanLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


