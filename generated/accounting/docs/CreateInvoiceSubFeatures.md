# CreateInvoiceSubFeatures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateCustomer** | **bool** | Indicates if a new customer can be created in the external accounting system. | 
**Taxes** | **bool** | Indicates if taxes can be specified by the HubSpot user for line items. | 
**ExchangeRates** | **bool** | Indicates if the external accounting system supports fetching exchange rates when create an invoice in HubSpot for a customer billed in a different currency. | 
**Terms** | **bool** | Indicates if the external accounting system supports fetching payment terms. | 
**InvoiceComments** | **bool** | Indicates if a visible comment can be added to invoices. | 
**InvoiceDiscounts** | **bool** | Indicates if invoice-level discounts can be added to invoices. | 

## Methods

### NewCreateInvoiceSubFeatures

`func NewCreateInvoiceSubFeatures(createCustomer bool, taxes bool, exchangeRates bool, terms bool, invoiceComments bool, invoiceDiscounts bool, ) *CreateInvoiceSubFeatures`

NewCreateInvoiceSubFeatures instantiates a new CreateInvoiceSubFeatures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInvoiceSubFeaturesWithDefaults

`func NewCreateInvoiceSubFeaturesWithDefaults() *CreateInvoiceSubFeatures`

NewCreateInvoiceSubFeaturesWithDefaults instantiates a new CreateInvoiceSubFeatures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateCustomer

`func (o *CreateInvoiceSubFeatures) GetCreateCustomer() bool`

GetCreateCustomer returns the CreateCustomer field if non-nil, zero value otherwise.

### GetCreateCustomerOk

`func (o *CreateInvoiceSubFeatures) GetCreateCustomerOk() (*bool, bool)`

GetCreateCustomerOk returns a tuple with the CreateCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateCustomer

`func (o *CreateInvoiceSubFeatures) SetCreateCustomer(v bool)`

SetCreateCustomer sets CreateCustomer field to given value.


### GetTaxes

`func (o *CreateInvoiceSubFeatures) GetTaxes() bool`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *CreateInvoiceSubFeatures) GetTaxesOk() (*bool, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *CreateInvoiceSubFeatures) SetTaxes(v bool)`

SetTaxes sets Taxes field to given value.


### GetExchangeRates

`func (o *CreateInvoiceSubFeatures) GetExchangeRates() bool`

GetExchangeRates returns the ExchangeRates field if non-nil, zero value otherwise.

### GetExchangeRatesOk

`func (o *CreateInvoiceSubFeatures) GetExchangeRatesOk() (*bool, bool)`

GetExchangeRatesOk returns a tuple with the ExchangeRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRates

`func (o *CreateInvoiceSubFeatures) SetExchangeRates(v bool)`

SetExchangeRates sets ExchangeRates field to given value.


### GetTerms

`func (o *CreateInvoiceSubFeatures) GetTerms() bool`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *CreateInvoiceSubFeatures) GetTermsOk() (*bool, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *CreateInvoiceSubFeatures) SetTerms(v bool)`

SetTerms sets Terms field to given value.


### GetInvoiceComments

`func (o *CreateInvoiceSubFeatures) GetInvoiceComments() bool`

GetInvoiceComments returns the InvoiceComments field if non-nil, zero value otherwise.

### GetInvoiceCommentsOk

`func (o *CreateInvoiceSubFeatures) GetInvoiceCommentsOk() (*bool, bool)`

GetInvoiceCommentsOk returns a tuple with the InvoiceComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceComments

`func (o *CreateInvoiceSubFeatures) SetInvoiceComments(v bool)`

SetInvoiceComments sets InvoiceComments field to given value.


### GetInvoiceDiscounts

`func (o *CreateInvoiceSubFeatures) GetInvoiceDiscounts() bool`

GetInvoiceDiscounts returns the InvoiceDiscounts field if non-nil, zero value otherwise.

### GetInvoiceDiscountsOk

`func (o *CreateInvoiceSubFeatures) GetInvoiceDiscountsOk() (*bool, bool)`

GetInvoiceDiscountsOk returns a tuple with the InvoiceDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDiscounts

`func (o *CreateInvoiceSubFeatures) SetInvoiceDiscounts(v bool)`

SetInvoiceDiscounts sets InvoiceDiscounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


