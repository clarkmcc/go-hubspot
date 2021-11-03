# HubDbTableRowV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the table row | [optional] 
**Path** | Pointer to **string** | Specifies the value for &#x60;hs_path&#x60; column, which will be used as slug in the dynamic pages | [optional] 
**Name** | Pointer to **string** | Specifies the value for &#x60;hs_name&#x60; column, which will be used as title in the dynamic pages | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp at which the row is created | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp at which the row is updated last time | [optional] 
**ChildTableId** | Pointer to **string** | Specifies the value for the column child table id | [optional] 
**Values** | **map[string]map[string]interface{}** | List of key value pairs with the column name and column value | 

## Methods

### NewHubDbTableRowV3

`func NewHubDbTableRowV3(values map[string]map[string]interface{}, ) *HubDbTableRowV3`

NewHubDbTableRowV3 instantiates a new HubDbTableRowV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubDbTableRowV3WithDefaults

`func NewHubDbTableRowV3WithDefaults() *HubDbTableRowV3`

NewHubDbTableRowV3WithDefaults instantiates a new HubDbTableRowV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HubDbTableRowV3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HubDbTableRowV3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HubDbTableRowV3) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HubDbTableRowV3) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *HubDbTableRowV3) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HubDbTableRowV3) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HubDbTableRowV3) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HubDbTableRowV3) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetName

`func (o *HubDbTableRowV3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HubDbTableRowV3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HubDbTableRowV3) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HubDbTableRowV3) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *HubDbTableRowV3) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HubDbTableRowV3) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HubDbTableRowV3) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HubDbTableRowV3) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *HubDbTableRowV3) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HubDbTableRowV3) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HubDbTableRowV3) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HubDbTableRowV3) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetChildTableId

`func (o *HubDbTableRowV3) GetChildTableId() string`

GetChildTableId returns the ChildTableId field if non-nil, zero value otherwise.

### GetChildTableIdOk

`func (o *HubDbTableRowV3) GetChildTableIdOk() (*string, bool)`

GetChildTableIdOk returns a tuple with the ChildTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildTableId

`func (o *HubDbTableRowV3) SetChildTableId(v string)`

SetChildTableId sets ChildTableId field to given value.

### HasChildTableId

`func (o *HubDbTableRowV3) HasChildTableId() bool`

HasChildTableId returns a boolean if a field has been set.

### GetValues

`func (o *HubDbTableRowV3) GetValues() map[string]map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *HubDbTableRowV3) GetValuesOk() (*map[string]map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *HubDbTableRowV3) SetValues(v map[string]map[string]interface{})`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


