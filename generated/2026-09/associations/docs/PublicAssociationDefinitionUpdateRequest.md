# PublicAssociationDefinitionUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationTypeId** | **int32** | An integer representing the unique identifier for the association type. | 
**InverseLabel** | Pointer to **string** | A string representing the inverse label for the association. | [optional] 
**Label** | **string** | A string representing the label for the association. | 

## Methods

### NewPublicAssociationDefinitionUpdateRequest

`func NewPublicAssociationDefinitionUpdateRequest(associationTypeId int32, label string, ) *PublicAssociationDefinitionUpdateRequest`

NewPublicAssociationDefinitionUpdateRequest instantiates a new PublicAssociationDefinitionUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAssociationDefinitionUpdateRequestWithDefaults

`func NewPublicAssociationDefinitionUpdateRequestWithDefaults() *PublicAssociationDefinitionUpdateRequest`

NewPublicAssociationDefinitionUpdateRequestWithDefaults instantiates a new PublicAssociationDefinitionUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationTypeId

`func (o *PublicAssociationDefinitionUpdateRequest) GetAssociationTypeId() int32`

GetAssociationTypeId returns the AssociationTypeId field if non-nil, zero value otherwise.

### GetAssociationTypeIdOk

`func (o *PublicAssociationDefinitionUpdateRequest) GetAssociationTypeIdOk() (*int32, bool)`

GetAssociationTypeIdOk returns a tuple with the AssociationTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationTypeId

`func (o *PublicAssociationDefinitionUpdateRequest) SetAssociationTypeId(v int32)`

SetAssociationTypeId sets AssociationTypeId field to given value.


### GetInverseLabel

`func (o *PublicAssociationDefinitionUpdateRequest) GetInverseLabel() string`

GetInverseLabel returns the InverseLabel field if non-nil, zero value otherwise.

### GetInverseLabelOk

`func (o *PublicAssociationDefinitionUpdateRequest) GetInverseLabelOk() (*string, bool)`

GetInverseLabelOk returns a tuple with the InverseLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverseLabel

`func (o *PublicAssociationDefinitionUpdateRequest) SetInverseLabel(v string)`

SetInverseLabel sets InverseLabel field to given value.

### HasInverseLabel

`func (o *PublicAssociationDefinitionUpdateRequest) HasInverseLabel() bool`

HasInverseLabel returns a boolean if a field has been set.

### GetLabel

`func (o *PublicAssociationDefinitionUpdateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PublicAssociationDefinitionUpdateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PublicAssociationDefinitionUpdateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


