# CardCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | The top-level title for this card. Displayed to users in the CRM UI. | 
**Fetch** | [**CardFetchBody**](CardFetchBody.md) |  | 
**Display** | [**CardDisplayBody**](CardDisplayBody.md) |  | 
**Actions** | [**CardActions**](CardActions.md) |  | 

## Methods

### NewCardCreateRequest

`func NewCardCreateRequest(title string, fetch CardFetchBody, display CardDisplayBody, actions CardActions, ) *CardCreateRequest`

NewCardCreateRequest instantiates a new CardCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardCreateRequestWithDefaults

`func NewCardCreateRequestWithDefaults() *CardCreateRequest`

NewCardCreateRequestWithDefaults instantiates a new CardCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CardCreateRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CardCreateRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CardCreateRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetFetch

`func (o *CardCreateRequest) GetFetch() CardFetchBody`

GetFetch returns the Fetch field if non-nil, zero value otherwise.

### GetFetchOk

`func (o *CardCreateRequest) GetFetchOk() (*CardFetchBody, bool)`

GetFetchOk returns a tuple with the Fetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetch

`func (o *CardCreateRequest) SetFetch(v CardFetchBody)`

SetFetch sets Fetch field to given value.


### GetDisplay

`func (o *CardCreateRequest) GetDisplay() CardDisplayBody`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CardCreateRequest) GetDisplayOk() (*CardDisplayBody, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CardCreateRequest) SetDisplay(v CardDisplayBody)`

SetDisplay sets Display field to given value.


### GetActions

`func (o *CardCreateRequest) GetActions() CardActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CardCreateRequest) GetActionsOk() (*CardActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CardCreateRequest) SetActions(v CardActions)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


