# PublicSingleSendEmail

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **string** | The From header for the email. | [optional] [default to null]
**To** | **string** | The recipient of the email. | [optional] [default to null]
**SendId** | **string** | ID for a particular send. No more than one email will be sent per sendId. | [optional] [default to null]
**ReplyTo** | **[]string** | List of Reply-To header values for the email. | [default to null]
**Cc** | **[]string** | List of email addresses to send as Cc. | [default to null]
**Bcc** | **[]string** | List of email addresses to send as Bcc. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

