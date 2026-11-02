# BatchReadInputPropertyName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | **bool** | Specifies whether the property is archived. | 
**DataSensitivity** | **string** | Indicates the sensitivity level of the data, such as &#39;non_sensitive&#39;, &#39;sensitive&#39;, or &#39;highly_sensitive&#39;. | 
**Inputs** | [**[]PropertyName**](PropertyName.md) | An array of property names to be included in the batch read operation. Each item in the array is a reference to a PropertyName object. | 

## Methods

### NewBatchReadInputPropertyName

`func NewBatchReadInputPropertyName(archived bool, dataSensitivity string, inputs []PropertyName, ) *BatchReadInputPropertyName`

NewBatchReadInputPropertyName instantiates a new BatchReadInputPropertyName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchReadInputPropertyNameWithDefaults

`func NewBatchReadInputPropertyNameWithDefaults() *BatchReadInputPropertyName`

NewBatchReadInputPropertyNameWithDefaults instantiates a new BatchReadInputPropertyName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *BatchReadInputPropertyName) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *BatchReadInputPropertyName) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *BatchReadInputPropertyName) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetDataSensitivity

`func (o *BatchReadInputPropertyName) GetDataSensitivity() string`

GetDataSensitivity returns the DataSensitivity field if non-nil, zero value otherwise.

### GetDataSensitivityOk

`func (o *BatchReadInputPropertyName) GetDataSensitivityOk() (*string, bool)`

GetDataSensitivityOk returns a tuple with the DataSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSensitivity

`func (o *BatchReadInputPropertyName) SetDataSensitivity(v string)`

SetDataSensitivity sets DataSensitivity field to given value.


### GetInputs

`func (o *BatchReadInputPropertyName) GetInputs() []PropertyName`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *BatchReadInputPropertyName) GetInputsOk() (*[]PropertyName, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *BatchReadInputPropertyName) SetInputs(v []PropertyName)`

SetInputs sets Inputs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


