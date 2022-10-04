# CardActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseUrls** | **[]string** | A list of URL prefixes that will be accepted for card action URLs. If your data fetch response includes an action URL that doesn&#39;t begin with one of these values, it will result in an error and the card will not be displayed. | 

## Methods

### NewCardActions

`func NewCardActions(baseUrls []string, ) *CardActions`

NewCardActions instantiates a new CardActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardActionsWithDefaults

`func NewCardActionsWithDefaults() *CardActions`

NewCardActionsWithDefaults instantiates a new CardActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseUrls

`func (o *CardActions) GetBaseUrls() []string`

GetBaseUrls returns the BaseUrls field if non-nil, zero value otherwise.

### GetBaseUrlsOk

`func (o *CardActions) GetBaseUrlsOk() (*[]string, bool)`

GetBaseUrlsOk returns a tuple with the BaseUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrls

`func (o *CardActions) SetBaseUrls(v []string)`

SetBaseUrls sets BaseUrls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


