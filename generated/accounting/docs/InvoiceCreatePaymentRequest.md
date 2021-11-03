# InvoiceCreatePaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountPaid** | **float32** | The amount that this payment is for. | 
**CurrencyCode** | **string** | The ISO 4217 currency code that represents the currency of the payment. | 
**PaymentDateTime** | **time.Time** | The datetime that this payment was received. | 
**ExternalPaymentId** | **string** | The id of this payment in the external accounting system. | 

## Methods

### NewInvoiceCreatePaymentRequest

`func NewInvoiceCreatePaymentRequest(amountPaid float32, currencyCode string, paymentDateTime time.Time, externalPaymentId string, ) *InvoiceCreatePaymentRequest`

NewInvoiceCreatePaymentRequest instantiates a new InvoiceCreatePaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceCreatePaymentRequestWithDefaults

`func NewInvoiceCreatePaymentRequestWithDefaults() *InvoiceCreatePaymentRequest`

NewInvoiceCreatePaymentRequestWithDefaults instantiates a new InvoiceCreatePaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountPaid

`func (o *InvoiceCreatePaymentRequest) GetAmountPaid() float32`

GetAmountPaid returns the AmountPaid field if non-nil, zero value otherwise.

### GetAmountPaidOk

`func (o *InvoiceCreatePaymentRequest) GetAmountPaidOk() (*float32, bool)`

GetAmountPaidOk returns a tuple with the AmountPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaid

`func (o *InvoiceCreatePaymentRequest) SetAmountPaid(v float32)`

SetAmountPaid sets AmountPaid field to given value.


### GetCurrencyCode

`func (o *InvoiceCreatePaymentRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *InvoiceCreatePaymentRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *InvoiceCreatePaymentRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetPaymentDateTime

`func (o *InvoiceCreatePaymentRequest) GetPaymentDateTime() time.Time`

GetPaymentDateTime returns the PaymentDateTime field if non-nil, zero value otherwise.

### GetPaymentDateTimeOk

`func (o *InvoiceCreatePaymentRequest) GetPaymentDateTimeOk() (*time.Time, bool)`

GetPaymentDateTimeOk returns a tuple with the PaymentDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDateTime

`func (o *InvoiceCreatePaymentRequest) SetPaymentDateTime(v time.Time)`

SetPaymentDateTime sets PaymentDateTime field to given value.


### GetExternalPaymentId

`func (o *InvoiceCreatePaymentRequest) GetExternalPaymentId() string`

GetExternalPaymentId returns the ExternalPaymentId field if non-nil, zero value otherwise.

### GetExternalPaymentIdOk

`func (o *InvoiceCreatePaymentRequest) GetExternalPaymentIdOk() (*string, bool)`

GetExternalPaymentIdOk returns a tuple with the ExternalPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPaymentId

`func (o *InvoiceCreatePaymentRequest) SetExternalPaymentId(v string)`

SetExternalPaymentId sets ExternalPaymentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


