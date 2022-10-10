# CollectionResponseWithTotalUrlMappingForwardPaging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**Results** | [**[]UrlMapping**](UrlMapping.md) |  | 
**Paging** | Pointer to [**ForwardPaging**](ForwardPaging.md) |  | [optional] 

## Methods

### NewCollectionResponseWithTotalUrlMappingForwardPaging

`func NewCollectionResponseWithTotalUrlMappingForwardPaging(total int32, results []UrlMapping, ) *CollectionResponseWithTotalUrlMappingForwardPaging`

NewCollectionResponseWithTotalUrlMappingForwardPaging instantiates a new CollectionResponseWithTotalUrlMappingForwardPaging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseWithTotalUrlMappingForwardPagingWithDefaults

`func NewCollectionResponseWithTotalUrlMappingForwardPagingWithDefaults() *CollectionResponseWithTotalUrlMappingForwardPaging`

NewCollectionResponseWithTotalUrlMappingForwardPagingWithDefaults instantiates a new CollectionResponseWithTotalUrlMappingForwardPaging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *CollectionResponseWithTotalUrlMappingForwardPaging) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CollectionResponseWithTotalUrlMappingForwardPaging) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CollectionResponseWithTotalUrlMappingForwardPaging) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetResults

`func (o *CollectionResponseWithTotalUrlMappingForwardPaging) GetResults() []UrlMapping`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponseWithTotalUrlMappingForwardPaging) GetResultsOk() (*[]UrlMapping, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponseWithTotalUrlMappingForwardPaging) SetResults(v []UrlMapping)`

SetResults sets Results field to given value.


### GetPaging

`func (o *CollectionResponseWithTotalUrlMappingForwardPaging) GetPaging() ForwardPaging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponseWithTotalUrlMappingForwardPaging) GetPagingOk() (*ForwardPaging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponseWithTotalUrlMappingForwardPaging) SetPaging(v ForwardPaging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponseWithTotalUrlMappingForwardPaging) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


