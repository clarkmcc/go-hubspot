# AccountingFeatures

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateInvoice** | [***CreateInvoiceFeature**](CreateInvoiceFeature.md) |  | [default to null]
**ImportInvoice** | [***ImportInvoiceFeature**](ImportInvoiceFeature.md) |  | [default to null]
**Sync** | [**map[string]ObjectSyncFeature**](ObjectSyncFeature.md) | Indicates if syncing objects from the external account system into HubSpot is supported for the integration. This is a map, where the key is one of &#x60;CONTACT&#x60; or &#x60;PRODUCT&#x60;, to indicate which type of object you do or don&#x27;t support syncing. For example: &#x60;&#x60;&#x60;   \&quot;sync\&quot;: {     \&quot;CONTACT\&quot;: {       \&quot;toHubSpot\&quot;: true     },     \&quot;PRODUCT\&quot;: {       \&quot;toHubSpot\&quot;: true     }   } &#x60;&#x60;&#x60;  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

