# TimelineEventTemplate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The template name. | [default to null]
**HeaderTemplate** | **string** | This uses Markdown syntax with Handlebars and event-specific data to render HTML on a timeline as a header. | [optional] [default to null]
**DetailTemplate** | **string** | This uses Markdown syntax with Handlebars and event-specific data to render HTML on a timeline when you expand the details. | [optional] [default to null]
**Tokens** | [**[]TimelineEventTemplateToken**](TimelineEventTemplateToken.md) | A collection of tokens that can be used as custom properties on the event and to create fully fledged CRM objects. | [default to null]
**Id** | **string** | The template ID. | [default to null]
**ObjectType** | **string** | The type of CRM object this template is for. [Contacts, companies, tickets, and deals] are supported. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date and time that the Event Template was created, as an ISO 8601 timestamp. Will be null if the template was created before Feb 18th, 2020. | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The date and time that the Event Template was last updated, as an ISO 8601 timestamp. Will be null if the template was created before Feb 18th, 2020. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

