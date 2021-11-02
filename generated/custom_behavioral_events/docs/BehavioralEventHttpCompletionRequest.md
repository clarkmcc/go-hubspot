# BehavioralEventHttpCompletionRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Utk** | **string** | User token | [optional] [default to null]
**Email** | **string** | Email of visitor | [optional] [default to null]
**EventName** | **string** | Internal name of the event-type to trigger | [default to null]
**Properties** | **map[string]string** | Map of properties for the event in the format property internal name - property value | [default to null]
**OccurredAt** | [**time.Time**](time.Time.md) | The time when this event occurred (if any). If this isn&#x27;t set, the current time will be used | [optional] [default to null]
**ObjectId** | **string** | The object id that this event occurred on. Could be a contact id or a visitor id. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

