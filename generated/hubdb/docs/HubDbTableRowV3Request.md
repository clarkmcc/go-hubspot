# HubDbTableRowV3Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | Specifies the value for &#x60;hs_path&#x60; column, which will be used as slug in the dynamic pages | [optional] 
**Name** | Pointer to **string** | Specifies the value for &#x60;hs_name&#x60; column, which will be used as title in the dynamic pages | [optional] 
**ChildTableId** | Pointer to **int32** | Specifies the value for the column child table id | [optional] 
**Values** | **map[string]map[string]interface{}** | List of key value pairs with the column name and column value | 

## Methods

### NewHubDbTableRowV3Request

`func NewHubDbTableRowV3Request(values map[string]map[string]interface{}, ) *HubDbTableRowV3Request`

NewHubDbTableRowV3Request instantiates a new HubDbTableRowV3Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubDbTableRowV3RequestWithDefaults

`func NewHubDbTableRowV3RequestWithDefaults() *HubDbTableRowV3Request`

NewHubDbTableRowV3RequestWithDefaults instantiates a new HubDbTableRowV3Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *HubDbTableRowV3Request) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HubDbTableRowV3Request) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HubDbTableRowV3Request) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HubDbTableRowV3Request) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetName

`func (o *HubDbTableRowV3Request) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HubDbTableRowV3Request) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HubDbTableRowV3Request) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HubDbTableRowV3Request) HasName() bool`

HasName returns a boolean if a field has been set.

### GetChildTableId

`func (o *HubDbTableRowV3Request) GetChildTableId() int32`

GetChildTableId returns the ChildTableId field if non-nil, zero value otherwise.

### GetChildTableIdOk

`func (o *HubDbTableRowV3Request) GetChildTableIdOk() (*int32, bool)`

GetChildTableIdOk returns a tuple with the ChildTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildTableId

`func (o *HubDbTableRowV3Request) SetChildTableId(v int32)`

SetChildTableId sets ChildTableId field to given value.

### HasChildTableId

`func (o *HubDbTableRowV3Request) HasChildTableId() bool`

HasChildTableId returns a boolean if a field has been set.

### GetValues

`func (o *HubDbTableRowV3Request) GetValues() map[string]map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *HubDbTableRowV3Request) GetValuesOk() (*map[string]map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *HubDbTableRowV3Request) SetValues(v map[string]map[string]interface{})`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


