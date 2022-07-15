# HubDbTableV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the table | [optional] 
**Name** | **string** | Name of the table | 
**Label** | **string** | Label of the table | 
**Columns** | Pointer to [**[]Column**](Column.md) | List of columns in the table | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**ColumnCount** | Pointer to **int32** | Number of columns including deleted | [optional] 
**RowCount** | Pointer to **int32** | Number of rows in the table | [optional] 
**CreatedBy** | Pointer to [**SimpleUser**](SimpleUser.md) |  | [optional] 
**UpdatedBy** | Pointer to [**SimpleUser**](SimpleUser.md) |  | [optional] 
**PublishedAt** | Pointer to **time.Time** | Timestamp at which the table is published recently | [optional] 
**DynamicMetaTags** | Pointer to **map[string]int32** | Specifies the key value pairs of the metadata fields with the associated column ids | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp at which the table is created | [optional] 
**Archived** | Pointer to **bool** | Specifies whether table is archived or not | [optional] 
**AllowPublicApiAccess** | Pointer to **bool** | Specifies whether the table can be read by public without authorization | [optional] 
**UseForPages** | Pointer to **bool** | Specifies whether the table can be used for creation of dynamic pages | [optional] 
**EnableChildTablePages** | Pointer to **bool** | Specifies creation of multi-level dynamic pages using child tables | [optional] 
**AllowChildTables** | Pointer to **bool** | Specifies whether child tables can be created | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp at which the table is updated recently | [optional] 

## Methods

### NewHubDbTableV3

`func NewHubDbTableV3(name string, label string, ) *HubDbTableV3`

NewHubDbTableV3 instantiates a new HubDbTableV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubDbTableV3WithDefaults

`func NewHubDbTableV3WithDefaults() *HubDbTableV3`

NewHubDbTableV3WithDefaults instantiates a new HubDbTableV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HubDbTableV3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HubDbTableV3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HubDbTableV3) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HubDbTableV3) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *HubDbTableV3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HubDbTableV3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HubDbTableV3) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *HubDbTableV3) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *HubDbTableV3) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *HubDbTableV3) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetColumns

`func (o *HubDbTableV3) GetColumns() []Column`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *HubDbTableV3) GetColumnsOk() (*[]Column, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *HubDbTableV3) SetColumns(v []Column)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *HubDbTableV3) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetPublished

`func (o *HubDbTableV3) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *HubDbTableV3) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *HubDbTableV3) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *HubDbTableV3) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetColumnCount

`func (o *HubDbTableV3) GetColumnCount() int32`

GetColumnCount returns the ColumnCount field if non-nil, zero value otherwise.

### GetColumnCountOk

`func (o *HubDbTableV3) GetColumnCountOk() (*int32, bool)`

GetColumnCountOk returns a tuple with the ColumnCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnCount

`func (o *HubDbTableV3) SetColumnCount(v int32)`

SetColumnCount sets ColumnCount field to given value.

### HasColumnCount

`func (o *HubDbTableV3) HasColumnCount() bool`

HasColumnCount returns a boolean if a field has been set.

### GetRowCount

`func (o *HubDbTableV3) GetRowCount() int32`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *HubDbTableV3) GetRowCountOk() (*int32, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *HubDbTableV3) SetRowCount(v int32)`

SetRowCount sets RowCount field to given value.

### HasRowCount

`func (o *HubDbTableV3) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### GetCreatedBy

`func (o *HubDbTableV3) GetCreatedBy() SimpleUser`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *HubDbTableV3) GetCreatedByOk() (*SimpleUser, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *HubDbTableV3) SetCreatedBy(v SimpleUser)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *HubDbTableV3) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *HubDbTableV3) GetUpdatedBy() SimpleUser`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *HubDbTableV3) GetUpdatedByOk() (*SimpleUser, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *HubDbTableV3) SetUpdatedBy(v SimpleUser)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *HubDbTableV3) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetPublishedAt

