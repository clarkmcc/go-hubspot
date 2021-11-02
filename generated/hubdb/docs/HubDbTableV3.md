# HubDbTableV3

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the table | [optional] [default to null]
**Name** | **string** | Name of the table | [default to null]
**Label** | **string** | Label of the table | [default to null]
**Columns** | [**[]Column**](Column.md) | List of columns in the table | [optional] [default to null]
**Published** | **bool** |  | [optional] [default to null]
**RowCount** | **int32** | Number of rows in the table | [optional] [default to null]
**CreatedBy** | [***SimpleUser**](SimpleUser.md) |  | [optional] [default to null]
**UpdatedBy** | [***SimpleUser**](SimpleUser.md) |  | [optional] [default to null]
**PublishedAt** | [**time.Time**](time.Time.md) | Timestamp at which the table is published recently | [optional] [default to null]
**DynamicMetaTags** | **map[string]int32** | Specifies the key value pairs of the metadata fields with the associated column ids | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Timestamp at which the table is created | [optional] [default to null]
**Archived** | **bool** | Specifies whether table is archived or not | [optional] [default to null]
**AllowPublicApiAccess** | **bool** | Specifies whether the table can be read by public without authorization | [optional] [default to null]
**UseForPages** | **bool** | Specifies whether the table can be used for creation of dynamic pages | [optional] [default to null]
**EnableChildTablePages** | **bool** | Specifies creation of multi-level dynamic pages using child tables | [optional] [default to null]
**ColumnCount** | **int32** | Number of columns including deleted | [optional] [default to null]
**AllowChildTables** | **bool** | Specifies whether child tables can be created | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | Timestamp at which the table is updated recently | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

