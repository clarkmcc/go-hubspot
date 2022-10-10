# ActionRevision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Definition** | [**ExtensionActionDefinition**](ExtensionActionDefinition.md) |  | 
**CreatedAt** | **time.Time** | The date the revision was created. | 
**Id** | **string** |  | 
**RevisionId** | **string** | The version number of the custom action. | 

## Methods

### NewActionRevision

`func NewActionRevision(definition ExtensionActionDefinition, createdAt time.Time, id string, revisionId string, ) *ActionRevision`

NewActionRevision instantiates a new ActionRevision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRevisionWithDefaults

`func NewActionRevisionWithDefaults() *ActionRevision`

NewActionRevisionWithDefaults instantiates a new ActionRevision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinition

`func (o *ActionRevision) GetDefinition() ExtensionActionDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *ActionRevision) GetDefinitionOk() (*ExtensionActionDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *ActionRevision) SetDefinition(v ExtensionActionDefinition)`

SetDefinition sets Definition field to given value.


### GetCreatedAt

`func (o *ActionRevision) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ActionRevision) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ActionRevision) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *ActionRevision) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionRevision) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionRevision) SetId(v string)`

SetId sets Id field to given value.


### GetRevisionId

`func (o *ActionRevision) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *ActionRevision) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *ActionRevision) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


