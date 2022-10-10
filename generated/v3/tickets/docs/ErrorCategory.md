# ErrorCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**HttpStatus** | **string** |  | 

## Methods

### NewErrorCategory

`func NewErrorCategory(name string, httpStatus string, ) *ErrorCategory`

NewErrorCategory instantiates a new ErrorCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorCategoryWithDefaults

`func NewErrorCategoryWithDefaults() *ErrorCategory`

NewErrorCategoryWithDefaults instantiates a new ErrorCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ErrorCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ErrorCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ErrorCategory) SetName(v string)`

SetName sets Name field to given value.


### GetHttpStatus

`func (o *ErrorCategory) GetHttpStatus() string`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *ErrorCategory) GetHttpStatusOk() (*string, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *ErrorCategory) SetHttpStatus(v string)`

SetHttpStatus sets HttpStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


