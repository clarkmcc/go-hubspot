# ImportRowCore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RowData** | **[]string** |  | 
**LineNumber** | **int32** |  | 
**PageName** | Pointer to **string** |  | [optional] 
**FileId** | **int32** |  | 

## Methods

### NewImportRowCore

`func NewImportRowCore(rowData []string, lineNumber int32, fileId int32, ) *ImportRowCore`

NewImportRowCore instantiates a new ImportRowCore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportRowCoreWithDefaults

`func NewImportRowCoreWithDefaults() *ImportRowCore`

NewImportRowCoreWithDefaults instantiates a new ImportRowCore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRowData

`func (o *ImportRowCore) GetRowData() []string`

GetRowData returns the RowData field if non-nil, zero value otherwise.

### GetRowDataOk

`func (o *ImportRowCore) GetRowDataOk() (*[]string, bool)`

GetRowDataOk returns a tuple with the RowData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowData

`func (o *ImportRowCore) SetRowData(v []string)`

SetRowData sets RowData field to given value.


### GetLineNumber

`func (o *ImportRowCore) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *ImportRowCore) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *ImportRowCore) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.


### GetPageName

`func (o *ImportRowCore) GetPageName() string`

GetPageName returns the PageName field if non-nil, zero value otherwise.

### GetPageNameOk

`func (o *ImportRowCore) GetPageNameOk() (*string, bool)`

GetPageNameOk returns a tuple with the PageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageName

`func (o *ImportRowCore) SetPageName(v string)`

SetPageName sets PageName field to given value.

### HasPageName

`func (o *ImportRowCore) HasPageName() bool`

HasPageName returns a boolean if a field has been set.

### GetFileId

`func (o *ImportRowCore) GetFileId() int32`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *ImportRowCore) GetFileIdOk() (*int32, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *ImportRowCore) SetFileId(v int32)`

SetFileId sets FileId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


