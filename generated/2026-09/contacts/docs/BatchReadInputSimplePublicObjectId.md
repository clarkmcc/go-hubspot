# BatchReadInputSimplePublicObjectId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdProperty** | Pointer to **string** | When using a custom unique value property to retrieve records, the name of the property. Do not include this parameter if retrieving by record ID. | [optional] 
**Inputs** | [**[]SimplePublicObjectId**](SimplePublicObjectId.md) | An array of SimplePublicObjectId objects, each containing an ID of an object to be read. | 
**Properties** | **[]string** | An array of strings representing the properties to be returned for each object. If a specified property is not present on an object, it will be ignored. | 
**PropertiesWithHistory** | **[]string** | An array of strings representing the properties to be returned along with their history of previous values. If a specified property is not present on an object, it will be ignored. | 

## Methods

### NewBatchReadInputSimplePublicObjectId

`func NewBatchReadInputSimplePublicObjectId(inputs []SimplePublicObjectId, properties []string, propertiesWithHistory []string, ) *BatchReadInputSimplePublicObjectId`

NewBatchReadInputSimplePublicObjectId instantiates a new BatchReadInputSimplePublicObjectId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchReadInputSimplePublicObjectIdWithDefaults

`func NewBatchReadInputSimplePublicObjectIdWithDefaults() *BatchReadInputSimplePublicObjectId`

NewBatchReadInputSimplePublicObjectIdWithDefaults instantiates a new BatchReadInputSimplePublicObjectId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdProperty

`func (o *BatchReadInputSimplePublicObjectId) GetIdProperty() string`

GetIdProperty returns the IdProperty field if non-nil, zero value otherwise.

### GetIdPropertyOk

`func (o *BatchReadInputSimplePublicObjectId) GetIdPropertyOk() (*string, bool)`

GetIdPropertyOk returns a tuple with the IdProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdProperty

`func (o *BatchReadInputSimplePublicObjectId) SetIdProperty(v string)`

SetIdProperty sets IdProperty field to given value.

### HasIdProperty

`func (o *BatchReadInputSimplePublicObjectId) HasIdProperty() bool`

HasIdProperty returns a boolean if a field has been set.

### GetInputs

`func (o *BatchReadInputSimplePublicObjectId) GetInputs() []SimplePublicObjectId`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *BatchReadInputSimplePublicObjectId) GetInputsOk() (*[]SimplePublicObjectId, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *BatchReadInputSimplePublicObjectId) SetInputs(v []SimplePublicObjectId)`

SetInputs sets Inputs field to given value.


### GetProperties

`func (o *BatchReadInputSimplePublicObjectId) GetProperties() []string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BatchReadInputSimplePublicObjectId) GetPropertiesOk() (*[]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BatchReadInputSimplePublicObjectId) SetProperties(v []string)`

SetProperties sets Properties field to given value.


### GetPropertiesWithHistory

`func (o *BatchReadInputSimplePublicObjectId) GetPropertiesWithHistory() []string`

GetPropertiesWithHistory returns the PropertiesWithHistory field if non-nil, zero value otherwise.

### GetPropertiesWithHistoryOk

`func (o *BatchReadInputSimplePublicObjectId) GetPropertiesWithHistoryOk() (*[]string, bool)`

GetPropertiesWithHistoryOk returns a tuple with the PropertiesWithHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesWithHistory

`func (o *BatchReadInputSimplePublicObjectId) SetPropertiesWithHistory(v []string)`

SetPropertiesWithHistory sets PropertiesWithHistory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


