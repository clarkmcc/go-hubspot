# ColumnRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Column Id | 
**Name** | **string** | Name of the column | 
**Label** | **string** | Label of the column | 
**Type** | **string** | Type of the column | 
**Options** | [**[]Option**](Option.md) | Options to choose for select and multi-select columns | 
**ForeignTableId** | Pointer to **int64** | The id of another table to which the column refers/points to. | [optional] 
**ForeignColumnId** | Pointer to **int32** | The id of the column from another table to which the column refers/points to. | [optional] 

## Methods

### NewColumnRequest

`func NewColumnRequest(id int32, name string, label string, type_ string, options []Option, ) *ColumnRequest`

NewColumnRequest instantiates a new ColumnRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColumnRequestWithDefaults

`func NewColumnRequestWithDefaults() *ColumnRequest`

NewColumnRequestWithDefaults instantiates a new ColumnRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ColumnRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ColumnRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ColumnRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ColumnRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ColumnRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ColumnRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *ColumnRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ColumnRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ColumnRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *ColumnRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ColumnRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ColumnRequest) SetType(v string)`

SetType sets Type field to given value.


### GetOptions

`func (o *ColumnRequest) GetOptions() []Option`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ColumnRequest) GetOptionsOk() (*[]Option, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ColumnRequest) SetOptions(v []Option)`

SetOptions sets Options field to given value.


### GetForeignTableId

`func (o *ColumnRequest) GetForeignTableId() int64`

GetForeignTableId returns the ForeignTableId field if non-nil, zero value otherwise.

### GetForeignTableIdOk

`func (o *ColumnRequest) GetForeignTableIdOk() (*int64, bool)`

GetForeignTableIdOk returns a tuple with the ForeignTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignTableId

`func (o *ColumnRequest) SetForeignTableId(v int64)`

SetForeignTableId sets ForeignTableId field to given value.

### HasForeignTableId

`func (o *ColumnRequest) HasForeignTableId() bool`

HasForeignTableId returns a boolean if a field has been set.

### GetForeignColumnId

`func (o *ColumnRequest) GetForeignColumnId() int32`

GetForeignColumnId returns the ForeignColumnId field if non-nil, zero value otherwise.

### GetForeignColumnIdOk

`func (o *ColumnRequest) GetForeignColumnIdOk() (*int32, bool)`

GetForeignColumnIdOk returns a tuple with the ForeignColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignColumnId

`func (o *ColumnRequest) SetForeignColumnId(v int32)`

SetForeignColumnId sets ForeignColumnId field to given value.

### HasForeignColumnId

`func (o *ColumnRequest) HasForeignColumnId() bool`

HasForeignColumnId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


