# CardListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]CardResponse**](CardResponse.md) | List of card definitions | 

## Methods

### NewCardListResponse

`func NewCardListResponse(results []CardResponse, ) *CardListResponse`

NewCardListResponse instantiates a new CardListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardListResponseWithDefaults

`func NewCardListResponseWithDefaults() *CardListResponse`

NewCardListResponseWithDefaults instantiates a new CardListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *CardListResponse) GetResults() []CardResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CardListResponse) GetResultsOk() (*[]CardResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CardListResponse) SetResults(v []CardResponse)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


