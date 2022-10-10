# TaskLocator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Links** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTaskLocator

`func NewTaskLocator(id string, ) *TaskLocator`

NewTaskLocator instantiates a new TaskLocator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskLocatorWithDefaults

`func NewTaskLocatorWithDefaults() *TaskLocator`

NewTaskLocatorWithDefaults instantiates a new TaskLocator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskLocator) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskLocator) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskLocator) SetId(v string)`

SetId sets Id field to given value.


### GetLinks

`func (o *TaskLocator) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TaskLocator) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TaskLocator) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TaskLocator) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


