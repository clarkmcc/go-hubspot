# EmailSendStatusView

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | [***EventIdView**](EventIdView.md) |  | [optional] [default to null]
**StatusId** | **string** | Identifier used to query the status of the send. | [default to null]
**SendResult** | **string** | Result of the send. | [optional] [default to null]
**RequestedAt** | [**time.Time**](time.Time.md) | Time when the send was requested. | [optional] [default to null]
**StartedAt** | [**time.Time**](time.Time.md) | Time when the send began processing. | [optional] [default to null]
**CompletedAt** | [**time.Time**](time.Time.md) | Time when the send was completed. | [optional] [default to null]
**Status** | **string** | Status of the send request. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

