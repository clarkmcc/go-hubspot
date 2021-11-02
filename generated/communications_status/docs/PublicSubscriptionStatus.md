# PublicSubscriptionStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID for the subscription. | [default to null]
**Name** | **string** | The name of the subscription. | [default to null]
**Description** | **string** | A description of the subscription. | [default to null]
**Status** | **string** | Whether the contact is subscribed. | [default to null]
**SourceOfStatus** | **string** | Where the status is determined from e.g. PORTAL_WIDE_STATUS if the contact opted out from the portal. | [default to null]
**BrandId** | **int64** | The ID of the brand that the subscription is associated with, if there is one. | [optional] [default to null]
**PreferenceGroupName** | **string** | The name of the preferences group that the subscription is associated with. | [optional] [default to null]
**LegalBasis** | **string** | The legal reason for the current status of the subscription. | [optional] [default to null]
**LegalBasisExplanation** | **string** | A more detailed explanation to go with the legal basis. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

