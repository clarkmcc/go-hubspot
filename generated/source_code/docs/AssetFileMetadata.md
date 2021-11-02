# AssetFileMetadata

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The path of the file in the CMS Developer File System. | [default to null]
**Name** | **string** | The name of the file. | [default to null]
**Folder** | **bool** | Determines whether or not this path points to a folder. | [default to null]
**Children** | **[]string** | If the object is a folder, contains the filenames of the files within the folder. | [optional] [default to null]
**UpdatedAt** | **int32** | Timestamp of when the object was last updated. | [default to null]
**CreatedAt** | **int32** | Timestamp of when the object was first created. | [default to null]
**ArchivedAt** | **int64** | Timestamp of when the object was archived (deleted). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

