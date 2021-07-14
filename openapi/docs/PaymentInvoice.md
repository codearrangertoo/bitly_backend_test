# PaymentInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalTax** | Pointer to **float32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Charges** | Pointer to [**Charges**](Charges.md) |  | [optional] 
**InvoiceId** | Pointer to **string** |  | [optional] 
**InvoiceDate** | Pointer to **string** |  | [optional] 
**Payments** | Pointer to [**[]Payments**](Payments.md) |  | [optional] 
**InvoiceNumber** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **float32** |  | [optional] 
**Subtotal** | Pointer to **float32** |  | [optional] 
**InvoiceDueDate** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentInvoice

`func NewPaymentInvoice() *PaymentInvoice`

NewPaymentInvoice instantiates a new PaymentInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInvoiceWithDefaults

`func NewPaymentInvoiceWithDefaults() *PaymentInvoice`

NewPaymentInvoiceWithDefaults instantiates a new PaymentInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalTax

`func (o *PaymentInvoice) GetTotalTax() float32`

GetTotalTax returns the TotalTax field if non-nil, zero value otherwise.

### GetTotalTaxOk

`func (o *PaymentInvoice) GetTotalTaxOk() (*float32, bool)`

GetTotalTaxOk returns a tuple with the TotalTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTax

`func (o *PaymentInvoice) SetTotalTax(v float32)`

SetTotalTax sets TotalTax field to given value.

### HasTotalTax

`func (o *PaymentInvoice) HasTotalTax() bool`

HasTotalTax returns a boolean if a field has been set.

### GetDescription

`func (o *PaymentInvoice) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentInvoice) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentInvoice) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentInvoice) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCharges

`func (o *PaymentInvoice) GetCharges() Charges`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *PaymentInvoice) GetChargesOk() (*Charges, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *PaymentInvoice) SetCharges(v Charges)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *PaymentInvoice) HasCharges() bool`

HasCharges returns a boolean if a field has been set.

### GetInvoiceId

`func (o *PaymentInvoice) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *PaymentInvoice) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *PaymentInvoice) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *PaymentInvoice) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetInvoiceDate

`func (o *PaymentInvoice) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *PaymentInvoice) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *PaymentInvoice) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *PaymentInvoice) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### GetPayments

`func (o *PaymentInvoice) GetPayments() []Payments`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *PaymentInvoice) GetPaymentsOk() (*[]Payments, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *PaymentInvoice) SetPayments(v []Payments)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *PaymentInvoice) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *PaymentInvoice) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *PaymentInvoice) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *PaymentInvoice) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *PaymentInvoice) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetTotal

`func (o *PaymentInvoice) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PaymentInvoice) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PaymentInvoice) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PaymentInvoice) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetSubtotal

`func (o *PaymentInvoice) GetSubtotal() float32`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *PaymentInvoice) GetSubtotalOk() (*float32, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *PaymentInvoice) SetSubtotal(v float32)`

SetSubtotal sets Subtotal field to given value.

### HasSubtotal

`func (o *PaymentInvoice) HasSubtotal() bool`

HasSubtotal returns a boolean if a field has been set.

### GetInvoiceDueDate

`func (o *PaymentInvoice) GetInvoiceDueDate() string`

GetInvoiceDueDate returns the InvoiceDueDate field if non-nil, zero value otherwise.

### GetInvoiceDueDateOk

`func (o *PaymentInvoice) GetInvoiceDueDateOk() (*string, bool)`

GetInvoiceDueDateOk returns a tuple with the InvoiceDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDueDate

`func (o *PaymentInvoice) SetInvoiceDueDate(v string)`

SetInvoiceDueDate sets InvoiceDueDate field to given value.

### HasInvoiceDueDate

`func (o *PaymentInvoice) HasInvoiceDueDate() bool`

HasInvoiceDueDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


