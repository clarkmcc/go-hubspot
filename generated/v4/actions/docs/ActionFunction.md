# ActionFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionSource** | **string** | The function source code. | 
**FunctionType** | **string** | The type of function. This determines when the function will be called. | 
**Id** | Pointer to **string** | The ID qualifier for the function. This is used to specify which input field a function is associated with for &#x60;PRE_FETCH_OPTIONS&#x60; and &#x60;POST_FETCH_OPTIONS&#x60; function types. | [optional] 

## Methods

### NewActionFunction

`func NewActionFunction(functionSource string, functionType string, ) *ActionFunction`

NewActionFunction instantiates a new ActionFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionFunctionWithDefaults

`func NewActionFunctionWithDefaults() *ActionFunction`

NewActionFunctionWithDefaults instantiates a new ActionFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionSource

`func (o *ActionFunction) GetFunctionSource() string`

GetFunctionSource returns the FunctionSource field if non-nil, zero value otherwise.

### GetFunctionSourceOk

`func (o *ActionFunction) GetFunctionSourceOk() (*string, bool)`

GetFunctionSourceOk returns a tuple with the FunctionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionSource

`func (o *ActionFunction) SetFunctionSource(v string)`

SetFunctionSource sets FunctionSource field to given value.


### GetFunctionType

`func (o *ActionFunction) GetFunctionType() string`

GetFunctionType returns the FunctionType field if non-nil, zero value otherwise.

### GetFunctionTypeOk

`func (o *ActionFunction) GetFunctionTypeOk() (*string, bool)`

GetFunctionTypeOk returns a tuple with the FunctionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionType

`func (o *ActionFunction) SetFunctionType(v string)`

SetFunctionType sets FunctionType field to given value.


### GetId

`func (o *ActionFunction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionFunction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionFunction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActionFunction) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


