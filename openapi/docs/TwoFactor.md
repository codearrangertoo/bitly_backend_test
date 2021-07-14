# TwoFactor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | **string** |  | 
**Status** | Pointer to **string** |  | [optional] 
**CountryCode** | **string** |  | 

## Methods

### NewTwoFactor

`func NewTwoFactor(phoneNumber string, countryCode string, ) *TwoFactor`

NewTwoFactor instantiates a new TwoFactor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwoFactorWithDefaults

`func NewTwoFactorWithDefaults() *TwoFactor`

NewTwoFactorWithDefaults instantiates a new TwoFactor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *TwoFactor) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *TwoFactor) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *TwoFactor) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetStatus

`func (o *TwoFactor) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TwoFactor) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TwoFactor) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TwoFactor) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCountryCode

`func (o *TwoFactor) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *TwoFactor) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *TwoFactor) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


