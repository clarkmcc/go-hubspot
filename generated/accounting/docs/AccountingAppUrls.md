# AccountingAppUrls

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GetInvoiceUrl** | **string** | A URL that specifies the endpoint where invoices can be retrieved. | 
**SearchCustomerUrl** | **string** | A URL that specifies the endpoint where a customer search can be performed. | 
**GetInvoicePdfUrl** | **string** | A URL that specifies the endpoint where an invoice PDF can be retrieved. | 
**CustomerUrlTemplate** | **string** | A template URL that indicates the endpoint where a customer can be fetched by ID. Note that ${CUSTOMER_ID} in this URL will be replaced by the actual customer ID. For example: https://myapp.com/api/customers/${CUSTOMER_ID} | 
**ProductUrlTemplate** | **string** | A template URL that indicates the endpoint where a product can be fetched by ID. Note that ${PRODUCT_ID} in this URL will be replaced by the actual product ID. For example: https://myapp.com/api/products/${PRODUCT_ID} | 
**InvoiceUrlTemplate** | **string** | A template URL that indicates the endpoint where an invoice can be fetched by ID. Note that ${INVOICE_ID} in this URL will be replaced by the actual invoice ID. For example: https://myapp.com/api/invoices/${INVOICE_ID} | 
**CreateInvoiceUrl** | Pointer to **string** | A URL that specifies the endpoint where an invoices can be created. | [optional] 
**SearchInvoiceUrl** | Pointer to **string** | A URL that specifies the endpoint where an invoice search can be performed. | [optional] 
**SearchProductUrl** | Pointer to **string** | A URL that specifies the endpoint where a product search can be performed. | [optional] 
**GetTermsUrl** | Pointer to **string** | A URL that specifies the endpoint where payment terms can be retrieved. | [optional] 
**CreateCustomerUrl** | Pointer to **string** | A URL that specifies the endpoint where a new customer can be created. | [optional] 
**SearchTaxUrl** | Pointer to **string** | A URL that specifies the endpoint where a tax search can be performed. | [optional] 
**ExchangeRateUrl** | Pointer to **string** | A URL that specifies the endpoint where exchange rates can be queried. | [optional] 
**SearchUrl** | Pointer to **string** |  | [optional] 
**SearchCountUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountingAppUrls

`func NewAccountingAppUrls(getInvoiceUrl string, searchCustomerUrl string, getInvoicePdfUrl string, customerUrlTemplate string, productUrlTemplate string, invoiceUrlTemplate string, ) *AccountingAppUrls`

NewAccountingAppUrls instantiates a new AccountingAppUrls object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountingAppUrlsWithDefaults

`func NewAccountingAppUrlsWithDefaults() *AccountingAppUrls`

NewAccountingAppUrlsWithDefaults instantiates a new AccountingAppUrls object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGetInvoiceUrl

`func (o *AccountingAppUrls) GetGetInvoiceUrl() string`

GetGetInvoiceUrl returns the GetInvoiceUrl field if non-nil, zero value otherwise.

### GetGetInvoiceUrlOk

`func (o *AccountingAppUrls) GetGetInvoiceUrlOk() (*string, bool)`

GetGetInvoiceUrlOk returns a tuple with the GetInvoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetInvoiceUrl

`func (o *AccountingAppUrls) SetGetInvoiceUrl(v string)`

SetGetInvoiceUrl sets GetInvoiceUrl field to given value.


### GetSearchCustomerUrl

`func (o *AccountingAppUrls) GetSearchCustomerUrl() string`

GetSearchCustomerUrl returns the SearchCustomerUrl field if non-nil, zero value otherwise.

### GetSearchCustomerUrlOk

`func (o *AccountingAppUrls) GetSearchCustomerUrlOk() (*string, bool)`

GetSearchCustomerUrlOk returns a tuple with the SearchCustomerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchCustomerUrl

`func (o *AccountingAppUrls) SetSearchCustomerUrl(v string)`

SetSearchCustomerUrl sets SearchCustomerUrl field to given value.


### GetGetInvoicePdfUrl

`func (o *AccountingAppUrls) GetGetInvoicePdfUrl() string`

GetGetInvoicePdfUrl returns the GetInvoicePdfUrl field if non-nil, zero value otherwise.

### GetGetInvoicePdfUrlOk

`func (o *AccountingAppUrls) GetGetInvoicePdfUrlOk() (*string, bool)`

