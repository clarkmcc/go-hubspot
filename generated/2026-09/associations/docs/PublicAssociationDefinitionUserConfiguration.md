# PublicAssociationDefinitionUserConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** | The category of the association definition. Valid values include &#39;HUBSPOT_DEFINED&#39;, &#39;USER_DEFINED&#39;, &#39;INTEGRATOR_DEFINED&#39;, and &#39;WORK&#39;. | 
**Label** | Pointer to **string** | A string representing the label for the association, which provides a human-readable name for the association type. | [optional] 
**TypeId** | **int32** | The unique identifier for the association type, represented as an integer. | 
**UserEnforcedMaxToObjectIds** | Pointer to **int32** | An integer specifying the maximum number of &#39;to&#39; object IDs that can be associated with a single &#39;from&#39; object ID, as enforced by the user. | [optional] 

## Methods

### NewPublicAssociationDefinitionUserConfiguration

`func NewPublicAssociationDefinitionUserConfiguration(category string, typeId int32, ) *PublicAssociationDefinitionUserConfiguration`

NewPublicAssociationDefinitionUserConfiguration instantiates a new PublicAssociationDefinitionUserConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAssociationDefinitionUserConfigurationWithDefaults

`func NewPublicAssociationDefinitionUserConfigurationWithDefaults() *PublicAssociationDefinitionUserConfiguration`

NewPublicAssociationDefinitionUserConfigurationWithDefaults instantiates a new PublicAssociationDefinitionUserConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *PublicAssociationDefinitionUserConfiguration) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PublicAssociationDefinitionUserConfiguration) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PublicAssociationDefinitionUserConfiguration) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetLabel

`func (o *PublicAssociationDefinitionUserConfiguration) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PublicAssociationDefinitionUserConfiguration) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PublicAssociationDefinitionUserConfiguration) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PublicAssociationDefinitionUserConfiguration) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetTypeId

`func (o *PublicAssociationDefinitionUserConfiguration) GetTypeId() int32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *PublicAssociationDefinitionUserConfiguration) GetTypeIdOk() (*int32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *PublicAssociationDefinitionUserConfiguration) SetTypeId(v int32)`

SetTypeId sets TypeId field to given value.


### GetUserEnforcedMaxToObjectIds

`func (o *PublicAssociationDefinitionUserConfiguration) GetUserEnforcedMaxToObjectIds() int32`

GetUserEnforcedMaxToObjectIds returns the UserEnforcedMaxToObjectIds field if non-nil, zero value otherwise.

### GetUserEnforcedMaxToObjectIdsOk

`func (o *PublicAssociationDefinitionUserConfiguration) GetUserEnforcedMaxToObjectIdsOk() (*int32, bool)`

GetUserEnforcedMaxToObjectIdsOk returns a tuple with the UserEnforcedMaxToObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEnforcedMaxToObjectIds

`func (o *PublicAssociationDefinitionUserConfiguration) SetUserEnforcedMaxToObjectIds(v int32)`

SetUserEnforcedMaxToObjectIds sets UserEnforcedMaxToObjectIds field to given value.

### HasUserEnforcedMaxToObjectIds

`func (o *PublicAssociationDefinitionUserConfiguration) HasUserEnforcedMaxToObjectIds() bool`

HasUserEnforcedMaxToObjectIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


