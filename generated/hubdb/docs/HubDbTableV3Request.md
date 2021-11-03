# HubDbTableV3Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the table | 
**Label** | **string** | Label of the table | 
**UseForPages** | Pointer to **bool** | Specifies whether the table can be used for creation of dynamic pages | [optional] 
**AllowPublicApiAccess** | Pointer to **bool** | Specifies whether the table can be read by public without authorization | [optional] 
**AllowChildTables** | Pointer to **bool** | Specifies whether child tables can be created | [optional] 
**EnableChildTablePages** | Pointer to **bool** | Specifies creation of multi-level dynamic pages using child tables | [optional] 
**Columns** | Pointer to [**[]ColumnRequest**](ColumnRequest.md) | List of columns in the table | [optional] 
**DynamicMetaTags** | Pointer to **map[string]int32** | Specifies the key value pairs of the metadata fields with the associated column ids | [optional] 

## Methods

### NewHubDbTableV3Request

`func NewHubDbTableV3Request(name string, label string, ) *HubDbTableV3Request`

NewHubDbTableV3Request instantiates a new HubDbTableV3Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubDbTableV3RequestWithDefaults

`func NewHubDbTableV3RequestWithDefaults() *HubDbTableV3Request`

NewHubDbTableV3RequestWithDefaults instantiates a new HubDbTableV3Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HubDbTableV3Request) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HubDbTableV3Request) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HubDbTableV3Request) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *HubDbTableV3Request) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *HubDbTableV3Request) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *HubDbTableV3Request) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetUseForPages

`func (o *HubDbTableV3Request) GetUseForPages() bool`

GetUseForPages returns the UseForPages field if non-nil, zero value otherwise.

### GetUseForPagesOk

`func (o *HubDbTableV3Request) GetUseForPagesOk() (*bool, bool)`

GetUseForPagesOk returns a tuple with the UseForPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForPages

`func (o *HubDbTableV3Request) SetUseForPages(v bool)`

SetUseForPages sets UseForPages field to given value.

### HasUseForPages

`func (o *HubDbTableV3Request) HasUseForPages() bool`

HasUseForPages returns a boolean if a field has been set.

### GetAllowPublicApiAccess

`func (o *HubDbTableV3Request) GetAllowPublicApiAccess() bool`

GetAllowPublicApiAccess returns the AllowPublicApiAccess field if non-nil, zero value otherwise.

### GetAllowPublicApiAccessOk

`func (o *HubDbTableV3Request) GetAllowPublicApiAccessOk() (*bool, bool)`

GetAllowPublicApiAccessOk returns a tuple with the AllowPublicApiAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPublicApiAccess

`func (o *HubDbTableV3Request) SetAllowPublicApiAccess(v bool)`

SetAllowPublicApiAccess sets AllowPublicApiAccess field to given value.

### HasAllowPublicApiAccess

`func (o *HubDbTableV3Request) HasAllowPublicApiAccess() bool`

HasAllowPublicApiAccess returns a boolean if a field has been set.

### GetAllowChildTables

`func (o *HubDbTableV3Request) GetAllowChildTables() bool`

GetAllowChildTables returns the AllowChildTables field if non-nil, zero value otherwise.

### GetAllowChildTablesOk

`func (o *HubDbTableV3Request) GetAllowChildTablesOk() (*bool, bool)`

GetAllowChildTablesOk returns a tuple with the AllowChildTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowChildTables

`func (o *HubDbTableV3Request) SetAllowChildTables(v bool)`

SetAllowChildTables sets AllowChildTables field to given value.

### HasAllowChildTables

`func (o *HubDbTableV3Request) HasAllowChildTables() bool`

HasAllowChildTables returns a boolean if a field has been set.

### GetEnableChildTablePages

`func (o *HubDbTableV3Request) GetEnableChildTablePages() bool`

GetEnableChildTablePages returns the EnableChildTablePages field if non-nil, zero value otherwise.

### GetEnableChildTablePagesOk

`func (o *HubDbTableV3Request) GetEnableChildTablePagesOk() (*bool, bool)`

GetEnableChildTablePagesOk returns a tuple with the EnableChildTablePages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableChildTablePages

`func (o *HubDbTableV3Request) SetEnableChildTablePages(v bool)`

SetEnableChildTablePages sets EnableChildTablePages field to given value.

### HasEnableChildTablePages

`func (o *HubDbTableV3Request) HasEnableChildTablePages() bool`

HasEnableChildTablePages returns a boolean if a field has been set.

### GetColumns

`func (o *HubDbTableV3Request) GetColumns() []ColumnRequest`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *HubDbTableV3Request) GetColumnsOk() (*[]ColumnRequest, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *HubDbTableV3Request) SetColumns(v []ColumnRequest)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *HubDbTableV3Request) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetDynamicMetaTags

`func (o *HubDbTableV3Request) GetDynamicMetaTags() map[string]int32`

GetDynamicMetaTags returns the DynamicMetaTags field if non-nil, zero value otherwise.

### GetDynamicMetaTagsOk

`func (o *HubDbTableV3Request) GetDynamicMetaTagsOk() (*map[string]int32, bool)`

GetDynamicMetaTagsOk returns a tuple with the DynamicMetaTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicMetaTags

`func (o *HubDbTableV3Request) SetDynamicMetaTags(v map[string]int32)`

SetDynamicMetaTags sets DynamicMetaTags field to given value.

### HasDynamicMetaTags

`func (o *HubDbTableV3Request) HasDynamicMetaTags() bool`

HasDynamicMetaTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


