# ErrorDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | A human readable message describing the error along with remediation steps where appropriate | 
**In** | Pointer to **string** | The name of the field or parameter in which the error was found. | [optional] 
**Code** | Pointer to **string** | The status code associated with the error detail | [optional] 
**SubCategory** | Pointer to **string** | A specific category that contains more specific detail about the error | [optional] 
**Context** | Pointer to **map[string][]string** | Context about the error condition | [optional] 

## Methods

### NewErrorDetail

`func NewErrorDetail(message string, ) *ErrorDetail`

NewErrorDetail instantiates a new ErrorDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorDetailWithDefaults

`func NewErrorDetailWithDefaults() *ErrorDetail`

NewErrorDetailWithDefaults instantiates a new ErrorDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ErrorDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorDetail) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetIn

`func (o *ErrorDetail) GetIn() string`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *ErrorDetail) GetInOk() (*string, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *ErrorDetail) SetIn(v string)`

SetIn sets In field to given value.

### HasIn

`func (o *ErrorDetail) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetCode

`func (o *ErrorDetail) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorDetail) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorDetail) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorDetail) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetSubCategory

`func (o *ErrorDetail) GetSubCategory() string`

GetSubCategory returns the SubCategory field if non-nil, zero value otherwise.

### GetSubCategoryOk

`func (o *ErrorDetail) GetSubCategoryOk() (*string, bool)`

GetSubCategoryOk returns a tuple with the SubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategory

`func (o *ErrorDetail) SetSubCategory(v string)`

SetSubCategory sets SubCategory field to given value.

### HasSubCategory

`func (o *ErrorDetail) HasSubCategory() bool`

HasSubCategory returns a boolean if a field has been set.

### GetContext

`func (o *ErrorDetail) GetContext() map[string][]string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ErrorDetail) GetContextOk() (*map[string][]string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ErrorDetail) SetContext(v map[string][]string)`

SetContext sets Context field to given value.

### HasContext

`func (o *ErrorDetail) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


