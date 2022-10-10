# IndexedField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | **map[string]interface{}** |  | 
**Values** | **[]map[string]interface{}** |  | 
**MetadataField** | **bool** |  | 

## Methods

### NewIndexedField

`func NewIndexedField(name string, value map[string]interface{}, values []map[string]interface{}, metadataField bool, ) *IndexedField`

NewIndexedField instantiates a new IndexedField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexedFieldWithDefaults

`func NewIndexedFieldWithDefaults() *IndexedField`

NewIndexedFieldWithDefaults instantiates a new IndexedField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IndexedField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IndexedField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IndexedField) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *IndexedField) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IndexedField) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IndexedField) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.


### GetValues

`func (o *IndexedField) GetValues() []map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *IndexedField) GetValuesOk() (*[]map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *IndexedField) SetValues(v []map[string]interface{})`

SetValues sets Values field to given value.


### GetMetadataField

`func (o *IndexedField) GetMetadataField() bool`

GetMetadataField returns the MetadataField field if non-nil, zero value otherwise.

### GetMetadataFieldOk

`func (o *IndexedField) GetMetadataFieldOk() (*bool, bool)`

GetMetadataFieldOk returns a tuple with the MetadataField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataField

`func (o *IndexedField) SetMetadataField(v bool)`

SetMetadataField sets MetadataField field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


