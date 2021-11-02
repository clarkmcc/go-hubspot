# SettingsResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetUrl** | **string** | A publicly available URL for Hubspot to call where event payloads will be delivered. See [link-so-some-doc](#) for details about the format of these event payloads. | [default to null]
**Throttling** | [***ThrottlingSettings**](ThrottlingSettings.md) |  | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | When this subscription was created. Formatted as milliseconds from the [Unix epoch](#). | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | When this subscription was last updated. Formatted as milliseconds from the [Unix epoch](#). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

