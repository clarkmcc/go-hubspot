# CollectionResponseAssociatedIdForwardPaging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]AssociatedId**](AssociatedId.md) |  | 
**Paging** | Pointer to [**ForwardPaging**](ForwardPaging.md) |  | [optional] 

## Methods

### NewCollectionResponseAssociatedIdForwardPaging

`func NewCollectionResponseAssociatedIdForwardPaging(results []AssociatedId, ) *CollectionResponseAssociatedIdForwardPaging`

NewCollectionResponseAssociatedIdForwardPaging instantiates a new CollectionResponseAssociatedIdForwardPaging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseAssociatedIdForwardPagingWithDefaults

`func NewCollectionResponseAssociatedIdForwardPagingWithDefaults() *CollectionResponseAssociatedIdForwardPaging`

NewCollectionResponseAssociatedIdForwardPagingWithDefaults instantiates a new CollectionResponseAssociatedIdForwardPaging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *CollectionResponseAssociatedIdForwardPaging) GetResults() []AssociatedId`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponseAssociatedIdForwardPaging) GetResultsOk() (*[]AssociatedId, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponseAssociatedIdForwardPaging) SetResults(v []AssociatedId)`

SetResults sets Results field to given value.


### GetPaging

`func (o *CollectionResponseAssociatedIdForwardPaging) GetPaging() ForwardPaging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponseAssociatedIdForwardPaging) GetPagingOk() (*ForwardPaging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponseAssociatedIdForwardPaging) SetPaging(v ForwardPaging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponseAssociatedIdForwardPaging) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


