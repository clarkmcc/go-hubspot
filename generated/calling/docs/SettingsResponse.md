# SettingsResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of your calling service to display to users. | [default to null]
**Url** | **string** | The URL to your phone/calling UI, built with the [Calling SDK](#). | [default to null]
**Height** | **int32** | The target height of the iframe that will contain your phone/calling UI. | [default to null]
**Width** | **int32** | The target width of the iframe that will contain your phone/calling UI. | [default to null]
**IsReady** | **bool** | When true, your service will appear as an option under the *Call* action in contact records of connected accounts. | [default to null]
**SupportsCustomObjects** | **bool** | When true, you are indicating that your service is compatible with engagement v2 service and can be used with custom objects. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | When this calling extension was created. | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The last time the settings for this calling extension were modified. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

