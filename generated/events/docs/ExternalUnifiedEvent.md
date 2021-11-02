# ExternalUnifiedEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** | The objectType for the object which did the event. | [default to null]
**ObjectId** | **string** | The objectId of the object which did the event. | [default to null]
**EventType** | **string** | The format of the &#x60;eventType&#x60; string is &#x60;ae{appId}_{eventTypeLabel}&#x60;, &#x60;pe{portalId}_{eventTypeLabel}&#x60;, or just &#x60;e_{eventTypeLabel}&#x60; for HubSpot events. | [default to null]
**OccurredAt** | [**time.Time**](time.Time.md) | An ISO 8601 timestamp when the event occurred. | [default to null]
**Id** | **string** | A unique identifier for the event. | [default to null]
**Properties** | **map[string]string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

