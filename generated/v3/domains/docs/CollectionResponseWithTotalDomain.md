# CollectionResponseWithTotalDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**Results** | [**[]Domain**](Domain.md) |  | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 

## Methods

### NewCollectionResponseWithTotalDomain

`func NewCollectionResponseWithTotalDomain(total int32, results []Domain, ) *CollectionResponseWithTotalDomain`

NewCollectionResponseWithTotalDomain instantiates a new CollectionResponseWithTotalDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseWithTotalDomainWithDefaults

`func NewCollectionResponseWithTotalDomainWithDefaults() *CollectionResponseWithTotalDomain`

NewCollectionResponseWithTotalDomainWithDefaults instantiates a new CollectionResponseWithTotalDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *CollectionResponseWithTotalDomain) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CollectionResponseWithTotalDomain) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CollectionResponseWithTotalDomain) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetResults

`func (o *CollectionResponseWithTotalDomain) GetResults() []Domain`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponseWithTotalDomain) GetResultsOk() (*[]Domain, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponseWithTotalDomain) SetResults(v []Domain)`

SetResults sets Results field to given value.


### GetPaging

`func (o *CollectionResponseWithTotalDomain) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponseWithTotalDomain) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponseWithTotalDomain) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponseWithTotalDomain) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


