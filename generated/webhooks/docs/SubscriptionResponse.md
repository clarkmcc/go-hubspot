# SubscriptionResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** | Type of event to listen for. Can be one of &#x60;create&#x60;, &#x60;delete&#x60;, &#x60;deletedForPrivacy&#x60;, or &#x60;propertyChange&#x60;. | [default to null]
**PropertyName** | **string** | The internal name of the property being monitored for changes. Only applies when &#x60;eventType&#x60; is &#x60;propertyChange&#x60;. | [optional] [default to null]
**Active** | **bool** | Determines if the subscription is active or paused. | [default to null]
**Id** | **string** | The unique ID of the subscription. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | When this subscription was created. Formatted as milliseconds from the [Unix epoch](#). | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | When this subscription was last updated. Formatted as milliseconds from the [Unix epoch](#). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

