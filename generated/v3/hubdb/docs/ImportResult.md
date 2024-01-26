# ImportResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RowLimitExceeded** | **bool** | Specifies whether row limit exceeded during import | 
**DuplicateRows** | **int32** | Specifies number of duplicate rows | 
**RowsImported** | **int32** | Specifies number of rows imported | 
**Errors** | [**[]Error**](Error.md) | List of errors during import | 

## Methods

### NewImportResult

`func NewImportResult(rowLimitExceeded bool, duplicateRows int32, rowsImported int32, errors []Error, ) *ImportResult`

NewImportResult instantiates a new ImportResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportResultWithDefaults

`func NewImportResultWithDefaults() *ImportResult`

NewImportResultWithDefaults instantiates a new ImportResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRowLimitExceeded

`func (o *ImportResult) GetRowLimitExceeded() bool`

GetRowLimitExceeded returns the RowLimitExceeded field if non-nil, zero value otherwise.

### GetRowLimitExceededOk

`func (o *ImportResult) GetRowLimitExceededOk() (*bool, bool)`

GetRowLimitExceededOk returns a tuple with the RowLimitExceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowLimitExceeded

`func (o *ImportResult) SetRowLimitExceeded(v bool)`

SetRowLimitExceeded sets RowLimitExceeded field to given value.


### GetDuplicateRows

`func (o *ImportResult) GetDuplicateRows() int32`

GetDuplicateRows returns the DuplicateRows field if non-nil, zero value otherwise.

### GetDuplicateRowsOk

`func (o *ImportResult) GetDuplicateRowsOk() (*int32, bool)`

GetDuplicateRowsOk returns a tuple with the DuplicateRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateRows

`func (o *ImportResult) SetDuplicateRows(v int32)`

SetDuplicateRows sets DuplicateRows field to given value.


### GetRowsImported

`func (o *ImportResult) GetRowsImported() int32`

GetRowsImported returns the RowsImported field if non-nil, zero value otherwise.

### GetRowsImportedOk

`func (o *ImportResult) GetRowsImportedOk() (*int32, bool)`

GetRowsImportedOk returns a tuple with the RowsImported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowsImported

`func (o *ImportResult) SetRowsImported(v int32)`

SetRowsImported sets RowsImported field to given value.


### GetErrors

`func (o *ImportResult) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ImportResult) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ImportResult) SetErrors(v []Error)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