GetGetInvoicePdfUrlOk returns a tuple with the GetInvoicePdfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetInvoicePdfUrl

`func (o *AccountingAppUrls) SetGetInvoicePdfUrl(v string)`

SetGetInvoicePdfUrl sets GetInvoicePdfUrl field to given value.


### GetCustomerUrlTemplate

`func (o *AccountingAppUrls) GetCustomerUrlTemplate() string`

GetCustomerUrlTemplate returns the CustomerUrlTemplate field if non-nil, zero value otherwise.

### GetCustomerUrlTemplateOk

`func (o *AccountingAppUrls) GetCustomerUrlTemplateOk() (*string, bool)`

GetCustomerUrlTemplateOk returns a tuple with the CustomerUrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUrlTemplate

`func (o *AccountingAppUrls) SetCustomerUrlTemplate(v string)`

SetCustomerUrlTemplate sets CustomerUrlTemplate field to given value.


### GetProductUrlTemplate

`func (o *AccountingAppUrls) GetProductUrlTemplate() string`

GetProductUrlTemplate returns the ProductUrlTemplate field if non-nil, zero value otherwise.

### GetProductUrlTemplateOk

`func (o *AccountingAppUrls) GetProductUrlTemplateOk() (*string, bool)`

GetProductUrlTemplateOk returns a tuple with the ProductUrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductUrlTemplate

`func (o *AccountingAppUrls) SetProductUrlTemplate(v string)`

SetProductUrlTemplate sets ProductUrlTemplate field to given value.


### GetInvoiceUrlTemplate

`func (o *AccountingAppUrls) GetInvoiceUrlTemplate() string`

GetInvoiceUrlTemplate returns the InvoiceUrlTemplate field if non-nil, zero value otherwise.

### GetInvoiceUrlTemplateOk

`func (o *AccountingAppUrls) GetInvoiceUrlTemplateOk() (*string, bool)`

GetInvoiceUrlTemplateOk returns a tuple with the InvoiceUrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceUrlTemplate

`func (o *AccountingAppUrls) SetInvoiceUrlTemplate(v string)`

SetInvoiceUrlTemplate sets InvoiceUrlTemplate field to given value.


### GetCreateInvoiceUrl

`func (o *AccountingAppUrls) GetCreateInvoiceUrl() string`

GetCreateInvoiceUrl returns the CreateInvoiceUrl field if non-nil, zero value otherwise.

### GetCreateInvoiceUrlOk

`func (o *AccountingAppUrls) GetCreateInvoiceUrlOk() (*string, bool)`

GetCreateInvoiceUrlOk returns a tuple with the CreateInvoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateInvoiceUrl

`func (o *AccountingAppUrls) SetCreateInvoiceUrl(v string)`

SetCreateInvoiceUrl sets CreateInvoiceUrl field to given value.

### HasCreateInvoiceUrl

`func (o *AccountingAppUrls) HasCreateInvoiceUrl() bool`

HasCreateInvoiceUrl returns a boolean if a field has been set.

### GetSearchInvoiceUrl

`func (o *AccountingAppUrls) GetSearchInvoiceUrl() string`

GetSearchInvoiceUrl returns the SearchInvoiceUrl field if non-nil, zero value otherwise.

### GetSearchInvoiceUrlOk

`func (o *AccountingAppUrls) GetSearchInvoiceUrlOk() (*string, bool)`

GetSearchInvoiceUrlOk returns a tuple with the SearchInvoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchInvoiceUrl

`func (o *AccountingAppUrls) SetSearchInvoiceUrl(v string)`

SetSearchInvoiceUrl sets SearchInvoiceUrl field to given value.

### HasSearchInvoiceUrl

`func (o *AccountingAppUrls) HasSearchInvoiceUrl() bool`

HasSearchInvoiceUrl returns a boolean if a field has been set.

### GetSearchProductUrl

`func (o *AccountingAppUrls) GetSearchProductUrl() string`

GetSearchProductUrl returns the SearchProductUrl field if non-nil, zero value otherwise.

### GetSearchProductUrlOk

`func (o *AccountingAppUrls) GetSearchProductUrlOk() (*string, bool)`

GetSearchProductUrlOk returns a tuple with the SearchProductUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchProductUrl

`func (o *AccountingAppUrls) SetSearchProductUrl(v string)`

SetSearchProductUrl sets SearchProductUrl field to given value.

### HasSearchProductUrl

`func (o *AccountingAppUrls) HasSearchProductUrl() bool`

