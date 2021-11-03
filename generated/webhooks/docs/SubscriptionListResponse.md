# SubscriptionListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]SubscriptionResponse**](SubscriptionResponse.md) | List of event subscriptions for your app | 

## Methods

### NewSubscriptionListResponse

`func NewSubscriptionListResponse(results []SubscriptionResponse, ) *SubscriptionListResponse`

NewSubscriptionListResponse instantiates a new SubscriptionListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionListResponseWithDefaults

`func NewSubscriptionListResponseWithDefaults() *SubscriptionListResponse`

NewSubscriptionListResponseWithDefaults instantiates a new SubscriptionListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *SubscriptionListResponse) GetResults() []SubscriptionResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SubscriptionListResponse) GetResultsOk() (*[]SubscriptionResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SubscriptionListResponse) SetResults(v []SubscriptionResponse)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


