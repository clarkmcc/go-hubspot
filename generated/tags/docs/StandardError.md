# StandardError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Error status. | 
**Id** | Pointer to **string** | Error ID. | [optional] 
**Category** | **map[string]interface{}** | Model definition for an error category. | 
**SubCategory** | Pointer to **map[string]interface{}** | Error subcategory. | [optional] 
**Message** | **string** | Error message. | 
**Errors** | [**[]ErrorDetail**](ErrorDetail.md) | List of error details. | 
**Context** | **map[string][]string** | Error context. | 
**Links** | **map[string]string** | Error links. | 

## Methods

### NewStandardError

`func NewStandardError(status string, category map[string]interface{}, message string, errors []ErrorDetail, context map[string][]string, links map[string]string, ) *StandardError`

NewStandardError instantiates a new StandardError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandardErrorWithDefaults

`func NewStandardErrorWithDefaults() *StandardError`

NewStandardErrorWithDefaults instantiates a new StandardError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *StandardError) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StandardError) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StandardError) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetId

`func (o *StandardError) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StandardError) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StandardError) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StandardError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCategory

`func (o *StandardError) GetCategory() map[string]interface{}`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *StandardError) GetCategoryOk() (*map[string]interface{}, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *StandardError) SetCategory(v map[string]interface{})`

SetCategory sets Category field to given value.


### GetSubCategory

`func (o *StandardError) GetSubCategory() map[string]interface{}`

GetSubCategory returns the SubCategory field if non-nil, zero value otherwise.

### GetSubCategoryOk

`func (o *StandardError) GetSubCategoryOk() (*map[string]interface{}, bool)`

GetSubCategoryOk returns a tuple with the SubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategory

`func (o *StandardError) SetSubCategory(v map[string]interface{})`

SetSubCategory sets SubCategory field to given value.

### HasSubCategory

`func (o *StandardError) HasSubCategory() bool`

HasSubCategory returns a boolean if a field has been set.

### GetMessage

`func (o *StandardError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StandardError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StandardError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrors

`func (o *StandardError) GetErrors() []ErrorDetail`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *StandardError) GetErrorsOk() (*[]ErrorDetail, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *StandardError) SetErrors(v []ErrorDetail)`

SetErrors sets Errors field to given value.


### GetContext

`func (o *StandardError) GetContext() map[string][]string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *StandardError) GetContextOk() (*map[string][]string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *StandardError) SetContext(v map[string][]string)`

SetContext sets Context field to given value.


### GetLinks

`func (o *StandardError) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StandardError) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StandardError) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


