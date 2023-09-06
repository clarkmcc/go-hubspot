# Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | A human readable message describing the error along with remediation steps where appropriate | 
**CorrelationId** | **string** | A unique identifier for the request. Include this value with any error reports or support tickets | 
**Category** | **string** | The error category | 
**SubCategory** | Pointer to **string** | A specific category that contains more specific detail about the error | [optional] 
**Errors** | Pointer to [**[]ErrorDetail**](ErrorDetail.md) | further information about the error | [optional] 
**Context** | Pointer to **map[string][]string** | Context about the error condition | [optional] 
**Links** | Pointer to **map[string]string** | A map of link names to associated URIs containing documentation about the error or recommended remediation steps | [optional] 

## Methods

### NewError

`func NewError(message string, correlationId string, category string, ) *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *Error) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Error) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Error) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCorrelationId

`func (o *Error) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *Error) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *Error) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.


### GetCategory

`func (o *Error) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Error) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Error) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetSubCategory

`func (o *Error) GetSubCategory() string`

GetSubCategory returns the SubCategory field if non-nil, zero value otherwise.

### GetSubCategoryOk

`func (o *Error) GetSubCategoryOk() (*string, bool)`

GetSubCategoryOk returns a tuple with the SubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategory

`func (o *Error) SetSubCategory(v string)`

SetSubCategory sets SubCategory field to given value.

### HasSubCategory

`func (o *Error) HasSubCategory() bool`

HasSubCategory returns a boolean if a field has been set.

### GetErrors

`func (o *Error) GetErrors() []ErrorDetail`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Error) GetErrorsOk() (*[]ErrorDetail, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Error) SetErrors(v []ErrorDetail)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Error) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetContext

`func (o *Error) GetContext() map[string][]string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Error) GetContextOk() (*map[string][]string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Error) SetContext(v map[string][]string)`

SetContext sets Context field to given value.

### HasContext

`func (o *Error) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLinks

`func (o *Error) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Error) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Error) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Error) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


