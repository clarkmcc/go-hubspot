# CollectionResponseWithTotalBlogPostForwardPaging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | Total number of blog posts. | 
**Results** | [**[]BlogPost**](BlogPost.md) | Collection of blog posts. | 
**Paging** | Pointer to [**ForwardPaging**](ForwardPaging.md) |  | [optional] 

## Methods

### NewCollectionResponseWithTotalBlogPostForwardPaging

`func NewCollectionResponseWithTotalBlogPostForwardPaging(total int32, results []BlogPost, ) *CollectionResponseWithTotalBlogPostForwardPaging`

NewCollectionResponseWithTotalBlogPostForwardPaging instantiates a new CollectionResponseWithTotalBlogPostForwardPaging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseWithTotalBlogPostForwardPagingWithDefaults

`func NewCollectionResponseWithTotalBlogPostForwardPagingWithDefaults() *CollectionResponseWithTotalBlogPostForwardPaging`

NewCollectionResponseWithTotalBlogPostForwardPagingWithDefaults instantiates a new CollectionResponseWithTotalBlogPostForwardPaging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *CollectionResponseWithTotalBlogPostForwardPaging) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CollectionResponseWithTotalBlogPostForwardPaging) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CollectionResponseWithTotalBlogPostForwardPaging) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetResults

`func (o *CollectionResponseWithTotalBlogPostForwardPaging) GetResults() []BlogPost`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponseWithTotalBlogPostForwardPaging) GetResultsOk() (*[]BlogPost, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponseWithTotalBlogPostForwardPaging) SetResults(v []BlogPost)`

SetResults sets Results field to given value.


### GetPaging

`func (o *CollectionResponseWithTotalBlogPostForwardPaging) GetPaging() ForwardPaging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponseWithTotalBlogPostForwardPaging) GetPagingOk() (*ForwardPaging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponseWithTotalBlogPostForwardPaging) SetPaging(v ForwardPaging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponseWithTotalBlogPostForwardPaging) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


