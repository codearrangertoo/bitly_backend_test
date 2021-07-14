# Payments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentNumber** | Pointer to **string** |  | [optional] 
**PaymentDate** | Pointer to **string** |  | [optional] 
**PaymentAmount** | Pointer to **float32** |  | [optional] 

## Methods

### NewPayments

`func NewPayments() *Payments`

NewPayments instantiates a new Payments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentsWithDefaults

`func NewPaymentsWithDefaults() *Payments`

NewPaymentsWithDefaults instantiates a new Payments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentNumber

`func (o *Payments) GetPaymentNumber() string`

GetPaymentNumber returns the PaymentNumber field if non-nil, zero value otherwise.

### GetPaymentNumberOk

`func (o *Payments) GetPaymentNumberOk() (*string, bool)`

GetPaymentNumberOk returns a tuple with the PaymentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentNumber

`func (o *Payments) SetPaymentNumber(v string)`

SetPaymentNumber sets PaymentNumber field to given value.

### HasPaymentNumber

`func (o *Payments) HasPaymentNumber() bool`

HasPaymentNumber returns a boolean if a field has been set.

### GetPaymentDate

`func (o *Payments) GetPaymentDate() string`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *Payments) GetPaymentDateOk() (*string, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *Payments) SetPaymentDate(v string)`

SetPaymentDate sets PaymentDate field to given value.

### HasPaymentDate

`func (o *Payments) HasPaymentDate() bool`

HasPaymentDate returns a boolean if a field has been set.

### GetPaymentAmount

`func (o *Payments) GetPaymentAmount() float32`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *Payments) GetPaymentAmountOk() (*float32, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *Payments) SetPaymentAmount(v float32)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *Payments) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


