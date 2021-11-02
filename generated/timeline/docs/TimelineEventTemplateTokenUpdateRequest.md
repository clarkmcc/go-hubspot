# TimelineEventTemplateTokenUpdateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Used for list segmentation and reporting. | [default to null]
**ObjectPropertyName** | **string** | The name of the CRM object property. This will populate the CRM object property associated with the event. With enough of these, you can fully build CRM objects via the Timeline API. | [optional] [default to null]
**Options** | [**[]TimelineEventTemplateTokenOption**](TimelineEventTemplateTokenOption.md) | If type is &#x60;enumeration&#x60;, we should have a list of options to choose from. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

