# PublicAssociationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** | A string indicating the category of the association. Valid categories include &#39;HUBSPOT_DEFINED&#39;, &#39;USER_DEFINED&#39;, &#39;INTEGRATOR_DEFINED&#39;, and &#39;WORK&#39;. | 
**TypeId** | **int32** | An integer representing the unique identifier for the association type. | 

## Methods

### NewPublicAssociationSpec

`func NewPublicAssociationSpec(category string, typeId int32, ) *PublicAssociationSpec`

NewPublicAssociationSpec instantiates a new PublicAssociationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAssociationSpecWithDefaults

`func NewPublicAssociationSpecWithDefaults() *PublicAssociationSpec`

NewPublicAssociationSpecWithDefaults instantiates a new PublicAssociationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *PublicAssociationSpec) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PublicAssociationSpec) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PublicAssociationSpec) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetTypeId

`func (o *PublicAssociationSpec) GetTypeId() int32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *PublicAssociationSpec) GetTypeIdOk() (*int32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *PublicAssociationSpec) SetTypeId(v int32)`

SetTypeId sets TypeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


