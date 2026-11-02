# PublicObjectWarning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** | A string indicating the category of the warning. | 
**Context** | **map[string]string** | An object providing additional context for the warning, with string values for each context key. | 
**Message** | **string** | A string containing the warning message. | 

## Methods

### NewPublicObjectWarning

`func NewPublicObjectWarning(category string, context map[string]string, message string, ) *PublicObjectWarning`

NewPublicObjectWarning instantiates a new PublicObjectWarning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicObjectWarningWithDefaults

`func NewPublicObjectWarningWithDefaults() *PublicObjectWarning`

NewPublicObjectWarningWithDefaults instantiates a new PublicObjectWarning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *PublicObjectWarning) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PublicObjectWarning) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PublicObjectWarning) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetContext

`func (o *PublicObjectWarning) GetContext() map[string]string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PublicObjectWarning) GetContextOk() (*map[string]string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PublicObjectWarning) SetContext(v map[string]string)`

SetContext sets Context field to given value.


### GetMessage

`func (o *PublicObjectWarning) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PublicObjectWarning) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PublicObjectWarning) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


