# BillingAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastFourDigits** | Pointer to **string** |  | [optional] 
**SubscriptionPrice** | Pointer to **float32** |  | [optional] 
**BillingInfo** | Pointer to [**BillingInfo**](BillingInfo.md) |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** | ISO timestamp | [optional] 
**ContactInfo** | Pointer to [**ContactInfo**](ContactInfo.md) |  | [optional] 
**AccountNum** | Pointer to **string** |  | [optional] 
**CardType** | Pointer to **string** |  | [optional] 
**SubscriptionStatus** | Pointer to **string** |  | [optional] 
**RenewalDate** | Pointer to **string** | ISO timestamp | [optional] 

## Methods

### NewBillingAccount

`func NewBillingAccount() *BillingAccount`

NewBillingAccount instantiates a new BillingAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingAccountWithDefaults

`func NewBillingAccountWithDefaults() *BillingAccount`

NewBillingAccountWithDefaults instantiates a new BillingAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastFourDigits

`func (o *BillingAccount) GetLastFourDigits() string`

GetLastFourDigits returns the LastFourDigits field if non-nil, zero value otherwise.

### GetLastFourDigitsOk

`func (o *BillingAccount) GetLastFourDigitsOk() (*string, bool)`

GetLastFourDigitsOk returns a tuple with the LastFourDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFourDigits

`func (o *BillingAccount) SetLastFourDigits(v string)`

SetLastFourDigits sets LastFourDigits field to given value.

### HasLastFourDigits

`func (o *BillingAccount) HasLastFourDigits() bool`

HasLastFourDigits returns a boolean if a field has been set.

### GetSubscriptionPrice

`func (o *BillingAccount) GetSubscriptionPrice() float32`

GetSubscriptionPrice returns the SubscriptionPrice field if non-nil, zero value otherwise.

### GetSubscriptionPriceOk

`func (o *BillingAccount) GetSubscriptionPriceOk() (*float32, bool)`

GetSubscriptionPriceOk returns a tuple with the SubscriptionPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPrice

`func (o *BillingAccount) SetSubscriptionPrice(v float32)`

SetSubscriptionPrice sets SubscriptionPrice field to given value.

### HasSubscriptionPrice

`func (o *BillingAccount) HasSubscriptionPrice() bool`

HasSubscriptionPrice returns a boolean if a field has been set.

### GetBillingInfo

`func (o *BillingAccount) GetBillingInfo() BillingInfo`

GetBillingInfo returns the BillingInfo field if non-nil, zero value otherwise.

### GetBillingInfoOk

`func (o *BillingAccount) GetBillingInfoOk() (*BillingInfo, bool)`

GetBillingInfoOk returns a tuple with the BillingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInfo

`func (o *BillingAccount) SetBillingInfo(v BillingInfo)`

SetBillingInfo sets BillingInfo field to given value.

### HasBillingInfo

`func (o *BillingAccount) HasBillingInfo() bool`

HasBillingInfo returns a boolean if a field has been set.

### GetAccountId

`func (o *BillingAccount) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BillingAccount) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BillingAccount) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *BillingAccount) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetEndDate

`func (o *BillingAccount) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BillingAccount) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BillingAccount) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BillingAccount) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetContactInfo

`func (o *BillingAccount) GetContactInfo() ContactInfo`

GetContactInfo returns the ContactInfo field if non-nil, zero value otherwise.

### GetContactInfoOk

`func (o *BillingAccount) GetContactInfoOk() (*ContactInfo, bool)`

GetContactInfoOk returns a tuple with the ContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInfo

`func (o *BillingAccount) SetContactInfo(v ContactInfo)`

SetContactInfo sets ContactInfo field to given value.

### HasContactInfo

`func (o *BillingAccount) HasContactInfo() bool`

HasContactInfo returns a boolean if a field has been set.

### GetAccountNum

`func (o *BillingAccount) GetAccountNum() string`

GetAccountNum returns the AccountNum field if non-nil, zero value otherwise.

### GetAccountNumOk

`func (o *BillingAccount) GetAccountNumOk() (*string, bool)`

GetAccountNumOk returns a tuple with the AccountNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNum

`func (o *BillingAccount) SetAccountNum(v string)`

SetAccountNum sets AccountNum field to given value.

### HasAccountNum

`func (o *BillingAccount) HasAccountNum() bool`

HasAccountNum returns a boolean if a field has been set.

### GetCardType

`func (o *BillingAccount) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *BillingAccount) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *BillingAccount) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *BillingAccount) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetSubscriptionStatus

`func (o *BillingAccount) GetSubscriptionStatus() string`

GetSubscriptionStatus returns the SubscriptionStatus field if non-nil, zero value otherwise.

### GetSubscriptionStatusOk

`func (o *BillingAccount) GetSubscriptionStatusOk() (*string, bool)`

GetSubscriptionStatusOk returns a tuple with the SubscriptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatus

`func (o *BillingAccount) SetSubscriptionStatus(v string)`

SetSubscriptionStatus sets SubscriptionStatus field to given value.

### HasSubscriptionStatus

`func (o *BillingAccount) HasSubscriptionStatus() bool`

HasSubscriptionStatus returns a boolean if a field has been set.

### GetRenewalDate

`func (o *BillingAccount) GetRenewalDate() string`

GetRenewalDate returns the RenewalDate field if non-nil, zero value otherwise.

### GetRenewalDateOk

`func (o *BillingAccount) GetRenewalDateOk() (*string, bool)`

GetRenewalDateOk returns a tuple with the RenewalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalDate

`func (o *BillingAccount) SetRenewalDate(v string)`

SetRenewalDate sets RenewalDate field to given value.

### HasRenewalDate

`func (o *BillingAccount) HasRenewalDate() bool`

HasRenewalDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


