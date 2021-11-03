# ResultIdAccountingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Designates if the response is a success (&#39;OK&#39;) or failure (&#39;ERR&#39;). | 
**Id** | **string** | The ID of created entity. | 

## Methods

### NewResultIdAccountingResponse

`func NewResultIdAccountingResponse(result string, id string, ) *ResultIdAccountingResponse`

NewResultIdAccountingResponse instantiates a new ResultIdAccountingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultIdAccountingResponseWithDefaults

`func NewResultIdAccountingResponseWithDefaults() *ResultIdAccountingResponse`

NewResultIdAccountingResponseWithDefaults instantiates a new ResultIdAccountingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ResultIdAccountingResponse) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ResultIdAccountingResponse) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ResultIdAccountingResponse) SetResult(v string)`

SetResult sets Result field to given value.


### GetId

`func (o *ResultIdAccountingResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResultIdAccountingResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResultIdAccountingResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


