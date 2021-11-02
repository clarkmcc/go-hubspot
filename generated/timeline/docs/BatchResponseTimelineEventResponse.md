# BatchResponseTimelineEventResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The status of the batch response. Should always be COMPLETED if processed. | [default to null]
**Results** | [**[]TimelineEventResponse**](TimelineEventResponse.md) | Successfully created events. | [default to null]
**RequestedAt** | [**time.Time**](time.Time.md) | The time the request occurred. | [optional] [default to null]
**StartedAt** | [**time.Time**](time.Time.md) | The time the request began processing. | [default to null]
**CompletedAt** | [**time.Time**](time.Time.md) | The time the request was completed. | [default to null]
**Links** | **map[string]string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

