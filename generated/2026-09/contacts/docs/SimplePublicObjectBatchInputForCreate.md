# SimplePublicObjectBatchInputForCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Associations** | [**[]PublicAssociationsForObject**](PublicAssociationsForObject.md) | An array of associations for the CRM objects, where each association specifies the related object IDs and types. | 
**ObjectWriteTraceId** | Pointer to **string** | A string used to trace the write operation for the objects being created. | [optional] 
**Properties** | **map[string]string** | An object containing key-value pairs representing the properties of the CRM objects to be created. Each property is a string. | 

## Methods

### NewSimplePublicObjectBatchInputForCreate

`func NewSimplePublicObjectBatchInputForCreate(associations []PublicAssociationsForObject, properties map[string]string, ) *SimplePublicObjectBatchInputForCreate`

NewSimplePublicObjectBatchInputForCreate instantiates a new SimplePublicObjectBatchInputForCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplePublicObjectBatchInputForCreateWithDefaults

`func NewSimplePublicObjectBatchInputForCreateWithDefaults() *SimplePublicObjectBatchInputForCreate`

NewSimplePublicObjectBatchInputForCreateWithDefaults instantiates a new SimplePublicObjectBatchInputForCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociations

`func (o *SimplePublicObjectBatchInputForCreate) GetAssociations() []PublicAssociationsForObject`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *SimplePublicObjectBatchInputForCreate) GetAssociationsOk() (*[]PublicAssociationsForObject, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *SimplePublicObjectBatchInputForCreate) SetAssociations(v []PublicAssociationsForObject)`

SetAssociations sets Associations field to given value.


### GetObjectWriteTraceId

`func (o *SimplePublicObjectBatchInputForCreate) GetObjectWriteTraceId() string`

GetObjectWriteTraceId returns the ObjectWriteTraceId field if non-nil, zero value otherwise.

### GetObjectWriteTraceIdOk

`func (o *SimplePublicObjectBatchInputForCreate) GetObjectWriteTraceIdOk() (*string, bool)`

GetObjectWriteTraceIdOk returns a tuple with the ObjectWriteTraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectWriteTraceId

`func (o *SimplePublicObjectBatchInputForCreate) SetObjectWriteTraceId(v string)`

SetObjectWriteTraceId sets ObjectWriteTraceId field to given value.

### HasObjectWriteTraceId

`func (o *SimplePublicObjectBatchInputForCreate) HasObjectWriteTraceId() bool`

HasObjectWriteTraceId returns a boolean if a field has been set.

### GetProperties

`func (o *SimplePublicObjectBatchInputForCreate) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SimplePublicObjectBatchInputForCreate) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SimplePublicObjectBatchInputForCreate) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


