# ExchangeRateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Designates if the response is a success (&#39;OK&#39;) or failure (&#39;ERR&#39;). | 
**ExchangeRate** | **float32** | The exchange rate between the 2 currencies | 
**SourceCurrencyCode** | **string** | The ISO 4217 currency code that represents the source currency of the exchange rate. | 
**TargetCurrencyCode** | **string** | The ISO 4217 currency code that represents the target currency of the exchange rate. | 

## Methods

### NewExchangeRateResponse

`func NewExchangeRateResponse(result string, exchangeRate float32, sourceCurrencyCode string, targetCurrencyCode string, ) *ExchangeRateResponse`

NewExchangeRateResponse instantiates a new ExchangeRateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeRateResponseWithDefaults

`func NewExchangeRateResponseWithDefaults() *ExchangeRateResponse`

NewExchangeRateResponseWithDefaults instantiates a new ExchangeRateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ExchangeRateResponse) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ExchangeRateResponse) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ExchangeRateResponse) SetResult(v string)`

SetResult sets Result field to given value.


### GetExchangeRate

`func (o *ExchangeRateResponse) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *ExchangeRateResponse) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *ExchangeRateResponse) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.


### GetSourceCurrencyCode

`func (o *ExchangeRateResponse) GetSourceCurrencyCode() string`

GetSourceCurrencyCode returns the SourceCurrencyCode field if non-nil, zero value otherwise.

### GetSourceCurrencyCodeOk

`func (o *ExchangeRateResponse) GetSourceCurrencyCodeOk() (*string, bool)`

GetSourceCurrencyCodeOk returns a tuple with the SourceCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCurrencyCode

`func (o *ExchangeRateResponse) SetSourceCurrencyCode(v string)`

SetSourceCurrencyCode sets SourceCurrencyCode field to given value.


### GetTargetCurrencyCode

`func (o *ExchangeRateResponse) GetTargetCurrencyCode() string`

GetTargetCurrencyCode returns the TargetCurrencyCode field if non-nil, zero value otherwise.

### GetTargetCurrencyCodeOk

`func (o *ExchangeRateResponse) GetTargetCurrencyCodeOk() (*string, bool)`

GetTargetCurrencyCodeOk returns a tuple with the TargetCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCurrencyCode

`func (o *ExchangeRateResponse) SetTargetCurrencyCode(v string)`

SetTargetCurrencyCode sets TargetCurrencyCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


