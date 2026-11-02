# PublicAssociationDefinitionConfigurationUpdateResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** | The category of the association definition. Valid values include &#39;HUBSPOT_DEFINED&#39;, &#39;USER_DEFINED&#39;, &#39;INTEGRATOR_DEFINED&#39;, and &#39;WORK&#39;. | 
**TypeId** | **int32** | The unique identifier for the type of association being updated. | 
**UserEnforcedMaxToObjectIds** | Pointer to **int32** | An integer representing the maximum number of &#39;to&#39; object IDs that can be associated, as enforced by the user. | [optional] 

## Methods

### NewPublicAssociationDefinitionConfigurationUpdateResult

`func NewPublicAssociationDefinitionConfigurationUpdateResult(category string, typeId int32, ) *PublicAssociationDefinitionConfigurationUpdateResult`

NewPublicAssociationDefinitionConfigurationUpdateResult instantiates a new PublicAssociationDefinitionConfigurationUpdateResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAssociationDefinitionConfigurationUpdateResultWithDefaults

`func NewPublicAssociationDefinitionConfigurationUpdateResultWithDefaults() *PublicAssociationDefinitionConfigurationUpdateResult`

NewPublicAssociationDefinitionConfigurationUpdateResultWithDefaults instantiates a new PublicAssociationDefinitionConfigurationUpdateResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *PublicAssociationDefinitionConfigurationUpdateResult) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PublicAssociationDefinitionConfigurationUpdateResult) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PublicAssociationDefinitionConfigurationUpdateResult) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetTypeId

`func (o *PublicAssociationDefinitionConfigurationUpdateResult) GetTypeId() int32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *PublicAssociationDefinitionConfigurationUpdateResult) GetTypeIdOk() (*int32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *PublicAssociationDefinitionConfigurationUpdateResult) SetTypeId(v int32)`

SetTypeId sets TypeId field to given value.


### GetUserEnforcedMaxToObjectIds

`func (o *PublicAssociationDefinitionConfigurationUpdateResult) GetUserEnforcedMaxToObjectIds() int32`

GetUserEnforcedMaxToObjectIds returns the UserEnforcedMaxToObjectIds field if non-nil, zero value otherwise.

### GetUserEnforcedMaxToObjectIdsOk

`func (o *PublicAssociationDefinitionConfigurationUpdateResult) GetUserEnforcedMaxToObjectIdsOk() (*int32, bool)`

GetUserEnforcedMaxToObjectIdsOk returns a tuple with the UserEnforcedMaxToObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEnforcedMaxToObjectIds

`func (o *PublicAssociationDefinitionConfigurationUpdateResult) SetUserEnforcedMaxToObjectIds(v int32)`

SetUserEnforcedMaxToObjectIds sets UserEnforcedMaxToObjectIds field to given value.

### HasUserEnforcedMaxToObjectIds

`func (o *PublicAssociationDefinitionConfigurationUpdateResult) HasUserEnforcedMaxToObjectIds() bool`

HasUserEnforcedMaxToObjectIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


