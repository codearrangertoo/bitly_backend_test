# PaymentInvoiceDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceSummary** | Pointer to [**PaymentInvoice**](PaymentInvoice.md) |  | [optional] 
**BillingAccount** | Pointer to [**BillingAccount**](BillingAccount.md) |  | [optional] 

## Methods

### NewPaymentInvoiceDetail

`func NewPaymentInvoiceDetail() *PaymentInvoiceDetail`

NewPaymentInvoiceDetail instantiates a new PaymentInvoiceDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInvoiceDetailWithDefaults

`func NewPaymentInvoiceDetailWithDefaults() *PaymentInvoiceDetail`

NewPaymentInvoiceDetailWithDefaults instantiates a new PaymentInvoiceDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceSummary

`func (o *PaymentInvoiceDetail) GetInvoiceSummary() PaymentInvoice`

GetInvoiceSummary returns the InvoiceSummary field if non-nil, zero value otherwise.

### GetInvoiceSummaryOk

`func (o *PaymentInvoiceDetail) GetInvoiceSummaryOk() (*PaymentInvoice, bool)`

GetInvoiceSummaryOk returns a tuple with the InvoiceSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceSummary

`func (o *PaymentInvoiceDetail) SetInvoiceSummary(v PaymentInvoice)`

SetInvoiceSummary sets InvoiceSummary field to given value.

### HasInvoiceSummary

`func (o *PaymentInvoiceDetail) HasInvoiceSummary() bool`

HasInvoiceSummary returns a boolean if a field has been set.

### GetBillingAccount

`func (o *PaymentInvoiceDetail) GetBillingAccount() BillingAccount`

GetBillingAccount returns the BillingAccount field if non-nil, zero value otherwise.

### GetBillingAccountOk

`func (o *PaymentInvoiceDetail) GetBillingAccountOk() (*BillingAccount, bool)`

GetBillingAccountOk returns a tuple with the BillingAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAccount

`func (o *PaymentInvoiceDetail) SetBillingAccount(v BillingAccount)`

SetBillingAccount sets BillingAccount field to given value.

### HasBillingAccount

`func (o *PaymentInvoiceDetail) HasBillingAccount() bool`

HasBillingAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


