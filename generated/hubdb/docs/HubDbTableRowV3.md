# HubDbTableRowV3

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the table row | [optional] [default to null]
**Path** | **string** | Specifies the value for &#x60;hs_path&#x60; column, which will be used as slug in the dynamic pages | [optional] [default to null]
**Name** | **string** | Specifies the value for &#x60;hs_name&#x60; column, which will be used as title in the dynamic pages | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Timestamp at which the row is created | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | Timestamp at which the row is updated last time | [optional] [default to null]
**ChildTableId** | **string** | Specifies the value for the column child table id | [optional] [default to null]
**Values** | [**map[string]interface{}**](interface{}.md) | List of key value pairs with the column name and column value | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

