# CollectionResponseWithTotalSimplePublicObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 
**Results** | [**[]SimplePublicObject**](SimplePublicObject.md) | An array of SimplePublicObject items, each representing an individual CRM object. | 
**Total** | **int32** | The number of available results | 

## Methods

### NewCollectionResponseWithTotalSimplePublicObject

`func NewCollectionResponseWithTotalSimplePublicObject(results []SimplePublicObject, total int32, ) *CollectionResponseWithTotalSimplePublicObject`

NewCollectionResponseWithTotalSimplePublicObject instantiates a new CollectionResponseWithTotalSimplePublicObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseWithTotalSimplePublicObjectWithDefaults

`func NewCollectionResponseWithTotalSimplePublicObjectWithDefaults() *CollectionResponseWithTotalSimplePublicObject`

NewCollectionResponseWithTotalSimplePublicObjectWithDefaults instantiates a new CollectionResponseWithTotalSimplePublicObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaging

`func (o *CollectionResponseWithTotalSimplePublicObject) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponseWithTotalSimplePublicObject) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponseWithTotalSimplePublicObject) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponseWithTotalSimplePublicObject) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *CollectionResponseWithTotalSimplePublicObject) GetResults() []SimplePublicObject`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponseWithTotalSimplePublicObject) GetResultsOk() (*[]SimplePublicObject, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponseWithTotalSimplePublicObject) SetResults(v []SimplePublicObject)`

SetResults sets Results field to given value.


### GetTotal

`func (o *CollectionResponseWithTotalSimplePublicObject) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CollectionResponseWithTotalSimplePublicObject) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CollectionResponseWithTotalSimplePublicObject) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


