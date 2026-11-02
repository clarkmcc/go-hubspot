# PublicAssociationDefinitionConfigurationUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** | A string representing the category of the association. Valid values include &#39;HUBSPOT_DEFINED&#39;, &#39;USER_DEFINED&#39;, &#39;INTEGRATOR_DEFINED&#39;, and &#39;WORK&#39;. | 
**MaxToObjectIds** | **int32** | An integer specifying the maximum number of object IDs that can be associated. | 
**TypeId** | **int32** | An integer that uniquely identifies the type of the association. | 

## Methods

### NewPublicAssociationDefinitionConfigurationUpdateRequest

`func NewPublicAssociationDefinitionConfigurationUpdateRequest(category string, maxToObjectIds int32, typeId int32, ) *PublicAssociationDefinitionConfigurationUpdateRequest`

NewPublicAssociationDefinitionConfigurationUpdateRequest instantiates a new PublicAssociationDefinitionConfigurationUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAssociationDefinitionConfigurationUpdateRequestWithDefaults

`func NewPublicAssociationDefinitionConfigurationUpdateRequestWithDefaults() *PublicAssociationDefinitionConfigurationUpdateRequest`

NewPublicAssociationDefinitionConfigurationUpdateRequestWithDefaults instantiates a new PublicAssociationDefinitionConfigurationUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *PublicAssociationDefinitionConfigurationUpdateRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PublicAssociationDefinitionConfigurationUpdateRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PublicAssociationDefinitionConfigurationUpdateRequest) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetMaxToObjectIds

`func (o *PublicAssociationDefinitionConfigurationUpdateRequest) GetMaxToObjectIds() int32`

GetMaxToObjectIds returns the MaxToObjectIds field if non-nil, zero value otherwise.

### GetMaxToObjectIdsOk

`func (o *PublicAssociationDefinitionConfigurationUpdateRequest) GetMaxToObjectIdsOk() (*int32, bool)`

GetMaxToObjectIdsOk returns a tuple with the MaxToObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxToObjectIds

`func (o *PublicAssociationDefinitionConfigurationUpdateRequest) SetMaxToObjectIds(v int32)`

SetMaxToObjectIds sets MaxToObjectIds field to given value.


### GetTypeId

`func (o *PublicAssociationDefinitionConfigurationUpdateRequest) GetTypeId() int32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *PublicAssociationDefinitionConfigurationUpdateRequest) GetTypeIdOk() (*int32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *PublicAssociationDefinitionConfigurationUpdateRequest) SetTypeId(v int32)`

SetTypeId sets TypeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


