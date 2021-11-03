# SingleFieldDependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DependencyType** | **string** |  | [default to "SINGLE_FIELD"]
**DependentFieldNames** | **[]string** |  | 
**ControllingFieldName** | **string** |  | 

## Methods

### NewSingleFieldDependency

`func NewSingleFieldDependency(dependencyType string, dependentFieldNames []string, controllingFieldName string, ) *SingleFieldDependency`

NewSingleFieldDependency instantiates a new SingleFieldDependency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleFieldDependencyWithDefaults

`func NewSingleFieldDependencyWithDefaults() *SingleFieldDependency`

NewSingleFieldDependencyWithDefaults instantiates a new SingleFieldDependency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDependencyType

`func (o *SingleFieldDependency) GetDependencyType() string`

GetDependencyType returns the DependencyType field if non-nil, zero value otherwise.

### GetDependencyTypeOk

`func (o *SingleFieldDependency) GetDependencyTypeOk() (*string, bool)`

GetDependencyTypeOk returns a tuple with the DependencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyType

`func (o *SingleFieldDependency) SetDependencyType(v string)`

SetDependencyType sets DependencyType field to given value.


### GetDependentFieldNames

`func (o *SingleFieldDependency) GetDependentFieldNames() []string`

GetDependentFieldNames returns the DependentFieldNames field if non-nil, zero value otherwise.

### GetDependentFieldNamesOk

`func (o *SingleFieldDependency) GetDependentFieldNamesOk() (*[]string, bool)`

GetDependentFieldNamesOk returns a tuple with the DependentFieldNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentFieldNames

`func (o *SingleFieldDependency) SetDependentFieldNames(v []string)`

SetDependentFieldNames sets DependentFieldNames field to given value.


### GetControllingFieldName

`func (o *SingleFieldDependency) GetControllingFieldName() string`

GetControllingFieldName returns the ControllingFieldName field if non-nil, zero value otherwise.

### GetControllingFieldNameOk

`func (o *SingleFieldDependency) GetControllingFieldNameOk() (*string, bool)`

GetControllingFieldNameOk returns a tuple with the ControllingFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllingFieldName

`func (o *SingleFieldDependency) SetControllingFieldName(v string)`

SetControllingFieldName sets ControllingFieldName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