HasSearchProductUrl returns a boolean if a field has been set.

### GetGetTermsUrl

`func (o *AccountingAppUrls) GetGetTermsUrl() string`

GetGetTermsUrl returns the GetTermsUrl field if non-nil, zero value otherwise.

### GetGetTermsUrlOk

`func (o *AccountingAppUrls) GetGetTermsUrlOk() (*string, bool)`

GetGetTermsUrlOk returns a tuple with the GetTermsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetTermsUrl

`func (o *AccountingAppUrls) SetGetTermsUrl(v string)`

SetGetTermsUrl sets GetTermsUrl field to given value.

### HasGetTermsUrl

`func (o *AccountingAppUrls) HasGetTermsUrl() bool`

HasGetTermsUrl returns a boolean if a field has been set.

### GetCreateCustomerUrl

`func (o *AccountingAppUrls) GetCreateCustomerUrl() string`

GetCreateCustomerUrl returns the CreateCustomerUrl field if non-nil, zero value otherwise.

### GetCreateCustomerUrlOk

`func (o *AccountingAppUrls) GetCreateCustomerUrlOk() (*string, bool)`

GetCreateCustomerUrlOk returns a tuple with the CreateCustomerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateCustomerUrl

`func (o *AccountingAppUrls) SetCreateCustomerUrl(v string)`

SetCreateCustomerUrl sets CreateCustomerUrl field to given value.

### HasCreateCustomerUrl

`func (o *AccountingAppUrls) HasCreateCustomerUrl() bool`

HasCreateCustomerUrl returns a boolean if a field has been set.

### GetSearchTaxUrl

`func (o *AccountingAppUrls) GetSearchTaxUrl() string`

GetSearchTaxUrl returns the SearchTaxUrl field if non-nil, zero value otherwise.

### GetSearchTaxUrlOk

`func (o *AccountingAppUrls) GetSearchTaxUrlOk() (*string, bool)`

GetSearchTaxUrlOk returns a tuple with the SearchTaxUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTaxUrl

`func (o *AccountingAppUrls) SetSearchTaxUrl(v string)`

SetSearchTaxUrl sets SearchTaxUrl field to given value.

### HasSearchTaxUrl

`func (o *AccountingAppUrls) HasSearchTaxUrl() bool`

HasSearchTaxUrl returns a boolean if a field has been set.

### GetExchangeRateUrl

`func (o *AccountingAppUrls) GetExchangeRateUrl() string`

GetExchangeRateUrl returns the ExchangeRateUrl field if non-nil, zero value otherwise.

### GetExchangeRateUrlOk

`func (o *AccountingAppUrls) GetExchangeRateUrlOk() (*string, bool)`

GetExchangeRateUrlOk returns a tuple with the ExchangeRateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRateUrl

`func (o *AccountingAppUrls) SetExchangeRateUrl(v string)`

SetExchangeRateUrl sets ExchangeRateUrl field to given value.

### HasExchangeRateUrl

`func (o *AccountingAppUrls) HasExchangeRateUrl() bool`

HasExchangeRateUrl returns a boolean if a field has been set.

### GetSearchUrl

`func (o *AccountingAppUrls) GetSearchUrl() string`

GetSearchUrl returns the SearchUrl field if non-nil, zero value otherwise.

### GetSearchUrlOk

`func (o *AccountingAppUrls) GetSearchUrlOk() (*string, bool)`

GetSearchUrlOk returns a tuple with the SearchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchUrl

`func (o *AccountingAppUrls) SetSearchUrl(v string)`

SetSearchUrl sets SearchUrl field to given value.

### HasSearchUrl

`func (o *AccountingAppUrls) HasSearchUrl() bool`

HasSearchUrl returns a boolean if a field has been set.

### GetSearchCountUrl

`func (o *AccountingAppUrls) GetSearchCountUrl() string`

GetSearchCountUrl returns the SearchCountUrl field if non-nil, zero value otherwise.

### GetSearchCountUrlOk

`func (o *AccountingAppUrls) GetSearchCountUrlOk() (*string, bool)`

GetSearchCountUrlOk returns a tuple with the SearchCountUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchCountUrl

`func (o *AccountingAppUrls) SetSearchCountUrl(v string)`

SetSearchCountUrl sets SearchCountUrl field to given value.

### HasSearchCountUrl

`func (o *AccountingAppUrls) HasSearchCountUrl() bool`

HasSearchCountUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


