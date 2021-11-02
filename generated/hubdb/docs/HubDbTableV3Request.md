# HubDbTableV3Request

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the table | [default to null]
**Label** | **string** | Label of the table | [default to null]
**UseForPages** | **bool** | Specifies whether the table can be used for creation of dynamic pages | [optional] [default to null]
**AllowPublicApiAccess** | **bool** | Specifies whether the table can be read by public without authorization | [optional] [default to null]
**AllowChildTables** | **bool** | Specifies whether child tables can be created | [optional] [default to null]
**EnableChildTablePages** | **bool** | Specifies creation of multi-level dynamic pages using child tables | [optional] [default to null]
**Columns** | [**[]ColumnRequest**](ColumnRequest.md) | List of columns in the table | [optional] [default to null]
**DynamicMetaTags** | **map[string]int32** | Specifies the key value pairs of the metadata fields with the associated column ids | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

