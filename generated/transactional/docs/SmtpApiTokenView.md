# SmtpApiTokenView

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | User name to log into the HubSpot SMTP server. | [default to null]
**CreatedBy** | **string** | Email address of the user that sent the token creation request. | [default to null]
**Password** | **string** | Password used to log into the HubSpot SMTP server. | [optional] [default to null]
**EmailCampaignId** | **string** | Identifier assigned to the campaign provided in the token creation request. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Timestamp generated when a token is created. | [default to null]
**CreateContact** | **bool** | Indicates whether a contact should be created for recipients of emails. | [default to null]
**CampaignName** | **string** | A name for the campaign tied to the token. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

