# UpgradeOrgBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingInfo** | Pointer to [**BillingInfo**](BillingInfo.md) |  | [optional] 
**PaymentProvider** | Pointer to **string** |  | [optional] [default to "zuora"]
**StripeToken** | Pointer to **string** |  | [optional] 
**RatePlanName** | Pointer to **string** |  | [optional] 
**OrgGuid** | Pointer to **string** |  | [optional] 
**PaymentMethodId** | Pointer to **string** |  | [optional] 

## Methods

### NewUpgradeOrgBody

`func NewUpgradeOrgBody() *UpgradeOrgBody`

NewUpgradeOrgBody instantiates a new UpgradeOrgBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeOrgBodyWithDefaults

`func NewUpgradeOrgBodyWithDefaults() *UpgradeOrgBody`

NewUpgradeOrgBodyWithDefaults instantiates a new UpgradeOrgBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingInfo

`func (o *UpgradeOrgBody) GetBillingInfo() BillingInfo`

GetBillingInfo returns the BillingInfo field if non-nil, zero value otherwise.

### GetBillingInfoOk

`func (o *UpgradeOrgBody) GetBillingInfoOk() (*BillingInfo, bool)`

GetBillingInfoOk returns a tuple with the BillingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInfo

`func (o *UpgradeOrgBody) SetBillingInfo(v BillingInfo)`

SetBillingInfo sets BillingInfo field to given value.

### HasBillingInfo

`func (o *UpgradeOrgBody) HasBillingInfo() bool`

HasBillingInfo returns a boolean if a field has been set.

### GetPaymentProvider

`func (o *UpgradeOrgBody) GetPaymentProvider() string`

GetPaymentProvider returns the PaymentProvider field if non-nil, zero value otherwise.

### GetPaymentProviderOk

`func (o *UpgradeOrgBody) GetPaymentProviderOk() (*string, bool)`

GetPaymentProviderOk returns a tuple with the PaymentProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentProvider

`func (o *UpgradeOrgBody) SetPaymentProvider(v string)`

SetPaymentProvider sets PaymentProvider field to given value.

### HasPaymentProvider

`func (o *UpgradeOrgBody) HasPaymentProvider() bool`

HasPaymentProvider returns a boolean if a field has been set.

### GetStripeToken

`func (o *UpgradeOrgBody) GetStripeToken() string`

GetStripeToken returns the StripeToken field if non-nil, zero value otherwise.

### GetStripeTokenOk

`func (o *UpgradeOrgBody) GetStripeTokenOk() (*string, bool)`

GetStripeTokenOk returns a tuple with the StripeToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeToken

`func (o *UpgradeOrgBody) SetStripeToken(v string)`

SetStripeToken sets StripeToken field to given value.

### HasStripeToken

`func (o *UpgradeOrgBody) HasStripeToken() bool`

HasStripeToken returns a boolean if a field has been set.

### GetRatePlanName

`func (o *UpgradeOrgBody) GetRatePlanName() string`

GetRatePlanName returns the RatePlanName field if non-nil, zero value otherwise.

### GetRatePlanNameOk

`func (o *UpgradeOrgBody) GetRatePlanNameOk() (*string, bool)`

GetRatePlanNameOk returns a tuple with the RatePlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePlanName

`func (o *UpgradeOrgBody) SetRatePlanName(v string)`

SetRatePlanName sets RatePlanName field to given value.

### HasRatePlanName

`func (o *UpgradeOrgBody) HasRatePlanName() bool`

HasRatePlanName returns a boolean if a field has been set.

### GetOrgGuid

`func (o *UpgradeOrgBody) GetOrgGuid() string`

GetOrgGuid returns the OrgGuid field if non-nil, zero value otherwise.

### GetOrgGuidOk

`func (o *UpgradeOrgBody) GetOrgGuidOk() (*string, bool)`

GetOrgGuidOk returns a tuple with the OrgGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgGuid

`func (o *UpgradeOrgBody) SetOrgGuid(v string)`

SetOrgGuid sets OrgGuid field to given value.

### HasOrgGuid

`func (o *UpgradeOrgBody) HasOrgGuid() bool`

HasOrgGuid returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *UpgradeOrgBody) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *UpgradeOrgBody) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *UpgradeOrgBody) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *UpgradeOrgBody) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


