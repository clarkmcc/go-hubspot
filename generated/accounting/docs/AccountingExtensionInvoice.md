# AccountingExtensionInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountDue** | **float32** | The total amount due. | 
**Balance** | Pointer to **float32** | The remaining outstanding balance due. | [optional] 
**DueDate** | **string** | The due date for payment of the invoice, in ISO-8601 date format (yyyy-MM-dd) | 
**InvoiceNumber** | Pointer to **string** | The invoice number | [optional] 
**CustomerId** | Pointer to **string** | The ID of the customer that this invoice is for. | [optional] 
**Currency** | **string** | The ISO 4217 currency code that represents the currency of this invoice. | 
**InvoiceLink** | **string** | A link to the invoice in the external accounting system. | 
**CustomerName** | **string** | The name of the customer that this invoice is for. | 
**Status** | **string** | The currency status of the invoice. | 

## Methods

### NewAccountingExtensionInvoice

`func NewAccountingExtensionInvoice(amountDue float32, dueDate string, currency string, invoiceLink string, customerName string, status string, ) *AccountingExtensionInvoice`

NewAccountingExtensionInvoice instantiates a new AccountingExtensionInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountingExtensionInvoiceWithDefaults

`func NewAccountingExtensionInvoiceWithDefaults() *AccountingExtensionInvoice`

NewAccountingExtensionInvoiceWithDefaults instantiates a new AccountingExtensionInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountDue

`func (o *AccountingExtensionInvoice) GetAmountDue() float32`

GetAmountDue returns the AmountDue field if non-nil, zero value otherwise.

### GetAmountDueOk

`func (o *AccountingExtensionInvoice) GetAmountDueOk() (*float32, bool)`

GetAmountDueOk returns a tuple with the AmountDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDue

`func (o *AccountingExtensionInvoice) SetAmountDue(v float32)`

SetAmountDue sets AmountDue field to given value.


### GetBalance

`func (o *AccountingExtensionInvoice) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *AccountingExtensionInvoice) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *AccountingExtensionInvoice) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *AccountingExtensionInvoice) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetDueDate

`func (o *AccountingExtensionInvoice) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *AccountingExtensionInvoice) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *AccountingExtensionInvoice) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.


### GetInvoiceNumber

`func (o *AccountingExtensionInvoice) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *AccountingExtensionInvoice) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *AccountingExtensionInvoice) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *AccountingExtensionInvoice) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetCustomerId

`func (o *AccountingExtensionInvoice) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AccountingExtensionInvoice) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AccountingExtensionInvoice) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *AccountingExtensionInvoice) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCurrency

`func (o *AccountingExtensionInvoice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AccountingExtensionInvoice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AccountingExtensionInvoice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetInvoiceLink

`func (o *AccountingExtensionInvoice) GetInvoiceLink() string`

GetInvoiceLink returns the InvoiceLink field if non-nil, zero value otherwise.

### GetInvoiceLinkOk

`func (o *AccountingExtensionInvoice) GetInvoiceLinkOk() (*string, bool)`

GetInvoiceLinkOk returns a tuple with the InvoiceLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceLink

`func (o *AccountingExtensionInvoice) SetInvoiceLink(v string)`

SetInvoiceLink sets InvoiceLink field to given value.


### GetCustomerName

`func (o *AccountingExtensionInvoice) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *AccountingExtensionInvoice) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *AccountingExtensionInvoice) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.


### GetStatus

`func (o *AccountingExtensionInvoice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountingExtensionInvoice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountingExtensionInvoice) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


