# ActionFunctionIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionType** | **string** | The type of function. This determines when the function will be called. | 
**Id** | Pointer to **string** | The ID qualifier for the function. This is used to specify which input field a function is associated with for &#x60;PRE_FETCH_OPTIONS&#x60; and &#x60;POST_FETCH_OPTIONS&#x60; function types. | [optional] 

## Methods

### NewActionFunctionIdentifier

`func NewActionFunctionIdentifier(functionType string, ) *ActionFunctionIdentifier`

NewActionFunctionIdentifier instantiates a new ActionFunctionIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionFunctionIdentifierWithDefaults

`func NewActionFunctionIdentifierWithDefaults() *ActionFunctionIdentifier`

NewActionFunctionIdentifierWithDefaults instantiates a new ActionFunctionIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionType

`func (o *ActionFunctionIdentifier) GetFunctionType() string`

GetFunctionType returns the FunctionType field if non-nil, zero value otherwise.

### GetFunctionTypeOk

`func (o *ActionFunctionIdentifier) GetFunctionTypeOk() (*string, bool)`

GetFunctionTypeOk returns a tuple with the FunctionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionType

`func (o *ActionFunctionIdentifier) SetFunctionType(v string)`

SetFunctionType sets FunctionType field to given value.


### GetId

`func (o *ActionFunctionIdentifier) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionFunctionIdentifier) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionFunctionIdentifier) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActionFunctionIdentifier) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


