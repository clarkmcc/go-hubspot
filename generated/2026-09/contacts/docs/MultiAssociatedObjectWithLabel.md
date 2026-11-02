# MultiAssociatedObjectWithLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationTypes** | [**[]AssociationSpecWithLabel**](AssociationSpecWithLabel.md) | An array of association types, each including a label. Each item in the array is an object that specifies the type of association and its label. | 
**ToObjectId** | **string** | The unique identifier of the target object to which associations are made. It is a string. | 

## Methods

### NewMultiAssociatedObjectWithLabel

`func NewMultiAssociatedObjectWithLabel(associationTypes []AssociationSpecWithLabel, toObjectId string, ) *MultiAssociatedObjectWithLabel`

NewMultiAssociatedObjectWithLabel instantiates a new MultiAssociatedObjectWithLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiAssociatedObjectWithLabelWithDefaults

`func NewMultiAssociatedObjectWithLabelWithDefaults() *MultiAssociatedObjectWithLabel`

NewMultiAssociatedObjectWithLabelWithDefaults instantiates a new MultiAssociatedObjectWithLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationTypes

`func (o *MultiAssociatedObjectWithLabel) GetAssociationTypes() []AssociationSpecWithLabel`

GetAssociationTypes returns the AssociationTypes field if non-nil, zero value otherwise.

### GetAssociationTypesOk

`func (o *MultiAssociatedObjectWithLabel) GetAssociationTypesOk() (*[]AssociationSpecWithLabel, bool)`

GetAssociationTypesOk returns a tuple with the AssociationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationTypes

`func (o *MultiAssociatedObjectWithLabel) SetAssociationTypes(v []AssociationSpecWithLabel)`

SetAssociationTypes sets AssociationTypes field to given value.


### GetToObjectId

`func (o *MultiAssociatedObjectWithLabel) GetToObjectId() string`

GetToObjectId returns the ToObjectId field if non-nil, zero value otherwise.

### GetToObjectIdOk

`func (o *MultiAssociatedObjectWithLabel) GetToObjectIdOk() (*string, bool)`

GetToObjectIdOk returns a tuple with the ToObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToObjectId

`func (o *MultiAssociatedObjectWithLabel) SetToObjectId(v string)`

SetToObjectId sets ToObjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


