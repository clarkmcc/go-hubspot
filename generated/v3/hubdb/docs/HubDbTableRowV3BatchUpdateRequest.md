# HubDbTableRowV3BatchUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | Specifies the value for &#x60;hs_path&#x60; column, which will be used as slug in the dynamic pages | [optional] 
**ChildTableId** | Pointer to **int32** | Specifies the value for the column child table id | [optional] 
**Values** | **map[string]map[string]interface{}** | List of key value pairs with the column name and column value | 
**Name** | Pointer to **string** | Specifies the value for &#x60;hs_name&#x60; column, which will be used as title in the dynamic pages | [optional] 
**Id** | **string** | The id of the table row | 
**DisplayIndex** | Pointer to **int32** |  | [optional] 

## Methods

### NewHubDbTableRowV3BatchUpdateRequest

`func NewHubDbTableRowV3BatchUpdateRequest(values map[string]map[string]interface{}, id string, ) *HubDbTableRowV3BatchUpdateRequest`

NewHubDbTableRowV3BatchUpdateRequest instantiates a new HubDbTableRowV3BatchUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubDbTableRowV3BatchUpdateRequestWithDefaults

`func NewHubDbTableRowV3BatchUpdateRequestWithDefaults() *HubDbTableRowV3BatchUpdateRequest`

NewHubDbTableRowV3BatchUpdateRequestWithDefaults instantiates a new HubDbTableRowV3BatchUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *HubDbTableRowV3BatchUpdateRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HubDbTableRowV3BatchUpdateRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HubDbTableRowV3BatchUpdateRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HubDbTableRowV3BatchUpdateRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetChildTableId

`func (o *HubDbTableRowV3BatchUpdateRequest) GetChildTableId() int32`

GetChildTableId returns the ChildTableId field if non-nil, zero value otherwise.

### GetChildTableIdOk

`func (o *HubDbTableRowV3BatchUpdateRequest) GetChildTableIdOk() (*int32, bool)`

GetChildTableIdOk returns a tuple with the ChildTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildTableId

`func (o *HubDbTableRowV3BatchUpdateRequest) SetChildTableId(v int32)`

SetChildTableId sets ChildTableId field to given value.

### HasChildTableId

`func (o *HubDbTableRowV3BatchUpdateRequest) HasChildTableId() bool`

HasChildTableId returns a boolean if a field has been set.

### GetValues

`func (o *HubDbTableRowV3BatchUpdateRequest) GetValues() map[string]map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *HubDbTableRowV3BatchUpdateRequest) GetValuesOk() (*map[string]map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *HubDbTableRowV3BatchUpdateRequest) SetValues(v map[string]map[string]interface{})`

SetValues sets Values field to given value.


### GetName

`func (o *HubDbTableRowV3BatchUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HubDbTableRowV3BatchUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HubDbTableRowV3BatchUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HubDbTableRowV3BatchUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *HubDbTableRowV3BatchUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HubDbTableRowV3BatchUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HubDbTableRowV3BatchUpdateRequest) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayIndex

`func (o *HubDbTableRowV3BatchUpdateRequest) GetDisplayIndex() int32`

GetDisplayIndex returns the DisplayIndex field if non-nil, zero value otherwise.

### GetDisplayIndexOk

`func (o *HubDbTableRowV3BatchUpdateRequest) GetDisplayIndexOk() (*int32, bool)`

GetDisplayIndexOk returns a tuple with the DisplayIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayIndex

`func (o *HubDbTableRowV3BatchUpdateRequest) SetDisplayIndex(v int32)`

SetDisplayIndex sets DisplayIndex field to given value.

### HasDisplayIndex

`func (o *HubDbTableRowV3BatchUpdateRequest) HasDisplayIndex() bool`

HasDisplayIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


