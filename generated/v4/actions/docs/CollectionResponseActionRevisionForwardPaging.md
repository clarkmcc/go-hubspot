# CollectionResponseActionRevisionForwardPaging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]ActionRevision**](ActionRevision.md) |  | 
**Paging** | Pointer to [**ForwardPaging**](ForwardPaging.md) |  | [optional] 

## Methods

### NewCollectionResponseActionRevisionForwardPaging

`func NewCollectionResponseActionRevisionForwardPaging(results []ActionRevision, ) *CollectionResponseActionRevisionForwardPaging`

NewCollectionResponseActionRevisionForwardPaging instantiates a new CollectionResponseActionRevisionForwardPaging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseActionRevisionForwardPagingWithDefaults

`func NewCollectionResponseActionRevisionForwardPagingWithDefaults() *CollectionResponseActionRevisionForwardPaging`

NewCollectionResponseActionRevisionForwardPagingWithDefaults instantiates a new CollectionResponseActionRevisionForwardPaging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *CollectionResponseActionRevisionForwardPaging) GetResults() []ActionRevision`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponseActionRevisionForwardPaging) GetResultsOk() (*[]ActionRevision, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponseActionRevisionForwardPaging) SetResults(v []ActionRevision)`

SetResults sets Results field to given value.


### GetPaging

`func (o *CollectionResponseActionRevisionForwardPaging) GetPaging() ForwardPaging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponseActionRevisionForwardPaging) GetPagingOk() (*ForwardPaging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponseActionRevisionForwardPaging) SetPaging(v ForwardPaging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponseActionRevisionForwardPaging) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


