# AssociationSpecWithLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** | The category of the association, indicating its origin or purpose. Valid values include &#39;HUBSPOT_DEFINED&#39;, &#39;USER_DEFINED&#39;, &#39;INTEGRATOR_DEFINED&#39;, and &#39;WORK&#39;. | 
**FromObjectTypeId** | Pointer to **string** | The string identifier for the object type from which the association originates. | [optional] 
**Label** | Pointer to **string** | A string label that provides additional context or description for the association. | [optional] 
**ToObjectTypeId** | Pointer to **string** | The string identifier for the object type to which the association is directed. | [optional] 
**TypeId** | **int32** | An integer representing the unique identifier for the type of association. | 

## Methods

### NewAssociationSpecWithLabel

`func NewAssociationSpecWithLabel(category string, typeId int32, ) *AssociationSpecWithLabel`

NewAssociationSpecWithLabel instantiates a new AssociationSpecWithLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociationSpecWithLabelWithDefaults

`func NewAssociationSpecWithLabelWithDefaults() *AssociationSpecWithLabel`

NewAssociationSpecWithLabelWithDefaults instantiates a new AssociationSpecWithLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *AssociationSpecWithLabel) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AssociationSpecWithLabel) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AssociationSpecWithLabel) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetFromObjectTypeId

`func (o *AssociationSpecWithLabel) GetFromObjectTypeId() string`

GetFromObjectTypeId returns the FromObjectTypeId field if non-nil, zero value otherwise.

### GetFromObjectTypeIdOk

`func (o *AssociationSpecWithLabel) GetFromObjectTypeIdOk() (*string, bool)`

GetFromObjectTypeIdOk returns a tuple with the FromObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromObjectTypeId

`func (o *AssociationSpecWithLabel) SetFromObjectTypeId(v string)`

SetFromObjectTypeId sets FromObjectTypeId field to given value.

### HasFromObjectTypeId

`func (o *AssociationSpecWithLabel) HasFromObjectTypeId() bool`

HasFromObjectTypeId returns a boolean if a field has been set.

### GetLabel

`func (o *AssociationSpecWithLabel) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AssociationSpecWithLabel) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AssociationSpecWithLabel) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AssociationSpecWithLabel) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetToObjectTypeId

`func (o *AssociationSpecWithLabel) GetToObjectTypeId() string`

GetToObjectTypeId returns the ToObjectTypeId field if non-nil, zero value otherwise.

### GetToObjectTypeIdOk

`func (o *AssociationSpecWithLabel) GetToObjectTypeIdOk() (*string, bool)`

GetToObjectTypeIdOk returns a tuple with the ToObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToObjectTypeId

`func (o *AssociationSpecWithLabel) SetToObjectTypeId(v string)`

SetToObjectTypeId sets ToObjectTypeId field to given value.

### HasToObjectTypeId

`func (o *AssociationSpecWithLabel) HasToObjectTypeId() bool`

HasToObjectTypeId returns a boolean if a field has been set.

### GetTypeId

`func (o *AssociationSpecWithLabel) GetTypeId() int32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *AssociationSpecWithLabel) GetTypeIdOk() (*int32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *AssociationSpecWithLabel) SetTypeId(v int32)`

SetTypeId sets TypeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


