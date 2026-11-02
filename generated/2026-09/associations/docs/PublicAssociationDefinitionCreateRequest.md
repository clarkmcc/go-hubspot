# PublicAssociationDefinitionCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InverseLabel** | Pointer to **string** | A string representing the inverse label for the association. This property is optional. | [optional] 
**Label** | **string** | A string representing the label for the association. This is a required property. | 
**Name** | **string** | A string representing the name of the association definition. This is a required property. | 

## Methods

### NewPublicAssociationDefinitionCreateRequest

`func NewPublicAssociationDefinitionCreateRequest(label string, name string, ) *PublicAssociationDefinitionCreateRequest`

NewPublicAssociationDefinitionCreateRequest instantiates a new PublicAssociationDefinitionCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAssociationDefinitionCreateRequestWithDefaults

`func NewPublicAssociationDefinitionCreateRequestWithDefaults() *PublicAssociationDefinitionCreateRequest`

NewPublicAssociationDefinitionCreateRequestWithDefaults instantiates a new PublicAssociationDefinitionCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInverseLabel

`func (o *PublicAssociationDefinitionCreateRequest) GetInverseLabel() string`

GetInverseLabel returns the InverseLabel field if non-nil, zero value otherwise.

### GetInverseLabelOk

`func (o *PublicAssociationDefinitionCreateRequest) GetInverseLabelOk() (*string, bool)`

GetInverseLabelOk returns a tuple with the InverseLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverseLabel

`func (o *PublicAssociationDefinitionCreateRequest) SetInverseLabel(v string)`

SetInverseLabel sets InverseLabel field to given value.

### HasInverseLabel

`func (o *PublicAssociationDefinitionCreateRequest) HasInverseLabel() bool`

HasInverseLabel returns a boolean if a field has been set.

### GetLabel

`func (o *PublicAssociationDefinitionCreateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PublicAssociationDefinitionCreateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PublicAssociationDefinitionCreateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *PublicAssociationDefinitionCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicAssociationDefinitionCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicAssociationDefinitionCreateRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


