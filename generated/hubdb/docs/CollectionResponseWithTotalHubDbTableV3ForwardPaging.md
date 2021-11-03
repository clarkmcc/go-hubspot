# CollectionResponseWithTotalHubDbTableV3ForwardPaging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**Results** | [**[]HubDbTableV3**](HubDbTableV3.md) |  | 
**Paging** | Pointer to [**ForwardPaging**](ForwardPaging.md) |  | [optional] 

## Methods

### NewCollectionResponseWithTotalHubDbTableV3ForwardPaging

`func NewCollectionResponseWithTotalHubDbTableV3ForwardPaging(total int32, results []HubDbTableV3, ) *CollectionResponseWithTotalHubDbTableV3ForwardPaging`

NewCollectionResponseWithTotalHubDbTableV3ForwardPaging instantiates a new CollectionResponseWithTotalHubDbTableV3ForwardPaging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseWithTotalHubDbTableV3ForwardPagingWithDefaults

`func NewCollectionResponseWithTotalHubDbTableV3ForwardPagingWithDefaults() *CollectionResponseWithTotalHubDbTableV3ForwardPaging`

NewCollectionResponseWithTotalHubDbTableV3ForwardPagingWithDefaults instantiates a new CollectionResponseWithTotalHubDbTableV3ForwardPaging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *CollectionResponseWithTotalHubDbTableV3ForwardPaging) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CollectionResponseWithTotalHubDbTableV3ForwardPaging) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CollectionResponseWithTotalHubDbTableV3ForwardPaging) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetResults

`func (o *CollectionResponseWithTotalHubDbTableV3ForwardPaging) GetResults() []HubDbTableV3`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponseWithTotalHubDbTableV3ForwardPaging) GetResultsOk() (*[]HubDbTableV3, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponseWithTotalHubDbTableV3ForwardPaging) SetResults(v []HubDbTableV3)`

SetResults sets Results field to given value.


### GetPaging

`func (o *CollectionResponseWithTotalHubDbTableV3ForwardPaging) GetPaging() ForwardPaging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponseWithTotalHubDbTableV3ForwardPaging) GetPagingOk() (*ForwardPaging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponseWithTotalHubDbTableV3ForwardPaging) SetPaging(v ForwardPaging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponseWithTotalHubDbTableV3ForwardPaging) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


