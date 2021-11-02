# UpdatedProduct

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyncAction** | **string** | The operation to be performed. | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The timestamp (ISO8601 format) when the product was updated in the external accounting system. | [default to null]
**Price** | **float64** | The price of the product. | [default to null]
**CurrencyCode** | **string** | The ISO 4217 currency code that represents the currency of the product price. | [optional] [default to null]
**Id** | **string** | The ID of the product in the external accounting system. | [default to null]
**Properties** | **map[string]string** | A map of key-value product properties to be imported. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

