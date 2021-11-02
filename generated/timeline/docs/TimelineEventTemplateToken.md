# TimelineEventTemplateToken

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Used for list segmentation and reporting. | [default to null]
**ObjectPropertyName** | **string** | The name of the CRM object property. This will populate the CRM object property associated with the event. With enough of these, you can fully build CRM objects via the Timeline API. | [optional] [default to null]
**Options** | [**[]TimelineEventTemplateTokenOption**](TimelineEventTemplateTokenOption.md) | If type is &#x60;enumeration&#x60;, we should have a list of options to choose from. | [default to null]
**Name** | **string** | The name of the token referenced in the templates. This must be unique for the specific template. It may only contain alphanumeric characters, periods, dashes, or underscores (. - _). | [default to null]
**Type_** | **string** | The data type of the token. You can currently choose from [string, number, date, enumeration]. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date and time that the Event Template Token was created, as an ISO 8601 timestamp. Will be null if the template was created before Feb 18th, 2020. | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The date and time that the Event Template Token was last updated, as an ISO 8601 timestamp. Will be null if the template was created before Feb 18th, 2020. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

