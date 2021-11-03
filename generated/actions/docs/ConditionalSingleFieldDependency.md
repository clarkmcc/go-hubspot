# ConditionalSingleFieldDependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DependencyType** | **string** |  | [default to "CONDITIONAL_SINGLE_FIELD"]
**DependentFieldNames** | **[]string** |  | 
**ControllingFieldName** | **string** |  | 
**ControllingFieldValue** | **string** |  | 

## Methods

### NewConditionalSingleFieldDependency

`func NewConditionalSingleFieldDependency(dependencyType string, dependentFieldNames []string, controllingFieldName string, controllingFieldValue string, ) *ConditionalSingleFieldDependency`

NewConditionalSingleFieldDependency instantiates a new ConditionalSingleFieldDependency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionalSingleFieldDependencyWithDefaults

`func NewConditionalSingleFieldDependencyWithDefaults() *ConditionalSingleFieldDependency`

NewConditionalSingleFieldDependencyWithDefaults instantiates a new ConditionalSingleFieldDependency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDependencyType

`func (o *ConditionalSingleFieldDependency) GetDependencyType() string`

GetDependencyType returns the DependencyType field if non-nil, zero value otherwise.

### GetDependencyTypeOk

`func (o *ConditionalSingleFieldDependency) GetDependencyTypeOk() (*string, bool)`

GetDependencyTypeOk returns a tuple with the DependencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyType

`func (o *ConditionalSingleFieldDependency) SetDependencyType(v string)`

SetDependencyType sets DependencyType field to given value.


### GetDependentFieldNames

`func (o *ConditionalSingleFieldDependency) GetDependentFieldNames() []string`

GetDependentFieldNames returns the DependentFieldNames field if non-nil, zero value otherwise.

### GetDependentFieldNamesOk

`func (o *ConditionalSingleFieldDependency) GetDependentFieldNamesOk() (*[]string, bool)`

GetDependentFieldNamesOk returns a tuple with the DependentFieldNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentFieldNames

`func (o *ConditionalSingleFieldDependency) SetDependentFieldNames(v []string)`

SetDependentFieldNames sets DependentFieldNames field to given value.


### GetControllingFieldName

`func (o *ConditionalSingleFieldDependency) GetControllingFieldName() string`

GetControllingFieldName returns the ControllingFieldName field if non-nil, zero value otherwise.

### GetControllingFieldNameOk

`func (o *ConditionalSingleFieldDependency) GetControllingFieldNameOk() (*string, bool)`

GetControllingFieldNameOk returns a tuple with the ControllingFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllingFieldName

`func (o *ConditionalSingleFieldDependency) SetControllingFieldName(v string)`

SetControllingFieldName sets ControllingFieldName field to given value.


### GetControllingFieldValue

`func (o *ConditionalSingleFieldDependency) GetControllingFieldValue() string`

GetControllingFieldValue returns the ControllingFieldValue field if non-nil, zero value otherwise.

### GetControllingFieldValueOk

`func (o *ConditionalSingleFieldDependency) GetControllingFieldValueOk() (*string, bool)`

GetControllingFieldValueOk returns a tuple with the ControllingFieldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllingFieldValue

`func (o *ConditionalSingleFieldDependency) SetControllingFieldValue(v string)`

SetControllingFieldValue sets ControllingFieldValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


