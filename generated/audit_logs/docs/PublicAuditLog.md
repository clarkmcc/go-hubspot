# PublicAuditLog

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectId** | **string** | The ID of the object. | [default to null]
**UserId** | **string** | The ID of the user who caused the event. | [default to null]
**Timestamp** | [**time.Time**](time.Time.md) | The timestamp at which the event occurred. | [default to null]
**ObjectName** | **string** | The internal name of the object in HubSpot. | [default to null]
**FullName** | **string** | The name of the user who caused the event. | [default to null]
**Event** | **string** | The type of event that took place (CREATED, UPDATED, PUBLISHED, DELETED, UNPUBLISHED). | [default to null]
**ObjectType** | **string** | The type of the object (BLOG, LANDING_PAGE, DOMAIN, HUBDB_TABLE etc.) | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

