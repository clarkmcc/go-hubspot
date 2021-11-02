# TimelineEventTemplateUpdateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The template name. | [default to null]
**HeaderTemplate** | **string** | This uses Markdown syntax with Handlebars and event-specific data to render HTML on a timeline as a header. | [optional] [default to null]
**DetailTemplate** | **string** | This uses Markdown syntax with Handlebars and event-specific data to render HTML on a timeline when you expand the details. | [optional] [default to null]
**Tokens** | [**[]TimelineEventTemplateToken**](TimelineEventTemplateToken.md) | A collection of tokens that can be used as custom properties on the event and to create fully fledged CRM objects. | [default to null]
**Id** | **string** | The template ID. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

