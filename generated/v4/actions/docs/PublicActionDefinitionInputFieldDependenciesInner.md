# PublicActionDefinitionInputFieldDependenciesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DependencyType** | **string** |  | [default to "CONDITIONAL_SINGLE_FIELD"]
**DependentFieldNames** | **[]string** |  | 
**ControllingFieldName** | **string** |  | 
**ControllingFieldValue** | **string** |  | 

## Methods

### NewPublicActionDefinitionInputFieldDependenciesInner

`func NewPublicActionDefinitionInputFieldDependenciesInner(dependencyType string, dependentFieldNames []string, controllingFieldName string, controllingFieldValue string, ) *PublicActionDefinitionInputFieldDependenciesInner`

NewPublicActionDefinitionInputFieldDependenciesInner instantiates a new PublicActionDefinitionInputFieldDependenciesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicActionDefinitionInputFieldDependenciesInnerWithDefaults

`func NewPublicActionDefinitionInputFieldDependenciesInnerWithDefaults() *PublicActionDefinitionInputFieldDependenciesInner`

NewPublicActionDefinitionInputFieldDependenciesInnerWithDefaults instantiates a new PublicActionDefinitionInputFieldDependenciesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDependencyType

`func (o *PublicActionDefinitionInputFieldDependenciesInner) GetDependencyType() string`

GetDependencyType returns the DependencyType field if non-nil, zero value otherwise.

### GetDependencyTypeOk

`func (o *PublicActionDefinitionInputFieldDependenciesInner) GetDependencyTypeOk() (*string, bool)`

GetDependencyTypeOk returns a tuple with the DependencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyType

`func (o *PublicActionDefinitionInputFieldDependenciesInner) SetDependencyType(v string)`

SetDependencyType sets DependencyType field to given value.


### GetDependentFieldNames

`func (o *PublicActionDefinitionInputFieldDependenciesInner) GetDependentFieldNames() []string`

GetDependentFieldNames returns the DependentFieldNames field if non-nil, zero value otherwise.

### GetDependentFieldNamesOk

`func (o *PublicActionDefinitionInputFieldDependenciesInner) GetDependentFieldNamesOk() (*[]string, bool)`

GetDependentFieldNamesOk returns a tuple with the DependentFieldNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentFieldNames

`func (o *PublicActionDefinitionInputFieldDependenciesInner) SetDependentFieldNames(v []string)`

SetDependentFieldNames sets DependentFieldNames field to given value.


### GetControllingFieldName

`func (o *PublicActionDefinitionInputFieldDependenciesInner) GetControllingFieldName() string`

GetControllingFieldName returns the ControllingFieldName field if non-nil, zero value otherwise.

### GetControllingFieldNameOk

`func (o *PublicActionDefinitionInputFieldDependenciesInner) GetControllingFieldNameOk() (*string, bool)`

GetControllingFieldNameOk returns a tuple with the ControllingFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllingFieldName

`func (o *PublicActionDefinitionInputFieldDependenciesInner) SetControllingFieldName(v string)`

SetControllingFieldName sets ControllingFieldName field to given value.


### GetControllingFieldValue

`func (o *PublicActionDefinitionInputFieldDependenciesInner) GetControllingFieldValue() string`

GetControllingFieldValue returns the ControllingFieldValue field if non-nil, zero value otherwise.

### GetControllingFieldValueOk

`func (o *PublicActionDefinitionInputFieldDependenciesInner) GetControllingFieldValueOk() (*string, bool)`

GetControllingFieldValueOk returns a tuple with the ControllingFieldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllingFieldValue

`func (o *PublicActionDefinitionInputFieldDependenciesInner) SetControllingFieldValue(v string)`

SetControllingFieldValue sets ControllingFieldValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


