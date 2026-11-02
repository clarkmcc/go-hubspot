# SimplePublicObjectBatchInputUpsert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A string representing the unique identifier for the CRM object. | 
**IdProperty** | Pointer to **string** | The name of a property whose values are unique for this object | [optional] 
**ObjectWriteTraceId** | Pointer to **string** | A string used to trace the write operation for the object, useful for debugging or tracking changes. | [optional] 
**Properties** | **map[string]string** | An object containing key-value pairs where keys are property names and values are property values for the CRM object. Each property is a string. | 

## Methods

### NewSimplePublicObjectBatchInputUpsert

`func NewSimplePublicObjectBatchInputUpsert(id string, properties map[string]string, ) *SimplePublicObjectBatchInputUpsert`

NewSimplePublicObjectBatchInputUpsert instantiates a new SimplePublicObjectBatchInputUpsert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplePublicObjectBatchInputUpsertWithDefaults

`func NewSimplePublicObjectBatchInputUpsertWithDefaults() *SimplePublicObjectBatchInputUpsert`

NewSimplePublicObjectBatchInputUpsertWithDefaults instantiates a new SimplePublicObjectBatchInputUpsert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimplePublicObjectBatchInputUpsert) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimplePublicObjectBatchInputUpsert) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimplePublicObjectBatchInputUpsert) SetId(v string)`

SetId sets Id field to given value.


### GetIdProperty

`func (o *SimplePublicObjectBatchInputUpsert) GetIdProperty() string`

GetIdProperty returns the IdProperty field if non-nil, zero value otherwise.

### GetIdPropertyOk

`func (o *SimplePublicObjectBatchInputUpsert) GetIdPropertyOk() (*string, bool)`

GetIdPropertyOk returns a tuple with the IdProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdProperty

`func (o *SimplePublicObjectBatchInputUpsert) SetIdProperty(v string)`

SetIdProperty sets IdProperty field to given value.

### HasIdProperty

`func (o *SimplePublicObjectBatchInputUpsert) HasIdProperty() bool`

HasIdProperty returns a boolean if a field has been set.

### GetObjectWriteTraceId

`func (o *SimplePublicObjectBatchInputUpsert) GetObjectWriteTraceId() string`

GetObjectWriteTraceId returns the ObjectWriteTraceId field if non-nil, zero value otherwise.

### GetObjectWriteTraceIdOk

`func (o *SimplePublicObjectBatchInputUpsert) GetObjectWriteTraceIdOk() (*string, bool)`

GetObjectWriteTraceIdOk returns a tuple with the ObjectWriteTraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectWriteTraceId

`func (o *SimplePublicObjectBatchInputUpsert) SetObjectWriteTraceId(v string)`

SetObjectWriteTraceId sets ObjectWriteTraceId field to given value.

### HasObjectWriteTraceId

`func (o *SimplePublicObjectBatchInputUpsert) HasObjectWriteTraceId() bool`

HasObjectWriteTraceId returns a boolean if a field has been set.

### GetProperties

`func (o *SimplePublicObjectBatchInputUpsert) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SimplePublicObjectBatchInputUpsert) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SimplePublicObjectBatchInputUpsert) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


