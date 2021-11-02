# Column

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the column | [default to null]
**Label** | **string** | Label of the column | [default to null]
**Id** | **string** | Column Id | [optional] [default to null]
**Width** | **int32** | Column width for HubDB UI | [optional] [default to null]
**ForeignTableId** | **int64** | Foreign table id referenced | [optional] [default to null]
**ForeignColumnId** | **int32** | Foreign Column id | [optional] [default to null]
**ForeignIds** | [**[]ForeignId**](ForeignId.md) | Foreign Ids | [optional] [default to null]
**ForeignIdsByName** | [**map[string]ForeignId**](ForeignId.md) | Foreign ids by name | [optional] [default to null]
**ForeignIdsById** | [**map[string]ForeignId**](ForeignId.md) | Foreign ids | [optional] [default to null]
**Type_** | **string** | Type of the column | [default to null]
**OptionCount** | **int32** | Number of options available | [optional] [default to null]
**Archived** | **bool** | Specifies whether the column is archived | [optional] [default to null]
**Options** | [**[]Option**](Option.md) | Options to choose for select and multi-select columns | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

