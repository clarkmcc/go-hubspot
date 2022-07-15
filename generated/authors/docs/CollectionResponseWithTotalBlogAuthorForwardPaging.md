# CollectionResponseWithTotalBlogAuthorForwardPaging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | Total number of blog authors. | 
**Results** | [**[]BlogAuthor**](BlogAuthor.md) | Collection of blog authors. | 
**Paging** | Pointer to [**ForwardPaging**](ForwardPaging.md) |  | [optional] 

## Methods

### NewCollectionResponseWithTotalBlogAuthorForwardPaging

`func NewCollectionResponseWithTotalBlogAuthorForwardPaging(total int32, results []BlogAuthor, ) *CollectionResponseWithTotalBlogAuthorForwardPaging`

NewCollectionResponseWithTotalBlogAuthorForwardPaging instantiates a new CollectionResponseWithTotalBlogAuthorForwardPaging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseWithTotalBlogAuthorForwardPagingWithDefaults

`func NewCollectionResponseWithTotalBlogAuthorForwardPagingWithDefaults() *CollectionResponseWithTotalBlogAuthorForwardPaging`

NewCollectionResponseWithTotalBlogAuthorForwardPagingWithDefaults instantiates a new CollectionResponseWithTotalBlogAuthorForwardPaging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *CollectionResponseWithTotalBlogAuthorForwardPaging) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CollectionResponseWithTotalBlogAuthorForwardPaging) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CollectionResponseWithTotalBlogAuthorForwardPaging) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetResults

`func (o *CollectionResponseWithTotalBlogAuthorForwardPaging) GetResults() []BlogAuthor`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponseWithTotalBlogAuthorForwardPaging) GetResultsOk() (*[]BlogAuthor, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponseWithTotalBlogAuthorForwardPaging) SetResults(v []BlogAuthor)`

SetResults sets Results field to given value.


### GetPaging

`func (o *CollectionResponseWithTotalBlogAuthorForwardPaging) GetPaging() ForwardPaging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponseWithTotalBlogAuthorForwardPaging) GetPagingOk() (*ForwardPaging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponseWithTotalBlogAuthorForwardPaging) SetPaging(v ForwardPaging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponseWithTotalBlogAuthorForwardPaging) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


