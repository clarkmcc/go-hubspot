# TimelineEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTemplateId** | **string** | The event template ID. | [default to null]
**Email** | **string** | The email address used for contact-specific events. This can be used to identify existing contacts, create new ones, or change the email for an existing contact (if paired with the &#x60;objectId&#x60;). | [optional] [default to null]
**ObjectId** | **string** | The CRM object identifier. This is required for every event other than contacts (where utk or email can be used). | [optional] [default to null]
**Utk** | **string** | Use the &#x60;utk&#x60; parameter to associate an event with a contact by &#x60;usertoken&#x60;. This is recommended if you don&#x27;t know a user&#x27;s email, but have an identifying user token in your cookie. | [optional] [default to null]
**Domain** | **string** | The event domain (often paired with utk). | [optional] [default to null]
**Timestamp** | [**time.Time**](time.Time.md) | The time the event occurred. If not passed in, the curren time will be assumed. This is used to determine where an event is shown on a CRM object&#x27;s timeline. | [optional] [default to null]
**Tokens** | **map[string]string** | A collection of token keys and values associated with the template tokens. | [default to null]
**ExtraData** | [***interface{}**](interface{}.md) | Additional event-specific data that can be interpreted by the template&#x27;s markdown. | [optional] [default to null]
**TimelineIFrame** | [***TimelineEventIFrame**](TimelineEventIFrame.md) |  | [optional] [default to null]
**Id** | **string** | Identifier for the event. This is optional, and we recommend you do not pass this in. We will create one for you if you omit this. You can also use &#x60;{{uuid}}&#x60; anywhere in the ID to generate a unique string, guaranteeing uniqueness. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

