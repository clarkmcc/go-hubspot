# AssociationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationCategory** | **string** | For [labeled associations](https://developers.hubspot.com/docs/guides/api/crm/associations/associations-v4#associate-records-with-a-label), the category of the association. | 
**AssociationTypeId** | **int32** | The [association type ID](https://developers.hubspot.com/docs/guides/api/crm/associations/associations-v4#association-type-id-values) (e.g., &#x60;4&#x60; for contact-to-company associations). | 

## Methods

### NewAssociationSpec

`func NewAssociationSpec(associationCategory string, associationTypeId int32, ) *AssociationSpec`

NewAssociationSpec instantiates a new AssociationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociationSpecWithDefaults

`func NewAssociationSpecWithDefaults() *AssociationSpec`

NewAssociationSpecWithDefaults instantiates a new AssociationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationCategory

`func (o *AssociationSpec) GetAssociationCategory() string`

GetAssociationCategory returns the AssociationCategory field if non-nil, zero value otherwise.

### GetAssociationCategoryOk

`func (o *AssociationSpec) GetAssociationCategoryOk() (*string, bool)`

GetAssociationCategoryOk returns a tuple with the AssociationCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationCategory

`func (o *AssociationSpec) SetAssociationCategory(v string)`

SetAssociationCategory sets AssociationCategory field to given value.


### GetAssociationTypeId

`func (o *AssociationSpec) GetAssociationTypeId() int32`

GetAssociationTypeId returns the AssociationTypeId field if non-nil, zero value otherwise.

### GetAssociationTypeIdOk

`func (o *AssociationSpec) GetAssociationTypeIdOk() (*int32, bool)`

GetAssociationTypeIdOk returns a tuple with the AssociationTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationTypeId

`func (o *AssociationSpec) SetAssociationTypeId(v int32)`

SetAssociationTypeId sets AssociationTypeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


