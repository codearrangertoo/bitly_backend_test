# DeactivateUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feedback** | Pointer to [**Feedback**](Feedback.md) |  | [optional] 
**ConfirmText** | Pointer to **string** |  | [optional] 

## Methods

### NewDeactivateUser

`func NewDeactivateUser() *DeactivateUser`

NewDeactivateUser instantiates a new DeactivateUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeactivateUserWithDefaults

`func NewDeactivateUserWithDefaults() *DeactivateUser`

NewDeactivateUserWithDefaults instantiates a new DeactivateUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeedback

`func (o *DeactivateUser) GetFeedback() Feedback`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *DeactivateUser) GetFeedbackOk() (*Feedback, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *DeactivateUser) SetFeedback(v Feedback)`

SetFeedback sets Feedback field to given value.

### HasFeedback

`func (o *DeactivateUser) HasFeedback() bool`

HasFeedback returns a boolean if a field has been set.

### GetConfirmText

`func (o *DeactivateUser) GetConfirmText() string`

GetConfirmText returns the ConfirmText field if non-nil, zero value otherwise.

### GetConfirmTextOk

`func (o *DeactivateUser) GetConfirmTextOk() (*string, bool)`

GetConfirmTextOk returns a tuple with the ConfirmText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmText

`func (o *DeactivateUser) SetConfirmText(v string)`

SetConfirmText sets ConfirmText field to given value.

### HasConfirmText

`func (o *DeactivateUser) HasConfirmText() bool`

HasConfirmText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


