# CardFetchBodyPatch

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetUrl** | **string** | URL to a service endpoint that will respond with details for this card. HubSpot will call this endpoint each time a user visits a CRM record page where this card should be displayed. | [optional] [default to null]
**ObjectTypes** | [**[]CardObjectTypeBody**](CardObjectTypeBody.md) | An array of CRM object types where this card should be displayed. HubSpot will call your target URL whenever a user visits a record page of the types defined here. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

