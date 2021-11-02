# IntegratorCardPayloadResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** | The total number of card properties that will be sent in this response. | [default to null]
**AllItemsLinkUrl** | **string** | URL to a page the integrator has built that displays all details for this card. This URL will be displayed to users under a &#x60;See more [x]&#x60; link if there are more than five items in your response, where &#x60;[x]&#x60; is the value of &#x60;itemLabel&#x60;. | [optional] [default to null]
**CardLabel** | **string** | The label to be used for the &#x60;allItemsLinkUrl&#x60; link (e.g. &#x27;See more tickets&#x27;). If not provided, this falls back to the card&#x27;s title. | [optional] [default to null]
**TopLevelActions** | [***TopLevelActions**](TopLevelActions.md) |  | [optional] [default to null]
**Sections** | [**[]IntegratorObjectResult**](IntegratorObjectResult.md) | A list of up to five valid card sub categories. | [optional] [default to null]
**ResponseVersion** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

