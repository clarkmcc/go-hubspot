# SimplePublicObjectBatchInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID to be updated. This can be the object ID, or the unique property value of the &#x60;idProperty&#x60; property. | 
**IdProperty** | Pointer to **string** | The name of a property whose values are unique for this object | [optional] 
**ObjectWriteTraceId** | Pointer to **string** | In each input object, set this field to a unique ID value to enable more granular debugging for error responses. Learn more about [multi-status errors](https://developers.hubspot.com/docs/reference/api/other-resources/error-handling#multi-status-errors). | [optional] 
**Properties** | **map[string]string** | The company property values to set. | 

## Methods

### NewSimplePublicObjectBatchInput

`func NewSimplePublicObjectBatchInput(id string, properties map[string]string, ) *SimplePublicObjectBatchInput`

NewSimplePublicObjectBatchInput instantiates a new SimplePublicObjectBatchInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplePublicObjectBatchInputWithDefaults

`func NewSimplePublicObjectBatchInputWithDefaults() *SimplePublicObjectBatchInput`

NewSimplePublicObjectBatchInputWithDefaults instantiates a new SimplePublicObjectBatchInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimplePublicObjectBatchInput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimplePublicObjectBatchInput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimplePublicObjectBatchInput) SetId(v string)`

SetId sets Id field to given value.


### GetIdProperty

`func (o *SimplePublicObjectBatchInput) GetIdProperty() string`

GetIdProperty returns the IdProperty field if non-nil, zero value otherwise.

### GetIdPropertyOk

`func (o *SimplePublicObjectBatchInput) GetIdPropertyOk() (*string, bool)`

GetIdPropertyOk returns a tuple with the IdProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdProperty

`func (o *SimplePublicObjectBatchInput) SetIdProperty(v string)`

SetIdProperty sets IdProperty field to given value.

### HasIdProperty

`func (o *SimplePublicObjectBatchInput) HasIdProperty() bool`

HasIdProperty returns a boolean if a field has been set.

### GetObjectWriteTraceId

`func (o *SimplePublicObjectBatchInput) GetObjectWriteTraceId() string`

GetObjectWriteTraceId returns the ObjectWriteTraceId field if non-nil, zero value otherwise.

### GetObjectWriteTraceIdOk

`func (o *SimplePublicObjectBatchInput) GetObjectWriteTraceIdOk() (*string, bool)`

GetObjectWriteTraceIdOk returns a tuple with the ObjectWriteTraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectWriteTraceId

`func (o *SimplePublicObjectBatchInput) SetObjectWriteTraceId(v string)`

SetObjectWriteTraceId sets ObjectWriteTraceId field to given value.

### HasObjectWriteTraceId

`func (o *SimplePublicObjectBatchInput) HasObjectWriteTraceId() bool`

HasObjectWriteTraceId returns a boolean if a field has been set.

### GetProperties

`func (o *SimplePublicObjectBatchInput) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SimplePublicObjectBatchInput) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SimplePublicObjectBatchInput) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


