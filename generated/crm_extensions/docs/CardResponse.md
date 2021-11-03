# CardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Generated unique ID for card. | 
**CreatedAt** | Pointer to **time.Time** | When this card was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The last time this card was updated. | [optional] 
**Title** | **string** | Displayed title of this card. | 
**Fetch** | [**CardFetchBody**](CardFetchBody.md) |  | 
**Display** | [**CardDisplayBody**](CardDisplayBody.md) |  | 
**Actions** | [**CardActions**](CardActions.md) |  | 

## Methods

### NewCardResponse

`func NewCardResponse(id string, title string, fetch CardFetchBody, display CardDisplayBody, actions CardActions, ) *CardResponse`

NewCardResponse instantiates a new CardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardResponseWithDefaults

`func NewCardResponseWithDefaults() *CardResponse`

NewCardResponseWithDefaults instantiates a new CardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CardResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *CardResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CardResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CardResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CardResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CardResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CardResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CardResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CardResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTitle

`func (o *CardResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CardResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CardResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetFetch

`func (o *CardResponse) GetFetch() CardFetchBody`

GetFetch returns the Fetch field if non-nil, zero value otherwise.

### GetFetchOk

`func (o *CardResponse) GetFetchOk() (*CardFetchBody, bool)`

GetFetchOk returns a tuple with the Fetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetch

`func (o *CardResponse) SetFetch(v CardFetchBody)`

SetFetch sets Fetch field to given value.


### GetDisplay

`func (o *CardResponse) GetDisplay() CardDisplayBody`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CardResponse) GetDisplayOk() (*CardDisplayBody, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CardResponse) SetDisplay(v CardDisplayBody)`

SetDisplay sets Display field to given value.


### GetActions

`func (o *CardResponse) GetActions() CardActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CardResponse) GetActionsOk() (*CardActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CardResponse) SetActions(v CardActions)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


