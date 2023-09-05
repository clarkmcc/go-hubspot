# PublicImportError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorType** | **string** |  | 
**SourceData** | [**ImportRowCore**](ImportRowCore.md) |  | 
**ObjectType** | Pointer to **string** |  | [optional] 
**InvalidValue** | Pointer to **string** |  | [optional] 
**ExtraContext** | Pointer to **string** |  | [optional] 
**ObjectTypeId** | Pointer to **string** |  | [optional] 
**KnownColumnNumber** | Pointer to **int32** |  | [optional] 
**CreatedAt** | **int32** |  | 
**Id** | **string** |  | 

## Methods

### NewPublicImportError

`func NewPublicImportError(errorType string, sourceData ImportRowCore, createdAt int32, id string, ) *PublicImportError`

NewPublicImportError instantiates a new PublicImportError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicImportErrorWithDefaults

`func NewPublicImportErrorWithDefaults() *PublicImportError`

NewPublicImportErrorWithDefaults instantiates a new PublicImportError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorType

`func (o *PublicImportError) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *PublicImportError) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *PublicImportError) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.


### GetSourceData

`func (o *PublicImportError) GetSourceData() ImportRowCore`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *PublicImportError) GetSourceDataOk() (*ImportRowCore, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *PublicImportError) SetSourceData(v ImportRowCore)`

SetSourceData sets SourceData field to given value.


### GetObjectType

`func (o *PublicImportError) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PublicImportError) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PublicImportError) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *PublicImportError) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetInvalidValue

`func (o *PublicImportError) GetInvalidValue() string`

GetInvalidValue returns the InvalidValue field if non-nil, zero value otherwise.

### GetInvalidValueOk

`func (o *PublicImportError) GetInvalidValueOk() (*string, bool)`

GetInvalidValueOk returns a tuple with the InvalidValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidValue

`func (o *PublicImportError) SetInvalidValue(v string)`

SetInvalidValue sets InvalidValue field to given value.

### HasInvalidValue

`func (o *PublicImportError) HasInvalidValue() bool`

HasInvalidValue returns a boolean if a field has been set.

### GetExtraContext

`func (o *PublicImportError) GetExtraContext() string`

GetExtraContext returns the ExtraContext field if non-nil, zero value otherwise.

### GetExtraContextOk

`func (o *PublicImportError) GetExtraContextOk() (*string, bool)`

GetExtraContextOk returns a tuple with the ExtraContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraContext

`func (o *PublicImportError) SetExtraContext(v string)`

SetExtraContext sets ExtraContext field to given value.

### HasExtraContext

`func (o *PublicImportError) HasExtraContext() bool`

HasExtraContext returns a boolean if a field has been set.

### GetObjectTypeId

`func (o *PublicImportError) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *PublicImportError) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *PublicImportError) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.

### HasObjectTypeId

`func (o *PublicImportError) HasObjectTypeId() bool`

HasObjectTypeId returns a boolean if a field has been set.

### GetKnownColumnNumber

`func (o *PublicImportError) GetKnownColumnNumber() int32`

GetKnownColumnNumber returns the KnownColumnNumber field if non-nil, zero value otherwise.

### GetKnownColumnNumberOk

`func (o *PublicImportError) GetKnownColumnNumberOk() (*int32, bool)`

GetKnownColumnNumberOk returns a tuple with the KnownColumnNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownColumnNumber

`func (o *PublicImportError) SetKnownColumnNumber(v int32)`

SetKnownColumnNumber sets KnownColumnNumber field to given value.

### HasKnownColumnNumber

`func (o *PublicImportError) HasKnownColumnNumber() bool`

HasKnownColumnNumber returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PublicImportError) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicImportError) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicImportError) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *PublicImportError) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicImportError) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicImportError) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