`func (o *HubDbTableV3) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *HubDbTableV3) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *HubDbTableV3) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *HubDbTableV3) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetDynamicMetaTags

`func (o *HubDbTableV3) GetDynamicMetaTags() map[string]int32`

GetDynamicMetaTags returns the DynamicMetaTags field if non-nil, zero value otherwise.

### GetDynamicMetaTagsOk

`func (o *HubDbTableV3) GetDynamicMetaTagsOk() (*map[string]int32, bool)`

GetDynamicMetaTagsOk returns a tuple with the DynamicMetaTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicMetaTags

`func (o *HubDbTableV3) SetDynamicMetaTags(v map[string]int32)`

SetDynamicMetaTags sets DynamicMetaTags field to given value.

### HasDynamicMetaTags

`func (o *HubDbTableV3) HasDynamicMetaTags() bool`

HasDynamicMetaTags returns a boolean if a field has been set.

### GetCreatedAt

`func (o *HubDbTableV3) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HubDbTableV3) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HubDbTableV3) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HubDbTableV3) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetArchived

`func (o *HubDbTableV3) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *HubDbTableV3) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *HubDbTableV3) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *HubDbTableV3) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetAllowPublicApiAccess

`func (o *HubDbTableV3) GetAllowPublicApiAccess() bool`

GetAllowPublicApiAccess returns the AllowPublicApiAccess field if non-nil, zero value otherwise.

### GetAllowPublicApiAccessOk

`func (o *HubDbTableV3) GetAllowPublicApiAccessOk() (*bool, bool)`

GetAllowPublicApiAccessOk returns a tuple with the AllowPublicApiAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPublicApiAccess

`func (o *HubDbTableV3) SetAllowPublicApiAccess(v bool)`

SetAllowPublicApiAccess sets AllowPublicApiAccess field to given value.

### HasAllowPublicApiAccess

`func (o *HubDbTableV3) HasAllowPublicApiAccess() bool`

HasAllowPublicApiAccess returns a boolean if a field has been set.

### GetUseForPages

`func (o *HubDbTableV3) GetUseForPages() bool`

GetUseForPages returns the UseForPages field if non-nil, zero value otherwise.

### GetUseForPagesOk

`func (o *HubDbTableV3) GetUseForPagesOk() (*bool, bool)`

GetUseForPagesOk returns a tuple with the UseForPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForPages

`func (o *HubDbTableV3) SetUseForPages(v bool)`

SetUseForPages sets UseForPages field to given value.

### HasUseForPages

`func (o *HubDbTableV3) HasUseForPages() bool`

HasUseForPages returns a boolean if a field has been set.

### GetEnableChildTablePages

`func (o *HubDbTableV3) GetEnableChildTablePages() bool`

GetEnableChildTablePages returns the EnableChildTablePages field if non-nil, zero value otherwise.

### GetEnableChildTablePagesOk

`func (o *HubDbTableV3) GetEnableChildTablePagesOk() (*bool, bool)`

GetEnableChildTablePagesOk returns a tuple with the EnableChildTablePages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableChildTablePages

`func (o *HubDbTableV3) SetEnableChildTablePages(v bool)`

SetEnableChildTablePages sets EnableChildTablePages field to given value.

### HasEnableChildTablePages

`func (o *HubDbTableV3) HasEnableChildTablePages() bool`

HasEnableChildTablePages returns a boolean if a field has been set.

### GetAllowChildTables

`func (o *HubDbTableV3) GetAllowChildTables() bool`

GetAllowChildTables returns the AllowChildTables field if non-nil, zero value otherwise.

### GetAllowChildTablesOk

`func (o *HubDbTableV3) GetAllowChildTablesOk() (*bool, bool)`

GetAllowChildTablesOk returns a tuple with the AllowChildTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowChildTables

`func (o *HubDbTableV3) SetAllowChildTables(v bool)`

SetAllowChildTables sets AllowChildTables field to given value.

### HasAllowChildTables

`func (o *HubDbTableV3) HasAllowChildTables() bool`

HasAllowChildTables returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *HubDbTableV3) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HubDbTableV3) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HubDbTableV3) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HubDbTableV3) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


