# CardPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The top-level title for this card. Displayed to users in the CRM UI. | [optional] 
**Fetch** | Pointer to [**CardFetchBodyPatch**](CardFetchBodyPatch.md) |  | [optional] 
**Display** | Pointer to [**CardDisplayBody**](CardDisplayBody.md) |  | [optional] 
**Actions** | Pointer to [**CardActions**](CardActions.md) |  | [optional] 

## Methods

### NewCardPatchRequest

`func NewCardPatchRequest() *CardPatchRequest`

NewCardPatchRequest instantiates a new CardPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardPatchRequestWithDefaults

`func NewCardPatchRequestWithDefaults() *CardPatchRequest`

NewCardPatchRequestWithDefaults instantiates a new CardPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CardPatchRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CardPatchRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CardPatchRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CardPatchRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFetch

`func (o *CardPatchRequest) GetFetch() CardFetchBodyPatch`

GetFetch returns the Fetch field if non-nil, zero value otherwise.

### GetFetchOk

`func (o *CardPatchRequest) GetFetchOk() (*CardFetchBodyPatch, bool)`

GetFetchOk returns a tuple with the Fetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetch

`func (o *CardPatchRequest) SetFetch(v CardFetchBodyPatch)`

SetFetch sets Fetch field to given value.

### HasFetch

`func (o *CardPatchRequest) HasFetch() bool`

HasFetch returns a boolean if a field has been set.

### GetDisplay

`func (o *CardPatchRequest) GetDisplay() CardDisplayBody`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CardPatchRequest) GetDisplayOk() (*CardDisplayBody, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CardPatchRequest) SetDisplay(v CardDisplayBody)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *CardPatchRequest) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetActions

`func (o *CardPatchRequest) GetActions() CardActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CardPatchRequest) GetActionsOk() (*CardActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CardPatchRequest) SetActions(v CardActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *CardPatchRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


