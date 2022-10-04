# CollectionResponsePublicAuditLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]PublicAuditLog**](PublicAuditLog.md) |  | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 

## Methods

### NewCollectionResponsePublicAuditLog

`func NewCollectionResponsePublicAuditLog(results []PublicAuditLog, ) *CollectionResponsePublicAuditLog`

NewCollectionResponsePublicAuditLog instantiates a new CollectionResponsePublicAuditLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponsePublicAuditLogWithDefaults

`func NewCollectionResponsePublicAuditLogWithDefaults() *CollectionResponsePublicAuditLog`

NewCollectionResponsePublicAuditLogWithDefaults instantiates a new CollectionResponsePublicAuditLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *CollectionResponsePublicAuditLog) GetResults() []PublicAuditLog`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponsePublicAuditLog) GetResultsOk() (*[]PublicAuditLog, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponsePublicAuditLog) SetResults(v []PublicAuditLog)`

SetResults sets Results field to given value.


### GetPaging

`func (o *CollectionResponsePublicAuditLog) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponsePublicAuditLog) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponsePublicAuditLog) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponsePublicAuditLog) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


